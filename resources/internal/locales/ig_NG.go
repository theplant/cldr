// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_ig_NG() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ig_NG",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'na' {0}", Long: "{1} 'na' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jen", Feb: "Feb", Mar: "Maa", Apr: "Epr", May: "Mee", Jun: "Juu", Jul: "Jul", Aug: "Ọgọ", Sep: "Sep", Oct: "Ọkt", Nov: "Nov", Dec: "Dis"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "E", May: "M", Jun: "J", Jul: "J", Aug: "Ọ", Sep: "S", Oct: "Ọ", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Jenụwarị", Feb: "Febrụwarị", Mar: "Maachị", Apr: "Epreel", May: "Mee", Jun: "Juun", Jul: "Julaị", Aug: "Ọgọọst", Sep: "Septemba", Oct: "Ọktoba", Nov: "Novemba", Dec: "Disemba"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ụka", Mon: "Mọn", Tue: "Tiu", Wed: "Wen", Thu: "Tọọ", Fri: "Fraị", Sat: "Sat"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Sọn", Mon: "Mọn", Tue: "Tiu", Wed: "Wen", Thu: "Tọọ", Fri: "Fraị", Sat: "Sat"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sọndee", Mon: "Mọnde", Tue: "Tiuzdee", Wed: "Wenezdee", Thu: "Tọọzdee", Fri: "Fraịdee", Sat: "Satọdee"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "A.M.", PM: "P.M."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "A.M.", PM: "P.M."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "A.M.", PM: "P.M."}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "\u200f-", Percent: "٪\u200f", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "¤#,##0.00", CurrencyAccounting: "", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Ego Dirham obodo United Arab Emirates", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Ego Afghani Obodo Afghanistan", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Ego Lek Obodo Albania", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Ego Dram obodo Armenia", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "Ego Antillean Guilder obodo Netherlands", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Ego Kwanza obodo Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Ego Peso obodo Argentina", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Ego Dollar obodo Australia", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Ego Florin obodo Aruba", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Ego Manat obodo Azerbaijan", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Akara mgbanwe ego obodo Bosnia-Herzegovina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Ego Dollar obodo Barbados", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Ego Taka obodo Bangladesh", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Ego Lev mba Bulgaria", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Ego Dinar Obodo Bahrain", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Ego Franc obodo Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Dollar Bermuda", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Ego Dollar obodo Brunei", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Ego Boliviano obodo Bolivia", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Real Brazil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Ego Dollar Obodo Bahamas", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Ego Ngultrum obodo Bhutan", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Ego Pula obodo Bostwana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Ego Ruble mba Belarus", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "Dollar Belize", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dollar Canada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Ego Franc obodo Congo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Ego Franc mba Switzerland", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Ego Peso obodo Chile", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Ego Yuan Obodo China (ndị bi na mmiri)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Yuan China", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Ego Peso obodo Columbia", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Ego Colón obodo Costa Rica", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Ego Peso e nwere ike ịgbanwe nke obodo Cuba", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Ego Peso obodo Cuba", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo Caboverdiano", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Ego Koruna obodo Czech", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Ego Franc obodo Djibouti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Ego Krone Obodo Denmark", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Ego Peso Obodo Dominica", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Ego Dinar Obodo Algeria", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Ego Pound obodo Egypt", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Ego Nakfa obodo Eritrea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ego Birr obodo Ethiopia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Ego Dollar obodo Fiji", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Ego Pound obodo Falkland Islands", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pound British", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Ego Lari Obodo Georgia", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "Ego Cedi obodo Ghana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Ego Pound obodo Gibraltar", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Ego Dalasi obodo Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Ego Franc obodo Guinea", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "Ego Quetzal obodo Guatemala", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Ego Dollar obodo Guyana", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Ego Dollar Obodo Honk Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Ego Lempira obodo Honduras", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Ego Kuna obodo Croatia", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Ego Gourde obodo Haiti", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Ego Forint obodo Hungary", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Ego Rupiah Obodo Indonesia", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "Ego Shekel ọhụrụ obodo Israel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupee India", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Ego Dinar obodo Iraq", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Ego Rial obodo Iran", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Ego Króna obodo Iceland", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Ego Dollar obodo Jamaica", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Ego Dinar Obodo Jordan", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Yen Japan", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "Ego Shilling obodo Kenya", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Ego Som Obodo Kyrgyzstan", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "Ego Riel obodo Cambodia", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Ego Franc obodo Comoros", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Ego Won Obodo North Korea", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Ego Won Obodo South Korea", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Ego Dinar Obodo Kuwait", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Ego Dollar obodo Cayman Islands", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Ego Tenge obodo Kazakhstani", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Ego Kip Obodo Laos", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Ego Pound obodo Lebanon", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Ego Rupee obodo Sri Lanka", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Ego Dollar obodo Liberia", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Ego Dinar obodo Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Ego Dirham obodo Morocco", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Ego Leu obodo Moldova", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ego Ariary obodo Madagascar", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Ego Denar Obodo Macedonia", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Ego Kyat obodo Myanmar", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Ego Turgik Obodo Mongolia", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Ego Pataca ndị Obodo Macanese", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ego Ouguiya Obodo Mauritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Ego Rupee obodo Mauritania", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Ego Rufiyaa obodo Moldova", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Ego Kwacha obodo Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Ego Peso obodo Mexico", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ego Ringgit obodo Malaysia", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "Ego Metical obodo Mozambique", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Ego Dollar obodo Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naịra", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Ego Córodoba obodo Nicaragua", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Ego Krone Obodo Norway", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Ego Rupee obodo Nepal", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Ego Dollar obodo New Zealand", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Ego Rial obodo Oman", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Ego Balboa obodo Panama", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Ego Sol obodo Peru", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Ego Kina obodo Papua New Guinea", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Ego piso obodo Philippine", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Ego Rupee obodo Pakistan", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Ego Zloty mba Poland", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Ego Guarani obodo Paraguay", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Ego Rial obodo Qatar", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Ego Leu obodo Romania", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Ego Dinar obodo Serbia", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Ruble Russia", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Ego Franc obodo Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Ego Riyal obodo Saudi", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Ego Dollar obodo Solomon Islands", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Ego Rupee obodo Seychelles", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Ego Pound obodo Sudan", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Ego Krona Obodo Sweden", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Ego Dollar obodo Singapore", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Ego Pound obodo St Helena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Ego Leone obodo Sierra Leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Ego shilling obodo Somali", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Dollar Surinamese", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Ego Pound obodo South Sudan", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "Ego Dobra nke obodo Sāo Tomé na Principe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Ego Pound obodo Syria", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Ego Lilangeni obodo Swaziland", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Ego Baht obodo Thai", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Who Somoni obodo Tajikistan", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Ego Manat Obodo Turkmenistan", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Ego Dinar Obodo Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Ego paʻanga obodo Tonga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Ego Lira obodo Turkey", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Dollar Trinidad & Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Dollar obodo New Taiwan", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Ego Shilling Obodo Tanzania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Ego Hryvnia obodo Ukraine", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Ego Shilling obodo Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dollar US", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "Ego Peso obodo Uruguay", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Ego Som obodo Uzbekistan", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Ego Bolivar obodo Venezuela", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Ego Dong obodo Vietnam", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Ego Vatu obodo Vanuatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Ego Tala obodo Samoa", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Ego Franc mba etiti Africa", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Ego Dollar obodo East Carribbean", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Ego CFA Franc obodo West Africa", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Ego Franc obodo CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Ego Amaghị", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Ego Rial obodo Yemeni", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Ego Rand obodo South Africa", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "Ego Kwacha Obodo Zambia", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:      "Akan",
			language.AM:      "Amariikị",
			language.AR:      "Arabiikị",
			language.AR_001:  "Ụdị Arabiikị nke oge a",
			language.BE:      "Belaruusu",
			language.BG:      "Bọlụgarịa",
			language.BN:      "Bengali",
			language.CS:      "Cheekị",
			language.DE:      "Asụsụ Jaman",
			language.DE_AT:   "Jaman ndị Austria",
			language.DE_CH:   "Jaman Izugbe ndị Switzerland",
			language.EL:      "Giriikị",
			language.EN:      "Asụsụ Bekee",
			language.EN_AU:   "Bekee ndị Australia",
			language.EN_CA:   "Bekee ndị Canada",
			language.EN_GB:   "Bekee ndị UK",
			language.EN_US:   "Bekee ndị US",
			language.ES:      "Asụsụ Spanish",
			language.ES_419:  "Asụsụ Spanish ndị Latin America",
			language.ES_ES:   "Asụsụ Spanish ndị Europe",
			language.ES_MX:   "Asụsụ Spanish ndị Mexico",
			language.FA:      "Peshan",
			language.FR:      "Asụsụ Fụrench",
			language.FR_CA:   "Fụrench ndị Canada",
			language.FR_CH:   "Fụrench ndị Switzerland",
			language.HA:      "Awụsa",
			language.HI:      "Hindi",
			language.HU:      "Magịya",
			language.ID:      "Indonisia",
			language.IG:      "Asụsụ Igbo",
			language.IT:      "Asụsụ Italian",
			language.JA:      "Asụsụ Japanese",
			language.JV:      "Java",
			language.KM:      "Keme, Etiti",
			language.KO:      "Koria",
			language.MS:      "Maleyi",
			language.MY:      "Mịanma",
			language.NE:      "Nepali",
			language.NL:      "Dọọch",
			language.PA:      "Punjabi",
			language.PL:      "Poliishi",
			language.PT:      "Asụsụ Portuguese",
			language.PT_BR:   "Asụsụ Portuguese ndị Brazil",
			language.PT_PT:   "Asụsụ Portuguese ndị Europe",
			language.RO:      "Rumenia",
			language.RU:      "Asụsụ Russian",
			language.RW:      "Rụwanda",
			language.SO:      "Somali",
			language.SV:      "Sụwidiishi",
			language.TA:      "Tamụlụ",
			language.TH:      "Taị",
			language.TR:      "Tọkiishi",
			language.UK:      "Ukureenị",
			language.UND:     "Asụsụ amaghị",
			language.UR:      "Urudu",
			language.VI:      "Viyetịnaamụ",
			language.YO:      "Yoruba",
			language.ZH:      "Mandarịịnị",
			language.ZH_HANS: "Asụsụ Chinese dị mfe",
			language.ZH_HANT: "Asụsụ Chinese Izugbe",
			language.ZU:      "Zulu",
		},
		Territories: cldr.Territories{
			territory.V_001: "Uwa",
			territory.V_002: "Afrika",
			territory.V_003: "Mpaghara Ugwu Amerịka",
			territory.V_005: "Mpaghara Mgbada Ugwu America",
			territory.V_009: "Oceania",
			territory.V_011: "Mpaghara Ọdịda Anyanwụ Afrịka",
			territory.V_013: "Etiti America",
			territory.V_014: "Mpaghara Ọwụwa Anyanwụ Afrịka",
			territory.V_015: "Mpaghara Ugwu Afrịka",
			territory.V_017: "Etiti Afrịka",
			territory.V_018: "Mpaghara Mgbada Ugwu Afrịka",
			territory.V_019: "Amerịka",
			territory.V_021: "Mpaghara Ugwu America",
			territory.V_029: "Onye Carrabean",
			territory.V_030: "Mpaghara Ọwụwa Anyanwụ Asia",
			territory.V_034: "Mpaghara Mgbada Ugwu Asia",
			territory.V_035: "Mpaghara Mgbada Ugwu Asia dị na Ọwụwa Anyanwụ",
			territory.V_039: "Mpaghara Mgbada Ugwu Europe",
			territory.V_053: "Australasia",
			territory.V_054: "Melanesia",
			territory.V_057: "Mpaghara Micronesian",
			territory.V_061: "Polynesia",
			territory.V_142: "Asia",
			territory.V_143: "Etiti Asia",
			territory.V_145: "Mpaghara Ọdịda Anyanwụ Asia",
			territory.V_150: "Europe",
			territory.V_151: "Mpaghara Ọwụwa Anyanwụ Europe",
			territory.V_154: "Mpaghara Ugwu Europe",
			territory.V_155: "Mpaghara Ọdịda Anyanwụ Europe",
			territory.V_202: "Sub-Saharan Afrịka",
			territory.V_419: "Latin America",
			territory.AC:    "Ascension Island",
			territory.AD:    "Andorra",
			territory.AE:    "Obodo United Arab Emirates",
			territory.AF:    "Mba Afghanistan",
			territory.AG:    "Antigua & Barbuda",
			territory.AI:    "Anguilla",
			territory.AL:    "Albania",
			territory.AM:    "Obodo Armenia",
			territory.AO:    "Angola",
			territory.AQ:    "Antarctica",
			territory.AR:    "Argentina",
			territory.AS:    "American Samoa",
			territory.AT:    "Austria",
			territory.AU:    "Australia",
			territory.AW:    "Aruba",
			territory.AX:    "Agwaetiti Aland",
			territory.AZ:    "Obodo Azerbaijan",
			territory.BA:    "Bosnia & Herzegovina",
			territory.BB:    "Barbados",
			territory.BD:    "Obodo Bangladesh",
			territory.BE:    "Belgium",
			territory.BF:    "Burkina Faso",
			territory.BG:    "Bulgaria",
			territory.BH:    "Obodo Bahrain",
			territory.BI:    "Burundi",
			territory.BJ:    "Binin",
			territory.BL:    "St. Barthélemy",
			territory.BM:    "Bemuda",
			territory.BN:    "Brunei",
			territory.BO:    "Bolivia",
			territory.BQ:    "Caribbean Netherlands",
			territory.BR:    "Mba Brazil",
			territory.BS:    "Bahamas",
			territory.BT:    "Obodo Bhutan",
			territory.BV:    "Agwaetiti Bouvet",
			territory.BW:    "Botswana",
			territory.BY:    "Belarus",
			territory.BZ:    "Belize",
			territory.CA:    "Kanada",
			territory.CC:    "Agwaetiti Cocos (Keeling)",
			territory.CD:    "Congo - Kinshasa",
			territory.CF:    "Central African Republik",
			territory.CG:    "Congo",
			territory.CH:    "Switzerland",
			territory.CI:    "Côte d’Ivoire",
			territory.CK:    "Agwaetiti Cook",
			territory.CL:    "Chile",
			territory.CM:    "Cameroon",
			territory.CN:    "Mba China",
			territory.CO:    "Colombia",
			territory.CP:    "Agwaetiti Clipperton",
			territory.CR:    "Kosta Rika",
			territory.CU:    "Cuba",
			territory.CV:    "Cape Verde",
			territory.CW:    "Kurakao",
			territory.CX:    "Agwaetiti Christmas",
			territory.CY:    "Obodo Cyprus",
			territory.CZ:    "Czechia",
			territory.DE:    "Mba Germany",
			territory.DG:    "Diego Garcia",
			territory.DJ:    "Djibouti",
			territory.DK:    "Denmark",
			territory.DM:    "Dominika",
			territory.DO:    "Dominican Republik",
			territory.DZ:    "Algeria",
			territory.EA:    "Ceuta & Melilla",
			territory.EC:    "Ecuador",
			territory.EE:    "Estonia",
			territory.EG:    "Egypt",
			territory.EH:    "Ọdịda Anyanwụ Sahara",
			territory.ER:    "Eritrea",
			territory.ES:    "Spain",
			territory.ET:    "Ethiopia",
			territory.EU:    "Otu nzukọ mba Europe",
			territory.EZ:    "Gburugburu Euro",
			territory.FI:    "Finland",
			territory.FJ:    "Fiji",
			territory.FK:    "Agwaetiti Falkland",
			territory.FM:    "Micronesia",
			territory.FO:    "Agwaetiti Faroe",
			territory.FR:    "Mba France",
			territory.GA:    "Gabon",
			territory.GB:    "Mba United Kingdom",
			territory.GD:    "Grenada",
			territory.GE:    "Obodo Georgia",
			territory.GF:    "Frenchi Guiana",
			territory.GG:    "Guernsey",
			territory.GH:    "Ghana",
			territory.GI:    "Gibraltar",
			territory.GL:    "Greenland",
			territory.GM:    "Gambia",
			territory.GN:    "Guinea",
			territory.GP:    "Guadeloupe",
			territory.GQ:    "Equatorial Guinea",
			territory.GR:    "Greece",
			territory.GS:    "South Georgia na Agwaetiti South Sandwich",
			territory.GT:    "Guatemala",
			territory.GU:    "Guam",
			territory.GW:    "Guinea-Bissau",
			territory.GY:    "Guyana",
			territory.HK:    "Honk Kong mba nwere ndozi pụrụ iche n’obodo China",
			territory.HM:    "Agwaetiti Heard na Agwaetiti McDonald",
			territory.HN:    "Honduras",
			territory.HR:    "Croatia",
			territory.HT:    "Hati",
			territory.HU:    "Hungary",
			territory.IC:    "Agwaetiti Kanarị",
			territory.ID:    "Indonesia",
			territory.IE:    "Ireland",
			territory.IL:    "Obodo Israel",
			territory.IM:    "Isle of Man",
			territory.IN:    "Mba India",
			territory.IO:    "British Indian Ocean Territory",
			territory.IQ:    "Obodo Iraq",
			territory.IR:    "Obodo Iran",
			territory.IS:    "Iceland",
			territory.IT:    "Mba Italy",
			territory.JE:    "Jersey",
			territory.JM:    "Jamaika",
			territory.JO:    "Obodo Jordan",
			territory.JP:    "Mba Japan",
			territory.KE:    "Kenya",
			territory.KG:    "Obodo Kyrgyzstan",
			territory.KH:    "Cambodia",
			territory.KI:    "Kiribati",
			territory.KM:    "Comorosu",
			territory.KN:    "St. Kitts & Nevis",
			territory.KP:    "Mba Ugwu Korea",
			territory.KR:    "Mba South Korea",
			territory.KW:    "Obodo Kuwait",
			territory.KY:    "Agwaetiti Cayman",
			territory.KZ:    "Obodo Kazakhstan",
			territory.LA:    "Laos",
			territory.LB:    "Obodo Lebanon",
			territory.LC:    "St. Lucia",
			territory.LI:    "Liechtenstein",
			territory.LK:    "Obodo Sri Lanka",
			territory.LR:    "Liberia",
			territory.LS:    "Lesotho",
			territory.LT:    "Lithuania",
			territory.LU:    "Luxembourg",
			territory.LV:    "Latvia",
			territory.LY:    "Libyia",
			territory.MA:    "Morocco",
			territory.MC:    "Monaco",
			territory.MD:    "Moldova",
			territory.ME:    "Montenegro",
			territory.MF:    "St. Martin",
			territory.MG:    "Madagaskar",
			territory.MH:    "Agwaetiti Marshall",
			territory.MK:    "North Macedonia",
			territory.ML:    "Mali",
			territory.MM:    "Myanmar (Burma)",
			territory.MN:    "Obodo Mongolia",
			territory.MO:    "Obodo Macao nwere ndozi pụrụ iche na mba China",
			territory.MP:    "Agwaetiti Northern Mariana",
			territory.MQ:    "Martinique",
			territory.MR:    "Mauritania",
			territory.MS:    "Montserrat",
			territory.MT:    "Malta",
			territory.MU:    "Mauritius",
			territory.MV:    "Maldivesa",
			territory.MW:    "Malawi",
			territory.MX:    "Mexico",
			territory.MY:    "Malaysia",
			territory.MZ:    "Mozambik",
			territory.NA:    "Namibia",
			territory.NC:    "New Caledonia",
			territory.NE:    "Niger",
			territory.NF:    "Agwaetiti Norfolk",
			territory.NG:    "Naịjịrịa",
			territory.NI:    "Nicaragua",
			territory.NL:    "Netherlands",
			territory.NO:    "Norway",
			territory.NP:    "Obodo Nepal",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "New Zealand",
			territory.OM:    "Obodo Oman",
			territory.PA:    "Panama",
			territory.PE:    "Peru",
			territory.PF:    "Frenchi Polynesia",
			territory.PG:    "Papua New Guinea",
			territory.PH:    "Philippines",
			territory.PK:    "Obodo Pakistan",
			territory.PL:    "Poland",
			territory.PM:    "St. Pierre & Miquelon",
			territory.PN:    "Agwaetiti Pitcairn",
			territory.PR:    "Puerto Rico",
			territory.PS:    "Obodo dị iche iche dị n’okpuru mba Palestine",
			territory.PT:    "Portugal",
			territory.PW:    "Palau",
			territory.PY:    "Paraguay",
			territory.QA:    "Obodo Qatar",
			territory.QO:    "Outlying Oceania",
			territory.RE:    "Réunion",
			territory.RO:    "Romania",
			territory.RS:    "Serbia",
			territory.RU:    "Mba Russia",
			territory.RW:    "Rwanda",
			territory.SA:    "Obodo Saudi Arabia",
			territory.SB:    "Agwaetiti Solomon",
			territory.SC:    "Seychelles",
			territory.SD:    "Sudan",
			territory.SE:    "Sweden",
			territory.SG:    "Singapore",
			territory.SH:    "St. Helena",
			territory.SI:    "Slovenia",
			territory.SJ:    "Svalbard & Jan Mayen",
			territory.SK:    "Slovakia",
			territory.SL:    "Sierra Leone",
			territory.SM:    "San Marino",
			territory.SN:    "Senegal",
			territory.SO:    "Somalia",
			territory.SR:    "Suriname",
			territory.SS:    "South Sudan",
			territory.ST:    "São Tomé & Príncipe",
			territory.SV:    "El Salvador",
			territory.SX:    "Sint Maarten",
			territory.SY:    "Obodo Syria",
			territory.SZ:    "Eswatini",
			territory.TA:    "Tristan da Cunha",
			territory.TC:    "Agwaetiti Turks na Caicos",
			territory.TD:    "Chad",
			territory.TF:    "Ụmụ ngalaba Frenchi Southern",
			territory.TG:    "Togo",
			territory.TH:    "Thailand",
			territory.TJ:    "Obodo Tajikistan",
			territory.TK:    "Tokelau",
			territory.TL:    "Timor-Leste",
			territory.TM:    "Obodo Turkmenistan",
			territory.TN:    "Tunisia",
			territory.TO:    "Tonga",
			territory.TR:    "Obodo Turkey",
			territory.TT:    "Trinidad & Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Obodo Taiwan",
			territory.TZ:    "Tanzania",
			territory.UA:    "Ukraine",
			territory.UG:    "Uganda",
			territory.UM:    "Obere Agwaetiti Dị Na Mpụga U.S",
			territory.UN:    "Mba Ụwa Jikọrọ Ọnụ",
			territory.US:    "Mba United States",
			territory.UY:    "Uruguay",
			territory.UZ:    "Obodo Uzbekistan",
			territory.VA:    "Vatican City",
			territory.VC:    "St. Vincent & Grenadines",
			territory.VE:    "Venezuela",
			territory.VG:    "Agwaetiti British Virgin",
			territory.VI:    "Agwaetiti Virgin nke US",
			territory.VN:    "Vietnam",
			territory.VU:    "Vanuatu",
			territory.WF:    "Wallis & Futuna",
			territory.WS:    "Samoa",
			territory.XA:    "Pseudo-Accents",
			territory.XB:    "Pseudo-Bidi",
			territory.XK:    "Kosovo",
			territory.YE:    "Obodo Yemen",
			territory.YT:    "Mayotte",
			territory.ZA:    "South Africa",
			territory.ZM:    "Zambia",
			territory.ZW:    "Zimbabwe",
			territory.ZZ:    "Mpaghara Amaghị",
		},
	}
}
