// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ne() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ne",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "yy/M/d"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "जनवरी", Feb: "फेब्रुअरी", Mar: "मार्च", Apr: "अप्रिल", May: "मे", Jun: "जुन", Jul: "जुलाई", Aug: "अगस्ट", Sep: "सेप्टेम्बर", Oct: "अक्टोबर", Nov: "नोभेम्बर", Dec: "डिसेम्बर"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "जन", Feb: "फेेब", Mar: "मार्च", Apr: "अप्र", May: "मे", Jun: "जुन", Jul: "जुल", Aug: "अग", Sep: "सेप", Oct: "अक्टो", Nov: "नोभे", Dec: "डिसे"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "जनवरी", Feb: "फेब्रुअरी", Mar: "मार्च", Apr: "अप्रिल", May: "मे", Jun: "जुन", Jul: "जुलाई", Aug: "अगस्ट", Sep: "सेप्टेम्बर", Oct: "अक्टोबर", Nov: "नोभेम्बर", Dec: "डिसेम्बर"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "आइत", Mon: "सोम", Tue: "मङ्गल", Wed: "बुध", Thu: "बिहि", Fri: "शुक्र", Sat: "शनि"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "आ", Mon: "सो", Tue: "म", Wed: "बु", Thu: "बि", Fri: "शु", Sat: "श"}, Short: cldr.CalendarDayFormatNameValue{Sun: "आइत", Mon: "सोम", Tue: "मङ्गल", Wed: "बुध", Thu: "बिहि", Fri: "शुक्र", Sat: "शनि"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "आइतबार", Mon: "सोमबार", Tue: "मङ्गलबार", Wed: "बुधबार", Thu: "बिहिबार", Fri: "शुक्रबार", Sat: "शनिबार"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "पूर्वाह्न", PM: "अपराह्न"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "पूर्वाह्न", PM: "अपराह्न"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "पूर्वाह्न", PM: "अपराह्न"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ne]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##,##0.###", Currency: "¤\u00a0#,##,##0.00", CurrencyAccounting: "", Percent: "#,##,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "संयुक्त अरब एमिराट्स डिर्हाम", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "अफ्गानी(१९२७–२००२)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "अफ्गान अफ्गानी", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "अल्बानियन लेक", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "आर्मेनियाली ड्राम", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "नेदरल्याण्ड्स एन्टिलियन गिल्डर", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "एङ्गोलान क्वान्जा", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "अर्जेन्टिनी पेसो", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "अष्ट्रेलियन डलर", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "आरूबन फ्लोरिन", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "अजरबैजानी मानात", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "बोस्निया-हर्जगोभिनिया रूपान्तरयोग्य मार्क", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "बर्बाडियन डलर", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "बङ्गलादेशी टाका", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "बुल्गारियाली लेभ", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "बाहारैनी डिनार", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "बुरूण्डियाली फ्रान्क", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "बर्मुडन डलर", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "ब्रुनाई डलर", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "बोलिभियन बोलिभियानो", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "ब्राजिलियन रियल", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "बहामियाली डलर", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "भुटानी एन्\u200cगुल्ट्रुम", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "बोट्सवानान पुला", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "बेलारूसी रूबल", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "बेलारूसी रूबल (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "वेलिज डलर", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "क्यानाडियाली डलर", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "कङ्गोली फ्रान्क", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "स्विस् फ्रैङ्क", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "चिलियन पेसो", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "चिनियाँ युआन(तटवर्ती)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "चिनिँया युआन", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "कोलम्वियन पेसो", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "कोष्टारिकन कोलोन", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "क्यूवाली रूपान्तरणयोग्य पेसो", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "क्यूवाली पेसो", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "केप भर्डियन एस्कुडो", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "चेख गणतञ्त्र कोरूना", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "जिबौंटियाली फ्रान्क", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "ड्यानिश क्रोन", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "डोमिनिकन पेसो", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "अल्जेरियाली डिनार", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "इजिप्सियन पाउन्ड", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "एरिट्रियन नाक्फा", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "इथियोपियाली बिर", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "युरो", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "फिजीयाली डलर", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "फक्\u200cल्याण्ड टापुहरूका पाउन्ड", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "बेलायती पाउण्ड स्टर्लिङ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "जर्जियाली लारी", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "घानाली सेडी", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "जिब्राल्टर पाउण्ड", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "गाम्वियाली डालासी", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "गिनियाली फ्रान्क", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "ग्वाटेमाला क्वेट्जाल", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "गाइनिज डलर", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "हङकङ डलर", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "होन्डुरान लेम्पिरा", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "क्रोएशियाली कुना", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "हैटियाली गुर्ड", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "हङ्गेरियन फोरिन्ट", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "इण्डोनेशियाली रूपियाँ", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "इजरायली नयाँ शेकेल", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "भारतीय रूपिँया", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "इराकी डिनार", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "इरानियाली रियाल", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "आइसल्याण्डिक क्रोना", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "जमाइकाली डलर", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "जोर्डानियाली डलर", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "जापानी येन", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "केन्याली शिलिङ", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "किर्गिस्तानी सोम", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "कम्बोडिनेयाली रियल", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "कोमोरियन फ्रान्क", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "उत्तर कोरियाली वन", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "दक्षिण कोरियाली वन", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "कुवेती डिनार", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "केम्यान टापुहरूका डलर", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "काजाखस्तानी टेन्ज", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "लाओशियन किप", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "लेबनाली पाउन्ड", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "श्रीलङ्काली रूपिया", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "लिबेरियाली डलर", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "लिथुनियाली लिटास", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "लाट्भियाली लाट्स", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "लिवियाली डिनार", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "मोरोक्काली डिर्\u200cहाम", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "माल्डोभन लेउ", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "मालागासी एरिआरी", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "म्यासेडोनियाली डेनार", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "म्यान्मार क्याट", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "मङ्गोलियाली टुग्रिक", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "माकानिज पटाका", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "माउरिटानियानली औगुइया (१९७३–२०१७)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "माउरिटानियानली औगुइया", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "माउरिटियन रूपी", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "मालडिभियाली रूफियाँ", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "मलाविअन क्वाचा", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "मेक्सिकन पेसो", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "मलेशियाली रिङ्गेट", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "मोजाम्विकन मेटिकल", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "नामिबियन डलर", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "नाइजेरियन नाइरा", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "निकारागुवान कोर्डोवा", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "नर्वेजियाली क्रोन", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "नेपाली रूपैयाँ", Symbol: "नेरू"},
				currency.NZD: cldr.Currency{DisplayName: "न्यूजिल्याण्ड डलर", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ओमनी रियल", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "पानामानियाली बाल्बोआ", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "पेरूभियाली सोल", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "पपुआ न्यू गिनियाली किना", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "फिलिपिनी पेसो", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "पाकिस्तानी रूपियाँ", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "पोलिश ज्लोटाई", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "पारागुयाली गुरानी", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "कतारी रियल", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "रोमानियाली लेऊ", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "सर्बियाली डिनार", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "रूसी रूबल", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "र्\u200cवाण्डाली फ्रान्क", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "साउदी रियालहरू", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "सोलोमन टापुहरूका डलर", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "सेचेलोइस रूपी", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "सुडानी पाउन्ड", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "स्विडिश क्रोना", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "सिङ्गापुर डलर", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "सेन्ट हेलेना पाउन्ड", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "सियरा लियोनेन लियोन", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "सोमाली शिलिङ", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "सुरिनामिज डलर", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "दक्षिण सुडानी पाउन्ड", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "साओ टोम र प्रिन्सिप डोब्रा (१९७७–२०१७)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "साओ टोम र प्रिन्सिप डोब्रा", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "सिरियाली पाउन्ड", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "स्वाजी लिलान्गेनी", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "थाई भाट", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "ताजिक्स्तानी सोमोनी", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "टुर्क्मेनिस्तानी मानात", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "टुनिसियाली डिनार", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "टङ्गन पाङ्गा", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "टर्किश लिरा", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "त्रिनिडाड र टोबागो डलर", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "नयाँ ताइवान डलर", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "ताञ्जानियाली शिलिङ", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "युक्रेनी हिर्भिनिया", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "युगाण्डाली शिलिङ", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "अमेरिकी डलर", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "उरूगुवायाली पेसो", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "उज्बेकिस्तान सोम", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "भेनेजुएलन बोलिभर (२००८–२०१८)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "भेनेजुएलन बोलिभर-2", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "भियतनामी डङ्", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "भानुआतू भातु", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "सामोआन ताला", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "सीएफ्\u200cए फ्रान्क बीइएसी", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "पूर्वी क्यारिबियन डलर", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "सीएफ्\u200cए फ्रान्क बीसीइएओ", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "सीएफ्\u200cपी फ्रान्क", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "अज्ञात मुद्रा", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "येमेनी रियाल", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "दक्षिण अफ्रिकी र्\u200dयान्ड", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "जाम्बियाली क्वाचा (१९६८–२०१२)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "जाम्बियाली क्वाचा", Symbol: "ZMW"},
			},
		},
	}
}
