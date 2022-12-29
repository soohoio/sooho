/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { EpochTracker } from "./epoch_tracker";
import { HostZone } from "./host_zone";
import { ICAAccount } from "./ica_account";
import { Params } from "./params";

export const protobufPackage = "stayking.stakeibc";

/** GenesisState defines the stakeibc module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  portId: string;
  /** list of zones that are registered by the protocol */
  icaAccount: ICAAccount | undefined;
  hostZoneList: HostZone[];
  hostZoneCount: number;
  /** stores a map from hostZone base denom to hostZone */
  denomToHostZone: { [key: string]: string };
  epochTrackerList: EpochTracker[];
}

export interface GenesisState_DenomToHostZoneEntry {
  key: string;
  value: string;
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    portId: "",
    icaAccount: undefined,
    hostZoneList: [],
    hostZoneCount: 0,
    denomToHostZone: {},
    epochTrackerList: [],
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.portId !== "") {
      writer.uint32(18).string(message.portId);
    }
    if (message.icaAccount !== undefined) {
      ICAAccount.encode(message.icaAccount, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.hostZoneList) {
      HostZone.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.hostZoneCount !== 0) {
      writer.uint32(48).uint64(message.hostZoneCount);
    }
    Object.entries(message.denomToHostZone).forEach(([key, value]) => {
      GenesisState_DenomToHostZoneEntry.encode({ key: key as any, value }, writer.uint32(74).fork()).ldelim();
    });
    for (const v of message.epochTrackerList) {
      EpochTracker.encode(v!, writer.uint32(82).fork()).ldelim();
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
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.portId = reader.string();
          break;
        case 4:
          message.icaAccount = ICAAccount.decode(reader, reader.uint32());
          break;
        case 5:
          message.hostZoneList.push(HostZone.decode(reader, reader.uint32()));
          break;
        case 6:
          message.hostZoneCount = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          const entry9 = GenesisState_DenomToHostZoneEntry.decode(reader, reader.uint32());
          if (entry9.value !== undefined) {
            message.denomToHostZone[entry9.key] = entry9.value;
          }
          break;
        case 10:
          message.epochTrackerList.push(EpochTracker.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      portId: isSet(object.portId) ? String(object.portId) : "",
      icaAccount: isSet(object.icaAccount) ? ICAAccount.fromJSON(object.icaAccount) : undefined,
      hostZoneList: Array.isArray(object?.hostZoneList)
        ? object.hostZoneList.map((e: any) => HostZone.fromJSON(e))
        : [],
      hostZoneCount: isSet(object.hostZoneCount) ? Number(object.hostZoneCount) : 0,
      denomToHostZone: isObject(object.denomToHostZone)
        ? Object.entries(object.denomToHostZone).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
      epochTrackerList: Array.isArray(object?.epochTrackerList)
        ? object.epochTrackerList.map((e: any) => EpochTracker.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.portId !== undefined && (obj.portId = message.portId);
    message.icaAccount !== undefined
      && (obj.icaAccount = message.icaAccount ? ICAAccount.toJSON(message.icaAccount) : undefined);
    if (message.hostZoneList) {
      obj.hostZoneList = message.hostZoneList.map((e) => e ? HostZone.toJSON(e) : undefined);
    } else {
      obj.hostZoneList = [];
    }
    message.hostZoneCount !== undefined && (obj.hostZoneCount = Math.round(message.hostZoneCount));
    obj.denomToHostZone = {};
    if (message.denomToHostZone) {
      Object.entries(message.denomToHostZone).forEach(([k, v]) => {
        obj.denomToHostZone[k] = v;
      });
    }
    if (message.epochTrackerList) {
      obj.epochTrackerList = message.epochTrackerList.map((e) => e ? EpochTracker.toJSON(e) : undefined);
    } else {
      obj.epochTrackerList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.portId = object.portId ?? "";
    message.icaAccount = (object.icaAccount !== undefined && object.icaAccount !== null)
      ? ICAAccount.fromPartial(object.icaAccount)
      : undefined;
    message.hostZoneList = object.hostZoneList?.map((e) => HostZone.fromPartial(e)) || [];
    message.hostZoneCount = object.hostZoneCount ?? 0;
    message.denomToHostZone = Object.entries(object.denomToHostZone ?? {}).reduce<{ [key: string]: string }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = String(value);
        }
        return acc;
      },
      {},
    );
    message.epochTrackerList = object.epochTrackerList?.map((e) => EpochTracker.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGenesisState_DenomToHostZoneEntry(): GenesisState_DenomToHostZoneEntry {
  return { key: "", value: "" };
}

export const GenesisState_DenomToHostZoneEntry = {
  encode(message: GenesisState_DenomToHostZoneEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState_DenomToHostZoneEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState_DenomToHostZoneEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState_DenomToHostZoneEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? String(object.value) : "" };
  },

  toJSON(message: GenesisState_DenomToHostZoneEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState_DenomToHostZoneEntry>, I>>(
    object: I,
  ): GenesisState_DenomToHostZoneEntry {
    const message = createBaseGenesisState_DenomToHostZoneEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
