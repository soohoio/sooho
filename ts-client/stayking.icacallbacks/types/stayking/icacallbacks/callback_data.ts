/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stayking.icacallbacks";

export interface CallbackData {
  callbackKey: string;
  portId: string;
  channelId: string;
  sequence: number;
  callbackId: string;
  callbackArgs: Uint8Array;
}

function createBaseCallbackData(): CallbackData {
  return { callbackKey: "", portId: "", channelId: "", sequence: 0, callbackId: "", callbackArgs: new Uint8Array() };
}

export const CallbackData = {
  encode(message: CallbackData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.callbackKey !== "") {
      writer.uint32(10).string(message.callbackKey);
    }
    if (message.portId !== "") {
      writer.uint32(18).string(message.portId);
    }
    if (message.channelId !== "") {
      writer.uint32(26).string(message.channelId);
    }
    if (message.sequence !== 0) {
      writer.uint32(32).uint64(message.sequence);
    }
    if (message.callbackId !== "") {
      writer.uint32(42).string(message.callbackId);
    }
    if (message.callbackArgs.length !== 0) {
      writer.uint32(50).bytes(message.callbackArgs);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CallbackData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCallbackData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.callbackKey = reader.string();
          break;
        case 2:
          message.portId = reader.string();
          break;
        case 3:
          message.channelId = reader.string();
          break;
        case 4:
          message.sequence = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.callbackId = reader.string();
          break;
        case 6:
          message.callbackArgs = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CallbackData {
    return {
      callbackKey: isSet(object.callbackKey) ? String(object.callbackKey) : "",
      portId: isSet(object.portId) ? String(object.portId) : "",
      channelId: isSet(object.channelId) ? String(object.channelId) : "",
      sequence: isSet(object.sequence) ? Number(object.sequence) : 0,
      callbackId: isSet(object.callbackId) ? String(object.callbackId) : "",
      callbackArgs: isSet(object.callbackArgs) ? bytesFromBase64(object.callbackArgs) : new Uint8Array(),
    };
  },

  toJSON(message: CallbackData): unknown {
    const obj: any = {};
    message.callbackKey !== undefined && (obj.callbackKey = message.callbackKey);
    message.portId !== undefined && (obj.portId = message.portId);
    message.channelId !== undefined && (obj.channelId = message.channelId);
    message.sequence !== undefined && (obj.sequence = Math.round(message.sequence));
    message.callbackId !== undefined && (obj.callbackId = message.callbackId);
    message.callbackArgs !== undefined
      && (obj.callbackArgs = base64FromBytes(
        message.callbackArgs !== undefined ? message.callbackArgs : new Uint8Array(),
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CallbackData>, I>>(object: I): CallbackData {
    const message = createBaseCallbackData();
    message.callbackKey = object.callbackKey ?? "";
    message.portId = object.portId ?? "";
    message.channelId = object.channelId ?? "";
    message.sequence = object.sequence ?? 0;
    message.callbackId = object.callbackId ?? "";
    message.callbackArgs = object.callbackArgs ?? new Uint8Array();
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

function bytesFromBase64(b64: string): Uint8Array {
  if (globalThis.Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (globalThis.Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
}

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
