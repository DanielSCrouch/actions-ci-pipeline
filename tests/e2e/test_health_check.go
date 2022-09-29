package e2e

import (
	"encoding/json"
	"github-actions/cmd/app"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func handleTestError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestStatusCheck(t *testing.T, url url.URL) {
	t.Cleanup(func() {})

	resp, err := http.Get(url.String())
	handleTestError(t, err)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got status code '%d', want '%d'", resp.StatusCode, http.StatusOK)
	}

	payload, err := ioutil.ReadAll(resp.Body)
	handleTestError(t, err)

	status := app.Status{}
	err = json.Unmarshal(payload, &status)
	handleTestError(t, err)

	if status.Status != "online" {
		t.Fatalf("got app status '%s', want '%s'", status.Status, "online")
	}

}
