package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockRepo struct {
}

func (m MockRepo) SaveDNA(dna Dna) error {
	return nil
}

func (m MockRepo) GetStats() (Stats, error) {
	return Stats{CountMutantsDna: 40, CountTotalDnas: 140}, nil
}

func TestController_VerifyDNA_Mutant(t *testing.T) {
	c := Controller{DB: MockRepo{}}
	s := httptest.NewServer(http.HandlerFunc(c.VerifyDNA))
	defer s.Close()
	client := s.Client()
	body := bytes.NewBuffer([]byte(`
{
"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}
`))
	res, err := client.Post(s.URL+"/mutants", "application/json", body)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestController_VerifyDNA_notMutant(t *testing.T) {
	c := Controller{DB: MockRepo{}}
	s := httptest.NewServer(http.HandlerFunc(c.VerifyDNA))
	defer s.Close()
	client := s.Client()
	body := bytes.NewBuffer([]byte(`
{
"dna":["ATGCGTA","CCGTGAA","TTATGTA","AGAAGGG","CACCTAA","TCACTGA", "TCACTGA"]
}
`))
	res, err := client.Post(s.URL+"/mutants", "application/json", body)
	require.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, res.StatusCode)
}

func TestController_VerifyDNA_invalidCharacter(t *testing.T) {
	c := Controller{DB: MockRepo{}}
	s := httptest.NewServer(http.HandlerFunc(c.VerifyDNA))
	defer s.Close()
	client := s.Client()
	body := bytes.NewBuffer([]byte(`
{
"dna":["ATGCGTA","CCGTGAA","TTAJGTA","AGAAGGG","CACCTAA","TCACTGA", "TCACTGA"]
}
`))
	res, err := client.Post(s.URL+"/mutants", "application/json", body)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestController_VerifyDNA_invalidStringLen(t *testing.T) {
	c := Controller{DB: MockRepo{}}
	s := httptest.NewServer(http.HandlerFunc(c.VerifyDNA))
	defer s.Close()
	client := s.Client()
	body := bytes.NewBuffer([]byte(`
{
"dna":["ATGCGTA","CCGTGAA","TTATTGTA","AGAAGGG","CACCTAA","TCACTGA", "TCACTGA"]
}
`))
	res, err := client.Post(s.URL+"/mutants", "application/json", body)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestController_VerifyDNA_invalidArraySize(t *testing.T) {
	c := Controller{DB: MockRepo{}}
	s := httptest.NewServer(http.HandlerFunc(c.VerifyDNA))
	defer s.Close()
	client := s.Client()
	body := bytes.NewBuffer([]byte(`
{
"dna":["ATGCGTA","CCGTGAA","TTATGTA","AGAAGGG","CACCTAA","TCACTGA", "TCACTGA", "AGAAGGG"]
}
`))
	res, err := client.Post(s.URL+"/mutants", "application/json", body)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestController_GetStats(t *testing.T) {
	c := Controller{DB: MockRepo{}}
	s := httptest.NewServer(http.HandlerFunc(c.GetStats))
	defer s.Close()
	client := s.Client()
	res, err := client.Get(s.URL+"/stats")
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)

	var sr StatsResponse
	err = json.Unmarshal(body, &sr)
	assert.Equal(t, StatsResponse{
		CountHumanDna: 100,
		CountMutantDna: 40,
		Ratio: 0.4,
	}, sr)
	assert.NoError(t, err)
}