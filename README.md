# razor-1/CLDR

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/razor-1/cldr/Build?style=flat-square)](https://github.com/razor-1/cldr/actions?query=workflow%3ABuild)
[![Go Report Card](https://goreportcard.com/badge/github.com/razor-1/cldr?style=flat-square)](https://goreportcard.com/report/github.com/razor-1/cldr)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.15-61CFDD.svg?style=flat-square)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/razor-1/cldr)](https://pkg.go.dev/github.com/razor-1/cldr)


This started as a fork of https://github.com/theplant/cldr. It incorporates a few significant changes:
* Merges data from parent locales in the CLDR. This is not a complete implementation of all rules in 
http://unicode.org/reports/tr35/#Common_Elements, but it's substantially better than not doing this at all
* Supports formatting timezone components using the GMT offset format
* Uses the plural code generation logic from https://github.com/nicksnyder/go-i18n
* The resources package is refactored to allow easy dynamic loading of locales at runtime

You're welcome to use this package on its own, but it might be easier to use https://github.com/razor-1/localizer. It
will allow you to get a locale, populated with this CLDR data, load translations into it, and use it for all your
localization needs. 
