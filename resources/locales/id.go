// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_id() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Agu", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Maret", Apr: "April", May: "Mei", Jun: "Juni", Jul: "Juli", Aug: "Agustus", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Desember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Min", Mon: "Sen", Tue: "Sel", Wed: "Rab", Thu: "Kam", Fri: "Jum", Sat: "Sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "M", Mon: "S", Tue: "S", Wed: "R", Thu: "K", Fri: "J", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Min", Mon: "Sen", Tue: "Sel", Wed: "Rab", Thu: "Kam", Fri: "Jum", Sat: "Sab"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Minggu", Mon: "Senin", Tue: "Selasa", Wed: "Rabu", Thu: "Kamis", Fri: "Jumat", Sat: "Sabtu"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_id]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Peseta Andorra", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "Dirham Uni Emirat Arab", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Afgani Afganistan (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afgani Afganistan", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Lek Albania", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Dram Armenia", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Guilder Antilla Belanda", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza Angola", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Kwanza Angola (1977–1991)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Kwanza Baru Angola (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Kwanza Angola yang Disesuaikan Lagi (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Austral Argentina", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "Peso Ley Argentina (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "Peso Argentina (1881–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Peso Argentina (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Peso Argentina", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Schilling Austria", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Dolar Australia", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "Florin Aruba", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Manat Azerbaijan (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Manat Azerbaijan", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Dinar Bosnia-Herzegovina (1992–1994)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Mark Konvertibel Bosnia-Herzegovina", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "Dinar Baru Bosnia-Herzegovina (1994–1997)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "Dolar Barbados", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Taka Bangladesh", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Franc Belgia (konvertibel)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Franc Belgia", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Franc Belgia (keuangan)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Hard Lev Bulgaria", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "Socialist Lev Bulgaria", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Lev Bulgaria", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Lev Bulgaria (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Dinar Bahrain", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Franc Burundi", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Dolar Bermuda", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Dolar Brunei", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Boliviano Bolivia (1863–1963)", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "Peso Bolivia", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Mvdol Bolivia", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Cruzeiro Baru Brasil (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Cruzado Brasil (1986–1989)", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Cruzeiro Brasil (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Real Brasil", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Cruzado Baru Brasil (1989–1990)", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Cruzeiro Brasil (1993–1994)", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "Cruzeiro Brasil (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Dolar Bahama", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrum Bhutan", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Kyat Burma", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Pula Botswana", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Rubel Baru Belarus (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Rubel Belarusia", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Rubel Belarusia (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Dolar Belize", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Dolar Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Franc Kongo", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "Euro WIR", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Franc Swiss", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "Franc WIR", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "Escudo Cile", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Satuan Hitung (UF) Cile", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Peso Cile", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Yuan Tiongkok (luar negeri)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Tiongkok", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso Kolombia", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Unit Nilai Nyata Kolombia", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "Colon Kosta Rika", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Dinar Serbia (2002–2006)", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Hard Koruna Cheska", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Peso Konvertibel Kuba", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Peso Kuba", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo Tanjung Verde", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Pound Siprus", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Koruna Cheska", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Mark Jerman Timur", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Mark Jerman", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Franc Jibuti", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Krone Denmark", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Peso Dominika", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar Algeria", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Sucre Ekuador", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Satuan Nilai Tetap Ekuador", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Kroon Estonia", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Pound Mesir", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa Eritrea", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "Peseta Spanyol (akun)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Peseta Spanyol (konvertibel)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Peseta Spanyol", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr Etiopia", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Markka Finlandia", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Dolar Fiji", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Pound Kepulauan Falkland", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Franc Prancis", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Pound Inggris", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Kupon Larit Georgia", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Lari Georgia", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Cedi Ghana (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Cedi Ghana", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Pound Gibraltar", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi Gambia", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Franc Guinea", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Syli Guinea", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Ekuele Guinea Ekuatorial", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Drachma Yunani", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal Guatemala", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Escudo Guinea Portugal", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Peso Guinea-Bissau", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Dolar Guyana", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Dolar Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira Honduras", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Dinar Kroasia", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Kuna Kroasia", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Gourde Haiti", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Forint Hungaria", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Rupiah Indonesia", Symbol: "Rp"},
				currency.IEP: cldr.Currency{DisplayName: "Pound Irlandia", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Pound Israel", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "Shekel Israel", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Shekel Baru Israel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupee India", Symbol: "Rs"},
				currency.IQD: cldr.Currency{DisplayName: "Dinar Irak", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Rial Iran", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "Krona Islandia (1918–1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Krona Islandia", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Lira Italia", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Dolar Jamaika", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Dinar Yordania", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Yen Jepang", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilling Kenya", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Som Kirgistan", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Riel Kamboja", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Franc Komoro", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Won Korea Utara", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "Hwan Korea Selatan (1953–1962)", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "Won Korea Selatan (1945–1953)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "Won Korea Selatan", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinar Kuwait", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Dolar Kepulauan Cayman", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge Kazakstan", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Kip Laos", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Pound Lebanon", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Rupee Sri Lanka", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Dolar Liberia", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Loti Lesotho", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litas Lituania", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Talonas Lituania", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Franc Konvertibel Luksemburg", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Franc Luksemburg", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Financial Franc Luksemburg", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Lats Latvia", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Rubel Latvia", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Dinar Libya", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Dirham Maroko", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Franc Maroko", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "Franc Monegasque", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "Cupon Moldova", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Leu Moldova", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Ariary Madagaskar", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Franc Malagasi", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Denar Makedonia", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Denar Makedonia (1992–1993)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Franc Mali", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Kyat Myanmar", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik Mongolia", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca Makau", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya Mauritania (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya Mauritania", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Lira Malta", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Pound Malta", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupee Mauritius", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "Rufiyaa Maladewa (1947–1981)", Symbol: ""},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyaa Maladewa", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha Malawi", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Peso Meksiko", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Peso Silver Meksiko (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Unit Investasi Meksiko", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Ringgit Malaysia", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Escudo Mozambik", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Metical Mozambik (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Metical Mozambik", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Dolar Namibia", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Naira Nigeria", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Cordoba Nikaragua (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Cordoba Nikaragua", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Guilder Belanda", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Krone Norwegia", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Rupee Nepal", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Dolar Selandia Baru", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Rial Oman", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Balboa Panama", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Inti Peru", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Sol Peru", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Sol Peru (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kina Papua Nugini", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Peso Filipina", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Rupee Pakistan", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Polandia Zloty", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Zloty Polandia (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Escudo Portugal", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Guarani Paraguay", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Rial Qatar", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Dolar Rhodesia", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Leu Rumania (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Leu Rumania", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Dinar Serbia", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rubel Rusia", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Rubel Rusia (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Franc Rwanda", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal Arab Saudi", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Dolar Kepulauan Solomon", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Rupee Seychelles", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Dinar Sudan (1992–2007)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Pound Sudan", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Pound Sudan (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Krona Swedia", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Dolar Singapura", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Pound Saint Helena", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Tolar Slovenia", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Koruna Slovakia", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Leone Sierra Leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Shilling Somalia", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Dolar Suriname", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Guilder Suriname", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "Pound Sudan Selatan", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Dobra Sao Tome dan Principe (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Dobra Sao Tome dan Principe", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Rubel Soviet", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Colon El Savador", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Pound Suriah", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni Swaziland", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Baht Thailand", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Rubel Tajikistan", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Somoni Tajikistan", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Manat Turkmenistan (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Manat Turkimenistan", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Dinar Tunisia", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Paʻanga Tonga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Escudo Timor", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Lira Turki (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Lira Turki", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Dolar Trinidad dan Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Dolar Baru Taiwan", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilling Tanzania", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Hryvnia Ukraina", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Karbovanet Ukraina", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Shilling Uganda (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Shilling Uganda", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Dolar Amerika Serikat", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "Dolar AS (Hari berikutnya)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "Dolar AS (Hari yang sama)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "Peso Uruguay (Unit Diindeks)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Peso Uruguay (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Peso Uruguay", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Som Uzbekistan", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Bolivar Venezuela (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Bolivar Venezuela (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Bolivar Venezuela", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Dong Vietnam", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "Dong Vietnam (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "Vatu Vanuatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Tala Samoa", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Franc CFA BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Silver", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Emas", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Unit Gabungan Eropa", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Unit Keuangan Eropa", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Satuan Hitung Eropa (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Satuan Hitung Eropa (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Dolar Karibia Timur", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Hak Khusus Menggambar", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "Satuan Mata Uang Eropa", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Franc Gold Perancis", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Franc UIC Perancis", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "Franc CFA BCEAO", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "Franc CFP", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platinum", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "Dana RINET", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Kode Mata Uang Pengujian", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Mata Uang Tidak Dikenal", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "Dinar Yaman", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Rial Yaman", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Hard Dinar Yugoslavia (1966–1990)", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Dinar Baru Yugoslavia (1994–2002)", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Dinar Konvertibel Yugoslavia (1990–1992)", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "Dinar Reformasi Yugoslavia (1992–1993)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Rand Afrika Selatan (Keuangan)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Rand Afrika Selatan", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha Zambia", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Zaire Baru Zaire (1993–1998)", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Zaire Zaire (1971–1993)", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Dolar Zimbabwe (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Dolar Zimbabwe (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Dolar Zimbabwe (2008)", Symbol: ""},
			},
		},
	}
}
