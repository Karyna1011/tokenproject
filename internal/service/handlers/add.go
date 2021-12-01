package handlers

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/log"
	"gitlab.com/distributed_lab/ape"
	"io/ioutil"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {
	var p []string
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ape.Render(w, "my array")
	ape.Render(w, p)

	file, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		log.Error("error during marshaling")
	}
	err = ioutil.WriteFile("test.json", file, 0644)
	if err != nil {
		log.Error("error during writing file")
	}

	w.WriteHeader(http.StatusOK)
}
