import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgClaimFreeAmount } from "./types/stayking/claim/tx";
import { MsgDeleteAirdrop } from "./types/stayking/claim/tx";
import { MsgSetAirdropAllocations } from "./types/stayking/claim/tx";
import { MsgCreateAirdrop } from "./types/stayking/claim/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stayking.claim.MsgClaimFreeAmount", MsgClaimFreeAmount],
    ["/stayking.claim.MsgDeleteAirdrop", MsgDeleteAirdrop],
    ["/stayking.claim.MsgSetAirdropAllocations", MsgSetAirdropAllocations],
    ["/stayking.claim.MsgCreateAirdrop", MsgCreateAirdrop],
    
];

export { msgTypes }