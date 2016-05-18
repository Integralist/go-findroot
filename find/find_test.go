package find

import (
	"log"
	"testing"
)

func TestRootIsFound(t *testing.T) {
	response, err := Repo()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	expectation := "go-findroot"

	if response.Name != expectation {
		t.Errorf("The response '%s' didn't match the expectation '%s'", response.Name, expectation)
	}
}
