package types

import "encoding/json"

type GraphqlDefaultHeaders struct {
	Map map[string]string
}

func (g *GraphqlDefaultHeaders) New() {
	g.Map = make(map[string]string)
	g.Map["User-Agent"] = "RivianApp/707 CFNetwork/1237 Darwin/20.4.0"
	g.Map["Accept"] = "application/json"
	g.Map["Content-Type"] = "application/json"
	g.Map["Apollographql-Client-Name"] = "com.rivian.ios.consumer-apollo-ios"
}

func (g *GraphqlDefaultHeaders) Add(key string, value string) {
	g.Map[key] = value
}

type GraphqlBody struct {
	Json []byte
}

func (g *GraphqlBody) New(operation string, query string, variables map[string]interface{}) {
	b := make(map[string]interface{})
	b["operationName"] = operation
	b["query"] = query
	b["variables"] = variables
	m, _ := json.Marshal(b)
	g.Json = m
}
