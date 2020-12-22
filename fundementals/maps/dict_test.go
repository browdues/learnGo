package dictionary

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func assertErrorOccurred(t *testing.T, err error) {
	if err == nil {
		t.Fatal("Expected an error, but did not get one.")
	}
}

func assertError(t *testing.T, got error, want error, word string) {
	if got != want {
		t.Errorf("want %q, got %q given %q", got, want, word)
	}
}

func assert(t *testing.T, got string, want string, word string) {
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func TestSearch(t *testing.T) {
	t.Run("finds a word", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "testing 1 2 3",
		}

		got, _ := dictionary.Search("test")
		want := "testing 1 2 3"

		assert(t, got, want, "test")
	})

	t.Run("throws if word not found", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "testing 1 2 3",
		}

		_, err := dictionary.Search("taco")

		assertErrorOccurred(t, err)
		assertError(t, err, ErrNotFound, "taco")
	})

	t.Run("can add a word", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "testing 1 2 3",
		}

		dictionary.Add("bacon", "delicious part of breakfast")

		got, _ := dictionary.Search("bacon")
		want := "delicious part of breakfast"

		assert(t, got, want, "bacon")
	})

	t.Run("diplicates are forbidden", func(t *testing.T) {
		dictionary := Dictionary{
			"test": "testing 1 2 3",
		}

		err := dictionary.Add("test", "testing 1 2 3")

		assertErrorOccurred(t, err)
		assertError(t, err, ErrDuplicate, "test")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("can update an existing entry", func(t *testing.T) {
		word := "test"
		definition := "testing 1 2 3"
		dictionary := Dictionary{word: definition}
		newDefinition := "is this thing on"
		dictionary.Update(word, newDefinition)

		Equal(t, newDefinition, dictionary[word])
	})

	t.Run("returns err if updating an invalid word", func(t *testing.T) {
		word := "test"
		newDef := "testing 1 2 3"
		d := Dictionary{}
		err := d.Update(word, newDef)

		ErrorIs(t, ErrNotFound, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("can delete an entry", func(t *testing.T) {
		d := Dictionary{"Nicholas": "A very handsome man"}

		err := d.Delete("Nicholas")
		NoError(t, err)

		def, err := d.Search("Nicholas")

		ErrorIs(t, ErrNotFound, err)
		Equal(t, "", def)
	})

	t.Run("return err when deleting an invalid entry", func(t *testing.T) {
		d := Dictionary{}

		err := d.Delete("test")

		ErrorIs(t, ErrNotFound, err)
	})
}
