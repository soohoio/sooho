import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRegisterHostZone } from "./types/stayking/stakeibc/tx";
import { MsgDeleteValidator } from "./types/stayking/stakeibc/tx";
import { MsgLiquidStake } from "./types/stayking/stakeibc/tx";
import { MsgChangeValidatorWeight } from "./types/stayking/stakeibc/tx";
import { MsgRebalanceValidators } from "./types/stayking/stakeibc/tx";
import { MsgAddValidator } from "./types/stayking/stakeibc/tx";
import { MsgRedeemStake } from "./types/stayking/stakeibc/tx";
import { MsgRestoreInterchainAccount } from "./types/stayking/stakeibc/tx";
import { MsgClearBalance } from "./types/stayking/stakeibc/tx";
import { MsgClaimUndelegatedTokens } from "./types/stayking/stakeibc/tx";
import { MsgUpdateValidatorSharesExchRate } from "./types/stayking/stakeibc/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stayking.stakeibc.MsgRegisterHostZone", MsgRegisterHostZone],
    ["/stayking.stakeibc.MsgDeleteValidator", MsgDeleteValidator],
    ["/stayking.stakeibc.MsgLiquidStake", MsgLiquidStake],
    ["/stayking.stakeibc.MsgChangeValidatorWeight", MsgChangeValidatorWeight],
    ["/stayking.stakeibc.MsgRebalanceValidators", MsgRebalanceValidators],
    ["/stayking.stakeibc.MsgAddValidator", MsgAddValidator],
    ["/stayking.stakeibc.MsgRedeemStake", MsgRedeemStake],
    ["/stayking.stakeibc.MsgRestoreInterchainAccount", MsgRestoreInterchainAccount],
    ["/stayking.stakeibc.MsgClearBalance", MsgClearBalance],
    ["/stayking.stakeibc.MsgClaimUndelegatedTokens", MsgClaimUndelegatedTokens],
    ["/stayking.stakeibc.MsgUpdateValidatorSharesExchRate", MsgUpdateValidatorSharesExchRate],
    
];

export { msgTypes }