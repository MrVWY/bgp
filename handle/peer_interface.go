package handle

import (
	"bgp/logger"
	"bgp/operation"
	"encoding/json"
	"golang.org/x/net/context"
	"net/http"
)

func CreatePeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	var err error
	var c newPeer
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

	reply, err := operation.AddPeers(context.Background(), c.Description, c.NeighborAddress, c.PeerAs, c.SendCommunity)
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


func DeletePeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	var err error
	var c peer
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

	reply, err := operation.DeletePeers(context.Background(), c.NeighborAddress)
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

func AddPolicyToPeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	var err error
	var c policyToPeer
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

	reply, err := operation.AddPolicyToPeer(context.Background(), c.NeighborAddress, c.PolicyAssignmentName, c.Direction, c.RouteAction, c.importPolicyName, c.exportPolicy, c.ImOrOut)
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

func ListPeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	var c peer
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

	reply, err := operation.ListPeers(context.Background(), c.NeighborAddress)
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