// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_he_IL() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "he_IL",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d בMMMM y", Long: "d בMMMM y", Medium: "d בMMM y", Short: "d.M.y"}, Time: cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} בשעה {0}", Long: "{1} בשעה {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}\u200e"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ינו׳", Feb: "פבר׳", Mar: "מרץ", Apr: "אפר׳", May: "מאי", Jun: "יוני", Jul: "יולי", Aug: "אוג׳", Sep: "ספט׳", Oct: "אוק׳", Nov: "נוב׳", Dec: "דצמ׳"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ינואר", Feb: "פברואר", Mar: "מרץ", Apr: "אפריל", May: "מאי", Jun: "יוני", Jul: "יולי", Aug: "אוגוסט", Sep: "ספטמבר", Oct: "אוקטובר", Nov: "נובמבר", Dec: "דצמבר"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "יום א׳", Mon: "יום ב׳", Tue: "יום ג׳", Wed: "יום ד׳", Thu: "יום ה׳", Fri: "יום ו׳", Sat: "שבת"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "א׳", Mon: "ב׳", Tue: "ג׳", Wed: "ד׳", Thu: "ה׳", Fri: "ו׳", Sat: "ש׳"}, Short: cldr.CalendarDayFormatNameValue{Sun: "א׳", Mon: "ב׳", Tue: "ג׳", Wed: "ד׳", Thu: "ה׳", Fri: "ו׳", Sat: "ש׳"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "יום ראשון", Mon: "יום שני", Tue: "יום שלישי", Wed: "יום רביעי", Thu: "יום חמישי", Fri: "יום שישי", Sat: "יום שבת"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "לפנה״צ", PM: "אחה״צ"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_he]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "\u200f#,##0.00\u00a0¤;\u200f-#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "פזטה אנדורית", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "דירהם של איחוד הנסיכויות הערביות", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "אפגני אפגני", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "לק אלבני", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "דראם ארמני", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "גילדר של האנטילים ההולנדיים", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "קואנזה אנגולי", Symbol: "AOA"},
				currency.AON: cldr.Currency{DisplayName: "קואנזה חדש אנגולי (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "קואנזה רג׳וסטדו אנגולי (1995–1999)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "פזו ארגנטינאי (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "פסו ארגנטינאי", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "שילינג אוסטרי", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "דולר אוסטרלי", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "פלורין של ארובה", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "מנאט אזרביג׳אני (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "מאנאט אזרבייג׳ני", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "דינר של בוסניה־הרצגובינה", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "מארק סחיר של בוסניה והרצגובינה", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "דולר ברבדיאני", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "טאקה בנגלדשי", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "פרנק בלגי (בר המרה)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "פרנק בלגי", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "לב בולגרי ישן", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "לב בולגרי", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "דינר בחרייני", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "פרנק בורונדי", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "דולר ברמודה", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "דולר ברוניי", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "בוליביאנו", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "פזו בוליבי", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "קרוזיארו חדש ברזילאי (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "קרוזדו ברזילאי", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "ריאל ברזילאי", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "דולר בהאמי", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "נגולטרום בהוטני", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "פולה בוטסואני", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "רובל בלרוסי", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "רובל בלרוסי (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "דולר בליזי", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "דולר קנדי", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "פרנק קונגולזי", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "פרנק שוויצרי", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "פסו צ׳ילאני", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "יואן סיני (CNH)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "יואן סיני", Symbol: "\u200eCN¥\u200e"},
				currency.COP: cldr.Currency{DisplayName: "פסו קולומביאני", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "קולון קוסטה־ריקני", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "דינר סרבי ישן", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "פזו קובני להמרה", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "פזו קובני", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "אסקודו כף ורדה", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "לירה קפריסאית", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "קורונה צ׳כית", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "מרק מזרח גרמני", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "מרק גרמני", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "פרנק ג׳יבוטי", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "כתר דני", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "פזו דומיניקני", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "דינר אלג׳ירי", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "סוקר אקואדורי", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "קרון אסטוני", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "לירה מצרית", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "נאקפה אריתראי", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "פזטה [ESA]", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "פזטה [ESB]", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "פסטה ספרדי", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ביר אתיופי", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "אירו", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "מרק פיני", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "דולר פיג׳י", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "לירה של איי פוקלנד", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "פרנק צרפתי", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "לירה שטרלינג", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "לארי גאורגי", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "סדי גאני", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "פאונד גיברלטר", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "דלסי גמבי", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "פרנק גינאי", Symbol: "GNF"},
				currency.GRD: cldr.Currency{DisplayName: "דרכמה", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "קצאל גואטמלי", Symbol: "GTQ"},
				currency.GWP: cldr.Currency{DisplayName: "פזו גינאי", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "דולר גיאני", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "דולר הונג קונגי", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "למפירה הונדורי", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "קונה קרואטי", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "גורד האיטי", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "פורינט הונגרי", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "רופיה אינדונזית", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "לירה אירית", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "לירה ישראלית", Symbol: "ל״י"},
				currency.ILS: cldr.Currency{DisplayName: "שקל חדש", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "רופי הודי", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "דינר עיראקי", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ריאל איראני", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "כתר איסלנדי", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "לירה איטלקית", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "דולר ג׳מייקני", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "דינר ירדני", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "ין יפני", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "שילינג קנייתי", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "סום קירגיזי", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "ריל קמבודי", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "פרנק קומורואי", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "וון צפון קוריאני", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "וון דרום קוריאני", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "דינר כוויתי", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "דולר קיימני", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "טנגה קזחסטני", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "קיפ לאי", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "לירה לבנונית", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "רופי סרי לנקי", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "דולר ליברי", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "לוטי לסותי", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "ליטא ליטאי", Symbol: "LTL"},
				currency.LUF: cldr.Currency{DisplayName: "פרנק לוקסמבורגי", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "לט לטבי", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "דינר לובי", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "דירהם מרוקאי", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "פרנק מרוקאי", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "לאו מולדובני", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "אריארי מלגשי", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "פרנק מדגסקארי", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "דינר מקדוני", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "קיאט מיאנמרי", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "טוגרוג מונגולי", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "פטקה של מקאו", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "אואוגויה מאוריטני (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "אואוגויה מאוריטני", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "לירה מלטית", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "רופי מאוריציני", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "רופיה מלדיבית", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "קואצ׳ה מלאווי", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "פזו מקסיקני", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "פזו מקסיקני (1861 – 1992)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "רינגיט מלזי", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "אסקודו מוזמביקי", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "מטיקל", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "מטיקל מוזמביני", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "דולר נמיבי", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "נאירה ניגרי", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "קורדובה ניקרגואה", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "גילדן הולנדי", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "כתר נורווגי", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "רופי נפאלי", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "דולר ניו זילנדי", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ריאל עומאני", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "בלבואה פנמי", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "סול פרואני", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "קינה של פפואה גינאה החדשה", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "פזו פיליפיני", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "רופי פקיסטני", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "זלוטי פולני", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "זלוטי (1950 – 1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "אסקודו פורטוגלי", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "גוארני פרגוואי", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "ריאל קטארי", Symbol: "QAR"},
				currency.ROL: cldr.Currency{DisplayName: "לאו רומני ישן", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "לאו רומני", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "דינר סרבי", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "רובל רוסי", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "רובל רוסי (1991 – 1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "פרנק רואנדי", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "ריאל סעודי", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "דולר איי שלמה", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "רופי סיישלי", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "דינר סודני", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "לירה סודנית", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "לירה סודנית (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "כתר שוודי", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "דולר סינגפורי", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "פאונד סנט הלני", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "טולאר סלובני", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "קורונה סלובקי", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "ליאון סיירה לאוני", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "שילינג סומלי", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "דולר סורינאמי", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "גילדר סורינאמי", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "לירה דרום-סודנית", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "דוברה של סן טומה ופרינסיפה (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "דוברה של סאו טומה ופרינסיפה", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "רובל סובייטי", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "קולון סלבדורי", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "לירה סורית", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "לילנגני סווזילנדי", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "בהט תאילנדי", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "סומוני טג׳קיסטני", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "מנאט טורקמאני", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "מאנאט טורקמני", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "דינר טוניסאי", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "פאנגה טונגי", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "אסקודו טימוראי", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "לירה טורקית", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "לירה טורקית חדשה", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "דולר טרינידדי", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "דולר טייוואני חדש", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "שילינג טנזני", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "הריבנה אוקראיני", Symbol: "UAH"},
				currency.UGS: cldr.Currency{DisplayName: "שילינג אוגנדי (1966 – 1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "שילינג אוגנדי", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "דולר אמריקאי", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "דולר אמריקאי (היום הבא)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "דולר אמריקאי (היום הזה)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "פסו אורוגוואי", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "סום אוזבקי", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "בוליבר ונצואלי (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "בוליבר ונצואלי (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "בוליבר ונצואלי", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "דונג וייטנאמי", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "ואטו של ונואטו", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "טאלה סמואי", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "פרנק CFA מרכז אפריקני", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "כסף", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "זהב", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "דולר מזרח קריבי", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "זכויות משיכה מיוחדות", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "פרנק זהב", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "פרנק CFA מערב אפריקני", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "פלדיום", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "פרנק פולינזיה הצרפתית", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "פלטינה", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "סימון למטרות בדיקה", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "מטבע שאינו ידוע", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "דינר תימני", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "ריאל תימני", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "דינר יגוסלבי חדש", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "דינר יגוסלבי", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "ראנד דרום אפריקאי (כספי)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "ראנד דרום אפריקאי", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "קוואצ׳ה זמבית (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "קוואצ׳ה זמבי", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "זאיר חדש", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "דולר זימבבואי", Symbol: ""},
			},
		},
	}
}
