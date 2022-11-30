package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestVoteCount(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(serveRoot)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("expected sucessful response status. got %v", status)
	}

    // Validate vote counts
	votes := []Vote{
        {CandidateID: "Roost", VoterID: "1"},
		{CandidateID: "Rancher", VoterID: "2"},
		{CandidateID: "Docker Desktop", VoterID: "3"},
	}
    initialVotesInStore := len(candidateVotesStore)
	for _, v := range votes {
		b, _ := json.Marshal(v)
		req = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
		rec = httptest.NewRecorder()

		handler = http.HandlerFunc(serveRoot)
		handler.ServeHTTP(rec, req)

		if status := rec.Code; status != http.StatusOK {
			t.Errorf("expected successful response status. got %v", status)
		}
	}

	newVotes := len(candidateVotesStore) - initialVotesInStore
	givenVotes := len(votes)
	if newVotes != givenVotes {
		t.Logf("expected total votes count %v, got %v", givenVotes, newVotes)
		t.FailNow()
	}
}
