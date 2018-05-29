package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	fromFile()
	fromStruct()

}
func fromFile() {
	bundle := &i18n.Bundle{DefaultLanguage: language.English}
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("ja.toml")
	bundle.MustLoadMessageFile("active.ja.toml")
	bundle.MustLoadMessageFile("active.en.toml")
	bundle.MustLoadMessageFile("active.es.toml")
	localizer := i18n.NewLocalizer(bundle, "en")
	//localized, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: "Hello"})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(localized)
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Hello"}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Hello", PluralCount: 0}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Hello", PluralCount: 1}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Hello", PluralCount: "2.5"}))
	p := map[string]string{"Name": "John"}
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "InvalidMissing", TemplateData: p}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "InvalidMissing", TemplateData: p, PluralCount: 0}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "InvalidMissing", TemplateData: p, PluralCount: 1}))
	p["Do"] = "もういない"
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "InvalidMissing", TemplateData: p, PluralCount: 2}))
}
func fromStruct() {
	bundle := &i18n.Bundle{DefaultLanguage: language.English}
	localizer := i18n.NewLocalizer(bundle, "ja")
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "Hello",
			One: "Hello, one",
			Other: "Hello, other",
		},
	}))
	localizer2 := i18n.NewLocalizer(bundle, "ja")
	fmt.Println(localizer2.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage:&i18n.Message{
			ID: "Invalid",
			One: "{{.Name}} is invalid, one",
			Other: "{{.Name}} are invalid, other",
		},
		PluralCount:1,
	}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "MyUnreadEmails",
			Description: "The number of unread emails I have",
			One:         "I have {{.PluralCount}} unread email.",
			Other:       "I have {{.PluralCount}} unread emails.",
		},
		PluralCount: 1,
	}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "MyUnreadEmails",
			Description: "The number of unread emails I have",
			One:         "I have {{.PluralCount}} unread email.",
			Other:       "I have {{.PluralCount}} unread emails.",
		},
		PluralCount: 2,
	}))


}
