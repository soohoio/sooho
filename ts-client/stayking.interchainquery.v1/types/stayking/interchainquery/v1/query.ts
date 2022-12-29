/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Query } from "./genesis";

export const protobufPackage = "stayking.interchainquery.v1";

export interface QueryPendingQueriesRequest {
}

export interface QueryPendingQueriesResponse {
  pendingQueries: Query[];
}

function createBaseQueryPendingQueriesRequest(): QueryPendingQueriesRequest {
  return {};
}

export const QueryPendingQueriesRequest = {
  encode(_: QueryPendingQueriesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryPendingQueriesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryPendingQueriesRequest();
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

  fromJSON(_: any): QueryPendingQueriesRequest {
    return {};
  },

  toJSON(_: QueryPendingQueriesRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryPendingQueriesRequest>, I>>(_: I): QueryPendingQueriesRequest {
    const message = createBaseQueryPendingQueriesRequest();
    return message;
  },
};

function createBaseQueryPendingQueriesResponse(): QueryPendingQueriesResponse {
  return { pendingQueries: [] };
}

export const QueryPendingQueriesResponse = {
  encode(message: QueryPendingQueriesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.pendingQueries) {
      Query.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryPendingQueriesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryPendingQueriesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pendingQueries.push(Query.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryPendingQueriesResponse {
    return {
      pendingQueries: Array.isArray(object?.pendingQueries)
        ? object.pendingQueries.map((e: any) => Query.fromJSON(e))
        : [],
    };
  },

  toJSON(message: QueryPendingQueriesResponse): unknown {
    const obj: any = {};
    if (message.pendingQueries) {
      obj.pendingQueries = message.pendingQueries.map((e) => e ? Query.toJSON(e) : undefined);
    } else {
      obj.pendingQueries = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryPendingQueriesResponse>, I>>(object: I): QueryPendingQueriesResponse {
    const message = createBaseQueryPendingQueriesResponse();
    message.pendingQueries = object.pendingQueries?.map((e) => Query.fromPartial(e)) || [];
    return message;
  },
};

export interface QueryService {
  PendingQueries(request: QueryPendingQueriesRequest): Promise<QueryPendingQueriesResponse>;
}

export class QueryServiceClientImpl implements QueryService {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.PendingQueries = this.PendingQueries.bind(this);
  }
  PendingQueries(request: QueryPendingQueriesRequest): Promise<QueryPendingQueriesResponse> {
    const data = QueryPendingQueriesRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.interchainquery.v1.QueryService", "PendingQueries", data);
    return promise.then((data) => QueryPendingQueriesResponse.decode(new _m0.Reader(data)));
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
