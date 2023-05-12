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
	store := ctx.KVStore(k.storeKey)
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

func (k Keeper) GetPositionByDenomAndSender(ctx sdk.Context, denom string, sender string) (position types.Position, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixPosition)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	found = false
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Position
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.GetSender() == sender && val.GetDenom() == denom {
			position = val
			found = true
			break
		}
	}

	return position, found
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
		panic("loan not found")
	}

	position, found := k.GetPositionByLoanId(ctx, loanId)

	if !found {
		errorsmod.Wrap(types.ErrPositionNotFound, fmt.Sprintf("err: position not found by loanId : %v", loanId))
		panic("position not found")
	}

	performanceFeeRate, _ := sdk.NewDecFromStr(strconv.FormatUint(k.GetParam(ctx, types.KeyLiquidationPerformanceFee), 10))
	performanceFeeRate = performanceFeeRate.Quo(sdk.NewDec(100))
	performanceFee := sdk.NewDecFromInt(position.StTokenAmount).Mul(performanceFeeRate).TruncateInt()
	remainingTotalStAsset := sdk.NewDecFromInt(position.StTokenAmount).Mul(sdk.OneDec().Sub(performanceFeeRate)).TruncateInt()
	// TODO: 2. 유저가 포지션 잡은 총 Asset 의 stTokenAmount 를 가져옴
	k.Logger(ctx).Info(fmt.Sprintf("Liquidated Position, TotalStToken Value :: position.StTokenAmount : %v, PerformanceFee : %v, RemainingTotalStAsset : %v", position.StTokenAmount, performanceFee, remainingTotalStAsset))
	position.StTokenAmount = remainingTotalStAsset
	k.SetPosition(ctx, position)
	liquidationFeeAccount, err := sdk.AccAddressFromBech32(types.LiquidationFeeAccount)
	if err != nil {
		errorsmod.Wrap(types.ErrInvalidAccount, fmt.Sprintf("err: invalid fee account"))
		panic("err: invalid fee account")
	}

	// TODO: 3. Performance Fee 와 Performance Fee 를 제외한 Asset 을 Unstaking 함
	// 수수료 지급 이후에 남은 stAsset 을 exit_msg_undelegate 쪽의 함수를 호출하므로 기존 로직을 탄다.
	hostZone, found := k.GetHostZoneByHostDenom(ctx, position.Denom)

	if !found {
		errorsmod.Wrap(types.ErrHostZoneNotFound, fmt.Sprintf("err: host not found"))
		panic("err: host not found)
	}
	zoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)
	if !found {
		errorsmod.Wrap(types.ErrInvalidAccount, fmt.Sprintf("err: invalid hostzone account"))
		panic("err: invalid hostzone address")
	}
	// 청산 수수료를 Module key.go 에 있는 LiquidationFeeAddress 로 계산된 stToken 을 전송함
	err = k.bankKeeper.SendCoins(ctx, zoneAddress, liquidationFeeAccount, sdk.NewCoins(sdk.NewCoin(types.StAssetDenomFromHostZoneDenom(position.Denom), performanceFee)))
	if err != nil {
		errorsmod.Wrap(types.ErrFailureTransferLiquidationFee, fmt.Sprintf("err: transfer liquidation fee fail"))
		panic("err: transfer liquidation failed")
	}
	//@TODO empty reciever field는 임시 값입니다. 추후에 ExitLeverageStake 커맨드에서 reciever필드를 뺄때 모든 코드에서 unstaking receiver field를 제거해야합니다.
	//@TODO 왜냐하면 levstakeibc에서는 receiver를 따로 지정하지 않고 본래 user acct로 보내주기 때문입니다.
	err = k.UnStakeWithLeverage(ctx, position.Sender, position.Id, hostZone.ChainId, "")

	if err != nil {
		errorsmod.Wrap(types.ErrInvalidUnStakeWithLeverage, fmt.Sprintf("err: faild Unstake with leverage for loanId: %v", loanId))
		panic("err: UnstakeWithLeverage Faield")
	}
	//Get position needs to be recalled. Because UnstakeWithLeverage will change the status of the position
	position, found = k.GetPositionByLoanId(ctx, loanId)
	if !found {
		errorsmod.Wrap(types.ErrPositionNotFound, fmt.Sprintf("err: position not found by loanId : %v", loanId))
		panic("position not found")
	}
	position.Liquidated = true
	k.SetPosition(ctx, position)

	return
}
