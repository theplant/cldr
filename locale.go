//go:generate go run ./cmd/make_resources

package cldr

import "golang.org/x/text/feature/plural"

type Locale struct {
	Locale   string
	Number   Number
	Calendar Calendar
	Plural   Plural
}

type Plural struct {
	Cardinal PluralData
	Ordinal  PluralData
}

type PluralData struct {
	Forms []plural.Form
	Func  func(ops *PluralOperands) plural.Form
}

//maps a currency code (e.g. "USD") to a Currency
type Currencies map[string]Currency

type Number struct {
	Symbols    Symbols
	Formats    NumberFormats
	Currencies Currencies
}

type Symbols struct {
	Decimal  string
	Group    string
	Negative string
	Percent  string
	PerMille string
}

type NumberFormats struct {
	Decimal  string
	Currency string
	Percent  string
}

type Currency struct {
	DisplayName string
	Symbol      string
}

type Calendar struct {
	Formats     CalendarFormats
	FormatNames CalendarFormatNames
}

type CalendarFormats struct {
	Date     CalendarDateFormat
	Time     CalendarDateFormat
	DateTime CalendarDateFormat
	GMT      string
}

type CalendarDateFormat struct{ Full, Long, Medium, Short string }

type CalendarFormatNames struct {
	Months  CalendarMonthFormatNames
	Days    CalendarDayFormatNames
	Periods CalendarPeriodFormatNames
}

type CalendarMonthFormatNames struct {
	Abbreviated CalendarMonthFormatNameValue
	Narrow      CalendarMonthFormatNameValue
	Short       CalendarMonthFormatNameValue
	Wide        CalendarMonthFormatNameValue
}

type CalendarMonthFormatNameValue struct {
	Jan, Feb, Mar, Apr, May, Jun, Jul, Aug, Sep, Oct, Nov, Dec string
}

type CalendarDayFormatNames struct {
	Abbreviated CalendarDayFormatNameValue
	Narrow      CalendarDayFormatNameValue
	Short       CalendarDayFormatNameValue
	Wide        CalendarDayFormatNameValue
}

type CalendarDayFormatNameValue struct {
	Sun, Mon, Tue, Wed, Thu, Fri, Sat string
}

type CalendarPeriodFormatNames struct {
	Abbreviated CalendarPeriodFormatNameValue
	Narrow      CalendarPeriodFormatNameValue
	Short       CalendarPeriodFormatNameValue
	Wide        CalendarPeriodFormatNameValue
}

type CalendarPeriodFormatNameValue struct {
	AM, PM string
}
