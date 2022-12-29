/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stayking.mint.v1beta1";

/** Minter represents the minting state. */
export interface Minter {
  /** current epoch provisions */
  epochProvisions: string;
}

/** next id: 5 */
export interface DistributionProportions {
  /**
   * staking defines the proportion of the minted minted_denom that is to be
   * allocated as staking rewards.
   */
  staking: string;
  /**
   * community_pool defines the proportion of the minted mint_denom that is
   * to be allocated to the community pool: growth.
   */
  communityPoolGrowth: string;
  /**
   * community_pool defines the proportion of the minted mint_denom that is
   * to be allocated to the community pool: security budget.
   */
  communityPoolSecurityBudget: string;
  /**
   * strategic_reserve defines the proportion of the minted mint_denom that is
   * to be allocated to the pool: strategic reserve.
   */
  strategicReserve: string;
}

/** Params holds parameters for the mint module. */
export interface Params {
  /** type of coin to mint */
  mintDenom: string;
  /** epoch provisions from the first epoch */
  genesisEpochProvisions: string;
  /** mint epoch identifier */
  epochIdentifier: string;
  /** number of epochs take to reduce rewards */
  reductionPeriodInEpochs: number;
  /** reduction multiplier to execute on each period */
  reductionFactor: string;
  /** distribution_proportions defines the proportion of the minted denom */
  distributionProportions:
    | DistributionProportions
    | undefined;
  /** start epoch to distribute minting rewards */
  mintingRewardsDistributionStartEpoch: number;
}

function createBaseMinter(): Minter {
  return { epochProvisions: "" };
}

export const Minter = {
  encode(message: Minter, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.epochProvisions !== "") {
      writer.uint32(10).string(message.epochProvisions);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Minter {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMinter();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.epochProvisions = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Minter {
    return { epochProvisions: isSet(object.epochProvisions) ? String(object.epochProvisions) : "" };
  },

  toJSON(message: Minter): unknown {
    const obj: any = {};
    message.epochProvisions !== undefined && (obj.epochProvisions = message.epochProvisions);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Minter>, I>>(object: I): Minter {
    const message = createBaseMinter();
    message.epochProvisions = object.epochProvisions ?? "";
    return message;
  },
};

function createBaseDistributionProportions(): DistributionProportions {
  return { staking: "", communityPoolGrowth: "", communityPoolSecurityBudget: "", strategicReserve: "" };
}

export const DistributionProportions = {
  encode(message: DistributionProportions, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.staking !== "") {
      writer.uint32(10).string(message.staking);
    }
    if (message.communityPoolGrowth !== "") {
      writer.uint32(18).string(message.communityPoolGrowth);
    }
    if (message.communityPoolSecurityBudget !== "") {
      writer.uint32(26).string(message.communityPoolSecurityBudget);
    }
    if (message.strategicReserve !== "") {
      writer.uint32(34).string(message.strategicReserve);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DistributionProportions {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDistributionProportions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.staking = reader.string();
          break;
        case 2:
          message.communityPoolGrowth = reader.string();
          break;
        case 3:
          message.communityPoolSecurityBudget = reader.string();
          break;
        case 4:
          message.strategicReserve = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DistributionProportions {
    return {
      staking: isSet(object.staking) ? String(object.staking) : "",
      communityPoolGrowth: isSet(object.communityPoolGrowth) ? String(object.communityPoolGrowth) : "",
      communityPoolSecurityBudget: isSet(object.communityPoolSecurityBudget)
        ? String(object.communityPoolSecurityBudget)
        : "",
      strategicReserve: isSet(object.strategicReserve) ? String(object.strategicReserve) : "",
    };
  },

  toJSON(message: DistributionProportions): unknown {
    const obj: any = {};
    message.staking !== undefined && (obj.staking = message.staking);
    message.communityPoolGrowth !== undefined && (obj.communityPoolGrowth = message.communityPoolGrowth);
    message.communityPoolSecurityBudget !== undefined
      && (obj.communityPoolSecurityBudget = message.communityPoolSecurityBudget);
    message.strategicReserve !== undefined && (obj.strategicReserve = message.strategicReserve);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DistributionProportions>, I>>(object: I): DistributionProportions {
    const message = createBaseDistributionProportions();
    message.staking = object.staking ?? "";
    message.communityPoolGrowth = object.communityPoolGrowth ?? "";
    message.communityPoolSecurityBudget = object.communityPoolSecurityBudget ?? "";
    message.strategicReserve = object.strategicReserve ?? "";
    return message;
  },
};

function createBaseParams(): Params {
  return {
    mintDenom: "",
    genesisEpochProvisions: "",
    epochIdentifier: "",
    reductionPeriodInEpochs: 0,
    reductionFactor: "",
    distributionProportions: undefined,
    mintingRewardsDistributionStartEpoch: 0,
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.mintDenom !== "") {
      writer.uint32(10).string(message.mintDenom);
    }
    if (message.genesisEpochProvisions !== "") {
      writer.uint32(18).string(message.genesisEpochProvisions);
    }
    if (message.epochIdentifier !== "") {
      writer.uint32(26).string(message.epochIdentifier);
    }
    if (message.reductionPeriodInEpochs !== 0) {
      writer.uint32(32).int64(message.reductionPeriodInEpochs);
    }
    if (message.reductionFactor !== "") {
      writer.uint32(42).string(message.reductionFactor);
    }
    if (message.distributionProportions !== undefined) {
      DistributionProportions.encode(message.distributionProportions, writer.uint32(50).fork()).ldelim();
    }
    if (message.mintingRewardsDistributionStartEpoch !== 0) {
      writer.uint32(56).int64(message.mintingRewardsDistributionStartEpoch);
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
          message.mintDenom = reader.string();
          break;
        case 2:
          message.genesisEpochProvisions = reader.string();
          break;
        case 3:
          message.epochIdentifier = reader.string();
          break;
        case 4:
          message.reductionPeriodInEpochs = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.reductionFactor = reader.string();
          break;
        case 6:
          message.distributionProportions = DistributionProportions.decode(reader, reader.uint32());
          break;
        case 7:
          message.mintingRewardsDistributionStartEpoch = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    return {
      mintDenom: isSet(object.mintDenom) ? String(object.mintDenom) : "",
      genesisEpochProvisions: isSet(object.genesisEpochProvisions) ? String(object.genesisEpochProvisions) : "",
      epochIdentifier: isSet(object.epochIdentifier) ? String(object.epochIdentifier) : "",
      reductionPeriodInEpochs: isSet(object.reductionPeriodInEpochs) ? Number(object.reductionPeriodInEpochs) : 0,
      reductionFactor: isSet(object.reductionFactor) ? String(object.reductionFactor) : "",
      distributionProportions: isSet(object.distributionProportions)
        ? DistributionProportions.fromJSON(object.distributionProportions)
        : undefined,
      mintingRewardsDistributionStartEpoch: isSet(object.mintingRewardsDistributionStartEpoch)
        ? Number(object.mintingRewardsDistributionStartEpoch)
        : 0,
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.mintDenom !== undefined && (obj.mintDenom = message.mintDenom);
    message.genesisEpochProvisions !== undefined && (obj.genesisEpochProvisions = message.genesisEpochProvisions);
    message.epochIdentifier !== undefined && (obj.epochIdentifier = message.epochIdentifier);
    message.reductionPeriodInEpochs !== undefined
      && (obj.reductionPeriodInEpochs = Math.round(message.reductionPeriodInEpochs));
    message.reductionFactor !== undefined && (obj.reductionFactor = message.reductionFactor);
    message.distributionProportions !== undefined && (obj.distributionProportions = message.distributionProportions
      ? DistributionProportions.toJSON(message.distributionProportions)
      : undefined);
    message.mintingRewardsDistributionStartEpoch !== undefined
      && (obj.mintingRewardsDistributionStartEpoch = Math.round(message.mintingRewardsDistributionStartEpoch));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.mintDenom = object.mintDenom ?? "";
    message.genesisEpochProvisions = object.genesisEpochProvisions ?? "";
    message.epochIdentifier = object.epochIdentifier ?? "";
    message.reductionPeriodInEpochs = object.reductionPeriodInEpochs ?? 0;
    message.reductionFactor = object.reductionFactor ?? "";
    message.distributionProportions =
      (object.distributionProportions !== undefined && object.distributionProportions !== null)
        ? DistributionProportions.fromPartial(object.distributionProportions)
        : undefined;
    message.mintingRewardsDistributionStartEpoch = object.mintingRewardsDistributionStartEpoch ?? 0;
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
