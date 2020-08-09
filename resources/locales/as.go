// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_as() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "dd-MM-y", Short: "d-M-y"}, Time: cldr.CalendarDateFormat{Full: "a h.mm.ss zzzz", Long: "a h.mm.ss z", Medium: "a h.mm.ss", Short: "a h.mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "জানু", Feb: "ফেব্ৰু", Mar: "মাৰ্চ", Apr: "এপ্ৰিল", May: "মে’", Jun: "জুন", Jul: "জুলাই", Aug: "আগ", Sep: "ছেপ্তে", Oct: "অক্টো", Nov: "নৱে", Dec: "ডিচে"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "জ", Feb: "ফ", Mar: "ম", Apr: "এ", May: "ম", Jun: "জ", Jul: "জ", Aug: "আ", Sep: "ছ", Oct: "অ", Nov: "ন", Dec: "ড"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "জানুৱাৰী", Feb: "ফেব্ৰুৱাৰী", Mar: "মাৰ্চ", Apr: "এপ্ৰিল", May: "মে’", Jun: "জুন", Jul: "জুলাই", Aug: "আগষ্ট", Sep: "ছেপ্তেম্বৰ", Oct: "অক্টোবৰ", Nov: "নৱেম্বৰ", Dec: "ডিচেম্বৰ"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "দেও", Mon: "সোম", Tue: "মঙ্গল", Wed: "বুধ", Thu: "বৃহ", Fri: "শুক্ৰ", Sat: "শনি"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "দ", Mon: "স", Tue: "ম", Wed: "ব", Thu: "ব", Fri: "শ", Sat: "শ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "দেও", Mon: "সোম", Tue: "মঙ্গল", Wed: "বুধ", Thu: "বৃহ", Fri: "শুক্ৰ", Sat: "শনি"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "দেওবাৰ", Mon: "সোমবাৰ", Tue: "মঙ্গলবাৰ", Wed: "বুধবাৰ", Thu: "বৃহস্পতিবাৰ", Fri: "শুক্ৰবাৰ", Sat: "শনিবাৰ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "পূৰ্বাহ্ন", PM: "অপৰাহ্ন"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "পূৰ্বাহ্ন", PM: "অপৰাহ্ন"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "পূৰ্বাহ্ন", PM: "অপৰাহ্ন"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_as]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##,##0.###", Currency: "¤\u00a0#,##,##0.00", Percent: "#,##,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "সংযুক্ত আৰব আমিৰাত ডিৰহেম", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "আফগান আফগানী", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "আলবেনীয় লেক", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "আৰ্মেনিয়ান ড্ৰাম", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "নেডাৰলেণ্ডছ এণ্টিলিয়েন গিল্ডাৰ", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "এংগোলান কোৱাঞ্জা", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "আৰ্জেণ্টাইন পেছো", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "অষ্ট্ৰেলিয়ান ডলাৰ", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "আৰুবান ফ্ল’ৰিন", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "আজেৰবাইজানী মানাত", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "ব’ছনিয়া আৰু হাৰ্জেগ’ভিনা কনভাৰ্টিব্\u200cল মাৰ্ক", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "বাৰ্বাডিয়ান ডলাৰ", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "বাংলাদেশী টাকা", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "বুলগেৰীয় লেভ", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "বাহৰেইনী ডিনাৰ", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "বুৰুণ্ডিয়ান ফ্ৰেংক", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "বাৰ্মুডান ডলাৰ", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "ব্ৰুনেই ডলাৰ", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "বলিভিয়ান বলিভিয়ানো", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "ব্ৰাজিলিয়ান ৰিয়েল", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "বাহামিয়ান ডলাৰ", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ভুটানী নংগলট্ৰাম", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "ব’টচোৱানান পুলা", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "বেলাৰুছীয় ৰুবেল", Symbol: "BYN"},
				currency.BZD: cldr.Currency{DisplayName: "বেলিজ ডলাৰ", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "কানাডিয়ান ডলাৰ", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "কংগো ফ্ৰেংক", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "চুইছ ফ্ৰেংক", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "চিলিয়ান পেছো", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "চীনা ইউৱান (অফশ্ব’ৰ)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "চীনা ইউৱান", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "কলম্বিয়ান পেছো", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "কোষ্টা ৰিকান কোলন", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "কিউবান ৰূপান্তৰযোগ্য পেছো", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "কিউবান পেছো", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "কেপ ভাৰ্দে এছকুডো", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "চেক কোৰুনা", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "জিবুটি ফ্ৰেংক", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "ডেনিচ ক্ৰোন", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "ড’মিনিকান পেছো", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "আলজেৰীয় ডিনাৰ", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "ইজিপ্তৰ পাউণ্ড", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "এৰিট্ৰিয়ন নাক্\u200cফা", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ইথিঅ’পিয়ান বিৰ", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "ইউৰো", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "ফিজিয়ান ডলাৰ", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "ফকলেণ্ড দ্বীপপুঞ্জৰ পাউণ্ড", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "ব্ৰিটিছ পাউণ্ড", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "জৰ্জিয়ান লাৰি", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "ঘানাৰ চেডি", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "জিব্ৰাল্টৰ পাউণ্ড", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "গাম্বিয়া ডালাছি", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "গিনি ফ্ৰেংক", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "গুৱাটেমালা কুৱেৎজাল", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "গায়ানিজ ডলাৰ", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "হং কং ডলাৰ", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "হোন্দুৰান লেম্পিৰা", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "ক্ৰোৱেছিয়ান কুনা", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "হাইটিয়ান গৌৰ্ড", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "হাংগেৰীয়ান ফ’ৰিণ্ট", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "ইণ্ডোনেচিয়ান ৰুপিয়াহ", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "ইজৰাইলী নিউ শ্বেকেল", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "ভাৰতীয় ৰুপী", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "ইৰাকী ডিনাৰ", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ইৰানীয়ান ৰিয়েল", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "আইচলেণ্ডিক ক্ৰোনা", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "জামাইকান ডলাৰ", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "জৰ্ডানিয়ান ডিনাৰ", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "জাপানী য়েন", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "কেনিয়ান শ্বিলিং", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "কিৰ্গিস্তানী ছোম", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "কেম্ব’ডিয়ান ৰিয়েল", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "ক’মোৰিয়ান ফ্ৰেংক", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "উত্তৰ কোৰিয়াৰ ওৱান", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "দক্ষিণ কোৰিয়াৰ ওৱান", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "কুৱেইটি ডিনাৰ", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "কেইমেন দ্বীপপুঞ্জৰ ডলাৰ", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "কাজাখস্তানী তেঞ্জ", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "লাওচিয়ান কিপ", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "লেবানীজ পাউণ্ড", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "শ্ৰীলংকান ৰুপী", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "লাইবেৰিয়ান ডলাৰ", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "লিবিয়ান ডিনাৰ", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "মৰোক্কান ডিৰহাম", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "মোলডোভান লেউ", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "মালাগাছী এৰিয়াৰী", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "মেচিডোনীয় ডেনাৰ", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "ম্যানমাৰ কিয়াট", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "মঙ্গোলিয়ান টুৰ্গিক", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "মেকানীজ পাটাকা", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "ম’ৰিটেনিয়ান ঔগুইয়া (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "ম’ৰিটেনিয়ান ঔগুইয়া", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "মৰিচিয়ান ৰুপী", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "মালডিভিয়ান ৰুফিয়া", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "মালাউইয়ান কোৱাচা", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "মেক্সিকান পেছো", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "মালায়েচিয়ান ৰিংগিট", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "মোজাম্বিকান মেটিকল", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "নামিবিয়ান ডলাৰ", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "নাইজেৰিয়ান নাইৰা", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "নিকাৰাগুৱান কোৰ্ডোবা", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "নৰৱেজিয়ান ক্ৰোন", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "নেপালী ৰুপী", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "নিউজিলেণ্ড ডলাৰ", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ওমানি ৰিয়েল", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "পানামেনিয়ান বাল্বোৱা", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "পেৰুভিয়ান ছ’ল", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "পাপুৱা নিউ গিনি কিনা", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "ফিলিপিন পেইছ’", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "পাকিস্তানী ৰুপী", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "প’লিচ জোল্টী", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "পাৰাগুয়ান গুৱাৰানি", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "কাটাৰি ৰিয়েল", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "ৰোমানীয় লেউ", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "চাৰ্বিয়ান ডিনাৰ", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "ৰাছিয়ান ৰুব্\u200cল", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ৰোৱান্দান ফ্ৰেংক", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "চৌডি ৰিয়েল", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "চোলোমোন দ্বীপপুঞ্জৰ ডলাৰ", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "ছেচেলৱা ৰুপী", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "চুডানী পাউণ্ড", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "চুইডিছ ক্ৰোনা", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "ছিংগাপুৰ ডলাৰ", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "ছেইণ্ট হেলেনা পাউণ্ড", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "চিয়েৰা লিঅ’নৰ লিঅ’ন", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "চোমালি শ্বিলিং", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "ছুৰিনামী ডলাৰ", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "দক্ষিণ চুডানীজ পাউণ্ড", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "চাও টোমে আৰু প্ৰিনচিপে ডোব্\u200cৰা (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "চাও টোমে আৰু প্ৰিনচিপে ডোব্\u200cৰা", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "চিৰিয়ান পাউণ্ড", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "স্বাজি লিলেংগেনি", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "থাই বাত", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "তাজিকিস্তানী ছোমনি", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "তুৰ্কমেনিস্তানী মানাত", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "টুনিচিয়ান ডিনাৰ", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "টংগান পাআংগা", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "তুৰ্কীৰ লিৰা", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ট্ৰিনিডাড আৰু টোবাগো ডলাৰ", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "নিউ টাইৱান ডলাৰ", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "টানজানিয়ান শ্বিলিং", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ইউক্ৰেইনীয় হৃভনিয়া", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "উগাণ্ডান শ্বিলিং", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "ইউ. এছ. ডলাৰ", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "উৰুগুৱেয়ান পেছো", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "উজবেকিস্তানী ছোম", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "ভেনিজুৱেলান বলিভাৰ (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "ভেনিজুৱেলান বলিভাৰ (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "ভেনিজুৱেলান বলিভাৰ", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "ভিয়েটনামীজ ডং", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "ভানাটুৰ ভাটু", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "ছামোৱান টালা", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "মধ্য আফ্ৰিকান CFA ফ্ৰেংক", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "ইষ্ট কেৰিবিয়ান ডলাৰ", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "পশ্চিম আফ্ৰিকান CFA ফ্ৰেংক", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP ফ্ৰেংক", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "অজ্ঞাত মুদ্ৰা", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "য়েমেনী ৰিয়েল", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "দক্ষিণ আফ্ৰিকাৰ ৰাণ্ড", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "জাম্বিয়ান কোৱাচা", Symbol: "ZMW"},
			},
		},
	}
}
