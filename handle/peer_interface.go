package handle

import (
	"net/http"
)

func CreatePeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json")

	if r.Body == nil {
		w.WriteHeader(400)
		msg, _ := Json("400", MessageTagOne)
		_, _ = w.Write(msg)
		//日志
	}

	//Description, NeighborAddress string, LocalAs, PeerAs, SendCommunity int

}


func DeletePeer(w http.ResponseWriter, r *http.Request) {

}