package types

func FormatICAAccountOwner(chainId string, accountType ICAType) (result string) {
	return ModuleName + "." + chainId + "." + accountType.String()
}
