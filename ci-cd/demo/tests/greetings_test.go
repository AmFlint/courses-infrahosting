package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const endpoint = "http://localhost:8080/greetings"

type greetingsResponse struct {
	Message string
}

func TestGreetingsHandler(t *testing.T) {
	t.Run("send person testing", func(t *testing.T) {
		person := "testing"

		response, err := http.Get(fmt.Sprintf("%s?person=%s", endpoint, person))
		if err != nil {
			t.Errorf("Got error: %v", err)
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Errorf("got error while reading response body: %v", err)
		}

		greetings := greetingsResponse{}
		if err = json.Unmarshal(bodyBytes, &greetings); err != nil {
			t.Errorf("got error while unmarshalling response body: %v", err)
		}

		expected := fmt.Sprintf("Hello, %s", person)
		if greetings.Message != expected {
			t.Errorf("got message %s, expected %s", greetings.Message, expected)
		}
	})

	t.Run("send person testing", func(t *testing.T) {
		response, err := http.Get(endpoint)
		if err != nil {
			t.Errorf("Got error: %v", err)
		}

		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Errorf("got error while reading response body: %v", err)
		}

		greetings := greetingsResponse{}
		if err = json.Unmarshal(bodyBytes, &greetings); err != nil {
			t.Errorf("got error while unmarshalling response body: %v", err)
		}

		expected := "Hello, World"
		if greetings.Message != expected {
			t.Errorf("got message %s, expected %s", greetings.Message, expected)
		}
	})
}
