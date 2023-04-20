package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"encoding/binary"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k Keeper) Liquidate(ctx sdk.Context, loanId uint64) {
	k.Logger(ctx).Info(fmt.Sprintf("Liquidate PositionId : %v", loanId))
	// TODO: 1. Loan 정보를 x/lendingpool 로 부터 불러옴
	loan, found := k.LendingPoolKeeper.GetPool(ctx, loanId)

	if !found {
		errorsmod.Wrap(types.ErrPositionNotFound, fmt.Sprintf("err: position not found by loanId : %v", loanId))
	}

	position, found := k.GetPositionByLoanId(ctx, loanId)

	if !found {
		errorsmod.Wrap(types.ErrPositionNotFound, fmt.Sprintf("err: position not found by loanId : %v", loanId))
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

	// TODO: 4. Position 상태 변경하기
	position.Liquidated = true
	k.SetPosition(ctx, position)

	return
}
