/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { EpochTracker } from "./epoch_tracker";
import { HostZone } from "./host_zone";
import { ICAAccount } from "./ica_account";
import { Params } from "./params";
import { Validator } from "./validator";

export const protobufPackage = "stayking.stakeibc";

/**
 * QueryInterchainAccountFromAddressRequest is the request type for the
 * Query/InterchainAccountAddress RPC
 */
export interface QueryInterchainAccountFromAddressRequest {
  owner: string;
  connectionId: string;
}

/**
 * QueryInterchainAccountFromAddressResponse the response type for the
 * Query/InterchainAccountAddress RPC
 */
export interface QueryInterchainAccountFromAddressResponse {
  interchainAccountAddress: string;
}

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetValidatorsRequest {
  chainId: string;
}

export interface QueryGetValidatorsResponse {
  validators: Validator[];
}

export interface QueryGetICAAccountRequest {
}

export interface QueryGetICAAccountResponse {
  icaAccount: ICAAccount | undefined;
}

export interface QueryGetHostZoneRequest {
  chainId: string;
}

export interface QueryGetHostZoneResponse {
  hostZone: HostZone | undefined;
}

export interface QueryAllHostZoneRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllHostZoneResponse {
  hostZone: HostZone[];
  pagination: PageResponse | undefined;
}

export interface QueryModuleAddressRequest {
  name: string;
}

export interface QueryModuleAddressResponse {
  addr: string;
}

export interface QueryGetEpochTrackerRequest {
  epochIdentifier: string;
}

export interface QueryGetEpochTrackerResponse {
  epochTracker: EpochTracker | undefined;
}

export interface QueryAllEpochTrackerRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllEpochTrackerResponse {
  epochTracker: EpochTracker[];
  pagination: PageResponse | undefined;
}

function createBaseQueryInterchainAccountFromAddressRequest(): QueryInterchainAccountFromAddressRequest {
  return { owner: "", connectionId: "" };
}

export const QueryInterchainAccountFromAddressRequest = {
  encode(message: QueryInterchainAccountFromAddressRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.connectionId !== "") {
      writer.uint32(18).string(message.connectionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryInterchainAccountFromAddressRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryInterchainAccountFromAddressRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.connectionId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryInterchainAccountFromAddressRequest {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      connectionId: isSet(object.connectionId) ? String(object.connectionId) : "",
    };
  },

  toJSON(message: QueryInterchainAccountFromAddressRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.connectionId !== undefined && (obj.connectionId = message.connectionId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryInterchainAccountFromAddressRequest>, I>>(
    object: I,
  ): QueryInterchainAccountFromAddressRequest {
    const message = createBaseQueryInterchainAccountFromAddressRequest();
    message.owner = object.owner ?? "";
    message.connectionId = object.connectionId ?? "";
    return message;
  },
};

function createBaseQueryInterchainAccountFromAddressResponse(): QueryInterchainAccountFromAddressResponse {
  return { interchainAccountAddress: "" };
}

export const QueryInterchainAccountFromAddressResponse = {
  encode(message: QueryInterchainAccountFromAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.interchainAccountAddress !== "") {
      writer.uint32(10).string(message.interchainAccountAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryInterchainAccountFromAddressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryInterchainAccountFromAddressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.interchainAccountAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryInterchainAccountFromAddressResponse {
    return {
      interchainAccountAddress: isSet(object.interchainAccountAddress) ? String(object.interchainAccountAddress) : "",
    };
  },

  toJSON(message: QueryInterchainAccountFromAddressResponse): unknown {
    const obj: any = {};
    message.interchainAccountAddress !== undefined && (obj.interchainAccountAddress = message.interchainAccountAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryInterchainAccountFromAddressResponse>, I>>(
    object: I,
  ): QueryInterchainAccountFromAddressResponse {
    const message = createBaseQueryInterchainAccountFromAddressResponse();
    message.interchainAccountAddress = object.interchainAccountAddress ?? "";
    return message;
  },
};

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

function createBaseQueryGetValidatorsRequest(): QueryGetValidatorsRequest {
  return { chainId: "" };
}

export const QueryGetValidatorsRequest = {
  encode(message: QueryGetValidatorsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.chainId !== "") {
      writer.uint32(10).string(message.chainId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetValidatorsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetValidatorsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.chainId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetValidatorsRequest {
    return { chainId: isSet(object.chainId) ? String(object.chainId) : "" };
  },

  toJSON(message: QueryGetValidatorsRequest): unknown {
    const obj: any = {};
    message.chainId !== undefined && (obj.chainId = message.chainId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetValidatorsRequest>, I>>(object: I): QueryGetValidatorsRequest {
    const message = createBaseQueryGetValidatorsRequest();
    message.chainId = object.chainId ?? "";
    return message;
  },
};

function createBaseQueryGetValidatorsResponse(): QueryGetValidatorsResponse {
  return { validators: [] };
}

export const QueryGetValidatorsResponse = {
  encode(message: QueryGetValidatorsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.validators) {
      Validator.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetValidatorsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetValidatorsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validators.push(Validator.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetValidatorsResponse {
    return {
      validators: Array.isArray(object?.validators) ? object.validators.map((e: any) => Validator.fromJSON(e)) : [],
    };
  },

  toJSON(message: QueryGetValidatorsResponse): unknown {
    const obj: any = {};
    if (message.validators) {
      obj.validators = message.validators.map((e) => e ? Validator.toJSON(e) : undefined);
    } else {
      obj.validators = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetValidatorsResponse>, I>>(object: I): QueryGetValidatorsResponse {
    const message = createBaseQueryGetValidatorsResponse();
    message.validators = object.validators?.map((e) => Validator.fromPartial(e)) || [];
    return message;
  },
};

function createBaseQueryGetICAAccountRequest(): QueryGetICAAccountRequest {
  return {};
}

export const QueryGetICAAccountRequest = {
  encode(_: QueryGetICAAccountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetICAAccountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetICAAccountRequest();
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

  fromJSON(_: any): QueryGetICAAccountRequest {
    return {};
  },

  toJSON(_: QueryGetICAAccountRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetICAAccountRequest>, I>>(_: I): QueryGetICAAccountRequest {
    const message = createBaseQueryGetICAAccountRequest();
    return message;
  },
};

function createBaseQueryGetICAAccountResponse(): QueryGetICAAccountResponse {
  return { icaAccount: undefined };
}

export const QueryGetICAAccountResponse = {
  encode(message: QueryGetICAAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.icaAccount !== undefined) {
      ICAAccount.encode(message.icaAccount, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetICAAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetICAAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.icaAccount = ICAAccount.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetICAAccountResponse {
    return { icaAccount: isSet(object.icaAccount) ? ICAAccount.fromJSON(object.icaAccount) : undefined };
  },

  toJSON(message: QueryGetICAAccountResponse): unknown {
    const obj: any = {};
    message.icaAccount !== undefined
      && (obj.icaAccount = message.icaAccount ? ICAAccount.toJSON(message.icaAccount) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetICAAccountResponse>, I>>(object: I): QueryGetICAAccountResponse {
    const message = createBaseQueryGetICAAccountResponse();
    message.icaAccount = (object.icaAccount !== undefined && object.icaAccount !== null)
      ? ICAAccount.fromPartial(object.icaAccount)
      : undefined;
    return message;
  },
};

function createBaseQueryGetHostZoneRequest(): QueryGetHostZoneRequest {
  return { chainId: "" };
}

export const QueryGetHostZoneRequest = {
  encode(message: QueryGetHostZoneRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.chainId !== "") {
      writer.uint32(10).string(message.chainId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetHostZoneRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetHostZoneRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.chainId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetHostZoneRequest {
    return { chainId: isSet(object.chainId) ? String(object.chainId) : "" };
  },

  toJSON(message: QueryGetHostZoneRequest): unknown {
    const obj: any = {};
    message.chainId !== undefined && (obj.chainId = message.chainId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetHostZoneRequest>, I>>(object: I): QueryGetHostZoneRequest {
    const message = createBaseQueryGetHostZoneRequest();
    message.chainId = object.chainId ?? "";
    return message;
  },
};

function createBaseQueryGetHostZoneResponse(): QueryGetHostZoneResponse {
  return { hostZone: undefined };
}

export const QueryGetHostZoneResponse = {
  encode(message: QueryGetHostZoneResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.hostZone !== undefined) {
      HostZone.encode(message.hostZone, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetHostZoneResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetHostZoneResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.hostZone = HostZone.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetHostZoneResponse {
    return { hostZone: isSet(object.hostZone) ? HostZone.fromJSON(object.hostZone) : undefined };
  },

  toJSON(message: QueryGetHostZoneResponse): unknown {
    const obj: any = {};
    message.hostZone !== undefined && (obj.hostZone = message.hostZone ? HostZone.toJSON(message.hostZone) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetHostZoneResponse>, I>>(object: I): QueryGetHostZoneResponse {
    const message = createBaseQueryGetHostZoneResponse();
    message.hostZone = (object.hostZone !== undefined && object.hostZone !== null)
      ? HostZone.fromPartial(object.hostZone)
      : undefined;
    return message;
  },
};

function createBaseQueryAllHostZoneRequest(): QueryAllHostZoneRequest {
  return { pagination: undefined };
}

export const QueryAllHostZoneRequest = {
  encode(message: QueryAllHostZoneRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllHostZoneRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllHostZoneRequest();
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

  fromJSON(object: any): QueryAllHostZoneRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllHostZoneRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllHostZoneRequest>, I>>(object: I): QueryAllHostZoneRequest {
    const message = createBaseQueryAllHostZoneRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllHostZoneResponse(): QueryAllHostZoneResponse {
  return { hostZone: [], pagination: undefined };
}

export const QueryAllHostZoneResponse = {
  encode(message: QueryAllHostZoneResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.hostZone) {
      HostZone.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllHostZoneResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllHostZoneResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.hostZone.push(HostZone.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllHostZoneResponse {
    return {
      hostZone: Array.isArray(object?.hostZone) ? object.hostZone.map((e: any) => HostZone.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllHostZoneResponse): unknown {
    const obj: any = {};
    if (message.hostZone) {
      obj.hostZone = message.hostZone.map((e) => e ? HostZone.toJSON(e) : undefined);
    } else {
      obj.hostZone = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllHostZoneResponse>, I>>(object: I): QueryAllHostZoneResponse {
    const message = createBaseQueryAllHostZoneResponse();
    message.hostZone = object.hostZone?.map((e) => HostZone.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryModuleAddressRequest(): QueryModuleAddressRequest {
  return { name: "" };
}

export const QueryModuleAddressRequest = {
  encode(message: QueryModuleAddressRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryModuleAddressRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryModuleAddressRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryModuleAddressRequest {
    return { name: isSet(object.name) ? String(object.name) : "" };
  },

  toJSON(message: QueryModuleAddressRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryModuleAddressRequest>, I>>(object: I): QueryModuleAddressRequest {
    const message = createBaseQueryModuleAddressRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseQueryModuleAddressResponse(): QueryModuleAddressResponse {
  return { addr: "" };
}

export const QueryModuleAddressResponse = {
  encode(message: QueryModuleAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.addr !== "") {
      writer.uint32(10).string(message.addr);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryModuleAddressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryModuleAddressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.addr = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryModuleAddressResponse {
    return { addr: isSet(object.addr) ? String(object.addr) : "" };
  },

  toJSON(message: QueryModuleAddressResponse): unknown {
    const obj: any = {};
    message.addr !== undefined && (obj.addr = message.addr);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryModuleAddressResponse>, I>>(object: I): QueryModuleAddressResponse {
    const message = createBaseQueryModuleAddressResponse();
    message.addr = object.addr ?? "";
    return message;
  },
};

function createBaseQueryGetEpochTrackerRequest(): QueryGetEpochTrackerRequest {
  return { epochIdentifier: "" };
}

export const QueryGetEpochTrackerRequest = {
  encode(message: QueryGetEpochTrackerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.epochIdentifier !== "") {
      writer.uint32(10).string(message.epochIdentifier);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEpochTrackerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEpochTrackerRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.epochIdentifier = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEpochTrackerRequest {
    return { epochIdentifier: isSet(object.epochIdentifier) ? String(object.epochIdentifier) : "" };
  },

  toJSON(message: QueryGetEpochTrackerRequest): unknown {
    const obj: any = {};
    message.epochIdentifier !== undefined && (obj.epochIdentifier = message.epochIdentifier);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEpochTrackerRequest>, I>>(object: I): QueryGetEpochTrackerRequest {
    const message = createBaseQueryGetEpochTrackerRequest();
    message.epochIdentifier = object.epochIdentifier ?? "";
    return message;
  },
};

function createBaseQueryGetEpochTrackerResponse(): QueryGetEpochTrackerResponse {
  return { epochTracker: undefined };
}

export const QueryGetEpochTrackerResponse = {
  encode(message: QueryGetEpochTrackerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.epochTracker !== undefined) {
      EpochTracker.encode(message.epochTracker, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEpochTrackerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEpochTrackerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.epochTracker = EpochTracker.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEpochTrackerResponse {
    return { epochTracker: isSet(object.epochTracker) ? EpochTracker.fromJSON(object.epochTracker) : undefined };
  },

  toJSON(message: QueryGetEpochTrackerResponse): unknown {
    const obj: any = {};
    message.epochTracker !== undefined
      && (obj.epochTracker = message.epochTracker ? EpochTracker.toJSON(message.epochTracker) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEpochTrackerResponse>, I>>(object: I): QueryGetEpochTrackerResponse {
    const message = createBaseQueryGetEpochTrackerResponse();
    message.epochTracker = (object.epochTracker !== undefined && object.epochTracker !== null)
      ? EpochTracker.fromPartial(object.epochTracker)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEpochTrackerRequest(): QueryAllEpochTrackerRequest {
  return { pagination: undefined };
}

export const QueryAllEpochTrackerRequest = {
  encode(message: QueryAllEpochTrackerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEpochTrackerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEpochTrackerRequest();
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

  fromJSON(object: any): QueryAllEpochTrackerRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllEpochTrackerRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEpochTrackerRequest>, I>>(object: I): QueryAllEpochTrackerRequest {
    const message = createBaseQueryAllEpochTrackerRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEpochTrackerResponse(): QueryAllEpochTrackerResponse {
  return { epochTracker: [], pagination: undefined };
}

export const QueryAllEpochTrackerResponse = {
  encode(message: QueryAllEpochTrackerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.epochTracker) {
      EpochTracker.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEpochTrackerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEpochTrackerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.epochTracker.push(EpochTracker.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllEpochTrackerResponse {
    return {
      epochTracker: Array.isArray(object?.epochTracker)
        ? object.epochTracker.map((e: any) => EpochTracker.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllEpochTrackerResponse): unknown {
    const obj: any = {};
    if (message.epochTracker) {
      obj.epochTracker = message.epochTracker.map((e) => e ? EpochTracker.toJSON(e) : undefined);
    } else {
      obj.epochTracker = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEpochTrackerResponse>, I>>(object: I): QueryAllEpochTrackerResponse {
    const message = createBaseQueryAllEpochTrackerResponse();
    message.epochTracker = object.epochTracker?.map((e) => EpochTracker.fromPartial(e)) || [];
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
  /** Queries a Validator by host zone. */
  Validators(request: QueryGetValidatorsRequest): Promise<QueryGetValidatorsResponse>;
  /** Queries a ICAAccount by index. */
  ICAAccount(request: QueryGetICAAccountRequest): Promise<QueryGetICAAccountResponse>;
  /** Queries a HostZone by id. */
  HostZone(request: QueryGetHostZoneRequest): Promise<QueryGetHostZoneResponse>;
  /** Queries a list of HostZone items. */
  HostZoneAll(request: QueryAllHostZoneRequest): Promise<QueryAllHostZoneResponse>;
  /** Queries a list of ModuleAddress items. */
  ModuleAddress(request: QueryModuleAddressRequest): Promise<QueryModuleAddressResponse>;
  /**
   * QueryInterchainAccountFromAddress returns the interchain account for given
   * owner address on a given connection pair
   */
  InterchainAccountFromAddress(
    request: QueryInterchainAccountFromAddressRequest,
  ): Promise<QueryInterchainAccountFromAddressResponse>;
  /** Queries a EpochTracker by index. */
  EpochTracker(request: QueryGetEpochTrackerRequest): Promise<QueryGetEpochTrackerResponse>;
  /** Queries a list of EpochTracker items. */
  EpochTrackerAll(request: QueryAllEpochTrackerRequest): Promise<QueryAllEpochTrackerResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Validators = this.Validators.bind(this);
    this.ICAAccount = this.ICAAccount.bind(this);
    this.HostZone = this.HostZone.bind(this);
    this.HostZoneAll = this.HostZoneAll.bind(this);
    this.ModuleAddress = this.ModuleAddress.bind(this);
    this.InterchainAccountFromAddress = this.InterchainAccountFromAddress.bind(this);
    this.EpochTracker = this.EpochTracker.bind(this);
    this.EpochTrackerAll = this.EpochTrackerAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Validators(request: QueryGetValidatorsRequest): Promise<QueryGetValidatorsResponse> {
    const data = QueryGetValidatorsRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Query", "Validators", data);
    return promise.then((data) => QueryGetValidatorsResponse.decode(new _m0.Reader(data)));
  }

  ICAAccount(request: QueryGetICAAccountRequest): Promise<QueryGetICAAccountResponse> {
    const data = QueryGetICAAccountRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Query", "ICAAccount", data);
    return promise.then((data) => QueryGetICAAccountResponse.decode(new _m0.Reader(data)));
  }

  HostZone(request: QueryGetHostZoneRequest): Promise<QueryGetHostZoneResponse> {
    const data = QueryGetHostZoneRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Query", "HostZone", data);
    return promise.then((data) => QueryGetHostZoneResponse.decode(new _m0.Reader(data)));
  }

  HostZoneAll(request: QueryAllHostZoneRequest): Promise<QueryAllHostZoneResponse> {
    const data = QueryAllHostZoneRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Query", "HostZoneAll", data);
    return promise.then((data) => QueryAllHostZoneResponse.decode(new _m0.Reader(data)));
  }

  ModuleAddress(request: QueryModuleAddressRequest): Promise<QueryModuleAddressResponse> {
    const data = QueryModuleAddressRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Query", "ModuleAddress", data);
    return promise.then((data) => QueryModuleAddressResponse.decode(new _m0.Reader(data)));
  }

  InterchainAccountFromAddress(
    request: QueryInterchainAccountFromAddressRequest,
  ): Promise<QueryInterchainAccountFromAddressResponse> {
    const data = QueryInterchainAccountFromAddressRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Query", "InterchainAccountFromAddress", data);
    return promise.then((data) => QueryInterchainAccountFromAddressResponse.decode(new _m0.Reader(data)));
  }

  EpochTracker(request: QueryGetEpochTrackerRequest): Promise<QueryGetEpochTrackerResponse> {
    const data = QueryGetEpochTrackerRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Query", "EpochTracker", data);
    return promise.then((data) => QueryGetEpochTrackerResponse.decode(new _m0.Reader(data)));
  }

  EpochTrackerAll(request: QueryAllEpochTrackerRequest): Promise<QueryAllEpochTrackerResponse> {
    const data = QueryAllEpochTrackerRequest.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Query", "EpochTrackerAll", data);
    return promise.then((data) => QueryAllEpochTrackerResponse.decode(new _m0.Reader(data)));
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
