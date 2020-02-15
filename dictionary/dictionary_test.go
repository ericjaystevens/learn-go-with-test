package dictionary

import (
	"testing"
)

func TestDictionary(t *testing.T) {

	t.Run("test look up", func(t *testing.T) {
		dict := GoDict{"pointer": "refrence to Data", "map": "key value store"}
		assertString(t, dict, "key value store")
	})
	t.Run("look up without entry", func(t *testing.T) {
		dict := GoDict{"pointer": "refrence to Data", "map": "key value store"}
		_, e := dict.Search("pizza")
		want := "could not find the word you were looking for"
		if e == nil {
			t.Errorf("expected Error, but did not get one")
		}
		if e.Error() != want {
			t.Errorf("got '%s' wanted '%s'", e.Error(), want)
		}

	})
}

func assertString(t *testing.T, dict GoDict, want string) {
	t.Helper()
	got, _ := dict.Search("map")
	if got != want {
		t.Errorf("got %s but wanted %s", got, want)
	}

}
