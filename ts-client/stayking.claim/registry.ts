import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSetAirdropAllocations } from "./types/stayking/claim/tx";
import { MsgClaimFreeAmount } from "./types/stayking/claim/tx";
import { MsgCreateAirdrop } from "./types/stayking/claim/tx";
import { MsgDeleteAirdrop } from "./types/stayking/claim/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stayking.claim.MsgSetAirdropAllocations", MsgSetAirdropAllocations],
    ["/stayking.claim.MsgClaimFreeAmount", MsgClaimFreeAmount],
    ["/stayking.claim.MsgCreateAirdrop", MsgCreateAirdrop],
    ["/stayking.claim.MsgDeleteAirdrop", MsgDeleteAirdrop],
    
];

export { msgTypes }