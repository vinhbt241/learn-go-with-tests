package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dict Dictionary) Search(key string) (string, error) {
	val, ok := dict[key]

	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}

func (dict Dictionary) Add(key, val string) error {
	_, err := dict.Search(key)

	if err == ErrNotFound {
		dict[key] = val
		return nil
	}

	return ErrWordExists
}

func (dict Dictionary) Update(key, val string) error {
	_, err := dict.Search(key)

	switch err {
	case nil:
		dict[key] = val
	case ErrNotFound:
		return ErrWordDoesNotExist
	}

	return nil
}

func (dict Dictionary) Delete(key string) error {
	_, err := dict.Search(key)

	switch err {
	case nil:
		delete(dict, key)
	case ErrNotFound:
		return ErrWordDoesNotExist
	}

	return nil
}
