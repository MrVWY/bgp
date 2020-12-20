package handle

import (
	"bgp/logger"
	"bgp/operation"
	"encoding/json"
	"golang.org/x/net/context"
	"net/http"
)

func CreatePolicy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/json")
	var err error
	var c createPolicy
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

	reply, err := operation.AddPolicies(context.Background(), c.PolicyName, c.StatementsName, c.PrefixSetName, c.NeighborSetName, c.CommunitySetName, c.CommunityAction, c.action, c.Community, c.NextHopAddress)
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

func DeletePolicy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/json")
	var c policy
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

	reply, err := operation.DeletePolicies(context.Background(), c.PolicyName)
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

func AddStatementToPolicy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/json")
	var c addStatementToPolicy
	if r.Body == nil || r.Method != "POST" {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagOne)
		_, _ = w.Write(msg)
		logger.Logger.Error("Illegal request")
	}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagTwo)
		_, _ = w.Write(msg)
		logger.Logger.Errorf("happen a err that is %v", err)
	}

	reply, err := operation.AddStatementToPolicy(context.Background(), c.PolicyName, c.StatementName)
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

func AddPolicyToGlobal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/json")
	var c policy
	if r.Body == nil || r.Method != "POST" {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagOne)
		_, _ = w.Write(msg)
		logger.Logger.Error("Illegal request")
	}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagTwo)
		_, _ = w.Write(msg)
		logger.Logger.Errorf("happen a err that is %v", err)
	}

	reply, err := operation.EditGlobalParameter(context.Background(), c.PolicyName)
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

func ListPolicy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/json")
	var c policy
	if r.Body == nil || r.Method != "POST" {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagOne)
		_, _ = w.Write(msg)
		logger.Logger.Error("Illegal request")
	}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagTwo)
		_, _ = w.Write(msg)
		logger.Logger.Errorf("happen a err that is %v", err)
	}

	reply, err := operation.ListPolicies(context.Background(), c.PolicyName)
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