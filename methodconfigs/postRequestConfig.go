package methodconfigs

type PostRequestConfig struct{
	Auth bool `json:"auth"`
	Data string `json:"data"`
	Validator []string `json:"validator"`
	Result string `json:"result"`
}
