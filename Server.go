package main

import (
	"bgp/handle"
	"net/http"
)

func main() {
	handle.Init()
	//handle.CreatePrefixSet(context.Background(), "PREFIX", "text", "10.1.1.0/24", "24", "32")

	http.HandleFunc("/CreatePolicy", handle.CreatePolicy)

	http.HandleFunc("/CreateStatement", handle.CreateStatement)

	http.HandleFunc("/CreatePrefixSet", handle.CreatePrefixSet)

	http.HandleFunc("/StartBGP", handle.StartBGP)

	http.HandleFunc("/CreateGlobalPolicy", handle.CreateGlobalPolicy)

	http.HandleFunc("/AddStatementToPolicy", handle.AddStatementToPolicy)

	_ = http.ListenAndServe(":8000", nil)
}
