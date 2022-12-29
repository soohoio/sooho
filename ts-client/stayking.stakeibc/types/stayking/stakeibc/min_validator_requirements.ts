/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stayking.stakeibc";

export interface MinValidatorRequirements {
  commissionRate: number;
  uptime: number;
}

function createBaseMinValidatorRequirements(): MinValidatorRequirements {
  return { commissionRate: 0, uptime: 0 };
}

export const MinValidatorRequirements = {
  encode(message: MinValidatorRequirements, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.commissionRate !== 0) {
      writer.uint32(8).int32(message.commissionRate);
    }
    if (message.uptime !== 0) {
      writer.uint32(16).int32(message.uptime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MinValidatorRequirements {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMinValidatorRequirements();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.commissionRate = reader.int32();
          break;
        case 2:
          message.uptime = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MinValidatorRequirements {
    return {
      commissionRate: isSet(object.commissionRate) ? Number(object.commissionRate) : 0,
      uptime: isSet(object.uptime) ? Number(object.uptime) : 0,
    };
  },

  toJSON(message: MinValidatorRequirements): unknown {
    const obj: any = {};
    message.commissionRate !== undefined && (obj.commissionRate = Math.round(message.commissionRate));
    message.uptime !== undefined && (obj.uptime = Math.round(message.uptime));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MinValidatorRequirements>, I>>(object: I): MinValidatorRequirements {
    const message = createBaseMinValidatorRequirements();
    message.commissionRate = object.commissionRate ?? 0;
    message.uptime = object.uptime ?? 0;
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
