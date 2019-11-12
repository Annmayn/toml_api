package methodconfigs

type PatchRequestConfig struct{
	Auth bool   `json:"auth"`
	Schema string	`json:"schema"`
	Data string 	`json:"data"`
	Validator []string	`json:"validator"`
	Result string	`json:"result"`
}
