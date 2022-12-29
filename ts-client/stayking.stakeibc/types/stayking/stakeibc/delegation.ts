/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Validator } from "./validator";

export const protobufPackage = "stayking.stakeibc";

export interface Delegation {
  delegateAcctAddress: string;
  validator: Validator | undefined;
  amt: string;
}

function createBaseDelegation(): Delegation {
  return { delegateAcctAddress: "", validator: undefined, amt: "" };
}

export const Delegation = {
  encode(message: Delegation, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.delegateAcctAddress !== "") {
      writer.uint32(10).string(message.delegateAcctAddress);
    }
    if (message.validator !== undefined) {
      Validator.encode(message.validator, writer.uint32(18).fork()).ldelim();
    }
    if (message.amt !== "") {
      writer.uint32(26).string(message.amt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Delegation {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDelegation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.delegateAcctAddress = reader.string();
          break;
        case 2:
          message.validator = Validator.decode(reader, reader.uint32());
          break;
        case 3:
          message.amt = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Delegation {
    return {
      delegateAcctAddress: isSet(object.delegateAcctAddress) ? String(object.delegateAcctAddress) : "",
      validator: isSet(object.validator) ? Validator.fromJSON(object.validator) : undefined,
      amt: isSet(object.amt) ? String(object.amt) : "",
    };
  },

  toJSON(message: Delegation): unknown {
    const obj: any = {};
    message.delegateAcctAddress !== undefined && (obj.delegateAcctAddress = message.delegateAcctAddress);
    message.validator !== undefined
      && (obj.validator = message.validator ? Validator.toJSON(message.validator) : undefined);
    message.amt !== undefined && (obj.amt = message.amt);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Delegation>, I>>(object: I): Delegation {
    const message = createBaseDelegation();
    message.delegateAcctAddress = object.delegateAcctAddress ?? "";
    message.validator = (object.validator !== undefined && object.validator !== null)
      ? Validator.fromPartial(object.validator)
      : undefined;
    message.amt = object.amt ?? "";
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
