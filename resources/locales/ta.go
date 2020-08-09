// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ta() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ta",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "a h:mm:ss zzzz", Long: "a h:mm:ss z", Medium: "a h:mm:ss", Short: "a h:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} ’அன்று’ {0}", Long: "{1} ’அன்று’ {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ஜன.", Feb: "பிப்.", Mar: "மார்.", Apr: "ஏப்.", May: "மே", Jun: "ஜூன்", Jul: "ஜூலை", Aug: "ஆக.", Sep: "செப்.", Oct: "அக்.", Nov: "நவ.", Dec: "டிச."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ஜ", Feb: "பி", Mar: "மா", Apr: "ஏ", May: "மே", Jun: "ஜூ", Jul: "ஜூ", Aug: "ஆ", Sep: "செ", Oct: "அ", Nov: "ந", Dec: "டி"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ஜனவரி", Feb: "பிப்ரவரி", Mar: "மார்ச்", Apr: "ஏப்ரல்", May: "மே", Jun: "ஜூன்", Jul: "ஜூலை", Aug: "ஆகஸ்ட்", Sep: "செப்டம்பர்", Oct: "அக்டோபர்", Nov: "நவம்பர்", Dec: "டிசம்பர்"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ஞாயி.", Mon: "திங்.", Tue: "செவ்.", Wed: "புத.", Thu: "வியா.", Fri: "வெள்.", Sat: "சனி"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "ஞா", Mon: "தி", Tue: "செ", Wed: "பு", Thu: "வி", Fri: "வெ", Sat: "ச"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ஞா", Mon: "தி", Tue: "செ", Wed: "பு", Thu: "வி", Fri: "வெ", Sat: "ச"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "ஞாயிறு", Mon: "திங்கள்", Tue: "செவ்வாய்", Wed: "புதன்", Thu: "வியாழன்", Fri: "வெள்ளி", Sat: "சனி"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "முற்பகல்", PM: "பிற்பகல்"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "மு.ப", PM: "பி.ப"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "முற்பகல்", PM: "பிற்பகல்"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ta]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##,##0.###", Currency: "¤\u00a0#,##,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "ஐக்கிய அரபு எமிரேட்ஸ் திர்ஹாம்", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "ஆஃப்கான் ஆஃப்கானி", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "அல்பேனியன் லெக்", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "ஆர்மேனியன் ட்ராம்", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "நெதர்லேண்ட்ஸ் அன்டிலியன் கில்டர்", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "அங்கோலன் க்வான்ஸா", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "அர்ஜென்டைன் பெசோ", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "ஆஸ்திரேலிய டாலர்", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "அருபன் ஃப்ளோரின்", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "அசர்பைஜானி மனத்", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "போஸ்னியா-ஹெர்ஸேகோவினா கன்வெர்டிபில் மார்க்", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "பார்பேடியன் டாலர்", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "பங்களாதேஷி டாகா", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "பல்கேரியன் லேவ்", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "பஹ்ரைனி தினார்", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "புருண்டியன் ஃப்ராங்க்", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "பெர்முடன் டாலர்", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "புரூனே டாலர்", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "பொலிவியன் பொலிவியானோ", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "பிரேசிலியன் ரியால்", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "பஹாமியன் டாலர்", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "பூட்டானீஸ் குல்ட்ரம்", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "போட்ஸ்வானன் புலா", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "பெலருசியன் ரூபிள்", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "பெலருசியன் ரூபிள் (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "பெலீஸ் டாலர்", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "கனடியன் டாலர்", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "காங்கோலீஸ் ஃப்ராங்க்", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "சுவிஸ் ஃப்ராங்க்", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "சிலியன் பெசோ", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "சீன யுவான் (ஆஃப்ஷோர்)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "சீன யுவான்", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "கொலம்பியன் பெசோ", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "கோஸ்டா ரிகன் கொலோன்", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "கியூபன் கன்வெர்டிபில் பெசோ", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "கியூபன் பெசோ", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "கேப் வெர்டியன் எஸ்குடோ", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "செக் குடியரசு கொருனா", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "ஜிபவ்டியென் ஃப்ராங்க்", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "டேனிஷ் க்ரோன்", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "டொமினிக்கன் பெசோ", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "அல்ஜீரியன் தினார்", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "எகிப்திய பவுண்டு", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "எரித்ரியன் நக்ஃபா", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "எத்தியோப்பியன் பிர்", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "யூரோ", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "ஃபிஜியன் டாலர்", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "ஃபாக்லாந்து தீவுகள் பவுண்டு", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "பிரிட்டிஷ் பவுண்டு", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "ஜார்ஜியன் லாரி", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "கானயன் சேடி", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "ஜிப்ரால்டர் பவுண்டு", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "கேம்பியன் தலாசி", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "கினியன் ஃப்ராங்க்", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "குவாதெமாலன் க்யுட்ஸல்", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "கயானீஸ் டாலர்", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "ஹாங்காங் டாலர்", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "ஹோன்டூரன் லெம்பீரா", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "குரோஷியன் குனா", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "ஹைட்டியன் கோர்டே", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "ஹங்கேரியன் ஃபோரின்ட்", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "இந்தோனேஷியன் ருபியா", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "இஸ்ரேலி நியூ ஷிகேல்", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "இந்திய ரூபாய்", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "ஈராக்கி தினார்", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ஈரானியன் ரியால்", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "ஐஸ்லாண்டிக் க்ரோனா", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "ஜமைக்கன் டாலர்", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "ஜோர்டானிய தினார்", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "ஜப்பானிய யென்", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "கென்யன் ஷில்லிங்", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "கிர்கிஸ்தானி சோம்", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "கம்போடியன் ரியெல்", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "கமோரியன் ஃப்ராங்க்", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "வட கொரிய வான்", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "தென் கொரிய வான்", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "குவைத்தி தினார்", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "கேமன் தீவுகள் டாலர்", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "கஸகஸ்தானி டென்கே", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "லவோஷியன் கிப்", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "லெபனீஸ் பவுண்டு", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "இலங்கை ரூபாய்", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "லைபீரியன் டாலர்", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "லெசோதோ லோட்டி", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "லிதுவேனியன் லிடஸ்", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "லத்வியன் லாட்ஸ்", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "லிபியன் தினார்", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "மொராக்கன் திர்ஹாம்", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "மால்டோவன் லியூ", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "மலகாசி ஏரியரி", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "மாசிடோனியன் டேனார்", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "மியான்மர் கியாத்", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "மங்கோலியன் டுக்ரிக்", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "மெகனீஸ் படாகா", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "மொரிஷானியன் ஒகுயா (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "மொரிஷானியன் ஒகுயா", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "மொரீஷியன் ருபீ", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "மாலத்தீவு ருஃபியா", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "மலாவியன் குவாச்சா", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "மெக்ஸிகன் பெசோ", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "மலேஷியன் ரிங்கிட்", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "மொசாம்பிகன் மெடிகல்", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "நமீபியன் டாலர்", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "நைஜீரியன் நைரா", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "நிகரகுவன் கோர்டோபா", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "நார்வேஜியன் க்ரோன்", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "நேபாளீஸ் ரூபாய்", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "நியூசிலாந்து டாலர்", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ஓமானி ரியால்", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "பனாமானியன் பால்போவா", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "பெரூவியன் சோல்", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "பபுவா நியூ கினியன் கினா", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "பிலிப்பைன் பெசோ", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "பாகிஸ்தானி ரூபாய்", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "போலிஷ் ஸ்லாட்டி", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "பராகுவன் குவாரானி", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "கத்தாரி ரியால்", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "ரோமானியன் லியூ", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "செர்பியன் தினார்", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "ரஷியன் ரூபிள்", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ருவாண்டன் ஃப்ராங்க்", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "சவுதி ரியால்", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "சாலமன் தீவுகள் டாலர்", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "சிசீலோயிஸ் ருபீ", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "சூடானீஸ் பவுண்டு", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "ஸ்வீடிஷ் க்ரோனா", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "சிங்கப்பூர் டாலர்", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "செயின்ட் ஹெலேனா பவுண்டு", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "சியாரா லியோனியன் லியோன்", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "சோமாலி ஷில்லிங்", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "சுரினாமீஸ் டாலர்", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "தெற்கு சூடானீஸ் பவுண்டு", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "சாவ் டோமி மற்றும் பிரின்ஸ்பி டோப்ரா (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "சாவ் டோமி மற்றும் பிரின்ஸ்பி டோப்ரா", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "சிரியன் பவுண்டு", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "சுவாஸி லிலாங்கனி", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "தாய் பாட்", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "தஜிகிஸ்தானி சோமோனி", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "துர்க்மெனிஸ்தானி மனத்", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "துனிஷியன் தினார்", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "தொங்கான் பங்கா", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "துருக்கிஷ் லீரா", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "டிரினிடாட் மற்றும் டோபாகோ டாலர்", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "புதிய தைவான் டாலர்", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "தான்சானியன் ஷில்லிங்", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "உக்ரைனியன் ஹிரைவ்னியா", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "உகாண்டன் ஷில்லிங்", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "அமெரிக்க டாலர்", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "உருகுவேயன் பெசோ", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "உஸ்பெக்கிஸ்தானி சோம்", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "வெனிசுலன் போலிவர் (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "வெனிசுலன் போலிவர்", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "வியட்நாமீஸ் டாங்", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "வனுவாட்டு வாட்டு", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "சமோவான் தாலா", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "மத்திய ஆப்பிரிக்க CFA ஃப்ராங்க்", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "கிழக்கு கரீபியன் டாலர்", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "மேற்கு ஆப்பிரிக்க CFA ஃப்ராங்க்", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "ஃப்ராங்க் (CFP)", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "தெரியாத நாணயம்", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "ஏமனி ரியால்", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "தென் ஆப்ரிக்க ராண்ட்", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "ஸாம்பியன் குவாசா (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "ஸாம்பியன் குவாச்சா", Symbol: "ZMW"},
			},
		},
	}
}
