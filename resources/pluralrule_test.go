package resources

import (
	"strconv"
	"strings"
	"testing"

	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"

	"github.com/razor-1/cldr"
)

type pluralFormTest struct {
	num  interface{}
	form plural.Form
}

func runTests(t *testing.T, tag language.Tag, tests []pluralFormTest) {
	if tag.IsRoot() {
		return
	}
	if rule := cardinalPlural[tag]; rule != nil {
		for _, test := range tests {
			ops, err := cldr.NewOperands(test.num)
			if err != nil {
				t.Errorf("%s: NewOperands(%d) errored with %s", tag.String(), test.num, err)
				break
			}
			pluralData := rule()
			if pluralForm := pluralData.Func(ops); pluralForm != test.form {
				t.Errorf("%s: PluralFormFunc(%#v) returned %q, %v; expected %q", tag.String(), ops, pluralForm, err, test.form)
			}
		}
	} else {
		t.Errorf("could not find plural rule for locale %s", tag.String())
	}

}

func appendIntegerTests(tests []pluralFormTest, form plural.Form, examples []string) []pluralFormTest {
	for _, ex := range expandExamples(examples) {
		i, err := strconv.ParseInt(ex, 10, 64)
		if err != nil {
			panic(err)
		}
		tests = append(tests, pluralFormTest{ex, form}, pluralFormTest{i, form})
	}
	return tests
}

func appendDecimalTests(tests []pluralFormTest, form plural.Form, examples []string) []pluralFormTest {
	for _, ex := range expandExamples(examples) {
		tests = append(tests, pluralFormTest{ex, form})
	}
	return tests
}

func expandExamples(examples []string) []string {
	var expanded []string
	for _, ex := range examples {
		if parts := strings.Split(ex, "~"); len(parts) == 2 {
			for ex := parts[0]; ; ex = increment(ex) {
				expanded = append(expanded, ex)
				if ex == parts[1] {
					break
				}
			}
		} else {
			expanded = append(expanded, ex)
		}
	}
	return expanded
}

func increment(dec string) string {
	runes := []rune(dec)
	carry := true
	for i := len(runes) - 1; carry && i >= 0; i-- {
		switch runes[i] {
		case '.':
			continue
		case '9':
			runes[i] = '0'
		default:
			runes[i]++
			carry = false
		}
	}
	if carry {
		runes = append([]rune{'1'}, runes...)
	}
	return string(runes)
}
