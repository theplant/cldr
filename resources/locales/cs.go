// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_cs() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d. MMMM y", Long: "d. MMMM y", Medium: "d. M. y", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "led", Feb: "úno", Mar: "bře", Apr: "dub", May: "kvě", Jun: "čvn", Jul: "čvc", Aug: "srp", Sep: "zář", Oct: "říj", Nov: "lis", Dec: "pro"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "leden", Feb: "únor", Mar: "březen", Apr: "duben", May: "květen", Jun: "červen", Jul: "červenec", Aug: "srpen", Sep: "září", Oct: "říjen", Nov: "listopad", Dec: "prosinec"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ne", Mon: "po", Tue: "út", Wed: "st", Thu: "čt", Fri: "pá", Sat: "so"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "N", Mon: "P", Tue: "Ú", Wed: "S", Thu: "Č", Fri: "P", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ne", Mon: "po", Tue: "út", Wed: "st", Thu: "čt", Fri: "pá", Sat: "so"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "neděle", Mon: "pondělí", Tue: "úterý", Wed: "středa", Thu: "čtvrtek", Fri: "pátek", Sat: "sobota"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "dop.", PM: "odp."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "dop.", PM: "odp."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "dop.", PM: "odp."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_cs]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "\u061c-", Percent: "٪\u061c", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "andorrská peseta", Symbol: "ADP"},
				currency.AED: cldr.Currency{DisplayName: "SAE dirham", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "afghánský afghán (1927–2002)", Symbol: "AFA"},
				currency.AFN: cldr.Currency{DisplayName: "afghánský afghán", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "albánský lek (1946–1965)", Symbol: "ALK"},
				currency.ALL: cldr.Currency{DisplayName: "albánský lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "arménský dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "nizozemskoantilský gulden", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "angolská kwanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "angolská kwanza (1977–1991)", Symbol: "AOK"},
				currency.AON: cldr.Currency{DisplayName: "angolská kwanza (1990–2000)", Symbol: "AON"},
				currency.AOR: cldr.Currency{DisplayName: "angolská kwanza (1995–1999)", Symbol: "AOR"},
				currency.ARA: cldr.Currency{DisplayName: "argentinský austral", Symbol: "ARA"},
				currency.ARL: cldr.Currency{DisplayName: "argentinské peso ley (1970–1983)", Symbol: "ARL"},
				currency.ARM: cldr.Currency{DisplayName: "argentinské peso (1881–1970)", Symbol: "ARM"},
				currency.ARP: cldr.Currency{DisplayName: "argentinské peso (1983–1985)", Symbol: "ARP"},
				currency.ARS: cldr.Currency{DisplayName: "argentinské peso", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "rakouský šilink", Symbol: "ATS"},
				currency.AUD: cldr.Currency{DisplayName: "australský dolar", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "arubský zlatý", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "ázerbájdžánský manat (1993–2006)", Symbol: "AZM"},
				currency.AZN: cldr.Currency{DisplayName: "ázerbájdžánský manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "bosenský dinár (1992–1994)", Symbol: "BAD"},
				currency.BAM: cldr.Currency{DisplayName: "bosenská konvertibilní marka", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "bosenský nový dinár (1994–1997)", Symbol: "BAN"},
				currency.BBD: cldr.Currency{DisplayName: "barbadoský dolar", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "bangladéšská taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "belgický konvertibilní frank", Symbol: "BEC"},
				currency.BEF: cldr.Currency{DisplayName: "belgický frank", Symbol: "BEF"},
				currency.BEL: cldr.Currency{DisplayName: "belgický finanční frank", Symbol: "BEL"},
				currency.BGL: cldr.Currency{DisplayName: "bulharský tvrdý leva", Symbol: "BGL"},
				currency.BGM: cldr.Currency{DisplayName: "bulharský socialistický leva", Symbol: "BGM"},
				currency.BGN: cldr.Currency{DisplayName: "bulharský leva", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "bulharský lev (1879–1952)", Symbol: "BGO"},
				currency.BHD: cldr.Currency{DisplayName: "bahrajnský dinár", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "burundský frank", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "bermudský dolar", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "brunejský dolar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "bolivijský boliviano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "bolivijský boliviano (1863–1963)", Symbol: "BOL"},
				currency.BOP: cldr.Currency{DisplayName: "bolivijské peso", Symbol: "BOP"},
				currency.BOV: cldr.Currency{DisplayName: "bolivijský mvdol", Symbol: "BOV"},
				currency.BRB: cldr.Currency{DisplayName: "brazilské nové cruzeiro (1967–1986)", Symbol: "BRB"},
				currency.BRC: cldr.Currency{DisplayName: "brazilské cruzado (1986–1989)", Symbol: "BRC"},
				currency.BRE: cldr.Currency{DisplayName: "brazilské cruzeiro (1990–1993)", Symbol: "BRE"},
				currency.BRL: cldr.Currency{DisplayName: "brazilský real", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "brazilské nové cruzado (1989–1990)", Symbol: "BRN"},
				currency.BRR: cldr.Currency{DisplayName: "brazilské cruzeiro (1993–1994)", Symbol: "BRR"},
				currency.BRZ: cldr.Currency{DisplayName: "brazilské cruzeiro (1942–1967)", Symbol: "BRZ"},
				currency.BSD: cldr.Currency{DisplayName: "bahamský dolar", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "bhútánský ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "barmský kyat", Symbol: "BUK"},
				currency.BWP: cldr.Currency{DisplayName: "botswanská pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "běloruský rubl (1994–1999)", Symbol: "BYB"},
				currency.BYN: cldr.Currency{DisplayName: "běloruský rubl", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "běloruský rubl (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "belizský dolar", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "kanadský dolar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "konžský frank", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "švýcarské WIR-euro", Symbol: "CHE"},
				currency.CHF: cldr.Currency{DisplayName: "švýcarský frank", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "švýcarský WIR-frank", Symbol: "CHW"},
				currency.CLE: cldr.Currency{DisplayName: "chilské escudo", Symbol: "CLE"},
				currency.CLF: cldr.Currency{DisplayName: "chilská účetní jednotka (UF)", Symbol: "CLF"},
				currency.CLP: cldr.Currency{DisplayName: "chilské peso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "čínský jüan (offshore)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "čínský dolar ČLB", Symbol: "CNX"},
				currency.CNY: cldr.Currency{DisplayName: "čínský jüan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "kolumbijské peso", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "kolumbijská jednotka reálné hodnoty", Symbol: "COU"},
				currency.CRC: cldr.Currency{DisplayName: "kostarický colón", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "srbský dinár (2002–2006)", Symbol: "CSD"},
				currency.CSK: cldr.Currency{DisplayName: "československá koruna", Symbol: "Kčs"},
				currency.CUC: cldr.Currency{DisplayName: "kubánské konvertibilní peso", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "kubánské peso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "kapverdské escudo", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "kyperská libra", Symbol: "CYP"},
				currency.CZK: cldr.Currency{DisplayName: "česká koruna", Symbol: "Kč"},
				currency.DDM: cldr.Currency{DisplayName: "východoněmecká marka", Symbol: "DDM"},
				currency.DEM: cldr.Currency{DisplayName: "německá marka", Symbol: "DEM"},
				currency.DJF: cldr.Currency{DisplayName: "džibutský frank", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "dánská koruna", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "dominikánské peso", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "alžírský dinár", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "ekvádorský sucre", Symbol: "ECS"},
				currency.ECV: cldr.Currency{DisplayName: "ekvádorská jednotka konstantní hodnoty", Symbol: "ECV"},
				currency.EEK: cldr.Currency{DisplayName: "estonská koruna", Symbol: "EEK"},
				currency.EGP: cldr.Currency{DisplayName: "egyptská libra", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "eritrejská nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "španělská peseta („A“ účet)", Symbol: "ESA"},
				currency.ESB: cldr.Currency{DisplayName: "španělská peseta (konvertibilní účet)", Symbol: "ESB"},
				currency.ESP: cldr.Currency{DisplayName: "španělská peseta", Symbol: "ESP"},
				currency.ETB: cldr.Currency{DisplayName: "etiopský birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "finská marka", Symbol: "FIM"},
				currency.FJD: cldr.Currency{DisplayName: "fidžijský dolar", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "falklandská libra", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "francouzský frank", Symbol: "FRF"},
				currency.GBP: cldr.Currency{DisplayName: "britská libra", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "gruzínské kuponové lari", Symbol: "GEK"},
				currency.GEL: cldr.Currency{DisplayName: "gruzínské lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "ghanský cedi (1979–2007)", Symbol: "GHC"},
				currency.GHS: cldr.Currency{DisplayName: "ghanský cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "gibraltarská libra", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "gambijský dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "guinejský frank", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "guinejský syli", Symbol: "GNS"},
				currency.GQE: cldr.Currency{DisplayName: "rovníkovoguinejský ekwele", Symbol: "GQE"},
				currency.GRD: cldr.Currency{DisplayName: "řecká drachma", Symbol: "GRD"},
				currency.GTQ: cldr.Currency{DisplayName: "guatemalský quetzal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "portugalskoguinejské escudo", Symbol: "GWE"},
				currency.GWP: cldr.Currency{DisplayName: "guinejsko-bissauské peso", Symbol: "GWP"},
				currency.GYD: cldr.Currency{DisplayName: "guyanský dolar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "hongkongský dolar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "honduraská lempira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "chorvatský dinár", Symbol: "HRD"},
				currency.HRK: cldr.Currency{DisplayName: "chorvatská kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "haitský gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "maďarský forint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "indonéská rupie", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "irská libra", Symbol: "IEP"},
				currency.ILP: cldr.Currency{DisplayName: "izraelská libra", Symbol: "ILP"},
				currency.ILR: cldr.Currency{DisplayName: "izraelský šekel (1980–1985)", Symbol: "ILR"},
				currency.ILS: cldr.Currency{DisplayName: "izraelský nový šekel", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "indická rupie", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "irácký dinár", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "íránský rijál", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "islandská koruna (1918–1981)", Symbol: "ISJ"},
				currency.ISK: cldr.Currency{DisplayName: "islandská koruna", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "italská lira", Symbol: "ITL"},
				currency.JMD: cldr.Currency{DisplayName: "jamajský dolar", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "jordánský dinár", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "japonský jen", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "keňský šilink", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "kyrgyzský som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "kambodžský riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "komorský frank", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "severokorejský won", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "jihokorejský hwan (1953–1962)", Symbol: "KRH"},
				currency.KRO: cldr.Currency{DisplayName: "jihokorejský won (1945–1953)", Symbol: "KRO"},
				currency.KRW: cldr.Currency{DisplayName: "jihokorejský won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "kuvajtský dinár", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "kajmanský dolar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "kazašské tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "laoský kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "libanonská libra", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "srílanská rupie", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "liberijský dolar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "lesothský loti", Symbol: "LSL"},
				currency.LTL: cldr.Currency{DisplayName: "litevský litas", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "litevský talonas", Symbol: "LTT"},
				currency.LUC: cldr.Currency{DisplayName: "lucemburský konvertibilní frank", Symbol: "LUC"},
				currency.LUF: cldr.Currency{DisplayName: "lucemburský frank", Symbol: "LUF"},
				currency.LUL: cldr.Currency{DisplayName: "lucemburský finanční frank", Symbol: "LUL"},
				currency.LVL: cldr.Currency{DisplayName: "lotyšský lat", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "lotyšský rubl", Symbol: "LVR"},
				currency.LYD: cldr.Currency{DisplayName: "libyjský dinár", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "marocký dinár", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "marocký frank", Symbol: "MAF"},
				currency.MCF: cldr.Currency{DisplayName: "monacký frank", Symbol: "MCF"},
				currency.MDC: cldr.Currency{DisplayName: "moldavský kupon", Symbol: "MDC"},
				currency.MDL: cldr.Currency{DisplayName: "moldavský leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "madagaskarský ariary", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "madagaskarský frank", Symbol: "MGF"},
				currency.MKD: cldr.Currency{DisplayName: "makedonský denár", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "makedonský denár (1992–1993)", Symbol: "MKN"},
				currency.MLF: cldr.Currency{DisplayName: "malijský frank", Symbol: "MLF"},
				currency.MMK: cldr.Currency{DisplayName: "myanmarský kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "mongolský tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "macajská pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "mauritánská ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "mauritánská ouguiya", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "maltská lira", Symbol: "MTL"},
				currency.MTP: cldr.Currency{DisplayName: "maltská libra", Symbol: "MTP"},
				currency.MUR: cldr.Currency{DisplayName: "mauricijská rupie", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "maledivská rupie (1947–1981)", Symbol: "MVP"},
				currency.MVR: cldr.Currency{DisplayName: "maledivská rupie", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "malawijská kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "mexické peso", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "mexické stříbrné peso (1861–1992)", Symbol: "MXP"},
				currency.MXV: cldr.Currency{DisplayName: "mexická investiční jednotka", Symbol: "MXV"},
				currency.MYR: cldr.Currency{DisplayName: "malajsijský ringgit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "mosambický escudo", Symbol: "MZE"},
				currency.MZM: cldr.Currency{DisplayName: "mosambický metical (1980–2006)", Symbol: "MZM"},
				currency.MZN: cldr.Currency{DisplayName: "mozambický metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "namibijský dolar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "nigerijská naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "nikaragujská córdoba (1988–1991)", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "nikaragujská córdoba", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "nizozemský gulden", Symbol: "NLG"},
				currency.NOK: cldr.Currency{DisplayName: "norská koruna", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "nepálská rupie", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "novozélandský dolar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ománský rijál", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "panamská balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "peruánská inti", Symbol: "PEI"},
				currency.PEN: cldr.Currency{DisplayName: "peruánský sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "peruánský sol (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "papuánská nová kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "filipínské peso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "pákistánská rupie", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "polský zlotý", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "polský zlotý (1950–1995)", Symbol: "PLZ"},
				currency.PTE: cldr.Currency{DisplayName: "portugalské escudo", Symbol: "PTE"},
				currency.PYG: cldr.Currency{DisplayName: "paraguajské guarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "katarský rijál", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "rhodéský dolar", Symbol: "RHD"},
				currency.ROL: cldr.Currency{DisplayName: "rumunské leu (1952–2006)", Symbol: "ROL"},
				currency.RON: cldr.Currency{DisplayName: "rumunský leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "srbský dinár", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "ruský rubl", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "ruský rubl (1991–1998)", Symbol: "RUR"},
				currency.RWF: cldr.Currency{DisplayName: "rwandský frank", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "saúdský rijál", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "šalamounský dolar", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "seychelská rupie", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "súdánský dinár (1992–2007)", Symbol: "SDD"},
				currency.SDG: cldr.Currency{DisplayName: "súdánská libra", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "súdánská libra (1957–1998)", Symbol: "SDP"},
				currency.SEK: cldr.Currency{DisplayName: "švédská koruna", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "singapurský dolar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "svatohelenská libra", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "slovinský tolar", Symbol: "SIT"},
				currency.SKK: cldr.Currency{DisplayName: "slovenská koruna", Symbol: "SKK"},
				currency.SLL: cldr.Currency{DisplayName: "sierro-leonský leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "somálský šilink", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "surinamský dolar", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "surinamský zlatý", Symbol: "SRG"},
				currency.SSP: cldr.Currency{DisplayName: "jihosúdánská libra", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "svatotomášská dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "svatotomášská dobra", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "sovětský rubl", Symbol: "SUR"},
				currency.SVC: cldr.Currency{DisplayName: "salvadorský colón", Symbol: "SVC"},
				currency.SYP: cldr.Currency{DisplayName: "syrská libra", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "svazijský lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "thajský baht", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "tádžický rubl", Symbol: "TJR"},
				currency.TJS: cldr.Currency{DisplayName: "tádžické somoni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "turkmenský manat (1993–2009)", Symbol: "TMM"},
				currency.TMT: cldr.Currency{DisplayName: "turkmenský manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "tuniský dinár", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "tonžská paanga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "timorské escudo", Symbol: "TPE"},
				currency.TRL: cldr.Currency{DisplayName: "turecká lira (1922–2005)", Symbol: "TRL"},
				currency.TRY: cldr.Currency{DisplayName: "turecká lira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "trinidadský dolar", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "tchajwanský dolar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "tanzanský šilink", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ukrajinská hřivna", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "ukrajinský karbovanec", Symbol: "UAK"},
				currency.UGS: cldr.Currency{DisplayName: "ugandský šilink (1966–1987)", Symbol: "UGS"},
				currency.UGX: cldr.Currency{DisplayName: "ugandský šilink", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "americký dolar", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "americký dolar (příští den)", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "americký dolar (týž den)", Symbol: "USS"},
				currency.UYI: cldr.Currency{DisplayName: "uruguayské peso (v indexovaných jednotkách)", Symbol: "UYI"},
				currency.UYP: cldr.Currency{DisplayName: "uruguayské peso (1975–1993)", Symbol: "UYP"},
				currency.UYU: cldr.Currency{DisplayName: "uruguayské peso", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "uzbecký sum", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "venezuelský bolívar (1871–2008)", Symbol: "VEB"},
				currency.VEF: cldr.Currency{DisplayName: "venezuelský bolívar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "venezuelský bolívar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "vietnamský dong", Symbol: "VND"},
				currency.VNN: cldr.Currency{DisplayName: "vietnamský dong (1978–1985)", Symbol: "VNN"},
				currency.VUV: cldr.Currency{DisplayName: "vanuatský vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "samojská tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "CFA/BEAC frank", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "stříbro", Symbol: "XAG"},
				currency.XAU: cldr.Currency{DisplayName: "zlato", Symbol: "XAU"},
				currency.XBA: cldr.Currency{DisplayName: "evropská smíšená jednotka", Symbol: "XBA"},
				currency.XBB: cldr.Currency{DisplayName: "evropská peněžní jednotka", Symbol: "XBB"},
				currency.XBC: cldr.Currency{DisplayName: "evropská jednotka účtu 9 (XBC)", Symbol: "XBC"},
				currency.XBD: cldr.Currency{DisplayName: "evropská jednotka účtu 17 (XBD)", Symbol: "XBD"},
				currency.XCD: cldr.Currency{DisplayName: "východokaribský dolar", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "SDR", Symbol: "XDR"},
				currency.XEU: cldr.Currency{DisplayName: "evropská měnová jednotka", Symbol: "ECU"},
				currency.XFO: cldr.Currency{DisplayName: "francouzský zlatý frank", Symbol: "XFO"},
				currency.XFU: cldr.Currency{DisplayName: "francouzský UIC frank", Symbol: "XFU"},
				currency.XOF: cldr.Currency{DisplayName: "CFA/BCEAO frank", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "palladium", Symbol: "XPD"},
				currency.XPF: cldr.Currency{DisplayName: "CFP frank", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "platina", Symbol: "XPT"},
				currency.XRE: cldr.Currency{DisplayName: "kód fondů RINET", Symbol: "XRE"},
				currency.XSU: cldr.Currency{DisplayName: "sucre", Symbol: "XSU"},
				currency.XTS: cldr.Currency{DisplayName: "kód zvlášť vyhrazený pro testovací účely", Symbol: "XTS"},
				currency.XUA: cldr.Currency{DisplayName: "", Symbol: "XUA"},
				currency.XXX: cldr.Currency{DisplayName: "neznámá měna", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "jemenský dinár", Symbol: "YDD"},
				currency.YER: cldr.Currency{DisplayName: "jemenský rijál", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "jugoslávský dinár (1966–1990)", Symbol: "YUD"},
				currency.YUM: cldr.Currency{DisplayName: "jugoslávský nový dinár (1994–2002)", Symbol: "YUM"},
				currency.YUN: cldr.Currency{DisplayName: "jugoslávský konvertibilní dinár (1990–1992)", Symbol: "YUN"},
				currency.YUR: cldr.Currency{DisplayName: "jugoslávský reformovaný dinár (1992–1993)", Symbol: "YUR"},
				currency.ZAL: cldr.Currency{DisplayName: "jihoafrický finanční rand", Symbol: "ZAL"},
				currency.ZAR: cldr.Currency{DisplayName: "jihoafrický rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "zambijská kwacha (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "zambijská kwacha", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "zairský nový zaire (1993–1998)", Symbol: "ZRN"},
				currency.ZRZ: cldr.Currency{DisplayName: "zairský zaire (1971–1993)", Symbol: "ZRZ"},
				currency.ZWD: cldr.Currency{DisplayName: "zimbabwský dolar (1980–2008)", Symbol: "ZWD"},
				currency.ZWL: cldr.Currency{DisplayName: "zimbabwský dolar (2009)", Symbol: "ZWL"},
				currency.ZWR: cldr.Currency{DisplayName: "zimbabwský dolar (2008)", Symbol: "ZWR"},
			},
		},
	}
}
