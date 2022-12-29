/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Duration } from "../../google/protobuf/duration";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "stayking.claim";

/** Params defines the claim module's parameters. */
export interface Params {
  airdrops: Airdrop[];
}

export interface Airdrop {
  airdropIdentifier: string;
  /** seconds */
  airdropStartTime:
    | Date
    | undefined;
  /** seconds */
  airdropDuration:
    | Duration
    | undefined;
  /** denom of claimable asset */
  claimDenom: string;
  /** airdrop distribution account */
  distributorAddress: string;
  /** ustrd tokens claimed so far in the current period */
  claimedSoFar: string;
}

function createBaseParams(): Params {
  return { airdrops: [] };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.airdrops) {
      Airdrop.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.airdrops.push(Airdrop.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    return { airdrops: Array.isArray(object?.airdrops) ? object.airdrops.map((e: any) => Airdrop.fromJSON(e)) : [] };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    if (message.airdrops) {
      obj.airdrops = message.airdrops.map((e) => e ? Airdrop.toJSON(e) : undefined);
    } else {
      obj.airdrops = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.airdrops = object.airdrops?.map((e) => Airdrop.fromPartial(e)) || [];
    return message;
  },
};

function createBaseAirdrop(): Airdrop {
  return {
    airdropIdentifier: "",
    airdropStartTime: undefined,
    airdropDuration: undefined,
    claimDenom: "",
    distributorAddress: "",
    claimedSoFar: "",
  };
}

export const Airdrop = {
  encode(message: Airdrop, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.airdropIdentifier !== "") {
      writer.uint32(10).string(message.airdropIdentifier);
    }
    if (message.airdropStartTime !== undefined) {
      Timestamp.encode(toTimestamp(message.airdropStartTime), writer.uint32(18).fork()).ldelim();
    }
    if (message.airdropDuration !== undefined) {
      Duration.encode(message.airdropDuration, writer.uint32(26).fork()).ldelim();
    }
    if (message.claimDenom !== "") {
      writer.uint32(34).string(message.claimDenom);
    }
    if (message.distributorAddress !== "") {
      writer.uint32(42).string(message.distributorAddress);
    }
    if (message.claimedSoFar !== "") {
      writer.uint32(50).string(message.claimedSoFar);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Airdrop {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAirdrop();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.airdropIdentifier = reader.string();
          break;
        case 2:
          message.airdropStartTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 3:
          message.airdropDuration = Duration.decode(reader, reader.uint32());
          break;
        case 4:
          message.claimDenom = reader.string();
          break;
        case 5:
          message.distributorAddress = reader.string();
          break;
        case 6:
          message.claimedSoFar = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Airdrop {
    return {
      airdropIdentifier: isSet(object.airdropIdentifier) ? String(object.airdropIdentifier) : "",
      airdropStartTime: isSet(object.airdropStartTime) ? fromJsonTimestamp(object.airdropStartTime) : undefined,
      airdropDuration: isSet(object.airdropDuration) ? Duration.fromJSON(object.airdropDuration) : undefined,
      claimDenom: isSet(object.claimDenom) ? String(object.claimDenom) : "",
      distributorAddress: isSet(object.distributorAddress) ? String(object.distributorAddress) : "",
      claimedSoFar: isSet(object.claimedSoFar) ? String(object.claimedSoFar) : "",
    };
  },

  toJSON(message: Airdrop): unknown {
    const obj: any = {};
    message.airdropIdentifier !== undefined && (obj.airdropIdentifier = message.airdropIdentifier);
    message.airdropStartTime !== undefined && (obj.airdropStartTime = message.airdropStartTime.toISOString());
    message.airdropDuration !== undefined
      && (obj.airdropDuration = message.airdropDuration ? Duration.toJSON(message.airdropDuration) : undefined);
    message.claimDenom !== undefined && (obj.claimDenom = message.claimDenom);
    message.distributorAddress !== undefined && (obj.distributorAddress = message.distributorAddress);
    message.claimedSoFar !== undefined && (obj.claimedSoFar = message.claimedSoFar);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Airdrop>, I>>(object: I): Airdrop {
    const message = createBaseAirdrop();
    message.airdropIdentifier = object.airdropIdentifier ?? "";
    message.airdropStartTime = object.airdropStartTime ?? undefined;
    message.airdropDuration = (object.airdropDuration !== undefined && object.airdropDuration !== null)
      ? Duration.fromPartial(object.airdropDuration)
      : undefined;
    message.claimDenom = object.claimDenom ?? "";
    message.distributorAddress = object.distributorAddress ?? "";
    message.claimedSoFar = object.claimedSoFar ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
