package main

import (
	"errors"
	"fmt"
)

var (
	errNotFound error = errors.New("not found")
)

type languageErr struct {
	lang string
	err  error
}

func (le languageErr) Error() string {
	return fmt.Sprintf("%s language error: %v", le.lang, le.err)
}

func (le languageErr) Unwrap() error {
	return le.err
}

// -----------------------------------------------------------------------

func getValue(m map[string]string, key string) (string, error) {
	val, ok := m[key]
	if !ok {
		return "", errNotFound
	}
	return val, nil
}

type languages map[string]string

func (l languages) describe(lang string) (string, error) {
	descr, err := getValue(l, lang)
	if err != nil {
		return "", languageErr{lang, err}
	}
	return descr, nil
}

var langs languages = languages{
	"go":     "is awesome",
	"python": "is everywhere",
	"php":    "just is",
}

// -----------------------------------------------------------------------

func main() {
	{
		langName := "go"
		if descr, err := langs.describe(langName); err != nil {
			panic(fmt.Sprintf("unexpected error %v", err))
		} else {
			fmt.Printf("descr for `%v`: %v\n", langName, descr)
		}
	}

	{
		descr, err := langs.describe("java")
		err = fmt.Errorf("wrap once more: %w", err)

		var langErr languageErr // concrete type!
		if errors.As(err, &langErr) {
			fmt.Println("language error:", langErr.lang)
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(descr)
	}
}
