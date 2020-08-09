// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_am_ET() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "ጂ ኤም ቲ{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ጃንዩ", Feb: "ፌብሩ", Mar: "ማርች", Apr: "ኤፕሪ", May: "ሜይ", Jun: "ጁን", Jul: "ጁላይ", Aug: "ኦገስ", Sep: "ሴፕቴ", Oct: "ኦክቶ", Nov: "ኖቬም", Dec: "ዲሴም"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ጃ", Feb: "ፌ", Mar: "ማ", Apr: "ኤ", May: "ሜ", Jun: "ጁ", Jul: "ጁ", Aug: "ኦ", Sep: "ሴ", Oct: "ኦ", Nov: "ኖ", Dec: "ዲ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ጃንዩወሪ", Feb: "ፌብሩወሪ", Mar: "ማርች", Apr: "ኤፕሪል", May: "ሜይ", Jun: "ጁን", Jul: "ጁላይ", Aug: "ኦገስት", Sep: "ሴፕቴምበር", Oct: "ኦክቶበር", Nov: "ኖቬምበር", Dec: "ዲሴምበር"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "እሑድ", Mon: "ሰኞ", Tue: "ማክሰ", Wed: "ረቡዕ", Thu: "ሐሙስ", Fri: "ዓርብ", Sat: "ቅዳሜ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "እ", Mon: "ሰ", Tue: "ማ", Wed: "ረ", Thu: "ሐ", Fri: "ዓ", Sat: "ቅ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "እ", Mon: "ሰ", Tue: "ማ", Wed: "ረ", Thu: "ሐ", Fri: "ዓ", Sat: "ቅ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "እሑድ", Mon: "ሰኞ", Tue: "ማክሰኞ", Wed: "ረቡዕ", Thu: "ሐሙስ", Fri: "ዓርብ", Sat: "ቅዳሜ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ጥዋት", PM: "ከሰዓት"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ጠ", PM: "ከ"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ጥዋት", PM: "ከሰዓት"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_am]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "የተባበሩት የአረብ ኤምሬትስ ድርሀም", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "የአፍጋን አፍጋኒ", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "የአልባንያ ሌክ", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "የአርመን ድራም", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "ኔዘርላንድስ አንቲሊአን ጊልደር", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "የአንጎላ ኩዋንዛ", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "የአርጀንቲና ፔሶ", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "የአውስትራሊያ ዶላር", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "አሩባን ፍሎሪን", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "የአዛርባጃን ማናት", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "የቦስኒያ ሄርዞጎቪና የሚመነዘር ማርክ", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "የባርቤዶስ ዶላር", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "የባንግላዲሽ ታካ", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "የቡልጋሪያ ሌቭ", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "የባኽሬን ዲናር", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "የብሩንዲ ፍራንክ", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "የቤርሙዳ ዶላር", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "የብሩኔ ዶላር", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "የቦሊቪያ ቦሊቪያኖ", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "የብራዚል ሪል", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "የባሃማስ ዶላር", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ብሁታኒዝ ንጉልትረም", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "የቦትስዋና ፑላ", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "የቤላሩስያ ሩብል", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "የቤላሩስያ ሩብል (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "የቤሊዝ ዶላር", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "የካናዳ ዶላር", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "የኮንጐ ፍራንክ ኮንጐሌዝ", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "የስዊስ ፍራንክ", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "የቺሊ ፔሶ", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "የቻይና ዩዋን (የውጭ ምንዛሪ)", Symbol: "የቻይና ዩዋን"},
				currency.CNY: cldr.Currency{DisplayName: "የቻይና የን", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "የኮሎምቢያ ፔሶ", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "የኮስታሪካ ኮሎን", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "የኩባ የሚመነዘር ፔሶ", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "የኩባ ፔሶ", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "የኬፕ ቫርዲ ኤስኩዶ", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "ቼክ ሪፐብሊክ ኮሩና", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "የጅቡቲ ፍራንክ", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "የዴንማርክ ክሮን", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "የዶሚኒክ ፔሶ", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "የአልጄሪያ ዲናር", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "የግብጽ ፓውንድ", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "የኤርትራ ናቅፋ", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "የኢትዮጵያ ብር", Symbol: "ብር"},
				currency.EUR: cldr.Currency{DisplayName: "ዩሮ", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "የፊጂ ዶላር", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "የፎክላንድ ደሴቶች ፓውንድ", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "የእንግሊዝ ፓውንድ ስተርሊንግ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "የጆርጅያ ላሪ", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "የጋና ሴዲ", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "የጋና ሲዲ", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "ጂብራልተር ፓውንድ", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "የጋምቢያ ዳላሲ", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "የጊኒ ፍራንክ", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "ጓቲማላን ኩቲዛል", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "የጉየና ዶላር", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "የሆንግኮንግ ዶላር", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "የሃንዱራ ሌምፓአይራ", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "የክሮሽያ ኩና", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "የሃያቲ ጓርዴ", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "የሃንጋሪያን ፎሪንት", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "የኢንዶኔዥያ ሩፒሃ", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "የእስራኤል አዲስ ሽቅል", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "የሕንድ ሩፒ", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "የኢራቅ ዲናር", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "የኢራን ሪአል", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "የአይስላንድ ክሮና", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "የጃማይካ ዶላር", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "የጆርዳን ዲናር", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "የጃፓን የን", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "የኬኒያ ሺሊንግ", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "የኪርጊስታን ሶም", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "የካምቦዲያ ሬል", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "የኮሞሮ ፍራንክ", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "የሰሜን ኮሪያ ዎን", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "የደቡብ ኮሪያ ዎን", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "የኩዌት ዲናር", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "የካይማን ደሴቶች ዶላር", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "የካዛኪስታን ተንጌ", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "የላኦቲ ኪፕ", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "የሊባኖስ ፓውንድ", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "የሲሪላንካ ሩፒ", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "የላይቤሪያ ዶላር", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "የሌሶቶ ሎቲ", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "ሊቱዌንያን ሊታስ", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "የላቲቫ ላትስ", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "የሊቢያ ዲናር", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "የሞሮኮ ዲርሀም", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "ሞልዶቫን ሊኡ", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "የማደጋስካር ማላጋስይ አሪያርይ", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "የሜቆድንያ ዲናር", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "የማያናማር ክያት", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "የሞንጎሊያን ቱግሪክ", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "የማካኔዝ ፓታካ", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "የሞሪቴኒያ ኦውጉያ (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "የሞሪቴኒያ ኦውጉያ", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "የሞሪሸስ ሩፒ", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "የማልዲቫ ሩፊያ", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "የማላዊ ኩዋቻ", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "የሜክሲኮ ፔሶ", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "የማሌዥያ ሪንጊት", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "የሞዛምቢክ ሜቲካል", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "የናሚቢያ ዶላር", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "የናይጄሪያ ናይራ", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "የኒካራጓ ኮርዶባ", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "የኖርዌይ ክሮን", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "የኔፓል ሩፒ", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "የኒውዚላንድ ዶላር", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "የኦማን ሪአል", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "ፓናማኒአን ባልቦአ", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "የፔሩቪያ ሶል", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "የፓፕዋ ኒው ጊኒ ኪና", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "የፊሊፒንስ ፔሶ", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "የፓኪስታን ሩፒ", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "የፖላንድ ዝሎቲ", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "የፓራጓይ ጉአራኒ", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "የኳታር ሪአል", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "የሮማኒያ ለው", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "የሰርቢያ ዲናር", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "የሩስያ ሩብል", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "የሩዋንዳ ፍራንክ", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "የሳውዲ ሪያል", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "የሰለሞን ደሴቶች ዶላር", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "የሲሼል ሩፒ", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "የሱዳን ፓውንድ", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "የሱዳን ፓውንድ (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "የስዊድን ክሮና", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "የሲንጋፖር ዶላር", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "የሴይንት ሔሌና ፓውንድ", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "የሴራሊዎን ሊዎን", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "የሶማሌ ሺሊንግ", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "የሰርናሜዝ ዶላር", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "የደቡብ ሱዳን ፓውንድ", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "የሳኦ ቶሜ እና ፕሪንሲፔ ዶብራ (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "የሳኦ ቶሜ እና ፕሪንሲፔ ዶብራ", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "የሲሪያ ፓውንድ", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "የስዋዚላንድ ሊላንገኒ", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "የታይላንድ ባህት", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "የታጂክስታን ሶሞኒ", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "ቱርክሜኒስታኒ ማናት", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "የቱኒዚያ ዲናር", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "ቶንጋን ፓ’አንጋ", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "የቱርክ ሊራ", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "የትሪንዳድ እና ቶቤጎዶላር", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "የአዲሷ ታይዋን ዶላር", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "የታንዛኒያ ሺሊንግ", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "የዩክሬን ሀሪይቭኒአ", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "የዩጋንዳ ሺሊንግ", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "የአሜሪካን ዶላር", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "የኡራጓይ ፔሶ", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "የኡዝፔኪስታን ሶም", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "የቬንዝዌላ ቦሊቫር (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "የቬንዝዌላ-ቦሊቫር", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "የቭየትናም ዶንግ", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "የቫንዋንቱ ቫቱ", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "ሳሞአን ታላ", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "የመካከለኛው አፍሪካ ሴፋ ፍራንክ", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "የምዕራብ ካሪብያን ዶላር", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "የምዕራብ አፍሪካ ሴፋ ፍራንክ", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "ሲ ኤፍ ፒ ፍራንክ", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "ያልታወቀ ገንዘብ", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "የየመን ሪአል", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "የደቡብ አፍሪካ ራንድ", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "የዛምቢያ ክዋቻ (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "የዛምቢያ ክዋቻ", Symbol: "ZMW"},
				currency.ZWD: cldr.Currency{DisplayName: "የዚምቧቡዌ ዶላር", Symbol: ""},
			},
		},
	}
}
