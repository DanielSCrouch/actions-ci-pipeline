//go:build e2e

package e2e_test

import (
	"fmt"
	"github-actions/tests/e2e"
	"net/url"
	"os"
	"strconv"
	"testing"
)

func TestApp(t *testing.T) {

	apiProtocol := os.Getenv("API_PROTOCOL")
	if apiProtocol == "" {
		apiProtocol = "http"
	}

	hostAddr := os.Getenv("HOST_ADDR")
	if hostAddr == "" {
		hostAddr = "127.0.0.1"
	}

	hostPortStr := os.Getenv("HOST_PORT")
	if hostPortStr == "" {
		hostPortStr = "8080"
	}

	hostPort, err := strconv.Atoi(hostPortStr)
	if err != nil {
		panic(err)
	}

	apiVersion := os.Getenv("API_VERSION")
	if apiVersion == "" {
		apiVersion = "1"
	}

	url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", hostAddr, hostPort),
		Path:   "/status",
	}

	t.Logf("testing url: %s", url.String())

	e2e.TestStatusCheck(t, url)
}
