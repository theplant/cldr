package main

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"golang.org/x/text/unicode/cldr"

	i18n "github.com/razor-1/cldr"
)

const (
	resourcesDir = "resources"
	internalDir  = "internal"
	localesDir   = "locales"
	currencyDir  = "currency"
	templatesDir = "cmd/make_resources/templates"
	cldrPackage  = "github.com/razor-1/cldr"

	dateTypeFull   = "full"
	dateTypeLong   = "long"
	dateTypeMedium = "medium"
	dateTypeShort  = "short"

	typeAbbreviated = "abbreviated"
	typeNarrow      = "narrow"
	typeShort       = "short"
	typeWide        = "wide"
)

type templateData struct {
	LocaleCode       string
	PluralLocaleCode string
	CLDRPackage      string
	Symbols          i18n.Symbols
	NumberFormats    i18n.NumberFormats
	Calendar         i18n.Calendar
	Currencies       i18n.Currencies
}

//makePath is a helper to create a path if needed, and panic if MkdirAll encounters an error
func makePath(path string) {
	if _, err := os.Stat(path); err != nil {
		if err = os.MkdirAll(path, 0777); err != nil {
			panic(err)
		}
	}
}

func main() {
	var decoder cldr.Decoder
	//maybe have this decode the zip, and download it if not present
	unicodeCLDR, err := decoder.DecodePath("data/core")
	if err != nil {
		panic(err)
	}

	numbers, calendars, allLocales := processCLDR(unicodeCLDR)

	path := filepath.Join(resourcesDir, internalDir, localesDir)
	makePath(path)

	currencyPath := filepath.Join(resourcesDir, currencyDir)
	makePath(currencyPath)

	pluralLocales := pluralRules(unicodeCLDR.Supplemental())
	// some locales seem to show up in plural rules but not in common/main...keep track of everything in allLocales
	for loc := range pluralLocales {
		allLocales[loc] = true
	}

	var wg sync.WaitGroup
	wg.Add(len(numbers))
	sem := make(chan bool, 20)
	for locale, number := range numbers {
		allLocales[locale] = true
		sem <- true
		go func(locale string, number i18n.Number) {
			defer func() {
				<-sem
				wg.Done()
			}()
			localeFile := filepath.Join(path, locale+".go")

			tplData := templateData{
				LocaleCode:       locale,
				PluralLocaleCode: pluralLocale(locale, pluralLocales),
				CLDRPackage:      cldrPackage,
				Symbols:          number.Symbols,
				NumberFormats:    number.Formats,
				Calendar:         calendars[locale],
				Currencies:       number.Currencies,
			}
			err = executeAndWrite(filepath.Join(templatesDir, "locales.tpl"), tplData, localeFile)
			if err != nil {
				panic(err)
			}

			if locale == "en" {
				// create the currency constants. only need to do this once; picking a common locale
				currencyFile := filepath.Join(currencyPath, "currency.go")
				err = executeAndWrite(filepath.Join(templatesDir, "currency.tpl"), number.Currencies, currencyFile)
				if err != nil {
					panic(err)
				}
			}
		}(locale, number)
	}
	wg.Wait()

	type allTemplateData struct {
		CLDRPackage string
		Numbers     Numbers
		Tags        map[string]bool
	}
	allData := allTemplateData{
		CLDRPackage: cldrPackage,
		Numbers:     numbers,
		Tags:        allLocales,
	}
	allFile := filepath.Join(resourcesDir, "gen_locales.go")
	err = executeAndWrite(filepath.Join(templatesDir, "all.tpl"), allData, allFile)
	if err != nil {
		panic(err)
	}
}

func executeAndWrite(templateFile string, data interface{}, outFileName string) error {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	bb := new(bytes.Buffer)
	err = tmpl.Execute(bb, data)
	if err != nil {
		return err
	}

	formatted, err := format.Source(bb.Bytes())
	if err != nil {
		return err
	}

	outFile, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = outFile.Write(formatted)
	if err != nil {
		return err
	}

	return nil
}
