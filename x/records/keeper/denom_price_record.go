package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/records/types"
)

// TODO: 현재는 Denom Price 에 대해 Overwrite 하는 방식으로 최근 가격 기준만 들고 있음
// TODO: 향후 이상치 데이터 없애고 값을 평균하여 줄 버퍼 형태로 갈 것인지 고민해야 함
func (k Keeper) SetDenomPriceRecord(ctx sdk.Context, denomPriceRecord types.DenomPriceRecord) {
	fmt.Println("DEBUG: Records module StoreKey in Keeper:", k.storeKey)
	k.Logger(ctx).Info("[SetDenomPrice Debug1] ")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomPriceRecordKey))
	k.Logger(ctx).Info("[SetDenomPrice Debug2] ")
	b := k.Cdc.MustMarshal(&denomPriceRecord)
	k.Logger(ctx).Info("[SetDenomPrice Debug3] ")
	store.Set(GetDenomPriceRecordIDBytes(denomPriceRecord.GetBaseDenom()+"-"+denomPriceRecord.GetTargetDenom()), b)
	k.Logger(ctx).Info("[SetDenomPrice Debug4] ")
}

func (k Keeper) GetDenomPriceRecord(ctx sdk.Context, denomId string) (val types.DenomPriceRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomPriceRecordKey))
	b := store.Get(GetDenomPriceRecordIDBytes(denomId))
	if b == nil {
		return val, false
	}
	k.Cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllDenomPriceRecord(ctx sdk.Context) (list []types.DenomPriceRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomPriceRecordKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var elem types.DenomPriceRecord
		k.Cdc.MustUnmarshal(iterator.Value(), &elem)
		list = append(list, elem)
	}

	return
}

func GetDenomPriceRecordIDBytes(denomId string) []byte {
	return []byte(denomId)
}
