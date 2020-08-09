// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_pt_AO() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d 'de' MMM 'de' y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "fev.", Mar: "mar.", Apr: "abr.", May: "mai.", Jun: "jun.", Jul: "jul.", Aug: "ago.", Sep: "set.", Oct: "out.", Nov: "nov.", Dec: "dez."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "janeiro", Feb: "fevereiro", Mar: "março", Apr: "abril", May: "maio", Jun: "junho", Jul: "julho", Aug: "agosto", Sep: "setembro", Oct: "outubro", Nov: "novembro", Dec: "dezembro"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dom.", Mon: "seg.", Tue: "ter.", Wed: "qua.", Thu: "qui.", Fri: "sex.", Sat: "sáb."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "S", Tue: "T", Wed: "Q", Thu: "Q", Fri: "S", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "domingo", Mon: "segunda-feira", Tue: "terça-feira", Wed: "quarta-feira", Thu: "quinta-feira", Fri: "sexta-feira", Sat: "sábado"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_pt]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Peseta de Andorra", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "Dirham dos Emirados Árabes Unidos", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Afegane (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afegane afegão", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "Lek Albanês (1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Lek albanês", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Dram armênio", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Florim das Antilhas Holandesas", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza angolano", Symbol: "Kz"},
				currency.AOK: cldr.Currency{DisplayName: "Cuanza angolano (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Novo cuanza angolano (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Cuanza angolano reajustado (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Austral argentino", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "Peso lei argentino (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "Peso argentino (1881–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Peso argentino (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Peso argentino", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Xelim austríaco", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Dólar australiano", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "Florim arubano", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Manat azerbaijano (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Manat azeri", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Dinar da Bósnia-Herzegovina (1992–1994)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Marco conversível da Bósnia e Herzegovina", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "Novo dinar da Bósnia-Herzegovina (1994–1997)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "Dólar barbadense", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Taka bengali", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Franco belga (conversível)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Franco belga", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Franco belga (financeiro)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Lev forte búlgaro", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "Lev socialista búlgaro", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Lev búlgaro", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Lev búlgaro (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Dinar bareinita", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Franco burundiano", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Dólar bermudense", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Dólar bruneano", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano da Bolívia", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Boliviano (1863–1963)", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "Peso boliviano", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Mvdol boliviano", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Cruzeiro novo brasileiro (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Cruzado brasileiro (1986–1989)", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Cruzeiro brasileiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Real brasileiro", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Cruzado novo brasileiro (1989–1990)", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Cruzeiro brasileiro (1993–1994)", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "Cruzeiro brasileiro (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Dólar bahamense", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrum butanês", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Kyat birmanês", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Pula botsuanesa", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Rublo novo bielo-russo (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Rublo bielorrusso", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Rublo bielorrusso (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Dólar belizenho", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Dólar canadense", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Franco congolês", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "Euro WIR", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Franco suíço", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "Franco WIR", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "Escudo chileno", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Unidades de Fomento chilenas", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Peso chileno", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Yuan chinês (offshore)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "Dólar do Banco Popular da China", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Yuan chinês", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso colombiano", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Unidade de Valor Real", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "Colón costarriquenho", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Dinar sérvio (2002–2006)", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Coroa Forte checoslovaca", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Peso cubano conversível", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Peso cubano", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo cabo-verdiano", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Libra cipriota", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Coroa tcheca", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Ostmark da Alemanha Oriental", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Marco alemão", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Franco djiboutiano", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Coroa dinamarquesa", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Peso dominicano", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar argelino", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Sucre equatoriano", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Unidade de Valor Constante (UVC) do Equador", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Coroa estoniana", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Libra egípcia", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa da Eritreia", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "Peseta espanhola (conta A)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Peseta espanhola (conta conversível)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Peseta espanhola", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr etíope", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Marca finlandesa", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Dólar fijiano", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Libra malvinense", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Franco francês", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Libra esterlina", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Cupom Lari georgiano", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Lari georgiano", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Cedi de Gana (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Cedi ganês", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Libra de Gibraltar", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi gambiano", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Franco guineano", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Syli da Guiné", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Ekwele da Guiné Equatorial", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Dracma grego", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal guatemalteco", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Escudo da Guiné Portuguesa", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Peso da Guiné-Bissau", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Dólar guianense", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Dólar de Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira hondurenha", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Dinar croata", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Kuna croata", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Gourde haitiano", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Florim húngaro", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Rupia indonésia", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Libra irlandesa", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Libra israelita", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "Sheqel antigo israelita", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Novo shekel israelense", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupia indiana", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinar iraquiano", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Rial iraniano", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "Coroa antiga islandesa", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Coroa islandesa", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Lira italiana", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Dólar jamaicano", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Dinar jordaniano", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Iene japonês", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Xelim queniano", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Som quirguiz", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Riel cambojano", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Franco comoriano", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Won norte-coreano", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "Hwan da Coreia do Sul (1953–1962)", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "Won da Coreia do Sul (1945–1953)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "Won sul-coreano", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinar kuwaitiano", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Dólar das Ilhas Cayman", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge cazaque", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Kip laosiano", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Libra libanesa", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Rupia do Sri Lanka", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Dólar liberiano", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Loti do Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litas lituano", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "Talonas lituano", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Franco conversível de Luxemburgo", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Franco luxemburguês", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Franco financeiro de Luxemburgo", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Lats letão", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "Rublo letão", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Dinar líbio", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Dirham marroquino", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Franco marroquino", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "Franco monegasco", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "Cupon moldávio", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Leu moldávio", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Ariary malgaxe", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Franco de Madagascar", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Dinar macedônio", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Dinar macedônio (1992–1993)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Franco de Mali", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Quiat de Myanmar", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik mongol", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca de Macau", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya mauritana (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya mauritana", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Lira maltesa", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Libra maltesa", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupia mauriciana", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Rupia maldivana", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha malauiana", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Peso mexicano", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Peso Prata mexicano (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Unidade Mexicana de Investimento (UDI)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Ringgit malaio", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Escudo de Moçambique", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Metical de Moçambique (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Metical moçambicano", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Dólar namibiano", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Naira nigeriana", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Córdoba nicaraguense (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Córdoba nicaraguense", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Florim holandês", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Coroa norueguesa", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Rupia nepalesa", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Dólar neozelandês", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Rial omanense", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Balboa panamenho", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Inti peruano", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Novo sol peruano", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Sol peruano (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kina papuásia", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Peso filipino", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Rupia paquistanesa", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Zloty polonês", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Zloti polonês (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Escudo português", Symbol: "Esc."},
				currency.PYG: cldr.Currency{DisplayName: "Guarani paraguaio", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Rial catariano", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Dólar rodesiano", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Leu romeno (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Leu romeno", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Dinar sérvio", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rublo russo", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Rublo russo (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Franco ruandês", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal saudita", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Dólar das Ilhas Salomão", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Rupia seichelense", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Dinar sudanês (1992–2007)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Libra sudanesa", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Libra sudanesa (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Coroa sueca", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Dólar singapuriano", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Libra de Santa Helena", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Tolar Bons esloveno", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Coroa eslovaca", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Leone de Serra Leoa", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Xelim somali", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Dólar surinamês", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Florim do Suriname", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "Libra sul-sudanesa", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Dobra de São Tomé e Príncipe (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Dobra de São Tomé e Príncipe", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Rublo soviético", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Colom salvadorenho", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Libra síria", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni suazi", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Baht tailandês", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Rublo do Tadjiquistão", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Somoni tadjique", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Manat do Turcomenistão (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Manat turcomeno", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Dinar tunisiano", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Paʻanga tonganesa", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Escudo timorense", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Lira turca (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Lira turca", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Dólar de Trinidad e Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Novo dólar taiwanês", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Xelim tanzaniano", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Hryvnia ucraniano", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Karbovanetz ucraniano", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Xelim ugandense (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Xelim ugandense", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Dólar americano", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "Dólar norte-americano (Dia seguinte)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "Dólar norte-americano (Mesmo dia)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "Peso uruguaio en unidades indexadas", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Peso uruguaio (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Peso uruguaio", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Som uzbeque", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Bolívar venezuelano (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Bolívar venezuelano (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Bolívar venezuelano", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Dong vietnamita", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "Dong vietnamita (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "Vatu de Vanuatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Tala samoano", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Franco CFA de BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Prata", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Ouro", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Unidade Composta Europeia", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Unidade Monetária Europeia", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Unidade de Conta Europeia (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Unidade de Conta Europeia (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Dólar do Caribe Oriental", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Direitos Especiais de Giro", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "Unidade de Moeda Europeia", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Franco-ouro francês", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Franco UIC francês", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "Franco CFA de BCEAO", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Paládio", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "Franco CFP", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platina", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "Fundos RINET", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Código de Moeda de Teste", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Moeda desconhecida", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Dinar iemenita", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Rial iemenita", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Dinar forte iugoslavo (1966–1990)", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Dinar noviy iugoslavo (1994–2002)", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Dinar conversível iugoslavo (1990–1992)", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "Dinar reformado iugoslavo (1992–1993)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Rand sul-africano (financeiro)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Rand sul-africano", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Cuacha zambiano (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha zambiano", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Zaire Novo zairense (1993–1998)", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Zaire zairense (1971–1993)", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Dólar do Zimbábue (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Dólar do Zimbábue (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Dólar do Zimbábue (2008)", Symbol: ""},
			},
		},
	}
}
