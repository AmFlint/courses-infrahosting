package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var endpoint = "http://localhost:8080/greetings"

type greetingsResponse struct {
	Message string
}

func TestGreetingsHandler(t *testing.T) {
	if customEndpoint := os.Getenv("TEST_ENDPOINT"); customEndpoint != "" {
		endpoint = customEndpoint
	}

	t.Run("send person testing", func(t *testing.T) {
		person := "testing"

		response, err := http.Get(fmt.Sprintf("%s?person=%s", endpoint, person))
		if err != nil {
			t.Errorf("Got error: %v", err)
			return
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Errorf("got error while reading response body: %v", err)
			return
		}

		greetings := greetingsResponse{}
		if err = json.Unmarshal(bodyBytes, &greetings); err != nil {
			t.Errorf("got error while unmarshalling response body: %v", err)
			return
		}

		expected := fmt.Sprintf("Hello, %s", person)
		if greetings.Message != expected {
			t.Errorf("got message %s, expected %s", greetings.Message, expected)
			return
		}
	})

	t.Run("send person testing", func(t *testing.T) {
		response, err := http.Get(endpoint)
		if err != nil {
			t.Errorf("Got error: %v", err)
			return
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Errorf("got error while reading response body: %v", err)
			return
		}

		greetings := greetingsResponse{}
		if err = json.Unmarshal(bodyBytes, &greetings); err != nil {
			t.Errorf("got error while unmarshalling response body: %v", err)
			return
		}

		expected := "Hello, World"
		if greetings.Message != expected {
			t.Errorf("got message %s, expected %s", greetings.Message, expected)
			return
		}
	})
}
