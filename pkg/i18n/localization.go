package i18n

import (
	"encoding/json"
	"fmt"
	"github.com/Mohamadreza-shad/auth/config"
	"github.com/Mohamadreza-shad/auth/pkg/exception"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"strings"
)

// I18n holds localization bundle and localizers for supported languages
type I18n struct {
	bundle     *i18n.Bundle
	localizers map[string]*i18n.Localizer
	langs      []string
}

// NewI18n initializes I18n, loading resources from resourcePath for each language in langs
func NewI18n(config config.Config) (I18n, error) {
	resourcePath := config.I18n.ResourcePath
	langs := strings.Split(config.I18n.Languages, ",")
	if langs == nil || len(langs) == 0 {
		return I18n{}, exception.ErrInvalidConfig.WithTrace("languages config is empty")
	}
	b := i18n.NewBundle(language.English)
	b.RegisterUnmarshalFunc("json", json.Unmarshal)

	localizers := make(map[string]*i18n.Localizer)
	for _, lang := range langs {
		langFile := filepath.Join(resourcePath, fmt.Sprintf("%s.json", lang))
		if _, err := os.Stat(langFile); err != nil {
			return I18n{}, exception.ErrMissingI18nResource.WithError(err)
		}
		_, err := b.LoadMessageFile(langFile)
		if err != nil {
			return I18n{}, exception.ErrLoadingI18nResource.WithError(err)
		}
		localizers[lang] = i18n.NewLocalizer(b, lang)
	}

	return I18n{
		bundle:     b,
		localizers: localizers,
		langs:      langs,
	}, nil
}

// GetLocalizedMessage returns localized message for messageKey and lang, or falls back to English
func (i *I18n) GetLocalizedMessage(messageKey, lang string) (string, error) {
	localizer, exists := i.localizers[lang]
	if !exists {
		return "", exception.ErrMissingI18nResource
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageKey})
	if err != nil {
		return "", exception.ErrMissingTranslationKey
	}
	return msg, nil
}
