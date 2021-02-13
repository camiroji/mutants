package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type DnaRequest struct {
	Dna []string `json:"dna"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var dna DnaRequest
	err = json.Unmarshal(body, &dna)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ok := IsMutant(dna.Dna)
	if !ok {
		http.Error(w, "Is not mutant", http.StatusForbidden)
		return
	}

	w.Write([]byte("Is Mutant"))
}

func main(){
	http.HandleFunc("/mutant/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
