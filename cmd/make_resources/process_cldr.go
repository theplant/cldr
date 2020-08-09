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

type Numbers map[string]i18n.Number
type Calendars map[string]i18n.Calendar

func processCLDR(unicodeCLDR *cldr.CLDR) (Numbers, Calendars, map[string]bool) {
	//size based on a check on 2020-07-25 of how many entries they ended up with: 464 numbers, 358 calendars
	numbers := make(Numbers, 500)
	calendars := make(Calendars, 400)

	//quick & easy way to know if a locale exists
	allLocales := make(map[string]bool, len(unicodeCLDR.Locales()))
	for _, loc := range unicodeCLDR.Locales() {
		allLocales[loc] = true
	}

	for loc := range allLocales {
		numbers[loc], calendars[loc] = getCLDRData(allLocales, unicodeCLDR, loc)
	}

	return numbers, calendars, allLocales
}

//getCLDRData turns CLDR data into our Number and Calendar types, recursively merging data
//so that information from parent locales is inherited. This isn't perfect and doesn't obey all the rules described in
//http://unicode.org/reports/tr35/#Common_Elements, but it should do a pretty good job most of the time.
func getCLDRData(allLocales map[string]bool, unicodeCLDR *cldr.CLDR, loc string) (number i18n.Number, calendar i18n.Calendar) {
	ldml := unicodeCLDR.RawLDML(loc)
	number = processNumbers(ldml.Numbers)
	calendar = processCalendar(ldml)

	parentLoc, isRoot := findParentLocale(loc, allLocales)
	if isRoot {
		// loc is already root
		return
	}

	parentNumber, parentCalendar := getCLDRData(allLocales, unicodeCLDR, parentLoc)

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

func processNumbers(ldmlNumbers *cldr.Numbers) (number i18n.Number) {
	if ldmlNumbers == nil {
		return
	}
	if len(ldmlNumbers.Symbols) > 0 {
		symbol := ldmlNumbers.Symbols[0]
		if len(symbol.Decimal) > 0 {
			number.Symbols.Decimal = symbol.Decimal[0].Data()
		}
		if len(symbol.Group) > 0 {
			number.Symbols.Group = symbol.Group[0].Data()
		}
		if len(symbol.MinusSign) > 0 {
			number.Symbols.Negative = symbol.MinusSign[0].Data()
		}
		if len(symbol.PercentSign) > 0 {
			number.Symbols.Percent = symbol.PercentSign[0].Data()
		}
		if len(symbol.PerMille) > 0 {
			number.Symbols.PerMille = symbol.PerMille[0].Data()
		}
	}
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
	//if len(ldmlNumbers.CurrencyFormats) > 0 && len(ldmlNumbers.CurrencyFormats[0].CurrencyFormatLength) > 0 {
	//	number.Formats.Currency = ldmlNumbers.CurrencyFormats[0].CurrencyFormatLength[0].CurrencyFormat[0].Pattern[0].Data()
	//}
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
	if ldml.Dates != nil && ldml.Dates.Calendars != nil {
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
		if ldmlCar.Months != nil {
			for _, monthctx := range ldmlCar.Months.MonthContext {
				for _, months := range monthctx.MonthWidth {
					var i18nMonth i18n.CalendarMonthFormatNameValue
					for _, m := range months.Month {
						switch m.Type {
						case "1":
							i18nMonth.Jan = m.Data()
						case "2":
							i18nMonth.Feb = m.Data()
						case "3":
							i18nMonth.Mar = m.Data()
						case "4":
							i18nMonth.Apr = m.Data()
						case "5":
							i18nMonth.May = m.Data()
						case "6":
							i18nMonth.Jun = m.Data()
						case "7":
							i18nMonth.Jul = m.Data()
						case "8":
							i18nMonth.Aug = m.Data()
						case "9":
							i18nMonth.Sep = m.Data()
						case "10":
							i18nMonth.Oct = m.Data()
						case "11":
							i18nMonth.Nov = m.Data()
						case "12":
							i18nMonth.Dec = m.Data()
						}
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
		if ldmlCar.Days != nil {
			for _, dayctx := range ldmlCar.Days.DayContext {
				for _, days := range dayctx.DayWidth {
					var i18nDay i18n.CalendarDayFormatNameValue
					for _, d := range days.Day {
						switch d.Type {
						case "sun":
							i18nDay.Sun = d.Data()
						case "mon":
							i18nDay.Mon = d.Data()
						case "tue":
							i18nDay.Tue = d.Data()
						case "wed":
							i18nDay.Wed = d.Data()
						case "thu":
							i18nDay.Thu = d.Data()
						case "fri":
							i18nDay.Fri = d.Data()
						case "sat":
							i18nDay.Sat = d.Data()
						}
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

	return
}
