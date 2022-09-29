package app_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github-actions/cmd/app"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestApp(t *testing.T) {

	hostAddr := "127.0.0.1"
	hostPort := 8080

	server := app.App{
		HostAddr: hostAddr,
		HostPort: hostPort,
	}

	go func() {
		server.Start()
	}()
	time.Sleep(1000 * time.Millisecond)

	cleanup := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Stop(ctx)
	}
	t.Cleanup(cleanup)

	resp, err := http.Get(fmt.Sprintf("http://%s:%d/status", hostAddr, hostPort))
	if err != nil {
		t.Fatalf("failed to get status, error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("response status '%d', want '%d'", resp.StatusCode, http.StatusOK)
	}

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body, error: %s", err)
	}

	status := app.Status{}
	err = json.Unmarshal(payload, &status)
	if err != nil {
		t.Fatalf("unexpected response body, error: %s", err)
	}

	if status.Status != "online" {
		t.Fatalf("unexpected response status, want %s, got %s", "online", status.Status)
	}

}
