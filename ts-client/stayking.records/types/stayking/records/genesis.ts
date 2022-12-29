/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stayking.records";

export interface UserRedemptionRecord {
  /** {chain_id}.{epoch}.{sender} */
  id: string;
  sender: string;
  receiver: string;
  amount: string;
  denom: string;
  hostZoneId: string;
  epochNumber: number;
  claimIsPending: boolean;
}

/** Params defines the parameters for the module. */
export interface Params {
}

export interface RecordsPacketData {
  /** this line is used by starport scaffolding # ibc/packet/proto/field */
  noData: NoData | undefined;
}

export interface NoData {
}

export interface DepositRecord {
  id: number;
  amount: string;
  denom: string;
  hostZoneId: string;
  status: DepositRecord_Status;
  depositEpochNumber: number;
  source: DepositRecord_Source;
}

export enum DepositRecord_Status {
  /** TRANSFER_QUEUE - in transfer queue to be sent to the delegation ICA */
  TRANSFER_QUEUE = 0,
  /** TRANSFER_IN_PROGRESS - transfer in progress (IBC packet sent, ack not received) */
  TRANSFER_IN_PROGRESS = 2,
  /** DELEGATION_QUEUE - in staking queue on delegation ICA */
  DELEGATION_QUEUE = 1,
  /** DELEGATION_IN_PROGRESS - staking in progress (ICA packet sent, ack not received) */
  DELEGATION_IN_PROGRESS = 3,
  UNRECOGNIZED = -1,
}

export function depositRecord_StatusFromJSON(object: any): DepositRecord_Status {
  switch (object) {
    case 0:
    case "TRANSFER_QUEUE":
      return DepositRecord_Status.TRANSFER_QUEUE;
    case 2:
    case "TRANSFER_IN_PROGRESS":
      return DepositRecord_Status.TRANSFER_IN_PROGRESS;
    case 1:
    case "DELEGATION_QUEUE":
      return DepositRecord_Status.DELEGATION_QUEUE;
    case 3:
    case "DELEGATION_IN_PROGRESS":
      return DepositRecord_Status.DELEGATION_IN_PROGRESS;
    case -1:
    case "UNRECOGNIZED":
    default:
      return DepositRecord_Status.UNRECOGNIZED;
  }
}

export function depositRecord_StatusToJSON(object: DepositRecord_Status): string {
  switch (object) {
    case DepositRecord_Status.TRANSFER_QUEUE:
      return "TRANSFER_QUEUE";
    case DepositRecord_Status.TRANSFER_IN_PROGRESS:
      return "TRANSFER_IN_PROGRESS";
    case DepositRecord_Status.DELEGATION_QUEUE:
      return "DELEGATION_QUEUE";
    case DepositRecord_Status.DELEGATION_IN_PROGRESS:
      return "DELEGATION_IN_PROGRESS";
    case DepositRecord_Status.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum DepositRecord_Source {
  STRIDE = 0,
  WITHDRAWAL_ICA = 1,
  UNRECOGNIZED = -1,
}

export function depositRecord_SourceFromJSON(object: any): DepositRecord_Source {
  switch (object) {
    case 0:
    case "STRIDE":
      return DepositRecord_Source.STRIDE;
    case 1:
    case "WITHDRAWAL_ICA":
      return DepositRecord_Source.WITHDRAWAL_ICA;
    case -1:
    case "UNRECOGNIZED":
    default:
      return DepositRecord_Source.UNRECOGNIZED;
  }
}

export function depositRecord_SourceToJSON(object: DepositRecord_Source): string {
  switch (object) {
    case DepositRecord_Source.STRIDE:
      return "STRIDE";
    case DepositRecord_Source.WITHDRAWAL_ICA:
      return "WITHDRAWAL_ICA";
    case DepositRecord_Source.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface HostZoneUnbonding {
  stTokenAmount: string;
  nativeTokenAmount: string;
  denom: string;
  hostZoneId: string;
  unbondingTime: number;
  status: HostZoneUnbonding_Status;
  userRedemptionRecords: string[];
}

export enum HostZoneUnbonding_Status {
  /** UNBONDING_QUEUE - tokens bonded on delegate account */
  UNBONDING_QUEUE = 0,
  UNBONDING_IN_PROGRESS = 3,
  /** EXIT_TRANSFER_QUEUE - unbonding completed on delegate account */
  EXIT_TRANSFER_QUEUE = 1,
  EXIT_TRANSFER_IN_PROGRESS = 4,
  /** CLAIMABLE - transfer success */
  CLAIMABLE = 2,
  UNRECOGNIZED = -1,
}

export function hostZoneUnbonding_StatusFromJSON(object: any): HostZoneUnbonding_Status {
  switch (object) {
    case 0:
    case "UNBONDING_QUEUE":
      return HostZoneUnbonding_Status.UNBONDING_QUEUE;
    case 3:
    case "UNBONDING_IN_PROGRESS":
      return HostZoneUnbonding_Status.UNBONDING_IN_PROGRESS;
    case 1:
    case "EXIT_TRANSFER_QUEUE":
      return HostZoneUnbonding_Status.EXIT_TRANSFER_QUEUE;
    case 4:
    case "EXIT_TRANSFER_IN_PROGRESS":
      return HostZoneUnbonding_Status.EXIT_TRANSFER_IN_PROGRESS;
    case 2:
    case "CLAIMABLE":
      return HostZoneUnbonding_Status.CLAIMABLE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return HostZoneUnbonding_Status.UNRECOGNIZED;
  }
}

export function hostZoneUnbonding_StatusToJSON(object: HostZoneUnbonding_Status): string {
  switch (object) {
    case HostZoneUnbonding_Status.UNBONDING_QUEUE:
      return "UNBONDING_QUEUE";
    case HostZoneUnbonding_Status.UNBONDING_IN_PROGRESS:
      return "UNBONDING_IN_PROGRESS";
    case HostZoneUnbonding_Status.EXIT_TRANSFER_QUEUE:
      return "EXIT_TRANSFER_QUEUE";
    case HostZoneUnbonding_Status.EXIT_TRANSFER_IN_PROGRESS:
      return "EXIT_TRANSFER_IN_PROGRESS";
    case HostZoneUnbonding_Status.CLAIMABLE:
      return "CLAIMABLE";
    case HostZoneUnbonding_Status.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface EpochUnbondingRecord {
  epochNumber: number;
  hostZoneUnbondings: HostZoneUnbonding[];
}

/**
 * GenesisState defines the recordπs module's genesis state.
 * next id: 9
 */
export interface GenesisState {
  params: Params | undefined;
  portId: string;
  userRedemptionRecordList: UserRedemptionRecord[];
  userRedemptionRecordCount: number;
  epochUnbondingRecordList: EpochUnbondingRecord[];
  depositRecordList: DepositRecord[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  depositRecordCount: number;
}

function createBaseUserRedemptionRecord(): UserRedemptionRecord {
  return {
    id: "",
    sender: "",
    receiver: "",
    amount: "",
    denom: "",
    hostZoneId: "",
    epochNumber: 0,
    claimIsPending: false,
  };
}

export const UserRedemptionRecord = {
  encode(message: UserRedemptionRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.sender !== "") {
      writer.uint32(18).string(message.sender);
    }
    if (message.receiver !== "") {
      writer.uint32(26).string(message.receiver);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.denom !== "") {
      writer.uint32(42).string(message.denom);
    }
    if (message.hostZoneId !== "") {
      writer.uint32(50).string(message.hostZoneId);
    }
    if (message.epochNumber !== 0) {
      writer.uint32(56).uint64(message.epochNumber);
    }
    if (message.claimIsPending === true) {
      writer.uint32(64).bool(message.claimIsPending);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserRedemptionRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserRedemptionRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.sender = reader.string();
          break;
        case 3:
          message.receiver = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        case 5:
          message.denom = reader.string();
          break;
        case 6:
          message.hostZoneId = reader.string();
          break;
        case 7:
          message.epochNumber = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.claimIsPending = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserRedemptionRecord {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      sender: isSet(object.sender) ? String(object.sender) : "",
      receiver: isSet(object.receiver) ? String(object.receiver) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      denom: isSet(object.denom) ? String(object.denom) : "",
      hostZoneId: isSet(object.hostZoneId) ? String(object.hostZoneId) : "",
      epochNumber: isSet(object.epochNumber) ? Number(object.epochNumber) : 0,
      claimIsPending: isSet(object.claimIsPending) ? Boolean(object.claimIsPending) : false,
    };
  },

  toJSON(message: UserRedemptionRecord): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.sender !== undefined && (obj.sender = message.sender);
    message.receiver !== undefined && (obj.receiver = message.receiver);
    message.amount !== undefined && (obj.amount = message.amount);
    message.denom !== undefined && (obj.denom = message.denom);
    message.hostZoneId !== undefined && (obj.hostZoneId = message.hostZoneId);
    message.epochNumber !== undefined && (obj.epochNumber = Math.round(message.epochNumber));
    message.claimIsPending !== undefined && (obj.claimIsPending = message.claimIsPending);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UserRedemptionRecord>, I>>(object: I): UserRedemptionRecord {
    const message = createBaseUserRedemptionRecord();
    message.id = object.id ?? "";
    message.sender = object.sender ?? "";
    message.receiver = object.receiver ?? "";
    message.amount = object.amount ?? "";
    message.denom = object.denom ?? "";
    message.hostZoneId = object.hostZoneId ?? "";
    message.epochNumber = object.epochNumber ?? 0;
    message.claimIsPending = object.claimIsPending ?? false;
    return message;
  },
};

function createBaseParams(): Params {
  return {};
}

export const Params = {
  encode(_: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();
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

  fromJSON(_: any): Params {
    return {};
  },

  toJSON(_: Params): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(_: I): Params {
    const message = createBaseParams();
    return message;
  },
};

function createBaseRecordsPacketData(): RecordsPacketData {
  return { noData: undefined };
}

export const RecordsPacketData = {
  encode(message: RecordsPacketData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.noData !== undefined) {
      NoData.encode(message.noData, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RecordsPacketData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRecordsPacketData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.noData = NoData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RecordsPacketData {
    return { noData: isSet(object.noData) ? NoData.fromJSON(object.noData) : undefined };
  },

  toJSON(message: RecordsPacketData): unknown {
    const obj: any = {};
    message.noData !== undefined && (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RecordsPacketData>, I>>(object: I): RecordsPacketData {
    const message = createBaseRecordsPacketData();
    message.noData = (object.noData !== undefined && object.noData !== null)
      ? NoData.fromPartial(object.noData)
      : undefined;
    return message;
  },
};

function createBaseNoData(): NoData {
  return {};
}

export const NoData = {
  encode(_: NoData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NoData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNoData();
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

  fromJSON(_: any): NoData {
    return {};
  },

  toJSON(_: NoData): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NoData>, I>>(_: I): NoData {
    const message = createBaseNoData();
    return message;
  },
};

function createBaseDepositRecord(): DepositRecord {
  return { id: 0, amount: "", denom: "", hostZoneId: "", status: 0, depositEpochNumber: 0, source: 0 };
}

export const DepositRecord = {
  encode(message: DepositRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    if (message.denom !== "") {
      writer.uint32(26).string(message.denom);
    }
    if (message.hostZoneId !== "") {
      writer.uint32(34).string(message.hostZoneId);
    }
    if (message.status !== 0) {
      writer.uint32(48).int32(message.status);
    }
    if (message.depositEpochNumber !== 0) {
      writer.uint32(56).uint64(message.depositEpochNumber);
    }
    if (message.source !== 0) {
      writer.uint32(64).int32(message.source);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DepositRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDepositRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.amount = reader.string();
          break;
        case 3:
          message.denom = reader.string();
          break;
        case 4:
          message.hostZoneId = reader.string();
          break;
        case 6:
          message.status = reader.int32() as any;
          break;
        case 7:
          message.depositEpochNumber = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.source = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DepositRecord {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      amount: isSet(object.amount) ? String(object.amount) : "",
      denom: isSet(object.denom) ? String(object.denom) : "",
      hostZoneId: isSet(object.hostZoneId) ? String(object.hostZoneId) : "",
      status: isSet(object.status) ? depositRecord_StatusFromJSON(object.status) : 0,
      depositEpochNumber: isSet(object.depositEpochNumber) ? Number(object.depositEpochNumber) : 0,
      source: isSet(object.source) ? depositRecord_SourceFromJSON(object.source) : 0,
    };
  },

  toJSON(message: DepositRecord): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.amount !== undefined && (obj.amount = message.amount);
    message.denom !== undefined && (obj.denom = message.denom);
    message.hostZoneId !== undefined && (obj.hostZoneId = message.hostZoneId);
    message.status !== undefined && (obj.status = depositRecord_StatusToJSON(message.status));
    message.depositEpochNumber !== undefined && (obj.depositEpochNumber = Math.round(message.depositEpochNumber));
    message.source !== undefined && (obj.source = depositRecord_SourceToJSON(message.source));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DepositRecord>, I>>(object: I): DepositRecord {
    const message = createBaseDepositRecord();
    message.id = object.id ?? 0;
    message.amount = object.amount ?? "";
    message.denom = object.denom ?? "";
    message.hostZoneId = object.hostZoneId ?? "";
    message.status = object.status ?? 0;
    message.depositEpochNumber = object.depositEpochNumber ?? 0;
    message.source = object.source ?? 0;
    return message;
  },
};

function createBaseHostZoneUnbonding(): HostZoneUnbonding {
  return {
    stTokenAmount: "",
    nativeTokenAmount: "",
    denom: "",
    hostZoneId: "",
    unbondingTime: 0,
    status: 0,
    userRedemptionRecords: [],
  };
}

export const HostZoneUnbonding = {
  encode(message: HostZoneUnbonding, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.stTokenAmount !== "") {
      writer.uint32(10).string(message.stTokenAmount);
    }
    if (message.nativeTokenAmount !== "") {
      writer.uint32(18).string(message.nativeTokenAmount);
    }
    if (message.denom !== "") {
      writer.uint32(26).string(message.denom);
    }
    if (message.hostZoneId !== "") {
      writer.uint32(34).string(message.hostZoneId);
    }
    if (message.unbondingTime !== 0) {
      writer.uint32(40).uint64(message.unbondingTime);
    }
    if (message.status !== 0) {
      writer.uint32(48).int32(message.status);
    }
    for (const v of message.userRedemptionRecords) {
      writer.uint32(58).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HostZoneUnbonding {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHostZoneUnbonding();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.stTokenAmount = reader.string();
          break;
        case 2:
          message.nativeTokenAmount = reader.string();
          break;
        case 3:
          message.denom = reader.string();
          break;
        case 4:
          message.hostZoneId = reader.string();
          break;
        case 5:
          message.unbondingTime = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.status = reader.int32() as any;
          break;
        case 7:
          message.userRedemptionRecords.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): HostZoneUnbonding {
    return {
      stTokenAmount: isSet(object.stTokenAmount) ? String(object.stTokenAmount) : "",
      nativeTokenAmount: isSet(object.nativeTokenAmount) ? String(object.nativeTokenAmount) : "",
      denom: isSet(object.denom) ? String(object.denom) : "",
      hostZoneId: isSet(object.hostZoneId) ? String(object.hostZoneId) : "",
      unbondingTime: isSet(object.unbondingTime) ? Number(object.unbondingTime) : 0,
      status: isSet(object.status) ? hostZoneUnbonding_StatusFromJSON(object.status) : 0,
      userRedemptionRecords: Array.isArray(object?.userRedemptionRecords)
        ? object.userRedemptionRecords.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: HostZoneUnbonding): unknown {
    const obj: any = {};
    message.stTokenAmount !== undefined && (obj.stTokenAmount = message.stTokenAmount);
    message.nativeTokenAmount !== undefined && (obj.nativeTokenAmount = message.nativeTokenAmount);
    message.denom !== undefined && (obj.denom = message.denom);
    message.hostZoneId !== undefined && (obj.hostZoneId = message.hostZoneId);
    message.unbondingTime !== undefined && (obj.unbondingTime = Math.round(message.unbondingTime));
    message.status !== undefined && (obj.status = hostZoneUnbonding_StatusToJSON(message.status));
    if (message.userRedemptionRecords) {
      obj.userRedemptionRecords = message.userRedemptionRecords.map((e) => e);
    } else {
      obj.userRedemptionRecords = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<HostZoneUnbonding>, I>>(object: I): HostZoneUnbonding {
    const message = createBaseHostZoneUnbonding();
    message.stTokenAmount = object.stTokenAmount ?? "";
    message.nativeTokenAmount = object.nativeTokenAmount ?? "";
    message.denom = object.denom ?? "";
    message.hostZoneId = object.hostZoneId ?? "";
    message.unbondingTime = object.unbondingTime ?? 0;
    message.status = object.status ?? 0;
    message.userRedemptionRecords = object.userRedemptionRecords?.map((e) => e) || [];
    return message;
  },
};

function createBaseEpochUnbondingRecord(): EpochUnbondingRecord {
  return { epochNumber: 0, hostZoneUnbondings: [] };
}

export const EpochUnbondingRecord = {
  encode(message: EpochUnbondingRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.epochNumber !== 0) {
      writer.uint32(8).uint64(message.epochNumber);
    }
    for (const v of message.hostZoneUnbondings) {
      HostZoneUnbonding.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EpochUnbondingRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEpochUnbondingRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.epochNumber = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.hostZoneUnbondings.push(HostZoneUnbonding.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EpochUnbondingRecord {
    return {
      epochNumber: isSet(object.epochNumber) ? Number(object.epochNumber) : 0,
      hostZoneUnbondings: Array.isArray(object?.hostZoneUnbondings)
        ? object.hostZoneUnbondings.map((e: any) => HostZoneUnbonding.fromJSON(e))
        : [],
    };
  },

  toJSON(message: EpochUnbondingRecord): unknown {
    const obj: any = {};
    message.epochNumber !== undefined && (obj.epochNumber = Math.round(message.epochNumber));
    if (message.hostZoneUnbondings) {
      obj.hostZoneUnbondings = message.hostZoneUnbondings.map((e) => e ? HostZoneUnbonding.toJSON(e) : undefined);
    } else {
      obj.hostZoneUnbondings = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EpochUnbondingRecord>, I>>(object: I): EpochUnbondingRecord {
    const message = createBaseEpochUnbondingRecord();
    message.epochNumber = object.epochNumber ?? 0;
    message.hostZoneUnbondings = object.hostZoneUnbondings?.map((e) => HostZoneUnbonding.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    portId: "",
    userRedemptionRecordList: [],
    userRedemptionRecordCount: 0,
    epochUnbondingRecordList: [],
    depositRecordList: [],
    depositRecordCount: 0,
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
    for (const v of message.userRedemptionRecordList) {
      UserRedemptionRecord.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.userRedemptionRecordCount !== 0) {
      writer.uint32(32).uint64(message.userRedemptionRecordCount);
    }
    for (const v of message.epochUnbondingRecordList) {
      EpochUnbondingRecord.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.depositRecordList) {
      DepositRecord.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    if (message.depositRecordCount !== 0) {
      writer.uint32(64).uint64(message.depositRecordCount);
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
        case 3:
          message.userRedemptionRecordList.push(UserRedemptionRecord.decode(reader, reader.uint32()));
          break;
        case 4:
          message.userRedemptionRecordCount = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.epochUnbondingRecordList.push(EpochUnbondingRecord.decode(reader, reader.uint32()));
          break;
        case 7:
          message.depositRecordList.push(DepositRecord.decode(reader, reader.uint32()));
          break;
        case 8:
          message.depositRecordCount = longToNumber(reader.uint64() as Long);
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
      userRedemptionRecordList: Array.isArray(object?.userRedemptionRecordList)
        ? object.userRedemptionRecordList.map((e: any) => UserRedemptionRecord.fromJSON(e))
        : [],
      userRedemptionRecordCount: isSet(object.userRedemptionRecordCount) ? Number(object.userRedemptionRecordCount) : 0,
      epochUnbondingRecordList: Array.isArray(object?.epochUnbondingRecordList)
        ? object.epochUnbondingRecordList.map((e: any) => EpochUnbondingRecord.fromJSON(e))
        : [],
      depositRecordList: Array.isArray(object?.depositRecordList)
        ? object.depositRecordList.map((e: any) => DepositRecord.fromJSON(e))
        : [],
      depositRecordCount: isSet(object.depositRecordCount) ? Number(object.depositRecordCount) : 0,
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.portId !== undefined && (obj.portId = message.portId);
    if (message.userRedemptionRecordList) {
      obj.userRedemptionRecordList = message.userRedemptionRecordList.map((e) =>
        e ? UserRedemptionRecord.toJSON(e) : undefined
      );
    } else {
      obj.userRedemptionRecordList = [];
    }
    message.userRedemptionRecordCount !== undefined
      && (obj.userRedemptionRecordCount = Math.round(message.userRedemptionRecordCount));
    if (message.epochUnbondingRecordList) {
      obj.epochUnbondingRecordList = message.epochUnbondingRecordList.map((e) =>
        e ? EpochUnbondingRecord.toJSON(e) : undefined
      );
    } else {
      obj.epochUnbondingRecordList = [];
    }
    if (message.depositRecordList) {
      obj.depositRecordList = message.depositRecordList.map((e) => e ? DepositRecord.toJSON(e) : undefined);
    } else {
      obj.depositRecordList = [];
    }
    message.depositRecordCount !== undefined && (obj.depositRecordCount = Math.round(message.depositRecordCount));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.portId = object.portId ?? "";
    message.userRedemptionRecordList = object.userRedemptionRecordList?.map((e) => UserRedemptionRecord.fromPartial(e))
      || [];
    message.userRedemptionRecordCount = object.userRedemptionRecordCount ?? 0;
    message.epochUnbondingRecordList = object.epochUnbondingRecordList?.map((e) => EpochUnbondingRecord.fromPartial(e))
      || [];
    message.depositRecordList = object.depositRecordList?.map((e) => DepositRecord.fromPartial(e)) || [];
    message.depositRecordCount = object.depositRecordCount ?? 0;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
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
