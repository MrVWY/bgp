package handle

import (
	"bgp/logger"
	"bgp/operation"
	"encoding/json"
	"golang.org/x/net/context"
	"net/http"
)

func CreatePrefixSet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/json")
	var c createPrefixSet
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

	reply, err := operation.AddDefinedSetPrefixSet(context.Background(), c.Type, c.PrefixSetName, c.ipPrefix, c.MaskMin, c.MaskMax)
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

func CreateCommunitySet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/json")
	var c createCommunitySet
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

	reply, err := operation.AddDefinedSetCommunitySet(context.Background(), c.CommunitySetName, c.Type, c.list)
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

func CreateNeighborSet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/json")
	var c createNeighborSet
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

	reply, err := operation.AddDefinedSetNeighborSet(context.Background(), c.NeighborSetName, c.Type, c.list)
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

func DeleteDefined(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type","text/json")
	var c deleteDefinedSet
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

	reply, err := operation.DeleteDefinedSets(context.Background(), c.DefinedSetName)
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