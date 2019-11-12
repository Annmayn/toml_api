package methodconfigs

type GetRequestConfig struct{
	Auth bool  `json:"auth"`
	Query string `json:"query"`
	QueryParams string `json:"query_params"`
	Display interface{} `json:"display"`
	Attachments []string `json:"attachments"`
	Result  string   `json:"result"`
}


