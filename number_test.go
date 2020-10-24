package cldr_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"

	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources"
)

const (
	USD = "USD"
)

type fmtCurrencyTest struct {
	inCurrency string
	inNumber   float64
	err        error
	cur        string
}

func currencyRunner(t *testing.T, tests []fmtCurrencyTest, fmtFunc func(string, float64) (string, error)) {
	t.Helper()
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %f", test.inCurrency, test.inNumber), func(t *testing.T) {
			cur, err := fmtFunc(test.inCurrency, test.inNumber)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.cur, cur)
		})
	}
}

func TestFmtCurrency(t *testing.T) {
	en, _ := resources.GetLocale(language.Make("en"))

	tests := []fmtCurrencyTest{
		{
			inCurrency: USD,
			inNumber:   12345.6789,
			cur:        "$12,345.68",
		},
		{
			inCurrency: "WHAT???",
			inNumber:   12345.6789,
			cur:        "12,345.68",
		},
		{
			inCurrency: USD,
			inNumber:   12345000000000.6789,
			cur:        "$12,345,000,000,000.68",
		},
		{
			inCurrency: USD,
			inNumber:   0.0084,
			cur:        "$0.01",
		},
	}

	currencyRunner(t, tests, en.Number.FmtCurrency)
}

func TestFmtCurrencyAccounting(t *testing.T) {
	en, _ := resources.GetLocale(language.Make("en"))

	tests := []fmtCurrencyTest{
		{
			inCurrency: USD,
			inNumber:   12345.6789,
			cur:        "$12,345.68",
		},
		{
			inCurrency: USD,
			inNumber:   -12345.6789,
			cur:        "($12,345.68)",
		},
		{
			inCurrency: USD,
			inNumber:   -12345000000000.6789,
			cur:        "($12,345,000,000,000.68)",
		},
	}

	currencyRunner(t, tests, en.Number.FmtCurrencyAccounting)

	wholeTests := []fmtCurrencyTest{
		{
			inCurrency: USD,
			inNumber:   12345.6789,
			cur:        "$12,346",
		},
		{
			inCurrency: USD,
			inNumber:   -12345.6789,
			cur:        "($12,346)",
		},
		{
			inCurrency: USD,
			inNumber:   12345000000000.6789,
			cur:        "$12,345,000,000,001",
		},
		{
			inCurrency: USD,
			inNumber:   -12345000000000.6789,
			cur:        "($12,345,000,000,001)",
		},
	}

	currencyRunner(t, wholeTests, en.Number.FmtCurrencyAccountingWhole)
}

func TestFmtCurrencyWhole(t *testing.T) {
	en, _ := resources.GetLocale(language.Make("en"))

	tests := []fmtCurrencyTest{
		{
			inCurrency: USD,
			inNumber:   12345.6789,
			cur:        "$12,346",
		},
		{
			inCurrency: USD,
			inNumber:   -12345.6789,
			cur:        "-$12,346",
		},
		{
			inCurrency: "WHAT???",
			inNumber:   12345.6789,
			cur:        "12,346",
		},
		{
			inCurrency: USD,
			inNumber:   12345000000000.6789,
			cur:        "$12,345,000,000,001",
		},
		{
			inCurrency: USD,
			inNumber:   -12345000000000.6789,
			cur:        "-$12,345,000,000,001",
		},
	}

	currencyRunner(t, tests, en.Number.FmtCurrencyWhole)
}

func TestFmtNumber(t *testing.T) {
	en, _ := resources.GetLocale(language.Make("en"))
	// check Hindi - different group sizes
	hi, _ := resources.GetLocale(language.Make("hi"))
	// check Uzbek - something with a partial fallback
	uz, _ := resources.GetLocale(language.Make("uz"))

	tests := []struct {
		locale *cldr.Locale
		in     float64
		out    string
	}{
		{
			locale: en,
			in:     12345.6789,
			out:    "12,345.679",
		},
		{
			locale: en,
			in:     -12345.6789,
			out:    "-12,345.679",
		},
		{
			locale: en,
			in:     123456789,
			out:    "123,456,789",
		},
		{
			locale: hi,
			in:     12345.6789,
			out:    "12,345.679",
		},
		{
			locale: hi,
			in:     -12345.6789,
			out:    "-12,345.679",
		},
		{
			locale: hi,
			in:     123456789,
			out:    "12,34,56,789",
		},
		{
			locale: uz,
			in:     12345.6789,
			out:    "12٬345٫679",
		},
		{
			locale: uz,
			in:     -12345.6789,
			out:    "-12٬345٫679",
		},
		{
			locale: uz,
			in:     123456789,
			out:    "123٬456٬789",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %f", test.locale.Locale, test.in), func(t *testing.T) {
			out := test.locale.Number.FmtNumber(test.in)
			assert.Equal(t, test.out, out)
		})
	}

	assert.Equal(t, "12,346", en.Number.FmtNumberWhole(12345.6789))
	assert.Equal(t, "-12,346", en.Number.FmtNumberWhole(-12345.6789))
}

func TestFmtPercent(t *testing.T) {
	en, _ := resources.GetLocale(language.Make("en"))

	tests := []struct {
		in  float64
		out string
	}{
		{
			in:  0.01234,
			out: "1%",
		},
		{
			in:  0.1234,
			out: "12%",
		},
		{
			in:  1.234,
			out: "123%",
		},
		{
			in:  12.34,
			out: "1,234%",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%f", test.in), func(t *testing.T) {
			out := en.Number.FmtPercent(test.in)
			assert.Equal(t, test.out, out)
		})
	}
}
