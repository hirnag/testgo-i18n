package main

import (
	"github.com/nicksnyder/go-i18n/i18n/bundle"
	"fmt"
)

func main() {
	b := bundle.New()
	b.MustLoadTranslationFile("ja.yaml")

	T, err := b.Tfunc("ja")
	if err != nil {
		panic(err)
	}
	fmt.Println(T("program_greeting"))
	fmt.Println(T("person_greeting", "John"))
	fmt.Println(T("person_greeting", map[string]interface{}{
		"Person": "Bob",
	}))
	fmt.Println(T("person_greeting", struct {
		Person string
	}{
		Person: "John",
	}))
	fmt.Println(T("my_height_in_meters", 0))
	fmt.Println(T("my_height_in_meters", 1))
	fmt.Println(T("my_height_in_meters", 2))
	fmt.Println(T("my_height_in_meters", 3))
	fmt.Println(T("your_unread_email_count", map[string]interface{}{
		"Person": "Bob",
		"Count":1,
	}))
	fmt.Println(T("person_unread_email_count" ))
	fmt.Println(T("person_unread_email_count", []interface{}{
		1, struct {
			Person string
			Count int
		}{
			Person: "John",
			Count: 1,
		},
	}))



}
