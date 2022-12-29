/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stayking.stakeibc";

/**
 * Params defines the parameters for the module.
 * next id: 18
 */
export interface Params {
  /** define epoch lengths, in stride_epochs */
  rewardsInterval: number;
  delegateInterval: number;
  depositInterval: number;
  redemptionRateInterval: number;
  strideCommission: number;
  /**
   * zone_com_address stores which addresses to
   * send the Stride commission too, as well as what portion
   * of the fee each address is entitled to
   * TODO implement this
   */
  zoneComAddress: { [key: string]: string };
  reinvestInterval: number;
  validatorRebalancingThreshold: number;
  icaTimeoutNanos: number;
  bufferSize: number;
  ibcTimeoutBlocks: number;
  feeTransferTimeoutNanos: number;
  maxStakeIcaCallsPerEpoch: number;
  safetyMinRedemptionRateThreshold: number;
  safetyMaxRedemptionRateThreshold: number;
  ibcTransferTimeoutNanos: number;
  safetyNumValidators: number;
}

export interface Params_ZoneComAddressEntry {
  key: string;
  value: string;
}

function createBaseParams(): Params {
  return {
    rewardsInterval: 0,
    delegateInterval: 0,
    depositInterval: 0,
    redemptionRateInterval: 0,
    strideCommission: 0,
    zoneComAddress: {},
    reinvestInterval: 0,
    validatorRebalancingThreshold: 0,
    icaTimeoutNanos: 0,
    bufferSize: 0,
    ibcTimeoutBlocks: 0,
    feeTransferTimeoutNanos: 0,
    maxStakeIcaCallsPerEpoch: 0,
    safetyMinRedemptionRateThreshold: 0,
    safetyMaxRedemptionRateThreshold: 0,
    ibcTransferTimeoutNanos: 0,
    safetyNumValidators: 0,
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.rewardsInterval !== 0) {
      writer.uint32(8).uint64(message.rewardsInterval);
    }
    if (message.delegateInterval !== 0) {
      writer.uint32(48).uint64(message.delegateInterval);
    }
    if (message.depositInterval !== 0) {
      writer.uint32(16).uint64(message.depositInterval);
    }
    if (message.redemptionRateInterval !== 0) {
      writer.uint32(24).uint64(message.redemptionRateInterval);
    }
    if (message.strideCommission !== 0) {
      writer.uint32(32).uint64(message.strideCommission);
    }
    Object.entries(message.zoneComAddress).forEach(([key, value]) => {
      Params_ZoneComAddressEntry.encode({ key: key as any, value }, writer.uint32(42).fork()).ldelim();
    });
    if (message.reinvestInterval !== 0) {
      writer.uint32(56).uint64(message.reinvestInterval);
    }
    if (message.validatorRebalancingThreshold !== 0) {
      writer.uint32(64).uint64(message.validatorRebalancingThreshold);
    }
    if (message.icaTimeoutNanos !== 0) {
      writer.uint32(72).uint64(message.icaTimeoutNanos);
    }
    if (message.bufferSize !== 0) {
      writer.uint32(80).uint64(message.bufferSize);
    }
    if (message.ibcTimeoutBlocks !== 0) {
      writer.uint32(88).uint64(message.ibcTimeoutBlocks);
    }
    if (message.feeTransferTimeoutNanos !== 0) {
      writer.uint32(96).uint64(message.feeTransferTimeoutNanos);
    }
    if (message.maxStakeIcaCallsPerEpoch !== 0) {
      writer.uint32(104).uint64(message.maxStakeIcaCallsPerEpoch);
    }
    if (message.safetyMinRedemptionRateThreshold !== 0) {
      writer.uint32(112).uint64(message.safetyMinRedemptionRateThreshold);
    }
    if (message.safetyMaxRedemptionRateThreshold !== 0) {
      writer.uint32(120).uint64(message.safetyMaxRedemptionRateThreshold);
    }
    if (message.ibcTransferTimeoutNanos !== 0) {
      writer.uint32(128).uint64(message.ibcTransferTimeoutNanos);
    }
    if (message.safetyNumValidators !== 0) {
      writer.uint32(136).uint64(message.safetyNumValidators);
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
          message.rewardsInterval = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.delegateInterval = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.depositInterval = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.redemptionRateInterval = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.strideCommission = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          const entry5 = Params_ZoneComAddressEntry.decode(reader, reader.uint32());
          if (entry5.value !== undefined) {
            message.zoneComAddress[entry5.key] = entry5.value;
          }
          break;
        case 7:
          message.reinvestInterval = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.validatorRebalancingThreshold = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          message.icaTimeoutNanos = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.bufferSize = longToNumber(reader.uint64() as Long);
          break;
        case 11:
          message.ibcTimeoutBlocks = longToNumber(reader.uint64() as Long);
          break;
        case 12:
          message.feeTransferTimeoutNanos = longToNumber(reader.uint64() as Long);
          break;
        case 13:
          message.maxStakeIcaCallsPerEpoch = longToNumber(reader.uint64() as Long);
          break;
        case 14:
          message.safetyMinRedemptionRateThreshold = longToNumber(reader.uint64() as Long);
          break;
        case 15:
          message.safetyMaxRedemptionRateThreshold = longToNumber(reader.uint64() as Long);
          break;
        case 16:
          message.ibcTransferTimeoutNanos = longToNumber(reader.uint64() as Long);
          break;
        case 17:
          message.safetyNumValidators = longToNumber(reader.uint64() as Long);
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
      rewardsInterval: isSet(object.rewardsInterval) ? Number(object.rewardsInterval) : 0,
      delegateInterval: isSet(object.delegateInterval) ? Number(object.delegateInterval) : 0,
      depositInterval: isSet(object.depositInterval) ? Number(object.depositInterval) : 0,
      redemptionRateInterval: isSet(object.redemptionRateInterval) ? Number(object.redemptionRateInterval) : 0,
      strideCommission: isSet(object.strideCommission) ? Number(object.strideCommission) : 0,
      zoneComAddress: isObject(object.zoneComAddress)
        ? Object.entries(object.zoneComAddress).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
      reinvestInterval: isSet(object.reinvestInterval) ? Number(object.reinvestInterval) : 0,
      validatorRebalancingThreshold: isSet(object.validatorRebalancingThreshold)
        ? Number(object.validatorRebalancingThreshold)
        : 0,
      icaTimeoutNanos: isSet(object.icaTimeoutNanos) ? Number(object.icaTimeoutNanos) : 0,
      bufferSize: isSet(object.bufferSize) ? Number(object.bufferSize) : 0,
      ibcTimeoutBlocks: isSet(object.ibcTimeoutBlocks) ? Number(object.ibcTimeoutBlocks) : 0,
      feeTransferTimeoutNanos: isSet(object.feeTransferTimeoutNanos) ? Number(object.feeTransferTimeoutNanos) : 0,
      maxStakeIcaCallsPerEpoch: isSet(object.maxStakeIcaCallsPerEpoch) ? Number(object.maxStakeIcaCallsPerEpoch) : 0,
      safetyMinRedemptionRateThreshold: isSet(object.safetyMinRedemptionRateThreshold)
        ? Number(object.safetyMinRedemptionRateThreshold)
        : 0,
      safetyMaxRedemptionRateThreshold: isSet(object.safetyMaxRedemptionRateThreshold)
        ? Number(object.safetyMaxRedemptionRateThreshold)
        : 0,
      ibcTransferTimeoutNanos: isSet(object.ibcTransferTimeoutNanos) ? Number(object.ibcTransferTimeoutNanos) : 0,
      safetyNumValidators: isSet(object.safetyNumValidators) ? Number(object.safetyNumValidators) : 0,
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.rewardsInterval !== undefined && (obj.rewardsInterval = Math.round(message.rewardsInterval));
    message.delegateInterval !== undefined && (obj.delegateInterval = Math.round(message.delegateInterval));
    message.depositInterval !== undefined && (obj.depositInterval = Math.round(message.depositInterval));
    message.redemptionRateInterval !== undefined
      && (obj.redemptionRateInterval = Math.round(message.redemptionRateInterval));
    message.strideCommission !== undefined && (obj.strideCommission = Math.round(message.strideCommission));
    obj.zoneComAddress = {};
    if (message.zoneComAddress) {
      Object.entries(message.zoneComAddress).forEach(([k, v]) => {
        obj.zoneComAddress[k] = v;
      });
    }
    message.reinvestInterval !== undefined && (obj.reinvestInterval = Math.round(message.reinvestInterval));
    message.validatorRebalancingThreshold !== undefined
      && (obj.validatorRebalancingThreshold = Math.round(message.validatorRebalancingThreshold));
    message.icaTimeoutNanos !== undefined && (obj.icaTimeoutNanos = Math.round(message.icaTimeoutNanos));
    message.bufferSize !== undefined && (obj.bufferSize = Math.round(message.bufferSize));
    message.ibcTimeoutBlocks !== undefined && (obj.ibcTimeoutBlocks = Math.round(message.ibcTimeoutBlocks));
    message.feeTransferTimeoutNanos !== undefined
      && (obj.feeTransferTimeoutNanos = Math.round(message.feeTransferTimeoutNanos));
    message.maxStakeIcaCallsPerEpoch !== undefined
      && (obj.maxStakeIcaCallsPerEpoch = Math.round(message.maxStakeIcaCallsPerEpoch));
    message.safetyMinRedemptionRateThreshold !== undefined
      && (obj.safetyMinRedemptionRateThreshold = Math.round(message.safetyMinRedemptionRateThreshold));
    message.safetyMaxRedemptionRateThreshold !== undefined
      && (obj.safetyMaxRedemptionRateThreshold = Math.round(message.safetyMaxRedemptionRateThreshold));
    message.ibcTransferTimeoutNanos !== undefined
      && (obj.ibcTransferTimeoutNanos = Math.round(message.ibcTransferTimeoutNanos));
    message.safetyNumValidators !== undefined && (obj.safetyNumValidators = Math.round(message.safetyNumValidators));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.rewardsInterval = object.rewardsInterval ?? 0;
    message.delegateInterval = object.delegateInterval ?? 0;
    message.depositInterval = object.depositInterval ?? 0;
    message.redemptionRateInterval = object.redemptionRateInterval ?? 0;
    message.strideCommission = object.strideCommission ?? 0;
    message.zoneComAddress = Object.entries(object.zoneComAddress ?? {}).reduce<{ [key: string]: string }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = String(value);
        }
        return acc;
      },
      {},
    );
    message.reinvestInterval = object.reinvestInterval ?? 0;
    message.validatorRebalancingThreshold = object.validatorRebalancingThreshold ?? 0;
    message.icaTimeoutNanos = object.icaTimeoutNanos ?? 0;
    message.bufferSize = object.bufferSize ?? 0;
    message.ibcTimeoutBlocks = object.ibcTimeoutBlocks ?? 0;
    message.feeTransferTimeoutNanos = object.feeTransferTimeoutNanos ?? 0;
    message.maxStakeIcaCallsPerEpoch = object.maxStakeIcaCallsPerEpoch ?? 0;
    message.safetyMinRedemptionRateThreshold = object.safetyMinRedemptionRateThreshold ?? 0;
    message.safetyMaxRedemptionRateThreshold = object.safetyMaxRedemptionRateThreshold ?? 0;
    message.ibcTransferTimeoutNanos = object.ibcTransferTimeoutNanos ?? 0;
    message.safetyNumValidators = object.safetyNumValidators ?? 0;
    return message;
  },
};

function createBaseParams_ZoneComAddressEntry(): Params_ZoneComAddressEntry {
  return { key: "", value: "" };
}

export const Params_ZoneComAddressEntry = {
  encode(message: Params_ZoneComAddressEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params_ZoneComAddressEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams_ZoneComAddressEntry();
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

  fromJSON(object: any): Params_ZoneComAddressEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? String(object.value) : "" };
  },

  toJSON(message: Params_ZoneComAddressEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params_ZoneComAddressEntry>, I>>(object: I): Params_ZoneComAddressEntry {
    const message = createBaseParams_ZoneComAddressEntry();
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
