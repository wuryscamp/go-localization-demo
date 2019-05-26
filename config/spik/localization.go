package spik

import (
	"io/ioutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	yaml "gopkg.in/yaml.v2"
)

// Spik struct
type Spik struct {
	bundle *i18n.Bundle
}

// New reads language files and registers them in i18n bundle
func New(langFiles []string) (*Spik, error) {
	// Bundle stores a set of messages
	bundle := &i18n.Bundle{DefaultLanguage: language.English}

	// Enable bundle to understand yaml
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	var translations []byte
	var err error
	for _, file := range langFiles {

		// Read our language yaml file
		translations, err = ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		// It parses the bytes in buffer to add translations to the bundle
		bundle.MustParseMessageFileBytes(translations, file)
	}

	return &Spik{bundle}, nil
}

// Translate function will lookup key and get the value from localization file and translate it to specific language
func (s *Spik) Translate(key, locale string) (string, error) {
	localizer := i18n.NewLocalizer(s.bundle, locale)
	msg, err := localizer.Localize(
		&i18n.LocalizeConfig{
			MessageID: key,
		},
	)
	return msg, err
}

func ParseMessage() {

}
