// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_gsw_CH() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "gsw_CH",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mär", Apr: "Apr", May: "Mai", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dez"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januar", Feb: "Februar", Mar: "März", Apr: "April", May: "Mai", Jun: "Juni", Jul: "Juli", Aug: "Auguscht", Sep: "Septämber", Oct: "Oktoober", Nov: "Novämber", Dec: "Dezämber"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Su.", Mon: "Mä.", Tue: "Zi.", Wed: "Mi.", Thu: "Du.", Fri: "Fr.", Sat: "Sa."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "M", Thu: "D", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sunntig", Mon: "Määntig", Tue: "Ziischtig", Wed: "Mittwuch", Thu: "Dunschtig", Fri: "Friitig", Sat: "Samschtig"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "vorm.", PM: "nam."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Vormittag", PM: "Namittag"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_gsw]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: "’", Negative: "−", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorranischi Peseete", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "UAE Dirham", Symbol: ""},
				currency.AFA: cldr.Currency{DisplayName: "Afghani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afghani", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Lek", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Dram", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "Niderländischi-Antille-Gulde", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza", Symbol: "Kz"},
				currency.AOK: cldr.Currency{DisplayName: "Angolanische Kwanza (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Nöie Kwanza", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Kwanza Reajustado", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Argentinische Auschtral", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Argentinische Peso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Argentinische Peso", Symbol: "$"},
				currency.ATS: cldr.Currency{DisplayName: "Öschtriichische Schilling", Symbol: "öS"},
				currency.AUD: cldr.Currency{DisplayName: "Auschtralische Dollar", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruba Florin", Symbol: ""},
				currency.AZM: cldr.Currency{DisplayName: "Aserbeidschanische Manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Aserbeidschanische Manat", Symbol: ""},
				currency.BAD: cldr.Currency{DisplayName: "Bosnie-und-Herzegowina-Dinar", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Konvertierbari Mark vo Bosnie und Herzegowina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbados-Dollar", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Taka", Symbol: "৳"},
				currency.BEC: cldr.Currency{DisplayName: "Belgische Franc (konvertibel)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Belgische Franc", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Belgische Finanz-Franc", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Lew (1962–1999)", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Bulgarische Lew", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Bahrain-Dinar", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Burundi-Franc", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda-Dollar", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Brunei-Dollar", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano", Symbol: "Bs"},
				currency.BOP: cldr.Currency{DisplayName: "Bolivianische Peso", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Bolivianische Mvdol", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Brasilianische Cruzeiro Novo (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Brasilianische Cruzado", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Brasilianische Cruzeiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Brasilianische Real", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Brasilianische Cruzado Novo", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Brasilianische Cruzeiro", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Bahama-Dollar", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Bhutanische Ngultrum", Symbol: ""},
				currency.BUK: cldr.Currency{DisplayName: "Birmanische Kyat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Botswanische Pula", Symbol: "P"},
				currency.BYB: cldr.Currency{DisplayName: "Belarus-Rubel (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Belarus Rubel", Symbol: "р."},
				currency.BYR: cldr.Currency{DisplayName: "Belarus Rubel (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Belize-Dollar", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Kanadische Dollar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kongolesische Franc", Symbol: ""},
				currency.CHE: cldr.Currency{DisplayName: "WIR-Euro", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Schwiizer Franke", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR-Franke", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Tschileenische Unidad de Fomento", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Tschileenische Peso", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Renminbi Yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolumbianische Peso", Symbol: "$"},
				currency.COU: cldr.Currency{DisplayName: "Unidad de Valor Real", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "Costa Rica Colon", Symbol: "₡"},
				currency.CSD: cldr.Currency{DisplayName: "Alte Serbische Dinar", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Tschechoslowakischi Chroone", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Kubanische Peso", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Kap Verde Escudo", Symbol: ""},
				currency.CYP: cldr.Currency{DisplayName: "Zypere-Pfund", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Tschechischi Chroone", Symbol: "Kč"},
				currency.DDM: cldr.Currency{DisplayName: "DDR-Mark", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Tüütschi Mark", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Dschibuti-Franc", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Tänischi Chroone", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Tominikanische Peso", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Algeerischi Dinar", Symbol: ""},
				currency.ECS: cldr.Currency{DisplayName: "Ecuadorianische Sucre", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Verrächnigsäiheit für EC", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Eestnischi Chroone", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Ägüptischs Pfund", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Eritreische Nakfa", Symbol: ""},
				currency.ESA: cldr.Currency{DisplayName: "Schpanischi Peseeta (A–Kontene)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Schpanischi Peseeta (konvertibel)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Schpanischi Peseeta", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Äthiopische Birr", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Finnischi Mark", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Fidschi Dollar", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Falkland-Pfund", Symbol: "£"},
				currency.FRF: cldr.Currency{DisplayName: "Französische Franc", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Pfund Schtörling", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Georgische Kupon Larit", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Georgische Lari", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Ghanaische Cedi (GHC)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Ghanaische Cedi (GHS)", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltar-Pfund", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Gambische Dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Guinea-Franc", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Guineische Syli", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Äquatorialguinea-Ekwele", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Griechische Trachme", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal", Symbol: "Q"},
				currency.GWE: cldr.Currency{DisplayName: "Portugiisische Guinea Escudo", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Guinea-Bissau-Peso", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Guyana-Dollar", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Hongkong-Dollar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira", Symbol: "L"},
				currency.HRD: cldr.Currency{DisplayName: "Kroazische Dinar", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Kuna", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Gourde", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Forint", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Indonesischi Rupie", Symbol: "Rp"},
				currency.IEP: cldr.Currency{DisplayName: "Iirischs Pfund", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Israelischs Pfund", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Schekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Indischi Rupie", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Irak-Dinar", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Rial", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Iisländischi Chroone", Symbol: "kr"},
				currency.ITL: cldr.Currency{DisplayName: "Italiänischi Lira", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Jamaika-Dollar", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Jordaanische Dinar", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Yen", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "Kenia-Schilling", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Som", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "Riel", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Komore-Franc", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Nordkoreanische Won", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Süüdkoreanische Won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Kuwait-Dinar", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Kaiman-Dollar", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Kip", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Libaneesischs Pfund", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Sri-Lanka-Rupie", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Liberiaanische Dollar", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litauische Litas", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "Litauische Talonas", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Luxemburgische Franc (konvertibel)", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Luxemburgische Franc", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Luxemburgischer Finanz-Franc", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Lettische Lats", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "Lettische Rubel", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Lüübische Dinar", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Marokkanische Dirham", Symbol: ""},
				currency.MAF: cldr.Currency{DisplayName: "Marokkanischer Franc", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Moldau-Löi", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Madagaschkar-Ariary", Symbol: "Ar"},
				currency.MGF: cldr.Currency{DisplayName: "Madagaschkar-Franc", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Denar", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Malische Franc", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Kyat", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "Malteesischi Lira", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Malteesischs Pfund", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Maurizius-Rupie", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyaa", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Malawi-Kwacha", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Mexikanische Peso", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Mexikanische Silber-Peso (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Mexikanische Unidad de Inversion (UDI)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Malaysische Ringgit", Symbol: "RM"},
				currency.MZE: cldr.Currency{DisplayName: "Mosambikanische Escudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Alte Metical", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Metical", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Namibia-Dollar", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira", Symbol: "₦"},
				currency.NIC: cldr.Currency{DisplayName: "Cordoba", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Nicaragua-Córdoba", Symbol: "C$"},
				currency.NLG: cldr.Currency{DisplayName: "Holländische Gulde", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Norweegischi Chroone", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Nepaleesischi Rupie", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Neuseeland-Dollar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Rial Omani", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Balboa", Symbol: ""},
				currency.PEI: cldr.Currency{DisplayName: "Peruanische Inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Sol", Symbol: ""},
				currency.PES: cldr.Currency{DisplayName: "Sol (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kina", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Philippiinische Peso", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Pakischtanischi Rupie", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Zloty", Symbol: "zł"},
				currency.PLZ: cldr.Currency{DisplayName: "Zloty (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Portugiisische Escudo", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Guarani", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Katar-Riyal", Symbol: ""},
				currency.RHD: cldr.Currency{DisplayName: "Rhodesische Dollar", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Löi", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Rumäänische Löi", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Serbische Dinar", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Russische Rubel", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "Russische Rubel (alt)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Ruanda-Franc", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudi-Riyal", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Salomone-Dollar", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Seyschelle-Rupie", Symbol: ""},
				currency.SDD: cldr.Currency{DisplayName: "Sudaneesische Dinar", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Sudaneesischs Pfund", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Sudaneesischs Pfund (alt)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Schweedischi Chroone", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Singapur-Dollar", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "St.-Helena-Pfund", Symbol: "£"},
				currency.SIT: cldr.Currency{DisplayName: "Tolar", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Slowakischi Chroone", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Somalia-Schilling", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Surinamische Dollar", Symbol: "$"},
				currency.SRG: cldr.Currency{DisplayName: "Surinamische Gulde", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "Süüdsudaneesischs Pfund", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra", Symbol: "Db"},
				currency.SUR: cldr.Currency{DisplayName: "Sowjetische Rubel", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "El-Salvador-Colon", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Süürischs Pfund", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Baht", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Tadschikischtan-Rubel", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Tadschikischtan-Somoni", Symbol: ""},
				currency.TMM: cldr.Currency{DisplayName: "Turkmeenischtan-Manat", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Tuneesische Dinar", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Paʻanga", Symbol: "T$"},
				currency.TPE: cldr.Currency{DisplayName: "Timor-Escudo", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Türkischi Liire", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Nöii Türkischi Liire", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad-und-Tobago-Dollar", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Nöii Taiwan-Dollar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tansania-Schilling", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Hryvnia", Symbol: "₴"},
				currency.UAK: cldr.Currency{DisplayName: "Ukraiinische Karbovanetz", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Uganda-Schilling (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Uganda-Schilling", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "US-Dollar", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "US Dollar (Nöchschte Taag)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "US Dollar (Gliiche Taag)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Uruguayische Nöie Peso (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Uruguayische Peso", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Usbeekischtan-Sum", Symbol: ""},
				currency.VEB: cldr.Currency{DisplayName: "Bolivar (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Bolivar (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Bolivar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Dong", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Tala", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "CFA-Franc (Äquatoriaal)", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Silber", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Gold", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Europääischi Rächnigseinheit", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Europääischi Währigseinheit (XBB)", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Europääischi Rächnigseinheit (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Europääischi Rächnigseinheit (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Oschtkaribische Dollar", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Sunderziäigsrächt", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "Europääischi Währigseinheit (XEU)", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Französische Gold-Franc", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Französische UIC-Franc", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "CFA-Franc (Wescht)", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP-Franc", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platin", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET-Funds", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Teschtwährig", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Unbekannti Währig", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Jeme-Dinar", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Jeme-Rial", Symbol: ""},
				currency.YUD: cldr.Currency{DisplayName: "Jugoslawische Dinar (1966–1990)", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Nöii Dinar", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Jugoslawische Dinar (konvertibel)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Rand", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha", Symbol: "ZK"},
				currency.ZRN: cldr.Currency{DisplayName: "Nöie Zaire", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Zaire", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Simbabwe-Dollar", Symbol: ""},
			},
		},
	}
}
