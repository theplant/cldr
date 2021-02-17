# CLDR

[![GoDoc](https://godoc.org/github.com/ContextLogic/cldr?status.svg)](http://godoc.org/github.com/ContextLogic/cldr)

cldr is a golang library using Common Locale Data Repository to format dates, plurals (and more in the future), inspired by [twitter-cldr-rb](https://github.com/twitter/twitter-cldr-rb) and borrowing some codes from [github.com/vube/i18n](https://github.com/vube/i18n).

# How to use

cldr embeds CLDR data in pure go and it doesn't import all those locale data by default. If you are using specific locale data, you can import that package as bellow:

```go
package main

import (
	"github.com/ContextLogic/cldr"
	_ "github.com/ContextLogic/cldr/resources/locales/en"
)

func main() {
	cldr.Parse(
		"en",
		`{{p "Count" (one "{{.Count}} item") (other "{{.Count}} items")}}`,
		map[string]int{"Count": 1},
	) // "1 item in Your Cart"
}
```

If you don't like hand-importing locales, you can import `github.com/ContextLogic/cldr/resources/locales`, which import all available locales in cldr pacakge.

More API could be found [here](https://godoc.org/github.com/ContextLogic/cldr).

# How to add locales

```go
cldr.RegisterLocale(Locale{...})
```

# How to override default locales

```go
// solution 1
// using the same locale name

import _ github.com/ContextLogic/cldr/resources/locales/en
cldr.RegisterLocale(Locale{Locale: "en"})

// solution 2
// update the exported locale directly

import github.com/ContextLogic/cldr/resources/locales/en
en.Locale.PluralRule = "2A"
```
