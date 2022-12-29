/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { CallbackData } from "./callback_data";
import { Params } from "./params";

export const protobufPackage = "stayking.icacallbacks";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetCallbackDataRequest {
  callbackKey: string;
}

export interface QueryGetCallbackDataResponse {
  callbackData: CallbackData | undefined;
}

export interface QueryAllCallbackDataRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllCallbackDataResponse {
  callbackData: CallbackData[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetCallbackDataRequest(): QueryGetCallbackDataRequest {
  return { callbackKey: "" };
}

export const QueryGetCallbackDataRequest = {
  encode(message: QueryGetCallbackDataRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.callbackKey !== "") {
      writer.uint32(10).string(message.callbackKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetCallbackDataRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetCallbackDataRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.callbackKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCallbackDataRequest {
    return { callbackKey: isSet(object.callbackKey) ? String(object.callbackKey) : "" };
  },

  toJSON(message: QueryGetCallbackDataRequest): unknown {
    const obj: any = {};
    message.callbackKey !== undefined && (obj.callbackKey = message.callbackKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetCallbackDataRequest>, I>>(object: I): QueryGetCallbackDataRequest {
    const message = createBaseQueryGetCallbackDataRequest();
    message.callbackKey = object.callbackKey ?? "";
    return message;
  },
};

function createBaseQueryGetCallbackDataResponse(): QueryGetCallbackDataResponse {
  return { callbackData: undefined };
}

export const QueryGetCallbackDataResponse = {
  encode(message: QueryGetCallbackDataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.callbackData !== undefined) {
      CallbackData.encode(message.callbackData, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetCallbackDataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetCallbackDataResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.callbackData = CallbackData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCallbackDataResponse {
    return { callbackData: isSet(object.callbackData) ? CallbackData.fromJSON(object.callbackData) : undefined };
  },

  toJSON(message: QueryGetCallbackDataResponse): unknown {
    const obj: any = {};
    message.callbackData !== undefined
      && (obj.callbackData = message.callbackData ? CallbackData.toJSON(message.callbackData) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetCallbackDataResponse>, I>>(object: I): QueryGetCallbackDataResponse {
    const message = createBaseQueryGetCallbackDataResponse();
    message.callbackData = (object.callbackData !== undefined && object.callbackData !== null)
      ? CallbackData.fromPartial(object.callbackData)
      : undefined;
    return message;
  },
};

function createBaseQueryAllCallbackDataRequest(): QueryAllCallbackDataRequest {
  return { pagination: undefined };
}

export const QueryAllCallbackDataRequest = {
  encode(message: QueryAllCallbackDataRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllCallbackDataRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllCallbackDataRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllCallbackDataRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllCallbackDataRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllCallbackDataRequest>, I>>(object: I): QueryAllCallbackDataRequest {
    const message = createBaseQueryAllCallbackDataRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllCallbackDataResponse(): QueryAllCallbackDataResponse {
  return { callbackData: [], pagination: undefined };
}

export const QueryAllCallbackDataResponse = {
  encode(message: QueryAllCallbackDataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.callbackData) {
      CallbackData.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllCallbackDataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllCallbackDataResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.callbackData.push(CallbackData.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllCallbackDataResponse {
    return {
      callbackData: Array.isArray(object?.callbackData)
        ? object.callbackData.map((e: any) => CallbackData.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllCallbackDataResponse): unknown {
    const obj: any = {};
    if (message.callbackData) {
      obj.callbackData = message.callbackData.map((e) => e ? CallbackData.toJSON(e) : undefined);
    } else {
      obj.callbackData = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllCallbackDataResponse>, I>>(object: I): QueryAllCallbackDataResponse {
    const message = createBaseQueryAllCallbackDataResponse();
    message.callbackData = object.callbackData?.map((e) => CallbackData.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a CallbackData by index. */
  CallbackData(request: QueryGetCallbackDataRequest): Promise<QueryGetCallbackDataResponse>;
  /** Queries a list of CallbackData items. */
  CallbackDataAll(request: QueryAllCallbackDataRequest): Promise<QueryAllCallbackDataResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.CallbackData = this.CallbackData.bind(this);
    this.CallbackDataAll = this.CallbackDataAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.icacallbacks.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  CallbackData(request: QueryGetCallbackDataRequest): Promise<QueryGetCallbackDataResponse> {
    const data = QueryGetCallbackDataRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.icacallbacks.Query", "CallbackData", data);
    return promise.then((data) => QueryGetCallbackDataResponse.decode(new _m0.Reader(data)));
  }

  CallbackDataAll(request: QueryAllCallbackDataRequest): Promise<QueryAllCallbackDataResponse> {
    const data = QueryAllCallbackDataRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.icacallbacks.Query", "CallbackDataAll", data);
    return promise.then((data) => QueryAllCallbackDataResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
