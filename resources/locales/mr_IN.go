// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_mr_IN() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} रोजी {0}", Long: "{1} रोजी {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "[GMT]{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "जाने", Feb: "फेब्रु", Mar: "मार्च", Apr: "एप्रि", May: "मे", Jun: "जून", Jul: "जुलै", Aug: "ऑग", Sep: "सप्टें", Oct: "ऑक्टो", Nov: "नोव्हें", Dec: "डिसें"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "जा", Feb: "फे", Mar: "मा", Apr: "ए", May: "मे", Jun: "जू", Jul: "जु", Aug: "ऑ", Sep: "स", Oct: "ऑ", Nov: "नो", Dec: "डि"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "जानेवारी", Feb: "फेब्रुवारी", Mar: "मार्च", Apr: "एप्रिल", May: "मे", Jun: "जून", Jul: "जुलै", Aug: "ऑगस्ट", Sep: "सप्टेंबर", Oct: "ऑक्टोबर", Nov: "नोव्हेंबर", Dec: "डिसेंबर"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "रवि", Mon: "सोम", Tue: "मंगळ", Wed: "बुध", Thu: "गुरु", Fri: "शुक्र", Sat: "शनि"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "र", Mon: "सो", Tue: "मं", Wed: "बु", Thu: "गु", Fri: "शु", Sat: "श"}, Short: cldr.CalendarDayFormatNameValue{Sun: "र", Mon: "सो", Tue: "मं", Wed: "बु", Thu: "गु", Fri: "शु", Sat: "श"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "रविवार", Mon: "सोमवार", Tue: "मंगळवार", Wed: "बुधवार", Thu: "गुरुवार", Fri: "शुक्रवार", Sat: "शनिवार"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "म.पू.", PM: "म.उ."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "म.पू.", PM: "म.उ."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "म.पू.", PM: "म.उ."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_mr]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "संयुक्त अरब अमीरात दिरहॅम", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "अफगाण अफगाणी", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "अल्बानियन लेक", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "आर्मेनियन द्रॅम", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "नेदरलँडचा अँटिलीन गिल्डर", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "अँगोलन क्वॅन्झा", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "अर्जेंटाइन पेसो", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "ऑस्ट्रेलियन डॉलर", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "अरुबा फ्लोरिन", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "अझरबैझानी मानाट", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "बोस्निया-हर्जेगोविना विनिमय मार्क", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "बार्बाडियन डॉलर", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "बांगलादेशी टका", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "बल्गेरियन लेव", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "बाहरिनी दिनार", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "बुरुंडियन फ्रँक", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "बर्मुडा डॉलर", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "ब्रुनेई डॉलर", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "बोलिव्हियन बोलिव्हियानो", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "ब्राझिलियन रियाल", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "बहामी डॉलर", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "भूतानी एंगल्ट्रम", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "बोट्सवानन पुला", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "बेलारुशियन रुबल", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "बेलारुशियन रुबल (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "बेलीझ डॉलर", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "कॅनडियन डॉलर", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "काँगोलीज फ्रँक", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "स्विस फ्रँक", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "चिली पेसो", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "चिनी युआन (ऑफशोर)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "चीनी युआन", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "कोलंबियन पेसो", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "कोस्टा रिका कोलोन", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "क्यूबन विनिमय पेसो", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "क्यूबन पेसो", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "केप व्हर्डेयन एस्कुडो", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "झेक प्रजासत्ताक कोरुना", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "जिबौटियन फ्रँक", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "डॅनिश क्रोन", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "डोमिनिकन पेसो", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "अल्जेरियन दिनार", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "इजिप्शियन पाउंड", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "इरिट्रियन नाक्फा", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "इथिओपियन बिर", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "युरो", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "फिजियन डॉलर", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "फॉकलंड आयलंड पाउंड", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "ब्रिटिश पाऊंड", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "जॉर्जियन लारी", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "घानीयन सेडी", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "जिब्राल्टर पाउंड", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "गाम्बियन डालासी", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "गिनी फ्रँक", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "ग्वाटेमालाचे क्वेत्झाल", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "गयाना डॉलर", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "हाँगकाँग डॉलर", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "होन्डुरन लेंपिरा", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "क्रोएशियन कूना", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "हैती गोअर्ड", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "हंगेरियन फॉरिन्ट", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "इंडोनेशियन रुपिया", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "ईस्त्रायली न्यू शेकेल", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "भारतीय रुपया", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "इराकी दिनार", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "इराणी रियाल", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "आइसलँडिक क्रोना", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "जमैकन डॉलर", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "जॉर्डनियन दिनार", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "जपानी येन", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "केनियन शिलिंग", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "किरगिस्तानी सॉम", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "कंबोडियन रियेल", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "कोमोरियन फ्रँक", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "उत्तर कोरियन वॉन", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "दक्षिण कोरियन वॉन", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "कुवैती दिनार", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "केमेन आयलॅंड डॉलर", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "कझाकिस्तानी तेंगे", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "लाओशियन किप", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "लेबनीज पाउंड", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "श्रीलंकन रुपया", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "लाइबेरियन डॉलर", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "लेसोटो लोटी", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "लिथुआनियन लिटास", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "लाट्व्हियन लाट्झ", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "लिबियाचा दिनार", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "मोरोक्को दिरहॅम", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "मोल्डोवन लेउ", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "मालागासी एरियारी", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "मॅसेडोनियन देनार", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "म्यानमार क्याट", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "मंगोलियन टुग्रिक", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "मॅकॅनीज् पटाका", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "मॉरिटानियन ओगिया (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "मॉरिटानियन ओगिया", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "मॉरिशियन रुपी", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "मालदीवियन रुफिया", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "मालावियन क्वाचा", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "मेक्सिको पेसो", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "मलेशियन रिंगिट", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "मोझांबिकन मेटिकल", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "नमिबियन डॉलर", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "नायजेरियन नायरा", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "निकाराग्वेचा कोर्डोबा", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "नॉर्वेजियन क्रोन", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "नेपाळी रुपया", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "न्यूझीलँड डॉलर", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ओमानी रियाल", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "पनामा बाल्बोआ", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "पेरुवियन सोल", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "पापुआ न्यू गिनीयन किना", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "फिलिपिनी पेसो", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "पाकिस्तानी रुपया", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "पोलिश झ्लॉटी", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "पराग्वे ग्वारानी", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "कतारी रियाल", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "रोमानियन लेऊ", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "सर्बियन दिनार", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "रशियन रुबल", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "रवांडा फ्रँक", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "सौदी रियाल", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "सोलोमन आयलँड्स डॉलर", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "सेशेलोईस रुपी", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "सुदानी पाउंड", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "स्वीडिश क्रोना", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "सिंगापूर डॉलर", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "सेंट हेलेना पाउंड", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "सिएरा लिऑनचा लिऑन", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "सोमाली शिलिंग", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "सुरिनामी डॉलर", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "दक्षिण सुदानी पाउंड", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "साओ टोम आणि प्रिन्सिपे डोबरा (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "साओ टोम आणि प्रिन्सिपे डोबरा", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "सीरियन पाउंड", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "स्वाझी लीलांगेनी", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "थाई बाहत", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "ताजकीस्तानी सोमोनी", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "तुर्कमेनिस्तानी मानाट", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "ट्यूनिशियन दिनार", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "टोंगन पाआंगा", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "तुर्की लिरा", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "त्रिनिदाद आणि टोबॅगो डॉलर", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "न्यू तैवान डॉलर", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "टांझानियन शिलिंग", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "युक्रेनियन रिवनिया", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "युगांडा शिलिंग", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "यूएस डॉलर", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "उरुग्वेचा पेसो", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "उझबेकिस्तानी सोम", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "व्हेनेझुएला बोलिव्हार (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "व्हेनेझुएला बोलिव्हार", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "व्हिएतनामी डोंग", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "वानुआतु वाटु", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "सामोअन टाला", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "मध्य आफ्रिकन [CFA] फ्रँक", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "पूर्व कॅरीबियन डॉलर", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "पश्चिम आफ्रिकन [CFA] फ्रँक", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "[CFP] फ्रँक", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "अज्ञात चलन", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "येमेनी रियाल", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "दक्षिण आफ्रिकी रँड", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "झांबियन क्वाचा (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "झांबियन क्वाचा", Symbol: "ZMW"},
			},
		},
	}
}
