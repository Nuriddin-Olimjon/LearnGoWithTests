package maps

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("given word already exists in dictionary")
	ErrWordDoesNotExist = DictionaryErr("given word does not exist in dictionary")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
