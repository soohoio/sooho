/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stayking.stakeibc";

export interface EpochTracker {
  epochIdentifier: string;
  epochNumber: number;
  nextEpochStartTime: number;
  duration: number;
}

function createBaseEpochTracker(): EpochTracker {
  return { epochIdentifier: "", epochNumber: 0, nextEpochStartTime: 0, duration: 0 };
}

export const EpochTracker = {
  encode(message: EpochTracker, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.epochIdentifier !== "") {
      writer.uint32(10).string(message.epochIdentifier);
    }
    if (message.epochNumber !== 0) {
      writer.uint32(16).uint64(message.epochNumber);
    }
    if (message.nextEpochStartTime !== 0) {
      writer.uint32(24).uint64(message.nextEpochStartTime);
    }
    if (message.duration !== 0) {
      writer.uint32(32).uint64(message.duration);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EpochTracker {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEpochTracker();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.epochIdentifier = reader.string();
          break;
        case 2:
          message.epochNumber = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.nextEpochStartTime = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.duration = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EpochTracker {
    return {
      epochIdentifier: isSet(object.epochIdentifier) ? String(object.epochIdentifier) : "",
      epochNumber: isSet(object.epochNumber) ? Number(object.epochNumber) : 0,
      nextEpochStartTime: isSet(object.nextEpochStartTime) ? Number(object.nextEpochStartTime) : 0,
      duration: isSet(object.duration) ? Number(object.duration) : 0,
    };
  },

  toJSON(message: EpochTracker): unknown {
    const obj: any = {};
    message.epochIdentifier !== undefined && (obj.epochIdentifier = message.epochIdentifier);
    message.epochNumber !== undefined && (obj.epochNumber = Math.round(message.epochNumber));
    message.nextEpochStartTime !== undefined && (obj.nextEpochStartTime = Math.round(message.nextEpochStartTime));
    message.duration !== undefined && (obj.duration = Math.round(message.duration));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EpochTracker>, I>>(object: I): EpochTracker {
    const message = createBaseEpochTracker();
    message.epochIdentifier = object.epochIdentifier ?? "";
    message.epochNumber = object.epochNumber ?? 0;
    message.nextEpochStartTime = object.nextEpochStartTime ?? 0;
    message.duration = object.duration ?? 0;
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
