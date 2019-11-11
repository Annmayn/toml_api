package expand

import (
	"strings"
	"toml_api/getresource"
)

//Endpoints : Returns expanded url
func Endpoints(config interface{}, url string) string {
	if url[0] == '$' {
		//url = "$api.root.endpoint/:id"
		//split -> [$api.root.endpoint id]
		//split[0] = "$api.root.endpoint"
		//split[1:] = "api.root.endpoint"
		ref := strings.Split(url, "/")[0][1:]
		dir := strings.Split(ref, ".")
		dirInter := make([]interface{}, len(dir))
		for i, val := range dir {
			dirInter[i] = val
		}
		res := getresource.GetResource(config, dirInter...).(string)
		url = res + url[len(ref)+1:]

		if url[0] == '$' {
			url = Endpoints(config, url)
		}

		return url
	}
	return url
}
