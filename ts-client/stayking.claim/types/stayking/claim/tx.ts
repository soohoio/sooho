/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "stayking.claim";

export interface MsgSetAirdropAllocations {
  allocator: string;
  airdropIdentifier: string;
  users: string[];
  weights: string[];
}

export interface MsgSetAirdropAllocationsResponse {
}

export interface MsgClaimFreeAmount {
  user: string;
}

export interface MsgClaimFreeAmountResponse {
  claimedAmount: Coin[];
}

export interface MsgCreateAirdrop {
  distributor: string;
  identifier: string;
  startTime: number;
  duration: number;
  denom: string;
}

export interface MsgCreateAirdropResponse {
}

export interface MsgDeleteAirdrop {
  distributor: string;
  identifier: string;
}

export interface MsgDeleteAirdropResponse {
}

function createBaseMsgSetAirdropAllocations(): MsgSetAirdropAllocations {
  return { allocator: "", airdropIdentifier: "", users: [], weights: [] };
}

export const MsgSetAirdropAllocations = {
  encode(message: MsgSetAirdropAllocations, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.allocator !== "") {
      writer.uint32(10).string(message.allocator);
    }
    if (message.airdropIdentifier !== "") {
      writer.uint32(18).string(message.airdropIdentifier);
    }
    for (const v of message.users) {
      writer.uint32(26).string(v!);
    }
    for (const v of message.weights) {
      writer.uint32(34).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetAirdropAllocations {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetAirdropAllocations();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.allocator = reader.string();
          break;
        case 2:
          message.airdropIdentifier = reader.string();
          break;
        case 3:
          message.users.push(reader.string());
          break;
        case 4:
          message.weights.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetAirdropAllocations {
    return {
      allocator: isSet(object.allocator) ? String(object.allocator) : "",
      airdropIdentifier: isSet(object.airdropIdentifier) ? String(object.airdropIdentifier) : "",
      users: Array.isArray(object?.users) ? object.users.map((e: any) => String(e)) : [],
      weights: Array.isArray(object?.weights) ? object.weights.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgSetAirdropAllocations): unknown {
    const obj: any = {};
    message.allocator !== undefined && (obj.allocator = message.allocator);
    message.airdropIdentifier !== undefined && (obj.airdropIdentifier = message.airdropIdentifier);
    if (message.users) {
      obj.users = message.users.map((e) => e);
    } else {
      obj.users = [];
    }
    if (message.weights) {
      obj.weights = message.weights.map((e) => e);
    } else {
      obj.weights = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetAirdropAllocations>, I>>(object: I): MsgSetAirdropAllocations {
    const message = createBaseMsgSetAirdropAllocations();
    message.allocator = object.allocator ?? "";
    message.airdropIdentifier = object.airdropIdentifier ?? "";
    message.users = object.users?.map((e) => e) || [];
    message.weights = object.weights?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgSetAirdropAllocationsResponse(): MsgSetAirdropAllocationsResponse {
  return {};
}

export const MsgSetAirdropAllocationsResponse = {
  encode(_: MsgSetAirdropAllocationsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetAirdropAllocationsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSetAirdropAllocationsResponse();
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

  fromJSON(_: any): MsgSetAirdropAllocationsResponse {
    return {};
  },

  toJSON(_: MsgSetAirdropAllocationsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetAirdropAllocationsResponse>, I>>(
    _: I,
  ): MsgSetAirdropAllocationsResponse {
    const message = createBaseMsgSetAirdropAllocationsResponse();
    return message;
  },
};

function createBaseMsgClaimFreeAmount(): MsgClaimFreeAmount {
  return { user: "" };
}

export const MsgClaimFreeAmount = {
  encode(message: MsgClaimFreeAmount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== "") {
      writer.uint32(10).string(message.user);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimFreeAmount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimFreeAmount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimFreeAmount {
    return { user: isSet(object.user) ? String(object.user) : "" };
  },

  toJSON(message: MsgClaimFreeAmount): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimFreeAmount>, I>>(object: I): MsgClaimFreeAmount {
    const message = createBaseMsgClaimFreeAmount();
    message.user = object.user ?? "";
    return message;
  },
};

function createBaseMsgClaimFreeAmountResponse(): MsgClaimFreeAmountResponse {
  return { claimedAmount: [] };
}

export const MsgClaimFreeAmountResponse = {
  encode(message: MsgClaimFreeAmountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.claimedAmount) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimFreeAmountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimFreeAmountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 3:
          message.claimedAmount.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimFreeAmountResponse {
    return {
      claimedAmount: Array.isArray(object?.claimedAmount) ? object.claimedAmount.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: MsgClaimFreeAmountResponse): unknown {
    const obj: any = {};
    if (message.claimedAmount) {
      obj.claimedAmount = message.claimedAmount.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.claimedAmount = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimFreeAmountResponse>, I>>(object: I): MsgClaimFreeAmountResponse {
    const message = createBaseMsgClaimFreeAmountResponse();
    message.claimedAmount = object.claimedAmount?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMsgCreateAirdrop(): MsgCreateAirdrop {
  return { distributor: "", identifier: "", startTime: 0, duration: 0, denom: "" };
}

export const MsgCreateAirdrop = {
  encode(message: MsgCreateAirdrop, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.distributor !== "") {
      writer.uint32(10).string(message.distributor);
    }
    if (message.identifier !== "") {
      writer.uint32(18).string(message.identifier);
    }
    if (message.startTime !== 0) {
      writer.uint32(24).uint64(message.startTime);
    }
    if (message.duration !== 0) {
      writer.uint32(32).uint64(message.duration);
    }
    if (message.denom !== "") {
      writer.uint32(42).string(message.denom);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAirdrop {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAirdrop();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.distributor = reader.string();
          break;
        case 2:
          message.identifier = reader.string();
          break;
        case 3:
          message.startTime = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.duration = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.denom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAirdrop {
    return {
      distributor: isSet(object.distributor) ? String(object.distributor) : "",
      identifier: isSet(object.identifier) ? String(object.identifier) : "",
      startTime: isSet(object.startTime) ? Number(object.startTime) : 0,
      duration: isSet(object.duration) ? Number(object.duration) : 0,
      denom: isSet(object.denom) ? String(object.denom) : "",
    };
  },

  toJSON(message: MsgCreateAirdrop): unknown {
    const obj: any = {};
    message.distributor !== undefined && (obj.distributor = message.distributor);
    message.identifier !== undefined && (obj.identifier = message.identifier);
    message.startTime !== undefined && (obj.startTime = Math.round(message.startTime));
    message.duration !== undefined && (obj.duration = Math.round(message.duration));
    message.denom !== undefined && (obj.denom = message.denom);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAirdrop>, I>>(object: I): MsgCreateAirdrop {
    const message = createBaseMsgCreateAirdrop();
    message.distributor = object.distributor ?? "";
    message.identifier = object.identifier ?? "";
    message.startTime = object.startTime ?? 0;
    message.duration = object.duration ?? 0;
    message.denom = object.denom ?? "";
    return message;
  },
};

function createBaseMsgCreateAirdropResponse(): MsgCreateAirdropResponse {
  return {};
}

export const MsgCreateAirdropResponse = {
  encode(_: MsgCreateAirdropResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAirdropResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAirdropResponse();
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

  fromJSON(_: any): MsgCreateAirdropResponse {
    return {};
  },

  toJSON(_: MsgCreateAirdropResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAirdropResponse>, I>>(_: I): MsgCreateAirdropResponse {
    const message = createBaseMsgCreateAirdropResponse();
    return message;
  },
};

function createBaseMsgDeleteAirdrop(): MsgDeleteAirdrop {
  return { distributor: "", identifier: "" };
}

export const MsgDeleteAirdrop = {
  encode(message: MsgDeleteAirdrop, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.distributor !== "") {
      writer.uint32(10).string(message.distributor);
    }
    if (message.identifier !== "") {
      writer.uint32(18).string(message.identifier);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteAirdrop {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteAirdrop();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.distributor = reader.string();
          break;
        case 2:
          message.identifier = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteAirdrop {
    return {
      distributor: isSet(object.distributor) ? String(object.distributor) : "",
      identifier: isSet(object.identifier) ? String(object.identifier) : "",
    };
  },

  toJSON(message: MsgDeleteAirdrop): unknown {
    const obj: any = {};
    message.distributor !== undefined && (obj.distributor = message.distributor);
    message.identifier !== undefined && (obj.identifier = message.identifier);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteAirdrop>, I>>(object: I): MsgDeleteAirdrop {
    const message = createBaseMsgDeleteAirdrop();
    message.distributor = object.distributor ?? "";
    message.identifier = object.identifier ?? "";
    return message;
  },
};

function createBaseMsgDeleteAirdropResponse(): MsgDeleteAirdropResponse {
  return {};
}

export const MsgDeleteAirdropResponse = {
  encode(_: MsgDeleteAirdropResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteAirdropResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteAirdropResponse();
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

  fromJSON(_: any): MsgDeleteAirdropResponse {
    return {};
  },

  toJSON(_: MsgDeleteAirdropResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteAirdropResponse>, I>>(_: I): MsgDeleteAirdropResponse {
    const message = createBaseMsgDeleteAirdropResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  SetAirdropAllocations(request: MsgSetAirdropAllocations): Promise<MsgSetAirdropAllocationsResponse>;
  ClaimFreeAmount(request: MsgClaimFreeAmount): Promise<MsgClaimFreeAmountResponse>;
  CreateAirdrop(request: MsgCreateAirdrop): Promise<MsgCreateAirdropResponse>;
  DeleteAirdrop(request: MsgDeleteAirdrop): Promise<MsgDeleteAirdropResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.SetAirdropAllocations = this.SetAirdropAllocations.bind(this);
    this.ClaimFreeAmount = this.ClaimFreeAmount.bind(this);
    this.CreateAirdrop = this.CreateAirdrop.bind(this);
    this.DeleteAirdrop = this.DeleteAirdrop.bind(this);
  }
  SetAirdropAllocations(request: MsgSetAirdropAllocations): Promise<MsgSetAirdropAllocationsResponse> {
    const data = MsgSetAirdropAllocations.encode(request).finish();
    const promise = this.rpc.request("stayking.claim.Msg", "SetAirdropAllocations", data);
    return promise.then((data) => MsgSetAirdropAllocationsResponse.decode(new _m0.Reader(data)));
  }

  ClaimFreeAmount(request: MsgClaimFreeAmount): Promise<MsgClaimFreeAmountResponse> {
    const data = MsgClaimFreeAmount.encode(request).finish();
    const promise = this.rpc.request("stayking.claim.Msg", "ClaimFreeAmount", data);
    return promise.then((data) => MsgClaimFreeAmountResponse.decode(new _m0.Reader(data)));
  }

  CreateAirdrop(request: MsgCreateAirdrop): Promise<MsgCreateAirdropResponse> {
    const data = MsgCreateAirdrop.encode(request).finish();
    const promise = this.rpc.request("stayking.claim.Msg", "CreateAirdrop", data);
    return promise.then((data) => MsgCreateAirdropResponse.decode(new _m0.Reader(data)));
  }

  DeleteAirdrop(request: MsgDeleteAirdrop): Promise<MsgDeleteAirdropResponse> {
    const data = MsgDeleteAirdrop.encode(request).finish();
    const promise = this.rpc.request("stayking.claim.Msg", "DeleteAirdrop", data);
    return promise.then((data) => MsgDeleteAirdropResponse.decode(new _m0.Reader(data)));
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
