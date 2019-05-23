package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle
var lang string
var loc *i18n.Localizer

func init() {
	lang = "zh_cn"
	bundle = &i18n.Bundle{DefaultLanguage: language.English}
	loc = i18n.NewLocalizer(bundle, lang)
	if err := loadTranslateFiles("language/"); err != nil {
		panic(err)
	}
}

func loadTranslateFiles(root string) error {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if _, err := bundle.LoadMessageFile(path); err != nil {
			return err
		}
		return nil
	})
	return nil
}

func I18nStr(key string, param ...interface{}) string {
	lc := &i18n.LocalizeConfig{
		MessageID: key,
	}
	if len(param) > 1 {
		lc.TemplateData = map[string]interface{}{
			param[0].(string): param[1],
		}
	}

	translation, err := loc.Localize(lc)
	if err != nil {
		fmt.Println("err=", err)
		return key
	}
	return translation
}

func main() {
	key := "hello_world"
	fmt.Println(lang, I18nStr(key))

	key = "missing_param"
	fmt.Println(lang, I18nStr(key, "Key", "operator"))
}
