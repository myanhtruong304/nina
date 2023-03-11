package request

type AddProject struct {
	ProjectName     string `json:"project_name" binding:"required"`
	Symbol          string `json:"symbol" binding:"required"`
	ContractAddress string `json:"contract_address" binding:"required"`
	Owner           string `json:"owner" binding:"required"`
}
