package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"encoding/binary"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"strconv"
)

func (k Keeper) SetPosition(ctx sdk.Context, posistion types.Position) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&posistion)
	store.Set(types.GetPositionKey(posistion.Id), bz)
}

func (k Keeper) GetPosition(ctx sdk.Context, id uint64) (types.Position, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPositionKey(id))
	if bz == nil {
		return types.Position{}, false
	}
	var position types.Position
	k.cdc.MustUnmarshal(bz, &position)
	return position, true
}

// RemovePosition removes a position from the store
func (k Keeper) RemovePosition(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixPosition)
	store.Delete(types.GetPositionKey(id))
}

func (k Keeper) GetNextPositionID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixNextPositionID)
	return binary.LittleEndian.Uint64(bz)
}

func (k Keeper) SetNextPositionID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, id)
	store.Set(types.KeyPrefixNextPositionID, bz)
}

// GetAllPosition returns all Position
func (k Keeper) GetAllPosition(ctx sdk.Context) (list []types.Position) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixPosition)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Position
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

func (k Keeper) GetAllPositionByPage(ctx sdk.Context, page *query.PageRequest) (positions []types.Position, pageRes *query.PageResponse, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixPosition)

	pageRes, err = query.Paginate(store, page, func(key []byte, value []byte) error {
		var position types.Position
		if err := k.cdc.Unmarshal(value, &position); err != nil {
			return err
		}
		positions = append(positions, position)
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return positions, pageRes, nil
}

func (k Keeper) GetPositionListBySender(ctx sdk.Context, sender string) (positions []types.Position) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixPosition)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Position
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.GetSender() == sender {
			positions = append(positions, val)
		}
	}

	return positions
}

func (k Keeper) GetPositionByLoanId(ctx sdk.Context, loanId uint64) (position types.Position, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixPosition)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Position
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.GetLoanId() == loanId {
			position = val
			break
		}
	}

	return position, true
}

func (k Keeper) Liquidate(ctx sdk.Context, loanId uint64) {
	k.Logger(ctx).Info(fmt.Sprintf("Liquidate PositionId : %v", loanId))
	// TODO: 1. Loan 정보를 x/lendingpool 로 부터 불러옴
	_, found := k.LendingPoolKeeper.GetPool(ctx, loanId)

	if !found {
		errorsmod.Wrap(types.ErrPositionNotFound, fmt.Sprintf("err: position not found by loanId : %v", loanId))
		return
	}

	position, found := k.GetPositionByLoanId(ctx, loanId)

	if !found {
		errorsmod.Wrap(types.ErrPositionNotFound, fmt.Sprintf("err: position not found by loanId : %v", loanId))
		return
	}

	performanceFeeRate, _ := sdk.NewDecFromStr(strconv.FormatUint(k.GetParam(ctx, types.KeyLiquidationPerformanceFee), 10))
	performanceFeeRate = performanceFeeRate.Quo(sdk.NewDec(100))
	performanceFee := sdk.NewDecFromInt(position.StTokenAmount).Mul(performanceFeeRate).TruncateInt()
	remainingTotalStAsset := sdk.NewDecFromInt(position.StTokenAmount).Mul(sdk.OneDec().Sub(performanceFeeRate)).TruncateInt()
	// TODO: 2. 유저가 포지션 잡은 총 Asset 의 stTokenAmount 를 가져옴
	k.Logger(ctx).Info(fmt.Sprintf("Liquidated Position, TotalStToken Value :: position.StTokenAmount : %v, PerformanceFee : %v, RemainingTotalStAsset : %v", position.StTokenAmount, performanceFee, remainingTotalStAsset))

	// 청산 수수료를 Module key.go 에 있는 FeeAddress 로 계산된 stToken 을 전송함
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(types.FeeAccount), sdk.NewCoins(sdk.NewCoin(types.StAssetDenomFromHostZoneDenom(position.Denom), performanceFee)))
	// TODO: 3. Performance Fee 와 Performance Fee 를 제외한 Asset 을 Unstaking 함
	// 수수료 지급 이후에 남은 stAsset 을 exit_msg_undelegate 쪽의 함수를 호출하므로 기존 로직을 탄다.
	hostZone, found := k.GetHostZoneByHostDenom(ctx, position.Denom)

	if !found {
		errorsmod.Wrap(types.ErrHostZoneNotFound, fmt.Sprintf("err host not found"))
		return
	}

	err := k.UnStakeWithLeverage(ctx, hostZone.Address, position.Id, hostZone.ChainId, position.Receiver)

	if err != nil {
		errorsmod.Wrap(types.ErrInvalidUnStakeWithLeverage, fmt.Sprintf("err: position not found by loanId : %v", loanId))
		return
	}
	position.Liquidated = true
	k.SetPosition(ctx, position)

	return
}
