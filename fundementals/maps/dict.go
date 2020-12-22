package dictionary

const (
	// ErrNotFound signifies a word can't be located
	ErrNotFound = DictErr("Unable to find word")
	// ErrDuplicate signifies a word exists
	ErrDuplicate = DictErr("Duplicate error found")
)

// DictErr is the base type for a dictionary error
type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

// Dictionary maps words to definitions
type Dictionary map[string]string

// Search function return definition, given it exists in the suplied dictionary
func (d Dictionary) Search(word string) (string, error) {
	word, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return word, nil
}

// Add allows you to add a word and defintion to the dictionary
func (d Dictionary) Add(word string, definition string) error {
	if d[word] != "" {
		return ErrDuplicate
	}

	d[word] = definition

	return nil
}

// Update allows one to update an existing dict def
func (d Dictionary) Update(word string, definition string) error {
	if d[word] == "" {
		return ErrNotFound
	}

	d[word] = definition

	return nil
}

// Delete allows one to delete a word from the dictionary
func (d Dictionary) Delete(word string) error {
	if d[word] == "" {
		return ErrNotFound
	}

	delete(d, word)
	return nil
}
