// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_am() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "am",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "ጂ ኤም ቲ{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ጃንዩ", Feb: "ፌብሩ", Mar: "ማርች", Apr: "ኤፕሪ", May: "ሜይ", Jun: "ጁን", Jul: "ጁላይ", Aug: "ኦገስ", Sep: "ሴፕቴ", Oct: "ኦክቶ", Nov: "ኖቬም", Dec: "ዲሴም"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ጃ", Feb: "ፌ", Mar: "ማ", Apr: "ኤ", May: "ሜ", Jun: "ጁ", Jul: "ጁ", Aug: "ኦ", Sep: "ሴ", Oct: "ኦ", Nov: "ኖ", Dec: "ዲ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ጃንዩወሪ", Feb: "ፌብሩወሪ", Mar: "ማርች", Apr: "ኤፕሪል", May: "ሜይ", Jun: "ጁን", Jul: "ጁላይ", Aug: "ኦገስት", Sep: "ሴፕቴምበር", Oct: "ኦክቶበር", Nov: "ኖቬምበር", Dec: "ዲሴምበር"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "እሑድ", Mon: "ሰኞ", Tue: "ማክሰ", Wed: "ረቡዕ", Thu: "ሐሙስ", Fri: "ዓርብ", Sat: "ቅዳሜ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "እ", Mon: "ሰ", Tue: "ማ", Wed: "ረ", Thu: "ሐ", Fri: "ዓ", Sat: "ቅ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "እ", Mon: "ሰ", Tue: "ማ", Wed: "ረ", Thu: "ሐ", Fri: "ዓ", Sat: "ቅ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "እሑድ", Mon: "ሰኞ", Tue: "ማክሰኞ", Wed: "ረቡዕ", Thu: "ሐሙስ", Fri: "ዓርብ", Sat: "ቅዳሜ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ጥዋት", PM: "ከሰዓት"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ጠ", PM: "ከ"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ጥዋት", PM: "ከሰዓት"}}}},
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
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}፣{1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "አፋርኛ",
			language.AB:      "አብሐዚኛ",
			language.ACE:     "አቻይንኛ",
			language.ACH:     "አኮሊኛ",
			language.ADA:     "አዳንግሜ",
			language.ADY:     "አድይግሄ",
			language.AE:      "አቬስታን",
			language.AF:      "አፍሪካንኛ",
			language.AFH:     "አፍሪሂሊ",
			language.AGQ:     "አገም",
			language.AIN:     "አይኑ",
			language.AK:      "አካንኛ",
			language.AKK:     "አካዲያን",
			language.AKZ:     "አላባማ",
			language.ALE:     "አልዩት",
			language.ALT:     "ደቡባዊ አልታይ",
			language.AM:      "አማርኛ",
			language.AN:      "አራጎንስ",
			language.ANP:     "አንጊካ",
			language.AR:      "ዓረብኛ",
			language.AR_001:  "ዘመናዊ መደበኛ ዓረብኛ",
			language.ARC:     "አራማይክ",
			language.ARN:     "ማፑቼ",
			language.ARO:     "አራኦና",
			language.ARP:     "አራፓሆ",
			language.ARQ:     "የአልጄሪያ ዓረብኛ",
			language.ARW:     "አራዋክ",
			language.AS:      "አሳሜዛዊ",
			language.ASA:     "አሱ",
			language.ASE:     "የአሜሪካ የምልክት ቋንቋ",
			language.AST:     "አስቱሪያን",
			language.AV:      "አቫሪክ",
			language.AWA:     "አዋድሂ",
			language.AY:      "አያማርኛ",
			language.AZ:      "አዜሪ",
			language.BA:      "ባስኪርኛ",
			language.BAL:     "ባሉቺ",
			language.BAN:     "ባሊኔስ",
			language.BAR:     "ባቫሪያን",
			language.BAS:     "ባሳ",
			language.BAX:     "ባሙን",
			language.BBC:     "ባታካ ቶባ",
			language.BE:      "ቤላራሻኛ",
			language.BEJ:     "ቤጃ",
			language.BEM:     "ቤምባ",
			language.BEW:     "ቤታዊ",
			language.BEZ:     "ቤና",
			language.BFD:     "ባፉት",
			language.BFQ:     "ባዳጋ",
			language.BG:      "ቡልጋሪኛ",
			language.BGN:     "የምዕራብ ባሎቺ",
			language.BHO:     "ቦጁሪ",
			language.BI:      "ቢስላምኛ",
			language.BIK:     "ቢኮል",
			language.BIN:     "ቢኒ",
			language.BJN:     "ባንጃር",
			language.BLA:     "ሲክሲካ",
			language.BM:      "ባምባርኛ",
			language.BN:      "ቤንጋሊኛ",
			language.BO:      "ቲቤታንኛ",
			language.BPY:     "ቢሹንፑሪያ",
			language.BQI:     "ባክህቲያሪ",
			language.BR:      "ብሬቶንኛ",
			language.BRA:     "ብራጅ",
			language.BRH:     "ብራሁዪ",
			language.BRX:     "ቦዶ",
			language.BS:      "ቦስኒያንኛ",
			language.BSS:     "አኮስ",
			language.BUA:     "ቡሪያት",
			language.BUG:     "ቡጊኔዝ",
			language.BUM:     "ቡሉ",
			language.BYN:     "ብሊን",
			language.CA:      "ካታላንኛ",
			language.CAD:     "ካዶ",
			language.CAR:     "ካሪብ",
			language.CAY:     "ካዩጋ",
			language.CCH:     "አትሳም",
			language.CCP:     "ቻክማ",
			language.CE:      "ችችን",
			language.CEB:     "ካቡዋኖ",
			language.CGG:     "ቺጋኛ",
			language.CH:      "ቻሞሮ",
			language.CHB:     "ቺብቻ",
			language.CHG:     "ቻጋታይ",
			language.CHK:     "ቹክስ",
			language.CHM:     "ማሪ",
			language.CHN:     "ቺኑክ ጃርጎን",
			language.CHO:     "ቾክታዋ",
			language.CHP:     "ቺፔውያን",
			language.CHR:     "ቼሮኬኛ",
			language.CHY:     "ችዬኔ",
			language.CKB:     "የሶራኒ ኩርድኛ",
			language.CO:      "ኮርሲካኛ",
			language.COP:     "ኮፕቲክ",
			language.CPS:     "ካፒዝኖን",
			language.CR:      "ክሪ",
			language.CRH:     "ክሪሚያን ተርኪሽ",
			language.CRS:     "ሰሰላዊ ክሬኦሊ ፈረንሳይኛ",
			language.CS:      "ቼክኛ",
			language.CU:      "ቸርች ስላቪክ",
			language.CV:      "ቹቫሽ",
			language.CY:      "ወልሽ",
			language.DA:      "ዴኒሽ",
			language.DAK:     "ዳኮታ",
			language.DAR:     "ዳርግዋ",
			language.DAV:     "ታይታኛ",
			language.DE:      "ጀርመን",
			language.DE_AT:   "የኦስትሪያ ጀርመን",
			language.DE_CH:   "የስዊዝ ከፍተኛ ጀርመንኛ",
			language.DEL:     "ዳላዌር",
			language.DGR:     "ዶግሪብ",
			language.DIN:     "ዲንካ",
			language.DJE:     "ዛርማኛ",
			language.DOI:     "ዶግሪ",
			language.DSB:     "የታችኛው ሰርቢያንኛ",
			language.DTP:     "ሴንተራል ዱሰን",
			language.DUA:     "ዱዋላኛ",
			language.DV:      "ዲቬሂ",
			language.DYO:     "ጆላ ፎንያኛ",
			language.DYU:     "ድዩላ",
			language.DZ:      "ድዞንግኻኛ",
			language.DZG:     "ዳዛጋ",
			language.EBU:     "ኢቦኛ",
			language.EE:      "ኢዊ",
			language.EFI:     "ኤፊክ",
			language.EGY:     "የጥንታዊ ግብጽኛ",
			language.EKA:     "ኤካጁክ",
			language.EL:      "ግሪክኛ",
			language.EN:      "እንግሊዝኛ",
			language.EN_AU:   "የአውስትራሊያ እንግሊዝኛ",
			language.EN_CA:   "የካናዳ እንግሊዝኛ",
			language.EN_GB:   "የዩናይትድ ኪንግደም እንግሊዝኛ",
			language.EN_US:   "የዩ ኤስ እንግሊዝኛ",
			language.EO:      "ኤስፐራንቶ",
			language.ES:      "ስፓንሽኛ",
			language.ES_419:  "የላቲን አሜሪካ ስፓኒሽ",
			language.ES_ES:   "የአውሮፓ ስፓንሽኛ",
			language.ES_MX:   "የሜክሲኮ ስፓንሽኛ",
			language.ESU:     "ሴንተራል ዩፒክ",
			language.ET:      "ኢስቶኒያንኛ",
			language.EU:      "ባስክኛ",
			language.EWO:     "ኤዎንዶ",
			language.FA:      "ፐርሺያኛ",
			language.FA_AF:   "ዳሪኛ",
			language.FF:      "ፉላህ",
			language.FI:      "ፊኒሽ",
			language.FIL:     "ፊሊፒንኛ",
			language.FJ:      "ፊጂኛ",
			language.FO:      "ፋሮኛ",
			language.FON:     "ፎን",
			language.FR:      "ፈረንሳይኛ",
			language.FR_CA:   "የካናዳ ፈረንሳይኛ",
			language.FR_CH:   "የስዊዝ ፈረንሳይኛ",
			language.FRC:     "ካጁን ፍሬንች",
			language.FRP:     "አርፒታን",
			language.FUR:     "ፍሩሊያን",
			language.FY:      "ምዕራባዊ ፍሪሲኛ",
			language.GA:      "አይሪሽ",
			language.GAA:     "ጋ",
			language.GAG:     "ጋጉዝኛ",
			language.GAN:     "ጋን ቻይንኛ",
			language.GD:      "የስኮቲሽ ጌልክኛ",
			language.GEZ:     "ግዕዝኛ",
			language.GIL:     "ጅልበርትስ",
			language.GL:      "ጋሊሺያ",
			language.GN:      "ጓራኒኛ",
			language.GOR:     "ጎሮንታሎ",
			language.GRC:     "የጥንታዊ ግሪክ",
			language.GSW:     "የስዊዝ ጀርመን",
			language.GU:      "ጉጃርቲኛ",
			language.GUZ:     "ጉስሊኛ",
			language.GV:      "ማንክስኛ",
			language.GWI:     "ግዊቺን",
			language.HA:      "ሃውሳኛ",
			language.HAK:     "ሃካ ቻይንኛ",
			language.HAW:     "ሃዊያኛ",
			language.HE:      "ዕብራይስጥ\ufeff",
			language.HI:      "ሒንዱኛ",
			language.HIL:     "ሂሊጋይኖን",
			language.HMN:     "ህሞንግ",
			language.HR:      "ክሮሽያንኛ",
			language.HSB:     "የላይኛው ሶርቢያንኛ",
			language.HSN:     "ዢያንግ ቻይንኛ",
			language.HT:      "ሃይትኛ",
			language.HU:      "ሀንጋሪኛ",
			language.HUP:     "ሁፓ",
			language.HY:      "አርመናዊ",
			language.HZ:      "ሄሬሮ",
			language.IA:      "ኢንቴርሊንጓ",
			language.IBA:     "ኢባን",
			language.IBB:     "ኢቢቦ",
			language.ID:      "ኢንዶኔዥኛ",
			language.IE:      "እንተርሊንግወ",
			language.IG:      "ኢግቦኛ",
			language.II:      "ሲቹንዪኛ",
			language.IK:      "እኑፒያቅኛ",
			language.ILO:     "ኢሎኮ",
			language.INH:     "ኢንጉሽ",
			language.IO:      "ኢዶ",
			language.IS:      "አይስላንድኛ",
			language.IT:      "ጣሊያንኛ",
			language.IU:      "እኑክቲቱትኛ",
			language.JA:      "ጃፓንኛ",
			language.JBO:     "ሎጅባን",
			language.JGO:     "ንጎባኛ",
			language.JMC:     "ማቻሜኛ",
			language.JV:      "ጃቫንኛ",
			language.KA:      "ጆርጂያን",
			language.KAB:     "ካብይል",
			language.KAC:     "ካቺን",
			language.KAJ:     "ካጅ",
			language.KAM:     "ካምባ",
			language.KBD:     "ካባርዲያን",
			language.KCG:     "ታያፕ",
			language.KDE:     "ማኮንዴ",
			language.KEA:     "ካቡቨርዲያኑ",
			language.KFO:     "ኮሮ",
			language.KG:      "ኮንጎኛ",
			language.KHA:     "ክሃሲ",
			language.KHQ:     "ኮይራ ቺኒ",
			language.KI:      "ኪኩዩ",
			language.KJ:      "ኩንያማ",
			language.KK:      "ካዛክኛ",
			language.KKJ:     "ካኮ",
			language.KL:      "ካላሊሱትኛ",
			language.KLN:     "ካለንጂን",
			language.KM:      "ክህመርኛ",
			language.KMB:     "ኪምቡንዱ",
			language.KN:      "ካናዳኛ",
			language.KO:      "ኮሪያኛ",
			language.KOI:     "ኮሚ ፔርምያክ",
			language.KOK:     "ኮንካኒ",
			language.KPE:     "ክፔሌ",
			language.KR:      "ካኑሪ",
			language.KRC:     "ካራቻይ-ባልካር",
			language.KRL:     "ካረሊኛ",
			language.KRU:     "ኩሩክ",
			language.KS:      "ካሽሚርኛ",
			language.KSB:     "ሻምባላ",
			language.KSF:     "ባፊያ",
			language.KSH:     "ኮሎኝኛ",
			language.KU:      "ኩርድሽኛ",
			language.KUM:     "ኩማይክ",
			language.KV:      "ኮሚ",
			language.KW:      "ኮርኒሽ",
			language.KY:      "ኪርጊዝኛ",
			language.LA:      "ላቲንኛ",
			language.LAD:     "ላዲኖ",
			language.LAG:     "ላንጊ",
			language.LB:      "ሉክዘምበርኛ",
			language.LEZ:     "ሌዝጊያን",
			language.LG:      "ጋንዳኛ",
			language.LI:      "ሊምቡርጊሽ",
			language.LKT:     "ላኮታ",
			language.LN:      "ሊንጋላኛ",
			language.LO:      "ላኦኛ",
			language.LOZ:     "ሎዚኛ",
			language.LRC:     "ሰሜናዊ ሉሪ",
			language.LT:      "ሉቴንያንኛ",
			language.LU:      "ሉባ ካታንጋ",
			language.LUA:     "ሉባ-ሉሏ",
			language.LUN:     "ሉንዳ",
			language.LUO:     "ሉኦ",
			language.LUS:     "ሚዞ",
			language.LUY:     "ሉዪያ",
			language.LV:      "ላትቪያን",
			language.MAD:     "ማዱረስ",
			language.MAG:     "ማጋሂ",
			language.MAI:     "ማይተሊ",
			language.MAK:     "ማካሳር",
			language.MAS:     "ማሳይ",
			language.MDF:     "ሞክሻ",
			language.MEN:     "ሜንዴ",
			language.MER:     "ሜሩ",
			language.MFE:     "ሞሪሲየኛ",
			language.MG:      "ማላጋስኛ",
			language.MGH:     "ማኩዋ ሜቶ",
			language.MGO:     "ሜታ",
			language.MH:      "ማርሻሌዝኛ",
			language.MI:      "ማኦሪኛ",
			language.MIC:     "ሚክማክ",
			language.MIN:     "ሚናንግካባኡ",
			language.MK:      "ማሴዶንኛ",
			language.ML:      "ማላያላምኛ",
			language.MN:      "ሞንጎሊያኛ",
			language.MNI:     "ማኒፑሪ",
			language.MOH:     "ሞሃውክ",
			language.MOS:     "ሞሲ",
			language.MR:      "ማራቲኛ",
			language.MS:      "ማላይኛ",
			language.MT:      "ማልቲስኛ",
			language.MUA:     "ሙንዳንግ",
			language.MUL:     "ባለብዙ ቋንቋዎች",
			language.MUS:     "ክሪክ",
			language.MWL:     "ሚራንዴዝኛ",
			language.MY:      "ቡርማኛ",
			language.MYV:     "ኤርዝያ",
			language.MZN:     "ማዛንደራኒ",
			language.NA:      "ናኡሩ",
			language.NAN:     "ሚን ኛን ቻይንኛ",
			language.NAP:     "ኒአፖሊታን",
			language.NAQ:     "ናማ",
			language.NB:      "የኖርዌይ ቦክማል",
			language.ND:      "ሰሜን ንዴብሌ",
			language.NDS:     "የታችኛው ጀርመን",
			language.NDS_NL:  "የታችኛው ሳክሰን",
			language.NE:      "ኔፓሊኛ",
			language.NEW:     "ኒዋሪ(ኔፓል)",
			language.NG:      "ንዶንጋ",
			language.NIA:     "ኒአስ",
			language.NIU:     "ኒዩአንኛ",
			language.NJO:     "ኦ ናጋ",
			language.NL:      "ደች",
			language.NL_BE:   "ፍሌሚሽ",
			language.NMG:     "ክዋሲዮ",
			language.NN:      "የኖርዌይ ናይኖርስክ",
			language.NNH:     "ኒጊምቡን",
			language.NO:      "ኖርዌጂያን",
			language.NOG:     "ኖጋይ",
			language.NQO:     "ንኮ",
			language.NR:      "ደቡብ ንደቤሌ",
			language.NSO:     "ሰሜናዊ ሶቶ",
			language.NUS:     "ኑዌር",
			language.NV:      "ናቫጆ",
			language.NWC:     "ክላሲክ ኔዋሪ",
			language.NY:      "ንያንጃ",
			language.NYN:     "ኒያንኮልኛ",
			language.OC:      "ኦኪታንኛ",
			language.OM:      "ኦሮሞኛ",
			language.OR:      "ኦዲያኛ",
			language.OS:      "ኦሴቲክ",
			language.PA:      "ፑንጃብኛ",
			language.PAG:     "ፓንጋሲናንኛ",
			language.PAM:     "ፓምፓንጋ",
			language.PAP:     "ፓፒአሜንቶ",
			language.PAU:     "ፓላኡአን",
			language.PCM:     "የናይጄሪያ ፒጂን",
			language.PL:      "ፖሊሽኛ",
			language.PRG:     "ፐሩሳንኛ",
			language.PS:      "ፑሽቶ",
			language.PT:      "ፖርቹጋልኛ",
			language.PT_BR:   "የብራዚል ፖርቹጋልኛ",
			language.PT_PT:   "የአውሮፓ ፖርቹጋልኛ",
			language.QU:      "ኵቿኛ",
			language.QUC:     "ኪቼ",
			language.QUG:     "ቺምቦራዞ ሃይላንድ ኩቹዋ",
			language.RAP:     "ራፓኑኢ",
			language.RAR:     "ራሮቶንጋ",
			language.RM:      "ሮማንሽ",
			language.RN:      "ሩንዲኛ",
			language.RO:      "ሮማኒያን",
			language.RO_MD:   "ሞልዳቪያንኛ",
			language.ROF:     "ሮምቦ",
			language.ROOT:    "ሩት",
			language.RU:      "ራሽያኛ",
			language.RUP:     "አሮማንያን",
			language.RW:      "ኪንያርዋንድኛ",
			language.RWK:     "ርዋ",
			language.SA:      "ሳንስክሪትኛ",
			language.SAD:     "ሳንዳዌ",
			language.SAH:     "ሳክሃ",
			language.SAQ:     "ሳምቡሩ",
			language.SAT:     "ሳንታሊ",
			language.SBA:     "ንጋምባይ",
			language.SBP:     "ሳንጉ",
			language.SC:      "ሳርዲንያንኛ",
			language.SCN:     "ሲሲሊያንኛ",
			language.SCO:     "ስኮትስ",
			language.SD:      "ሲንድሂኛ",
			language.SDH:     "ደቡባዊ ኩርዲሽ",
			language.SE:      "ሰሜናዊ ሳሚ",
			language.SEH:     "ሴና",
			language.SES:     "ኮይራቦሮ ሴኒ",
			language.SG:      "ሳንጎኛ",
			language.SH:      "ሰርቦ-ክሮኤሽያኛ",
			language.SHI:     "ታቼልሂት",
			language.SHN:     "ሻን",
			language.SHU:     "ቻዲያን ዓረብኛ",
			language.SI:      "ሲንሃልኛ",
			language.SID:     "ሲዳምኛ",
			language.SK:      "ስሎቫክኛ",
			language.SL:      "ስሎቪኛ",
			language.SM:      "ሳሞአኛ",
			language.SMA:     "ደቡባዊ ሳሚ",
			language.SMJ:     "ሉሌ ሳሚ",
			language.SMN:     "ኢናሪ ሳሚ",
			language.SMS:     "ስኮልት ሳሚ",
			language.SN:      "ሾናኛ",
			language.SNK:     "ሶኒንኬ",
			language.SO:      "ሱማልኛ",
			language.SQ:      "አልባንያንኛ",
			language.SR:      "ሰርብያኛ",
			language.SRN:     "ስራናን ቶንጎ",
			language.SS:      "ስዋቲኛ",
			language.SSY:     "ሳሆኛ",
			language.ST:      "ደቡባዊ ሶቶ",
			language.SU:      "ሱዳንኛ",
			language.SUK:     "ሱኩማ",
			language.SV:      "ስዊድንኛ",
			language.SW:      "ስዋሂሊኛ",
			language.SW_CD:   "ኮንጎ ስዋሂሊ",
			language.SWB:     "ኮሞሪያን",
			language.SYC:     "ክላሲክ ኔይራ",
			language.SYR:     "ሲሪያክ",
			language.TA:      "ታሚልኛ",
			language.TE:      "ተሉጉኛ",
			language.TEM:     "ቲምኔ",
			language.TEO:     "ቴሶ",
			language.TET:     "ቴተም",
			language.TG:      "ታጂኪኛ",
			language.TH:      "ታይኛ",
			language.TI:      "ትግርኛ",
			language.TIG:     "ትግረ",
			language.TK:      "ቱርክሜንኛ",
			language.TL:      "ታጋሎገኛ",
			language.TLH:     "ክሊንጎንኛ",
			language.TN:      "ጽዋናዊኛ",
			language.TO:      "ቶንጋኛ",
			language.TPI:     "ቶክ ፒሲን",
			language.TR:      "ቱርክኛ",
			language.TRV:     "ታሮኮ",
			language.TS:      "ጾንጋኛ",
			language.TT:      "ታታርኛ",
			language.TUM:     "ቱምቡካ",
			language.TVL:     "ቱቫሉ",
			language.TW:      "ትዊኛ",
			language.TWQ:     "ታሳዋቅ",
			language.TY:      "ታሂታንኛ",
			language.TYV:     "ቱቪንያንኛ",
			language.TZM:     "መካከለኛው አትላስ ታማዚኛ",
			language.UDM:     "ኡድሙርት",
			language.UG:      "ኡይገር",
			language.UK:      "ዩክሬንኛ",
			language.UMB:     "ኡምቡንዱ",
			language.UND:     "ያልታወቀ ቋንቋ",
			language.UR:      "ኡርዱኛ",
			language.UZ:      "ኡዝቤክኛ",
			language.VAI:     "ቫይ",
			language.VE:      "ቬንዳ",
			language.VI:      "ቪየትናምኛ",
			language.VO:      "ቮላፑክኛ",
			language.VUN:     "ቩንጆ",
			language.WA:      "ዋሎን",
			language.WAE:     "ዋልሰር",
			language.WAL:     "ወላይትኛ",
			language.WAR:     "ዋራይ",
			language.WBP:     "ዋርልፒሪ",
			language.WO:      "ዎሎፍኛ",
			language.WUU:     "ዉ ቻይንኛ",
			language.XAL:     "ካልማይክ",
			language.XH:      "ዞሳኛ",
			language.XOG:     "ሶጋ",
			language.YAV:     "ያንግቤንኛ",
			language.YBB:     "የምባ",
			language.YI:      "ይዲሽኛ",
			language.YO:      "ዮሩባዊኛ",
			language.YUE:     "ቻይና፤ ካንቶንኛ",
			language.ZA:      "ዡዋንግኛ",
			language.ZBL:     "ብሊስይምቦልስ",
			language.ZGH:     "መደበኛ የሞሮኮ ታማዚግት",
			language.ZH:      "ማንድሪን ቻይንኛ",
			language.ZH_HANS: "ቀለል ያለ ማንድሪን ቻይንኛ",
			language.ZH_HANT: "ባህላዊ ማንድሪን ቻይንኛ",
			language.ZU:      "ዙሉኛ",
			language.ZUN:     "ዙኒ",
			language.ZXX:     "ቋንቋዊ ይዘት አይደለም",
			language.ZZA:     "ዛዛ",
		},
		Territories: cldr.Territories{
			territory.V_001: "ዓለም",
			territory.V_002: "አፍሪካ",
			territory.V_003: "ሰሜን አሜሪካ",
			territory.V_005: "ደቡብ አሜሪካ",
			territory.V_009: "ኦሽኒአ",
			territory.V_011: "ምስራቃዊ አፍሪካ",
			territory.V_013: "መካከለኛው አሜሪካ",
			territory.V_014: "ምዕራባዊ አፍሪካ",
			territory.V_015: "ሰሜናዊ አፍሪካ",
			territory.V_017: "መካከለኛው አፍሪካ",
			territory.V_018: "ደቡባዊ አፍሪካ",
			territory.V_019: "አሜሪካ",
			territory.V_021: "ሰሜናዊ አሜሪካ",
			territory.V_029: "ካሪቢያን",
			territory.V_030: "ምስራቃዊ እስያ",
			territory.V_034: "ደቡባዊ እሲያ",
			territory.V_035: "ምዕራባዊ ደቡብ እሲያ",
			territory.V_039: "ደቡባዊ አውሮፓ",
			territory.V_053: "አውስትራሌዥያ",
			territory.V_054: "ሜላኔዥያ",
			territory.V_057: "የማይክሮኔዥያን ክልል",
			territory.V_061: "ፖሊኔዥያ",
			territory.V_142: "እሲያ",
			territory.V_143: "መካከለኛው እሲያ",
			territory.V_145: "ምዕራባዊ እስያ",
			territory.V_150: "አውሮፓ",
			territory.V_151: "ምዕራባዊ አውሮፓ",
			territory.V_154: "ሰሜናዊ አውሮፓ",
			territory.V_155: "ምስራቃዊ አውሮፓ",
			territory.V_202: "ከሰሃራ በታች አፍሪካ",
			territory.V_419: "ላቲን አሜሪካ",
			territory.AC:    "አሴንሽን ደሴት",
			territory.AD:    "አንዶራ",
			territory.AE:    "የተባበሩት ዓረብ ኤምሬትስ",
			territory.AF:    "አፍጋኒስታን",
			territory.AG:    "አንቲጓ እና ባሩዳ",
			territory.AI:    "አንጉይላ",
			territory.AL:    "አልባኒያ",
			territory.AM:    "አርሜኒያ",
			territory.AO:    "አንጐላ",
			territory.AQ:    "አንታርክቲካ",
			territory.AR:    "አርጀንቲና",
			territory.AS:    "የአሜሪካ ሳሞአ",
			territory.AT:    "ኦስትሪያ",
			territory.AU:    "አውስትራልያ",
			territory.AW:    "አሩባ",
			territory.AX:    "የአላንድ ደሴቶች",
			territory.AZ:    "አዘርባጃን",
			territory.BA:    "ቦስኒያ እና ሄርዞጎቪኒያ",
			territory.BB:    "ባርቤዶስ",
			territory.BD:    "ባንግላዲሽ",
			territory.BE:    "ቤልጄም",
			territory.BF:    "ቡርኪና ፋሶ",
			territory.BG:    "ቡልጌሪያ",
			territory.BH:    "ባህሬን",
			territory.BI:    "ብሩንዲ",
			territory.BJ:    "ቤኒን",
			territory.BL:    "ቅዱስ በርቴሎሜ",
			territory.BM:    "ቤርሙዳ",
			territory.BN:    "ብሩኒ",
			territory.BO:    "ቦሊቪያ",
			territory.BQ:    "የካሪቢያን ኔዘርላንድስ",
			territory.BR:    "ብራዚል",
			territory.BS:    "ባሃማስ",
			territory.BT:    "ቡህታን",
			territory.BV:    "ቡቬት ደሴት",
			territory.BW:    "ቦትስዋና",
			territory.BY:    "ቤላሩስ",
			territory.BZ:    "በሊዝ",
			territory.CA:    "ካናዳ",
			territory.CC:    "ኮኮስ(ኬሊንግ) ደሴቶች",
			territory.CD:    "ኮንጎ (የዲሞክራቲክ ሪፐብሊክ ኮንጎ)",
			territory.CF:    "ማዕከላዊ አፍሪካ ሪፑብሊክ",
			territory.CG:    "ኮንጎ (ሪፐብሊክ)",
			territory.CH:    "ስዊዘርላንድ",
			territory.CI:    "አይቮሪኮስት",
			territory.CK:    "ኩክ ደሴቶች",
			territory.CL:    "ቺሊ",
			territory.CM:    "ካሜሩን",
			territory.CN:    "ቻይና",
			territory.CO:    "ኮሎምቢያ",
			territory.CP:    "ክሊፐርቶን ደሴት",
			territory.CR:    "ኮስታሪካ",
			territory.CU:    "ኩባ",
			territory.CV:    "ኬፕ ቬርዴ",
			territory.CW:    "ኩራሳዎ",
			territory.CX:    "ክሪስማስ ደሴት",
			territory.CY:    "ሳይፕረስ",
			territory.CZ:    "ቼክ ሪፑብሊክ",
			territory.DE:    "ጀርመን",
			territory.DG:    "ዲዬጎ ጋርሺያ",
			territory.DJ:    "ጂቡቲ",
			territory.DK:    "ዴንማርክ",
			territory.DM:    "ዶሚኒካ",
			territory.DO:    "ዶመኒካን ሪፑብሊክ",
			territory.DZ:    "አልጄሪያ",
			territory.EA:    "ሴኡታና ሜሊላ",
			territory.EC:    "ኢኳዶር",
			territory.EE:    "ኤስቶኒያ",
			territory.EG:    "ግብጽ",
			territory.EH:    "ምዕራባዊ ሳህራ",
			territory.ER:    "ኤርትራ",
			territory.ES:    "ስፔን",
			territory.ET:    "ኢትዮጵያ",
			territory.EU:    "የአውሮፓ ህብረት",
			territory.EZ:    "የአውሮፓ ዞን",
			territory.FI:    "ፊንላንድ",
			territory.FJ:    "ፊጂ",
			territory.FK:    "ፎክላንድ ደሴቶች (ኢስላስ ማልቪናስ)",
			territory.FM:    "ሚክሮኔዢያ",
			territory.FO:    "የፋሮ ደሴቶች",
			territory.FR:    "ፈረንሳይ",
			territory.GA:    "ጋቦን",
			territory.GB:    "ዩኬ",
			territory.GD:    "ግሬናዳ",
			territory.GE:    "ጆርጂያ",
			territory.GF:    "የፈረንሳይ ጉዊአና",
			territory.GG:    "ጉርነሲ",
			territory.GH:    "ጋና",
			territory.GI:    "ጂብራልተር",
			territory.GL:    "ግሪንላንድ",
			territory.GM:    "ጋምቢያ",
			territory.GN:    "ጊኒ",
			territory.GP:    "ጉዋደሉፕ",
			territory.GQ:    "ኢኳቶሪያል ጊኒ",
			territory.GR:    "ግሪክ",
			territory.GS:    "ደቡብ ጆርጂያ እና የደቡብ ሳንድዊች ደሴቶች",
			territory.GT:    "ጉዋቲማላ",
			territory.GU:    "ጉዋም",
			territory.GW:    "ጊኒ ቢሳኦ",
			territory.GY:    "ጉያና",
			territory.HK:    "ሆንግ ኮንግ",
			territory.HM:    "ኽርድ ደሴቶችና ማክዶናልድ ደሴቶች",
			territory.HN:    "ሆንዱራስ",
			territory.HR:    "ክሮኤሽያ",
			territory.HT:    "ሀይቲ",
			territory.HU:    "ሀንጋሪ",
			territory.IC:    "የካናሪ ደሴቶች",
			territory.ID:    "ኢንዶኔዢያ",
			territory.IE:    "አየርላንድ",
			territory.IL:    "እስራኤል",
			territory.IM:    "አይል ኦፍ ማን",
			territory.IN:    "ህንድ",
			territory.IO:    "የብሪታኒያ ህንድ ውቂያኖስ ግዛት",
			territory.IQ:    "ኢራቅ",
			territory.IR:    "ኢራን",
			territory.IS:    "አይስላንድ",
			territory.IT:    "ጣሊያን",
			territory.JE:    "ጀርሲ",
			territory.JM:    "ጃማይካ",
			territory.JO:    "ጆርዳን",
			territory.JP:    "ጃፓን",
			territory.KE:    "ኬንያ",
			territory.KG:    "ኪርጊስታን",
			territory.KH:    "ካምቦዲያ",
			territory.KI:    "ኪሪባቲ",
			territory.KM:    "ኮሞሮስ",
			territory.KN:    "ቅዱስ ኪትስ እና ኔቪስ",
			territory.KP:    "ሰሜን ኮሪያ",
			territory.KR:    "ደቡብ ኮሪያ",
			territory.KW:    "ክዌት",
			territory.KY:    "ካይማን ደሴቶች",
			territory.KZ:    "ካዛኪስታን",
			territory.LA:    "ላኦስ",
			territory.LB:    "ሊባኖስ",
			territory.LC:    "ሴንት ሉቺያ",
			territory.LI:    "ሊችተንስታይን",
			territory.LK:    "ሲሪላንካ",
			territory.LR:    "ላይቤሪያ",
			territory.LS:    "ሌሶቶ",
			territory.LT:    "ሊቱዌኒያ",
			territory.LU:    "ሉክሰምበርግ",
			territory.LV:    "ላትቪያ",
			territory.LY:    "ሊቢያ",
			territory.MA:    "ሞሮኮ",
			territory.MC:    "ሞናኮ",
			territory.MD:    "ሞልዶቫ",
			territory.ME:    "ሞንተኔግሮ",
			territory.MF:    "ሴንት ማርቲን",
			territory.MG:    "ማዳጋስካር",
			territory.MH:    "ማርሻል አይላንድ",
			territory.MK:    "ሰሜን መቄዶንያ",
			territory.ML:    "ማሊ",
			territory.MM:    "ማይናማር(በርማ)",
			territory.MN:    "ሞንጎሊያ",
			territory.MO:    "ማካኡ",
			territory.MP:    "የሰሜናዊ ማሪያና ደሴቶች",
			territory.MQ:    "ማርቲኒክ",
			territory.MR:    "ሞሪቴኒያ",
			territory.MS:    "ሞንትሴራት",
			territory.MT:    "ማልታ",
			territory.MU:    "ሞሪሸስ",
			territory.MV:    "ማልዲቭስ",
			territory.MW:    "ማላዊ",
			territory.MX:    "ሜክሲኮ",
			territory.MY:    "ማሌዢያ",
			territory.MZ:    "ሞዛምቢክ",
			territory.NA:    "ናሚቢያ",
			territory.NC:    "ኒው ካሌዶኒያ",
			territory.NE:    "ኒጀር",
			territory.NF:    "ኖርፎልክ ደሴት",
			territory.NG:    "ናይጄሪያ",
			territory.NI:    "ኒካራጓ",
			territory.NL:    "ኔዘርላንድ",
			territory.NO:    "ኖርዌይ",
			territory.NP:    "ኔፓል",
			territory.NR:    "ናኡሩ",
			territory.NU:    "ኒኡይ",
			territory.NZ:    "ኒው ዚላንድ",
			territory.OM:    "ኦማን",
			territory.PA:    "ፓናማ",
			territory.PE:    "ፔሩ",
			territory.PF:    "የፈረንሳይ ፖሊኔዢያ",
			territory.PG:    "ፓፑዋ ኒው ጊኒ",
			territory.PH:    "ፊሊፒንስ",
			territory.PK:    "ፓኪስታን",
			territory.PL:    "ፖላንድ",
			territory.PM:    "ቅዱስ ፒዬር እና ሚኩኤሎን",
			territory.PN:    "ፒትካኢርን አይስላንድ",
			territory.PR:    "ፖርታ ሪኮ",
			territory.PS:    "ፍልስጥኤም",
			territory.PT:    "ፖርቱጋል",
			territory.PW:    "ፓላው",
			territory.PY:    "ፓራጓይ",
			territory.QA:    "ኳታር",
			territory.QO:    "አውትላይንግ ኦሽንያ",
			territory.RE:    "ሪዩኒየን",
			territory.RO:    "ሮሜኒያ",
			territory.RS:    "ሰርብያ",
			territory.RU:    "ሩስያ",
			territory.RW:    "ሩዋንዳ",
			territory.SA:    "ሳውድአረቢያ",
			territory.SB:    "ሰሎሞን ደሴት",
			territory.SC:    "ሲሼልስ",
			territory.SD:    "ሱዳን",
			territory.SE:    "ስዊድን",
			territory.SG:    "ሲንጋፖር",
			territory.SH:    "ሴንት ሄለና",
			territory.SI:    "ስሎቬኒያ",
			territory.SJ:    "ስቫልባርድ እና ጃን ማየን",
			territory.SK:    "ስሎቫኪያ",
			territory.SL:    "ሴራሊዮን",
			territory.SM:    "ሳን ማሪኖ",
			territory.SN:    "ሴኔጋል",
			territory.SO:    "ሱማሌ",
			territory.SR:    "ሱሪናም",
			territory.SS:    "ደቡብ ሱዳን",
			territory.ST:    "ሳኦ ቶሜ እና ፕሪንሲፔ",
			territory.SV:    "ኤል ሳልቫዶር",
			territory.SX:    "ሲንት ማርተን",
			territory.SY:    "ሲሪያ",
			territory.SZ:    "ስዋዚላንድ",
			territory.TA:    "ትሪስታን ዲ ኩንሃ",
			territory.TC:    "የቱርኮችና የካኢኮስ ደሴቶች",
			territory.TD:    "ቻድ",
			territory.TF:    "የፈረንሳይ ደቡባዊ ግዛቶች",
			territory.TG:    "ቶጐ",
			territory.TH:    "ታይላንድ",
			territory.TJ:    "ታጃኪስታን",
			territory.TK:    "ቶክላው",
			territory.TL:    "ምስራቅ ቲሞር",
			territory.TM:    "ቱርክሜኒስታን",
			territory.TN:    "ቱኒዚያ",
			territory.TO:    "ቶንጋ",
			territory.TR:    "ቱርክ",
			territory.TT:    "ትሪናዳድ እና ቶቤጎ",
			territory.TV:    "ቱቫሉ",
			territory.TW:    "ታይዋን",
			territory.TZ:    "ታንዛኒያ",
			territory.UA:    "ዩክሬን",
			territory.UG:    "ዩጋንዳ",
			territory.UM:    "የዩ ኤስ ጠረፍ ላይ ያሉ ደሴቶች",
			territory.UN:    "የተመ",
			territory.US:    "ዩ ኤስ",
			territory.UY:    "ኡራጓይ",
			territory.UZ:    "ኡዝቤኪስታን",
			territory.VA:    "ቫቲካን ከተማ",
			territory.VC:    "ቅዱስ ቪንሴንት እና ግሬናዲንስ",
			territory.VE:    "ቬንዙዌላ",
			territory.VG:    "የእንግሊዝ ቨርጂን ደሴቶች",
			territory.VI:    "የአሜሪካ ቨርጂን ደሴቶች",
			territory.VN:    "ቬትናም",
			territory.VU:    "ቫኑአቱ",
			territory.WF:    "ዋሊስ እና ፉቱና ደሴቶች",
			territory.WS:    "ሳሞአ",
			territory.XA:    "የሀሰት ትእምርት",
			territory.XB:    "የሀሰት ባለሁለት አቅጣጫ",
			territory.XK:    "ኮሶቮ",
			territory.YE:    "የመን",
			territory.YT:    "ሜይኦቴ",
			territory.ZA:    "ደቡብ አፍሪካ",
			territory.ZM:    "ዛምቢያ",
			territory.ZW:    "ዚምቧቤ",
			territory.ZZ:    "ያልታወቀ ክልል",
		},
	}
}