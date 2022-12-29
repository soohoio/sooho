/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { ICAAccountType, iCAAccountTypeFromJSON, iCAAccountTypeToJSON } from "./ica_account";

export const protobufPackage = "stayking.stakeibc";

export interface MsgLiquidStake {
  creator: string;
  amount: string;
  /** TODO(TEST-86): Update Denom -> HostDenom */
  hostDenom: string;
}

export interface MsgLiquidStakeResponse {
}

export interface MsgClearBalance {
  creator: string;
  chainId: string;
  amount: string;
  channel: string;
}

export interface MsgClearBalanceResponse {
}

export interface MsgRedeemStake {
  creator: string;
  amount: string;
  hostZone: string;
  receiver: string;
}

export interface MsgRedeemStakeResponse {
}

/** next: 13 */
export interface MsgRegisterHostZone {
  connectionId: string;
  bech32prefix: string;
  hostDenom: string;
  ibcDenom: string;
  creator: string;
  transferChannelId: string;
  unbondingFrequency: number;
}

/**
 * TODO(TEST-53): Remove this pre-launch (no need for clients to create /
 * interact with ICAs)
 */
export interface MsgRegisterHostZoneResponse {
}

export interface MsgClaimUndelegatedTokens {
  creator: string;
  /** UserUnbondingRecords are keyed on {chain_id}.{epoch}.{sender} */
  hostZoneId: string;
  epoch: number;
  sender: string;
}

export interface MsgClaimUndelegatedTokensResponse {
}

export interface MsgRebalanceValidators {
  creator: string;
  hostZone: string;
  numRebalance: number;
}

export interface MsgRebalanceValidatorsResponse {
}

export interface MsgAddValidator {
  creator: string;
  hostZone: string;
  name: string;
  address: string;
  commission: number;
  weight: number;
}

export interface MsgAddValidatorResponse {
}

export interface MsgChangeValidatorWeight {
  creator: string;
  hostZone: string;
  valAddr: string;
  weight: number;
}

export interface MsgChangeValidatorWeightResponse {
}

export interface MsgDeleteValidator {
  creator: string;
  hostZone: string;
  valAddr: string;
}

export interface MsgDeleteValidatorResponse {
}

export interface MsgRestoreInterchainAccount {
  creator: string;
  chainId: string;
  accountType: ICAAccountType;
}

export interface MsgRestoreInterchainAccountResponse {
}

export interface MsgUpdateValidatorSharesExchRate {
  creator: string;
  chainId: string;
  valoper: string;
}

export interface MsgUpdateValidatorSharesExchRateResponse {
}

function createBaseMsgLiquidStake(): MsgLiquidStake {
  return { creator: "", amount: "", hostDenom: "" };
}

export const MsgLiquidStake = {
  encode(message: MsgLiquidStake, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    if (message.hostDenom !== "") {
      writer.uint32(26).string(message.hostDenom);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgLiquidStake {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgLiquidStake();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        case 3:
          message.hostDenom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgLiquidStake {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      hostDenom: isSet(object.hostDenom) ? String(object.hostDenom) : "",
    };
  },

  toJSON(message: MsgLiquidStake): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    message.hostDenom !== undefined && (obj.hostDenom = message.hostDenom);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgLiquidStake>, I>>(object: I): MsgLiquidStake {
    const message = createBaseMsgLiquidStake();
    message.creator = object.creator ?? "";
    message.amount = object.amount ?? "";
    message.hostDenom = object.hostDenom ?? "";
    return message;
  },
};

function createBaseMsgLiquidStakeResponse(): MsgLiquidStakeResponse {
  return {};
}

export const MsgLiquidStakeResponse = {
  encode(_: MsgLiquidStakeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgLiquidStakeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgLiquidStakeResponse();
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

  fromJSON(_: any): MsgLiquidStakeResponse {
    return {};
  },

  toJSON(_: MsgLiquidStakeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgLiquidStakeResponse>, I>>(_: I): MsgLiquidStakeResponse {
    const message = createBaseMsgLiquidStakeResponse();
    return message;
  },
};

function createBaseMsgClearBalance(): MsgClearBalance {
  return { creator: "", chainId: "", amount: "", channel: "" };
}

export const MsgClearBalance = {
  encode(message: MsgClearBalance, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.chainId !== "") {
      writer.uint32(18).string(message.chainId);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    if (message.channel !== "") {
      writer.uint32(34).string(message.channel);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClearBalance {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClearBalance();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.chainId = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        case 4:
          message.channel = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClearBalance {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      chainId: isSet(object.chainId) ? String(object.chainId) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      channel: isSet(object.channel) ? String(object.channel) : "",
    };
  },

  toJSON(message: MsgClearBalance): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.amount !== undefined && (obj.amount = message.amount);
    message.channel !== undefined && (obj.channel = message.channel);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClearBalance>, I>>(object: I): MsgClearBalance {
    const message = createBaseMsgClearBalance();
    message.creator = object.creator ?? "";
    message.chainId = object.chainId ?? "";
    message.amount = object.amount ?? "";
    message.channel = object.channel ?? "";
    return message;
  },
};

function createBaseMsgClearBalanceResponse(): MsgClearBalanceResponse {
  return {};
}

export const MsgClearBalanceResponse = {
  encode(_: MsgClearBalanceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClearBalanceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClearBalanceResponse();
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

  fromJSON(_: any): MsgClearBalanceResponse {
    return {};
  },

  toJSON(_: MsgClearBalanceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClearBalanceResponse>, I>>(_: I): MsgClearBalanceResponse {
    const message = createBaseMsgClearBalanceResponse();
    return message;
  },
};

function createBaseMsgRedeemStake(): MsgRedeemStake {
  return { creator: "", amount: "", hostZone: "", receiver: "" };
}

export const MsgRedeemStake = {
  encode(message: MsgRedeemStake, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    if (message.hostZone !== "") {
      writer.uint32(26).string(message.hostZone);
    }
    if (message.receiver !== "") {
      writer.uint32(34).string(message.receiver);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRedeemStake {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRedeemStake();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        case 3:
          message.hostZone = reader.string();
          break;
        case 4:
          message.receiver = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRedeemStake {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      hostZone: isSet(object.hostZone) ? String(object.hostZone) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
    };
  },

  toJSON(message: MsgRedeemStake): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    message.hostZone !== undefined && (obj.hostZone = message.hostZone);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRedeemStake>, I>>(object: I): MsgRedeemStake {
    const message = createBaseMsgRedeemStake();
    message.creator = object.creator ?? "";
    message.amount = object.amount ?? "";
    message.hostZone = object.hostZone ?? "";
    message.receiver = object.receiver ?? "";
    return message;
  },
};

function createBaseMsgRedeemStakeResponse(): MsgRedeemStakeResponse {
  return {};
}

export const MsgRedeemStakeResponse = {
  encode(_: MsgRedeemStakeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRedeemStakeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRedeemStakeResponse();
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

  fromJSON(_: any): MsgRedeemStakeResponse {
    return {};
  },

  toJSON(_: MsgRedeemStakeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRedeemStakeResponse>, I>>(_: I): MsgRedeemStakeResponse {
    const message = createBaseMsgRedeemStakeResponse();
    return message;
  },
};

function createBaseMsgRegisterHostZone(): MsgRegisterHostZone {
  return {
    connectionId: "",
    bech32prefix: "",
    hostDenom: "",
    ibcDenom: "",
    creator: "",
    transferChannelId: "",
    unbondingFrequency: 0,
  };
}

export const MsgRegisterHostZone = {
  encode(message: MsgRegisterHostZone, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.connectionId !== "") {
      writer.uint32(18).string(message.connectionId);
    }
    if (message.bech32prefix !== "") {
      writer.uint32(98).string(message.bech32prefix);
    }
    if (message.hostDenom !== "") {
      writer.uint32(34).string(message.hostDenom);
    }
    if (message.ibcDenom !== "") {
      writer.uint32(42).string(message.ibcDenom);
    }
    if (message.creator !== "") {
      writer.uint32(50).string(message.creator);
    }
    if (message.transferChannelId !== "") {
      writer.uint32(82).string(message.transferChannelId);
    }
    if (message.unbondingFrequency !== 0) {
      writer.uint32(88).uint64(message.unbondingFrequency);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterHostZone {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterHostZone();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.connectionId = reader.string();
          break;
        case 12:
          message.bech32prefix = reader.string();
          break;
        case 4:
          message.hostDenom = reader.string();
          break;
        case 5:
          message.ibcDenom = reader.string();
          break;
        case 6:
          message.creator = reader.string();
          break;
        case 10:
          message.transferChannelId = reader.string();
          break;
        case 11:
          message.unbondingFrequency = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterHostZone {
    return {
      connectionId: isSet(object.connectionId) ? String(object.connectionId) : "",
      bech32prefix: isSet(object.bech32prefix) ? String(object.bech32prefix) : "",
      hostDenom: isSet(object.hostDenom) ? String(object.hostDenom) : "",
      ibcDenom: isSet(object.ibcDenom) ? String(object.ibcDenom) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      transferChannelId: isSet(object.transferChannelId) ? String(object.transferChannelId) : "",
      unbondingFrequency: isSet(object.unbondingFrequency) ? Number(object.unbondingFrequency) : 0,
    };
  },

  toJSON(message: MsgRegisterHostZone): unknown {
    const obj: any = {};
    message.connectionId !== undefined && (obj.connectionId = message.connectionId);
    message.bech32prefix !== undefined && (obj.bech32prefix = message.bech32prefix);
    message.hostDenom !== undefined && (obj.hostDenom = message.hostDenom);
    message.ibcDenom !== undefined && (obj.ibcDenom = message.ibcDenom);
    message.creator !== undefined && (obj.creator = message.creator);
    message.transferChannelId !== undefined && (obj.transferChannelId = message.transferChannelId);
    message.unbondingFrequency !== undefined && (obj.unbondingFrequency = Math.round(message.unbondingFrequency));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterHostZone>, I>>(object: I): MsgRegisterHostZone {
    const message = createBaseMsgRegisterHostZone();
    message.connectionId = object.connectionId ?? "";
    message.bech32prefix = object.bech32prefix ?? "";
    message.hostDenom = object.hostDenom ?? "";
    message.ibcDenom = object.ibcDenom ?? "";
    message.creator = object.creator ?? "";
    message.transferChannelId = object.transferChannelId ?? "";
    message.unbondingFrequency = object.unbondingFrequency ?? 0;
    return message;
  },
};

function createBaseMsgRegisterHostZoneResponse(): MsgRegisterHostZoneResponse {
  return {};
}

export const MsgRegisterHostZoneResponse = {
  encode(_: MsgRegisterHostZoneResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterHostZoneResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterHostZoneResponse();
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

  fromJSON(_: any): MsgRegisterHostZoneResponse {
    return {};
  },

  toJSON(_: MsgRegisterHostZoneResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterHostZoneResponse>, I>>(_: I): MsgRegisterHostZoneResponse {
    const message = createBaseMsgRegisterHostZoneResponse();
    return message;
  },
};

function createBaseMsgClaimUndelegatedTokens(): MsgClaimUndelegatedTokens {
  return { creator: "", hostZoneId: "", epoch: 0, sender: "" };
}

export const MsgClaimUndelegatedTokens = {
  encode(message: MsgClaimUndelegatedTokens, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.hostZoneId !== "") {
      writer.uint32(18).string(message.hostZoneId);
    }
    if (message.epoch !== 0) {
      writer.uint32(24).uint64(message.epoch);
    }
    if (message.sender !== "") {
      writer.uint32(34).string(message.sender);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimUndelegatedTokens {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimUndelegatedTokens();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.hostZoneId = reader.string();
          break;
        case 3:
          message.epoch = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.sender = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimUndelegatedTokens {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      hostZoneId: isSet(object.hostZoneId) ? String(object.hostZoneId) : "",
      epoch: isSet(object.epoch) ? Number(object.epoch) : 0,
      sender: isSet(object.sender) ? String(object.sender) : "",
    };
  },

  toJSON(message: MsgClaimUndelegatedTokens): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.hostZoneId !== undefined && (obj.hostZoneId = message.hostZoneId);
    message.epoch !== undefined && (obj.epoch = Math.round(message.epoch));
    message.sender !== undefined && (obj.sender = message.sender);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimUndelegatedTokens>, I>>(object: I): MsgClaimUndelegatedTokens {
    const message = createBaseMsgClaimUndelegatedTokens();
    message.creator = object.creator ?? "";
    message.hostZoneId = object.hostZoneId ?? "";
    message.epoch = object.epoch ?? 0;
    message.sender = object.sender ?? "";
    return message;
  },
};

function createBaseMsgClaimUndelegatedTokensResponse(): MsgClaimUndelegatedTokensResponse {
  return {};
}

export const MsgClaimUndelegatedTokensResponse = {
  encode(_: MsgClaimUndelegatedTokensResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimUndelegatedTokensResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimUndelegatedTokensResponse();
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

  fromJSON(_: any): MsgClaimUndelegatedTokensResponse {
    return {};
  },

  toJSON(_: MsgClaimUndelegatedTokensResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimUndelegatedTokensResponse>, I>>(
    _: I,
  ): MsgClaimUndelegatedTokensResponse {
    const message = createBaseMsgClaimUndelegatedTokensResponse();
    return message;
  },
};

function createBaseMsgRebalanceValidators(): MsgRebalanceValidators {
  return { creator: "", hostZone: "", numRebalance: 0 };
}

export const MsgRebalanceValidators = {
  encode(message: MsgRebalanceValidators, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.hostZone !== "") {
      writer.uint32(18).string(message.hostZone);
    }
    if (message.numRebalance !== 0) {
      writer.uint32(24).uint64(message.numRebalance);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRebalanceValidators {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRebalanceValidators();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.hostZone = reader.string();
          break;
        case 3:
          message.numRebalance = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRebalanceValidators {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      hostZone: isSet(object.hostZone) ? String(object.hostZone) : "",
      numRebalance: isSet(object.numRebalance) ? Number(object.numRebalance) : 0,
    };
  },

  toJSON(message: MsgRebalanceValidators): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.hostZone !== undefined && (obj.hostZone = message.hostZone);
    message.numRebalance !== undefined && (obj.numRebalance = Math.round(message.numRebalance));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRebalanceValidators>, I>>(object: I): MsgRebalanceValidators {
    const message = createBaseMsgRebalanceValidators();
    message.creator = object.creator ?? "";
    message.hostZone = object.hostZone ?? "";
    message.numRebalance = object.numRebalance ?? 0;
    return message;
  },
};

function createBaseMsgRebalanceValidatorsResponse(): MsgRebalanceValidatorsResponse {
  return {};
}

export const MsgRebalanceValidatorsResponse = {
  encode(_: MsgRebalanceValidatorsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRebalanceValidatorsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRebalanceValidatorsResponse();
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

  fromJSON(_: any): MsgRebalanceValidatorsResponse {
    return {};
  },

  toJSON(_: MsgRebalanceValidatorsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRebalanceValidatorsResponse>, I>>(_: I): MsgRebalanceValidatorsResponse {
    const message = createBaseMsgRebalanceValidatorsResponse();
    return message;
  },
};

function createBaseMsgAddValidator(): MsgAddValidator {
  return { creator: "", hostZone: "", name: "", address: "", commission: 0, weight: 0 };
}

export const MsgAddValidator = {
  encode(message: MsgAddValidator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.hostZone !== "") {
      writer.uint32(18).string(message.hostZone);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.address !== "") {
      writer.uint32(34).string(message.address);
    }
    if (message.commission !== 0) {
      writer.uint32(40).uint64(message.commission);
    }
    if (message.weight !== 0) {
      writer.uint32(48).uint64(message.weight);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddValidator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddValidator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.hostZone = reader.string();
          break;
        case 3:
          message.name = reader.string();
          break;
        case 4:
          message.address = reader.string();
          break;
        case 5:
          message.commission = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.weight = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddValidator {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      hostZone: isSet(object.hostZone) ? String(object.hostZone) : "",
      name: isSet(object.name) ? String(object.name) : "",
      address: isSet(object.address) ? String(object.address) : "",
      commission: isSet(object.commission) ? Number(object.commission) : 0,
      weight: isSet(object.weight) ? Number(object.weight) : 0,
    };
  },

  toJSON(message: MsgAddValidator): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.hostZone !== undefined && (obj.hostZone = message.hostZone);
    message.name !== undefined && (obj.name = message.name);
    message.address !== undefined && (obj.address = message.address);
    message.commission !== undefined && (obj.commission = Math.round(message.commission));
    message.weight !== undefined && (obj.weight = Math.round(message.weight));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddValidator>, I>>(object: I): MsgAddValidator {
    const message = createBaseMsgAddValidator();
    message.creator = object.creator ?? "";
    message.hostZone = object.hostZone ?? "";
    message.name = object.name ?? "";
    message.address = object.address ?? "";
    message.commission = object.commission ?? 0;
    message.weight = object.weight ?? 0;
    return message;
  },
};

function createBaseMsgAddValidatorResponse(): MsgAddValidatorResponse {
  return {};
}

export const MsgAddValidatorResponse = {
  encode(_: MsgAddValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddValidatorResponse();
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

  fromJSON(_: any): MsgAddValidatorResponse {
    return {};
  },

  toJSON(_: MsgAddValidatorResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddValidatorResponse>, I>>(_: I): MsgAddValidatorResponse {
    const message = createBaseMsgAddValidatorResponse();
    return message;
  },
};

function createBaseMsgChangeValidatorWeight(): MsgChangeValidatorWeight {
  return { creator: "", hostZone: "", valAddr: "", weight: 0 };
}

export const MsgChangeValidatorWeight = {
  encode(message: MsgChangeValidatorWeight, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.hostZone !== "") {
      writer.uint32(18).string(message.hostZone);
    }
    if (message.valAddr !== "") {
      writer.uint32(26).string(message.valAddr);
    }
    if (message.weight !== 0) {
      writer.uint32(32).uint64(message.weight);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgChangeValidatorWeight {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgChangeValidatorWeight();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.hostZone = reader.string();
          break;
        case 3:
          message.valAddr = reader.string();
          break;
        case 4:
          message.weight = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgChangeValidatorWeight {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      hostZone: isSet(object.hostZone) ? String(object.hostZone) : "",
      valAddr: isSet(object.valAddr) ? String(object.valAddr) : "",
      weight: isSet(object.weight) ? Number(object.weight) : 0,
    };
  },

  toJSON(message: MsgChangeValidatorWeight): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.hostZone !== undefined && (obj.hostZone = message.hostZone);
    message.valAddr !== undefined && (obj.valAddr = message.valAddr);
    message.weight !== undefined && (obj.weight = Math.round(message.weight));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgChangeValidatorWeight>, I>>(object: I): MsgChangeValidatorWeight {
    const message = createBaseMsgChangeValidatorWeight();
    message.creator = object.creator ?? "";
    message.hostZone = object.hostZone ?? "";
    message.valAddr = object.valAddr ?? "";
    message.weight = object.weight ?? 0;
    return message;
  },
};

function createBaseMsgChangeValidatorWeightResponse(): MsgChangeValidatorWeightResponse {
  return {};
}

export const MsgChangeValidatorWeightResponse = {
  encode(_: MsgChangeValidatorWeightResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgChangeValidatorWeightResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgChangeValidatorWeightResponse();
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

  fromJSON(_: any): MsgChangeValidatorWeightResponse {
    return {};
  },

  toJSON(_: MsgChangeValidatorWeightResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgChangeValidatorWeightResponse>, I>>(
    _: I,
  ): MsgChangeValidatorWeightResponse {
    const message = createBaseMsgChangeValidatorWeightResponse();
    return message;
  },
};

function createBaseMsgDeleteValidator(): MsgDeleteValidator {
  return { creator: "", hostZone: "", valAddr: "" };
}

export const MsgDeleteValidator = {
  encode(message: MsgDeleteValidator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.hostZone !== "") {
      writer.uint32(18).string(message.hostZone);
    }
    if (message.valAddr !== "") {
      writer.uint32(26).string(message.valAddr);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteValidator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteValidator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.hostZone = reader.string();
          break;
        case 3:
          message.valAddr = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteValidator {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      hostZone: isSet(object.hostZone) ? String(object.hostZone) : "",
      valAddr: isSet(object.valAddr) ? String(object.valAddr) : "",
    };
  },

  toJSON(message: MsgDeleteValidator): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.hostZone !== undefined && (obj.hostZone = message.hostZone);
    message.valAddr !== undefined && (obj.valAddr = message.valAddr);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteValidator>, I>>(object: I): MsgDeleteValidator {
    const message = createBaseMsgDeleteValidator();
    message.creator = object.creator ?? "";
    message.hostZone = object.hostZone ?? "";
    message.valAddr = object.valAddr ?? "";
    return message;
  },
};

function createBaseMsgDeleteValidatorResponse(): MsgDeleteValidatorResponse {
  return {};
}

export const MsgDeleteValidatorResponse = {
  encode(_: MsgDeleteValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteValidatorResponse();
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

  fromJSON(_: any): MsgDeleteValidatorResponse {
    return {};
  },

  toJSON(_: MsgDeleteValidatorResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteValidatorResponse>, I>>(_: I): MsgDeleteValidatorResponse {
    const message = createBaseMsgDeleteValidatorResponse();
    return message;
  },
};

function createBaseMsgRestoreInterchainAccount(): MsgRestoreInterchainAccount {
  return { creator: "", chainId: "", accountType: 0 };
}

export const MsgRestoreInterchainAccount = {
  encode(message: MsgRestoreInterchainAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.chainId !== "") {
      writer.uint32(18).string(message.chainId);
    }
    if (message.accountType !== 0) {
      writer.uint32(24).int32(message.accountType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRestoreInterchainAccount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRestoreInterchainAccount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.chainId = reader.string();
          break;
        case 3:
          message.accountType = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRestoreInterchainAccount {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      chainId: isSet(object.chainId) ? String(object.chainId) : "",
      accountType: isSet(object.accountType) ? iCAAccountTypeFromJSON(object.accountType) : 0,
    };
  },

  toJSON(message: MsgRestoreInterchainAccount): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.accountType !== undefined && (obj.accountType = iCAAccountTypeToJSON(message.accountType));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRestoreInterchainAccount>, I>>(object: I): MsgRestoreInterchainAccount {
    const message = createBaseMsgRestoreInterchainAccount();
    message.creator = object.creator ?? "";
    message.chainId = object.chainId ?? "";
    message.accountType = object.accountType ?? 0;
    return message;
  },
};

function createBaseMsgRestoreInterchainAccountResponse(): MsgRestoreInterchainAccountResponse {
  return {};
}

export const MsgRestoreInterchainAccountResponse = {
  encode(_: MsgRestoreInterchainAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRestoreInterchainAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRestoreInterchainAccountResponse();
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

  fromJSON(_: any): MsgRestoreInterchainAccountResponse {
    return {};
  },

  toJSON(_: MsgRestoreInterchainAccountResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRestoreInterchainAccountResponse>, I>>(
    _: I,
  ): MsgRestoreInterchainAccountResponse {
    const message = createBaseMsgRestoreInterchainAccountResponse();
    return message;
  },
};

function createBaseMsgUpdateValidatorSharesExchRate(): MsgUpdateValidatorSharesExchRate {
  return { creator: "", chainId: "", valoper: "" };
}

export const MsgUpdateValidatorSharesExchRate = {
  encode(message: MsgUpdateValidatorSharesExchRate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.chainId !== "") {
      writer.uint32(18).string(message.chainId);
    }
    if (message.valoper !== "") {
      writer.uint32(26).string(message.valoper);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateValidatorSharesExchRate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateValidatorSharesExchRate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.chainId = reader.string();
          break;
        case 3:
          message.valoper = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateValidatorSharesExchRate {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      chainId: isSet(object.chainId) ? String(object.chainId) : "",
      valoper: isSet(object.valoper) ? String(object.valoper) : "",
    };
  },

  toJSON(message: MsgUpdateValidatorSharesExchRate): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.valoper !== undefined && (obj.valoper = message.valoper);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateValidatorSharesExchRate>, I>>(
    object: I,
  ): MsgUpdateValidatorSharesExchRate {
    const message = createBaseMsgUpdateValidatorSharesExchRate();
    message.creator = object.creator ?? "";
    message.chainId = object.chainId ?? "";
    message.valoper = object.valoper ?? "";
    return message;
  },
};

function createBaseMsgUpdateValidatorSharesExchRateResponse(): MsgUpdateValidatorSharesExchRateResponse {
  return {};
}

export const MsgUpdateValidatorSharesExchRateResponse = {
  encode(_: MsgUpdateValidatorSharesExchRateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateValidatorSharesExchRateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateValidatorSharesExchRateResponse();
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

  fromJSON(_: any): MsgUpdateValidatorSharesExchRateResponse {
    return {};
  },

  toJSON(_: MsgUpdateValidatorSharesExchRateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateValidatorSharesExchRateResponse>, I>>(
    _: I,
  ): MsgUpdateValidatorSharesExchRateResponse {
    const message = createBaseMsgUpdateValidatorSharesExchRateResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  LiquidStake(request: MsgLiquidStake): Promise<MsgLiquidStakeResponse>;
  RedeemStake(request: MsgRedeemStake): Promise<MsgRedeemStakeResponse>;
  /**
   * TODO(TEST-53): Remove this pre-launch (no need for clients to create /
   * interact with ICAs)
   */
  RegisterHostZone(request: MsgRegisterHostZone): Promise<MsgRegisterHostZoneResponse>;
  ClaimUndelegatedTokens(request: MsgClaimUndelegatedTokens): Promise<MsgClaimUndelegatedTokensResponse>;
  RebalanceValidators(request: MsgRebalanceValidators): Promise<MsgRebalanceValidatorsResponse>;
  AddValidator(request: MsgAddValidator): Promise<MsgAddValidatorResponse>;
  ChangeValidatorWeight(request: MsgChangeValidatorWeight): Promise<MsgChangeValidatorWeightResponse>;
  DeleteValidator(request: MsgDeleteValidator): Promise<MsgDeleteValidatorResponse>;
  RestoreInterchainAccount(request: MsgRestoreInterchainAccount): Promise<MsgRestoreInterchainAccountResponse>;
  UpdateValidatorSharesExchRate(
    request: MsgUpdateValidatorSharesExchRate,
  ): Promise<MsgUpdateValidatorSharesExchRateResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ClearBalance(request: MsgClearBalance): Promise<MsgClearBalanceResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.LiquidStake = this.LiquidStake.bind(this);
    this.RedeemStake = this.RedeemStake.bind(this);
    this.RegisterHostZone = this.RegisterHostZone.bind(this);
    this.ClaimUndelegatedTokens = this.ClaimUndelegatedTokens.bind(this);
    this.RebalanceValidators = this.RebalanceValidators.bind(this);
    this.AddValidator = this.AddValidator.bind(this);
    this.ChangeValidatorWeight = this.ChangeValidatorWeight.bind(this);
    this.DeleteValidator = this.DeleteValidator.bind(this);
    this.RestoreInterchainAccount = this.RestoreInterchainAccount.bind(this);
    this.UpdateValidatorSharesExchRate = this.UpdateValidatorSharesExchRate.bind(this);
    this.ClearBalance = this.ClearBalance.bind(this);
  }
  LiquidStake(request: MsgLiquidStake): Promise<MsgLiquidStakeResponse> {
    const data = MsgLiquidStake.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Msg", "LiquidStake", data);
    return promise.then((data) => MsgLiquidStakeResponse.decode(new _m0.Reader(data)));
  }

  RedeemStake(request: MsgRedeemStake): Promise<MsgRedeemStakeResponse> {
    const data = MsgRedeemStake.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Msg", "RedeemStake", data);
    return promise.then((data) => MsgRedeemStakeResponse.decode(new _m0.Reader(data)));
  }

  RegisterHostZone(request: MsgRegisterHostZone): Promise<MsgRegisterHostZoneResponse> {
    const data = MsgRegisterHostZone.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Msg", "RegisterHostZone", data);
    return promise.then((data) => MsgRegisterHostZoneResponse.decode(new _m0.Reader(data)));
  }

  ClaimUndelegatedTokens(request: MsgClaimUndelegatedTokens): Promise<MsgClaimUndelegatedTokensResponse> {
    const data = MsgClaimUndelegatedTokens.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Msg", "ClaimUndelegatedTokens", data);
    return promise.then((data) => MsgClaimUndelegatedTokensResponse.decode(new _m0.Reader(data)));
  }

  RebalanceValidators(request: MsgRebalanceValidators): Promise<MsgRebalanceValidatorsResponse> {
    const data = MsgRebalanceValidators.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Msg", "RebalanceValidators", data);
    return promise.then((data) => MsgRebalanceValidatorsResponse.decode(new _m0.Reader(data)));
  }

  AddValidator(request: MsgAddValidator): Promise<MsgAddValidatorResponse> {
    const data = MsgAddValidator.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Msg", "AddValidator", data);
    return promise.then((data) => MsgAddValidatorResponse.decode(new _m0.Reader(data)));
  }

  ChangeValidatorWeight(request: MsgChangeValidatorWeight): Promise<MsgChangeValidatorWeightResponse> {
    const data = MsgChangeValidatorWeight.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Msg", "ChangeValidatorWeight", data);
    return promise.then((data) => MsgChangeValidatorWeightResponse.decode(new _m0.Reader(data)));
  }

  DeleteValidator(request: MsgDeleteValidator): Promise<MsgDeleteValidatorResponse> {
    const data = MsgDeleteValidator.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Msg", "DeleteValidator", data);
    return promise.then((data) => MsgDeleteValidatorResponse.decode(new _m0.Reader(data)));
  }

  RestoreInterchainAccount(request: MsgRestoreInterchainAccount): Promise<MsgRestoreInterchainAccountResponse> {
    const data = MsgRestoreInterchainAccount.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Msg", "RestoreInterchainAccount", data);
    return promise.then((data) => MsgRestoreInterchainAccountResponse.decode(new _m0.Reader(data)));
  }

  UpdateValidatorSharesExchRate(
    request: MsgUpdateValidatorSharesExchRate,
  ): Promise<MsgUpdateValidatorSharesExchRateResponse> {
    const data = MsgUpdateValidatorSharesExchRate.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Msg", "UpdateValidatorSharesExchRate", data);
    return promise.then((data) => MsgUpdateValidatorSharesExchRateResponse.decode(new _m0.Reader(data)));
  }

  ClearBalance(request: MsgClearBalance): Promise<MsgClearBalanceResponse> {
    const data = MsgClearBalance.encode(request).finish();
    const promise = this.rpc.request("stayking.stakeibc.Msg", "ClearBalance", data);
    return promise.then((data) => MsgClearBalanceResponse.decode(new _m0.Reader(data)));
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
