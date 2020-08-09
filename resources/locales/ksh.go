// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_ksh() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ksh",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, 'dä' d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM. y", Short: "d. M. y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan.", Feb: "Fäb.", Mar: "Mäz.", Apr: "Apr.", May: "Mai", Jun: "Jun.", Jul: "Jul.", Aug: "Ouj.", Sep: "Säp.", Oct: "Okt.", Nov: "Nov.", Dec: "Dez."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "O", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Jannewa", Feb: "Fäbrowa", Mar: "Määz", Apr: "Aprell", May: "Mai", Jun: "Juuni", Jul: "Juuli", Aug: "Oujoß", Sep: "Septämber", Oct: "Oktohber", Nov: "Novämber", Dec: "Dezämber"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Su.", Mon: "Mo.", Tue: "Di.", Wed: "Me.", Thu: "Du.", Fri: "Fr.", Sat: "Sa."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "M", Thu: "D", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Su", Mon: "Mo", Tue: "Di", Wed: "Me", Thu: "Du", Fri: "Fr", Sat: "Sa"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sunndaach", Mon: "Mohndaach", Tue: "Dinnsdaach", Wed: "Metwoch", Thu: "Dunnersdaach", Fri: "Friidaach", Sat: "Samsdaach"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "v.M.", PM: "n.M."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Vörmeddaach", PM: "Nommendaach"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ksh]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "−", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "de vereineschte arraabesche Emiraate ier Dirham", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afjahni", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "albaanesche Lek", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "armeenesche Dram", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "neederlängsch antillesche Jullde", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "angjolaanesche Kwansa", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "ajentiinesche Peeso", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "austraalesche Dollaa", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "arubesche Florin", Symbol: ""},
				currency.AZM: cldr.Currency{DisplayName: "Asserbaidschaani Manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Asserbaidschaani Manat", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "ömtuuschbaa Mark us Boßnije un dä Hächejovvina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbados-Dollaa", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Taka us Bangladäsch", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "buljaaresche Lev", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Bachrainesche Denaa", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "burundesche Frang", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Bermuuda-Dollaa", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Brunei-Dollaa", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Bollivijano", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "brasilljaanesche Real", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "bahama’sche Dollaa", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "bhutanesesche Ngultrum", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Pula us Bozwaana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "wiißrußesche Rubel", Symbol: "р."},
				currency.BYR: cldr.Currency{DisplayName: "wiißrußesche Rubel (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "belizjaanesche Dollaa", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "kannaadesche Dollaa", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "kongjoleesesche Frang", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "schweijzer Fränkli", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "schileenesche Peeso", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "schineesesche Yuan Renminbi", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "kolumbesche Peso", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "koßtarikaanesche Colón", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "ömtuuschbaa kubaanesche Pesos", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "kubaanesche Peesos", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "kapverdesche Eskuudos", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "tschäschesche Kruhne", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Frang uß Dschibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "dänesche Kruhne", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "dommenikaanesche Peesos", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "aljeresche Denaa", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Äßnesche Kruhne", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "äjiptesche Pongk", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nafka uß Erritreja", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ätejoopesche Birr", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Fidschi-Dollaa", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Pongk vun de Falkland-Enselle", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "brittesche £", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "jeorjesche Lari", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "janaaesche Cedi", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "jibraltaa’sche Pongk", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "jambesche Dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Jineea-Frang", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "juatemalesche Quetzal", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Juaana-Dollaa", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Hongkong-Dollaa", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "hondureanesche Lempira", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "krowaatesche Kuna", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "haiitesche Gourde", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "unjarresche Forint", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "indoneesesche Ruupije", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "ißraeelesche Schekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "indesche Ruupije", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "iraakesche Denaa", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "persesche Rial", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "ißländesche Kruhne", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "jamaikaanesche Dollaa", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Jordaanesche Dollaa", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "japaanesche Jen", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Kenjaanesche Schillinge", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "kirjiisesche Som", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "kambodschaanesche Riel", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "kommooresche Frang", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "noodkorejaansche Won", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "söödkorejaansche Won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "kuwaitesche Denaa", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Kaimann-Dollaa", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "kasakesche Tenge", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "loaatesche Kip", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "libaneesesche Pongk", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "schrilankesche Ruupije", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "liberijaanesche Dollaa", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "lesoothesche Loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "littouesche Litas", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "lättesche Lats", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "libesche Denaa", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "marrokaanesche Dirhamm", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "moldaavesche Leu", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "madajaskesche Ariary", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "mazedoonesche Denaa", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "burmeesesche Kyat", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "mongjoolesche Tugrik", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "makaneesesche Pataca", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "mauretanesche Ouguiya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "mauretanesche Ouguiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "maurizjahnesche Ruupije", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "mallediivesche Rufiyaa", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "malaawesche Kwache", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "mexekaanesche Peeso", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "malaisesche Ringgit", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "mosambikaanesche Metical (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "mossambikaanesche Metical", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "namiibesche Dollaa", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "nijerijaanesche Naira", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "nikarajaanesche Córdoba", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "norrweejesche Kruhne", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "nepaleesesche Ruupije", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "neuseeländesche Dollaa", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ommaanesche Rijal", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "pannameesesche Balboa", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "perruaanesche Sol", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "papua neujinejaanesche Kina", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "fillipiinesche Pesos", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "pakestaanesche Ruupije", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "polnesche Złoty", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "parajuaanesche Juarani", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "kataaresche Rijal", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "romäänesche Leu (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "romäänesche Leu", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "särbesche Denaare", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "russesche Ruubel", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ruandesche Frang", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "saudesche Rijal", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "solomonesche Dollaa", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "seischellesche Ruupije", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "sudaneesesche Pongk", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "schweedesche Kruhne", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "singjapurejaanesche Dollaa", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Zint-Hellena-Pongk", Symbol: "£"},
				currency.SKK: cldr.Currency{DisplayName: "ßlovaakesche Kruhne", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "ẞjärra-lejoneesesche Leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "somaalesche Schillenge", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "sürinameesesche Dollaa", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "södsudaneesesche Pongk", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra vun São Tomé un Príncipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra vun São Tomé un Príncipe", Symbol: "Db"},
				currency.SVC: cldr.Currency{DisplayName: "asalvadorejaanesche Cosan", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "süüresche Pund", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "swasiländesche Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "tailändesche Baht", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "tadschikißtaanesche Somoni", Symbol: ""},
				currency.TMM: cldr.Currency{DisplayName: "turkmeneßtaanesche Manat (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "turkmeneßtaanesche Manat", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "tuneesesche Denaa", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "tongjanes Paʻangache", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "törkesche Liire", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Dollaa uß Trinidad un Tobääjo", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "neu taiwaneesesche Dollaa", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "tansaanesche Schillenge", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "ukraijnesche Hryvnia", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "ujandesche Schillenge", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "ammärrikaanesche Dollaa", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "urrujuwaische Peeso", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "ußbeekesche Som", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "venezuelaanesche Bolívar (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "venezuelaanesche Bolívar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "vijätnammeesesche Dong", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "vanuatesche Vatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "samowaanesche Tala", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Kmmeruhner Frang", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Sellver", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Jold", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "oß-karribbesche Dollaa", Symbol: "EC$"},
				currency.XFO: cldr.Currency{DisplayName: "franzüüsesche Joldfranke", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "Frang uß de Älfebeinköß", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Palladijum", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "polineesesche Frang", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Plaatin", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Währong zum Prööfe", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "onbikannte Währong", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "jemenitesche Rijal", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "södaffrekaanesche Rand", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "sambesche Kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "sambesche Kwacha", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "simbabwesche Dollaa (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "simbabwesche Dollaa (2009)", Symbol: ""},
			},
		},
	}
}
