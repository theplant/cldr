// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_sd() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "جنوري", Feb: "فيبروري", Mar: "مارچ", Apr: "اپريل", May: "مئي", Jun: "جون", Jul: "جولاءِ", Aug: "آگسٽ", Sep: "سيپٽمبر", Oct: "آڪٽوبر", Nov: "نومبر", Dec: "ڊسمبر"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "جنوري", Feb: "فيبروري", Mar: "مارچ", Apr: "اپريل", May: "مئي", Jun: "جون", Jul: "جولاءِ", Aug: "آگسٽ", Sep: "سيپٽمبر", Oct: "آڪٽوبر", Nov: "نومبر", Dec: "ڊسمبر"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "آچر", Mon: "سومر", Tue: "اڱارو", Wed: "اربع", Thu: "خميس", Fri: "جمعو", Sat: "ڇنڇر"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "آچر", Mon: "سو", Tue: "اڱارو", Wed: "اربع", Thu: "خم", Fri: "جمعو", Sat: "ڇنڇر"}, Short: cldr.CalendarDayFormatNameValue{Sun: "آچر", Mon: "سومر", Tue: "اڱارو", Wed: "اربع", Thu: "خميس", Fri: "جمعو", Sat: "ڇنڇر"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "آچر", Mon: "سومر", Tue: "اڱارو", Wed: "اربع", Thu: "خميس", Fri: "جمعو", Sat: "ڇنڇر"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "صبح، منجهند", PM: "منجهند، شام"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "صبح، منجهند", PM: "منجهند، شام"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "صبح، منجهند", PM: "منجهند، شام"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_sd]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "گڏيل عرب امارات درهم", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "افغاني افغاني", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "الباني ليڪ", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "ارماني ڊرم", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "نيڌرلينڊ انٽليئن گلڊر", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "انگوليائي ڪوانزا", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "ارجنٽائن پيسو", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "آسٽريلوي ڊالر", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "اروبن فلورن", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "آذربائيجاني منت", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "بوسنيا هرزگوينا ڪنورٽبل مارڪ", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "باربيڊين ڊالر", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "بنگلاديشي ٽڪا", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "بلغارین لیو", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "بحريني دينار", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "برونڊي فرينڪ", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "برمودي ڊالر", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "برونائي ڊالر", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "بولیوین بولیویانو", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "برازيلي ريل", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "بهاماني ڊادلر", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ڀوٽاني گلٽرم", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "بوستواني پولا", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "بیلاروسی ربل", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "بيليز ڊالر", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "ڪئينڊيائي ڊالر", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "ڪانگو فرينڪ", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "سوئس فرينڪ", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "چلي پيسو", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "چيني يوآن (غير ملڪي)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "چيني يوآن", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "ڪولمبيائي پيسو", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "ڪوسٽا ريڪا ڪولن", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "ڪيوبن ڪنورٽيبل پيسو", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "ڪيوبن پيسو", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "ڪيپ وردي ايسڪوڊو", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "چيڪ ڪرونا", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "جبوتي فرينڪ", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "دانش ڪرون", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "ڊومينيڪن پيسو", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "الجيريائي دينار", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "مصري پائونڊ", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "ايريٽيريائي ناڪفا", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ايٿوپيائي بر", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "يورو", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "فجي ڊالر", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "فاڪلينڊ ٻيٽ پائونڊ", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "برطانوي پائونڊ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "جارجيائي لاري", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "گهانين سيدي", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "جبرالٽر پائونڊ", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "گيمبيا دلاسائي", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "گني فرينڪ", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "گوئٽي مالا ڪٽزل", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "گيانا ڊالر", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "هانگ ڪانگ ڊالر", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "هونڊوراس ليمپرا", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "ڪروشيائي ڪونا", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "هيٽي گورڊي", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "هنگيرين فورنٽ", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "انڊونيشيائي رپيه", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "اسرائيلي نيو شيڪل", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "انڊين رپي", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "عراقي دينار", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ايراني ريال", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "آئيس لينڊي ڪرونا", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "جميڪائي ڊالر", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "اردني دينار", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "جاپاني يين", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "ڪينيائي سلنگ", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "ڪرغزستاني سوم", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "ڪمبوڊيائي ريال", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "ڪوموريائي فرينڪ", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "اتر ڪوريائي ون", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "ڏکڻ ڪوريائي ون", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "ڪويتي دينار", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "ڪيمين ٻيٽ ڊالر", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "قازقستان ٽينگا", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "لائوشيائي ڪپ", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "لبناني پائونڊ", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "سري لنڪن رپي", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "لائبیریائی ڊالر", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "لبيائي دينار", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "موروڪيائي درهم", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "مالديپ ليو", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ملاگاسي اریاری", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "ميسي ڊوني دينار", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "ميانمار ڪياٽ", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "منگولي تجرڪ", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "ميڪانيز پٽاڪا", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "موريشيائي اوگوئیا (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "موريشيائي اوگوئیا", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "ماريشيائي رپي", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "مالديپ روفيا", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "ملاوي ڪواچا", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "ميڪسيڪو پيسو", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "ملائيشيائي رنگٽ", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "موزمبيق ميٽيڪل", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "نميبائي ڊالر", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "نائجريائي نائرا", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "نڪارا گوا ڪارڊوبا", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "نارويائي ڪرون", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "نيپالي رپي", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "نيوزي لينڊي ڊالر", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "عماني ريال", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "پاناما پالبوا", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "پيرو سول", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "پاپوا نيو گني ڪنا", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "فلپائني پيسو", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "پاڪستاني رپي", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "پولش زلاٽي", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "پيراگوئي گاراني", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "قطري ريال", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "رومانیائي لیو", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "سربيا دينار", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "روسي ربل", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "روانڊا فرينڪ", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "سعودي ريال", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "سولومان ٻيٽ ڊالر", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "سشلي رپي", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "سوڊاني پائونڊ", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "سويڊني ڪرونا", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "سنگاپوري ڊالر", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "سينٽ هيلنا پائونڊ", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "سیرا لیونيائي لیون", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "سومالي شلنگ", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "سرينامي ڊالر", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "ڏکڻ سوڊاني پائونڊ", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "سائو ٽوم ۽ پرنسپي دوبرا (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "سائو ٽوم ۽ پرنسپي دوبرا", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "سيريائي پائونڊ", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "سوازي للانگيني", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "ٿائي باهت", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "تاجڪستاني سوموني", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "ترڪمانستان منت", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "تیونس دینار", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "تونگن پانگا", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "ترڪي لرا", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ٽرينڊيڊ ۽ ٽوباگو ڊالر", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "نيو تائيوان ڊالر", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "تنزانيائي شلنگ", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "یوڪرائن هریونیا", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "يگانڊا شلنگ", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "آمريڪي ڊالر", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "يوروگوئي پيسو", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "ازبڪستاني سوم", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "Venezuelan Bolívar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "وینزویلا بولیور", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "ويٽنامي ڊونگ", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "وانواتو واتو", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "ساموآن ٽالا", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "وچ آفريڪا فرينڪ", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "اوڀر ڪيريبين ڊالر", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "اولهه آفريڪا فرينڪ", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP فرينڪ", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "اڻڄاتل سڪو", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "يمني ريال", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "ڏکڻ آفريقي رانڊ", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "زمبائي ڪواچا", Symbol: "ZMW"},
			},
		},
	}
}
