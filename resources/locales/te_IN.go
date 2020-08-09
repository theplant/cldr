// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_te_IN() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "d, MMMM y, EEEE", Long: "d MMMM, y", Medium: "d MMM, y", Short: "dd-MM-yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}కి", Long: "{1} {0}కి", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "జన", Feb: "ఫిబ్ర", Mar: "మార్చి", Apr: "ఏప్రి", May: "మే", Jun: "జూన్", Jul: "జులై", Aug: "ఆగ", Sep: "సెప్టెం", Oct: "అక్టో", Nov: "నవం", Dec: "డిసెం"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "జ", Feb: "ఫి", Mar: "మా", Apr: "ఏ", May: "మే", Jun: "జూ", Jul: "జు", Aug: "ఆ", Sep: "సె", Oct: "అ", Nov: "న", Dec: "డి"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "జనవరి", Feb: "ఫిబ్రవరి", Mar: "మార్చి", Apr: "ఏప్రిల్", May: "మే", Jun: "జూన్", Jul: "జులై", Aug: "ఆగస్టు", Sep: "సెప్టెంబర్", Oct: "అక్టోబర్", Nov: "నవంబర్", Dec: "డిసెంబర్"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ఆది", Mon: "సోమ", Tue: "మంగళ", Wed: "బుధ", Thu: "గురు", Fri: "శుక్ర", Sat: "శని"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "ఆ", Mon: "సో", Tue: "మ", Wed: "బు", Thu: "గు", Fri: "శు", Sat: "శ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ఆది", Mon: "సోమ", Tue: "మం", Wed: "బుధ", Thu: "గురు", Fri: "శుక్ర", Sat: "శని"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "ఆదివారం", Mon: "సోమవారం", Tue: "మంగళవారం", Wed: "బుధవారం", Thu: "గురువారం", Fri: "శుక్రవారం", Sat: "శనివారం"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_te]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##,##0.###", Currency: "¤#,##,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "యునైటెడ్ ఆరబ్ ఎమిరేట్స్ దిరామ్", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "ఆఫ్ఘాన్ ఆఫ్ఘాని", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "ఆల్బేనియన్ లేక్", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "అమెరికన్ డ్రామ్", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "నెదర్లాండ్స్ యాంటిల్లియన్ గిల్\u200cడర్", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "అంగోలాన్ క్వాన్\u200cజా", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "అర్జెంటీనా పెసో", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "ఆస్ట్రేలియన్ డాలర్", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "అరుబన్ ఫ్లోరిన్", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "అజర్బైజాన్ మానట్", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "బోస్నియా-హెర్జగోవినా మార్పిడి చెయ్యగలిగే మార్క్", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "బర్బాడియన్ డాలర్", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "బాంగ్లాదేశ్ టాకా", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "బల్గేరియన్ లేవ్", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "బహ్రెయిన్ దినార్", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "బురిండియన్ ఫ్రాంక్", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "బెర్ముడన్ డాలర్", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "బ్రూనై డాలర్", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "బొలీవియన్ బొలీవియానో", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "బ్రెజిలియన్ రియల్", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "బహామియన్ డాలర్", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "భూటానీయుల గుల్\u200cట్రుమ్", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "బోట్స్\u200cవానా పులా", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "బెలరూసియన్ రూబల్", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "బెలరూసియన్ రూబల్ (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "బెలీజ్ డాలర్", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "కెనడియన్ డాలర్", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "కొంగోలిస్ ఫ్రాంక్", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "స్విస్ ఫ్రాంక్", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "చిలియన్ పెసో", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "చైనీస్ యూవాన్ (ఆఫ్\u200cషోర్)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "చైనా దేశ యువాన్", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "కొలంబియన్ పెసో", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "కోస్టా రికన్ కోలోన్", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "క్యూబన్ కన్వర్టబుల్ పెసో", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "క్యూబన్ పెసో", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "కేప్ వెర్డియన్ ఎస్కుడో", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "చెక్ రిపబ్లిక్ కోరునా", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "జిబోటియన్ ఫ్రాంక్", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "డానిష్ క్రోన్", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "డోమినికన్ పెసో", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "అల్జీరియన్ దీనార్", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "ఈజిప్షియన్ పౌండ్", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "ఎరిట్రీన్ నక్ఫా", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ఇథియోపియన్ బర్", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "యురొ", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "ఫీజియన్ డాలర్", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "ఫాక్\u200cల్యాండ్ దీవులు పౌండ్", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "బ్రిటిష్ పౌండ్", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "జార్జియన్ లారి", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "గానెయన్ సెడి", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "జిబ్రల్\u200cటూర్ పౌండ్", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "గాంబియన్ దలాసి", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "గ్వినియన్ ఫ్రాంక్", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "గ్యుటెమాలన్ క్వెట్\u200cజల్", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "గుయనియాస్ డాలర్", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "హాంకాంగ్ డాలర్", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "హోండురన్ లెమిపిరా", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "క్రొయేషియన్ క్యూన", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "హైటియన్ గ్వోర్డే", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "హంగేరియన్ ఫోరింట్", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "ఇండోనేషియా రూపాయి", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "ఇజ్రాయేలీ న్యూ షెకెల్", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "రూపాయి", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "ఇరాకీ దీనార్", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ఇరానియన్ రీయల్", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "ఐస్లాండిక్ క్రోనా", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "జమైకన్ డాలర్", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "జోర్\u200cడానియన్ దీనార్", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "జపాను దేశ యెన్", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "కెన్యాన్ షిల్లింగ్", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "కిర్గిస్థాని సౌమ్", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "కాంబోడియన్ రీల్", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "కొమోరియన్ ఫ్రాంక్", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "ఉత్తర కొరియా వోన్", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "దక్షిణ కొరియా వోన్", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "కువైట్ దీనార్", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "కేమాన్ దీవుల డాలర్", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "ఖజికిస్థాన్ టెంగే", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "లాటియన్ కిప్", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "లెబనీస్ పౌండ్", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "శ్రీలంక రూపాయి", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "లిబేరియన్ డాలర్", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "లెసోధో లోటి", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "లిథోనియన్ లీటాస్", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "లాత్వియన్ లాట్స్", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "లిబియన్ దీనార్", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "మోరోకన్ దిర్హుమ్", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "మోల్\u200cడోవన్ ల్యూ", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "మలగసీ అరియరీ", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "మెసిడోనియన్ దినార్", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "మయన్మార్ క్యాట్", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "మంగోలియన్ టుగ్రిక్", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "మకనీస్ పటాక", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "మౌరిటానియన్ ఒగ్యియా (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "మౌరిటానియన్ ఒగ్యియా", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "మారిషన్ రూపాయి", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "మాల్దీవుల రూపాయి", Symbol: ""},
				currency.MVR: cldr.Currency{DisplayName: "మాల్దీవియన్ రుఫియా", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "మలావియన్ క్వాచా", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "మెక్సికన్ పెసో", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "మలేషియా రింగ్గిట్", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "మొజాంబికన్ మెటికల్", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "నమిబియన్ డాలర్", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "నైజీరియన్ నైరా", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "నికరగ్యుయన్ కొర్\u200cడుబు", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "నార్వేజీయన్ క్రోన్", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "నేపాలీయుల రూపాయి", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "న్యూజిలాండ్ డాలర్", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ఒమాని రీయల్", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "పనామనియన్ బల్బోవ", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "పెరువియన్ సోల్", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "పప్యూ న్యూ గ్యినియన్ కినా", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "ఫిలిప్పిన్ పెసో", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "పాకిస్థాన్ రూపాయి", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "పోలిష్ జ్లోటీ", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "పరగ్వాయన్ గ్వారని", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "క్వాటరి రీయల్", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "రోమానియాన్ లెయు", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "సెర్బియన్ దీనార్", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "రష్యన్ రూబల్", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ర్వానడాన్ ఫ్రాంక్", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "సౌది రియల్", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "సోలోమన్ దీవుల డాలర్", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "సెయిచెల్లోయిస్ రూపాయి", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "సుడానీస్ పౌండ్", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "స్వీడిష్ క్రోనా", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "సింగపూర్ డాలర్", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "సెయింట్ హెలెనా పౌండ్", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "సీయిరు లియోనియన్ లీయోన్", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "సొమాలి షిల్లింగ్", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "సురినామీయుల డాలర్", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "దక్షిణ సుడానీస్ పౌండ్", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "సావో టోమ్ మరియు ప్రిన్సిపి డోబ్రా (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "సావో టోమ్ మరియు ప్రిన్సిపి డోబ్రా", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "సిరీయన్ పౌండ్", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "స్వాజి లిలాన్గేని", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "థాయ్ బాట్", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "తజికిస్థాన్ సమోని", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "తుర్క్\u200cమెనిస్థాని మనాట్", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "తునీషియన్ దీనార్", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "టోంగాన్ పాంʻగా", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "టర్కిస్ లీరా", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ట్రినిడాడ్ మరియు టొబాగో డాలర్", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "క్రొత్త తైవాన్ డాలర్", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "టాంజానియన్ షిల్లింగ్", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ఉక్రయినియన్ హ్రివ్\u200cనియా", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "ఉగాండన్ షిల్లింగ్", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "అమెరికా డాలర్", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "ఉరుగ్వెయన్ పెసో", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "ఉజ్\u200cబెకిస్తాన్ సౌమ్", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "వెనుజులా బోలివర్ (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "వెనుజులా బోలివర్", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "వియత్నామీయుల డాంగ్", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "వనాటు వటు", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "సమోయన్ తాలా", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "సెంట్రల్ ఆఫ్రికన్ సిఎఫ్\u200cఎ ఫ్రాంక్", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "వెండి", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "బంగారం", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "తూర్పు కరీబియన్ డాలర్", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "పశ్చిమ ఆఫ్రికన్ సిఏఫ్ఏ ఫ్రాంక్", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "సిఎఫ్\u200cపి ఫ్రాంక్", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "ప్లాటినం", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "తెలియని కరెన్సీ", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "ఎమునీ రీయల్", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "దక్షిణ ఆఫ్రికా ర్యాండ్", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "జాంబియన్ క్వాచా (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "జాంబియన్ క్వాచా", Symbol: "ZMW"},
			},
		},
	}
}
