/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { DepositRecord, EpochUnbondingRecord, Params, UserRedemptionRecord } from "./genesis";

export const protobufPackage = "stayking.records";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetDepositRecordRequest {
  id: number;
}

export interface QueryGetDepositRecordResponse {
  depositRecord: DepositRecord | undefined;
}

export interface QueryAllDepositRecordRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDepositRecordResponse {
  depositRecord: DepositRecord[];
  pagination: PageResponse | undefined;
}

export interface QueryGetUserRedemptionRecordRequest {
  id: string;
}

export interface QueryGetUserRedemptionRecordResponse {
  userRedemptionRecord: UserRedemptionRecord | undefined;
}

export interface QueryAllUserRedemptionRecordRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllUserRedemptionRecordResponse {
  userRedemptionRecord: UserRedemptionRecord[];
  pagination: PageResponse | undefined;
}

/** Query UserRedemptionRecords by chainId / userId pair */
export interface QueryAllUserRedemptionRecordForUserRequest {
  chainId: string;
  day: number;
  address: string;
  limit: number;
  pagination: PageRequest | undefined;
}

export interface QueryAllUserRedemptionRecordForUserResponse {
  userRedemptionRecord: UserRedemptionRecord[];
  pagination: PageResponse | undefined;
}

export interface QueryGetEpochUnbondingRecordRequest {
  epochNumber: number;
}

export interface QueryGetEpochUnbondingRecordResponse {
  epochUnbondingRecord: EpochUnbondingRecord | undefined;
}

export interface QueryAllEpochUnbondingRecordRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllEpochUnbondingRecordResponse {
  epochUnbondingRecord: EpochUnbondingRecord[];
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

function createBaseQueryGetDepositRecordRequest(): QueryGetDepositRecordRequest {
  return { id: 0 };
}

export const QueryGetDepositRecordRequest = {
  encode(message: QueryGetDepositRecordRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDepositRecordRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDepositRecordRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDepositRecordRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryGetDepositRecordRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDepositRecordRequest>, I>>(object: I): QueryGetDepositRecordRequest {
    const message = createBaseQueryGetDepositRecordRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryGetDepositRecordResponse(): QueryGetDepositRecordResponse {
  return { depositRecord: undefined };
}

export const QueryGetDepositRecordResponse = {
  encode(message: QueryGetDepositRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.depositRecord !== undefined) {
      DepositRecord.encode(message.depositRecord, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDepositRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDepositRecordResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.depositRecord = DepositRecord.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDepositRecordResponse {
    return { depositRecord: isSet(object.depositRecord) ? DepositRecord.fromJSON(object.depositRecord) : undefined };
  },

  toJSON(message: QueryGetDepositRecordResponse): unknown {
    const obj: any = {};
    message.depositRecord !== undefined
      && (obj.depositRecord = message.depositRecord ? DepositRecord.toJSON(message.depositRecord) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDepositRecordResponse>, I>>(
    object: I,
  ): QueryGetDepositRecordResponse {
    const message = createBaseQueryGetDepositRecordResponse();
    message.depositRecord = (object.depositRecord !== undefined && object.depositRecord !== null)
      ? DepositRecord.fromPartial(object.depositRecord)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDepositRecordRequest(): QueryAllDepositRecordRequest {
  return { pagination: undefined };
}

export const QueryAllDepositRecordRequest = {
  encode(message: QueryAllDepositRecordRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDepositRecordRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDepositRecordRequest();
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

  fromJSON(object: any): QueryAllDepositRecordRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllDepositRecordRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDepositRecordRequest>, I>>(object: I): QueryAllDepositRecordRequest {
    const message = createBaseQueryAllDepositRecordRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDepositRecordResponse(): QueryAllDepositRecordResponse {
  return { depositRecord: [], pagination: undefined };
}

export const QueryAllDepositRecordResponse = {
  encode(message: QueryAllDepositRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.depositRecord) {
      DepositRecord.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDepositRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDepositRecordResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.depositRecord.push(DepositRecord.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllDepositRecordResponse {
    return {
      depositRecord: Array.isArray(object?.depositRecord)
        ? object.depositRecord.map((e: any) => DepositRecord.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllDepositRecordResponse): unknown {
    const obj: any = {};
    if (message.depositRecord) {
      obj.depositRecord = message.depositRecord.map((e) => e ? DepositRecord.toJSON(e) : undefined);
    } else {
      obj.depositRecord = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDepositRecordResponse>, I>>(
    object: I,
  ): QueryAllDepositRecordResponse {
    const message = createBaseQueryAllDepositRecordResponse();
    message.depositRecord = object.depositRecord?.map((e) => DepositRecord.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetUserRedemptionRecordRequest(): QueryGetUserRedemptionRecordRequest {
  return { id: "" };
}

export const QueryGetUserRedemptionRecordRequest = {
  encode(message: QueryGetUserRedemptionRecordRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserRedemptionRecordRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserRedemptionRecordRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserRedemptionRecordRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: QueryGetUserRedemptionRecordRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserRedemptionRecordRequest>, I>>(
    object: I,
  ): QueryGetUserRedemptionRecordRequest {
    const message = createBaseQueryGetUserRedemptionRecordRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseQueryGetUserRedemptionRecordResponse(): QueryGetUserRedemptionRecordResponse {
  return { userRedemptionRecord: undefined };
}

export const QueryGetUserRedemptionRecordResponse = {
  encode(message: QueryGetUserRedemptionRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userRedemptionRecord !== undefined) {
      UserRedemptionRecord.encode(message.userRedemptionRecord, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserRedemptionRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserRedemptionRecordResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userRedemptionRecord = UserRedemptionRecord.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserRedemptionRecordResponse {
    return {
      userRedemptionRecord: isSet(object.userRedemptionRecord)
        ? UserRedemptionRecord.fromJSON(object.userRedemptionRecord)
        : undefined,
    };
  },

  toJSON(message: QueryGetUserRedemptionRecordResponse): unknown {
    const obj: any = {};
    message.userRedemptionRecord !== undefined && (obj.userRedemptionRecord = message.userRedemptionRecord
      ? UserRedemptionRecord.toJSON(message.userRedemptionRecord)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserRedemptionRecordResponse>, I>>(
    object: I,
  ): QueryGetUserRedemptionRecordResponse {
    const message = createBaseQueryGetUserRedemptionRecordResponse();
    message.userRedemptionRecord = (object.userRedemptionRecord !== undefined && object.userRedemptionRecord !== null)
      ? UserRedemptionRecord.fromPartial(object.userRedemptionRecord)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUserRedemptionRecordRequest(): QueryAllUserRedemptionRecordRequest {
  return { pagination: undefined };
}

export const QueryAllUserRedemptionRecordRequest = {
  encode(message: QueryAllUserRedemptionRecordRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserRedemptionRecordRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserRedemptionRecordRequest();
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

  fromJSON(object: any): QueryAllUserRedemptionRecordRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllUserRedemptionRecordRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserRedemptionRecordRequest>, I>>(
    object: I,
  ): QueryAllUserRedemptionRecordRequest {
    const message = createBaseQueryAllUserRedemptionRecordRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUserRedemptionRecordResponse(): QueryAllUserRedemptionRecordResponse {
  return { userRedemptionRecord: [], pagination: undefined };
}

export const QueryAllUserRedemptionRecordResponse = {
  encode(message: QueryAllUserRedemptionRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.userRedemptionRecord) {
      UserRedemptionRecord.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserRedemptionRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserRedemptionRecordResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userRedemptionRecord.push(UserRedemptionRecord.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllUserRedemptionRecordResponse {
    return {
      userRedemptionRecord: Array.isArray(object?.userRedemptionRecord)
        ? object.userRedemptionRecord.map((e: any) => UserRedemptionRecord.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllUserRedemptionRecordResponse): unknown {
    const obj: any = {};
    if (message.userRedemptionRecord) {
      obj.userRedemptionRecord = message.userRedemptionRecord.map((e) =>
        e ? UserRedemptionRecord.toJSON(e) : undefined
      );
    } else {
      obj.userRedemptionRecord = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserRedemptionRecordResponse>, I>>(
    object: I,
  ): QueryAllUserRedemptionRecordResponse {
    const message = createBaseQueryAllUserRedemptionRecordResponse();
    message.userRedemptionRecord = object.userRedemptionRecord?.map((e) => UserRedemptionRecord.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUserRedemptionRecordForUserRequest(): QueryAllUserRedemptionRecordForUserRequest {
  return { chainId: "", day: 0, address: "", limit: 0, pagination: undefined };
}

export const QueryAllUserRedemptionRecordForUserRequest = {
  encode(message: QueryAllUserRedemptionRecordForUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.chainId !== "") {
      writer.uint32(10).string(message.chainId);
    }
    if (message.day !== 0) {
      writer.uint32(16).uint64(message.day);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.limit !== 0) {
      writer.uint32(32).uint64(message.limit);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserRedemptionRecordForUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserRedemptionRecordForUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.chainId = reader.string();
          break;
        case 2:
          message.day = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.address = reader.string();
          break;
        case 4:
          message.limit = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUserRedemptionRecordForUserRequest {
    return {
      chainId: isSet(object.chainId) ? String(object.chainId) : "",
      day: isSet(object.day) ? Number(object.day) : 0,
      address: isSet(object.address) ? String(object.address) : "",
      limit: isSet(object.limit) ? Number(object.limit) : 0,
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllUserRedemptionRecordForUserRequest): unknown {
    const obj: any = {};
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.day !== undefined && (obj.day = Math.round(message.day));
    message.address !== undefined && (obj.address = message.address);
    message.limit !== undefined && (obj.limit = Math.round(message.limit));
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserRedemptionRecordForUserRequest>, I>>(
    object: I,
  ): QueryAllUserRedemptionRecordForUserRequest {
    const message = createBaseQueryAllUserRedemptionRecordForUserRequest();
    message.chainId = object.chainId ?? "";
    message.day = object.day ?? 0;
    message.address = object.address ?? "";
    message.limit = object.limit ?? 0;
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUserRedemptionRecordForUserResponse(): QueryAllUserRedemptionRecordForUserResponse {
  return { userRedemptionRecord: [], pagination: undefined };
}

export const QueryAllUserRedemptionRecordForUserResponse = {
  encode(message: QueryAllUserRedemptionRecordForUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.userRedemptionRecord) {
      UserRedemptionRecord.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserRedemptionRecordForUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserRedemptionRecordForUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userRedemptionRecord.push(UserRedemptionRecord.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllUserRedemptionRecordForUserResponse {
    return {
      userRedemptionRecord: Array.isArray(object?.userRedemptionRecord)
        ? object.userRedemptionRecord.map((e: any) => UserRedemptionRecord.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllUserRedemptionRecordForUserResponse): unknown {
    const obj: any = {};
    if (message.userRedemptionRecord) {
      obj.userRedemptionRecord = message.userRedemptionRecord.map((e) =>
        e ? UserRedemptionRecord.toJSON(e) : undefined
      );
    } else {
      obj.userRedemptionRecord = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserRedemptionRecordForUserResponse>, I>>(
    object: I,
  ): QueryAllUserRedemptionRecordForUserResponse {
    const message = createBaseQueryAllUserRedemptionRecordForUserResponse();
    message.userRedemptionRecord = object.userRedemptionRecord?.map((e) => UserRedemptionRecord.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetEpochUnbondingRecordRequest(): QueryGetEpochUnbondingRecordRequest {
  return { epochNumber: 0 };
}

export const QueryGetEpochUnbondingRecordRequest = {
  encode(message: QueryGetEpochUnbondingRecordRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.epochNumber !== 0) {
      writer.uint32(8).uint64(message.epochNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEpochUnbondingRecordRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEpochUnbondingRecordRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.epochNumber = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEpochUnbondingRecordRequest {
    return { epochNumber: isSet(object.epochNumber) ? Number(object.epochNumber) : 0 };
  },

  toJSON(message: QueryGetEpochUnbondingRecordRequest): unknown {
    const obj: any = {};
    message.epochNumber !== undefined && (obj.epochNumber = Math.round(message.epochNumber));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEpochUnbondingRecordRequest>, I>>(
    object: I,
  ): QueryGetEpochUnbondingRecordRequest {
    const message = createBaseQueryGetEpochUnbondingRecordRequest();
    message.epochNumber = object.epochNumber ?? 0;
    return message;
  },
};

function createBaseQueryGetEpochUnbondingRecordResponse(): QueryGetEpochUnbondingRecordResponse {
  return { epochUnbondingRecord: undefined };
}

export const QueryGetEpochUnbondingRecordResponse = {
  encode(message: QueryGetEpochUnbondingRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.epochUnbondingRecord !== undefined) {
      EpochUnbondingRecord.encode(message.epochUnbondingRecord, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEpochUnbondingRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEpochUnbondingRecordResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.epochUnbondingRecord = EpochUnbondingRecord.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEpochUnbondingRecordResponse {
    return {
      epochUnbondingRecord: isSet(object.epochUnbondingRecord)
        ? EpochUnbondingRecord.fromJSON(object.epochUnbondingRecord)
        : undefined,
    };
  },

  toJSON(message: QueryGetEpochUnbondingRecordResponse): unknown {
    const obj: any = {};
    message.epochUnbondingRecord !== undefined && (obj.epochUnbondingRecord = message.epochUnbondingRecord
      ? EpochUnbondingRecord.toJSON(message.epochUnbondingRecord)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEpochUnbondingRecordResponse>, I>>(
    object: I,
  ): QueryGetEpochUnbondingRecordResponse {
    const message = createBaseQueryGetEpochUnbondingRecordResponse();
    message.epochUnbondingRecord = (object.epochUnbondingRecord !== undefined && object.epochUnbondingRecord !== null)
      ? EpochUnbondingRecord.fromPartial(object.epochUnbondingRecord)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEpochUnbondingRecordRequest(): QueryAllEpochUnbondingRecordRequest {
  return { pagination: undefined };
}

export const QueryAllEpochUnbondingRecordRequest = {
  encode(message: QueryAllEpochUnbondingRecordRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEpochUnbondingRecordRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEpochUnbondingRecordRequest();
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

  fromJSON(object: any): QueryAllEpochUnbondingRecordRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllEpochUnbondingRecordRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEpochUnbondingRecordRequest>, I>>(
    object: I,
  ): QueryAllEpochUnbondingRecordRequest {
    const message = createBaseQueryAllEpochUnbondingRecordRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEpochUnbondingRecordResponse(): QueryAllEpochUnbondingRecordResponse {
  return { epochUnbondingRecord: [], pagination: undefined };
}

export const QueryAllEpochUnbondingRecordResponse = {
  encode(message: QueryAllEpochUnbondingRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.epochUnbondingRecord) {
      EpochUnbondingRecord.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEpochUnbondingRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEpochUnbondingRecordResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.epochUnbondingRecord.push(EpochUnbondingRecord.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllEpochUnbondingRecordResponse {
    return {
      epochUnbondingRecord: Array.isArray(object?.epochUnbondingRecord)
        ? object.epochUnbondingRecord.map((e: any) => EpochUnbondingRecord.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllEpochUnbondingRecordResponse): unknown {
    const obj: any = {};
    if (message.epochUnbondingRecord) {
      obj.epochUnbondingRecord = message.epochUnbondingRecord.map((e) =>
        e ? EpochUnbondingRecord.toJSON(e) : undefined
      );
    } else {
      obj.epochUnbondingRecord = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEpochUnbondingRecordResponse>, I>>(
    object: I,
  ): QueryAllEpochUnbondingRecordResponse {
    const message = createBaseQueryAllEpochUnbondingRecordResponse();
    message.epochUnbondingRecord = object.epochUnbondingRecord?.map((e) => EpochUnbondingRecord.fromPartial(e)) || [];
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
  /** Queries a UserRedemptionRecord by id. */
  UserRedemptionRecord(request: QueryGetUserRedemptionRecordRequest): Promise<QueryGetUserRedemptionRecordResponse>;
  /** Queries a list of UserRedemptionRecord items. */
  UserRedemptionRecordAll(request: QueryAllUserRedemptionRecordRequest): Promise<QueryAllUserRedemptionRecordResponse>;
  /** Queries a list of UserRedemptionRecord items by chainId / userId pair. */
  UserRedemptionRecordForUser(
    request: QueryAllUserRedemptionRecordForUserRequest,
  ): Promise<QueryAllUserRedemptionRecordForUserResponse>;
  /** Queries a EpochUnbondingRecord by id. */
  EpochUnbondingRecord(request: QueryGetEpochUnbondingRecordRequest): Promise<QueryGetEpochUnbondingRecordResponse>;
  /** Queries a list of EpochUnbondingRecord items. */
  EpochUnbondingRecordAll(request: QueryAllEpochUnbondingRecordRequest): Promise<QueryAllEpochUnbondingRecordResponse>;
  /** Queries a DepositRecord by id. */
  DepositRecord(request: QueryGetDepositRecordRequest): Promise<QueryGetDepositRecordResponse>;
  /** Queries a list of DepositRecord items. */
  DepositRecordAll(request: QueryAllDepositRecordRequest): Promise<QueryAllDepositRecordResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.UserRedemptionRecord = this.UserRedemptionRecord.bind(this);
    this.UserRedemptionRecordAll = this.UserRedemptionRecordAll.bind(this);
    this.UserRedemptionRecordForUser = this.UserRedemptionRecordForUser.bind(this);
    this.EpochUnbondingRecord = this.EpochUnbondingRecord.bind(this);
    this.EpochUnbondingRecordAll = this.EpochUnbondingRecordAll.bind(this);
    this.DepositRecord = this.DepositRecord.bind(this);
    this.DepositRecordAll = this.DepositRecordAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.records.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  UserRedemptionRecord(request: QueryGetUserRedemptionRecordRequest): Promise<QueryGetUserRedemptionRecordResponse> {
    const data = QueryGetUserRedemptionRecordRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.records.Query", "UserRedemptionRecord", data);
    return promise.then((data) => QueryGetUserRedemptionRecordResponse.decode(new _m0.Reader(data)));
  }

  UserRedemptionRecordAll(request: QueryAllUserRedemptionRecordRequest): Promise<QueryAllUserRedemptionRecordResponse> {
    const data = QueryAllUserRedemptionRecordRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.records.Query", "UserRedemptionRecordAll", data);
    return promise.then((data) => QueryAllUserRedemptionRecordResponse.decode(new _m0.Reader(data)));
  }

  UserRedemptionRecordForUser(
    request: QueryAllUserRedemptionRecordForUserRequest,
  ): Promise<QueryAllUserRedemptionRecordForUserResponse> {
    const data = QueryAllUserRedemptionRecordForUserRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.records.Query", "UserRedemptionRecordForUser", data);
    return promise.then((data) => QueryAllUserRedemptionRecordForUserResponse.decode(new _m0.Reader(data)));
  }

  EpochUnbondingRecord(request: QueryGetEpochUnbondingRecordRequest): Promise<QueryGetEpochUnbondingRecordResponse> {
    const data = QueryGetEpochUnbondingRecordRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.records.Query", "EpochUnbondingRecord", data);
    return promise.then((data) => QueryGetEpochUnbondingRecordResponse.decode(new _m0.Reader(data)));
  }

  EpochUnbondingRecordAll(request: QueryAllEpochUnbondingRecordRequest): Promise<QueryAllEpochUnbondingRecordResponse> {
    const data = QueryAllEpochUnbondingRecordRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.records.Query", "EpochUnbondingRecordAll", data);
    return promise.then((data) => QueryAllEpochUnbondingRecordResponse.decode(new _m0.Reader(data)));
  }

  DepositRecord(request: QueryGetDepositRecordRequest): Promise<QueryGetDepositRecordResponse> {
    const data = QueryGetDepositRecordRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.records.Query", "DepositRecord", data);
    return promise.then((data) => QueryGetDepositRecordResponse.decode(new _m0.Reader(data)));
  }

  DepositRecordAll(request: QueryAllDepositRecordRequest): Promise<QueryAllDepositRecordResponse> {
    const data = QueryAllDepositRecordRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.records.Query", "DepositRecordAll", data);
    return promise.then((data) => QueryAllDepositRecordResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
