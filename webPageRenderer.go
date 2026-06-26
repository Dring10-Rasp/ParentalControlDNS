package main

import (
	"embed"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var localeFS embed.FS

var i18nBundle *i18n.Bundle

type Page struct {
	Title string
	Body  []byte
}

func initI18n() {
	// 1. Set fallback language
	i18nBundle = i18n.NewBundle(language.English)
	i18nBundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// 2. Load translation files from embedded FS
	i18nBundle.LoadMessageFileFS(localeFS, "active.en.toml")
	i18nBundle.LoadMessageFileFS(localeFS, "active.es.toml")
}
func getLocalizer(r *http.Request) *i18n.Localizer {
	langParam := r.URL.Query().Get("lang")
	acceptHeader := r.Header.Get("Accept-Language")

	// Order of preference: URL param -> Header -> Bundle default (English)
	return i18n.NewLocalizer(i18nBundle, langParam, acceptHeader)
}
func lookup_Text(r *http.Request, id string) (*string, error) {
	localizer := getLocalizer(r)

	Message, errorLocalizer := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: id,
	})
	if errorLocalizer != nil {
		return nil, errorLocalizer
	}
	return &Message, nil
}

func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
