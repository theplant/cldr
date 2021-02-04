package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/unicode/cldr"
)

const (
	//pluralTypeOrdinal  = "ordinal"
	pluralTypeCardinal = "cardinal"
)

//pluralRules generates plural_ordinals.go and plural_cardinals.go based on the CLDR data. It returns a map of all
//locale codes for which there are plural rules, mapped to their function name
func pluralRules(data *cldr.SupplementalData) map[string]string {
	//pluralGroupsOrdinal := make([]PluralGroup, 0, 50)
	pluralGroupsCardinal := make([]PluralGroup, 0, 50)
	pluralLocales := make(map[string]bool, 250)
	for _, plurals := range data.Plurals {
		var pluralGroup *[]PluralGroup
		//pluralGroup := new([]PluralGroup)
		switch plurals.Type {
		//case pluralTypeOrdinal:
		//	pluralGroup = &pluralGroupsOrdinal
		case pluralTypeCardinal:
			pluralGroup = &pluralGroupsCardinal
		default:
			// unknown or no plural type. this exists in CLDR in pluralRanges, for example
			continue
		}
		for _, pg := range plurals.PluralRules {
			for _, loc := range strings.Split(pg.Locales, " ") {
				pluralLocales[loc] = true
			}
			rules := make([]PluralRule, len(pg.PluralRule))
			for i, pr := range pg.PluralRule {
				rules[i] = PluralRule{
					Count: pr.Count,
					Rule:  pr.CharData,
				}
			}
			*pluralGroup = append(*pluralGroup, PluralGroup{
				Locales:     pg.Locales,
				PluralRules: rules,
			})
		}
	}

	pluralTemplate := filepath.Join(templatesDir, "pluralfuncs.tpl")
	pluralTestTemplate := filepath.Join(templatesDir, "plurals_test.tpl")

	type pluralTemplateData struct {
		CLDRPackage  string
		PluralType   string
		PluralGroups []PluralGroup
		LocaleMap    map[string]string
	}
	//ordinalData := pluralTemplateData{
	//	PluralType:   strings.Title(pluralTypeOrdinal),
	//	PluralGroups: pluralGroupsOrdinal,
	//}
	//ordinalsFile := filepath.Join(resourcesDir, "aaa_plural_ordinals.go")
	//err := executeAndWrite(pluralTemplate, ordinalData, ordinalsFile)
	//if err != nil {
	//	panic(err)
	//}
	//ordinalsTestFile := filepath.Join(resourcesDir, "aaa_plural_ordinals_test.go")
	//err = executeAndWrite(pluralTestTemplate, ordinalData, ordinalsTestFile)
	//if err != nil {
	//	panic(err)
	//}
	localeMap := make(map[string]string)
	for i, group := range pluralGroupsCardinal {
		for _, locale := range group.SplitLocales() {
			localeMap[locale] = fmt.Sprintf("cardinal_%d", i)
		}
	}

	cardinalData := pluralTemplateData{
		CLDRPackage:  cldrPackage,
		PluralType:   strings.Title(pluralTypeCardinal),
		PluralGroups: pluralGroupsCardinal,
		LocaleMap:    localeMap,
	}
	cardinalsFile := filepath.Join(resourcesDir, "gen_plural_cardinals.go")
	err := executeAndWrite(pluralTemplate, cardinalData, cardinalsFile)
	if err != nil {
		panic(err)
	}
	cardinalsTestFile := filepath.Join(resourcesDir, "gen_plural_cardinals_test.go")
	err = executeAndWrite(pluralTestTemplate, cardinalData, cardinalsTestFile)
	if err != nil {
		panic(err)
	}

	return localeMap
}

//pluralLocaleFunc finds the best match for loc that exists in pluralLocales and returns the func name
//e.g. cardinal_1
func pluralLocale(loc string, pluralLocaleFuncs map[string]string) string {
	tag := language.Make(loc)
	for {
		tagStr := tag.String()
		if tag.IsRoot() {
			// there is no "und" plural rule
			return pluralLocaleFuncs["root"]
		}
		if plFunc, ok := pluralLocaleFuncs[tagStr]; ok {
			return plFunc
		}
		tag = tag.Parent()
	}
}
