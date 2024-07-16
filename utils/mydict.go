package utils

import "errors"

type Dictionary map[string]string

// type 은 method 를 가질 수 있다.

var errNotFound = errors.New("not found")
var errCantupdate = errors.New("cant update non-existing word")
var errWordExist = errors.New("that word already exists")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}

	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExist
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantupdate
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
