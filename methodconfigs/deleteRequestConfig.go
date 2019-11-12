package methodconfigs

type DeleteRequestConfig struct{
	Auth bool `json:"auth"`
	Data string	`json:"data"`
	Result string	`json:"result"`
}
