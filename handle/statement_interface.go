package handle

import (
	"bgp/operation"
	"encoding/json"
	"golang.org/x/net/context"
	"net/http"
)

func CreateStatement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/json")
	var c createStatement
	if r.Body == nil {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagOne)
		_, _ = w.Write(msg)
		//日志
	}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagTwo)
		_, _ = w.Write(msg)
		//日志
	}

	reply, err := operation.AddStatements(context.Background(), c.StatementsName, c.PrefixSetName, c.NeighborSetName, c.CommunitySetName, c.ExtCommunitySetName, c.action)
	if err != nil {
		w.WriteHeader(404)
		msg, _ := Json("404", err.Error())
		_, _ = w.Write(msg)
		//日志
	}

	w.WriteHeader(200)
	msg, _ := Json("200", reply)
	_, _ = w.Write(msg)
}

func DeleteStatement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/json")
	var c deleteStatement
	if r.Body == nil {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagOne)
		_, _ = w.Write(msg)
		//日志
	}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagTwo)
		_, _ = w.Write(msg)
		//日志
	}

	reply, err := operation.DeleteStatements(context.Background(), c.StatementsName)
	if err != nil {
		w.WriteHeader(404)
		msg, _ := Json("404", err.Error())
		_, _ = w.Write(msg)
		//日志
	}

	w.WriteHeader(200)
	msg, _ := Json("200", reply)
	_, _ = w.Write(msg)
}