/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stayking.records";

/** ---------------------- Transfer Callback ---------------------- // */
export interface TransferCallback {
  depositRecordId: number;
}

function createBaseTransferCallback(): TransferCallback {
  return { depositRecordId: 0 };
}

export const TransferCallback = {
  encode(message: TransferCallback, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.depositRecordId !== 0) {
      writer.uint32(8).uint64(message.depositRecordId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TransferCallback {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransferCallback();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.depositRecordId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TransferCallback {
    return { depositRecordId: isSet(object.depositRecordId) ? Number(object.depositRecordId) : 0 };
  },

  toJSON(message: TransferCallback): unknown {
    const obj: any = {};
    message.depositRecordId !== undefined && (obj.depositRecordId = Math.round(message.depositRecordId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<TransferCallback>, I>>(object: I): TransferCallback {
    const message = createBaseTransferCallback();
    message.depositRecordId = object.depositRecordId ?? 0;
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
