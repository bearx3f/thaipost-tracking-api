package v1_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	v1 "github.com/bearx3f/thaipost-tracking-api/v1"
)

type config struct {
	Token string `json:"token"`
}

func Test_Track(t *testing.T) {
	cfg := config{}
	fp, err := os.Open("../config.json")
	if err != nil {
		t.Fatalf("%v", err)
	}
	if err := json.NewDecoder(fp).Decode(&cfg); err != nil {
		t.Fatalf("%v", err)
	}

	api := v1.New(cfg.Token)
	res, err := api.Track(&v1.TrackRequest{
		Status:   "all",
		Language: v1.LanguageEN,
		Barcode: []string{
			"LP147586515SG",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v", res)
}
