package v1

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// API .
type API interface {
	GetToken() (*GetTokenResponse, error)
	Track(*TrackRequest) (*TrackResponse, error)
	TrackBatch() error
}

// TrackingAPI .
type api struct {
	Token string

	httpClient *http.Client
}

// New .
func New(token string) API {
	return &api{
		Token:      token,
		httpClient: new(http.Client),
	}
}

// GetToken .
func (a *api) GetToken() (*GetTokenResponse, error) {
	req, err := http.NewRequest(http.MethodPost, getTokenURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Token "+a.Token)
	req.Header.Set("Content-Type", contentJSON)

	res, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	payload := GetTokenResponse{}
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

// Track .
func (a *api) Track(rqd *TrackRequest) (*TrackResponse, error) {
	// convert request to json
	buf := bytes.Buffer{}

	json.NewEncoder(&buf).Encode(rqd)

	req, err := http.NewRequest(http.MethodPost, trackURL, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Token "+a.Token)
	req.Header.Set("Content-Type", contentJSON)

	res, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buf.Reset() // reset buffer
	io.Copy(&buf, res.Body)

	ioutil.WriteFile("Track_response.json", buf.Bytes(), 0644)

	payload := TrackResponse{}
	err = json.NewDecoder(&buf).Decode(&payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

// TrackBatch .
func (a *api) TrackBatch() error {
	return nil
}
