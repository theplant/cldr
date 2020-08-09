// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_bn() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT {0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "জানুয়ারী", Feb: "ফেব্রুয়ারী", Mar: "মার্চ", Apr: "এপ্রিল", May: "মে", Jun: "জুন", Jul: "জুলাই", Aug: "আগস্ট", Sep: "সেপ্টেম্বর", Oct: "অক্টোবর", Nov: "নভেম্বর", Dec: "ডিসেম্বর"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "জা", Feb: "ফে", Mar: "মা", Apr: "এ", May: "মে", Jun: "জুন", Jul: "জু", Aug: "আ", Sep: "সে", Oct: "অ", Nov: "ন", Dec: "ডি"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "জানুয়ারী", Feb: "ফেব্রুয়ারী", Mar: "মার্চ", Apr: "এপ্রিল", May: "মে", Jun: "জুন", Jul: "জুলাই", Aug: "আগস্ট", Sep: "সেপ্টেম্বর", Oct: "অক্টোবর", Nov: "নভেম্বর", Dec: "ডিসেম্বর"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "রবি", Mon: "সোম", Tue: "মঙ্গল", Wed: "বুধ", Thu: "বৃহস্পতি", Fri: "শুক্র", Sat: "শনি"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "র", Mon: "সো", Tue: "ম", Wed: "বু", Thu: "বৃ", Fri: "শু", Sat: "শ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "রঃ", Mon: "সোঃ", Tue: "মঃ", Wed: "বুঃ", Thu: "বৃঃ", Fri: "শুঃ", Sat: "শনি"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "রবিবার", Mon: "সোমবার", Tue: "মঙ্গলবার", Wed: "বুধবার", Thu: "বৃহস্পতিবার", Fri: "শুক্রবার", Sat: "শনিবার"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_bn]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##,##0.###", Currency: "#,##,##0.00¤", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "এ্যান্ডোরান পেসেতা", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "সংযুক্ত আরব আমিরাত দিরহাম", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "আফগানি (১৯২৭–২০০২)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "আফগান আফগানি", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "আলবেনিয়ান লেক", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "আরমেনিয়ান দ্রাম", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "নেদারল্যান্ড এ্যান্টিলিয়ান গুল্ডের", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "এ্যাঙ্গোলান কওয়ানজা", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "এ্যাঙ্গোলান কওয়ানজা (১৯৭৭–১৯৯০)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "এ্যাঙ্গোলান নতুন কওয়ানজা (১৯৯৫–২০০০)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "এ্যাঙ্গোলান কওয়ানজা (১৯৯৫–১৯৯৯)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "আর্জেন্টিনা অস্ট্রাল", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "আর্জেন্টিনা পেসো (১৯৮৩–১৯৮৫)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "আর্জেন্টিনা পেসো", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "অস্ট্রিয়ান শিলিং", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "অস্ট্রেলিয়ান ডলার", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "আরুবা গিল্ডার", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "আজারবাইজান মানাত (১৯৯৩–২০০৬)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "আজারবাইজান মানাত", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "বসনিয়া এবং হার্জেগোভিনা দিনার", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "বসনিয়া এবং হার্জেগোভিনা বিনিমেয় মার্ক", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "বার্বেডোজ ডলার", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "বাংলাদেশী টাকা", Symbol: "৳"},
				currency.BEC: cldr.Currency{DisplayName: "বেলজিয়ান ফ্রাঙ্ক (রূপান্তরযোগ্য)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "বেলজিয়ান ফ্রাঙ্ক", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "বেলজিয়ান ফ্রাঙ্ক (আর্থিক)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "বুলগেরীয় হার্ড লেভ", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "বুলগেরীয় লেভ", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "বাহরাইনি দিনার", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "বুরুন্ডি ফ্রাঙ্ক", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "বারমিউডান ডলার", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "ব্রুনেই ডলার", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "বলিভিয়ানো", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "বলিভিয়ান পেসো", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "বলিভিয়ান মভডোল", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "ব্রাজিলিয়ান ক্রুজেয়রোনোভো (১৯৬৭–১৯৮৬)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "ব্রাজিলিয়ান ক্রুজেইডাউ", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "ব্রাজিলিয়ান ক্রুজেয়রো (১৯৯০–১৯৯৩)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "ব্রাজিলিয়ান রিয়েল", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "ব্রাজিলিয়ান ক্রুজেইডো নোভো", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "ব্রাজিলিয়ান ক্রুজেয়রো", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "বাহামিয়ান ডলার", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ভুটানি এনগুল্ট্রুম", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "বর্মি কিয়াৎ", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "বতসোয়ানা পুলা", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "বেলারুশিয়ান নিউ রুবেল (১৯৯৪–১৯৯৯)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "বেলারুশিয়ান রুবেল", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "বেলারুশিয়ান রুবেল (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "বেলিজ ডলার", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "কানাডিয়ান ডলার", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "কঙ্গোলিস ফ্র্যাঙ্ক", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "সুইজারল্যান্ড ইউরো", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "সুইস ফ্রাঁ", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "সুইজারল্যান্ড ফ্রাঙ্ক", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "চিলিয়ান উনিদাদেস দি ফোমেন্তো", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "চিলি পেসো", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "চাইনিজ ইউয়ান (অফশোর)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "চীনা য়ুয়ান", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "কলোম্বিয়ান পেসো", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "উনিদাদ দি ভ্যালোর রিয়েল", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "কোস্টা রিকা কোলোন", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "প্রাচীন সারবিয়ান দিনার", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "চেকোস্লোভাক হার্ড কোরুনা", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "কিউবান রূপান্তরযোগ্য পেসো", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "কিউবান পেসো", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "কেপ ভার্দে এসকুডো", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "সাইপ্রাস পাউন্ড", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "চেক প্রজাতন্ত্র কোরুনা", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "পূর্ব জার্মান মার্ক", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "ডয়চ্ মার্ক", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "জিবুতি ফ্রাঙ্ক", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "ড্যানিশ ক্রোন", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "ডোমিনিকান পেসো", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "আলজেরীয় দিনার", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "ইকুয়াডোর সুক্রে", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "ইকুয়াডোর উনিদাদেস দি ভেলর কনসতান্তে (ইউভিসি)", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "এস্তোনিয়া ক্রুনি", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "মিশরীয় পাউন্ড", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "এরিট্রিয়েন নাকফা", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "স্প্যানিশ পেসেতা (একই হিসাব)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "স্প্যানিশ পেসেতা (রূপান্তরযোগ্য হিসাব)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "স্প্যানিশ পেসেতা", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ইথিওপিয়ান বির", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "ইউরো", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "ফিনিস মার্কা", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "ফিজি ডলার", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "ফকল্যান্ড দ্বীপপুঞ্জ পাউন্ড", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "ফরাসি ফ্রাঙ্ক", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "ব্রিটিশ পাউন্ড", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "জর্জিয়ান কুপন লারিট", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "জর্জিয়ান লারি", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "ঘানা সেডি (১৯৭৯–২০০৭)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "ঘানা সেডি", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "জিব্রাল্টার পাউন্ড", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "গাম্বিয়া ডালাসি", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "গিনি ফ্রাঙ্ক", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "গিনি সাইলি", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "ইকুয়েটোরিয়াল গিনি ইকুয়িলি", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "গ্রীক দ্রাচমা", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "গুয়াতেমালা কুয়েৎজাল", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "পর্তুগিজ গিনি এসকুডো", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "গিনি বিসাউ পেসো", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "গাইয়েনা ডলার", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "হংকং ডলার", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "হন্ডুরাস লেম্পিরা", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "ক্রোয়েশিয়ান দিনার", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "ক্রোয়েশিয়ান কুনা", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "হাইতি গৌর্দে", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "হাঙ্গেরিয়ান ফোরিন্ট", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "ইন্দোনেশিয়ান রুপিয়াহ", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "ইরিশ পাউন্ড", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "ইস্রাইলি পাউন্ড", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "ইস্রাইলি নতুন শেকেল", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "ভারতীয় রুপি", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "ইরাকি দিনার", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ইরানিয়ান রিয়াল", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "আইসল্যান্ডীয় ক্রোনা", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "ইতালীয় লিরা", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "জামাইকান ডলার", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "জর্ডানিয়ান দিনার", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "জাপানি ইয়েন", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "কেনিয়ান শিলিং", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "কিরগিজস্তান সোম", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "কম্বোডিয়ান রিয়েল", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "কম্বোরো ফ্রাঙ্ক", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "উত্তর কোরিয়ার ওন", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "দক্ষিণ কোরিয়ান ওন", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "কুয়েতি দিনার", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "কেম্যান দ্বীপপুঞ্জের ডলার", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "কাজাখাস্তানি টেঙ্গে", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "লেউশান কিপ", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "লেবানিজ পাউন্ড", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "শ্রীলঙ্কান রুপি", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "লিবেরিয়ান ডলার", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "লেসুটু লোটি", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "লিথুইনিয়ান লিটা", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "লিথুইনিয়ান টালোন্যাস", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "লুক্সেমবার্গ রুপান্তযোগ্য ফ্রাঙ্ক", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "লুক্সেমবার্গ ফ্রাঙ্ক", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "লুক্সেমবার্গ ফাইনেনশিয়াল ফ্রাঙ্ক", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "ল্যাটভিয়ান ল্যাট্\u200cস", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "ল্যাটভিয়ান রুবল", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "লিবিয়ান দিনার", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "মোরোক্কান দিরহাম", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "মোরোক্কান ফ্রাঙ্ক", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "মোল্ডোভান লেয়ু", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "মাদাগাস্কার আরিয়ারি", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "মাদাগাস্কার ফ্রাঙ্ক", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "ম্যাসেডোনিয়ান দিনার", Symbol: "MKD"},
				currency.MLF: cldr.Currency{DisplayName: "মালি ফ্রাঙ্ক", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "মায়ানমার কিয়াত", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "মঙ্গোলিয়ান তুগরিক", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "ম্যাক্যাও পাটাকা", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "মৌরিতানিয়ান ওউগুইয়া (১৯৭৩–২০১৭)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "মৌরিতানিয়ান ওউগুইয়া", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "মাল্টা লিরা", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "মাল্টা পাউন্ড", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "মৌরিতানিয়ান রুপি", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "মালদিভিয়ান রুফিয়া", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "মালাউইয়ান কওয়াচ", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "ম্যাক্সিকান পেসো", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "ম্যাক্সিকান সিলভার পেসো (১৮৬১–১৯৯২)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "মেক্সিকান উনিদাদ দি ইনভার্সান (UDI)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "মালয়েশিয়ান রিঙ্গিৎ", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "মোজাম্বিক এসকুডো", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "প্রাচীন মোজাম্বিক মেটিকেল", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "মোজাম্বিক মেটিকেল", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "নামিবিয়া ডলার", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "নাইজেরিয়ান নায়রা", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "নিকারাগুয়ান কর্ডোবা (১৯৮৮–১৯৯১)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "নিকারাগুয়ান কর্ডোবা", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "নেদারল্যান্ড গুল্ডের", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "নরওয়েজিয়ান ক্রোন", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "নেপালি রুপি", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "নিউজিল্যান্ড ডলার", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ওমানি রিয়াল", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "পানামা বেলবোয়া", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "পেরুভিয়ান ইন্তি", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "পেরুভিয়ান সোল", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "পেরুভিয়ান সোল (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "পাপুয়া নিউ গিনিয়ান কিনা", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "ফিলিপাইন পেসো", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "পাকিস্তানি রুপি", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "পোলিশ জ্লোটি", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "পোলিশ জ্লোটি (১৯৫০–১৯৯৫)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "পর্তুগিজ এসকুডো", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "প্যারাগুয়ান গুয়ারানি", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "কাতার রিয়াল", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "রোডেশিয়ান ডলার", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "প্রাচীন রুমানিয়া লেয়ু", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "রুমানিয়া লেয়ু", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "সারবিয়ান দিনার", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "রাশিয়ান রুবেল", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "রাশিয়ান রুবল (১৯৯১–১৯৯৮)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "রুয়ান্ডান ফ্রাঙ্ক", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "সৌদি রিয়াল", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "সলোমন দ্বীপপুঞ্জ ডলার", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "সেয়চেল্লোইস রুপি", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "প্রাচীন সুদানি দিনার", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "সুদানি পাউন্ড", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "প্রাচীন সুদানি পাউন্ড", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "সুইডিশ ক্রোনা", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "সিঙ্গাপুর ডলার", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "সেন্ট হেলেনা পাউন্ড", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "স্লোভানিয়া টোলার", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "স্লোভাক কোরুনা", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "সিয়েরালিয়ন লিয়ন", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "সোমালি শিলিং", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "সুরিনাম ডলার", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "সুরিনাম গিল্ডার", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "দক্ষিণ সুদানি পাউন্ড", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "সাও টোমে এবং প্রিন্সিপে ডোবরা (১৯৭৭–২০১৭)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "সাও টোমে এবং প্রিন্সিপে ডোবরা", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "সোভিয়েত রুবল", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "এল স্যালভোডোর কোলোন", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "সিরিয়ান পাউন্ড", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "সোয়াজিল্যান্ড লিলাঙ্গেনি", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "থাই বাত", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "তাজিকিস্তান রুবল", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "তাজিকিস্তান সোমোনি", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "তুর্কমেনিস্টানি মানাত", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "তুর্কমেনিস্তান মানত", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "তিউনেশিয়ান দিনার", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "টোঙ্গা পা’আঙ্গা", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "তিমুর এসকুডো", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "প্রাচীন তুর্কি লিরা", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "তুর্কি লিরা", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ত্রিনিদাদ এবং টোবাগো ডলার", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "নতুন তাইওয়ান ডলার", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "তাঞ্জনিয়া শিলিং", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ইউক্রেইন হৃভনিয়া", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "ইউক্রেইন কার্বোভ্যান্টস", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "উগান্ডান শিলিং (১৯৬৬–১৯৮৭)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "উগান্ডান শিলিং", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "মার্কিন ডলার", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "মার্কিন ডলার (পরবর্তী দিন)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "মার্কিন ডলার (একই দিন)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "উরুগুয়ায়ান পেসো এন উনিদাদেস ইনডেক্সেডাস", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "উরুগুয়ে পেসো (১৯৭৫–১৯৯৩)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "উরুগুয়ে পেসো", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "উজবেকিস্তানি সোম", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "ভেনিজুয়েলান বলিভার (১৮৭১–২০০৮)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "ভেনিজুয়েলীয় বলিভার (২০০৮–২০১৮)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "ভেনিজুয়েলীয় বলিভার", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "ভিয়েতনামি ডঙ্গ", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "ভানুয়াতু ভাতু", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "সামোয়ান টালা", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "মধ্য আফ্রিকান [CFA] ফ্র্যাঙ্ক", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "সিলভার", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "গোল্ড", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "ইউরোপীয় আর্থিক একক", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "পূর্ব ক্যারাবিয়ান ডলার", Symbol: "EC$"},
				currency.XEU: cldr.Currency{DisplayName: "ইউরোপীয় মুদ্রা একক", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "ফরাসি গোল্ড ফ্রাঙ্ক", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "ফরাসি ইউআইসি - ফ্রাঙ্ক", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "পশ্চিম আফ্রিকান [CFA] ফ্র্যাঙ্ক", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "প্যালেডিয়াম", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "সিএফপি ফ্র্যাঙ্ক", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "প্লাটিনাম", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "অজানা মুদ্রা", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "ইয়েমেনি দিনার", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "ইয়েমেনি রিয়াল", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "যুগোশ্লাভিয় হার্ড দিনার", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "যুগোশ্লাভিয় নোভি দিনার", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "যুগোশ্লাভিয় রুপান্তরযোগ্য দিনার", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "দক্ষিণ আফ্রিকান র\u200c্যান্ড", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "দক্ষিণ আফ্রিকান রেন্ড", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "জাম্বিয়ান কওয়াচা (১৯৬৮–২০১২)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "জাম্বিয়ান কওয়াচা", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "জাইরিয়ান নিউ জাইরে", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "জাইরিয়ান জাইরে", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "জিম্বাবুয়ে ডলার (১৯৮০–২০০৮)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "জিম্বাবুয়ে ডলার (২০০৯)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "জিম্বাবুয়ে ডলার (২০০৮)", Symbol: ""},
			},
		},
	}
}
