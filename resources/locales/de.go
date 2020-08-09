// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_de() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "de",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'um' {0}", Long: "{1} 'um' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mär", Apr: "Apr", May: "Mai", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dez"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januar", Feb: "Februar", Mar: "März", Apr: "April", May: "Mai", Jun: "Juni", Jul: "Juli", Aug: "August", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Dezember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "So", Mon: "Mo", Tue: "Di", Wed: "Mi", Thu: "Do", Fri: "Fr", Sat: "Sa"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "M", Thu: "D", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "So.", Mon: "Mo.", Tue: "Di.", Wed: "Mi.", Thu: "Do.", Fri: "Fr.", Sat: "Sa."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sonntag", Mon: "Montag", Tue: "Dienstag", Wed: "Mittwoch", Thu: "Donnerstag", Fri: "Freitag", Sat: "Samstag"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_de]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorranische Pesete", Symbol: "ADP"},
				currency.AED: cldr.Currency{DisplayName: "VAE-Dirham", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Afghanische Afghani (1927–2002)", Symbol: "AFA"},
				currency.AFN: cldr.Currency{DisplayName: "Afghanischer Afghani", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "Albanischer Lek (1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Albanischer Lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Armenischer Dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Niederländische-Antillen-Gulden", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Angolanischer Kwanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Angolanischer Kwanza (1977–1990)", Symbol: "AOK"},
				currency.AON: cldr.Currency{DisplayName: "Angolanischer Neuer Kwanza (1990–2000)", Symbol: "AON"},
				currency.AOR: cldr.Currency{DisplayName: "Angolanischer Kwanza Reajustado (1995–1999)", Symbol: "AOR"},
				currency.ARA: cldr.Currency{DisplayName: "Argentinischer Austral", Symbol: "ARA"},
				currency.ARL: cldr.Currency{DisplayName: "Argentinischer Peso Ley (1970–1983)", Symbol: "ARL"},
				currency.ARM: cldr.Currency{DisplayName: "Argentinischer Peso (1881–1970)", Symbol: "ARM"},
				currency.ARP: cldr.Currency{DisplayName: "Argentinischer Peso (1983–1985)", Symbol: "ARP"},
				currency.ARS: cldr.Currency{DisplayName: "Argentinischer Peso", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Österreichischer Schilling", Symbol: "öS"},
				currency.AUD: cldr.Currency{DisplayName: "Australischer Dollar", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruba-Florin", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Aserbaidschan-Manat (1993–2006)", Symbol: "AZM"},
				currency.AZN: cldr.Currency{DisplayName: "Aserbaidschan-Manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Bosnien und Herzegowina Dinar (1992–1994)", Symbol: "BAD"},
				currency.BAM: cldr.Currency{DisplayName: "Bosnien und Herzegowina Konvertierbare Mark", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "Bosnien und Herzegowina Neuer Dinar (1994–1997)", Symbol: "BAN"},
				currency.BBD: cldr.Currency{DisplayName: "Barbados-Dollar", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladesch-Taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Belgischer Franc (konvertibel)", Symbol: "BEC"},
				currency.BEF: cldr.Currency{DisplayName: "Belgischer Franc", Symbol: "BEF"},
				currency.BEL: cldr.Currency{DisplayName: "Belgischer Finanz-Franc", Symbol: "BEL"},
				currency.BGL: cldr.Currency{DisplayName: "Bulgarische Lew (1962–1999)", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "Bulgarischer Lew (1952–1962)", Symbol: "BGK"},
				currency.BGN: cldr.Currency{DisplayName: "Bulgarischer Lew", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Bulgarischer Lew (1879–1952)", Symbol: "BGJ"},
				currency.BHD: cldr.Currency{DisplayName: "Bahrain-Dinar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burundi-Franc", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda-Dollar", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Brunei-Dollar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Bolivianischer Boliviano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Bolivianischer Boliviano (1863–1963)", Symbol: "BOL"},
				currency.BOP: cldr.Currency{DisplayName: "Bolivianischer Peso", Symbol: "BOP"},
				currency.BOV: cldr.Currency{DisplayName: "Boliviansiche Mvdol", Symbol: "BOV"},
				currency.BRB: cldr.Currency{DisplayName: "Brasilianischer Cruzeiro Novo (1967–1986)", Symbol: "BRB"},
				currency.BRC: cldr.Currency{DisplayName: "Brasilianischer Cruzado (1986–1989)", Symbol: "BRC"},
				currency.BRE: cldr.Currency{DisplayName: "Brasilianischer Cruzeiro (1990–1993)", Symbol: "BRE"},
				currency.BRL: cldr.Currency{DisplayName: "Brasilianischer Real", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Brasilianischer Cruzado Novo (1989–1990)", Symbol: "BRN"},
				currency.BRR: cldr.Currency{DisplayName: "Brasilianischer Cruzeiro (1993–1994)", Symbol: "BRR"},
				currency.BRZ: cldr.Currency{DisplayName: "Brasilianischer Cruzeiro (1942–1967)", Symbol: "BRZ"},
				currency.BSD: cldr.Currency{DisplayName: "Bahamas-Dollar", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Bhutan-Ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Birmanischer Kyat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Botswanischer Pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Belarus-Rubel (1994–1999)", Symbol: "BYB"},
				currency.BYN: cldr.Currency{DisplayName: "Weißrussischer Rubel", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Weißrussischer Rubel (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Belize-Dollar", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kanadischer Dollar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kongo-Franc", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR-Euro", Symbol: "CHE"},
				currency.CHF: cldr.Currency{DisplayName: "Schweizer Franken", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR Franken", Symbol: "CHW"},
				currency.CLE: cldr.Currency{DisplayName: "Chilenischer Escudo", Symbol: "CLE"},
				currency.CLF: cldr.Currency{DisplayName: "Chilenische Unidades de Fomento", Symbol: "CLF"},
				currency.CLP: cldr.Currency{DisplayName: "Chilenischer Peso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Renminbi Yuan (Off–Shore)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "Dollar der Chinesischen Volksbank", Symbol: "CNX"},
				currency.CNY: cldr.Currency{DisplayName: "Renminbi Yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolumbianischer Peso", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Kolumbianische Unidades de valor real", Symbol: "COU"},
				currency.CRC: cldr.Currency{DisplayName: "Costa-Rica-Colón", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Serbischer Dinar (2002–2006)", Symbol: "CSD"},
				currency.CSK: cldr.Currency{DisplayName: "Tschechoslowakische Krone", Symbol: "CSK"},
				currency.CUC: cldr.Currency{DisplayName: "Kubanischer Peso (konvertibel)", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Kubanischer Peso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Cabo-Verde-Escudo", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Zypern-Pfund", Symbol: "CYP"},
				currency.CZK: cldr.Currency{DisplayName: "Tschechische Krone", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Mark der DDR", Symbol: "DDM"},
				currency.DEM: cldr.Currency{DisplayName: "Deutsche Mark", Symbol: "DM"},
				currency.DJF: cldr.Currency{DisplayName: "Dschibuti-Franc", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Dänische Krone", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Dominikanischer Peso", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Algerischer Dinar", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Ecuadorianischer Sucre", Symbol: "ECS"},
				currency.ECV: cldr.Currency{DisplayName: "Verrechnungseinheit für Ecuador", Symbol: "ECV"},
				currency.EEK: cldr.Currency{DisplayName: "Estnische Krone", Symbol: "EEK"},
				currency.EGP: cldr.Currency{DisplayName: "Ägyptisches Pfund", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritreischer Nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "Spanische Peseta (A–Konten)", Symbol: "ESA"},
				currency.ESB: cldr.Currency{DisplayName: "Spanische Peseta (konvertibel)", Symbol: "ESB"},
				currency.ESP: cldr.Currency{DisplayName: "Spanische Peseta", Symbol: "ESP"},
				currency.ETB: cldr.Currency{DisplayName: "Äthiopischer Birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Finnische Mark", Symbol: "FIM"},
				currency.FJD: cldr.Currency{DisplayName: "Fidschi-Dollar", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Falkland-Pfund", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Französischer Franc", Symbol: "FRF"},
				currency.GBP: cldr.Currency{DisplayName: "Britisches Pfund", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Georgischer Kupon Larit", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Georgischer Lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Ghanaischer Cedi (1979–2007)", Symbol: "GHC"},
				currency.GHS: cldr.Currency{DisplayName: "Ghanaischer Cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltar-Pfund", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gambia-Dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Guinea-Franc", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Guineischer Syli", Symbol: "GNS"},
				currency.GQE: cldr.Currency{DisplayName: "Äquatorialguinea-Ekwele", Symbol: "GQE"},
				currency.GRD: cldr.Currency{DisplayName: "Griechische Drachme", Symbol: "GRD"},
				currency.GTQ: cldr.Currency{DisplayName: "Guatemaltekischer Quetzal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Portugiesisch Guinea Escudo", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Guinea-Bissau Peso", Symbol: "GWP"},
				currency.GYD: cldr.Currency{DisplayName: "Guyana-Dollar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hongkong-Dollar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Honduras-Lempira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Kroatischer Dinar", Symbol: "HRD"},
				currency.HRK: cldr.Currency{DisplayName: "Kroatischer Kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Haitianische Gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Ungarischer Forint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indonesische Rupiah", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Irisches Pfund", Symbol: "IEP"},
				currency.ILP: cldr.Currency{DisplayName: "Israelisches Pfund", Symbol: "ILP"},
				currency.ILR: cldr.Currency{DisplayName: "Israelischer Schekel (1980–1985)", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Israelischer Neuer Schekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Indische Rupie", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Irakischer Dinar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Iranischer Rial", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "Isländische Krone (1918–1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Isländische Krone", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Italienische Lira", Symbol: "ITL"},
				currency.JMD: cldr.Currency{DisplayName: "Jamaika-Dollar", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Jordanischer Dinar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Japanischer Yen", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "Kenia-Schilling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Kirgisischer Som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kambodschanischer Riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Komoren-Franc", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Nordkoreanischer Won", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "Südkoreanischer Hwan (1953–1962)", Symbol: "KRH"},
				currency.KRO: cldr.Currency{DisplayName: "Südkoreanischer Won (1945–1953)", Symbol: "KRO"},
				currency.KRW: cldr.Currency{DisplayName: "Südkoreanischer Won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Kuwait-Dinar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Kaiman-Dollar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Kasachischer Tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laotischer Kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Libanesisches Pfund", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Sri-Lanka-Rupie", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Liberianischer Dollar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Loti", Symbol: "LSL"},
				currency.LTL: cldr.Currency{DisplayName: "Litauischer Litas", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Litauischer Talonas", Symbol: "LTT"},
				currency.LUC: cldr.Currency{DisplayName: "Luxemburgischer Franc (konvertibel)", Symbol: "LUC"},
				currency.LUF: cldr.Currency{DisplayName: "Luxemburgischer Franc", Symbol: "LUF"},
				currency.LUL: cldr.Currency{DisplayName: "Luxemburgischer Finanz-Franc", Symbol: "LUL"},
				currency.LVL: cldr.Currency{DisplayName: "Lettischer Lats", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Lettischer Rubel", Symbol: "LVR"},
				currency.LYD: cldr.Currency{DisplayName: "Libyscher Dinar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Marokkanischer Dirham", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Marokkanischer Franc", Symbol: "MAF"},
				currency.MCF: cldr.Currency{DisplayName: "Monegassischer Franc", Symbol: "MCF"},
				currency.MDC: cldr.Currency{DisplayName: "Moldau-Cupon", Symbol: "MDC"},
				currency.MDL: cldr.Currency{DisplayName: "Moldau-Leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Madagaskar-Ariary", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Madagaskar-Franc", Symbol: "MGF"},
				currency.MKD: cldr.Currency{DisplayName: "Mazedonischer Denar", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Mazedonischer Denar (1992–1993)", Symbol: "MKN"},
				currency.MLF: cldr.Currency{DisplayName: "Malischer Franc", Symbol: "MLF"},
				currency.MMK: cldr.Currency{DisplayName: "Myanmarischer Kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mongolischer Tögrög", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Macao-Pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Mauretanischer Ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Mauretanischer Ouguiya", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Maltesische Lira", Symbol: "MTL"},
				currency.MTP: cldr.Currency{DisplayName: "Maltesisches Pfund", Symbol: "MTP"},
				currency.MUR: cldr.Currency{DisplayName: "Mauritius-Rupie", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "Malediven-Rupie (alt)", Symbol: ""},
				currency.MVR: cldr.Currency{DisplayName: "Malediven-Rufiyaa", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malawi-Kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Mexikanischer Peso", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Mexikanischer Silber-Peso (1861–1992)", Symbol: "MXP"},
				currency.MXV: cldr.Currency{DisplayName: "Mexicanischer Unidad de Inversion (UDI)", Symbol: "MXV"},
				currency.MYR: cldr.Currency{DisplayName: "Malaysischer Ringgit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Mosambikanischer Escudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Mosambikanischer Metical (1980–2006)", Symbol: "MZM"},
				currency.MZN: cldr.Currency{DisplayName: "Mosambikanischer Metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namibia-Dollar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Nigerianischer Naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Nicaraguanischer Córdoba (1988–1991)", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "Nicaragua-Córdoba", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Niederländischer Gulden", Symbol: "NLG"},
				currency.NOK: cldr.Currency{DisplayName: "Norwegische Krone", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nepalesische Rupie", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Neuseeland-Dollar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Omanischer Rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Panamaischer Balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Peruanischer Inti", Symbol: "PEI"},
				currency.PEN: cldr.Currency{DisplayName: "Peruanischer Sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Peruanischer Sol (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "Papua-Neuguineischer Kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Philippinischer Peso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistanische Rupie", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Polnischer Złoty", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Polnischer Zloty (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Portugiesischer Escudo", Symbol: "PTE"},
				currency.PYG: cldr.Currency{DisplayName: "Paraguayischer Guaraní", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Katar-Riyal", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Rhodesischer Dollar", Symbol: "RHD"},
				currency.ROL: cldr.Currency{DisplayName: "Rumänischer Leu (1952–2006)", Symbol: "ROL"},
				currency.RON: cldr.Currency{DisplayName: "Rumänischer Leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Serbischer Dinar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Russischer Rubel", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Russischer Rubel (1991–1998)", Symbol: "RUR"},
				currency.RWF: cldr.Currency{DisplayName: "Ruanda-Franc", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudi-Rial", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Salomonen-Dollar", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Seychellen-Rupie", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Sudanesischer Dinar (1992–2007)", Symbol: "SDD"},
				currency.SDG: cldr.Currency{DisplayName: "Sudanesisches Pfund", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Sudanesisches Pfund (1957–1998)", Symbol: "SDP"},
				currency.SEK: cldr.Currency{DisplayName: "Schwedische Krone", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singapur-Dollar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "St. Helena-Pfund", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Slowenischer Tolar", Symbol: "SIT"},
				currency.SKK: cldr.Currency{DisplayName: "Slowakische Krone", Symbol: "SKK"},
				currency.SLL: cldr.Currency{DisplayName: "Sierra-leonischer Leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Somalia-Schilling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Suriname-Dollar", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Suriname Gulden", Symbol: "SRG"},
				currency.SSP: cldr.Currency{DisplayName: "Südsudanesisches Pfund", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "São-toméischer Dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "São-toméischer Dobra", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Sowjetischer Rubel", Symbol: "SUR"},
				currency.SVC: cldr.Currency{DisplayName: "El Salvador Colon", Symbol: "SVC"},
				currency.SYP: cldr.Currency{DisplayName: "Syrisches Pfund", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Swasiländischer Lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Thailändischer Baht", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Tadschikistan Rubel", Symbol: "TJR"},
				currency.TJS: cldr.Currency{DisplayName: "Tadschikistan-Somoni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Turkmenistan-Manat (1993–2009)", Symbol: "TMM"},
				currency.TMT: cldr.Currency{DisplayName: "Turkmenistan-Manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tunesischer Dinar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tongaischer Paʻanga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Timor-Escudo", Symbol: "TPE"},
				currency.TRL: cldr.Currency{DisplayName: "Türkische Lira (1922–2005)", Symbol: "TRL"},
				currency.TRY: cldr.Currency{DisplayName: "Türkische Lira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad und Tobago-Dollar", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Neuer Taiwan-Dollar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tansania-Schilling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Ukrainische Hrywnja", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Ukrainischer Karbovanetz", Symbol: "UAK"},
				currency.UGS: cldr.Currency{DisplayName: "Uganda-Schilling (1966–1987)", Symbol: "UGS"},
				currency.UGX: cldr.Currency{DisplayName: "Uganda-Schilling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "US-Dollar", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "US Dollar (Nächster Tag)", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "US Dollar (Gleicher Tag)", Symbol: "USS"},
				currency.UYI: cldr.Currency{DisplayName: "Uruguayischer Peso (Indexierte Rechnungseinheiten)", Symbol: "UYI"},
				currency.UYP: cldr.Currency{DisplayName: "Uruguayischer Peso (1975–1993)", Symbol: "UYP"},
				currency.UYU: cldr.Currency{DisplayName: "Uruguayischer Peso", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Usbekistan-Sum", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Venezolanischer Bolívar (1871–2008)", Symbol: "VEB"},
				currency.VEF: cldr.Currency{DisplayName: "Venezolanischer Bolívar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Venezolanischer Bolívar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Vietnamesischer Dong", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "Vietnamesischer Dong(1978–1985)", Symbol: "VNN"},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatu-Vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samoanischer Tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "CFA-Franc (BEAC)", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Unze Silber", Symbol: "XAG"},
				currency.XAU: cldr.Currency{DisplayName: "Unze Gold", Symbol: "XAU"},
				currency.XBA: cldr.Currency{DisplayName: "Europäische Rechnungseinheit", Symbol: "XBA"},
				currency.XBB: cldr.Currency{DisplayName: "Europäische Währungseinheit (XBB)", Symbol: "XBB"},
				currency.XBC: cldr.Currency{DisplayName: "Europäische Rechnungseinheit (XBC)", Symbol: "XBC"},
				currency.XBD: cldr.Currency{DisplayName: "Europäische Rechnungseinheit (XBD)", Symbol: "XBD"},
				currency.XCD: cldr.Currency{DisplayName: "Ostkaribischer Dollar", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Sonderziehungsrechte", Symbol: "XDR"},
				currency.XEU: cldr.Currency{DisplayName: "Europäische Währungseinheit (XEU)", Symbol: "XEU"},
				currency.XFO: cldr.Currency{DisplayName: "Französischer Gold-Franc", Symbol: "XFO"},
				currency.XFU: cldr.Currency{DisplayName: "Französischer UIC-Franc", Symbol: "XFU"},
				currency.XOF: cldr.Currency{DisplayName: "CFA-Franc (BCEAO)", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Unze Palladium", Symbol: "XPD"},
				currency.XPF: cldr.Currency{DisplayName: "CFP-Franc", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Unze Platin", Symbol: "XPT"},
				currency.XRE: cldr.Currency{DisplayName: "RINET Funds", Symbol: "XRE"},
				currency.XSU: cldr.Currency{DisplayName: "SUCRE", Symbol: "XSU"},
				currency.XTS: cldr.Currency{DisplayName: "Testwährung", Symbol: "XTS"},
				currency.XUA: cldr.Currency{DisplayName: "Rechnungseinheit der AfEB", Symbol: "XUA"},
				currency.XXX: cldr.Currency{DisplayName: "Unbekannte Währung", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "Jemen-Dinar", Symbol: "YDD"},
				currency.YER: cldr.Currency{DisplayName: "Jemen-Rial", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Jugoslawischer Dinar (1966–1990)", Symbol: "YUD"},
				currency.YUM: cldr.Currency{DisplayName: "Jugoslawischer Neuer Dinar (1994–2002)", Symbol: "YUM"},
				currency.YUN: cldr.Currency{DisplayName: "Jugoslawischer Dinar (konvertibel)", Symbol: "YUN"},
				currency.YUR: cldr.Currency{DisplayName: "Jugoslawischer reformierter Dinar (1992–1993)", Symbol: "YUR"},
				currency.ZAL: cldr.Currency{DisplayName: "Südafrikanischer Rand (Finanz)", Symbol: "ZAL"},
				currency.ZAR: cldr.Currency{DisplayName: "Südafrikanischer Rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Zaire-Neuer Zaïre (1993–1998)", Symbol: "ZRN"},
				currency.ZRZ: cldr.Currency{DisplayName: "Zaire-Zaïre (1971–1993)", Symbol: "ZRZ"},
				currency.ZWD: cldr.Currency{DisplayName: "Simbabwe-Dollar (1980–2008)", Symbol: "ZWD"},
				currency.ZWL: cldr.Currency{DisplayName: "Simbabwe-Dollar (2009)", Symbol: "ZWL"},
				currency.ZWR: cldr.Currency{DisplayName: "Simbabwe-Dollar (2008)", Symbol: "ZWR"},
			},
		},
	}
}
