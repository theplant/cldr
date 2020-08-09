// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_kn() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "kn",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "hh:mm:ss a zzzz", Long: "hh:mm:ss a z", Medium: "hh:mm:ss a", Short: "hh:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ಜನ", Feb: "ಫೆಬ್ರ", Mar: "ಮಾರ್ಚ್", Apr: "ಏಪ್ರಿ", May: "ಮೇ", Jun: "ಜೂನ್", Jul: "ಜುಲೈ", Aug: "ಆಗ", Sep: "ಸೆಪ್ಟೆಂ", Oct: "ಅಕ್ಟೋ", Nov: "ನವೆಂ", Dec: "ಡಿಸೆಂ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ಜ", Feb: "ಫೆ", Mar: "ಮಾ", Apr: "ಏ", May: "ಮೇ", Jun: "ಜೂ", Jul: "ಜು", Aug: "ಆ", Sep: "ಸೆ", Oct: "ಅ", Nov: "ನ", Dec: "ಡಿ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ಜನವರಿ", Feb: "ಫೆಬ್ರವರಿ", Mar: "ಮಾರ್ಚ್", Apr: "ಏಪ್ರಿಲ್", May: "ಮೇ", Jun: "ಜೂನ್", Jul: "ಜುಲೈ", Aug: "ಆಗಸ್ಟ್", Sep: "ಸೆಪ್ಟೆಂಬರ್", Oct: "ಅಕ್ಟೋಬರ್", Nov: "ನವೆಂಬರ್", Dec: "ಡಿಸೆಂಬರ್"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ಭಾನು", Mon: "ಸೋಮ", Tue: "ಮಂಗಳ", Wed: "ಬುಧ", Thu: "ಗುರು", Fri: "ಶುಕ್ರ", Sat: "ಶನಿ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "ಭಾ", Mon: "ಸೋ", Tue: "ಮಂ", Wed: "ಬು", Thu: "ಗು", Fri: "ಶು", Sat: "ಶ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ಭಾನು", Mon: "ಸೋಮ", Tue: "ಮಂಗಳ", Wed: "ಬುಧ", Thu: "ಗುರು", Fri: "ಶುಕ್ರ", Sat: "ಶನಿ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "ಭಾನುವಾರ", Mon: "ಸೋಮವಾರ", Tue: "ಮಂಗಳವಾರ", Wed: "ಬುಧವಾರ", Thu: "ಗುರುವಾರ", Fri: "ಶುಕ್ರವಾರ", Sat: "ಶನಿವಾರ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ಪೂರ್ವಾಹ್ನ", PM: "ಅಪರಾಹ್ನ"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ಪೂರ್ವಾಹ್ನ", PM: "ಅಪರಾಹ್ನ"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ಪೂರ್ವಾಹ್ನ", PM: "ಅಪರಾಹ್ನ"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_kn]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "ಸಂಯುಕ್ತ ಅರಬ್\u200c ಎಮಿರೇಟ್\u200c\u200cಗಳ ದಿರಾಮ್\u200c\u200c", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "ಅಫ್\u200cಘನ್ ಅಫಘಾನಿ", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "ಅಲ್\u200cಬೇನಿಯನ್ ಲೆಕ್", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "ಅರ್ಮೆನಿಯನ್ ಡ್ರಾಮ್", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "ನೆದರ್ಲೆಂಡ್ಸ್ ಆಂಟಿಲಿಯನ್ ಗಿಲ್ಡರ್", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "ಅಂಗೋಲಾದ ಕ್ವಾನ್ಝಾ", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "ಅರ್ಜೆಂಟಿನಾ ಪೆಸೊ", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "ಆಸ್ಟ್ರೇಲಿಯನ್ ಡಾಲರ್\u200c", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "ಅರುಬನ್ ಫ್ಲೊರೀನ್\u200c\u200c", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "ಅಝರ್\u200cಬೈಜಾನಿ ಮನಾತ್", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "ಬೋಸ್ನಿಯಾ-ಹರ್ಜ್\u200cಗೋವಿನ ಪರಿವರ್ತನೀಯ ಗುರುತು", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "ಬರ್ಬಾಡಿಯನ್ ಡಾಲರ್", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "ಬಾಂಗ್ಲಾದೇಶದ ಟಾಕಾ", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "ಬಲ್ಗೇರಿಯನ್ ಲೆವ್", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "ಬಹ್\u200c\u200cರೈನಿ ದಿನಾರ್", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "ಬುರುಂದಿಯನ್ ಫ್ರಾಂಕ್", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "ಬರ್ಮುಡನ್ ಡಾಲರ್", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "ಬ್ರೂನಿ ಡಾಲರ್", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "ಬೊಲಿವಿಯಾದ ಬೊಲಿವಿಯಾನೊ", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "ಬ್ರೆಜಿಲಿಯನ್\u200c ರಿಯಲ್", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "ಬಹಾಮಿಯನ್ ಡಾಲರ್", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ಭೂತಾನೀಸ್ ನುಲ್ತರಮ್", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "ಬೋಟ್ಸ್\u200cವಾನನ್ ಪುಲಾ", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "ಬೆಲಾರುಸಿಯನ್ ರೂಬಲ್", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "ಬೆಲಾರುಸಿಯನ್ ರೂಬಲ್ (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "ಬೆಲೀಜ್ ಡಾಲರ್", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "ಕೆನಡಾದ ಡಾಲರ್", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "ಕಾಂಗೋಲೀಸ್ ಫ್ರಾಂಕ್", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "ಸ್ವಿಸ್ ಫ್ರಾಂಕ್", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "ಚಿಲಿಯ ಪೆಸೊ", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "ಚೈನೀಸ್ ಯುವಾನ್ (ಆಫ್\u200cಶೋರ್)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "ಚೈನೀಸ್ ಯುವಾನ್", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "ಕೊಲೊಂಬಿಯೋದ ಪೆಸೊ", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "ಕೋಸ್ಟ ರಿಕನ್ ಕೊಲನ್", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "ಕ್ಯುಬಾದ ಪರಿವರ್ತನೀಯ ಪೆಸೊ", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "ಕ್ಯೂಬಾದ ಪೆಸೊ", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "ಕೇಪ್ ವರ್ಡಿನ್ ಎಸ್\u200cಕೂಡೊ", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "ಝೆಕ್ ಗಣರಾಜ್ಯ ಕೊರೂನ", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "ಜಿಬೊಟಿಯನ್ ಫ್ರಾಂಕ್", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "ಡ್ಯಾನಿಶ್ ಕ್ರೋನ್", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "ಡೊಮಿನಿಕನ್ ಪೆಸೊ", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "ಅಲ್ಜೀರಿಯನ್ ದಿನಾರ್", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "ಈಜಿಪ್ಷಿಯನ್ ಪೌಂಡ್\u200d", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "ಎರಿತ್ರಿಯನ್ ನಕ್ಫಾ", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ಇಥಿಯೋಪಿಯನ್ ಬಿರ್", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "ಯೂರೊ", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "ಫಿಜಿಯನ್ ಡಾಲರ್", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "ಫಾಲ್ಕ್\u200cಲ್ಯಾಂಡ್ ದ್ವೀಪಗಳ ಪೌಂಡ್", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "ಬ್ರಿಟೀಷ್ ಪೌಂಡ್", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "ಜಾರ್ಜಿಯಾದ ಲಾರಿ", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "ಘಾನಾದ ಸೆದಿ", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "ಗಿಬ್ರಾಲ್ಟರ್ ಪೌಂಡ್", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "ಗ್ಯಾಂಬಿಯಾದ ದಲಾಸಿ", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "ಗಿನಿಯನ್ ಫ್ರಾಂಕ್", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "ಗ್ವಾಟೆಮಾಲಾದ ಕುಯಿಟ್ಸಲ್\u200c\u200c", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "ಗುಯಾನೀಸ್\u200c ಡಾಲರ್\u200c", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "ಹಾಂಗ್ ಕಾಂಗ್ ಡಾಲರ್", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "ಹೊಂಡುರಾನ್\u200c ಲೆಂಪಿರಾ", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "ಕ್ರೊಯೆಷ್ಯಾದ ಕೂನಾ", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "ಹೈಟಿಯ ಗೋರ್ದೆ", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "ಹಂಗೇರಿಯನ್ ಫೋರಿಂಟ್", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "ಇಂಡೊನೇಷ್ಯಾ ರುಪೈ", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "ಇಸ್ರೇಲಿ ನ್ಯೂ ಶೇಖಲ್", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "ಭಾರತೀಯ ರೂಪಾಯಿ", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "ಇರಾಕಿ ದಿನಾರ್", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ಇರಾನಿಯನ್ ರಿಯಲ್", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "ಐಸ್\u200cಲ್ಯಾಂಡಿಕ್ ಕ್ರೋನಾ", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "ಜಮೈಕನ್ ಡಾಲರ್", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "ಜೋರ್ಡಾನಿಯನ್ ದಿನಾರ್", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "ಜಪಾನೀಸ್ ಯೆನ್", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "ಕೀನ್ಯಾದ ಶಿಲ್ಲಿಂಗ್\u200c", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "ಕಿರ್ಗಿಸ್ತಾನಿ ಸೋಮ್", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "ಕಾಂಬೋಡಿಯನ್ ರಿಯಲ್", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "ಕೊಮೊರಿಯನ್ ಫ್ರಾಂಕ್", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "ಉತ್ತರ ಕೊರಿಯನ್ ವೋನ್", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "ದಕ್ಷಿಣ ಕೊರಿಯನ್ ವೊನ್", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "ಕುವೈತೀ ದಿನಾರ್", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "ಕೆಮ್ಯಾನ್\u200c ಐಲ್ಯಾಂಡ್\u200cನ ಡಾಲರ್\u200c", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "ಕಜಾಕಿಸ್ತಾನಿ ತೆಂಗೆ", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "ಲೋಟಿಯನ್ ಕಿಪ್", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "ಲೆಬೆನೀಸ್ ಪೌಂಡ್", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "ಶ್ರೀಲಂಕಾದ ರುಪೀ", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "ಲಿಬೇರಿಯನ್ ಡಾಲರ್", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "ಲೆಸೊತೊ ಲೊತಿ", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "ಲಿಥುನಿಯನ್\u200c ಲಿತಾಸ್\u200c", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "ಲ್ಯಾಟ್ವಿಯನ್ ಲ್ಯಾಟ್ಸ್", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "ಲಿಬಿಯಾದ ದಿನಾರ್\u200c", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "ಮೊರೊಕನ್ ದಿರ್\u200cಹಮ್", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "ಮಲ್ದೋವಾದ ಲೆವೂ", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ಮಲಗಾಸಿ ಅರಿಯಾರಿ", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "ಮೆಸಡೋನಿಯನ್ ದಿನಾರ್", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "ಮ್ಯಾನ್ಮಾರ್ ಕ್ಯಾಟ್", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "ಮಂಗೋಲಿಯಾದ ತುಗ್ರಿಕ್\u200c\u200c", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "ಮಕಾನಿಸ್ ಪಟಾಕಾ", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "ಮೌರೀಶಿಯನಿಯನ್ ಒಗಿಯ (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "ಮೌರೀಶಿಯನಿಯನ್ ಒಗಿಯ", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "ಮಾರಿಷಿಯನ್ ರುಪಿ", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "ಮಾಲ್ಡೀವಿಯನ್ ರುಫಿಯಾ", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "ಮಲಾವಿಯ ಕ್ವಾಚ", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "ಮೆಕ್ಸಿಕೊದ ಪೆಸೊ", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "ಮಲೇಶಿಯನ್ ರಿಂಗಿಟ್", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "ಮೊಝಾಂಬಿಕನ್ ಮೆಟಿಕಲ್", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "ನಮೀಬಿಯನ್ ಡಾಲರ್", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "ನೈಜೀರಿಯಾದ ನೇರಾ", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "ನಿಕಾರಗ್ವಾದ ಕರ್ದೊಬಾ", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "ನಾರ್ವೇಯ ಕ್ರೋನ್", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "ನೇಪಾಳದ ರುಪೀ", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "ನ್ಯೂಜಿಲ್ಯಾಂಡ್ ಡಾಲರ್", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ಒಮಾನಿ ರಿಯಲ್", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "ಪನಾಮಾನಿಯನ್ ಬಲ್ಬೋವಾ", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "ಪೆರುವಿಯನ್ ಸೊಲ್", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "ಪಪುವಾ ನ್ಯೂ ಗಿನಿಯನ್ ಕಿನಾ", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "ಫಿಲಿಪ್ಪೈನ್ ಪಿಸೊ", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "ಪಾಕಿಸ್ತಾನದ ರುಪೀ", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "ಪೊಲಿಶ್ ಝ್ಲೋಟಿ", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "ಪೆರುಗ್ವೇಯ ಗ್ವಾರನೀ", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "ಖತಾರಿ ರಿಯಲ್", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "ರೊಮೇನಿಯನ್ ಲೆವು", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "ಸೆರ್ಬಿಯನ್ ದಿನಾರ್", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "ರಶಿಯನ್ ರೂಬಲ್", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ರುವಾಂಡನ್ ಫ್ರಾಂಕ್", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "ಸೌದಿ ರಿಯಾಲ್", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "ಸೊಲೊಮನ್ ದ್ವೀಪಗಳ ಡಾಲರ್", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "ಸೆಚೊಲಿಯೊಸ್ ರುಪಿ", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "ಸುಡಾನೀಸ್ ಪೌಂಡ್", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "ಸ್ವೀಡಿಷ್ ಕ್ರೋನಾ", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "ಸಿಂಗಾಪುರ್ ಡಾಲರ್\u200c", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "ಸೇಂಟ್ ಹೆಲೇನಾ ಪೌಂಡ್", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "ಸಿಯೆರಾ ಲಿಯೋನಿಯನ್ ಲಿಯೋನ್", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "ಸೊಮಾಲಿ ಶಿಲ್ಲಿಂಗ್", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "ಸುರಿನಾಮೀಸ್ ಡಾಲರ್", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "ದಕ್ಷಿಣ ಸೂಡಾನೀಸ್ ಪೌಂಡ್\u200d", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "ಸಾವೊ ಟೋಮ್ ಮತ್ತು ಪ್ರಿನ್ಸಿಪ್ ದೊಬ್ರಾ (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "ಸಾವೊ ಟೋಮ್ ಮತ್ತು ಪ್ರಿನ್ಸಿಪ್ ದೊಬ್ರಾ", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "ಸಿರಿಯನ್ ಪೌಂಡ್", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "ಸ್ವಾಜಿ ಲಿಲಂಗೆನಿ", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "ಥಾಯ್ ಬಹ್ತ್", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "ತಜಕಿಸ್ತಾನಿ ಸೊಮೋನಿ", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "ಟರ್ಕ್\u200dಮೆನಿಸ್ತಾನ್ ಮನಾತ್", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "ಟ್ಯುನೀಷಿಯನ್\u200c ದಿನಾರ್", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "ಟೊಂಗಾ ಪಾಂಗ", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "ಟರ್ಕಿಶ್ ಲಿರಾ", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ಟ್ರಿನಿಡಾಡ್ ಮತ್ತು ಟೊಬಾಗೊ ಡಾಲರ್", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "ನ್ಯೂ ತೈವಾನ್ ಡಾಲರ್", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "ತಾನ್\u200cಜೇನಿಯನ್ ಶಿಲ್ಲಿಂಗ್", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ಉಕ್ರೇನಿಯನ್ ಹ್ರಿವ್ನೀಯ", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "ಉಗಾಂಡನ್ ಶಿಲ್ಲಿಂಗ್", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "ಅಮೆರಿಕದ ಡಾಲರ್\u200c", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "ಉರುಗ್ವೆಯ ಪೆಸೊ", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "ಉಜ್ಬೇಕಿಸ್ತಾನ್ ಸೊಮ್", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "ವೆನಿಜುಲಿಯನ್ ಬೊಲಿವರ್ (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "ವೆನಿಜುಲಿಯನ್ ಬೊಲಿವರ್", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "ವಿಯೆಟ್ನಾಮೀಸ್ ಡಾಂಗ್", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "ವನೂತು ವತು", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "ಸಮೋನ್ ತಲಾ", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "ಮಧ್ಯ ಆಫ್ರಿಕನ್ CFA ಫ್ರಾಂಕ್", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "ಪೂರ್ವ ಕೆರೀಬಿಯನ್ ಡಾಲರ್", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "ಪಶ್ಚಿಮ ಆಫ್ರಿಕಾದ CFA ಫ್ರಾಂಕ್", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "[CFP] ಫ್ರಾಂಕ್", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "ಅಪರಿಚಿತ ಕರೆನ್ಸಿ", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "ಯೆಮೆನಿ ರಿಯಲ್", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "ದಕ್ಷಿಣ ಆಫ್ರಿಕನ್ ರಾಂಡ್", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "ಜಾಂಬಿಯಾ ಕ್ವಾಚ (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "ಜಾಂಬಿಯಾ ಕ್ವಾಚ", Symbol: "ZMW"},
			},
		},
	}
}
