package methodconfigs

type PatchRequestConfig struct{
	auth bool
	schema string
	data string 
	validator []string
	result string
}
