package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	"github.com/soohoio/stayking/v2/x/interchainquery/types"
	levstakeibctype "github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func GenerateQueryHash(connectionId string, chainId string, queryType string, request []byte, module string, callbackId string) string {
	return fmt.Sprintf("%x", crypto.Sha256(append([]byte(module+connectionId+chainId+queryType+callbackId), request...)))
}

func (k Keeper) NewQuery(ctx sdk.Context, module string, callbackId string, chainId string, connectionId string, queryType string, request []byte, ttl uint64) *types.Query {
	return &types.Query{
		Id:           GenerateQueryHash(connectionId, chainId, queryType, request, module, callbackId),
		ConnectionId: connectionId,
		ChainId:      chainId,
		QueryType:    queryType,
		Request:      request,
		CallbackId:   callbackId,
		Ttl:          ttl,
		RequestSent:  false,
	}
}

// GetQuery returns query
func (k Keeper) GetQuery(ctx sdk.Context, id string) (types.Query, bool) {
	query := types.Query{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixQuery)
	bz := store.Get([]byte(id))
	if len(bz) == 0 {
		return query, false
	}
	k.cdc.MustUnmarshal(bz, &query)
	return query, true
}

// SetQuery set query info
func (k Keeper) SetQuery(ctx sdk.Context, query types.Query) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixQuery)
	bz := k.cdc.MustMarshal(&query)
	store.Set([]byte(query.Id), bz)
}

// DeleteQuery delete query info
func (k Keeper) DeleteQuery(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixQuery)
	store.Delete([]byte(id))
}

// IterateQueries iterate through queries
func (k Keeper) IterateQueries(ctx sdk.Context, fn func(index int64, queryInfo types.Query) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixQuery)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		query := types.Query{}
		k.cdc.MustUnmarshal(iterator.Value(), &query)
		stop := fn(i, query)

		if stop {
			break
		}
		i++
	}
}

// AllQueries returns every queryInfo in the store
func (k Keeper) AllQueries(ctx sdk.Context) []types.Query {
	queries := []types.Query{}
	k.IterateQueries(ctx, func(_ int64, queryInfo types.Query) (stop bool) {
		queries = append(queries, queryInfo)
		return false
	})
	return queries
}

// Helper function to unmarshal a Balance query response across SDK versions
// Before SDK v46, the query response returned a sdk.Coin type. SDK v46 returns an int type
// https://github.com/cosmos/cosmos-sdk/pull/9832
func UnmarshalAmountFromBalanceQuery(cdc codec.BinaryCodec, queryResponseBz []byte) (amount sdkmath.Int, err error) {
	// An nil should not be possible, exit immediately if it occurs
	if queryResponseBz == nil {
		return sdkmath.Int{}, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "query response is nil")
	}

	// If the query response is empty, that means the account was never registed (and thus has a 0 balance)
	if len(queryResponseBz) == 0 {
		return sdkmath.ZeroInt(), nil
	}

	// First attempt to unmarshal as an Int (for SDK v46+)
	// If the result was serialized as a `Coin` type, it should contain a string (representing the denom)
	// which will cause the unmarshalling to throw an error
	intError := amount.Unmarshal(queryResponseBz)
	if intError == nil {
		return amount, nil
	}

	// If the Int unmarshaling was unsuccessful, attempt again using a Coin type (for SDK v45 and below)
	// If successful, return the amount field from the coin (if the coin is not nil)
	var coin sdk.Coin
	coinError := cdc.Unmarshal(queryResponseBz, &coin)
	if coinError == nil {
		return coin.Amount, nil
	}

	// If it failed unmarshaling with either data structure, return an error with the failure messages combined
	return sdkmath.Int{}, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
		"unable to unmarshal balance query response %v as sdkmath.Int (err: %s) or sdk.Coin (err: %s)", queryResponseBz, intError.Error(), coinError.Error())
}


// SetQueryRequest saves the query request
func (k Keeper) SetQueryRequest(ctx sdk.Context, packetSequence uint64, req banktypes.QueryAllBalancesRequest) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.QueryRequestStoreKey(packetSequence), k.cdc.MustMarshal(&req))
}
func (k Keeper) SetQuerySwapRequest(ctx sdk.Context, packetSequence uint64, req levstakeibctype.EstimateSwapExactAmountOutRequest) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.QueryRequestStoreKey(packetSequence), k.cdc.MustMarshal(&req))
}

// GetQueryRequest returns the query request by packet sequence
func (k Keeper) GetQueryRequest(ctx sdk.Context, packetSequence uint64) (levstakeibctype.EstimateSwapExactAmountOutRequest, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.QueryRequestStoreKey(packetSequence))
	if bz == nil {
		return levstakeibctype.EstimateSwapExactAmountOutRequest{}, sdkerrors.Wrapf(types.ErrSample,
			"GetQueryRequest: Result for packet sequence %d is not available.", packetSequence,
		)
	}
	var req levstakeibctype.EstimateSwapExactAmountOutRequest
	k.cdc.MustUnmarshal(bz, &req)
	return req, nil
}

// SetQueryResponse saves the query response
func (k Keeper) SetQueryResponse(ctx sdk.Context, packetSequence uint64, resp levstakeibctype.EstimateSwapExactAmountOutResponse) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.QueryResponseStoreKey(packetSequence), k.cdc.MustMarshal(&resp))
}

// GetQueryResponse returns the query response by packet sequence
func (k Keeper) GetQueryResponse(ctx sdk.Context, packetSequence uint64) (levstakeibctype.EstimateSwapExactAmountOutResponse, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.QueryResponseStoreKey(packetSequence))
	if bz == nil {
		return levstakeibctype.EstimateSwapExactAmountOutResponse{}, sdkerrors.Wrapf(types.ErrSample,
			"GetQueryResponse: Result for packet sequence %d is not available.", packetSequence,
		)
	}
	var resp levstakeibctype.EstimateSwapExactAmountOutResponse
	k.cdc.MustUnmarshal(bz, &resp)
	return resp, nil
}

// GetLastQueryPacketSeq return the id from the last query request
func (k Keeper) GetLastQueryPacketSeq(ctx sdk.Context) uint64 {
	bz := ctx.KVStore(k.storeKey).Get(types.KeyPrefix(types.LastQueryPacketSeqKey))
	uintV := gogotypes.UInt64Value{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &uintV)
	return uintV.GetValue()
}

// SetLastQueryPacketSeq saves the id from the last query request
func (k Keeper) SetLastQueryPacketSeq(ctx sdk.Context, packetSequence uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefix(types.LastQueryPacketSeqKey),
		k.cdc.MustMarshalLengthPrefixed(&gogotypes.UInt64Value{Value: packetSequence}))
}
