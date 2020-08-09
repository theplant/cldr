// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_lb_LU() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "lb_LU",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mäe", Apr: "Abr", May: "Mee", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dez"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januar", Feb: "Februar", Mar: "Mäerz", Apr: "Abrëll", May: "Mee", Jun: "Juni", Jul: "Juli", Aug: "August", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Dezember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Son", Mon: "Méi", Tue: "Dën", Wed: "Mët", Thu: "Don", Fri: "Fre", Sat: "Sam"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "M", Thu: "D", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "So.", Mon: "Mé.", Tue: "Dë.", Wed: "Më.", Thu: "Do.", Fri: "Fr.", Sat: "Sa."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sonndeg", Mon: "Méindeg", Tue: "Dënschdeg", Wed: "Mëttwoch", Thu: "Donneschdeg", Fri: "Freideg", Sat: "Samschdeg"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "moies", PM: "nomëttes"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "mo.", PM: "nomë."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "moies", PM: "nomëttes"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_lb]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorranesch Peseta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "VAE-Dirham", Symbol: ""},
				currency.AFA: cldr.Currency{DisplayName: "Afghanesch Afghani (1927–2002)", Symbol: "AFA"},
				currency.AFN: cldr.Currency{DisplayName: "Afghanesch Afghani", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Albanesche Lek", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Armeneschen Dram", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "Antillen-Gulden", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Angolanesche Kwanza", Symbol: "Kz"},
				currency.AOK: cldr.Currency{DisplayName: "Angolanesche Kwanza (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Angolaneschen Neie Kwanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Angolanesche Kwanza Reajustado (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Argentineschen Austral", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Argentinesche Peso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Argentinesche Peso", Symbol: "$"},
				currency.ATS: cldr.Currency{DisplayName: "Éisträichesche Schilling", Symbol: "öS"},
				currency.AUD: cldr.Currency{DisplayName: "Australeschen Dollar", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruba-Florin", Symbol: ""},
				currency.AZM: cldr.Currency{DisplayName: "Aserbaidschan-Manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Aserbaidschan-Manat", Symbol: ""},
				currency.BAD: cldr.Currency{DisplayName: "Bosnien an Herzegowina Dinar (1992–1994)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Bosnien an Herzegowina Konvertéierbar Mark", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbados-Dollar", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladesch-Taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Belsche Frang (konvertibel)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Belsche Frang", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Belsche Finanz-Frang", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Bulgaresch Lew (1962–1999)", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Bulgaresch Lew", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Bahrain-Dinar", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Burundi-Frang", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda-Dollar", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Brunei-Dollar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Bolivianesche Boliviano", Symbol: "Bs"},
				currency.BOP: cldr.Currency{DisplayName: "Bolivianesche Peso", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Bolivianseche Mvdol", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Brasilianesche Cruzeiro Novo (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Brasilianesche Cruzado (1986–1989)", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Brasilianesche Cruzeiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Brasilianesche Real", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Brasilianesche Cruzado Novo (1989–1990)", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Brasilianesche Cruzeiro (1993–1994)", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "Brasilianesche Cruzeiro (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Bahama-Dollar", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Bhutan-Ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Birmanesche Kyat", Symbol: "BUK"},
				currency.BWP: cldr.Currency{DisplayName: "Botswanesch Pula", Symbol: "P"},
				currency.BYB: cldr.Currency{DisplayName: "Wäissrussesche Rubel (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Wäissrussesche Rubel", Symbol: "р."},
				currency.BYR: cldr.Currency{DisplayName: "Wäissrussesche Rubel (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Belize-Dollar", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Kanadeschen Dollar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kongo-Frang", Symbol: ""},
				currency.CHE: cldr.Currency{DisplayName: "WIR-Euro", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Schwäizer Frang", Symbol: ""},
				currency.CHW: cldr.Currency{DisplayName: "WIR-Frang", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Chileneschen Unidad de Fomento", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Chilenesche Peso", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Renminbi Yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolumbianesche Peso", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Costa-Rica-Colón", Symbol: "₡"},
				currency.CSD: cldr.Currency{DisplayName: "Serbeschen Dinar (2002–2006)", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Tschechoslowakesch Kroun", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Kubanesche Peso (konvertibel)", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Kubanesche Peso", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Kap-Verde-Escudo", Symbol: ""},
				currency.CYP: cldr.Currency{DisplayName: "Zypern-Pond", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Tschechesch Kroun", Symbol: "Kč"},
				currency.DDM: cldr.Currency{DisplayName: "DDR-Mark", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Däitsch Mark", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Dschibuti-Frang", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Dänesch Kroun", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Dominikanesche Peso", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Algereschen Dinar", Symbol: ""},
				currency.ECS: cldr.Currency{DisplayName: "Ecuadorianesche Sucre", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Verrechnungseenheete fir Ecuador", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Estnesch Kroun", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Egyptescht Pond", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Eritréieschen Nakfa", Symbol: ""},
				currency.ESA: cldr.Currency{DisplayName: "Spuenesch Peseta (A–Konten)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Spuenesch Peseta (konvertibel)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Spuenesch Peseta", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ethiopescht Birr", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Finnesch Mark", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Fidschi-Dollar", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Falkland-Pond", Symbol: "£"},
				currency.FRF: cldr.Currency{DisplayName: "Franséische Frang", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Britescht Pond", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Georgesche Kupon Larit", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Georgesche Lari", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Ghanaeschen Cedi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Ghanaeschen Cedi", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltar-Pond", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Gambia-Dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Guinea-Frang", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Guinéiesche Syli", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Equatorialguinea-Ekwele", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Griichesch Drachme", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Guatemaltekesche Quetzal", Symbol: "Q"},
				currency.GWE: cldr.Currency{DisplayName: "Portugisesch-Guinea Escudo", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Guinea-Bissau Peso", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Guyana-Dollar", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Hong-Kong-Dollar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Honduras-Lempira", Symbol: "L"},
				currency.HRD: cldr.Currency{DisplayName: "Kroateschen Dinar", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Kroatesche Kuna", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Haitianesch Gourde", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Ungaresche Forint", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Indonesesch Rupiah", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Irescht Pond", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Israelescht Pond", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Israeleschen Neie Schekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Indesch Rupie", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Irakeschen Dinar", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Iranesch Rial", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Islännesch Kroun", Symbol: "kr"},
				currency.ITL: cldr.Currency{DisplayName: "Italienesch Lira", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Jamaika-Dollar", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Jordaneschen Dinar", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Japanesche Yen", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "Kenia-Schilling", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Kirgisesche Som", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "Kambodschanesche Riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Komore-Frang", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Nordkoreanesche Won", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Südkoreanesche Won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Kuwait-Dinar", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Kaiman-Dollar", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Kasacheschen Tenge", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Laoteschen Kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Libanesescht Pond", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Sri-Lanka-Rupie", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Liberianeschen Dollar", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litauesche Litas", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "Litaueschen Talonas", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Lëtzebuerger Frang (konvertibel)", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Lëtzebuerger Frang", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Lëtzebuerger Finanz-Frang", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Lettesche Lats", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "Lettesche Rubel", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Libeschen Dinar", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Marokkaneschen Dirham", Symbol: ""},
				currency.MAF: cldr.Currency{DisplayName: "Marokkanesche Frang", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Moldawesche Leu", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Madagaskar-Ariary", Symbol: "Ar"},
				currency.MGF: cldr.Currency{DisplayName: "Madagaskar-Frang", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Mazedoneschen Denar", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Malesche Frang", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Myanmaresche Kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mongoleschen Tögrög", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Macau-Pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Mauretaneschen Ouguiya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Mauretaneschen Ouguiya", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "Maltesesch Lira", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Maltesescht Pond", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Mauritius-Rupie", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Maldiven-Rupie", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malawi-Kwacha", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Mexikanesche Peso", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Mexikanesche Sëlwer-Peso (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Mexikaneschen Unidad de Inversion (UDI)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Malayseschen Ringgit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Mosambikaneschen Escudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Mosambikanesche Metical (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Mosambikanesche Metical", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Namibia-Dollar", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nigerianeschen Naira", Symbol: "₦"},
				currency.NIC: cldr.Currency{DisplayName: "Nicaraguanesche Córdoba (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Nicaraguanesche Córdoba", Symbol: "C$"},
				currency.NLG: cldr.Currency{DisplayName: "Hollännesche Gulden", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Norwegesch Kroun", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Nepalesesch Rupie", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Neiséiland-Dollar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Omanesche Rial", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Panamaesche Balboa", Symbol: ""},
				currency.PEI: cldr.Currency{DisplayName: "Peruaneschen Inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Peruaneschen Sol", Symbol: ""},
				currency.PES: cldr.Currency{DisplayName: "Peruaneschen Sol (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Papua-Neiguinéiesche Kina", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Philippinnesche Peso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistanesch Rupie", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Polneschen Zloty", Symbol: "zł"},
				currency.PLZ: cldr.Currency{DisplayName: "Polneschen Zloty (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Portugiseschen Escudo", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Paraguayeschen Guaraní", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Katar-Riyal", Symbol: ""},
				currency.RHD: cldr.Currency{DisplayName: "Rhodeseschen Dollar", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Rumänesche Leu (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Rumänesche Leu", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Serbeschen Dinar", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Russesche Rubel", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "Russesche Rubel (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Ruanda-Frang", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudi-Rial", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Salomonen-Dollar", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Seychellen-Rupie", Symbol: ""},
				currency.SDD: cldr.Currency{DisplayName: "Sudaneseschen Dinar (1992–2007)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Sudanesescht Pond", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Sudanesescht Pond (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Schwedesch Kroun", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Singapur-Dollar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "St. Helena-Pond", Symbol: "£"},
				currency.SIT: cldr.Currency{DisplayName: "Sloweneschen Tolar", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Slowakesch Kroun", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Sierra-leonesche Leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Somalia-Schilling", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Surinameschen Dollar", Symbol: "$"},
				currency.SRG: cldr.Currency{DisplayName: "Surinamesche Gulden", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "Südsudanesescht Pond", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "São-toméeschen Dobra (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "São-toméeschen Dobra", Symbol: "Db"},
				currency.SUR: cldr.Currency{DisplayName: "Sowjetesche Rubel", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "El-Salvador-Colón", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Syrescht Pond", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Swasilännesche Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Thailännesche Baht", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Tadschikistan-Rubel", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Tadschikistan-Somoni", Symbol: ""},
				currency.TMM: cldr.Currency{DisplayName: "Turkmenistan-Manat (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Turkmenistan-Manat", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Tuneseschen Dinar", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Tongaeschen Paʻanga", Symbol: "T$"},
				currency.TPE: cldr.Currency{DisplayName: "Timor-Escudo", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Tierkesch Lira (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Tierkesch Lira", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad-an-Tobago-Dollar", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Neien Taiwan-Dollar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tansania-Schilling", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Ukraineschen Hrywnja", Symbol: "₴"},
				currency.UAK: cldr.Currency{DisplayName: "Ukrainesche Karbovanetz", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Uganda-Schilling (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Uganda-Schilling", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "US-Dollar", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "US Dollar (Nächsten Dag)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "US Dollar (Selwechten Dag)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Uruguayesche Peso (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Uruguayesche Peso", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Usbekistan-Sum", Symbol: ""},
				currency.VEB: cldr.Currency{DisplayName: "Venezolanesche Bolívar (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Venezolanesche Bolívar (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Venezolanesche Bolívar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Vietnameseschen Dong", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatu-Vatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Samoaneschen Tala", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "CFA-Frang (BEAC)", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Onze Sëlwer", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Onze Gold", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Europäesch Rechnungseenheet", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Europäesch Währungseenheet (XBB)", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Europäesch Rechnungseenheet (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Europäesch Rechnungseenheet (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Ostkaribeschen Dollar", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Sonnerzéiungsrecht", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "Europäesch Währungseenheet (XEU)", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Franséische Gold-Frang", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Franséischen UIC-Frang", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "CFA-Frang (BCEAO)", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Onz Palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP-Frang", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Onz Platin", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET Funds", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Testwährung", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Onbekannt Währung", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Jemen-Dinar", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Jemen-Rial", Symbol: ""},
				currency.YUD: cldr.Currency{DisplayName: "Jugoslaweschen Dinar (1966–1990)", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Jugoslaweschen Neien Dinar (1994–2002)", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Jugoslaweschen Dinar (konvertibel)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Südafrikanesche Rand (Finanz)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Südafrikanesche Rand", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha", Symbol: "ZK"},
				currency.ZRN: cldr.Currency{DisplayName: "Zaire-Neien Zaïre (1993–1998)", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Zaire-Zaïre (1971–1993)", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Simbabwe-Dollar (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Simbabwe-Dollar (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Simbabwe-Dollar (2008)", Symbol: ""},
			},
		},
	}
}
