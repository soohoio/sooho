/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stayking.claim";

export enum Action {
  ACTION_FREE = 0,
  ACTION_LIQUID_STAKE = 1,
  ACTION_DELEGATE_STAKE = 2,
  UNRECOGNIZED = -1,
}

export function actionFromJSON(object: any): Action {
  switch (object) {
    case 0:
    case "ACTION_FREE":
      return Action.ACTION_FREE;
    case 1:
    case "ACTION_LIQUID_STAKE":
      return Action.ACTION_LIQUID_STAKE;
    case 2:
    case "ACTION_DELEGATE_STAKE":
      return Action.ACTION_DELEGATE_STAKE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Action.UNRECOGNIZED;
  }
}

export function actionToJSON(object: Action): string {
  switch (object) {
    case Action.ACTION_FREE:
      return "ACTION_FREE";
    case Action.ACTION_LIQUID_STAKE:
      return "ACTION_LIQUID_STAKE";
    case Action.ACTION_DELEGATE_STAKE:
      return "ACTION_DELEGATE_STAKE";
    case Action.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

/** A Claim Records is the metadata of claim data per address */
export interface ClaimRecord {
  /** airdrop identifier */
  airdropIdentifier: string;
  /** address of claim user */
  address: string;
  /** weight that represent the portion from total allocation */
  weight: string;
  /**
   * true if action is completed
   * index of bool in array refers to action enum #
   */
  actionCompleted: boolean[];
}

function createBaseClaimRecord(): ClaimRecord {
  return { airdropIdentifier: "", address: "", weight: "", actionCompleted: [] };
}

export const ClaimRecord = {
  encode(message: ClaimRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.airdropIdentifier !== "") {
      writer.uint32(10).string(message.airdropIdentifier);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.weight !== "") {
      writer.uint32(26).string(message.weight);
    }
    writer.uint32(34).fork();
    for (const v of message.actionCompleted) {
      writer.bool(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ClaimRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseClaimRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.airdropIdentifier = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.weight = reader.string();
          break;
        case 4:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.actionCompleted.push(reader.bool());
            }
          } else {
            message.actionCompleted.push(reader.bool());
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ClaimRecord {
    return {
      airdropIdentifier: isSet(object.airdropIdentifier) ? String(object.airdropIdentifier) : "",
      address: isSet(object.address) ? String(object.address) : "",
      weight: isSet(object.weight) ? String(object.weight) : "",
      actionCompleted: Array.isArray(object?.actionCompleted) ? object.actionCompleted.map((e: any) => Boolean(e)) : [],
    };
  },

  toJSON(message: ClaimRecord): unknown {
    const obj: any = {};
    message.airdropIdentifier !== undefined && (obj.airdropIdentifier = message.airdropIdentifier);
    message.address !== undefined && (obj.address = message.address);
    message.weight !== undefined && (obj.weight = message.weight);
    if (message.actionCompleted) {
      obj.actionCompleted = message.actionCompleted.map((e) => e);
    } else {
      obj.actionCompleted = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ClaimRecord>, I>>(object: I): ClaimRecord {
    const message = createBaseClaimRecord();
    message.airdropIdentifier = object.airdropIdentifier ?? "";
    message.address = object.address ?? "";
    message.weight = object.weight ?? "";
    message.actionCompleted = object.actionCompleted?.map((e) => e) || [];
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
