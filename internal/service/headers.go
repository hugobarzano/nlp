
package service

import (
	"encoding/json"
	"log"
	"net/http"
	"neuron-orch/internal/models"
)


func Ping(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Ready to orchestrate!")
}

func Receiver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	obj := &models.Msg{}
	err:= json.NewDecoder(r.Body).Decode(obj)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(obj)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
}