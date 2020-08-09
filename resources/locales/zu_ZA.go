// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_zu_ZA() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mas", Apr: "Eph", May: "Mey", Jun: "Jun", Jul: "Jul", Aug: "Aga", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dis"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januwari", Feb: "Februwari", Mar: "Mashi", Apr: "Ephreli", May: "Meyi", Jun: "Juni", Jul: "Julayi", Aug: "Agasti", Sep: "Septhemba", Oct: "Okthoba", Nov: "Novemba", Dec: "Disemba"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Son", Mon: "Mso", Tue: "Bil", Wed: "Tha", Thu: "Sin", Fri: "Hla", Sat: "Mgq"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "B", Wed: "T", Thu: "S", Fri: "H", Sat: "M"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Son", Mon: "Mso", Tue: "Bil", Wed: "Tha", Thu: "Sin", Fri: "Hla", Sat: "Mgq"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "ISonto", Mon: "UMsombuluko", Tue: "ULwesibili", Wed: "ULwesithathu", Thu: "ULwesine", Fri: "ULwesihlanu", Sat: "UMgqibelo"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_zu]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "i-Dirham yase-United Arab Emirates", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "i-Afghan Afghani", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "i-Albanian Lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "i-Armenian Dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "i-Netherlands Antillean Guilder", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "i-Angolan Kwanza", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "i-Argentina Peso", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "i-Austrilian Dollar", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "i-Aruban Florin", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "i-Azerbaijani Manat", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "i-Bosnia-Herzegovina Convertible Mark", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "i-Barbadian Dollar", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "i-Bangladeshi Taka", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "i-Bulgarian Lev", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "i-Bahraini Dinar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "i-Burundian Franc", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "i-Bermudan Dollar", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "i-Brunei Dollar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "i-Bolivian Boliviano", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "i-Brazilian Real", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "i-Bahamian Dollar", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "i-Bhutanese Ngultrum", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "i-Botswana Pula", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "i-Belarusian Ruble", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "i-Belarusian Ruble (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "i-Belize Dollar", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "i-Candian Dollar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "i-Congolese Franc", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "i-Swiss Franc", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "i-Chilean Peso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "i-Chinese yuan (offshore)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "i-Chinese Yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "i-Colombian Peso", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "i-Costa Rican Colón", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "i-Cuban Convertable Peso", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "i-Cuban Peso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "i-Cape Verdean Escudo", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "i-Czech Republic Koruna", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "i-Djiboutian Franc", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "i-Danish Krone", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "i-Dominican Peso", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "i-Algerian Dinar", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "i-Egyptian Pound", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "i-Eritrean Nakfa", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "i-Ethopian Birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "i-Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "i-Fijian Dollar", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "i-Falkland Islands Pound", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "i-British Pound", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "i-Georgian Lari", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "i-Ghanaian Cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "i-Gibraltar Pound", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "i-Gambian Dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "i-Gunean Franc", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "i-Guatemalan Quetzal", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "i-Guyanaese Dollar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "i-Hong Kong Dollar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "i-Honduran Lempira", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "i-Croatian Kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "i-Haitian Gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "i-Hungarian Forint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "i-Indonesian Rupiah", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "i-Israeli New Sheqel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "i-Indian Rupee", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "i-Iraqi Dinar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "i-Iranian Rial", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "i-Icelandic Króna", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "i-Jamaican Dollar", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "i-Jordanian Dinar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "i-Japanese Yen", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "i-Kenyan Shilling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "i-Kyrgystani Som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "i-Cambodian Riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "i-Comorian Franc", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "i-North Korean Won", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "i-South Korean Won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "i-Kuwaiti Dinar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "i-Cayman Islands Dollar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "i-Kazakhstani Tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "i-Laotian Kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "i-Lebanese Pound", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "i-Sri Lankan Rupee", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "i-Liberian Dollar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "i-Lesotho Loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "i-Lithuanian Litas", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "i-Latvian Lats", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "i-Libyan Dinar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "i-Moroccan Dirham", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "i-Moldovan Leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "i-Malagasy Ariary", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "i-Macedonian Denar", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "i-Myanma Kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "i-Mongolian Tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "i-Macanese Pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "i-Mauritanian Ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "i-Mauritanian Ouguiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "i-Mauritian Rupee", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "i-Maldivian Rufiyana", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "i-Malawian Kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "i-Mexican Peso", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "i-Malaysian Ringgit", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "i-Mozambican Metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "i-Namibian Dollar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "i-Nigerian Naira", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "i-Nicaraguan Córdoba", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "i-Norwegian Krone", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "i-Nepalese Rupee", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "i-New Zealand Dollar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "i-Omani Rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "i-Panamanian Balboa", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "i-Peruvian Nuevo Sol", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "i-Papua New Guinean Kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "i-Philippine Peso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "i-Pakistani Rupee", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "i-Polish Zloty", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "i-Paraguayan Guarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "i-Qatari Rial", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "i-Romanian Leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "i-Serbian Dinar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "i-Russian Ruble", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "i-Rwandan Franc", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "i-Saudi Riyal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "i-Solomon Islands Dollar", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "i-Seychellois Rupee", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "i-Sudanese Pound", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "i-Swedish Krona", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "i-Singapore Dollar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "i-Saint Helena Pound", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "i-Sierra Leonean Leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "i-Somali Shilling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "i-Surinamese Dollar", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "i-South Sudanese Pound", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "i-São Tomé kanye ne-Príncipe Dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "i-São Tomé kanye ne-Príncipe Dobra", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "i-Syrian Pound", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "i-Swazi Lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "i-Thai Baht", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "i-Tajikistani Somoni", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "i-Turkmenistani Manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "i-Tunisian Dinar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "i-Tongan Paʻanga", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "i-Turkish Lira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "i-Trinidad and Tobago Dollar", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "i-New Taiwan Dollar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "i-Tanzanian Shilling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "i-Ukrainian Hryvnia", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "i-Ugandan Shilling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "i-US Dollar", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "i-Uruguayan Peso", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "i-Uzbekistan Som", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "i-Venezuelan Bolívar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "i-Venezuelan Bolívar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "i-Vietnamese Dong", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "i-Vanuatu Vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "i-Samoan Tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "i-Central African CFA Franc", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "i-East Caribbean Dollar", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "i-West African CFA Franc", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "i-CFP Franc", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "imali engaziwa", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "i-Yemeni Rial", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "i-South African Rand", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "i-Zambian Kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "i-Zambian Kwacha", Symbol: "ZMW"},
			},
		},
	}
}
