package mydict

import "errors"

type Dictionary map[string]string

// type 은 method 를 가질 수 있다.

var errNotFound = errors.New("Not Found")
var errWordExist = errors.New("That word already exists")

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
