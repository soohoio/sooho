/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stayking.stakeibc";

export interface AddValidatorProposal {
  title: string;
  description: string;
  hostZone: string;
  validatorName: string;
  validatorAddress: string;
  deposit: string;
}

function createBaseAddValidatorProposal(): AddValidatorProposal {
  return { title: "", description: "", hostZone: "", validatorName: "", validatorAddress: "", deposit: "" };
}

export const AddValidatorProposal = {
  encode(message: AddValidatorProposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.hostZone !== "") {
      writer.uint32(26).string(message.hostZone);
    }
    if (message.validatorName !== "") {
      writer.uint32(34).string(message.validatorName);
    }
    if (message.validatorAddress !== "") {
      writer.uint32(42).string(message.validatorAddress);
    }
    if (message.deposit !== "") {
      writer.uint32(50).string(message.deposit);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddValidatorProposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddValidatorProposal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.description = reader.string();
          break;
        case 3:
          message.hostZone = reader.string();
          break;
        case 4:
          message.validatorName = reader.string();
          break;
        case 5:
          message.validatorAddress = reader.string();
          break;
        case 6:
          message.deposit = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AddValidatorProposal {
    return {
      title: isSet(object.title) ? String(object.title) : "",
      description: isSet(object.description) ? String(object.description) : "",
      hostZone: isSet(object.hostZone) ? String(object.hostZone) : "",
      validatorName: isSet(object.validatorName) ? String(object.validatorName) : "",
      validatorAddress: isSet(object.validatorAddress) ? String(object.validatorAddress) : "",
      deposit: isSet(object.deposit) ? String(object.deposit) : "",
    };
  },

  toJSON(message: AddValidatorProposal): unknown {
    const obj: any = {};
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined && (obj.description = message.description);
    message.hostZone !== undefined && (obj.hostZone = message.hostZone);
    message.validatorName !== undefined && (obj.validatorName = message.validatorName);
    message.validatorAddress !== undefined && (obj.validatorAddress = message.validatorAddress);
    message.deposit !== undefined && (obj.deposit = message.deposit);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AddValidatorProposal>, I>>(object: I): AddValidatorProposal {
    const message = createBaseAddValidatorProposal();
    message.title = object.title ?? "";
    message.description = object.description ?? "";
    message.hostZone = object.hostZone ?? "";
    message.validatorName = object.validatorName ?? "";
    message.validatorAddress = object.validatorAddress ?? "";
    message.deposit = object.deposit ?? "";
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
