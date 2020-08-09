// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_en_AU() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apr", May: "May", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Oct", Nov: "Nov", Dec: "Dec"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "January", Feb: "February", Mar: "March", Apr: "April", May: "May", Jun: "June", Jul: "July", Aug: "August", Sep: "September", Oct: "October", Nov: "November", Dec: "December"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sun", Mon: "Mon", Tue: "Tue", Wed: "Wed", Thu: "Thu", Fri: "Fri", Sat: "Sat"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Su.", Mon: "M.", Tue: "Tu.", Wed: "W.", Thu: "Th.", Fri: "F.", Sat: "Sa."}, Short: cldr.CalendarDayFormatNameValue{Sun: "Su", Mon: "Mon", Tue: "Tu", Wed: "Wed", Thu: "Th", Fri: "Fri", Sat: "Sat"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sunday", Mon: "Monday", Tue: "Tuesday", Wed: "Wednesday", Thu: "Thursday", Fri: "Friday", Sat: "Saturday"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_en]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorran Peseta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "United Arab Emirates Dirham", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Afghan Afghani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afghan Afghani", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "Albanian Lek (1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Albanian Lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Armenian Dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Netherlands Antillean Guilder", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Angolan Kwanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Angolan Kwanza (1977–1991)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Angolan New Kwanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Angolan Readjusted Kwanza (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Argentine Austral", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "Argentine Peso Ley (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "Argentine Peso (1881–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Argentine Peso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Argentine Peso", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Austrian Schilling", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Australian Dollar", Symbol: "$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruban Florin", Symbol: ""},
				currency.AZM: cldr.Currency{DisplayName: "Azerbaijani Manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Azerbaijani Manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Bosnia-Herzegovina Dinar (1992–1994)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Bosnia-Herzegovina Convertible Marka", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "Bosnia-Herzegovina New Dinar (1994–1997)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "Barbados Dollar", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladeshi Taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Belgian Franc (convertible)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Belgian Franc", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Belgian Franc (financial)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Bulgarian Hard Lev", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "Bulgarian Socialist Lev", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Bulgarian Lev", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Bulgarian Lev (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Bahraini Dinar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burundian Franc", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda Dollar", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Brunei Dollar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Bolivian Boliviano (1863–1963)", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "Bolivian Peso", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Bolivian Mvdol", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Brazilian New Cruzeiro (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Brazilian Cruzado (1986–1989)", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Brazilian Cruzeiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Brazilian Real", Symbol: "BRL"},
				currency.BRN: cldr.Currency{DisplayName: "Brazilian New Cruzado (1989–1990)", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Brazilian Cruzeiro (1993–1994)", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "Brazilian Cruzeiro (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Bahamian Dollar", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Bhutanese Ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Burmese Kyat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Botswanan Pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Belarusian Ruble (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Belarusian Ruble", Symbol: "р."},
				currency.BYR: cldr.Currency{DisplayName: "Belarusian Ruble (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Belize Dollar", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Canadian Dollar", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "Congolese Franc", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR Euro", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Swiss Franc", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR Franc", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "Chilean Escudo", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Chilean Unit of Account (UF)", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Chilean Peso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "CNH", Symbol: ""},
				currency.CNX: cldr.Currency{DisplayName: "Chinese People’s Bank Dollar", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Chinese Yuan", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "Colombian Peso", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Colombian Real Value Unit", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "Costa Rican Colón", Symbol: "₡"},
				currency.CSD: cldr.Currency{DisplayName: "Serbian Dinar (2002–2006)", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Czechoslovak Hard Koruna", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Cuban Convertible Peso", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Cuban Peso", Symbol: "₱"},
				currency.CVE: cldr.Currency{DisplayName: "Cape Verdean Escudo", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Cypriot Pound", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Czech Koruna", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "East German Mark", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "German Mark", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Djiboutian Franc", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Danish Krone", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Dominican Peso", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Algerian Dinar", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Ecuadorian Sucre", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Ecuadorian Unit of Constant Value", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Estonian Kroon", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Egyptian Pound", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritrean Nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "Spanish Peseta (A account)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Spanish Peseta (convertible account)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Spanish Peseta", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ethiopian Birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "EUR"},
				currency.FIM: cldr.Currency{DisplayName: "Finnish Markka", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Fijian Dollar", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Falkland Islands Pound", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "French Franc", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "British Pound", Symbol: "GBP"},
				currency.GEK: cldr.Currency{DisplayName: "Georgian Kupon Larit", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Georgian lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Ghanaian Cedi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Ghanaian Cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltar Pound", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gambian Dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Guinean Franc", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Guinean Syli", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Equatorial Guinean Ekwele", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Greek Drachma", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Guatemalan Quetzal", Symbol: "Q"},
				currency.GWE: cldr.Currency{DisplayName: "Portuguese Guinea Escudo", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Guinea-Bissau Peso", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Guyanaese Dollar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hong Kong Dollar", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "Honduran Lempira", Symbol: "L"},
				currency.HRD: cldr.Currency{DisplayName: "Croatian Dinar", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Croatian Kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Haitian Gourde", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Hungarian Forint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indonesian Rupiah", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Irish Pound", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Israeli Pound", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "Israeli Shekel (1980–1985)", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Israeli Shekel", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "Indian Rupee", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "Iraqi Dinar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Iranian Rial", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "Icelandic Króna (1918–1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Icelandic Króna", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Italian Lira", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Jamaican Dollar", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Jordanian Dinar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Japanese Yen", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "Kenyan Shilling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Kyrgystani Som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Cambodian Riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Comorian Franc", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "North Korean Won", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "South Korean Hwan (1953–1962)", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "South Korean Won (1945–1953)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "South Korean Won", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "Kuwaiti Dinar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Cayman Islands Dollar", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Kazakhstani tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laotian kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Lebanese Pound", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Sri Lankan Rupee", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Liberian Dollar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Lesotho Loti", Symbol: "LSL"},
				currency.LTL: cldr.Currency{DisplayName: "Lithuanian Litas", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "Lithuanian Talonas", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Luxembourgian Convertible Franc", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Luxembourgian Franc", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Luxembourg Financial Franc", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Latvian Lats", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "Latvian Ruble", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Libyan Dinar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Moroccan Dirham", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Moroccan Franc", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "Monegasque Franc", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "Moldovan Cupon", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Moldovan Leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Malagasy Ariary", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Malagasy Franc", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Macedonian denar", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Macedonian Denar (1992–1993)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Malian Franc", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Myanmar Kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mongolian Tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Macanese Pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Mauritanian Ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Mauritanian Ouguiya", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "Maltese Lira", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Maltese Pound", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Mauritian Rupee", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "Maldivian Rupee (1947–1981)", Symbol: ""},
				currency.MVR: cldr.Currency{DisplayName: "Maldivian Rufiyaa", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malawian Kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Mexican Peso", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "Mexican Silver Peso (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Mexican Investment Unit", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Malaysian Ringgit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Mozambican Escudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Mozambican Metical (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Mozambican Metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namibian Dollar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Nigerian Naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Nicaraguan Córdoba (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Nicaraguan Córdoba", Symbol: "C$"},
				currency.NLG: cldr.Currency{DisplayName: "Dutch Guilder", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Norwegian Krone", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nepalese Rupee", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "New Zealand Dollar", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "Omani Rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Panamanian Balboa", Symbol: ""},
				currency.PEI: cldr.Currency{DisplayName: "Peruvian Inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Peruvian Sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Peruvian Sol (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Papua New Guinean kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Philippine Piso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistani Rupee", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Polish Zloty", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Polish Zloty (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Portuguese Escudo", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Paraguayan Guarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Qatari Riyal", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Rhodesian Dollar", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Romanian Leu (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Romanian Leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Serbian Dinar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Russian Ruble", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Russian Ruble (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Rwandan Franc", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudi Riyal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Solomon Islands Dollar", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Seychellois Rupee", Symbol: "Rs"},
				currency.SDD: cldr.Currency{DisplayName: "Sudanese Dinar (1992–2007)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Sudanese Pound", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Sudanese Pound (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Swedish Krona", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singapore Dollar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "St. Helena Pound", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Slovenian Tolar", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Slovak Koruna", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Sierra Leonean Leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Somali Shilling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Suriname Dollar", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Surinamese Guilder", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "South Sudanese Pound", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "São Tomé & Príncipe Dobra (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "São Tomé & Príncipe Dobra", Symbol: "Db"},
				currency.SUR: cldr.Currency{DisplayName: "Soviet Rouble", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Salvadoran Colón", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Syrian Pound", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Swazi Lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Thai Baht", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Tajikistani Ruble", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Tajikistani Somoni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Turkmenistani Manat (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Turkmenistani Manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tunisian Dinar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tongan Paʻanga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Timorese Escudo", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Turkish Lira (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Turkish lira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad & Tobago Dollar", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "New Taiwan Dollar", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzanian Shilling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Ukrainian Hryvnia", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Ukrainian Karbovanets", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Ugandan Shilling (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Ugandan Shilling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "US Dollar", Symbol: "USD"},
				currency.USN: cldr.Currency{DisplayName: "US Dollar (Next day)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "US Dollar (Same day)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "Uruguayan Peso (Indexed Units)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Uruguayan Peso (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Peso Uruguayo", Symbol: "UYU"},
				currency.UYW: cldr.Currency{DisplayName: "Uruguayan Nominal Wage Index Unit", Symbol: ""},
				currency.UZS: cldr.Currency{DisplayName: "Uzbekistani som", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Venezuelan Bolívar (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Venezuelan bolívar", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "VES", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Vietnamese dong", Symbol: "VND"},
				currency.VNN: cldr.Currency{DisplayName: "Vietnamese Dong (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatu Vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samoan tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Central African CFA Franc", Symbol: "XAF"},
				currency.XAG: cldr.Currency{DisplayName: "Silver", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Gold", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "European Composite Unit", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "European Monetary Unit", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "European Unit of Account (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "European Unit of Account (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "East Caribbean Dollar", Symbol: "XCD"},
				currency.XDR: cldr.Currency{DisplayName: "Special Drawing Rights", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "European Currency Unit", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "French Gold Franc", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "French UIC-Franc", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "West African CFA Franc", Symbol: "XOF"},
				currency.XPD: cldr.Currency{DisplayName: "Palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP Franc", Symbol: "CFP"},
				currency.XPT: cldr.Currency{DisplayName: "Platinum", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET Funds", Symbol: ""},
				currency.XSU: cldr.Currency{DisplayName: "Sucre", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Testing Currency Code", Symbol: ""},
				currency.XUA: cldr.Currency{DisplayName: "ADB Unit of Account", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Unknown Currency", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Yemeni Dinar", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Yemeni Rial", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Yugoslavian Hard Dinar (1966–1990)", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Yugoslavian New Dinar (1994–2002)", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Yugoslavian Convertible Dinar (1990–1992)", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "Yugoslavian Reformed Dinar (1992–1993)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "South African Rand (financial)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "South African Rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambian Kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Zambian Kwacha", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Zairean New Zaire (1993–1998)", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Zairean Zaire (1971–1993)", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabwean Dollar (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Zimbabwean Dollar (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Zimbabwean Dollar (2008)", Symbol: ""},
			},
		},
	}
}
