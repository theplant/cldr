// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_ur_IN() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ur_IN",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "d MMM، y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT {0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "جنوری", Feb: "فروری", Mar: "مارچ", Apr: "اپریل", May: "مئی", Jun: "جون", Jul: "جولائی", Aug: "اگست", Sep: "ستمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "جنوری", Feb: "فروری", Mar: "مارچ", Apr: "اپریل", May: "مئی", Jun: "جون", Jul: "جولائی", Aug: "اگست", Sep: "ستمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "اتوار", Mon: "پیر", Tue: "منگل", Wed: "بدھ", Thu: "جمعرات", Fri: "جمعہ", Sat: "ہفتہ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "اتوار", Mon: "پیر", Tue: "منگل", Wed: "بدھ", Thu: "جمعرات", Fri: "جمعہ", Sat: "ہفتہ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "اتوار", Mon: "پیر", Tue: "منگل", Wed: "بدھ", Thu: "جمعرات", Fri: "جمعہ", Sat: "ہفتہ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "\u200e-\u200e", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##,##0.00", CurrencyAccounting: "", Percent: "#,##0%"},
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
				currency.CRC: cldr.Currency{DisplayName: "کوسٹا ریکا کولون", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "قابل منتقلی کیوبائی پیسو", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "کیوبائی پیسو", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "کیپ ورڈی اسکیوڈو", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "چیک کرونا", Symbol: "CZK"},
				currency.DEM: cldr.Currency{DisplayName: "ڈچ مارکس", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "جبوتی فرانک", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "ڈنمارک کرون", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "ڈومنیکن پیسو", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "الجیریائی دینار", Symbol: "DZD"},
				currency.EEK: cldr.Currency{DisplayName: "ایسٹونین کرون", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "مصری پاؤنڈ", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "اریٹیریائی ناکفا", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ایتھوپیائی بِرّ", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "یورو", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "فجی کا ڈالر", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "فاکلینڈ آئلینڈز پونڈ", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "فرانسیسی فرانک", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "برطانوی پاونڈ سٹرلنگ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "جارجیائی لاری", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "گھانی سیڈی", Symbol: "GHS"},
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
				currency.PKR: cldr.Currency{DisplayName: "پاکستانی روپیہ", Symbol: "PKR"},
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
				currency.WST: cldr.Currency{DisplayName: "ساموآئی ٹالا", Symbol: "WST"},
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
		Languages: cldr.Languages{
			language.AA:      "افار",
			language.AB:      "ابقازیان",
			language.ACE:     "اچائینیز",
			language.ACH:     "اکولی",
			language.ADA:     "ادانگمے",
			language.ADY:     "ادیگھے",
			language.AF:      "افریقی",
			language.AGQ:     "اغم",
			language.AIN:     "اینو",
			language.AK:      "اکان",
			language.ALE:     "الیوت",
			language.ALT:     "جنوبی الٹائی",
			language.AM:      "امہاری",
			language.AN:      "اراگونیز",
			language.ANP:     "انگیکا",
			language.AR:      "عربی",
			language.AR_001:  "جدید معیاری عربی",
			language.ARN:     "ماپوچے",
			language.ARP:     "اراپاہو",
			language.AS:      "آسامی",
			language.ASA:     "آسو",
			language.AST:     "اسٹوریائی",
			language.AV:      "اواری",
			language.AWA:     "اودھی",
			language.AY:      "ایمارا",
			language.AZ:      "ازیری",
			language.AZ_ARAB: "آزربائیجانی (عربی)",
			language.BA:      "باشکیر",
			language.BAN:     "بالینیز",
			language.BAS:     "باسا",
			language.BE:      "بیلاروسی",
			language.BEM:     "بیمبا",
			language.BEZ:     "بینا",
			language.BG:      "بلغاری",
			language.BGN:     "مغربی بلوچی",
			language.BHO:     "بھوجپوری",
			language.BI:      "بسلاما",
			language.BIN:     "بینی",
			language.BLA:     "سکسیکا",
			language.BM:      "بمبارا",
			language.BN:      "بنگلہ",
			language.BO:      "تبتی",
			language.BR:      "بریٹن",
			language.BRX:     "بوڈو",
			language.BS:      "بوسنیائی",
			language.BUG:     "بگینیز",
			language.BYN:     "بلین",
			language.CA:      "کیٹالان",
			language.CCP:     "چکمہ",
			language.CE:      "چیچن",
			language.CEB:     "سیبوآنو",
			language.CGG:     "چیگا",
			language.CH:      "چیمارو",
			language.CHK:     "چوکیز",
			language.CHM:     "ماری",
			language.CHO:     "چاکٹاؤ",
			language.CHR:     "چیروکی",
			language.CHY:     "چینّے",
			language.CKB:     "سورانی کردی",
			language.CO:      "کوراسیکن",
			language.CRS:     "سیسلوا کریولے فرانسیسی",
			language.CS:      "چیک",
			language.CU:      "چرچ سلاوک",
			language.CV:      "چوواش",
			language.CY:      "ویلش",
			language.DA:      "ڈینش",
			language.DAK:     "ڈاکوٹا",
			language.DAR:     "درگوا",
			language.DAV:     "تائتا",
			language.DE:      "جرمن",
			language.DE_AT:   "آسٹریائی جرمن",
			language.DE_CH:   "سوئس ہائی جرمن",
			language.DGR:     "دوگریب",
			language.DJE:     "زرمہ",
			language.DSB:     "ذیلی سربیائی",
			language.DUA:     "دوالا",
			language.DV:      "ڈیویہی",
			language.DYO:     "جولا فونيا",
			language.DZ:      "ژونگکھا",
			language.DZG:     "دزاگا",
			language.EBU:     "امبو",
			language.EE:      "ایو",
			language.EFI:     "ایفِک",
			language.EKA:     "ایکاجوی",
			language.EL:      "یونانی",
			language.EN:      "انگریزی",
			language.EN_AU:   "آسٹریلیائی انگریزی",
			language.EN_CA:   "کینیڈین انگریزی",
			language.EN_GB:   "برطانوی انگریزی",
			language.EN_US:   "امریکی انگریزی",
			language.EO:      "ایسپرانٹو",
			language.ES:      "ہسپانوی",
			language.ES_419:  "لاطینی امریکی ہسپانوی",
			language.ES_ES:   "یورپی ہسپانوی",
			language.ES_MX:   "میکسیکن ہسپانوی",
			language.ET:      "اسٹونین",
			language.EU:      "باسکی",
			language.EWO:     "ایوانڈو",
			language.FA:      "فارسی",
			language.FA_AF:   "دری",
			language.FF:      "فولہ",
			language.FI:      "فینیش",
			language.FIL:     "فلیپینو",
			language.FJ:      "فجی",
			language.FO:      "فیروئیز",
			language.FON:     "فون",
			language.FR:      "فرانسیسی",
			language.FR_CA:   "کینیڈین فرانسیسی",
			language.FR_CH:   "سوئس فرینچ",
			language.FRC:     "کاجن فرانسیسی",
			language.FUR:     "فریولیائی",
			language.FY:      "مغربی فریسیئن",
			language.GA:      "آئیرِش",
			language.GAA:     "گا",
			language.GAG:     "غاغاوز",
			language.GAN:     "gan",
			language.GD:      "سکاٹش گیلک",
			language.GEZ:     "گیز",
			language.GIL:     "گلبرتیز",
			language.GL:      "گالیشیائی",
			language.GN:      "گُارانی",
			language.GOR:     "گورانٹالو",
			language.GSW:     "سوئس جرمن",
			language.GU:      "گجراتی",
			language.GUZ:     "گسی",
			language.GV:      "مینکس",
			language.GWI:     "گوئچ ان",
			language.HA:      "ہؤسا",
			language.HAW:     "ہوائی",
			language.HE:      "عبرانی",
			language.HI:      "ہندی",
			language.HIL:     "ہالیگینون",
			language.HMN:     "ہمانگ",
			language.HR:      "کروشین",
			language.HSB:     "اپر سربیائی",
			language.HT:      "ہیتی",
			language.HU:      "ہنگیرین",
			language.HUP:     "ہیوپا",
			language.HY:      "آرمینیائی",
			language.HZ:      "ہریرو",
			language.IA:      "بین لسانیات",
			language.IBA:     "ایبان",
			language.IBB:     "ابی بیو",
			language.ID:      "انڈونیثیائی",
			language.IG:      "اِگبو",
			language.II:      "سچوان ای",
			language.ILO:     "ایلوکو",
			language.INH:     "انگوش",
			language.IO:      "ایڈو",
			language.IS:      "آئس لینڈک",
			language.IT:      "اطالوی",
			language.IU:      "اینُکٹیٹٹ",
			language.JA:      "جاپانی",
			language.JBO:     "لوجبان",
			language.JGO:     "نگومبا",
			language.JMC:     "ماشیم",
			language.JV:      "جاوانیز",
			language.KA:      "جارجيائى",
			language.KAB:     "قبائلی",
			language.KAC:     "کاچن",
			language.KAJ:     "جے جو",
			language.KAM:     "کامبا",
			language.KBD:     "کبارڈین",
			language.KCG:     "تیاپ",
			language.KDE:     "ماکونده",
			language.KEA:     "کابويرديانو",
			language.KFO:     "کورو",
			language.KG:      "کانگو",
			language.KHA:     "کھاسی",
			language.KHQ:     "کويرا شيني",
			language.KI:      "کیکویو",
			language.KJ:      "کونیاما",
			language.KK:      "قزاخ",
			language.KKJ:     "کاکو",
			language.KL:      "کلالیسٹ",
			language.KLN:     "کالينجين",
			language.KM:      "خمیر",
			language.KMB:     "کیمبونڈو",
			language.KN:      "کنڑ",
			language.KO:      "کوریائی",
			language.KOI:     "کومی پرمیاک",
			language.KOK:     "کونکنی",
			language.KPE:     "کیپیلّے",
			language.KR:      "کنوری",
			language.KRC:     "کراچے بالکر",
			language.KRL:     "کیرلین",
			language.KRU:     "کوروکھ",
			language.KS:      "کشمیری",
			language.KSB:     "شامبالا",
			language.KSF:     "بافيا",
			language.KSH:     "کولوگنیائی",
			language.KU:      "کرد",
			language.KUM:     "کومیک",
			language.KV:      "کومی",
			language.KW:      "کورنش",
			language.KY:      "کرغیزی",
			language.LA:      "لاطینی",
			language.LAD:     "لیڈینو",
			language.LAG:     "لانگی",
			language.LB:      "لکسمبرگیش",
			language.LEZ:     "لیزگیان",
			language.LG:      "گینڈا",
			language.LI:      "لیمبرگش",
			language.LKT:     "لاکوٹا",
			language.LN:      "لِنگَلا",
			language.LO:      "لاؤ",
			language.LOU:     "لوزیانا کریول",
			language.LOZ:     "لوزی",
			language.LRC:     "شمالی لری",
			language.LT:      "لیتھوینین",
			language.LU:      "لبا-کاتانجا",
			language.LUA:     "لیوبا لولوآ",
			language.LUN:     "لونڈا",
			language.LUO:     "لو",
			language.LUS:     "میزو",
			language.LUY:     "لویا",
			language.LV:      "لیٹوین",
			language.MAD:     "مدورسی",
			language.MAG:     "مگہی",
			language.MAI:     "میتھیلی",
			language.MAK:     "مکاسر",
			language.MAS:     "مسائی",
			language.MDF:     "موکشا",
			language.MEN:     "میندے",
			language.MER:     "میرو",
			language.MFE:     "موریسیین",
			language.MG:      "ملاگاسی",
			language.MGH:     "ماخاوا-ميتو",
			language.MGO:     "میٹا",
			language.MH:      "مارشلیز",
			language.MI:      "ماؤری",
			language.MIC:     "مکمیک",
			language.MIN:     "منانگکباؤ",
			language.MK:      "مقدونیائی",
			language.ML:      "مالایالم",
			language.MN:      "منگولین",
			language.MNI:     "منی پوری",
			language.MOH:     "موہاک",
			language.MOS:     "موسی",
			language.MR:      "مراٹهی",
			language.MS:      "مالے",
			language.MT:      "مالٹی",
			language.MUA:     "منڈانگ",
			language.MUL:     "متعدد زبانیں",
			language.MUS:     "کریک",
			language.MWL:     "میرانڈیز",
			language.MY:      "برمی",
			language.MYV:     "ارزیا",
			language.MZN:     "مزندرانی",
			language.NA:      "ناؤرو",
			language.NAP:     "نیاپولیٹن",
			language.NAQ:     "ناما",
			language.NB:      "نارویجین بوکمل",
			language.ND:      "شمالی دبیل",
			language.NDS:     "ادنی جرمن",
			language.NDS_NL:  "ادنی سیکسن",
			language.NE:      "نیپالی",
			language.NEW:     "نیواری",
			language.NG:      "نڈونگا",
			language.NIA:     "نیاس",
			language.NIU:     "نیویائی",
			language.NL:      "ڈچ",
			language.NL_BE:   "فلیمِش",
			language.NMG:     "کوايسو",
			language.NN:      "نارویجین نینورسک",
			language.NNH:     "نگیمبون",
			language.NO:      "نارویجین",
			language.NOG:     "نوگائی",
			language.NQO:     "اینکو",
			language.NR:      "جنوبی نڈیبیلی",
			language.NSO:     "شمالی سوتھو",
			language.NUS:     "نویر",
			language.NV:      "نواجو",
			language.NY:      "نیانجا",
			language.NYN:     "نینکول",
			language.OC:      "آکسیٹان",
			language.OM:      "اورومو",
			language.OR:      "اڑیہ",
			language.OS:      "اوسیٹک",
			language.PA:      "پنجابی",
			language.PAG:     "پنگاسنان",
			language.PAM:     "پامپنگا",
			language.PAP:     "پاپیامینٹو",
			language.PAU:     "پالاون",
			language.PCM:     "نائجیریائی پڈگن",
			language.PL:      "پولش",
			language.PRG:     "پارسی",
			language.PS:      "پشتو",
			language.PT:      "پُرتگالی",
			language.PT_BR:   "برازیلی پرتگالی",
			language.PT_PT:   "یورپی پرتگالی",
			language.QU:      "کویچوآ",
			language.QUC:     "کيشی",
			language.RAP:     "رپانوی",
			language.RAR:     "راروتونگان",
			language.RM:      "رومانش",
			language.RN:      "رونڈی",
			language.RO:      "رومینین",
			language.RO_MD:   "مالدووا",
			language.ROF:     "رومبو",
			language.ROOT:    "روٹ",
			language.RU:      "روسی",
			language.RUP:     "ارومانی",
			language.RW:      "کینیاروانڈا",
			language.RWK:     "روا",
			language.SA:      "سنسکرت",
			language.SAD:     "سنڈاوے",
			language.SAH:     "ساکھا",
			language.SAQ:     "سامبورو",
			language.SAT:     "سنتالی",
			language.SBA:     "نگامبے",
			language.SBP:     "سانگو",
			language.SC:      "سردینین",
			language.SCN:     "سیسیلین",
			language.SCO:     "سکاٹ",
			language.SD:      "سندھی",
			language.SDH:     "جنوبی کرد",
			language.SE:      "شمالی سامی",
			language.SEH:     "سینا",
			language.SES:     "کويرابورو سينی",
			language.SG:      "ساںغو",
			language.SH:      "سربو-کروئیشین",
			language.SHI:     "تشلحيت",
			language.SHN:     "شان",
			language.SI:      "سنہالا",
			language.SK:      "سلوواک",
			language.SL:      "سلووینیائی",
			language.SM:      "ساموآن",
			language.SMA:     "جنوبی سامی",
			language.SMJ:     "لول سامی",
			language.SMN:     "اناری سامی",
			language.SMS:     "سکولٹ سامی",
			language.SN:      "شونا",
			language.SNK:     "سوننکے",
			language.SO:      "صومالی",
			language.SQ:      "البانی",
			language.SR:      "سربین",
			language.SRN:     "سرانن ٹونگو",
			language.SS:      "سواتی",
			language.SSY:     "ساہو",
			language.ST:      "جنوبی سوتھو",
			language.SU:      "سنڈانیز",
			language.SUK:     "سکوما",
			language.SV:      "سویڈش",
			language.SW:      "سواحلی",
			language.SW_CD:   "کانگو سواحلی",
			language.SWB:     "کوموریائی",
			language.SYR:     "سریانی",
			language.TA:      "تمل",
			language.TE:      "تیلگو",
			language.TEM:     "ٹمنے",
			language.TEO:     "تیسو",
			language.TET:     "ٹیٹم",
			language.TG:      "تاجک",
			language.TH:      "تھائی",
			language.TI:      "ٹگرینیا",
			language.TIG:     "ٹگرے",
			language.TK:      "ترکمان",
			language.TL:      "ٹیگا لوگ",
			language.TLH:     "کلنگن",
			language.TN:      "سوانا",
			language.TO:      "ٹونگن",
			language.TPI:     "ٹوک پِسِن",
			language.TR:      "ترکی",
			language.TRV:     "ٹوروکو",
			language.TS:      "زونگا",
			language.TT:      "تاتار",
			language.TUM:     "ٹمبوکا",
			language.TVL:     "تووالو",
			language.TW:      "توی",
			language.TWQ:     "تاساواق",
			language.TY:      "تاہیتی",
			language.TYV:     "تووینین",
			language.TZM:     "سینٹرل ایٹلس ٹمازائٹ",
			language.UDM:     "ادمورت",
			language.UG:      "یوئگہر",
			language.UK:      "یوکرینیائی",
			language.UMB:     "اومبوندو",
			language.UND:     "نامعلوم زبان",
			language.UR:      "اردو",
			language.UZ:      "ازبیک",
			language.VAI:     "وائی",
			language.VE:      "وینڈا",
			language.VI:      "ویتنامی",
			language.VO:      "وولاپوک",
			language.VUN:     "ونجو",
			language.WA:      "والون",
			language.WAE:     "والسر",
			language.WAL:     "وولایتا",
			language.WAR:     "وارے",
			language.WBP:     "وارلپیری",
			language.WO:      "وولوف",
			language.XAL:     "کالمیک",
			language.XH:      "ژوسا",
			language.XOG:     "سوگا",
			language.YAV:     "یانگبین",
			language.YBB:     "یمبا",
			language.YI:      "یدش",
			language.YO:      "یوروبا",
			language.YUE:     "چینی، کینٹونیز",
			language.ZGH:     "معیاری مراقشی تمازیقی",
			language.ZH:      "چینی، مندارن",
			language.ZH_HANS: "آسان چینی",
			language.ZH_HANT: "روایتی مندارن چینی",
			language.ZU:      "زولو",
			language.ZUN:     "زونی",
			language.ZXX:     "کوئی لسانی مواد نہیں",
			language.ZZA:     "زازا",
		},
		Territories: cldr.Territories{
			territory.V_001: "دنیا",
			territory.V_002: "افریقہ",
			territory.V_003: "شمالی امریکہ",
			territory.V_005: "جنوبی امریکہ",
			territory.V_009: "اوشیانیا",
			territory.V_011: "مغربی افریقہ",
			territory.V_013: "وسطی امریکہ",
			territory.V_014: "مشرقی افریقہ",
			territory.V_015: "شمالی افریقہ",
			territory.V_017: "وسطی افریقہ",
			territory.V_018: "جنوبی افریقہ کے علاقہ",
			territory.V_019: "امیریکاز",
			territory.V_021: "شمالی امریکہ کا علاقہ",
			territory.V_029: "کریبیائی",
			territory.V_030: "مشرقی ایشیا",
			territory.V_034: "جنوبی ایشیا",
			territory.V_035: "جنوب مشرقی ایشیا",
			territory.V_039: "جنوبی یورپ",
			territory.V_053: "آسٹریلیشیا",
			territory.V_054: "مالینیشیا",
			territory.V_057: "مائکرونیشیائی علاقہ",
			territory.V_061: "پولینیشیا",
			territory.V_142: "ایشیا",
			territory.V_143: "وسطی ایشیا",
			territory.V_145: "مغربی ایشیا",
			territory.V_150: "یورپ",
			territory.V_151: "مشرقی یورپ",
			territory.V_154: "شمالی یورپ",
			territory.V_155: "مغربی یورپ",
			territory.V_202: "ذیلی صحارن افریقہ",
			territory.V_419: "لاطینی امریکہ",
			territory.AC:    "جزیرہ اسینشن",
			territory.AD:    "انڈورا",
			territory.AE:    "متحدہ عرب امارات",
			territory.AF:    "افغانستان",
			territory.AG:    "انٹیگوا اور باربودا",
			territory.AI:    "انگوئیلا",
			territory.AL:    "البانیہ",
			territory.AM:    "آرمینیا",
			territory.AO:    "انگولا",
			territory.AQ:    "انٹارکٹیکا",
			territory.AR:    "ارجنٹینا",
			territory.AS:    "امریکی ساموآ",
			territory.AT:    "آسٹریا",
			territory.AU:    "آسٹریلیا",
			territory.AW:    "اروبا",
			territory.AX:    "جزائر آلینڈ",
			territory.AZ:    "آذربائیجان",
			territory.BA:    "بوسنیا اور ہرزیگووینا",
			territory.BB:    "بارباڈوس",
			territory.BD:    "بنگلہ دیش",
			territory.BE:    "بیلجیم",
			territory.BF:    "برکینا فاسو",
			territory.BG:    "بلغاریہ",
			territory.BH:    "بحرین",
			territory.BI:    "برونڈی",
			territory.BJ:    "بینن",
			territory.BL:    "سینٹ برتھلیمی",
			territory.BM:    "برمودا",
			territory.BN:    "برونائی",
			territory.BO:    "بولیویا",
			territory.BQ:    "کریبیائی نیدرلینڈز",
			territory.BR:    "برازیل",
			territory.BS:    "بہاماس",
			territory.BT:    "بھوٹان",
			territory.BV:    "جزیرہ بوویت",
			territory.BW:    "بوتسوانا",
			territory.BY:    "بیلاروس",
			territory.BZ:    "بیلائز",
			territory.CA:    "کینیڈا",
			territory.CC:    "جزائر (کیلنگ) کوکوس",
			territory.CD:    "کانگو (DRC)",
			territory.CF:    "وسط افریقی جمہوریہ",
			territory.CG:    "کانگو (جمہوریہ)",
			territory.CH:    "سوئٹزر لینڈ",
			territory.CI:    "آئیوری کوسٹ",
			territory.CK:    "جزائر کک",
			territory.CL:    "چلی",
			territory.CM:    "کیمرون",
			territory.CN:    "چین",
			territory.CO:    "کولمبیا",
			territory.CP:    "جزیرہ کلپرٹن",
			territory.CR:    "کوسٹا ریکا",
			territory.CU:    "کیوبا",
			territory.CV:    "کیپ ورڈی",
			territory.CW:    "کیوراکاؤ",
			territory.CX:    "جزیرہ کرسمس",
			territory.CY:    "قبرص",
			territory.CZ:    "چیک جمہوریہ",
			territory.DE:    "جرمنی",
			territory.DG:    "ڈیگو گارشیا",
			territory.DJ:    "جبوتی",
			territory.DK:    "ڈنمارک",
			territory.DM:    "ڈومنیکا",
			territory.DO:    "جمہوریہ ڈومينيکن",
			territory.DZ:    "الجیریا",
			territory.EA:    "سیئوٹا اور میلیلا",
			territory.EC:    "ایکواڈور",
			territory.EE:    "اسٹونیا",
			territory.EG:    "مصر",
			territory.EH:    "مغربی صحارا",
			territory.ER:    "اریٹیریا",
			territory.ES:    "ہسپانیہ",
			territory.ET:    "ایتھوپیا",
			territory.EU:    "یوروپی یونین",
			territory.EZ:    "یوروزون",
			territory.FI:    "فن لینڈ",
			territory.FJ:    "فجی",
			territory.FK:    "جزائر فاکلینڈ (اسلاس مالویناس)",
			territory.FM:    "مائکرونیشیا",
			territory.FO:    "جزائر فیرو",
			territory.FR:    "فرانس",
			territory.GA:    "گیبون",
			territory.GB:    "یو کے",
			territory.GD:    "گریناڈا",
			territory.GE:    "جارجیا",
			territory.GF:    "فرانسیسی گیانا",
			territory.GG:    "گوئرنسی",
			territory.GH:    "گھانا",
			territory.GI:    "جبل الطارق",
			territory.GL:    "گرین لینڈ",
			territory.GM:    "گیمبیا",
			territory.GN:    "گنی",
			territory.GP:    "گواڈیلوپ",
			territory.GQ:    "استوائی گیانا",
			territory.GR:    "یونان",
			territory.GS:    "جنوبی جارجیا اور جنوبی سینڈوچ جزائر",
			territory.GT:    "گواٹے مالا",
			territory.GU:    "گوام",
			territory.GW:    "گنی بساؤ",
			territory.GY:    "گیانا",
			territory.HK:    "ہانگ کانگ",
			territory.HM:    "جزائر ہرڈ و مکڈونلڈ",
			territory.HN:    "ہونڈاروس",
			territory.HR:    "کروشیا",
			territory.HT:    "ہیٹی",
			territory.HU:    "ہنگری",
			territory.IC:    "جزائر کناری",
			territory.ID:    "انڈونیشیا",
			territory.IE:    "آئرلینڈ",
			territory.IL:    "اسرائیل",
			territory.IM:    "آئل آف مین",
			territory.IN:    "بھارت",
			territory.IO:    "برطانوی بحرہند خطہ",
			territory.IQ:    "عراق",
			territory.IR:    "ایران",
			territory.IS:    "آئس لینڈ",
			territory.IT:    "اٹلی",
			territory.JE:    "جرسی",
			territory.JM:    "جمائیکا",
			territory.JO:    "اردن",
			territory.JP:    "جاپان",
			territory.KE:    "کینیا",
			territory.KG:    "کرغزستان",
			territory.KH:    "کمبوڈیا",
			territory.KI:    "کریباتی",
			territory.KM:    "کوموروس",
			territory.KN:    "سینٹ کٹس اور نیویس",
			territory.KP:    "شمالی کوریا",
			territory.KR:    "جنوبی کوریا",
			territory.KW:    "کویت",
			territory.KY:    "کیمین آئلینڈز",
			territory.KZ:    "قزاخستان",
			territory.LA:    "لاؤس",
			territory.LB:    "لبنان",
			territory.LC:    "سینٹ لوسیا",
			territory.LI:    "لیشٹنسٹائن",
			territory.LK:    "سری لنکا",
			territory.LR:    "لائبیریا",
			territory.LS:    "لیسوتھو",
			territory.LT:    "لیتھونیا",
			territory.LU:    "لکسمبرگ",
			territory.LV:    "لٹویا",
			territory.LY:    "لیبیا",
			territory.MA:    "مراکش",
			territory.MC:    "موناکو",
			territory.MD:    "مالدووا",
			territory.ME:    "مونٹے نیگرو",
			territory.MF:    "سینٹ مارٹن",
			territory.MG:    "مڈغاسکر",
			territory.MH:    "جزائر مارشل",
			territory.MK:    "شمالی مقدونیہ",
			territory.ML:    "مالی",
			territory.MM:    "میانمار (برما)",
			territory.MN:    "منگولیا",
			territory.MO:    "مکاؤ",
			territory.MP:    "جزائر شمالی ماریانا",
			territory.MQ:    "مارٹینک",
			territory.MR:    "موریطانیہ",
			territory.MS:    "مونٹسیراٹ",
			territory.MT:    "مالٹا",
			territory.MU:    "ماریشس",
			territory.MV:    "مالدیپ",
			territory.MW:    "ملاوی",
			territory.MX:    "میکسیکو",
			territory.MY:    "ملائشیا",
			territory.MZ:    "موزمبیق",
			territory.NA:    "نامیبیا",
			territory.NC:    "نیو کلیڈونیا",
			territory.NE:    "نائجر",
			territory.NF:    "جزیرہ نارفوک",
			territory.NG:    "نائجیریا",
			territory.NI:    "نکاراگووا",
			territory.NL:    "نیدر لینڈز",
			territory.NO:    "ناروے",
			territory.NP:    "نیپال",
			territory.NR:    "نؤرو",
			territory.NU:    "نیئو",
			territory.NZ:    "نیوزی لینڈ",
			territory.OM:    "عمان",
			territory.PA:    "پانامہ",
			territory.PE:    "پیرو",
			territory.PF:    "فرانسیسی پولینیشیا",
			territory.PG:    "پاپوآ نیو گنی",
			territory.PH:    "فلپائن",
			territory.PK:    "پاکستان",
			territory.PL:    "پولینڈ",
			territory.PM:    "سینٹ پیئر اور میکلیئون",
			territory.PN:    "جزائر پٹکیرن",
			territory.PR:    "پیورٹو ریکو",
			territory.PS:    "فلسطین",
			territory.PT:    "پرتگال",
			territory.PW:    "پلاؤ",
			territory.PY:    "پیراگوئے",
			territory.QA:    "قطر",
			territory.QO:    "بیرونی اوشیانیا",
			territory.RE:    "ری یونین",
			territory.RO:    "رومانیہ",
			territory.RS:    "سربیا",
			territory.RU:    "روس",
			territory.RW:    "روانڈا",
			territory.SA:    "سعودی عرب",
			territory.SB:    "جزائر سلیمان",
			territory.SC:    "سشلیز",
			territory.SD:    "سوڈان",
			territory.SE:    "سویڈن",
			territory.SG:    "سنگاپور",
			territory.SH:    "سینٹ ہیلینا",
			territory.SI:    "سلووینیا",
			territory.SJ:    "سوالبرڈ اور جان ماین",
			territory.SK:    "سلوواکیہ",
			territory.SL:    "سیرالیون",
			territory.SM:    "سان مارینو",
			territory.SN:    "سینیگل",
			territory.SO:    "صومالیہ",
			territory.SR:    "سورینام",
			territory.SS:    "جنوبی سوڈان",
			territory.ST:    "ساؤ ٹومے اور پرنسپے",
			territory.SV:    "ال سلواڈور",
			territory.SX:    "سنٹ مارٹن",
			territory.SY:    "شام",
			territory.SZ:    "سوازی لینڈ",
			territory.TA:    "ترسٹان دا کونیا",
			territory.TC:    "جزائر کیکس و ترکیہ",
			territory.TD:    "چاڈ",
			territory.TF:    "فرانسیسی جنوبی خطے",
			territory.TG:    "ٹوگو",
			territory.TH:    "تھائی لینڈ",
			territory.TJ:    "تاجکستان",
			territory.TK:    "ٹوکیلاؤ",
			territory.TL:    "مشرقی تیمور",
			territory.TM:    "ترکمانستان",
			territory.TN:    "تونس",
			territory.TO:    "ٹونگا",
			territory.TR:    "ترکی",
			territory.TT:    "ترینیداد اور ٹوباگو",
			territory.TV:    "ٹووالو",
			territory.TW:    "تائیوان",
			territory.TZ:    "تنزانیہ",
			territory.UA:    "یوکرین",
			territory.UG:    "یوگنڈا",
			territory.UM:    "امریکی بیرونی جزائر",
			territory.UN:    "اقوام متحدہ",
			territory.US:    "امریکا",
			territory.UY:    "یوروگوئے",
			territory.UZ:    "ازبکستان",
			territory.VA:    "ویٹیکن سٹی",
			territory.VC:    "سینٹ ونسنٹ اور گرینیڈائنز",
			territory.VE:    "وینزوئیلا",
			territory.VG:    "برطانوی جزائر ورجن",
			territory.VI:    "امریکی جزائر ورجن",
			territory.VN:    "ویتنام",
			territory.VU:    "وینوآٹو",
			territory.WF:    "ویلیز اور فیوٹیونا",
			territory.WS:    "ساموآ",
			territory.XA:    "بناوٹی لہجے",
			territory.XB:    "مصنوعی بیڑی",
			territory.XK:    "کوسووو",
			territory.YE:    "یمن",
			territory.YT:    "مایوٹ",
			territory.ZA:    "جنوبی افریقہ",
			territory.ZM:    "زامبیا",
			territory.ZW:    "زمبابوے",
			territory.ZZ:    "نامعلوم علاقہ",
		},
	}
}
