package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"io/ioutil"
	"net/http"
	"strings"
	"regexp"
)

type Controller struct {
	DB Repository
}

type DnaRequest struct {
	Dna []string `json:"dna"`
}

type StatsResponse struct {
	CountMutantDna int     `json:"count_mutant_dna"`
	CountHumanDna  int     `json:"count_human_dna"`
	Ratio          float32 `json:"ratio"`
}

func (c Controller) VerifyDNA(w http.ResponseWriter, r *http.Request) {
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

	if !c.validateInput(dna.Dna) {
		http.Error(w, "Invalid character or size", http.StatusBadRequest)
		return
	}

	ok := IsMutant(dna.Dna)
	err = c.DB.SaveDNA(Dna{dna: aws.String(strings.Join(dna.Dna, "")), isMutant: aws.Bool(ok)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, "Is not mutant", http.StatusForbidden)
		return
	}

	w.Write([]byte("Is Mutant"))
}

func (c Controller) validateInput(s []string) bool {
	for _, d := range s {
		if !c.validLetters(d) || !c.validSize(len(s), d) {
			return false
		}
	}
	return true
}

func (c Controller) validLetters(s string) bool {
	hasInvalidLetters, err := regexp.MatchString("[^ATCG]", s)
	return err == nil && !hasInvalidLetters
}

func (c Controller) validSize(n int, s string) bool {
	return len(s) == n
}

func (c Controller) GetStats(w http.ResponseWriter, r *http.Request) {
	queryResponse, err := c.DB.GetStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	countHumansDna := queryResponse.CountTotalDnas - queryResponse.CountMutantsDna
	response := StatsResponse{
		CountMutantDna: queryResponse.CountMutantsDna,
		CountHumanDna:  countHumansDna,
		Ratio:          float32(queryResponse.CountMutantsDna) / float32(countHumansDna),
	}
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(json); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
