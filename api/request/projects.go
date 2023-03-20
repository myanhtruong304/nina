package request

type AddProject struct {
	ProjectName     string `json:"project_name" bind:"require"`
	Symbol          string `json:"symbol" bind:"require"`
	Website         string `json:"website" bind:"require"`
	Twitter         string `json:"twitter"`
	Telegram        string `json:"telegram"`
	Facebook        string `json:"facebook"`
	Linkedin        string `json:"linkedin"`
	Medium          string `json:"medium"`
	Whitepaper      string `json:"whitepaper"`
	Email           string `json:"email"`
	ContractAddress string `json:"contract_address"`
	Explorer        string `json:"explorer"`
	Owner           string `json:"owner" bind:"require"`
}
