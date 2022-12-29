/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { ICAAccount } from "./ica_account";
import { Validator } from "./validator";

export const protobufPackage = "stayking.stakeibc";

/** next id: 19 */
export interface HostZone {
  chainId: string;
  connectionId: string;
  bech32prefix: string;
  transferChannelId: string;
  validators: Validator[];
  blacklistedValidators: Validator[];
  withdrawalAccount: ICAAccount | undefined;
  feeAccount: ICAAccount | undefined;
  delegationAccount: ICAAccount | undefined;
  redemptionAccount:
    | ICAAccount
    | undefined;
  /** ibc denom on stride */
  ibcDenom: string;
  /** native denom on host zone */
  hostDenom: string;
  /**
   * TODO(TEST-68): Should we make this an array and store the last n redemption
   * rates then calculate a TWARR?
   */
  lastRedemptionRate: string;
  redemptionRate: string;
  /** stores how many days we should wait before issuing unbondings */
  unbondingFrequency: number;
  /** TODO(TEST-101) int to dec */
  stakedBal: string;
  address: string;
}

function createBaseHostZone(): HostZone {
  return {
    chainId: "",
    connectionId: "",
    bech32prefix: "",
    transferChannelId: "",
    validators: [],
    blacklistedValidators: [],
    withdrawalAccount: undefined,
    feeAccount: undefined,
    delegationAccount: undefined,
    redemptionAccount: undefined,
    ibcDenom: "",
    hostDenom: "",
    lastRedemptionRate: "",
    redemptionRate: "",
    unbondingFrequency: 0,
    stakedBal: "",
    address: "",
  };
}

export const HostZone = {
  encode(message: HostZone, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.chainId !== "") {
      writer.uint32(10).string(message.chainId);
    }
    if (message.connectionId !== "") {
      writer.uint32(18).string(message.connectionId);
    }
    if (message.bech32prefix !== "") {
      writer.uint32(138).string(message.bech32prefix);
    }
    if (message.transferChannelId !== "") {
      writer.uint32(98).string(message.transferChannelId);
    }
    for (const v of message.validators) {
      Validator.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.blacklistedValidators) {
      Validator.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.withdrawalAccount !== undefined) {
      ICAAccount.encode(message.withdrawalAccount, writer.uint32(42).fork()).ldelim();
    }
    if (message.feeAccount !== undefined) {
      ICAAccount.encode(message.feeAccount, writer.uint32(50).fork()).ldelim();
    }
    if (message.delegationAccount !== undefined) {
      ICAAccount.encode(message.delegationAccount, writer.uint32(58).fork()).ldelim();
    }
    if (message.redemptionAccount !== undefined) {
      ICAAccount.encode(message.redemptionAccount, writer.uint32(130).fork()).ldelim();
    }
    if (message.ibcDenom !== "") {
      writer.uint32(66).string(message.ibcDenom);
    }
    if (message.hostDenom !== "") {
      writer.uint32(74).string(message.hostDenom);
    }
    if (message.lastRedemptionRate !== "") {
      writer.uint32(82).string(message.lastRedemptionRate);
    }
    if (message.redemptionRate !== "") {
      writer.uint32(90).string(message.redemptionRate);
    }
    if (message.unbondingFrequency !== 0) {
      writer.uint32(112).uint64(message.unbondingFrequency);
    }
    if (message.stakedBal !== "") {
      writer.uint32(106).string(message.stakedBal);
    }
    if (message.address !== "") {
      writer.uint32(146).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HostZone {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHostZone();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.chainId = reader.string();
          break;
        case 2:
          message.connectionId = reader.string();
          break;
        case 17:
          message.bech32prefix = reader.string();
          break;
        case 12:
          message.transferChannelId = reader.string();
          break;
        case 3:
          message.validators.push(Validator.decode(reader, reader.uint32()));
          break;
        case 4:
          message.blacklistedValidators.push(Validator.decode(reader, reader.uint32()));
          break;
        case 5:
          message.withdrawalAccount = ICAAccount.decode(reader, reader.uint32());
          break;
        case 6:
          message.feeAccount = ICAAccount.decode(reader, reader.uint32());
          break;
        case 7:
          message.delegationAccount = ICAAccount.decode(reader, reader.uint32());
          break;
        case 16:
          message.redemptionAccount = ICAAccount.decode(reader, reader.uint32());
          break;
        case 8:
          message.ibcDenom = reader.string();
          break;
        case 9:
          message.hostDenom = reader.string();
          break;
        case 10:
          message.lastRedemptionRate = reader.string();
          break;
        case 11:
          message.redemptionRate = reader.string();
          break;
        case 14:
          message.unbondingFrequency = longToNumber(reader.uint64() as Long);
          break;
        case 13:
          message.stakedBal = reader.string();
          break;
        case 18:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): HostZone {
    return {
      chainId: isSet(object.chainId) ? String(object.chainId) : "",
      connectionId: isSet(object.connectionId) ? String(object.connectionId) : "",
      bech32prefix: isSet(object.bech32prefix) ? String(object.bech32prefix) : "",
      transferChannelId: isSet(object.transferChannelId) ? String(object.transferChannelId) : "",
      validators: Array.isArray(object?.validators) ? object.validators.map((e: any) => Validator.fromJSON(e)) : [],
      blacklistedValidators: Array.isArray(object?.blacklistedValidators)
        ? object.blacklistedValidators.map((e: any) => Validator.fromJSON(e))
        : [],
      withdrawalAccount: isSet(object.withdrawalAccount) ? ICAAccount.fromJSON(object.withdrawalAccount) : undefined,
      feeAccount: isSet(object.feeAccount) ? ICAAccount.fromJSON(object.feeAccount) : undefined,
      delegationAccount: isSet(object.delegationAccount) ? ICAAccount.fromJSON(object.delegationAccount) : undefined,
      redemptionAccount: isSet(object.redemptionAccount) ? ICAAccount.fromJSON(object.redemptionAccount) : undefined,
      ibcDenom: isSet(object.ibcDenom) ? String(object.ibcDenom) : "",
      hostDenom: isSet(object.hostDenom) ? String(object.hostDenom) : "",
      lastRedemptionRate: isSet(object.lastRedemptionRate) ? String(object.lastRedemptionRate) : "",
      redemptionRate: isSet(object.redemptionRate) ? String(object.redemptionRate) : "",
      unbondingFrequency: isSet(object.unbondingFrequency) ? Number(object.unbondingFrequency) : 0,
      stakedBal: isSet(object.stakedBal) ? String(object.stakedBal) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: HostZone): unknown {
    const obj: any = {};
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.connectionId !== undefined && (obj.connectionId = message.connectionId);
    message.bech32prefix !== undefined && (obj.bech32prefix = message.bech32prefix);
    message.transferChannelId !== undefined && (obj.transferChannelId = message.transferChannelId);
    if (message.validators) {
      obj.validators = message.validators.map((e) => e ? Validator.toJSON(e) : undefined);
    } else {
      obj.validators = [];
    }
    if (message.blacklistedValidators) {
      obj.blacklistedValidators = message.blacklistedValidators.map((e) => e ? Validator.toJSON(e) : undefined);
    } else {
      obj.blacklistedValidators = [];
    }
    message.withdrawalAccount !== undefined
      && (obj.withdrawalAccount = message.withdrawalAccount ? ICAAccount.toJSON(message.withdrawalAccount) : undefined);
    message.feeAccount !== undefined
      && (obj.feeAccount = message.feeAccount ? ICAAccount.toJSON(message.feeAccount) : undefined);
    message.delegationAccount !== undefined
      && (obj.delegationAccount = message.delegationAccount ? ICAAccount.toJSON(message.delegationAccount) : undefined);
    message.redemptionAccount !== undefined
      && (obj.redemptionAccount = message.redemptionAccount ? ICAAccount.toJSON(message.redemptionAccount) : undefined);
    message.ibcDenom !== undefined && (obj.ibcDenom = message.ibcDenom);
    message.hostDenom !== undefined && (obj.hostDenom = message.hostDenom);
    message.lastRedemptionRate !== undefined && (obj.lastRedemptionRate = message.lastRedemptionRate);
    message.redemptionRate !== undefined && (obj.redemptionRate = message.redemptionRate);
    message.unbondingFrequency !== undefined && (obj.unbondingFrequency = Math.round(message.unbondingFrequency));
    message.stakedBal !== undefined && (obj.stakedBal = message.stakedBal);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<HostZone>, I>>(object: I): HostZone {
    const message = createBaseHostZone();
    message.chainId = object.chainId ?? "";
    message.connectionId = object.connectionId ?? "";
    message.bech32prefix = object.bech32prefix ?? "";
    message.transferChannelId = object.transferChannelId ?? "";
    message.validators = object.validators?.map((e) => Validator.fromPartial(e)) || [];
    message.blacklistedValidators = object.blacklistedValidators?.map((e) => Validator.fromPartial(e)) || [];
    message.withdrawalAccount = (object.withdrawalAccount !== undefined && object.withdrawalAccount !== null)
      ? ICAAccount.fromPartial(object.withdrawalAccount)
      : undefined;
    message.feeAccount = (object.feeAccount !== undefined && object.feeAccount !== null)
      ? ICAAccount.fromPartial(object.feeAccount)
      : undefined;
    message.delegationAccount = (object.delegationAccount !== undefined && object.delegationAccount !== null)
      ? ICAAccount.fromPartial(object.delegationAccount)
      : undefined;
    message.redemptionAccount = (object.redemptionAccount !== undefined && object.redemptionAccount !== null)
      ? ICAAccount.fromPartial(object.redemptionAccount)
      : undefined;
    message.ibcDenom = object.ibcDenom ?? "";
    message.hostDenom = object.hostDenom ?? "";
    message.lastRedemptionRate = object.lastRedemptionRate ?? "";
    message.redemptionRate = object.redemptionRate ?? "";
    message.unbondingFrequency = object.unbondingFrequency ?? 0;
    message.stakedBal = object.stakedBal ?? "";
    message.address = object.address ?? "";
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
