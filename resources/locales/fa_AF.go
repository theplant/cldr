// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_fa_AF() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y/M/d"}, Time: cldr.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss (z)", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}، ساعت {0}", Long: "{1}، ساعت {0}", Medium: "{1}،\u200f {0}", Short: "{1}،\u200f {0}"}, GMT: "{0} گرینویچ"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "جنوری", Feb: "فبروری", Mar: "مارچ", Apr: "اپریل", May: "می", Jun: "جون", Jul: "جولای", Aug: "اگست", Sep: "سپتمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ج", Feb: "ف", Mar: "م", Apr: "ا", May: "م", Jun: "ج", Jul: "ج", Aug: "ا", Sep: "س", Oct: "ا", Nov: "ن", Dec: "د"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "جنوری", Feb: "فبروری", Mar: "مارچ", Apr: "اپریل", May: "می", Jun: "جون", Jul: "جولای", Aug: "اگست", Sep: "سپتمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "یکشنبه", Mon: "دوشنبه", Tue: "سه\u200cشنبه", Wed: "چهارشنبه", Thu: "پنجشنبه", Fri: "جمعه", Sat: "شنبه"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "ی", Mon: "د", Tue: "س", Wed: "چ", Thu: "پ", Fri: "ج", Sat: "ش"}, Short: cldr.CalendarDayFormatNameValue{Sun: "۱ش", Mon: "۲ش", Tue: "۳ش", Wed: "۴ش", Thu: "۵ش", Fri: "ج", Sat: "ش"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "یکشنبه", Mon: "دوشنبه", Tue: "سه\u200cشنبه", Wed: "چهارشنبه", Thu: "پنجشنبه", Fri: "جمعه", Sat: "شنبه"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ق.ظ.", PM: "ب.ظ."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ق", PM: "ب"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "قبل\u200cازظهر", PM: "بعدازظهر"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_fa]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "", Percent: "٪", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "0 هزار", Currency: "¤\u00a0#,##0.00", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "پزتای آندورا", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "درهم امارات متحدهٔ عربی", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "افغانی افغانستان (۱۹۲۷ تا ۲۰۰۲)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "افغانی افغانستان", Symbol: "؋"},
				currency.ALK: cldr.Currency{DisplayName: "لک آلبانی (۱۹۴۶ تا ۱۹۶۵)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "لک آلبانی", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "درام ارمنستان", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "گیلدر آنتیل هلند", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "کوانزای آنگولا", Symbol: "AOA"},
				currency.ARM: cldr.Currency{DisplayName: "پزوی آرژانتین (۱۸۸۱ تا ۱۹۷۰)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "پزوی آرژانتین (۱۹۸۳ تا ۱۹۸۵)\u200f", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "پزوی آرژانتین", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "شیلینگ اتریش", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "دالر آسترالیا", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "فلورین آروبا", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "منات جمهوری آذربایجان (۱۹۹۳ تا ۲۰۰۶)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "منات جمهوری آذربایجان", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "دینار بوسنی و هرزگوین (۱۹۹۲ تا ۱۹۹۴)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "مارک تبدیل\u200cپذیر بوسنی و هرزگوین", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "دلار باربادوس", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "تاکای بنگلادش", Symbol: "BDT"},
				currency.BEF: cldr.Currency{DisplayName: "فرانک بلژیک", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "فرانک بلژیک (مالی)", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "لف بلغارستان", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "دینار بحرین", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "فرانک بوروندی", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "دلار برمودا", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "دالر برونی", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "بولیویانوی بولیوی", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "پزوی بولیوی", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "رئال برزیل", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "دلار باهاما", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "انگولتروم بوتان", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "پولای بوتسوانا", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "روبل جدید بیلوروسی (۱۹۹۴ تا ۱۹۹۹)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "روبل روسیهٔ سفید", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "روبل روسیهٔ سفید (۲۰۰۰–۲۰۱۶)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "دلار بلیز", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "دالر کانادا", Symbol: "$CA"},
				currency.CDF: cldr.Currency{DisplayName: "فرانک کنگو", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "فرانک سویس", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "پزوی شیلی", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "یوآن چین (برون\u200cمرزی)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "یوآن چین", Symbol: "¥CN"},
				currency.COP: cldr.Currency{DisplayName: "پزوی کلمبیا", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "کولون کاستاریکا", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "دینار صربستان (۲۰۰۲ تا ۲۰۰۶)", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "پزوی تبدیل\u200cپذیر کوبا", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "پزوی کوبا", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "اسکودوی کیپ\u200cورد", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "پوند قبرس", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "کورونای جمهوری چک", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "مارک آلمان شرقی", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "مارک آلمان", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "فرانک جیبوتی", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "کرون دنمارک", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "پزوی جمهوری دومینیکن", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "دینار الجزایر", Symbol: "DZD"},
				currency.EEK: cldr.Currency{DisplayName: "کرون استونی", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "پوند مصر", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "ناکفای اریتره", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "پزتای اسپانیا", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "بیر اتیوپی", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "یورو", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "دلار فیجی", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "پوند جزایر فالکلند", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "فرانک فرانسه", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "پوند بریتانیا", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "لاری گرجستان", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "سدی غنا", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "پوند جبل\u200cالطارق", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "دالاسی گامبیا", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "فرانک گینه", Symbol: "GNF"},
				currency.GRD: cldr.Currency{DisplayName: "دراخمای یونان", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "کتزال گواتمالا", Symbol: "GTQ"},
				currency.GWP: cldr.Currency{DisplayName: "پزوی گینهٔ بیسائو", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "دلار گویانا", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "دلار هنگ\u200cکنگ", Symbol: "$HK"},
				currency.HNL: cldr.Currency{DisplayName: "لمپیرای هندوراس", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "دینار کرواسی", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "کونای کرواسی", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "گورد هائیتی", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "فورینت مجارستان", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "روپیهٔ اندونزی", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "پوند ایرلند", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "پوند اسرائیل", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "شقل اسرائیل (۱۹۸۰ تا ۱۹۸۵)", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "شقل جدید اسرائیل", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "روپیهٔ هند", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "دینار عراق", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ریال ایران", Symbol: "ریال"},
				currency.ISJ: cldr.Currency{DisplayName: "کرونای ایسلند (۱۹۱۸ تا ۱۹۸۱)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "کرونای ایسلند", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "لیرهٔ ایتالیا", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "دلار جامائیکا", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "دینار اردن", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "ین جاپان", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "شیلینگ کنیا", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "سوم قرقیزستان", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "ری\u200cیل کامبوج", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "فرانک کومورو", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "وون کرهٔ شمالی", Symbol: "KPW"},
				currency.KRO: cldr.Currency{DisplayName: "وون کرهٔ جنوبی (۱۹۴۵ تا ۱۹۵۳)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "وون کرهٔ جنوبی", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "دینار کویت", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "دلار جزایر کِیمن", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "تنگهٔ قزاقستان", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "کیپ لائوس", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "لیرهٔ لبنان", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "روپیهٔ سری\u200cلانکا", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "دلار لیبریا", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "لوتی لسوتو", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "لیتاس لیتوانی", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "تالوناس لیتوانی", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "فرانک لوکزامبورگ", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "فرانک مالی لوگزامبورگ", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "لاتس لتونی", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "روبل لتونی", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "دینار لیبی", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "درهم مراکش", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "فرانک مراکش", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "فرانک موناکو", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "لئوی مولداوی", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "آریاری مالاگاسی", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "فرانک ماداگاسکار", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "دینار مقدونیه", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "دینار مقدونیه (۱۹۹۲ تا ۱۹۹۳)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "فرانک مالی", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "کیات میانمار", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "توگریک مغولستان", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "پاتاکای ماکائو", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "اوگوئیای موریتانی (۱۹۷۳ تا ۲۰۱۷)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "اوگوئیای موریتانی", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "لیرهٔ مالت", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "پوند مالت", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "روپیهٔ موریس", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "روپیهٔ مالدیو (۱۹۴۷ تا ۱۹۸۱)", Symbol: ""},
				currency.MVR: cldr.Currency{DisplayName: "روپیهٔ مالدیو", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "کواچای مالاوی", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "پزوی مکسیکو", Symbol: "$MX"},
				currency.MXP: cldr.Currency{DisplayName: "پزوی نقرهٔ مکزیک (۱۸۶۱ تا ۱۹۹۲)", Symbol: "MXP"},
				currency.MYR: cldr.Currency{DisplayName: "رینگیت مالزی", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "اسکودوی موزامبیک", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "متیکال موزامبیک", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "دلار نامیبیا", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "نایرای نیجریه", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "کوردوبای نیکاراگوئه", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "گیلدر هالند", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "کرون ناروی", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "روپیهٔ نپال", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "دلار زلاند نو", Symbol: "$NZ"},
				currency.OMR: cldr.Currency{DisplayName: "ریال عمان", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "بالبوای پاناما", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "اینتی پرو", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "سول پرو", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "سول پرو (۱۸۶۳ تا ۱۹۶۵)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "کینای پاپوا گینهٔ نو", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "پزوی فیلیپین", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "روپیهٔ پاکستان", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "زواتی لهستان", Symbol: "PLN"},
				currency.PTE: cldr.Currency{DisplayName: "اسکودوی پرتغال", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "گوارانی پاراگوئه", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "ریال قطر", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "دلار رودزیا", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "لئوی رومانی", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "دینار صربستان", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "روبل روسیه", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "روبل روسیه (۱۹۹۱ تا ۱۹۹۸)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "فرانک رواندا", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "ریال سعودی", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "دلار جزایر سلیمان", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "روپیهٔ سیشل", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "دینار سودان (۱۹۹۲ تا ۲۰۰۷)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "پوند سودان", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "کرون سویدن", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "دالر سینگاپور", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "پوند سنت هلن", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "لئون سیرالئون", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "شیلینگ سومالی", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "دلار سورینام", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "گیلدر سورینام", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "پوند سودان جنوبی", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "دوبرای سائوتومه و پرنسیپ (۱۹۷۷ تا ۲۰۱۷)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "دوبرای سائوتومه و پرنسیپ", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "روبل شوروی", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "لیرهٔ سوریه", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "لیلانگنی سوازیلند", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "بات تایلند", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "روبل تاجیکستان", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "سامانی تاجکستان", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "منات ترکمنستان (۱۹۹۳ تا ۲۰۰۹)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "منات ترکمنستان", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "دینار تونس", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "پاآنگای تونگا", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "اسکودوی تیمور", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "لیرهٔ ترکیه (۱۹۲۲ تا ۲۰۰۵)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "لیرهٔ ترکیه", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "دلار ترینیداد و توباگو", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "دلار جدید تایوان", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "شیلینگ تانزانیا", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "هریونیای اوکراین", Symbol: "UAH"},
				currency.UGS: cldr.Currency{DisplayName: "شیلینگ اوگاندا (۱۹۶۶ تا ۱۹۸۷)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "شیلینگ اوگاندا", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "دالر امریکا", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "دلار امریکا (روز بعد)", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "دلار امریکا (همان روز)", Symbol: "USS"},
				currency.UYP: cldr.Currency{DisplayName: "پزوی اوروگوئه (۱۹۷۵ تا ۱۹۹۳)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "پزوی اوروگوئه", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "سوم ازبکستان", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "بولیوار ونزوئلا (۱۸۷۱ تا ۲۰۰۸)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "بولیوار ونزوئلا (۲۰۰۸ تا ۲۰۱۸)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "بولیوار ونزوئلا", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "دانگ ویتنام", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "واتوی وانوواتو", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "تالای ساموا", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "فرانک CFA مرکز افریقا", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "نقره", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "طلا", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "دلار شرق کارائیب", Symbol: "$EC"},
				currency.XFO: cldr.Currency{DisplayName: "فرانک طلای فرانسه", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "فرانک CFA غرب افریقا", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "پالادیم", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "فرانک اقیانوسیه", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "پلاتین", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "ارز آزمایشی", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "ارز نامشخص", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "دینار یمن", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "ریال یمن", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "راند افریقای جنوبی", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "کواچای زامبیا (۱۹۶۸ تا ۲۰۱۲)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "کواچای زامبیا", Symbol: "ZMW"},
				currency.ZWD: cldr.Currency{DisplayName: "دلار زیمبابوه (۱۹۸۰ تا ۲۰۰۸)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "دلار زیمبابوه (۲۰۰۹)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "دلار زیمبابوه (۲۰۰۸)", Symbol: ""},
			},
		},
	}
}
