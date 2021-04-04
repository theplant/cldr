package main

import (
	"fmt"

	"github.com/imdario/mergo"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/cldr"

	i18n "github.com/razor-1/cldr"
)

const (
	defaultGMTFormat = "GMT{0}"
)

type localeData struct {
	Locales        map[string]bool
	Numbers        numbers
	Calendars      calendars
	Languages      map[string]languages
	Territories    map[string]territories
	DisplayPattern map[string]i18n.LocaleDisplayPattern
}

type numbers map[string]i18n.Number
type calendars map[string]i18n.Calendar
type languages map[string]string
type territories map[string]string

func processCLDR(unicodeCLDR *cldr.CLDR) *localeData {
	//size based on a check on 2020-07-25 of how many entries they ended up with: 464 numbers, 358 calendars
	localeData := localeData{
		Locales:        make(map[string]bool, len(unicodeCLDR.Locales())),
		Numbers:        make(numbers, 500),
		Calendars:      make(calendars, 400),
		Languages:      make(map[string]languages, 500),
		Territories:    make(map[string]territories, 350),
		DisplayPattern: make(map[string]i18n.LocaleDisplayPattern, 500),
	}

	//quick & easy way to know if a locale exists
	for _, loc := range unicodeCLDR.Locales() {
		localeData.Locales[loc] = true
	}

	for loc := range localeData.Locales {
		localeData.Numbers[loc], localeData.Calendars[loc], localeData.Languages[loc],
			localeData.Territories[loc], localeData.DisplayPattern[loc] = getCLDRData(localeData.Locales, unicodeCLDR, loc)
	}

	return &localeData
}

//getCLDRData turns CLDR data into our Number and Calendar types, recursively merging data
//so that information from parent locales is inherited. This isn't perfect and doesn't obey all the rules described in
//http://unicode.org/reports/tr35/#Common_Elements, but it should do a pretty good job most of the time.
func getCLDRData(allLocales map[string]bool, unicodeCLDR *cldr.CLDR, loc string) (number i18n.Number,
	calendar i18n.Calendar, languages languages, territories territories, pattern i18n.LocaleDisplayPattern) {
	ldml := unicodeCLDR.RawLDML(loc)
	number = processNumbers(ldml.Numbers)
	calendar = processCalendar(ldml)
	languages = getLanguages(ldml.LocaleDisplayNames)
	territories = getTerritories(ldml.LocaleDisplayNames)
	pattern = getDisplayPattern(ldml.LocaleDisplayNames)

	parentLoc, isRoot := findParentLocale(loc, allLocales)
	if isRoot {
		// loc is already root
		return
	}

	//TODO can we check if parentLoc != loc and only do this in that case?
	parentNumber, parentCalendar, parentLanguages, parentTerritories, parentPattern :=
		getCLDRData(allLocales, unicodeCLDR, parentLoc)

	//merge them
	err := mergo.Merge(&number, parentNumber)
	if err != nil {
		fmt.Println("Number merge error", err)
	}

	//handle the currency map of structs (mergo doesn't do this)
	for curName, parentCurrency := range parentNumber.Currencies {
		mergedCurrency := number.Currencies[curName]
		if mergedCurrency.DisplayName == "" {
			mergedCurrency.DisplayName = parentCurrency.DisplayName
		}
		if mergedCurrency.Symbol == "" {
			mergedCurrency.Symbol = parentCurrency.Symbol
		}
		number.Currencies[curName] = mergedCurrency
	}

	err = mergo.Merge(&calendar, parentCalendar)
	if err != nil {
		fmt.Println("Calendar merge error", err)
	}

	//merge langs and territories
	for k, v := range parentLanguages {
		if _, ok := languages[k]; !ok {
			languages[k] = v
		}
	}

	for k, v := range parentTerritories {
		if _, ok := territories[k]; !ok {
			territories[k] = v
		}
	}

	err = mergo.Merge(&pattern, parentPattern)
	if err != nil {
		fmt.Println("pattern merge error", err)
	}

	return
}

//findParentLocale walks up the inheritance chain and returns the next parent locale that's present in allLocales
//if loc is root, the isRoot bool will be true, and parentLoc will be the empty string "".
func findParentLocale(loc string, allLocales map[string]bool) (parentLoc string, isRoot bool) {
	tag := language.Make(loc)
	if tag.IsRoot() {
		return "", true
	}
	parent := tag.Parent()
	if parent == language.Und {
		// we need to return "root" so that we can inherit from root
		return "root", false
	}
	parentLoc = parent.String()
	if _, ok := allLocales[parentLoc]; ok {
		return parentLoc, false
	}
	// it's not in allLocales; so go higher up the chain by recursing
	return findParentLocale(parentLoc, allLocales)
}

func getNumberSymbols(ldmlNumbers *cldr.Numbers) (symbol i18n.Symbols) {
	if len(ldmlNumbers.Symbols) > 0 {
		ldmlSymbol := ldmlNumbers.Symbols[0]
		if len(ldmlSymbol.Decimal) > 0 {
			symbol.Decimal = ldmlSymbol.Decimal[0].Data()
		}
		if len(ldmlSymbol.Group) > 0 {
			symbol.Group = ldmlSymbol.Group[0].Data()
		}
		if len(ldmlSymbol.MinusSign) > 0 {
			symbol.Negative = ldmlSymbol.MinusSign[0].Data()
		}
		if len(ldmlSymbol.PercentSign) > 0 {
			symbol.Percent = ldmlSymbol.PercentSign[0].Data()
		}
		if len(ldmlSymbol.PerMille) > 0 {
			symbol.PerMille = ldmlSymbol.PerMille[0].Data()
		}
	}

	return
}

func processNumbers(ldmlNumbers *cldr.Numbers) (number i18n.Number) {
	if ldmlNumbers == nil {
		return
	}
	number.Symbols = getNumberSymbols(ldmlNumbers)

	if len(ldmlNumbers.DecimalFormats) > 0 && len(ldmlNumbers.DecimalFormats[0].DecimalFormatLength) > 0 {
		number.Formats.Decimal = ldmlNumbers.DecimalFormats[0].DecimalFormatLength[0].DecimalFormat[0].Pattern[0].Data()
	}
	if len(ldmlNumbers.CurrencyFormats) > 0 && len(ldmlNumbers.CurrencyFormats[0].CurrencyFormatLength) > 0 {
		for _, currencyFormat := range ldmlNumbers.CurrencyFormats[0].CurrencyFormatLength[0].CurrencyFormat {
			switch currencyFormat.Type {
			case "standard":
				number.Formats.Currency = currencyFormat.Pattern[0].Data()
			case "accounting":
				number.Formats.CurrencyAccounting = currencyFormat.Pattern[0].Data()
			}
		}
	}

	if len(ldmlNumbers.PercentFormats) > 0 && len(ldmlNumbers.PercentFormats[0].PercentFormatLength) > 0 {
		number.Formats.Percent = ldmlNumbers.PercentFormats[0].PercentFormatLength[0].PercentFormat[0].Pattern[0].Data()
	}
	if ldmlNumbers.Currencies != nil {
		number.Currencies = make(i18n.Currencies, 350)
		for _, currency := range ldmlNumbers.Currencies.Currency {
			var c i18n.Currency
			if len(currency.DisplayName) > 0 {
				c.DisplayName = currency.DisplayName[0].Data()
			}
			if len(currency.Symbol) > 0 {
				c.Symbol = currency.Symbol[0].Data()
			}
			number.Currencies[currency.Type] = c
		}
	}

	return
}

func processCalendar(ldml *cldr.LDML) (calendar i18n.Calendar) {
	if ldml.Dates == nil || ldml.Dates.Calendars == nil {
		return
	}

	ldmlCar := ldml.Dates.Calendars.Calendar[0]
	for _, cal := range ldml.Dates.Calendars.Calendar {
		if cal.Type == "gregorian" {
			ldmlCar = cal
		}
	}

	if ldml.Dates.TimeZoneNames != nil {
		gmtFormat := ldml.Dates.TimeZoneNames.GmtFormat
		if len(gmtFormat) > 0 && gmtFormat[0].CharData != "" {
			calendar.Formats.GMT = gmtFormat[0].CharData
		} else {
			calendar.Formats.GMT = defaultGMTFormat
		}
	}

	processCalendarDateFormats(ldmlCar, &calendar)
	processCalendarTimeFormats(ldmlCar, &calendar)
	processCalendarMonths(ldmlCar, &calendar)
	processCalendarDays(ldmlCar, &calendar)
	processCalendarDayPeriods(ldmlCar, &calendar)

	return
}

func processCalendarDateFormats(ldmlCar *cldr.Calendar, calendar *i18n.Calendar) {
	if ldmlCar.DateFormats != nil {
		for _, datefmt := range ldmlCar.DateFormats.DateFormatLength {
			switch datefmt.Type {
			case dateTypeFull:
				calendar.Formats.Date.Full = datefmt.DateFormat[0].Pattern[0].Data()
			case dateTypeLong:
				calendar.Formats.Date.Long = datefmt.DateFormat[0].Pattern[0].Data()
			case dateTypeMedium:
				calendar.Formats.Date.Medium = datefmt.DateFormat[0].Pattern[0].Data()
			case dateTypeShort:
				calendar.Formats.Date.Short = datefmt.DateFormat[0].Pattern[0].Data()
			}
		}
	}

	if ldmlCar.DateTimeFormats != nil {
		for _, datefmt := range ldmlCar.DateTimeFormats.DateTimeFormatLength {
			switch datefmt.Type {
			case dateTypeFull:
				calendar.Formats.DateTime.Full = datefmt.DateTimeFormat[0].Pattern[0].Data()
			case dateTypeLong:
				calendar.Formats.DateTime.Long = datefmt.DateTimeFormat[0].Pattern[0].Data()
			case dateTypeMedium:
				calendar.Formats.DateTime.Medium = datefmt.DateTimeFormat[0].Pattern[0].Data()
			case dateTypeShort:
				calendar.Formats.DateTime.Short = datefmt.DateTimeFormat[0].Pattern[0].Data()
			}
		}
	}
}

func processCalendarTimeFormats(ldmlCar *cldr.Calendar, calendar *i18n.Calendar) {
	if ldmlCar.TimeFormats != nil {
		for _, datefmt := range ldmlCar.TimeFormats.TimeFormatLength {
			switch datefmt.Type {
			case dateTypeFull:
				calendar.Formats.Time.Full = datefmt.TimeFormat[0].Pattern[0].Data()
			case dateTypeLong:
				calendar.Formats.Time.Long = datefmt.TimeFormat[0].Pattern[0].Data()
			case dateTypeMedium:
				calendar.Formats.Time.Medium = datefmt.TimeFormat[0].Pattern[0].Data()
			case dateTypeShort:
				calendar.Formats.Time.Short = datefmt.TimeFormat[0].Pattern[0].Data()
			}
		}
	}

}

func processCalendarMonths(ldmlCar *cldr.Calendar, calendar *i18n.Calendar) {
	if ldmlCar.Months != nil {
		for _, monthctx := range ldmlCar.Months.MonthContext {
			for _, months := range monthctx.MonthWidth {
				var i18nMonth i18n.CalendarMonthFormatNameValue
				for _, m := range months.Month {
					setMonthName(m.Type, m.Data(), &i18nMonth)
				}
				switch months.Type {
				case typeAbbreviated:
					calendar.FormatNames.Months.Abbreviated = i18nMonth
				case typeNarrow:
					calendar.FormatNames.Months.Narrow = i18nMonth
				case typeShort:
					calendar.FormatNames.Months.Short = i18nMonth
				case typeWide:
					calendar.FormatNames.Months.Wide = i18nMonth
				}
			}
		}
	}
}

func setMonthName(monthNum, monthName string, i18nMonth *i18n.CalendarMonthFormatNameValue) {
	switch monthNum {
	case "1":
		i18nMonth.Jan = monthName
	case "2":
		i18nMonth.Feb = monthName
	case "3":
		i18nMonth.Mar = monthName
	case "4":
		i18nMonth.Apr = monthName
	case "5":
		i18nMonth.May = monthName
	case "6":
		i18nMonth.Jun = monthName
	case "7":
		i18nMonth.Jul = monthName
	case "8":
		i18nMonth.Aug = monthName
	case "9":
		i18nMonth.Sep = monthName
	case "10":
		i18nMonth.Oct = monthName
	case "11":
		i18nMonth.Nov = monthName
	case "12":
		i18nMonth.Dec = monthName
	}
}

func processCalendarDays(ldmlCar *cldr.Calendar, calendar *i18n.Calendar) {
	if ldmlCar.Days != nil {
		for _, dayctx := range ldmlCar.Days.DayContext {
			for _, days := range dayctx.DayWidth {
				var i18nDay i18n.CalendarDayFormatNameValue
				for _, d := range days.Day {
					setDayName(d.Type, d.Data(), &i18nDay)

				}
				switch days.Type {
				case typeAbbreviated:
					calendar.FormatNames.Days.Abbreviated = i18nDay
				case typeNarrow:
					calendar.FormatNames.Days.Narrow = i18nDay
				case typeShort:
					calendar.FormatNames.Days.Short = i18nDay
				case typeWide:
					calendar.FormatNames.Days.Wide = i18nDay
				}
			}
		}
	}
}

func setDayName(day, dayName string, i18nDay *i18n.CalendarDayFormatNameValue) {
	switch day {
	case "sun":
		i18nDay.Sun = dayName
	case "mon":
		i18nDay.Mon = dayName
	case "tue":
		i18nDay.Tue = dayName
	case "wed":
		i18nDay.Wed = dayName
	case "thu":
		i18nDay.Thu = dayName
	case "fri":
		i18nDay.Fri = dayName
	case "sat":
		i18nDay.Sat = dayName
	}
}

func processCalendarDayPeriods(ldmlCar *cldr.Calendar, calendar *i18n.Calendar) {
	if ldmlCar.DayPeriods != nil {
		for _, ctx := range ldmlCar.DayPeriods.DayPeriodContext {
			for _, width := range ctx.DayPeriodWidth {
				var i18nPeriod i18n.CalendarPeriodFormatNameValue
				for _, d := range width.DayPeriod {
					switch d.Type {
					case "am":
						if i18nPeriod.AM == "" {
							i18nPeriod.AM = d.Data()
						}
					case "pm":
						if i18nPeriod.PM == "" {
							i18nPeriod.PM = d.Data()
						}
					}
				}
				switch width.Type {
				case typeAbbreviated:
					calendar.FormatNames.Periods.Abbreviated = i18nPeriod
				case typeNarrow:
					calendar.FormatNames.Periods.Narrow = i18nPeriod
				case typeShort:
					calendar.FormatNames.Periods.Short = i18nPeriod
				case typeWide:
					calendar.FormatNames.Periods.Wide = i18nPeriod
				}
			}
		}
	}
}

func getLanguages(ldn *cldr.LocaleDisplayNames) (langs languages) {
	langs = make(languages, 500)

	if ldn == nil || ldn.Languages == nil {
		return
	}
	for _, lang := range ldn.Languages.Language {
		langs[lang.Type] = lang.Data()
	}

	return
}

func getTerritories(ldn *cldr.LocaleDisplayNames) (terrs territories) {
	terrs = make(territories, 500)

	if ldn == nil || ldn.Territories == nil {
		return
	}
	for _, terr := range ldn.Territories.Territory {
		if terr.Alt != "" {
			continue
		}
		terrs[terr.Type] = terr.Data()
	}

	return terrs
}

func getDisplayPattern(ldn *cldr.LocaleDisplayNames) (pattern i18n.LocaleDisplayPattern) {
	if ldn == nil || ldn.LocaleDisplayPattern == nil {
		return
	}
	ldp := ldn.LocaleDisplayPattern

	if len(ldp.LocalePattern) > 0 {
		pattern.Pattern = ldp.LocalePattern[0].Data()
	}
	if len(ldp.LocaleSeparator) > 0 {
		pattern.Separator = ldp.LocaleSeparator[0].Data()
	}
	if len(ldp.LocaleKeyTypePattern) > 0 {
		pattern.KeyTypePattern = ldp.LocaleKeyTypePattern[0].Data()
	}

	return
}
