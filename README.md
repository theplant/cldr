# razor-1/CLDR

This started as a fork of https://github.com/theplant/cldr. It incorporates a few significant changes:
* Merges data from parent locales in the CLDR. This is not a complete implementation of all rules in 
http://unicode.org/reports/tr35/#Common_Elements, but it's substantially better than not doing this at all
* Supports formatting timezone components using the GMT offset format
* Uses the plural code generation logic from https://github.com/nicksnyder/go-i18n
* The resources package is refactored to allow easy dynamic loading of locales at runtime

You're welcome to use this package on its own, but it might be easier to use https://github.com/razor-1/localizer. It
will allow you to get a locale, populated with this CLDR data, load translations into it, and use it for all your
localization needs. 
