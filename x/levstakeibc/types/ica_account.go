package types

func FormatICAAccountOwner(chainId string, accountType ICAType) (result string) {
	return chainId + "." + accountType.String()
}
