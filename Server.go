package main

import (
	"bgp/handle"
	"net/http"
)

func main() {
	handle.Init("0.0.0.0")
	//handle.CreatePrefixSet(context.Background(), "PREFIX", "text", "10.1.1.0/24", "24", "32")

	http.HandleFunc("/CreatePolicy", handle.CreatePolicy)
	http.HandleFunc("/DeletePolicy", handle.DeletePolicy)
	http.HandleFunc("/AddPolicyToGlobal", handle.AddPolicyToGlobal)
	http.HandleFunc("/ListPolicy", handle.ListPolicy)
	http.HandleFunc("/AddStatementToPolicy", handle.AddStatementToPolicy)

	http.HandleFunc("/CreateStatement", handle.CreateStatement)
	http.HandleFunc("/DeleteStatement", handle.DeleteStatement)
	http.HandleFunc("/ListStatement", handle.ListStatement)

	http.HandleFunc("/CreatePrefixSet", handle.CreatePrefixSet)
	http.HandleFunc("/CreateCommunitySet", handle.CreateCommunitySet)
	http.HandleFunc("/CreateNeighborSet", handle.CreateNeighborSet)
	http.HandleFunc("/DeleteDefined", handle.DeleteDefined)

	http.HandleFunc("/StartBGP", handle.StartBGP)
	http.HandleFunc("/CreateGlobalPolicy", handle.CreateGlobalPolicy)

	_ = http.ListenAndServe(":8000", nil)
}
