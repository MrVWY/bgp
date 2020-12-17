package handle

import (
	"bgp/logger"
	"bgp/operation"
	"encoding/json"
	"golang.org/x/net/context"
	"net/http"
)

func CreatePeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json")
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

	reply, err := operation.AddPeers(context.Background(), c.Description, c.NeighborAddress, c.LocalAs, c.PeerAs, c.SendCommunity)
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
	w.Header().Set("content-type", "text/json")
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

func UpdatePeer(w http.ResponseWriter, r *http.Request) {

}

func ListPeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/json")
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