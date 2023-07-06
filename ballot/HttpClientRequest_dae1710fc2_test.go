func httpClientRequest(method, url, path string, body io.Reader) (int, []byte, error) {
    client := &http.Client{}
    req, err := http.NewRequest(method, url+path, body)
    if err != nil {
        return 0, nil, err
    }

    resp, err := client.Do(req)
    if err != nil {
        return 0, nil, err
    }
    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return 0, nil, err
    }

    return resp.StatusCode, respBody, nil
}
