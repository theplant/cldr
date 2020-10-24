package resources

import (
	"errors"

	"github.com/razor-1/cldr"
	"golang.org/x/text/language"
)

//GetLocale retrieves the cldr.Locale object for the supplied language tag.
func GetLocale(tag language.Tag) (*cldr.Locale, error) {
	locFunc, ok := localeData[tag]
	if ok {
		loc := locFunc()
		if plFunc, ok := cardinalPluralLocales[tag]; ok {
			loc.Plural = cldr.Plural{Cardinal: plFunc()}
		}
		return loc, nil
	}

	return nil, errors.New("no data for tag")
}
