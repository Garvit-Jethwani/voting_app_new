func TestBallot(t *testing.T) {
	port := "8080" // TODO: Replace with the actual port number
	_, result, err := httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		t.Error(err)
		return
	}
	var initalRespData ResultBoard
	if err = json.Unmarshal(result, &initalRespData); err != nil {
		t.Error(err)
		return
	}

	var ballotvotereq Vote
	ballotvotereq.CandidateID = fmt.Sprint(rand.Intn(10))
	ballotvotereq.VoterID = fmt.Sprint(rand.Intn(10))
	reqBuff, err := json.Marshal(ballotvotereq)
	if err != nil {
		t.Error(err)
		return
	}

	_, result, err = httpClientRequest(http.MethodPost, net.JoinHostPort("", port), "/", bytes.NewReader(reqBuff))
	if err != nil {
		t.Error(err)
		return
	}
	var postballotResp Status
	if err = json.Unmarshal(result, &postballotResp); err != nil {
		t.Error(err)
		return
	}
	if postballotResp.Code != 201 {
		t.Error("post ballot resp status code")
		return
	}

	_, result, err = httpClientRequest(http.MethodGet, net.JoinHostPort("", port), "/", nil)
	if err != nil {
		t.Error(err)
		return
	}
	var finalRespData ResultBoard
	if err = json.Unmarshal(result, &finalRespData); err != nil {
		t.Error(err)
		return
	}
	if finalRespData.TotalVotes-initalRespData.TotalVotes != 1 {
		t.Error("ballot vote count error")
		return
	}
}
