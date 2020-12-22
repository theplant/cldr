package main

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
	"unicode"

	"golang.org/x/text/unicode/cldr"

	i18n "github.com/razor-1/cldr"
)

const (
	resourcesDir = "resources"
	internalDir  = "internal"
	localesDir   = "locales"
	currencyDir  = "currency"
	languageDir  = "language"
	territoryDir = "territory"
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
	LocaleCode    string
	CLDRPackage   string
	Symbols       i18n.Symbols
	NumberFormats i18n.NumberFormats
	Calendar      i18n.Calendar
	Currencies    i18n.Currencies
	Languages     Languages
	Territories   Territories
}

type aggregateData struct {
	Currencies  i18n.Currencies
	Languages   Languages
	Territories Territories
	mutex       sync.Mutex
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

	localeData := processCLDR(unicodeCLDR)

	path := filepath.Join(resourcesDir, internalDir, localesDir)
	makePath(path)

	pluralLocales := pluralRules(unicodeCLDR.Supplemental())
	// some locales seem to show up in plural rules but not in common/main...keep track of everything in allLocales
	for loc := range pluralLocales {
		localeData.Locales[loc] = true
	}

	//aggregate stores everything we've seen across all locales
	//this ensures that we generate our constants packages with all needed data
	aggregate := aggregateData{
		Currencies:  make(i18n.Currencies, 500),
		Languages:   make(Languages, 500),
		Territories: make(Territories, 500),
		mutex:       sync.Mutex{},
	}

	var wg sync.WaitGroup
	wg.Add(len(localeData.Numbers))
	sem := make(chan bool, 20)
	for locale, number := range localeData.Numbers {
		localeData.Locales[locale] = true
		sem <- true
		go func(locale string, number i18n.Number) {
			defer func() {
				<-sem
				wg.Done()
			}()
			localeFile := filepath.Join(path, locale+".go")

			tplData := templateData{
				LocaleCode:    locale,
				CLDRPackage:   cldrPackage,
				Symbols:       number.Symbols,
				NumberFormats: number.Formats,
				Calendar:      localeData.Calendars[locale],
				Currencies:    number.Currencies,
				Languages:     localeData.Languages[locale],
				Territories:   localeData.Territories[locale],
			}
			err = executeAndWrite(filepath.Join(templatesDir, "locales.tpl"), tplData, localeFile)
			if err != nil {
				panic(err)
			}

			// add this locale's info into the aggregate
			aggregate.mutex.Lock()
			defer aggregate.mutex.Unlock()
			for k, v := range tplData.Currencies {
				aggregate.Currencies[k] = v
			}
			for k, v := range tplData.Languages {
				aggregate.Languages[k] = v
			}
			for k, v := range tplData.Territories {
				aggregate.Territories[k] = v
			}

		}(locale, number)
	}
	wg.Wait()

	//create the currency, language and territory constants
	currencyPath := filepath.Join(resourcesDir, currencyDir)
	makePath(currencyPath)
	currencyFile := filepath.Join(currencyPath, "currency.go")
	err = executeAndWrite(filepath.Join(templatesDir, "currency.tpl"), aggregate.Currencies, currencyFile)
	if err != nil {
		panic(err)
	}

	languagePath := filepath.Join(resourcesDir, languageDir)
	makePath(languagePath)
	languageFile := filepath.Join(languagePath, "language.go")
	err = executeAndWrite(filepath.Join(templatesDir, "language.tpl"), aggregate.Languages, languageFile)
	if err != nil {
		panic(err)
	}

	territoryPath := filepath.Join(resourcesDir, territoryDir)
	makePath(territoryPath)
	territoryFile := filepath.Join(territoryPath, "territory.go")
	err = executeAndWrite(filepath.Join(templatesDir, "territory.tpl"), aggregate.Territories, territoryFile)
	if err != nil {
		panic(err)
	}

	//create a mapping of a locale to which tag to use for plurals. this ensures that all the locales, not just the
	//ones in the CLDR with plural rules, have them populated.
	//without doing this, locales like en_US or fr_CA wouldn't have plural rules, since they are only attached to
	//higher level locales like en and fr
	plFuncs := make(map[string]string, len(localeData.Locales))
	for loc := range localeData.Locales {
		plFuncs[loc] = pluralLocale(loc, pluralLocales)
	}

	type allTemplateData struct {
		CLDRPackage      string
		Numbers          Numbers
		Tags             map[string]bool
		PluralLocaleTags map[string]string
	}
	allData := allTemplateData{
		CLDRPackage:      cldrPackage,
		Numbers:          localeData.Numbers,
		Tags:             localeData.Locales,
		PluralLocaleTags: plFuncs,
	}
	allFile := filepath.Join(resourcesDir, "gen_locales.go")
	err = executeAndWrite(filepath.Join(templatesDir, "all.tpl"), allData, allFile)
	if err != nil {
		panic(err)
	}

	plLocalesFiles := filepath.Join(resourcesDir, "gen_plural_locales.go")
	//we don't need the rest of allData, but it's convenient...
	err = executeAndWrite(filepath.Join(templatesDir, "plural_locales.tpl"), allData, plLocalesFiles)
	if err != nil {
		panic(err)
	}
}

//upperIdentifier takes a string and does some basic things to make it a valid identifier. it doesn't do a
//check for reserved words, invalid characters, etc. The entire string is uppercased.
func upperIdentifier(s string) string {
	upper := strings.ToUpper(s)
	var first rune
	for _, r := range upper {
		first = r
		break
	}
	if !unicode.IsLetter(first) {
		upper = "V_" + upper
	}
	return upper
}

func executeAndWrite(templateFile string, data interface{}, outFileName string) error {
	funcMap := template.FuncMap{
		"ToUpperIdent": upperIdentifier,
	}
	tmpl, err := template.New(filepath.Base(templateFile)).Funcs(funcMap).ParseFiles(templateFile)
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
