// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_kok_IN() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "dd-MM-y", Short: "d-M-yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "जाने", Feb: "फेब्रु", Mar: "मार्च", Apr: "एप्री", May: "मे", Jun: "जून", Jul: "जुल", Aug: "ऑग", Sep: "सप्टें", Oct: "ऑक्टो", Nov: "नो", Dec: "डिसे"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "जानेवारी", Feb: "फेब्रुवारी", Mar: "मार्च", Apr: "एप्रिल", May: "मे", Jun: "जून", Jul: "जुलय", Aug: "ऑगस्ट", Sep: "सप्टेंबर", Oct: "ऑक्टोबर", Nov: "नोव्हेंबर", Dec: "डिसेंबर"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "आयतार", Mon: "सोमार", Tue: "मंगळार", Wed: "बुधवार", Thu: "बिरेस्तार", Fri: "शुक्रार", Sat: "शेनवार"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "आ", Mon: "सो", Tue: "मं", Wed: "बु", Thu: "ब", Fri: "शु", Sat: "शे"}, Short: cldr.CalendarDayFormatNameValue{Sun: "आय", Mon: "सोम", Tue: "मंगळ", Wed: "बुध", Thu: "बिरे", Fri: "शुक्र", Sat: "शेन"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "आयतार", Mon: "सोमार", Tue: "मंगळार", Wed: "बुधवार", Thu: "बिरेस्तार", Fri: "शुक्रार", Sat: "शेनवार"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_root]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "युनाइटेड अरब इमीरॅट्स दिरहम", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "अफगाण अफगाणी", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "अल्बेनियन लेक", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "अर्मेनियन ड्राम", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "नॅदरलँड अँटिलियन गिल्डर", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "अंगोलन क्वॉन्ज", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "अर्जेंटिना पेसो", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "ऑस्ट्रेलियाई डॉलर", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "अरुबान फ्लोरिन", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "अज़रबैजानी मनात", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "बोस्निया-हेर्जेगोविना रुपांतरीत मार्क", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "बार्बाडियान डॉलर", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "बांगलादेशी टाका", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "बल्गेरियन लेव", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "बहरिनी डिनार", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "बुरुंडी फ्रँक", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "बरमुदान डॉलर", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "ब्रूनेई डॉलर", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "बोलिव्हियन बोलिवियानो", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "ब्राझिलियन रियाल", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "बहामियन डॉलर", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "भुतानीज नागल्ट्रम", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "बोत्सवाना पुला", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "बैलोरुसियन् रूबल", Symbol: "BYN"},
				currency.BZD: cldr.Currency{DisplayName: "बेलिझ डॉलर", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "कॅनाडियन डॉलर", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "काँगोलिसी फ्रँक", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "स्विस फ्रँक", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "चिली पेसो", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "चिनी युआन (ऑफशोर)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "चिनी युआन", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "कोलंबियन पेसो", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "कोस्ता रिका कॉलॉन", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "क्युबान रुपांतरीत पेसो", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "क्युबान पेसो", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "केप वर्दे एस्कुडो", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "चेक कोरुना", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "जिबूती फ्रँक", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "डॅनिश क्रोन", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "डोमिनिकन पेसो", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "अल्जेरियाई डिनार", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "ईजिप्ती पावंड", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "इरिट्रियन नाक्फा", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "इथियोपियाई बिरर", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "युरो", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "फिजी डॉलर", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "फ़ॉकलैंड आइलैंड्स पावंड", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "ब्रिटिश पावंड", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "जॉर्जियन लारी", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "घानाई सेडी", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "जिब्राल्टर पावंड", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "गॅम्बियन दलासी", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "गिनीन फ्रँक", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "ग्वाटेमाला कुएट्झल", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "गयाना डॉलर", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "हाँग काँग डॉलर", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "होंडुरान लेम्पिरा", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "क्रोयेषियन् कुना", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "हैतीयन गौर्डे", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "हंगेरियन फोरिंट", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "इंडोनेशियन रुपिया", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "इस्त्रायली न्यु शेकेल", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "भारतीय रुपया", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "इराकी डिनार", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ईरानी रियाल", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "आईस्लान्डिक क्रोना", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "जमैकन डॉलर", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "जॉर्डनियन डिनार", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "जपानी येन", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "केनयाई शिलिंग", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "किरगिझस्तान सोम", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "कंबोडियन रियाल", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "कोमोरियन फ्रँक", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "उत्तर कोरियन वॉन", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "दक्षिण कोरियन वॉन", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "कुवेती डिनार", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "कैमेन आइलैंड्स डॉलर", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "कझाकस्तानी टेंग", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "लाओ किप", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "लिबानेस पावंड", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "श्री लंका रुपया", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "लायबेरियन डॉलर", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "लीबियान डिनार", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "मोरक्कन दिरहम", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "मोल्दोवान लियू", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "मलागासी एरियारी", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "मसीडोनियन् डिनर", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "म्यानमार क्यात", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "मंगोलियन तुगरिक", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "मकानेसे पटका", Symbol: "MOP"},
				currency.MRU: cldr.Currency{DisplayName: "मॉरिटानिया उगिया", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "मॉरिशस रुपी", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "मालदिवी रुफिया", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "मलावियन क्वाचा", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "मेक्सिकन पेसो", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "मलेशियाई रिंग्गित", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "मोझांबिकन मेटिकल", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "नामीबिया डॉलर", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "नायजेरियन नायरा", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "निकारागुआन कॉर्डोबा", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "नॉर्वेगन क्रोन", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "नेपाळी रुपया", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "न्युझीलॅन्ड डॉलर", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ओमानी रियाल", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "पानामानियन बाल्बोआ", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "पेरिवियन सोल", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "पापुआ न्यु गिनी किना", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "फिलिपिनी पेसो", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "पाकिस्तानी रुपया", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "पोलिष झ्लोटी", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "पराग्वेन गौरानी", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "कतारी रियाल", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "रोमानियन् लियू", Symbol: "रॉन"},
				currency.RSD: cldr.Currency{DisplayName: "सर्बियन डिनार", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "रुसी रुबल", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "रवांडा फ्रँक", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "सौदी रियाल", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "सोलोमन आयलँड्स डॉलर", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "सेशेल्लोइस रुपी", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "सुदानी पावंड", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "स्वीदीष क्रोन", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "सिंगापूरी डॉलर", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "सेंट हेलिना पावंड", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "सिएरा लियॉनी लियॉन", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "सोमाली शिलिंग", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "सुरीनामी डॉलर", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "दक्षिण सुडानी पावंड", Symbol: "SSP"},
				currency.STN: cldr.Currency{DisplayName: "साओ टोम आनी प्रिन्सिप डोब्रा", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "सिरियन पावंड", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "स्वाजी लिलांगेनी", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "थाई बाट", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "ताजिकिस्तानी सोमोनी", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "तुर्कमेनिस्तानी मनत", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "ट्यूनीशियन डिनार", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "टोंगन पांगा", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "तुर्किश लायरा", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ट्रिनीडाड आनी टोबॅगो डॉलर", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "न्यू तायवान डॉलर", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "तंजानिया शिलिंग", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "युक्रेनियन् रिव्निया", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "युगांडा शिलिंग", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "युएस डॉलर", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "उरुग्वेन पेसो", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "उज़्बेकिस्तानी सोम", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "विनेझुएला बोलिव्हर", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "वियतनामी डोंग", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "वानूआतू वातू", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "समोआई टाला", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "मध्य अफ्रीकी सीएफए फ्रँक", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "उदेंत कॅरिबियन डॉलर", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "अस्तंत आफ्रिकी सीएफए फ्रँक", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "सीएफपी फ्रँक", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "अज्ञात चलन", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "येमेनी रियाल", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "दक्षिण आफ्रिकन रँड", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "झांबियन क्वाचा", Symbol: "ZMW"},
			},
		},
	}
}
