/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "stayking.stakeibc";

/** ---------------------- Delegation Callbacks ---------------------- // */
export interface SplitDelegation {
  validator: string;
  amount: string;
}

export interface DelegateCallback {
  hostZoneId: string;
  depositRecordId: number;
  splitDelegations: SplitDelegation[];
}

export interface ClaimCallback {
  userRedemptionRecordId: string;
  chainId: string;
  epochNumber: number;
}

/** ---------------------- Reinvest Callback ---------------------- // */
export interface ReinvestCallback {
  reinvestAmount: Coin | undefined;
  hostZoneId: string;
}

/** ---------------------- Undelegation Callbacks ---------------------- // */
export interface UndelegateCallback {
  hostZoneId: string;
  splitDelegations: SplitDelegation[];
  epochUnbondingRecordIds: number[];
}

/** ---------------------- Redemption Callbacks ---------------------- // */
export interface RedemptionCallback {
  hostZoneId: string;
  epochUnbondingRecordIds: number[];
}

export interface Rebalancing {
  srcValidator: string;
  dstValidator: string;
  amt: string;
}

export interface RebalanceCallback {
  hostZoneId: string;
  rebalancings: Rebalancing[];
}

function createBaseSplitDelegation(): SplitDelegation {
  return { validator: "", amount: "" };
}

export const SplitDelegation = {
  encode(message: SplitDelegation, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.validator !== "") {
      writer.uint32(10).string(message.validator);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SplitDelegation {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSplitDelegation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validator = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SplitDelegation {
    return {
      validator: isSet(object.validator) ? String(object.validator) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
    };
  },

  toJSON(message: SplitDelegation): unknown {
    const obj: any = {};
    message.validator !== undefined && (obj.validator = message.validator);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SplitDelegation>, I>>(object: I): SplitDelegation {
    const message = createBaseSplitDelegation();
    message.validator = object.validator ?? "";
    message.amount = object.amount ?? "";
    return message;
  },
};

function createBaseDelegateCallback(): DelegateCallback {
  return { hostZoneId: "", depositRecordId: 0, splitDelegations: [] };
}

export const DelegateCallback = {
  encode(message: DelegateCallback, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.hostZoneId !== "") {
      writer.uint32(10).string(message.hostZoneId);
    }
    if (message.depositRecordId !== 0) {
      writer.uint32(16).uint64(message.depositRecordId);
    }
    for (const v of message.splitDelegations) {
      SplitDelegation.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DelegateCallback {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDelegateCallback();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.hostZoneId = reader.string();
          break;
        case 2:
          message.depositRecordId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.splitDelegations.push(SplitDelegation.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DelegateCallback {
    return {
      hostZoneId: isSet(object.hostZoneId) ? String(object.hostZoneId) : "",
      depositRecordId: isSet(object.depositRecordId) ? Number(object.depositRecordId) : 0,
      splitDelegations: Array.isArray(object?.splitDelegations)
        ? object.splitDelegations.map((e: any) => SplitDelegation.fromJSON(e))
        : [],
    };
  },

  toJSON(message: DelegateCallback): unknown {
    const obj: any = {};
    message.hostZoneId !== undefined && (obj.hostZoneId = message.hostZoneId);
    message.depositRecordId !== undefined && (obj.depositRecordId = Math.round(message.depositRecordId));
    if (message.splitDelegations) {
      obj.splitDelegations = message.splitDelegations.map((e) => e ? SplitDelegation.toJSON(e) : undefined);
    } else {
      obj.splitDelegations = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DelegateCallback>, I>>(object: I): DelegateCallback {
    const message = createBaseDelegateCallback();
    message.hostZoneId = object.hostZoneId ?? "";
    message.depositRecordId = object.depositRecordId ?? 0;
    message.splitDelegations = object.splitDelegations?.map((e) => SplitDelegation.fromPartial(e)) || [];
    return message;
  },
};

function createBaseClaimCallback(): ClaimCallback {
  return { userRedemptionRecordId: "", chainId: "", epochNumber: 0 };
}

export const ClaimCallback = {
  encode(message: ClaimCallback, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userRedemptionRecordId !== "") {
      writer.uint32(10).string(message.userRedemptionRecordId);
    }
    if (message.chainId !== "") {
      writer.uint32(18).string(message.chainId);
    }
    if (message.epochNumber !== 0) {
      writer.uint32(24).uint64(message.epochNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ClaimCallback {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseClaimCallback();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userRedemptionRecordId = reader.string();
          break;
        case 2:
          message.chainId = reader.string();
          break;
        case 3:
          message.epochNumber = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ClaimCallback {
    return {
      userRedemptionRecordId: isSet(object.userRedemptionRecordId) ? String(object.userRedemptionRecordId) : "",
      chainId: isSet(object.chainId) ? String(object.chainId) : "",
      epochNumber: isSet(object.epochNumber) ? Number(object.epochNumber) : 0,
    };
  },

  toJSON(message: ClaimCallback): unknown {
    const obj: any = {};
    message.userRedemptionRecordId !== undefined && (obj.userRedemptionRecordId = message.userRedemptionRecordId);
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.epochNumber !== undefined && (obj.epochNumber = Math.round(message.epochNumber));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ClaimCallback>, I>>(object: I): ClaimCallback {
    const message = createBaseClaimCallback();
    message.userRedemptionRecordId = object.userRedemptionRecordId ?? "";
    message.chainId = object.chainId ?? "";
    message.epochNumber = object.epochNumber ?? 0;
    return message;
  },
};

function createBaseReinvestCallback(): ReinvestCallback {
  return { reinvestAmount: undefined, hostZoneId: "" };
}

export const ReinvestCallback = {
  encode(message: ReinvestCallback, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.reinvestAmount !== undefined) {
      Coin.encode(message.reinvestAmount, writer.uint32(10).fork()).ldelim();
    }
    if (message.hostZoneId !== "") {
      writer.uint32(26).string(message.hostZoneId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReinvestCallback {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReinvestCallback();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.reinvestAmount = Coin.decode(reader, reader.uint32());
          break;
        case 3:
          message.hostZoneId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ReinvestCallback {
    return {
      reinvestAmount: isSet(object.reinvestAmount) ? Coin.fromJSON(object.reinvestAmount) : undefined,
      hostZoneId: isSet(object.hostZoneId) ? String(object.hostZoneId) : "",
    };
  },

  toJSON(message: ReinvestCallback): unknown {
    const obj: any = {};
    message.reinvestAmount !== undefined
      && (obj.reinvestAmount = message.reinvestAmount ? Coin.toJSON(message.reinvestAmount) : undefined);
    message.hostZoneId !== undefined && (obj.hostZoneId = message.hostZoneId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ReinvestCallback>, I>>(object: I): ReinvestCallback {
    const message = createBaseReinvestCallback();
    message.reinvestAmount = (object.reinvestAmount !== undefined && object.reinvestAmount !== null)
      ? Coin.fromPartial(object.reinvestAmount)
      : undefined;
    message.hostZoneId = object.hostZoneId ?? "";
    return message;
  },
};

function createBaseUndelegateCallback(): UndelegateCallback {
  return { hostZoneId: "", splitDelegations: [], epochUnbondingRecordIds: [] };
}

export const UndelegateCallback = {
  encode(message: UndelegateCallback, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.hostZoneId !== "") {
      writer.uint32(10).string(message.hostZoneId);
    }
    for (const v of message.splitDelegations) {
      SplitDelegation.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    writer.uint32(26).fork();
    for (const v of message.epochUnbondingRecordIds) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UndelegateCallback {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUndelegateCallback();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.hostZoneId = reader.string();
          break;
        case 2:
          message.splitDelegations.push(SplitDelegation.decode(reader, reader.uint32()));
          break;
        case 3:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.epochUnbondingRecordIds.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.epochUnbondingRecordIds.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UndelegateCallback {
    return {
      hostZoneId: isSet(object.hostZoneId) ? String(object.hostZoneId) : "",
      splitDelegations: Array.isArray(object?.splitDelegations)
        ? object.splitDelegations.map((e: any) => SplitDelegation.fromJSON(e))
        : [],
      epochUnbondingRecordIds: Array.isArray(object?.epochUnbondingRecordIds)
        ? object.epochUnbondingRecordIds.map((e: any) => Number(e))
        : [],
    };
  },

  toJSON(message: UndelegateCallback): unknown {
    const obj: any = {};
    message.hostZoneId !== undefined && (obj.hostZoneId = message.hostZoneId);
    if (message.splitDelegations) {
      obj.splitDelegations = message.splitDelegations.map((e) => e ? SplitDelegation.toJSON(e) : undefined);
    } else {
      obj.splitDelegations = [];
    }
    if (message.epochUnbondingRecordIds) {
      obj.epochUnbondingRecordIds = message.epochUnbondingRecordIds.map((e) => Math.round(e));
    } else {
      obj.epochUnbondingRecordIds = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UndelegateCallback>, I>>(object: I): UndelegateCallback {
    const message = createBaseUndelegateCallback();
    message.hostZoneId = object.hostZoneId ?? "";
    message.splitDelegations = object.splitDelegations?.map((e) => SplitDelegation.fromPartial(e)) || [];
    message.epochUnbondingRecordIds = object.epochUnbondingRecordIds?.map((e) => e) || [];
    return message;
  },
};

function createBaseRedemptionCallback(): RedemptionCallback {
  return { hostZoneId: "", epochUnbondingRecordIds: [] };
}

export const RedemptionCallback = {
  encode(message: RedemptionCallback, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.hostZoneId !== "") {
      writer.uint32(10).string(message.hostZoneId);
    }
    writer.uint32(18).fork();
    for (const v of message.epochUnbondingRecordIds) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RedemptionCallback {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRedemptionCallback();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.hostZoneId = reader.string();
          break;
        case 2:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.epochUnbondingRecordIds.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.epochUnbondingRecordIds.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RedemptionCallback {
    return {
      hostZoneId: isSet(object.hostZoneId) ? String(object.hostZoneId) : "",
      epochUnbondingRecordIds: Array.isArray(object?.epochUnbondingRecordIds)
        ? object.epochUnbondingRecordIds.map((e: any) => Number(e))
        : [],
    };
  },

  toJSON(message: RedemptionCallback): unknown {
    const obj: any = {};
    message.hostZoneId !== undefined && (obj.hostZoneId = message.hostZoneId);
    if (message.epochUnbondingRecordIds) {
      obj.epochUnbondingRecordIds = message.epochUnbondingRecordIds.map((e) => Math.round(e));
    } else {
      obj.epochUnbondingRecordIds = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RedemptionCallback>, I>>(object: I): RedemptionCallback {
    const message = createBaseRedemptionCallback();
    message.hostZoneId = object.hostZoneId ?? "";
    message.epochUnbondingRecordIds = object.epochUnbondingRecordIds?.map((e) => e) || [];
    return message;
  },
};

function createBaseRebalancing(): Rebalancing {
  return { srcValidator: "", dstValidator: "", amt: "" };
}

export const Rebalancing = {
  encode(message: Rebalancing, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.srcValidator !== "") {
      writer.uint32(10).string(message.srcValidator);
    }
    if (message.dstValidator !== "") {
      writer.uint32(18).string(message.dstValidator);
    }
    if (message.amt !== "") {
      writer.uint32(26).string(message.amt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Rebalancing {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRebalancing();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.srcValidator = reader.string();
          break;
        case 2:
          message.dstValidator = reader.string();
          break;
        case 3:
          message.amt = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Rebalancing {
    return {
      srcValidator: isSet(object.srcValidator) ? String(object.srcValidator) : "",
      dstValidator: isSet(object.dstValidator) ? String(object.dstValidator) : "",
      amt: isSet(object.amt) ? String(object.amt) : "",
    };
  },

  toJSON(message: Rebalancing): unknown {
    const obj: any = {};
    message.srcValidator !== undefined && (obj.srcValidator = message.srcValidator);
    message.dstValidator !== undefined && (obj.dstValidator = message.dstValidator);
    message.amt !== undefined && (obj.amt = message.amt);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Rebalancing>, I>>(object: I): Rebalancing {
    const message = createBaseRebalancing();
    message.srcValidator = object.srcValidator ?? "";
    message.dstValidator = object.dstValidator ?? "";
    message.amt = object.amt ?? "";
    return message;
  },
};

function createBaseRebalanceCallback(): RebalanceCallback {
  return { hostZoneId: "", rebalancings: [] };
}

export const RebalanceCallback = {
  encode(message: RebalanceCallback, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.hostZoneId !== "") {
      writer.uint32(10).string(message.hostZoneId);
    }
    for (const v of message.rebalancings) {
      Rebalancing.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RebalanceCallback {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRebalanceCallback();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.hostZoneId = reader.string();
          break;
        case 2:
          message.rebalancings.push(Rebalancing.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RebalanceCallback {
    return {
      hostZoneId: isSet(object.hostZoneId) ? String(object.hostZoneId) : "",
      rebalancings: Array.isArray(object?.rebalancings)
        ? object.rebalancings.map((e: any) => Rebalancing.fromJSON(e))
        : [],
    };
  },

  toJSON(message: RebalanceCallback): unknown {
    const obj: any = {};
    message.hostZoneId !== undefined && (obj.hostZoneId = message.hostZoneId);
    if (message.rebalancings) {
      obj.rebalancings = message.rebalancings.map((e) => e ? Rebalancing.toJSON(e) : undefined);
    } else {
      obj.rebalancings = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RebalanceCallback>, I>>(object: I): RebalanceCallback {
    const message = createBaseRebalanceCallback();
    message.hostZoneId = object.hostZoneId ?? "";
    message.rebalancings = object.rebalancings?.map((e) => Rebalancing.fromPartial(e)) || [];
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
