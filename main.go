package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {

	bundle := &i18n.Bundle{DefaultLanguage: language.Japanese}
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("./resources/ja.toml")
	lang := "ja"
	accept := "ja,en-US,en;q=0.5"
	localizer := i18n.NewLocalizer(bundle, lang, accept)
	//localized, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: "Hello"})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(localized)
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Hello"}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Hello", PluralCount: 0}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Hello", PluralCount: 1}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Hello", PluralCount: "2.5"}))
	p := map[string]interface{}{"Name": "John"}
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "InvalidMissing", TemplateData: p}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "InvalidMissing", TemplateData: p, PluralCount: 0}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "InvalidMissing", TemplateData: p, PluralCount: 1}))
	p["Do"] = "家に帰っていたようだ"
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "InvalidMissing", TemplateData: p, PluralCount: 2}))
}
