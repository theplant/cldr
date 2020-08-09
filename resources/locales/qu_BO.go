// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_qu_BO() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{0} {1}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ene", Feb: "Feb", Mar: "Mar", Apr: "Abr", May: "May", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Set", Oct: "Oct", Nov: "Nov", Dec: "Dic"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Enero", Feb: "Febrero", Mar: "Marzo", Apr: "Abril", May: "Mayo", Jun: "Junio", Jul: "Julio", Aug: "Agosto", Sep: "Setiembre", Oct: "Octubre", Nov: "Noviembre", Dec: "Diciembre"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dom", Mon: "Lun", Tue: "Mar", Wed: "Mié", Thu: "Jue", Fri: "Vie", Sat: "Sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "X", Thu: "J", Fri: "V", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Dom", Mon: "Lun", Tue: "Mar", Wed: "Mié", Thu: "Jue", Fri: "Vie", Sat: "Sab"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Domingo", Mon: "Lunes", Tue: "Martes", Wed: "Miércoles", Thu: "Jueves", Fri: "Viernes", Sat: "Sábado"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_root]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Afgani Afgano", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Lek albanés", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Florín Antillano Neerlandés", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza Angoleño", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Peso Argentino", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Dólar Australiano", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Florín Arubeño", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Marco Bosnioherzegovino", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Dólar de Barbados", Symbol: "BBG"},
				currency.BDT: cldr.Currency{DisplayName: "Taka Bangladesí", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Lev", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Dinar Bareiní", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Franco Burundés", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Dólar Bermudeño", Symbol: "DBM"},
				currency.BND: cldr.Currency{DisplayName: "Dólar de Brunéi", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Real Brasileño", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Dólar Bahameño", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrum Butanés", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Pula Botswano", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Nuevo Rublo Bielorruso", Symbol: "BYN"},
				currency.BZD: cldr.Currency{DisplayName: "Dólar Beliceño", Symbol: "DBZ"},
				currency.CAD: cldr.Currency{DisplayName: "Dólar Canadiense", Symbol: "$CA"},
				currency.CDF: cldr.Currency{DisplayName: "Franco Congoleño", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Franco Suizo", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Peso Chileno", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Yuan Chino (offshore)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Chino", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso Colombiano", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Colón Costarricense", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Peso Cubano Convertible", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Peso Cubano", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo Caboverdiano", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Corona Checa", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Franco Yibutiano", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Corona Danesa", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Peso Dominicano", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar Argelino", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Libra Egipcia", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa Eritreano", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr Etíope", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dólar Fiyiano", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Libra Malvinense", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Libra Esterlina", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Lari Georgiano", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Cedi Ganés", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Libra Gibraltareña", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Franco Guineano", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal Guatemalteco", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Dólar Guyanés", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Dólar de Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira Hondureño", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Kuna Croata", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Gourde Haitiano", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Florín Húngaro", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Rupia Indonesia", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Nuevo Séquel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupia India", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinar Iraquí", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Rial Iraní", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Corona Islandesa", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Dólar Jamaiquino", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Dinar Jordano", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Yen Japonés", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Chelín Keniano", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Som Kirguís", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Riel Camboyano", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Franco Comorense", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Won Norcoreano", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Won Surcoreano", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinar Kuwaití", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Dólar de las Islas Caimán", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge Kazajo", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Kip Laosiano", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Libra Libanesa", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Rupia de Sri Lanka", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Dólar Liberiano", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinar Libio", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Dírham Marroquí", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Leu Moldavo", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Ariary Malgache", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Dinar Macedonio", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Kyat Birmano", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik Mongol", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca Macaense", Symbol: "MOP"},
				currency.MRU: cldr.Currency{DisplayName: "Uguiya Mauritano", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "Rupia de Mauricio", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Rupia de Maldivas", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha Malauí", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Peso Mexicano", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ringgit Malayo", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Metical Mozambiqueño", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Dólar Namibio", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Naira Nigeriano", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Córdova Nicaragüense", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Corona Noruega", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Rupia Nepalí", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Dólar Neozelandés", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Rial Omaní", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Balboa Panameño", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Sol Peruano", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Kina Papuano", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Peso Filipino", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Rupia Pakistaní", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Zloty", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Guaraní Paraguayo", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Riyal Catarí", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Leu Rumano", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Dinar Serbio", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rublo Ruso", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Franco Ruandés", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal Saudí", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Dólar de las Islas Salomón", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Rupia de Seychelles", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Libra Sudanesa", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "Corona Sueca", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Dólar de Singapur", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Libra de Santa Helena", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Leone de Sierra Leona", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Chelín Somalí", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Dólar Surinamés", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Libra Sursudanesa", Symbol: "SSP"},
				currency.STN: cldr.Currency{DisplayName: "Dobra Santotomense", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Libra Siria", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni Swazi", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Baht Tailandés", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "Somoni Tayiko", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Manat Turcomano", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Dinar Tunecino", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Paʻanga Tongano", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Lira Turca", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Dólar de Trinidad y Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Nuevo Dólar Taiwanés", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Chelín Tanzano", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Grivna", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Chelín Ugandés", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Dólar Americano", Symbol: "$US"},
				currency.UYU: cldr.Currency{DisplayName: "Peso Uruguayo", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Som Ubzeko", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Bolívar Venezolano", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Dong Vietnamita", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu Vanuatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Tala Samoano", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Franco CFA de África Central", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dólar del Caribe Oriental", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Franco CFA de África Occidental", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Franco CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Mana riqsisqa Qullqi", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Rial Yemení", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Rand Sudafricano", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha Zambiano", Symbol: "ZMW"},
			},
		},
	}
}
