/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stayking.interchainquery.v1";

export interface Query {
  id: string;
  connectionId: string;
  chainId: string;
  queryType: string;
  request: Uint8Array;
  callbackId: string;
  ttl: number;
  requestSent: boolean;
}

export interface DataPoint {
  id: string;
  remoteHeight: string;
  localHeight: string;
  value: Uint8Array;
}

/** GenesisState defines the epochs module's genesis state. */
export interface GenesisState {
  queries: Query[];
}

function createBaseQuery(): Query {
  return {
    id: "",
    connectionId: "",
    chainId: "",
    queryType: "",
    request: new Uint8Array(),
    callbackId: "",
    ttl: 0,
    requestSent: false,
  };
}

export const Query = {
  encode(message: Query, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.connectionId !== "") {
      writer.uint32(18).string(message.connectionId);
    }
    if (message.chainId !== "") {
      writer.uint32(26).string(message.chainId);
    }
    if (message.queryType !== "") {
      writer.uint32(34).string(message.queryType);
    }
    if (message.request.length !== 0) {
      writer.uint32(42).bytes(message.request);
    }
    if (message.callbackId !== "") {
      writer.uint32(66).string(message.callbackId);
    }
    if (message.ttl !== 0) {
      writer.uint32(72).uint64(message.ttl);
    }
    if (message.requestSent === true) {
      writer.uint32(88).bool(message.requestSent);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Query {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuery();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.connectionId = reader.string();
          break;
        case 3:
          message.chainId = reader.string();
          break;
        case 4:
          message.queryType = reader.string();
          break;
        case 5:
          message.request = reader.bytes();
          break;
        case 8:
          message.callbackId = reader.string();
          break;
        case 9:
          message.ttl = longToNumber(reader.uint64() as Long);
          break;
        case 11:
          message.requestSent = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Query {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      connectionId: isSet(object.connectionId) ? String(object.connectionId) : "",
      chainId: isSet(object.chainId) ? String(object.chainId) : "",
      queryType: isSet(object.queryType) ? String(object.queryType) : "",
      request: isSet(object.request) ? bytesFromBase64(object.request) : new Uint8Array(),
      callbackId: isSet(object.callbackId) ? String(object.callbackId) : "",
      ttl: isSet(object.ttl) ? Number(object.ttl) : 0,
      requestSent: isSet(object.requestSent) ? Boolean(object.requestSent) : false,
    };
  },

  toJSON(message: Query): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.connectionId !== undefined && (obj.connectionId = message.connectionId);
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.queryType !== undefined && (obj.queryType = message.queryType);
    message.request !== undefined
      && (obj.request = base64FromBytes(message.request !== undefined ? message.request : new Uint8Array()));
    message.callbackId !== undefined && (obj.callbackId = message.callbackId);
    message.ttl !== undefined && (obj.ttl = Math.round(message.ttl));
    message.requestSent !== undefined && (obj.requestSent = message.requestSent);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Query>, I>>(object: I): Query {
    const message = createBaseQuery();
    message.id = object.id ?? "";
    message.connectionId = object.connectionId ?? "";
    message.chainId = object.chainId ?? "";
    message.queryType = object.queryType ?? "";
    message.request = object.request ?? new Uint8Array();
    message.callbackId = object.callbackId ?? "";
    message.ttl = object.ttl ?? 0;
    message.requestSent = object.requestSent ?? false;
    return message;
  },
};

function createBaseDataPoint(): DataPoint {
  return { id: "", remoteHeight: "", localHeight: "", value: new Uint8Array() };
}

export const DataPoint = {
  encode(message: DataPoint, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.remoteHeight !== "") {
      writer.uint32(18).string(message.remoteHeight);
    }
    if (message.localHeight !== "") {
      writer.uint32(26).string(message.localHeight);
    }
    if (message.value.length !== 0) {
      writer.uint32(34).bytes(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DataPoint {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDataPoint();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.remoteHeight = reader.string();
          break;
        case 3:
          message.localHeight = reader.string();
          break;
        case 4:
          message.value = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DataPoint {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      remoteHeight: isSet(object.remoteHeight) ? String(object.remoteHeight) : "",
      localHeight: isSet(object.localHeight) ? String(object.localHeight) : "",
      value: isSet(object.value) ? bytesFromBase64(object.value) : new Uint8Array(),
    };
  },

  toJSON(message: DataPoint): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.remoteHeight !== undefined && (obj.remoteHeight = message.remoteHeight);
    message.localHeight !== undefined && (obj.localHeight = message.localHeight);
    message.value !== undefined
      && (obj.value = base64FromBytes(message.value !== undefined ? message.value : new Uint8Array()));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DataPoint>, I>>(object: I): DataPoint {
    const message = createBaseDataPoint();
    message.id = object.id ?? "";
    message.remoteHeight = object.remoteHeight ?? "";
    message.localHeight = object.localHeight ?? "";
    message.value = object.value ?? new Uint8Array();
    return message;
  },
};

function createBaseGenesisState(): GenesisState {
  return { queries: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.queries) {
      Query.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.queries.push(Query.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return { queries: Array.isArray(object?.queries) ? object.queries.map((e: any) => Query.fromJSON(e)) : [] };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.queries) {
      obj.queries = message.queries.map((e) => e ? Query.toJSON(e) : undefined);
    } else {
      obj.queries = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.queries = object.queries?.map((e) => Query.fromPartial(e)) || [];
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
