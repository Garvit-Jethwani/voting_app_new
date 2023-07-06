package main

import (
	"testing"
	"sort"
)

type CandidateVotes struct {
	CandidateID string
	Votes       int
}

type ResultBoard struct {
	Results    []CandidateVotes
	TotalVotes int
}

func getCandidatesVote() map[string]int {
	return map[string]int{
		"A": 10,
		"B": 20,
		"C": 30,
	}
}

func countVote() (res ResultBoard, err error) {
	votes := getCandidatesVote()
	for candidateID, vote := range votes {
		res.Results = append(res.Results, CandidateVotes{candidateID, vote})
		res.TotalVotes += vote
	}

	sort.Slice(res.Results, func(i, j int) bool {
		return res.Results[i].Votes > res.Results[j].Votes
	})
	return res, err
}

func TestCountVote_0b2cc2f08c(t *testing.T) {
	res, err := countVote()
	if err != nil {
		t.Error("Expected no error, got ", err)
	}

	expectedTotalVotes := 60
	if res.TotalVotes != expectedTotalVotes {
		t.Error("Expected total votes to be ", expectedTotalVotes, ", got ", res.TotalVotes)
	}

	expectedWinner := "C"
	if res.Results[0].CandidateID != expectedWinner {
		t.Error("Expected winner to be ", expectedWinner, ", got ", res.Results[0].CandidateID)
	}

	expectedWinnerVotes := 30
	if res.Results[0].Votes != expectedWinnerVotes {
		t.Error("Expected winner votes to be ", expectedWinnerVotes, ", got ", res.Results[0].Votes)
	}
}
