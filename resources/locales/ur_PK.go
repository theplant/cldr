// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ur_PK() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "d MMM، y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT {0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "جنوری", Feb: "فروری", Mar: "مارچ", Apr: "اپریل", May: "مئی", Jun: "جون", Jul: "جولائی", Aug: "اگست", Sep: "ستمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "جنوری", Feb: "فروری", Mar: "مارچ", Apr: "اپریل", May: "مئی", Jun: "جون", Jul: "جولائی", Aug: "اگست", Sep: "ستمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "اتوار", Mon: "پیر", Tue: "منگل", Wed: "بدھ", Thu: "جمعرات", Fri: "جمعہ", Sat: "ہفتہ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "اتوار", Mon: "پیر", Tue: "منگل", Wed: "بدھ", Thu: "جمعرات", Fri: "جمعہ", Sat: "ہفتہ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "اتوار", Mon: "پیر", Tue: "منگل", Wed: "بدھ", Thu: "جمعرات", Fri: "جمعہ", Sat: "ہفتہ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ur]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "\u200e-\u200e", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "متحدہ عرب اماراتی درہم", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "افغان افغانی", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "البانیا کا لیک", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "آرمینیائی ڈرم", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "نیدر لینڈز انٹیلیئن گلڈر", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "انگولا کا کوانزا", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "ارجنٹائن پیسہ", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "آسٹریلین ڈالر", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "اروبن فلورِن", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "آذربائجانی منات", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "بوسنیا ہرزیگووینا کا قابل منتقلی نشان", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "باربیڈین ڈالر", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "بنگلہ دیشی ٹکا", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "بلغارین لیو", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "بحرینی دینار", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "برونڈیئن فرانک", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "برموڈا ڈالر", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "برونئی ڈالر", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "بولیوین بولیویانو", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "برازیلی ریئل", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "بہامانی ڈالر", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "بھوٹانی گُلٹرم", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "بوتسوانا کا پولا", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "بیلاروسی روبل", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "بیلاروسی روبل (۲۰۰۰–۲۰۱۶)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "بیلیز ڈالر", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "کنیڈین ڈالر", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "کانگولیز فرانک", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "سوئس فرانکس", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "چلّین پیسہ", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "چینی یوآن (آف شور)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "چینی یوآن", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "کولمبین پیسہ", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "کوسٹا ریکا کا کولن", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "کیوبا کا قابل منتقلی پیسو", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "کیوبا کا پیسو", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "کیپ ورڈی کا اسکیوڈو", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "چیک کرونا", Symbol: "CZK"},
				currency.DEM: cldr.Currency{DisplayName: "ڈچ مارکس", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "جبوتی فرانک", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "ڈنمارک کرون", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "ڈومنیکن پیسو", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "الجیریائی دینار", Symbol: "DZD"},
				currency.EEK: cldr.Currency{DisplayName: "ایسٹونین کرون", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "مصری پاؤنڈ", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "اریٹیریا کا نافکا", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ایتھوپیائی بِرّ", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "یورو", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "فجی کا ڈالر", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "فاکلینڈ آئلینڈز پونڈ", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "فرانسیسی فرانک", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "برطانوی پاؤنڈ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "جارجیائی لاری", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "گھانا کا سیڈی", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "جبل الطارق پونڈ", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "گامبیا کا ڈلاسی", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "گنی فرانک", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "گواٹے مالا کا کوئٹزل", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "گویانیز ڈالر", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "ھانگ کانگ ڈالر", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "ہونڈوران لیمپیرا", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "کروشین کونا", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "ہیتی کا گؤرڈی", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "ہنگرین فورنٹ", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "انڈونیشین روپیہ", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "اسرائیلی نیا شیکل", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "بھارتی روپیہ", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "عراقی دینار", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ایرانی ریال", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "آئس لينڈی کرونا", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "جمائیکن ڈالر", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "اردنی دینار", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "جاپانی ین", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "کینیائی شلنگ", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "کرغستانی سوم", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "کمبوڈیائی ریئل", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "کوموریئن فرانک", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "شمالی کوریائی وون", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "جنوبی کوریائی وون", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "کویتی دینار", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "کیمین آئلینڈز ڈالر", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "قزاخستانی ٹینگ", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "لاؤشیائی کِپ", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "لبنانی پونڈ", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "سری لنکائی روپیہ", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "لائبریائی ڈالر", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "لیسوتھو لوٹی", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "لیتھوینیائی لیٹاس", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "لاتویائی لیٹس", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "لیبیائی دینار", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "مراکشی درہم", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "مالدووی لیو", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ملاگاسی اریاری", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "مقدونیائی دینار", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "میانمار کیاٹ", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "منگولیائی ٹگرِ", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "میکانیز پٹاکا", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "موریطانیائی اوگوئیا (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "موریطانیائی اوگوئیا", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "ماریشس کا روپیہ", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "مالدیپ کا روفیہ", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "ملاوی کواچا", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "میکسیکی پیسہ", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "ملیشیائی رنگِٹ", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "موزامبیقی میٹیکل", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "نامیبیائی ڈالر", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "نائیجیریائی نائرا", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "نکارا گوا کا کورڈوبا", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "ناروے کرون", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "نیپالی روپیہ", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "نیوزی لینڈ ڈالر", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "عمانی ریال", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "پنامہ کا بالبوآ", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "پیروویئن سول", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "پاپوآ نیو گنی کا کینا", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "فلپائینی پیسہ", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "پاکستانی روپیہ", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "پولش زلوٹی", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "پیراگوئے کا گوآرنی", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "قطری ریال", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "رومانیائی لیو", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "سربین دینار", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "روسی روبل", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "روانڈا کا فرانک", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "سعودی ریال", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "سولومن آئلینڈز ڈالر", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "سشلی کا روپیہ", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "سوڈانی پاؤنڈ", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "سویڈن کرونا", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "سنگا پور ڈالر", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "سینٹ ہیلینا پاؤنڈ", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "سلوانین ٹولر", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "سلووک کرونا", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "سیئرا لیون لیون", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "صومالی شلنگ", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "سورینامی ڈالر", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "جنوبی سوڈانی پاؤنڈ", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "ساؤ ٹوم اور پرنسپے ڈوبرا (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "ساؤ ٹومے اور پرنسپے ڈوبرا", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "شامی پونڈ", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "سوازی لیلانجینی", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "تھائی باہت", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "تاجکستانی سومونی", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "ترکمانستانی منات", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "تیونیسیائی دینار", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "ٹونگن پانگا", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "ترکی لیرا", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ترینیداد اور ٹوباگو کا ڈالر", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "نیو تائیوان ڈالر", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "تنزانیائی شلنگ", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "یوکرینیائی ہریونیا", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "یوگانڈا شلنگ", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "امریکی ڈالر", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "یوروگویان پیسو", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "ازبکستانی سوم", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "وینزویلا بولیور (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "وینزویلا بولیور (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "وینزویلا بولیور", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "ویتنامی ڈانگ", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "وینوواتو واتو", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "ساموآ کا ٹالا", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "وسطی افریقی [CFA] فرانک", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "مشرقی کریبیا کا ڈالر", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "مغربی افریقی [CFA] فرانک", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP فرانک", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "نامعلوم کرنسی", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "یمنی ریال", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "جنوبی افریقی رانڈ", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "زامبیائی کواچا (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "زامبیائی کواچا", Symbol: "ZMW"},
			},
		},
	}
}
