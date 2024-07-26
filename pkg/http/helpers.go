package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func respond(w http.ResponseWriter, resp *Resp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	bts, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("Marshalling Reponse to JSON failed, err")
	}

	_, err = w.Write(bts)
	if err != nil {
		log.Fatal("Writing response to response-writer failed, err:")
	}
}
