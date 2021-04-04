// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_sd_Arab_PK() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "sd_Arab_PK",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "جنوري", Feb: "فيبروري", Mar: "مارچ", Apr: "اپريل", May: "مئي", Jun: "جون", Jul: "جولاءِ", Aug: "آگسٽ", Sep: "سيپٽمبر", Oct: "آڪٽوبر", Nov: "نومبر", Dec: "ڊسمبر"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "جنوري", Feb: "فيبروري", Mar: "مارچ", Apr: "اپريل", May: "مئي", Jun: "جون", Jul: "جولاءِ", Aug: "آگسٽ", Sep: "سيپٽمبر", Oct: "آڪٽوبر", Nov: "نومبر", Dec: "ڊسمبر"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "آچر", Mon: "سومر", Tue: "اڱارو", Wed: "اربع", Thu: "خميس", Fri: "جمعو", Sat: "ڇنڇر"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "آچر", Mon: "سو", Tue: "اڱارو", Wed: "اربع", Thu: "خم", Fri: "جمعو", Sat: "ڇنڇر"}, Short: cldr.CalendarDayFormatNameValue{Sun: "آچر", Mon: "سومر", Tue: "اڱارو", Wed: "اربع", Thu: "خميس", Fri: "جمعو", Sat: "ڇنڇر"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "آچر", Mon: "سومر", Tue: "اڱارو", Wed: "اربع", Thu: "خميس", Fri: "جمعو", Sat: "ڇنڇر"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "صبح، منجهند", PM: "منجهند، شام"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "صبح، منجهند", PM: "منجهند، شام"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "صبح، منجهند", PM: "منجهند، شام"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", CurrencyAccounting: "¤\u00a0#,##0.00", Percent: "#,##0%"},
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
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "افار",
			language.AB:      "ابقازیان",
			language.ACE:     "اچائينيز",
			language.ADA:     "ادنگمي",
			language.ADY:     "اديگهي",
			language.AF:      "آفريڪي",
			language.AGQ:     "اگهيم",
			language.AIN:     "آئينو",
			language.AK:      "اڪان",
			language.ALE:     "اليوٽ",
			language.ALT:     "ڏکڻ التائي",
			language.AM:      "امهاري",
			language.AN:      "ارگني",
			language.ANP:     "انجيڪا",
			language.AR:      "عربي",
			language.AR_001:  "جديد معياري عربي",
			language.ARN:     "ماپوچي",
			language.ARP:     "اراپائو",
			language.AS:      "آسامي",
			language.ASA:     "اسو",
			language.AST:     "اسٽورين",
			language.AV:      "اويرس",
			language.AWA:     "اواڌي",
			language.AY:      "ایمارا",
			language.AZ:      "ازري",
			language.BA:      "بشڪر",
			language.BAN:     "بالي",
			language.BAS:     "باسا",
			language.BE:      "بيلاروسي",
			language.BEM:     "بيمبا",
			language.BEZ:     "بينا",
			language.BG:      "بلغاريائي",
			language.BHO:     "ڀوجپوري",
			language.BI:      "بسلاما",
			language.BIN:     "بني",
			language.BLA:     "سڪسڪا",
			language.BM:      "بمبارا",
			language.BN:      "بنگلا",
			language.BO:      "تبيتائي",
			language.BR:      "بريٽن",
			language.BRX:     "بودو",
			language.BS:      "بوسنيائي",
			language.BUG:     "بگنيز",
			language.BYN:     "بلن",
			language.CA:      "ڪيٽالان",
			language.CCP:     "چمڪا",
			language.CE:      "چیچن",
			language.CEB:     "سبوانو",
			language.CGG:     "چگا",
			language.CH:      "چمورو",
			language.CHK:     "چڪيز",
			language.CHM:     "ماري",
			language.CHO:     "چوڪ تو",
			language.CHR:     "چروڪي",
			language.CHY:     "چايان",
			language.CKB:     "مرڪزي ڪردش",
			language.CO:      "ڪارسيڪائي",
			language.CRS:     "سيسلوا ڪريئول فرانسي",
			language.CS:      "چيڪ",
			language.CU:      "چرچ سلاوی",
			language.CV:      "چو واش",
			language.CY:      "ويلش",
			language.DA:      "ڊينش",
			language.DAK:     "ڊڪوٽا",
			language.DAR:     "ڊارگوا",
			language.DAV:     "تائيتا",
			language.DE:      "جرمن",
			language.DE_AT:   "آسٽريائي جرمن",
			language.DE_CH:   "سوئس هائي جرمن",
			language.DGR:     "داگرب",
			language.DJE:     "زارما",
			language.DSB:     "لوئر سوربين",
			language.DUA:     "ڊيولا",
			language.DV:      "دويهي",
			language.DYO:     "جولا فوني",
			language.DZ:      "زونخا",
			language.DZG:     "دزاگا",
			language.EBU:     "ايمبيو",
			language.EE:      "ايو",
			language.EFI:     "ايفڪ",
			language.EKA:     "ايڪاجڪ",
			language.EL:      "يوناني",
			language.EN:      "انگريزي",
			language.EN_AU:   "آسٽريليائي انگريزي",
			language.EN_CA:   "ڪينيڊيائي انگريزي",
			language.EN_GB:   "برطانوي انگريزي",
			language.EN_US:   "انگريزي (آمريڪا)",
			language.EO:      "ايسپرانٽو",
			language.ES:      "اسپيني",
			language.ES_419:  "لاطيني آمريڪي اسپينش",
			language.ES_ES:   "يورپي اسپيني",
			language.ES_MX:   "ميڪسيڪين اسپيني",
			language.ET:      "ايستونائي",
			language.EU:      "باسڪي",
			language.EWO:     "اوانڊو",
			language.FA:      "فارسي",
			language.FA_AF:   "دري",
			language.FF:      "فلاهه",
			language.FI:      "فنش",
			language.FIL:     "فلپائني",
			language.FJ:      "فجي",
			language.FO:      "فيروايس",
			language.FON:     "فون",
			language.FR:      "فرانسي",
			language.FR_CA:   "ڪينيڊيائي فرانسيسي",
			language.FR_CH:   "سوئس فرانسيسي",
			language.FUR:     "فرائي لئين",
			language.FY:      "مغربي فريشن",
			language.GA:      "آئرش",
			language.GAA:     "گا",
			language.GD:      "اسڪاٽش گيلڪ",
			language.GEZ:     "جيز",
			language.GIL:     "گلبرٽيز",
			language.GL:      "گليشئين",
			language.GN:      "گواراني",
			language.GOR:     "گورنٽلو",
			language.GSW:     "سوئس جرمن",
			language.GU:      "گجراتي",
			language.GUZ:     "گشي",
			language.GV:      "مينڪس",
			language.GWI:     "گوچن",
			language.HA:      "هوسا",
			language.HAW:     "هوائي",
			language.HE:      "عبراني",
			language.HI:      "هندي",
			language.HIL:     "هلي گيانان",
			language.HMN:     "مونگ",
			language.HR:      "ڪروشيائي",
			language.HSB:     "اپر سربيائي",
			language.HT:      "هيٽي ڪرولي",
			language.HU:      "هنگري",
			language.HUP:     "هوپا",
			language.HY:      "ارماني",
			language.HZ:      "هريرو",
			language.IA:      "انٽرلنگئا",
			language.IBA:     "ايبن",
			language.IBB:     "ابيبيو",
			language.ID:      "انڊونيشي",
			language.IG:      "اگبو",
			language.II:      "سچوان يي",
			language.ILO:     "الوڪو",
			language.INH:     "انگش",
			language.IO:      "ادو",
			language.IS:      "آئيس لينڊڪ",
			language.IT:      "اطالوي",
			language.IU:      "انو ڪتوت",
			language.JA:      "جاپاني",
			language.JBO:     "لوجبين",
			language.JGO:     "نغومبا",
			language.JMC:     "ميڪم",
			language.JV:      "جاونيز",
			language.KA:      "جارجين",
			language.KAB:     "ڪبائل",
			language.KAC:     "ڪچن",
			language.KAJ:     "پوڪيپسي",
			language.KAM:     "ڪئمبا",
			language.KBD:     "ڪبارڊيئن",
			language.KCG:     "تياپ",
			language.KDE:     "مڪوندي",
			language.KEA:     "ڪيبيو ويرڊيانو",
			language.KFO:     "ڪورو",
			language.KHA:     "خاسي",
			language.KHQ:     "ڪيورا چني",
			language.KI:      "اڪويو",
			language.KJ:      "ڪنياما",
			language.KK:      "قازق",
			language.KKJ:     "ڪڪو",
			language.KL:      "ڪالا ليسٽ",
			language.KLN:     "ڪيلين جن",
			language.KM:      "خمر",
			language.KMB:     "ڪمبونڊو",
			language.KN:      "ڪناڊا",
			language.KO:      "ڪوريائي",
			language.KOK:     "ڪونڪي",
			language.KPE:     "ڪپيل",
			language.KR:      "ڪنوري",
			language.KRC:     "ڪراچي بالڪر",
			language.KRL:     "ڪريلئين",
			language.KRU:     "ڪورخ",
			language.KS:      "ڪشميري",
			language.KSB:     "شمبالا",
			language.KSF:     "بافيا",
			language.KSH:     "ڪلونئين",
			language.KU:      "ڪردي",
			language.KUM:     "ڪومڪ",
			language.KV:      "ڪومي",
			language.KW:      "ڪورنش",
			language.KY:      "ڪرغيز",
			language.LA:      "لاطيني",
			language.LAD:     "لڊينو",
			language.LAG:     "لانگي",
			language.LB:      "لگزمبرگ",
			language.LEZ:     "ليزگهين",
			language.LG:      "گاندا",
			language.LI:      "لمبرگش",
			language.LKT:     "لڪوٽا",
			language.LN:      "لنگالا",
			language.LO:      "لائو",
			language.LOZ:     "لوزي",
			language.LRC:     "اتر لوري",
			language.LT:      "ليٿونيائي",
			language.LU:      "لوبا-ڪتانگا",
			language.LUA:     "لوبا-لولوا",
			language.LUN:     "لنڊا",
			language.LUO:     "لو",
			language.LUS:     "ميزو",
			language.LUY:     "لوهيا",
			language.LV:      "لاتوين",
			language.MAD:     "مدورائي",
			language.MAG:     "مگاهي",
			language.MAI:     "ميٿلي",
			language.MAK:     "مڪاسر",
			language.MAS:     "مسائي",
			language.MDF:     "موڪشا",
			language.MEN:     "مينڊي",
			language.MER:     "ميرو",
			language.MFE:     "موریسیین",
			language.MG:      "ملاگاسي",
			language.MGH:     "مخووا ميتو",
			language.MGO:     "ميتا",
			language.MH:      "مارشليز",
			language.MI:      "مائوري",
			language.MIC:     "ميڪ مڪ",
			language.MIN:     "مناڪابوا",
			language.MK:      "ميسي ڊونيائي",
			language.ML:      "مليالم",
			language.MN:      "منگولي",
			language.MNI:     "ماني پوري",
			language.MOH:     "موهاڪ",
			language.MOS:     "موسي",
			language.MR:      "مراٺي",
			language.MS:      "ملي",
			language.MT:      "مالٽي",
			language.MUA:     "من دانگ",
			language.MUL:     "هڪ کان وڌيڪ ٻوليون",
			language.MUS:     "ڪريڪ",
			language.MWL:     "مرانڊيز",
			language.MY:      "برمي",
			language.MYV:     "ايريزيا",
			language.MZN:     "مزيندراني",
			language.NA:      "نائو",
			language.NAP:     "نيپولٽن",
			language.NAQ:     "ناما",
			language.NB:      "نارويائي بوڪمال",
			language.ND:      "اتر دبيلي",
			language.NDS:     "لو جرمن",
			language.NE:      "نيپالي",
			language.NEW:     "نيواري",
			language.NG:      "ڊونگا",
			language.NIA:     "نياس",
			language.NIU:     "نووي",
			language.NL:      "ڊچ",
			language.NL_BE:   "فلیمش",
			language.NMG:     "ڪويسيو",
			language.NN:      "نارويائي نيوناسڪ",
			language.NNH:     "نغيمبون",
			language.NOG:     "نوگائي",
			language.NQO:     "نڪو",
			language.NR:      "ڏکڻ دبيلي",
			language.NSO:     "اتر سوٿو",
			language.NUS:     "نيور",
			language.NV:      "نواجو",
			language.NY:      "نيانجا",
			language.NYN:     "نايانڪول",
			language.OC:      "آڪسيٽن",
			language.OM:      "اورومو",
			language.OR:      "اوڊيا",
			language.OS:      "اوسيٽڪ",
			language.PA:      "پنجابي",
			language.PAG:     "پانگا سينان",
			language.PAM:     "پيم پينگا",
			language.PAP:     "پاپي امينٽو",
			language.PAU:     "پلون",
			language.PCM:     "نائيجرين پجن",
			language.PL:      "پولش",
			language.PRG:     "پرشن",
			language.PS:      "پشتو",
			language.PT:      "پورٽگليز",
			language.PT_BR:   "برازيلي پرتگالي",
			language.PT_PT:   "يورپي پرتگالي",
			language.QU:      "ڪيچوا",
			language.QUC:     "ڪچي",
			language.RAP:     "ريپنوئي",
			language.RAR:     "ريرو ٽينگو",
			language.RM:      "رومانش",
			language.RN:      "رونڊي",
			language.RO:      "روماني",
			language.RO_MD:   "مالديوي",
			language.ROF:     "رومبو",
			language.ROOT:    "روٽ",
			language.RU:      "روسي",
			language.RUP:     "ارومينين",
			language.RW:      "ڪنيار وانڊا",
			language.RWK:     "روا",
			language.SA:      "سنسڪرت",
			language.SAD:     "سنداوي",
			language.SAH:     "ساخا",
			language.SAQ:     "سيمبورو",
			language.SAT:     "سنتالي",
			language.SBA:     "نغمبي",
			language.SBP:     "سانگوو",
			language.SC:      "سارڊيني",
			language.SCN:     "سسلي",
			language.SCO:     "اسڪاٽس",
			language.SD:      "سنڌي",
			language.SE:      "اتر سامي",
			language.SEH:     "سينا",
			language.SES:     "ڪيورابورو سيني",
			language.SG:      "سانگو",
			language.SHI:     "تيچل هاتي",
			language.SHN:     "شان",
			language.SI:      "سنهالا",
			language.SK:      "سلواڪي",
			language.SL:      "سلوويني",
			language.SM:      "ساموآن",
			language.SMA:     "ڏکڻ سامي",
			language.SMJ:     "لولي سامي",
			language.SMN:     "اناري سامي",
			language.SMS:     "اسڪاٽ سامي",
			language.SN:      "شونا",
			language.SNK:     "سونينڪي",
			language.SO:      "سومالي",
			language.SQ:      "الباني",
			language.SR:      "سربيائي",
			language.SRN:     "سرانن تانگو",
			language.SS:      "سواتي",
			language.SSY:     "سهو",
			language.ST:      "ڏکڻ سوٿي",
			language.SU:      "سوڊاني",
			language.SUK:     "سڪوما",
			language.SV:      "سويڊني",
			language.SW:      "سواحيلي",
			language.SW_CD:   "ڪونگو سواحيلي",
			language.SWB:     "ڪمورين",
			language.SYR:     "شامي",
			language.TA:      "تامل",
			language.TE:      "تلگو",
			language.TEM:     "تمني",
			language.TEO:     "تيسو",
			language.TET:     "تيتم",
			language.TG:      "تاجڪي",
			language.TH:      "ٿائي",
			language.TI:      "تگرينيائي",
			language.TIG:     "تگري",
			language.TK:      "ترڪماني",
			language.TLH:     "ڪلون",
			language.TN:      "تسوانا",
			language.TO:      "تونگن",
			language.TPI:     "تاڪ پسن",
			language.TR:      "ترڪ",
			language.TRV:     "تاروڪو",
			language.TS:      "سونگا",
			language.TT:      "تاتري",
			language.TUM:     "تمبوڪا",
			language.TVL:     "توالو",
			language.TWQ:     "تساوڪي",
			language.TY:      "تاهيتي",
			language.TYV:     "تووينيائي",
			language.TZM:     "وچ اٽلس تمازائيٽ",
			language.UDM:     "ادمورتيا",
			language.UG:      "يوغور",
			language.UK:      "يوڪراني",
			language.UMB:     "اومبنڊو",
			language.UND:     "اڻڄاتل ٻولي",
			language.UR:      "اردو",
			language.UZ:      "ازبڪ",
			language.VAI:     "يا",
			language.VE:      "وينڊا",
			language.VI:      "ويتنامي",
			language.VO:      "والپڪ",
			language.VUN:     "ونجو",
			language.WA:      "ولون",
			language.WAE:     "والسر",
			language.WAL:     "وولايٽا",
			language.WAR:     "واري",
			language.WO:      "وولف",
			language.XAL:     "ڪيلمڪ",
			language.XH:      "زھوسا",
			language.XOG:     "سوگا",
			language.YAV:     "يانگ بين",
			language.YBB:     "ييمبا",
			language.YI:      "يدش",
			language.YO:      "يوروبا",
			language.YUE:     "چيني، ڪينٽونيز",
			language.ZGH:     "معياري مراڪشي تامازائيٽ",
			language.ZH:      "چيني، مندارن",
			language.ZH_HANS: "سادي مندارن چيني",
			language.ZH_HANT: "رواجي مندارن چيني",
			language.ZU:      "زولو",
			language.ZUN:     "زوني",
			language.ZXX:     "ڪوئي ٻولي جو مواد ڪونهي",
			language.ZZA:     "زازا",
		},
		Territories: cldr.Territories{
			territory.V_001: "دنيا",
			territory.V_002: "آفريڪا",
			territory.V_003: "اتر آمريڪا",
			territory.V_005: "ڏکڻ آمريڪا",
			territory.V_009: "سامونڊي",
			territory.V_011: "اولهه آفريقا",
			territory.V_013: "وچ آمريڪا",
			territory.V_014: "اوڀر آفريڪا",
			territory.V_015: "اترين آفريڪا",
			territory.V_017: "وچ آفريڪا",
			territory.V_018: "ڏاکڻي آمريڪا",
			territory.V_019: "آمريڪا",
			territory.V_021: "اترين آمريڪا",
			territory.V_029: "ڪيريبين",
			territory.V_030: "اوڀر ايشيا",
			territory.V_034: "ڏکڻ ايشيا",
			territory.V_035: "ڏکڻ اوڀر ايشيا",
			territory.V_039: "ڏکڻ يورپ",
			territory.V_053: "آسٽریلیشیا",
			territory.V_054: "میلانیشیا",
			territory.V_057: "مائڪرونيشائي خطو",
			territory.V_061: "پولینیشیا",
			territory.V_142: "ايشيا",
			territory.V_143: "وچ ايشيا",
			territory.V_145: "اولهه ايشيا",
			territory.V_150: "يورپ",
			territory.V_151: "اوڀر يورپ",
			territory.V_154: "اترين يورپ",
			territory.V_155: "اولهه يورپ",
			territory.V_202: "سب-سهارا آفريڪا",
			territory.V_419: "لاطيني آمريڪا",
			territory.AC:    "طلوع ٻيٽ",
			territory.AD:    "اندورا",
			territory.AE:    "متحده عرب امارات",
			territory.AF:    "افغانستان",
			territory.AG:    "انٽيگئا و بربودا",
			territory.AI:    "انگويلا",
			territory.AL:    "البانيا",
			territory.AM:    "ارمینیا",
			territory.AO:    "انگولا",
			territory.AQ:    "انٽارڪٽيڪا",
			territory.AR:    "ارجنٽينا",
			territory.AS:    "آمريڪي ساموا",
			territory.AT:    "آشٽريا",
			territory.AU:    "آسٽريليا",
			territory.AW:    "عروبا",
			territory.AX:    "الند ٻيٽ",
			territory.AZ:    "آذربائيجان",
			territory.BA:    "بوسنیا اور هرزیگوینا",
			territory.BB:    "باربڊوس",
			territory.BD:    "بنگلاديش",
			territory.BE:    "بيلجيم",
			territory.BF:    "برڪينا فاسو",
			territory.BG:    "بلغاريا",
			territory.BH:    "بحرين",
			territory.BI:    "برونڊي",
			territory.BJ:    "بينن",
			territory.BL:    "سینٽ برٿلیمی",
			territory.BM:    "برمودا",
			territory.BN:    "برونائي",
			territory.BO:    "بوليويا",
			territory.BQ:    "ڪيريبين نيدرلينڊ",
			territory.BR:    "برازيل",
			territory.BS:    "بهاماس",
			territory.BT:    "ڀوٽان",
			territory.BV:    "بووٽ ٻيٽ",
			territory.BW:    "بوٽسوانا",
			territory.BY:    "بیلارس",
			territory.BZ:    "بيليز",
			territory.CA:    "ڪئناڊا",
			territory.CC:    "ڪوڪوس ٻيٽ",
			territory.CD:    "ڪانگو",
			territory.CF:    "وچ آفريقي جمهوريه",
			territory.CG:    "ڪانگو (جمهوري)",
			territory.CH:    "سوئزرلينڊ",
			territory.CI:    "آئيوري ڪوسٽ",
			territory.CK:    "ڪوڪ ٻيٽ",
			territory.CL:    "چلي",
			territory.CM:    "ڪيمرون",
			territory.CN:    "چين",
			territory.CO:    "ڪولمبيا",
			territory.CP:    "ڪلپرٽن ٻيٽ",
			territory.CR:    "ڪوسٽا رڪا",
			territory.CU:    "ڪيوبا",
			territory.CV:    "ڪيپ وردي",
			territory.CW:    "ڪيوراسائو",
			territory.CX:    "ڪرسمس ٻيٽ",
			territory.CY:    "سائپرس",
			territory.CZ:    "چيڪ جهموريو",
			territory.DE:    "جرمني",
			territory.DG:    "ڊئيگو گارسيا",
			territory.DJ:    "ڊجبيوتي",
			territory.DK:    "ڊينمارڪ",
			territory.DM:    "ڊومينيڪا",
			territory.DO:    "ڊومينيڪن جمهوريه",
			territory.DZ:    "الجيريا",
			territory.EA:    "سیوٽا ۽ میلیلا",
			territory.EC:    "ايڪواڊور",
			territory.EE:    "ايسٽونيا",
			territory.EG:    "مصر",
			territory.EH:    "اولهه صحارا",
			territory.ER:    "ايريٽيريا",
			territory.ES:    "اسپين",
			territory.ET:    "ايٿوپيا",
			territory.EU:    "يورپين يونين",
			territory.EZ:    "يورو زون",
			territory.FI:    "فن لينڊ",
			territory.FJ:    "فجي",
			territory.FK:    "فلڪ لينڊ ٻيٽ",
			territory.FM:    "مائڪرونيشيا",
			territory.FO:    "فارو ٻيٽ",
			territory.FR:    "فرانس",
			territory.GA:    "گبون",
			territory.GB:    "برطانيه",
			territory.GD:    "گرينڊا",
			territory.GE:    "جارجيا",
			territory.GF:    "فرانسيسي گيانا",
			territory.GG:    "گورنسي",
			territory.GH:    "گهانا",
			territory.GI:    "جبرالٽر",
			territory.GL:    "گرين لينڊ",
			territory.GM:    "گيمبيا",
			territory.GN:    "گني",
			territory.GP:    "گواڊیلوپ",
			territory.GQ:    "ايڪوٽوريل گائينا",
			territory.GR:    "يونان",
			territory.GS:    "ڏکڻ جارجيا ۽ ڏکڻ سينڊوچ ٻيٽ",
			territory.GT:    "گوئٽي مالا",
			territory.GU:    "گوام",
			territory.GW:    "گني بسائو",
			territory.GY:    "گيانا",
			territory.HK:    "هانگ ڪانگ",
			territory.HM:    "هرڊ ۽ مڪڊونلڊ ٻيٽ",
			territory.HN:    "هنڊورس",
			territory.HR:    "ڪروئيشيا",
			territory.HT:    "هيٽي",
			territory.HU:    "چيڪ جهموريه",
			territory.IC:    "ڪينري ٻيٽ",
			territory.ID:    "انڊونيشيا",
			territory.IE:    "آئرلينڊ",
			territory.IL:    "اسرائيل",
			territory.IM:    "انسانن جو ٻيٽ",
			territory.IN:    "انڊيا",
			territory.IO:    "برطانوي هندي سمنڊ خطو",
			territory.IQ:    "عراق",
			territory.IR:    "ايران",
			territory.IS:    "آئس لينڊ",
			territory.IT:    "اٽلي",
			territory.JE:    "جرسي",
			territory.JM:    "جميڪا",
			territory.JO:    "اردن",
			territory.JP:    "جاپان",
			territory.KE:    "ڪينيا",
			territory.KG:    "ڪرغستان",
			territory.KH:    "ڪمبوڊيا",
			territory.KI:    "ڪرباتي",
			territory.KM:    "ڪوموروس",
			territory.KN:    "سينٽ ڪٽس و نيوس",
			territory.KP:    "اتر ڪوريا",
			territory.KR:    "ڏکڻ ڪوريا",
			territory.KW:    "ڪويت",
			territory.KY:    "ڪي مين ٻيٽ",
			territory.KZ:    "قازقستان",
			territory.LA:    "لائوس",
			territory.LB:    "لبنان",
			territory.LC:    "سينٽ لوسيا",
			territory.LI:    "لچي ٽينسٽين",
			territory.LK:    "سري لنڪا",
			territory.LR:    "لائبیریا",
			territory.LS:    "ليسوٿو",
			territory.LT:    "لٿونيا",
			territory.LU:    "لیگزمبرگ",
			territory.LV:    "لاتويا",
			territory.LY:    "لبيا",
			territory.MA:    "موروڪو",
			territory.MC:    "موناڪو",
			territory.MD:    "مالدووا",
			territory.ME:    "مونٽي نيگرو",
			territory.MF:    "سينٽ مارٽن",
			territory.MG:    "مداگيسڪر",
			territory.MH:    "مارشل ڀيٽ",
			territory.MK:    "شمالي مقدونيا",
			territory.ML:    "مالي",
			territory.MM:    "ميانمار (برما)",
			territory.MN:    "منگوليا",
			territory.MO:    "مڪائو",
			territory.MP:    "اتر مرينا ٻيٽ",
			territory.MQ:    "مارتينڪ",
			territory.MR:    "موريتانيا",
			territory.MS:    "مونٽسراٽ",
			territory.MT:    "مالٽا",
			territory.MU:    "موريشس",
			territory.MV:    "مالديپ",
			territory.MW:    "مالاوي",
			territory.MX:    "ميڪسيڪو",
			territory.MY:    "ملائيشيا",
			territory.MZ:    "موزمبیق",
			territory.NA:    "نيميبيا",
			territory.NC:    "نیو ڪالیڊونیا",
			territory.NE:    "نائيجر",
			territory.NF:    "نورفوڪ ٻيٽ",
			territory.NG:    "نائيجيريا",
			territory.NI:    "نڪراگوا",
			territory.NL:    "نيدرلينڊ",
			territory.NO:    "ناروي",
			territory.NP:    "نيپال",
			territory.NR:    "نائورو",
			territory.NU:    "نووي",
			territory.NZ:    "نيو زيلينڊ",
			territory.OM:    "عمان",
			territory.PA:    "پناما",
			territory.PE:    "پيرو",
			territory.PF:    "فرانسيسي پولينيشيا",
			territory.PG:    "پاپوا نیو گني",
			territory.PH:    "فلپائن",
			territory.PK:    "پاڪستان",
			territory.PL:    "پولينڊ",
			territory.PM:    "سینٽ پیئر و میڪوئیلون",
			territory.PN:    "پٽڪئرن ٻيٽ",
			territory.PR:    "پيوئرٽو ريڪو",
			territory.PS:    "فلسطين",
			territory.PT:    "پرتگال",
			territory.PW:    "پلائو",
			territory.PY:    "پيراگوءِ",
			territory.QA:    "قطر",
			territory.QO:    "بيروني سامونڊي",
			territory.RE:    "ري يونين",
			territory.RO:    "رومانيا",
			territory.RS:    "سربيا",
			territory.RU:    "روس",
			territory.RW:    "روانڊا",
			territory.SA:    "سعودي عرب",
			territory.SB:    "سولومون ٻيٽَ",
			territory.SC:    "شي شلز",
			territory.SD:    "سوڊان",
			territory.SE:    "سوئيڊن",
			territory.SG:    "سينگاپور",
			territory.SH:    "سينٽ ھيلينا",
			territory.SI:    "سلوینیا",
			territory.SJ:    "سوالبارڊ ۽ جان ماین",
			territory.SK:    "سلوواڪيا",
			territory.SL:    "سيرا ليون",
			territory.SM:    "سین مرینو",
			territory.SN:    "سينيگال",
			territory.SO:    "سوماليا",
			territory.SR:    "سورينام",
			territory.SS:    "ڏکڻ سوڊان",
			territory.ST:    "سائو ٽوم ۽ پرنسپیي",
			territory.SV:    "ال سلواڊور",
			territory.SX:    "سنٽ مارٽن",
			territory.SY:    "شام",
			territory.SZ:    "سوازيلينڊ",
			territory.TA:    "ٽرسٽن دا ڪوها",
			territory.TC:    "ترڪ ۽ ڪيڪوس ٻيٽ",
			territory.TD:    "چاڊ",
			territory.TF:    "فرانسيسي ڏاکڻي علائقا",
			territory.TG:    "توگو",
			territory.TH:    "ٿائيليند",
			territory.TJ:    "تاجڪستان",
			territory.TK:    "ٽوڪلائو",
			territory.TL:    "اوڀر تيمور",
			territory.TM:    "ترڪمانستان",
			territory.TN:    "تيونيسيا",
			territory.TO:    "ٽونگا",
			territory.TR:    "ترڪي",
			territory.TT:    "ٽريني ڊيڊ ۽ ٽوباگو ٻيٽ",
			territory.TV:    "توالو",
			territory.TW:    "تائیوان",
			territory.TZ:    "تنزانيا",
			territory.UA:    "يوڪرين",
			territory.UG:    "يوگنڊا",
			territory.UM:    "آمريڪي ٻاهريون ٻيٽ",
			territory.UN:    "اقوام متحده",
			territory.US:    "يوايس",
			territory.UY:    "يوروگوءِ",
			territory.UZ:    "ازبڪستان",
			territory.VA:    "ويٽڪين سٽي",
			territory.VC:    "سینٽ ونسنت ۽ گریناڊینز",
			territory.VE:    "وينزيلا",
			territory.VG:    "برطانوي ورجن ٻيٽ",
			territory.VI:    "آمريڪي ورجن ٻيٽ",
			territory.VN:    "ويتنام",
			territory.VU:    "وينيٽيو",
			territory.WF:    "والس ۽ فتونا",
			territory.WS:    "سموئا",
			territory.XA:    "سوڊو-لهجا",
			territory.XB:    "سوڊو-بي ڊي",
			territory.XK:    "ڪوسووو",
			territory.YE:    "يمن",
			territory.YT:    "مياتي",
			territory.ZA:    "ڏکڻ آفريقا",
			territory.ZM:    "زيمبيا",
			territory.ZW:    "زمبابوي",
			territory.ZZ:    "اڻڄاتل خطو",
		},
	}
}
