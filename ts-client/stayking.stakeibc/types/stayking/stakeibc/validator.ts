/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stayking.stakeibc";

export interface ValidatorExchangeRate {
  internalTokensToSharesRate: string;
  epochNumber: number;
}

export interface Validator {
  name: string;
  address: string;
  status: Validator_ValidatorStatus;
  commissionRate: number;
  delegationAmt: string;
  weight: number;
  internalExchangeRate: ValidatorExchangeRate | undefined;
}

export enum Validator_ValidatorStatus {
  ACTIVE = 0,
  INACTIVE = 1,
  UNRECOGNIZED = -1,
}

export function validator_ValidatorStatusFromJSON(object: any): Validator_ValidatorStatus {
  switch (object) {
    case 0:
    case "ACTIVE":
      return Validator_ValidatorStatus.ACTIVE;
    case 1:
    case "INACTIVE":
      return Validator_ValidatorStatus.INACTIVE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Validator_ValidatorStatus.UNRECOGNIZED;
  }
}

export function validator_ValidatorStatusToJSON(object: Validator_ValidatorStatus): string {
  switch (object) {
    case Validator_ValidatorStatus.ACTIVE:
      return "ACTIVE";
    case Validator_ValidatorStatus.INACTIVE:
      return "INACTIVE";
    case Validator_ValidatorStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

function createBaseValidatorExchangeRate(): ValidatorExchangeRate {
  return { internalTokensToSharesRate: "", epochNumber: 0 };
}

export const ValidatorExchangeRate = {
  encode(message: ValidatorExchangeRate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.internalTokensToSharesRate !== "") {
      writer.uint32(10).string(message.internalTokensToSharesRate);
    }
    if (message.epochNumber !== 0) {
      writer.uint32(16).uint64(message.epochNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ValidatorExchangeRate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseValidatorExchangeRate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.internalTokensToSharesRate = reader.string();
          break;
        case 2:
          message.epochNumber = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ValidatorExchangeRate {
    return {
      internalTokensToSharesRate: isSet(object.internalTokensToSharesRate)
        ? String(object.internalTokensToSharesRate)
        : "",
      epochNumber: isSet(object.epochNumber) ? Number(object.epochNumber) : 0,
    };
  },

  toJSON(message: ValidatorExchangeRate): unknown {
    const obj: any = {};
    message.internalTokensToSharesRate !== undefined
      && (obj.internalTokensToSharesRate = message.internalTokensToSharesRate);
    message.epochNumber !== undefined && (obj.epochNumber = Math.round(message.epochNumber));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ValidatorExchangeRate>, I>>(object: I): ValidatorExchangeRate {
    const message = createBaseValidatorExchangeRate();
    message.internalTokensToSharesRate = object.internalTokensToSharesRate ?? "";
    message.epochNumber = object.epochNumber ?? 0;
    return message;
  },
};

function createBaseValidator(): Validator {
  return {
    name: "",
    address: "",
    status: 0,
    commissionRate: 0,
    delegationAmt: "",
    weight: 0,
    internalExchangeRate: undefined,
  };
}

export const Validator = {
  encode(message: Validator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.status !== 0) {
      writer.uint32(24).int32(message.status);
    }
    if (message.commissionRate !== 0) {
      writer.uint32(32).uint64(message.commissionRate);
    }
    if (message.delegationAmt !== "") {
      writer.uint32(42).string(message.delegationAmt);
    }
    if (message.weight !== 0) {
      writer.uint32(48).uint64(message.weight);
    }
    if (message.internalExchangeRate !== undefined) {
      ValidatorExchangeRate.encode(message.internalExchangeRate, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Validator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseValidator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.status = reader.int32() as any;
          break;
        case 4:
          message.commissionRate = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.delegationAmt = reader.string();
          break;
        case 6:
          message.weight = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.internalExchangeRate = ValidatorExchangeRate.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Validator {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      address: isSet(object.address) ? String(object.address) : "",
      status: isSet(object.status) ? validator_ValidatorStatusFromJSON(object.status) : 0,
      commissionRate: isSet(object.commissionRate) ? Number(object.commissionRate) : 0,
      delegationAmt: isSet(object.delegationAmt) ? String(object.delegationAmt) : "",
      weight: isSet(object.weight) ? Number(object.weight) : 0,
      internalExchangeRate: isSet(object.internalExchangeRate)
        ? ValidatorExchangeRate.fromJSON(object.internalExchangeRate)
        : undefined,
    };
  },

  toJSON(message: Validator): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.address !== undefined && (obj.address = message.address);
    message.status !== undefined && (obj.status = validator_ValidatorStatusToJSON(message.status));
    message.commissionRate !== undefined && (obj.commissionRate = Math.round(message.commissionRate));
    message.delegationAmt !== undefined && (obj.delegationAmt = message.delegationAmt);
    message.weight !== undefined && (obj.weight = Math.round(message.weight));
    message.internalExchangeRate !== undefined && (obj.internalExchangeRate = message.internalExchangeRate
      ? ValidatorExchangeRate.toJSON(message.internalExchangeRate)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Validator>, I>>(object: I): Validator {
    const message = createBaseValidator();
    message.name = object.name ?? "";
    message.address = object.address ?? "";
    message.status = object.status ?? 0;
    message.commissionRate = object.commissionRate ?? 0;
    message.delegationAmt = object.delegationAmt ?? "";
    message.weight = object.weight ?? 0;
    message.internalExchangeRate = (object.internalExchangeRate !== undefined && object.internalExchangeRate !== null)
      ? ValidatorExchangeRate.fromPartial(object.internalExchangeRate)
      : undefined;
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
