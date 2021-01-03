package handle

import (
	"bgp/logger"
	"bgp/operation"
	"encoding/json"
	"golang.org/x/net/context"
	"net/http"
)

func CreateStatement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	var c createStatement
	var err error
	if r.Body == nil || r.Method != "POST" {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagOne)
		_, _ = w.Write(msg)
		logger.Logger.Error("Illegal request")
	}
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagTwo)
		_, _ = w.Write(msg)
		logger.Logger.Errorf("happen a err that is %v", err)
	}

	reply, err := operation.AddStatements(context.Background(), c.StatementsName, c.PrefixSetName, c.NeighborSetName, c.CommunitySetName, c.CommunityAction, c.action, c.Community)
	if err != nil {
		w.WriteHeader(404)
		msg, _ := Json("404", err.Error())
		_, _ = w.Write(msg)
		logger.Logger.Errorf("happen a err that is %v", err)
	}

	w.WriteHeader(200)
	msg, _ := Json("200", reply)
	_, _ = w.Write(msg)
}

func DeleteStatement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	var c Statements
	var err error
	if r.Body == nil || r.Method != "POST" {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagOne)
		_, _ = w.Write(msg)
		logger.Logger.Error("Illegal request")
	}
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagTwo)
		_, _ = w.Write(msg)
		logger.Logger.Errorf("happen a err that is %v", err)
	}

	reply, err := operation.DeleteStatements(context.Background(), c.StatementsName)
	if err != nil {
		w.WriteHeader(404)
		msg, _ := Json("404", err.Error())
		_, _ = w.Write(msg)
		logger.Logger.Errorf("happen a err that is %v", err)
	}

	w.WriteHeader(200)
	msg, _ := Json("200", reply)
	_, _ = w.Write(msg)
}

func ListStatement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	var c Statements
	var err error
	if r.Body == nil || r.Method != "POST" {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagOne)
		_, _ = w.Write(msg)
		logger.Logger.Error("Illegal request")
	}
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagTwo)
		_, _ = w.Write(msg)
		logger.Logger.Errorf("happen a err that is %v", err)
	}

	reply, err := operation.ListStatements(context.Background(), c.StatementsName)
	if err != nil {
		w.WriteHeader(404)
		msg, _ := Json("404", err.Error())
		_, _ = w.Write(msg)
		logger.Logger.Errorf("happen a err that is %v", err)
	}
	w.WriteHeader(200)
	re, err := json.Marshal(reply)
	if err != nil {
		logger.Logger.Errorf("happen a err that is %v", err)
	}
	_, _ = w.Write(re)
}