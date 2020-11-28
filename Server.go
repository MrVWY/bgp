package main

import (
	"bgp/handle"
	"net/http"
)

func main() {
	handle.Init("address")
	//handle.CreatePrefixSet(context.Background(), "PREFIX", "text", "10.1.1.0/24", "24", "32")

	http.HandleFunc("/CreatePolicy", handle.CreatePolicy)
	http.HandleFunc("/DeletePolicy", handle.DeletePolicy)
	http.HandleFunc("/AddStatementToPolicy", handle.AddStatementToPolicy)

	http.HandleFunc("/CreateStatement", handle.CreateStatement)
	http.HandleFunc("/DeleteStatement", handle.DeleteStatement)

	http.HandleFunc("/CreatePrefixSet", handle.CreatePrefixSet)
	http.HandleFunc("DeletePrefixSet", handle.DeletePrefixSet)

	http.HandleFunc("/StartBGP", handle.StartBGP)
	http.HandleFunc("/CreateGlobalPolicy", handle.CreateGlobalPolicy)

	_ = http.ListenAndServe(":8000", nil)
}
