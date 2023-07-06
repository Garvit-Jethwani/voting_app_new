package main

import (
	"testing"
	"sort"
)

type ResultBoard struct {
	Results []CandidateVotes
	TotalVotes int
}

type CandidateVotes struct {
	ID string
	Votes int
}

var getCandidatesVote func() map[string]int

func countVote() (res ResultBoard, err error) {
	votes := getCandidatesVote()
	for candidateID, voteCount := range votes {
		res.Results = append(res.Results, CandidateVotes{candidateID, voteCount})
		res.TotalVotes += voteCount
	}

	sort.Slice(res.Results, func(i, j int) bool {
		return res.Results[i].Votes > res.Results[j].Votes
	})
	return res, err
}

func TestCountVote_0b2cc2f08c(t *testing.T) {
	// Test case 1: Normal scenario
	getCandidatesVote = func() map[string]int {
		return map[string]int{
			"Candidate1": 100,
			"Candidate2": 200,
			"Candidate3": 150,
		}
	}
	res, err := countVote()
	if err != nil {
		t.Error("Failed to count votes: ", err)
	}
	if res.TotalVotes != 450 {
		t.Error("Total votes count is incorrect")
	}
	if len(res.Results) != 3 {
		t.Error("Results count is incorrect")
	}
	if res.Results[0].ID != "Candidate2" || res.Results[0].Votes != 200 {
		t.Error("Sorting of results is incorrect")
	}

	// Test case 2: No votes
	getCandidatesVote = func() map[string]int {
		return map[string]int{}
	}
	res, err = countVote()
	if err != nil {
		t.Error("Failed to count votes: ", err)
	}
	if res.TotalVotes != 0 {
		t.Error("Total votes count is incorrect")
	}
	if len(res.Results) != 0 {
		t.Error("Results count is incorrect")
	}
}
