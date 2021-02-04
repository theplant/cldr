//go:generate go run ./cmd/make_resources

package cldr

import "golang.org/x/text/feature/plural"

//Languages maps a language code to the localized name of that language
type Languages map[string]string

//Territories maps a territory code to the localized name of the territory
type Territories map[string]string

//Locale has all the relevant information from the CLDR for a locale
type Locale struct {
	Locale      string
	Number      Number
	Calendar    Calendar
	Plural      Plural
	Languages   Languages
	Territories Territories
}

//Plural contains both cardinal and ordinal plural data from the CLDR
type Plural struct {
	Cardinal PluralData
	Ordinal  PluralData
}

//PluralData contains information about the plural forms for the locale and how to get plural operators
type PluralData struct {
	Forms []plural.Form
	Func  func(ops *PluralOperands) plural.Form
}

//Currencies maps a currency code (e.g. "USD") to a Currency
type Currencies map[string]Currency

//Number contains information required for locale-specific number formatting
type Number struct {
	Symbols    Symbols
	Formats    NumberFormats
	Currencies Currencies
}

//Symbols has the various CLDR symbols for the locale
type Symbols struct {
	Decimal  string
	Group    string
	Negative string
	Percent  string
	PerMille string
}

//NumberFormats contains the different format strings from the CLDR
type NumberFormats struct {
	Decimal            string
	Currency           string
	CurrencyAccounting string
	Percent            string
}

//Currency has the basic currency information
type Currency struct {
	DisplayName string
	Symbol      string
}

//Calendar is populated with the CLDR calendar data
type Calendar struct {
	Formats     CalendarFormats
	FormatNames CalendarFormatNames
}

//CalendarFormats is populated with the CLDR calendar data
type CalendarFormats struct {
	Date     CalendarDateFormat
	Time     CalendarDateFormat
	DateTime CalendarDateFormat
	GMT      string
}

//CalendarDateFormat holds different date format strings
type CalendarDateFormat struct{ Full, Long, Medium, Short string }

//CalendarFormatNames contains various data related to localized names for calendar items
type CalendarFormatNames struct {
	Months  CalendarMonthFormatNames
	Days    CalendarDayFormatNames
	Periods CalendarPeriodFormatNames
}

//CalendarMonthFormatNames is the different forms of months
type CalendarMonthFormatNames struct {
	Abbreviated CalendarMonthFormatNameValue
	Narrow      CalendarMonthFormatNameValue
	Short       CalendarMonthFormatNameValue
	Wide        CalendarMonthFormatNameValue
}

//CalendarMonthFormatNameValue has a localized string for each month
type CalendarMonthFormatNameValue struct {
	Jan, Feb, Mar, Apr, May, Jun, Jul, Aug, Sep, Oct, Nov, Dec string
}

//CalendarDayFormatNames is the different forms of days
type CalendarDayFormatNames struct {
	Abbreviated CalendarDayFormatNameValue
	Narrow      CalendarDayFormatNameValue
	Short       CalendarDayFormatNameValue
	Wide        CalendarDayFormatNameValue
}

//CalendarDayFormatNameValue has a localized string for each day of the week
type CalendarDayFormatNameValue struct {
	Sun, Mon, Tue, Wed, Thu, Fri, Sat string
}

//CalendarPeriodFormatNames are the different forms of time periods
type CalendarPeriodFormatNames struct {
	Abbreviated CalendarPeriodFormatNameValue
	Narrow      CalendarPeriodFormatNameValue
	Short       CalendarPeriodFormatNameValue
	Wide        CalendarPeriodFormatNameValue
}

//CalendarPeriodFormatNameValue has a localized string for each time period
type CalendarPeriodFormatNameValue struct {
	AM, PM string
}
