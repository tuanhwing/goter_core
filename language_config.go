package gotercore

import language "github.com/moemoe89/go-localization"

func NewLanguageConfig() language.Config {
	var lang *language.Config

	// initiate the go-localization & bind some config
	cfg := language.New()
	// json file location
	cfg.BindPath("./languages.json")
	// default language
	cfg.BindMainLocale("en")

	var err error
	lang, err = cfg.Init()
	if err != nil {
		panic(err)
	}

	return *lang
}
