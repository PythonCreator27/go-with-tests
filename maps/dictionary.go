package dictionary

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound               = DictionaryErr("cannot find the word you are looking for")
	ErrWordExists             = DictionaryErr("the word that you are trying to add already exists")
	ErrWordDoesNotExistUpdate = DictionaryErr("the word that you are trying to update does not exist")
	ErrWordDoesNotExistDelete = DictionaryErr("the word that you are trying to delete does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, exists := d[word]

	if !exists {
		return "", ErrNotFound
	}

	return definition, nil
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
		return ErrWordDoesNotExistUpdate
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExistDelete
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
