package main

import "errors"

var (
	ErrNotFound = errors.New("there's no such word")
	ErrWordExists = errors.New("this word already exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if result, ok := d[word]; ok {
		return result, nil
	} else {
		return "", ErrNotFound
	}
}

func (d Dictionary) Add(word, definition string) (error) {
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

func (d Dictionary) Update(word, definition string) (error) {
	_, err := d.Search(word)
	
	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		d[word] = definition
	default:
		return err
	}
	
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}