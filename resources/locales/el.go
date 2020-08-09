// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func getLocale_el() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} - {0}", Long: "{1} - {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ιαν", Feb: "Φεβ", Mar: "Μάρ", Apr: "Απρ", May: "Μάι", Jun: "Ιούν", Jul: "Ιούλ", Aug: "Αύγ", Sep: "Σεπ", Oct: "Οκτ", Nov: "Νοέ", Dec: "Δεκ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Ι", Feb: "Φ", Mar: "Μ", Apr: "Α", May: "Μ", Jun: "Ι", Jul: "Ι", Aug: "Α", Sep: "Σ", Oct: "Ο", Nov: "Ν", Dec: "Δ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Ιανουάριος", Feb: "Φεβρουάριος", Mar: "Μάρτιος", Apr: "Απρίλιος", May: "Μάιος", Jun: "Ιούνιος", Jul: "Ιούλιος", Aug: "Αύγουστος", Sep: "Σεπτέμβριος", Oct: "Οκτώβριος", Nov: "Νοέμβριος", Dec: "Δεκέμβριος"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Κυρ", Mon: "Δευ", Tue: "Τρί", Wed: "Τετ", Thu: "Πέμ", Fri: "Παρ", Sat: "Σάβ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Κ", Mon: "Δ", Tue: "Τ", Wed: "Τ", Thu: "Π", Fri: "Π", Sat: "Σ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Κυ", Mon: "Δε", Tue: "Τρ", Wed: "Τε", Thu: "Πέ", Fri: "Πα", Sat: "Σά"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Κυριακή", Mon: "Δευτέρα", Tue: "Τρίτη", Wed: "Τετάρτη", Thu: "Πέμπτη", Fri: "Παρασκευή", Sat: "Σάββατο"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "π.μ.", PM: "μ.μ."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "πμ", PM: "μμ"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "π.μ.", PM: "μ.μ."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_el]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Πεσέτα Ανδόρας", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "Ντιράμ Ηνωμένων Αραβικών Εμιράτων", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Αφγανί Αφγανιστάν (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Αφγάνι Αφγανιστάν", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "Παλαιό λεκ Αλβανίας", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Λεκ Αλβανίας", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Ντραμ Αρμενίας", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Γκίλντα Ολλανδικών Αντιλλών", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Κουάνζα Ανγκόλας", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Κουάνζα Ανγκόλας (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Νέα Κουάνζα Ανγκόλας (1990–2000)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Ωστράλ Αργετινής", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "Πέσο λέι Αργετινής", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Πέσο Αργεντινής (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Πέσο Αργεντινής", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Σελίνι Αυστρίας", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Δολάριο Αυστραλίας", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Φλορίνι Αρούμπας", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Μανάτ Αζερμπαϊτζάν (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Μανάτ Αζερμπαϊτζάν", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Δηνάριο Βοσνίας-Ερζεγοβίνης", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Μετατρέψιμο Μάρκο Βοσνίας-Ερζεγοβίνης", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "Νέο δινάριο Βοσνίας-Ερζεγοβίνης", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "Δολάριο Μπαρμπέιντος", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Τάκα Μπαγκλαντές", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Φράγκο Βελγίου (μετατρέψιμο)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Φράγκο Βελγίου", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Φράγκο Βελγίου (οικονομικό)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Μεταλλικό Λεβ Βουλγαρίας", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "Σοσιαλιστικό λεβ Βουλγαρίας", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Λεβ Βουλγαρίας", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Παλαιό λεβ Βουλγαρίας", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Δηνάριο Μπαχρέιν", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Φράγκο Μπουρούντι", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Δολάριο Βερμούδων", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Δολάριο Μπρουνέι", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Μπολιβιάνο Βολιβίας", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Παλαιό βολιβιάνο Βολιβίας", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "Πέσο Βολιβίας", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Μβδολ Βολιβίας", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Νέο Κρουζιέρο Βραζιλίας (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Κρουζάντο Βραζιλίας", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Κρουζιέρο Βραζιλίας (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Ρεάλ Βραζιλίας", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Νέο Κρουζάντο Βραζιλίας", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Κρουζιέρο Βραζιλίας", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "Παλαιό κρουζέιρο Βραζιλίας", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Δολάριο Μπαχαμών", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Νγκούλτρουμ Μπουτάν", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Κιατ Βιρμανίας", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Πούλα Μποτσουάνας", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Νέο Ρούβλι Λευκορωσίας (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Ρούβλι Λευκορωσίας", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Ρούβλι Λευκορωσίας (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Δολάριο Μπελίζ", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Δολάριο Καναδά", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Φράγκο Κονγκό", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "Ευρώ WIR", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Φράγκο Ελβετίας", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "Φράγκο WIR", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "Εσκούδο Χιλής", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Ουνιδάδες ντε φομέντο Χιλής", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Πέσο Χιλής", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Γουάν Κίνας (υπεράκτιο)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "Δολάριο Λαϊκής Τράπεζας Κίνας", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Γουάν Κίνας", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Πέσο Κολομβίας", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Κολόν Κόστα Ρίκα", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Παλαιό Δηνάριο Σερβίας", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Σκληρή Κορόνα Τσεχοσλοβακίας", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Μετατρέψιμο πέσο Κούβας", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Πέσο Κούβας", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Εσκούδο Πράσινου Ακρωτηρίου", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Λίρα Κύπρου", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Κορόνα Τσεχίας", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Οστμάρκ Ανατολικής Γερμανίας", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Μάρκο Γερμανίας", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Φράγκο Τζιμπουτί", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Κορόνα Δανίας", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Πέσο Δομινικανής Δημοκρατίας", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Δηνάριο Αλγερίας", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Σούκρε Εκουαδόρ", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Κορόνα Εσθονίας", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Λίρα Αιγύπτου", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Νάκφα Ερυθραίας", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "πεσέτα Ισπανίας (λογαριασμός Α)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "πεσέτα Ισπανίας (μετατρέψιμος λογαριασμός)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Πεσέτα Ισπανίας", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Μπιρ Αιθιοπίας", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Ευρώ", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Μάρκο Φινλανδίας", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Δολάριο Φίτζι", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Λίρα Νήσων Φόκλαντ", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Φράγκο Γαλλίας", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Λίρα Στερλίνα Βρετανίας", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Κούπον Λάρι Γεωργίας", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Λάρι Γεωργίας", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Σέντι Γκάνας (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Σέντι Γκάνας", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Λίρα Γιβραλτάρ", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Νταλάσι Γκάμπιας", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Φράγκο Γουινέας", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Συλί Γουινέας", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Εκγουέλε Ισημερινής Γουινέας", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Δραχμή Ελλάδας", Symbol: "Δρχ"},
				currency.GTQ: cldr.Currency{DisplayName: "Κουετσάλ Γουατεμάλας", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Γκινέα Εσκούδο Πορτογαλίας", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Πέσο Γουινέας-Μπισάου", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Δολάριο Γουιάνας", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Δολάριο Χονγκ Κονγκ", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Λεμπίρα Ονδούρας", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Δηνάριο Κροατίας", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Κούνα Κροατίας", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Γκουρντ Αϊτής", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Φιορίνι Ουγγαρίας", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Ρουπία Ινδονησίας", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Λίρα Ιρλανδίας", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Λίρα Ισραήλ", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "παλιό σεκέλ Ισραήλ", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Νέο Σέκελ Ισραήλ", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Ρουπία Ινδίας", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Δηνάριο Ιράκ", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Ριάλ Ιράν", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "Παλιά κορόνα Ισλανδίας", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Κορόνα Ισλανδίας", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Λιρέτα Ιταλίας", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Δολάριο Τζαμάικας", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Δηνάριο Ιορδανίας", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Γιεν Ιαπωνίας", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Σελίνι Κένυας", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Σομ Κιργιζίας", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Ρίελ Καμπότζης", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Φράγκο Κομορών", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Γουόν Βόρειας Κορέας", Symbol: "KPW"},
				currency.KRO: cldr.Currency{DisplayName: "Παλιό γον Νότιας Κορέας", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "Γουόν Νότιας Κορέας", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Δηνάριο Κουβέιτ", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Δολάριο Νήσων Κέιμαν", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Τένγκε Καζακστάν", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Κιπ Λάος", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Λίρα Λιβάνου", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Ρουπία Σρι Λάνκα", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Δολάριο Λιβερίας", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Λότι Λεσότο", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Λίτα Λιθουανίας", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Ταλόνας Λιθουανίας", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Μετατρέψιμο Φράγκο Λουξεμβούργου", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Φράγκο Λουξεμβούργου", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Οικονομικό Φράγκο Λουξεμβούργου", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Λατς Λετονίας", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Ρούβλι Λετονίας", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Δηνάριο Λιβύης", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Ντιράμ Μαρόκου", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Φράγκο Μαρόκου", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "Φράγκο Μονακό", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "Κούπον Μολδαβίας", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Λέου Μολδαβίας", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Αριάρι Μαδαγασκάρης", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Φράγκο Μαδαγασκάρης", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Δηνάριο ΠΓΔΜ", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Παλιό δηνάριο ΠΓΔΜ", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Φράγκο Μαλί", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Κιάτ Μιανμάρ", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Τουγκρίκ Μογγολίας", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Πατάκα Μακάο", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Ουγκίγια Μαυριτανίας (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Ουγκίγια Μαυριτανίας", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Λιρέτα Μάλτας", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Λίρα Μάλτας", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Ρουπία Μαυρικίου", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Ρουφίγια Μαλδίβων", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Κουάτσα Μαλάουι", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Πέσο Μεξικού", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Ασημένιο Πέσο Μεξικού (1861–1992)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Ρινγκίτ Μαλαισίας", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Εσκούδο Μοζαμβίκης", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Παλαιό Μετικάλ Μοζαμβίκης", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Μετικάλ Μοζαμβίκης", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Δολάριο Ναμίμπιας", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Νάιρα Νιγηρίας", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Κόρδοβα Νικαράγουας", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Χρυσή Κόρδοβα Νικαράγουας", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Γκίλντα Ολλανδίας", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Κορόνα Νορβηγίας", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Ρουπία Νεπάλ", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Δολάριο Νέας Ζηλανδίας", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Ριάλ Ομάν", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Μπαλμπόα Παναμά", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Ίντι Περού", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Σολ Περού", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Σολ Περού (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Κίνα Παπούας Νέας Γουινέας", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Πέσο Φιλιππίνων", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Ρουπία Πακιστάν", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Ζλότι Πολωνίας", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Ζλότυ Πολωνίας (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Εσκούδο Πορτογαλίας", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Γκουαρανί Παραγουάης", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Ριάλ Κατάρ", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Δολάριο Ροδεσίας", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Λέι Ρουμανίας", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Λέου Ρουμανίας", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Δηνάριο Σερβίας", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Ρούβλι Ρωσίας", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Ρούβλι Ρωσίας (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Φράγκο Ρουάντας", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Ριάλ Σαουδικής Αραβίας", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Δολάριο Νήσων Σολομώντος", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Ρουπία Σεϋχελλών", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Δηνάριο Σουδάν", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Λίρα Σουδάν", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Παλαιά Λίρα Σουδάν", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Κορόνα Σουηδίας", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Δολάριο Σιγκαπούρης", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Λίρα Αγίας Ελένης", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Τόλαρ Σλοβενίας", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Κορόνα Σλοβενίας", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Λεόνε Σιέρα Λεόνε", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Σελίνι Σομαλίας", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Δολάριο Σουρινάμ", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Γκίλντα Σουρινάμ", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "Λίρα Νότιου Σουδάν", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Ντόμπρα Σάο Τομέ και Πρίνσιπε (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Ντόμπρα Σάο Τομέ και Πρίνσιπε", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Σοβιετικό Ρούβλι", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Κολόν Ελ Σαλβαδόρ", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Λίρα Συρίας", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Λιλανγκένι Σουαζιλάνδης", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Μπατ Ταϊλάνδης", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Ρούβλι Τατζικιστάν", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Σομόνι Τατζικιστάν", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Μανάτ Τουρκμενιστάν", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Μάνατ Τουρκμενιστάν", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Δηνάριο Τυνησίας", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Παάγκα Τόνγκα", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Εσκούδο Τιμόρ", Symbol: "TPE"},
				currency.TRL: cldr.Currency{DisplayName: "Παλιά Λίρα Τουρκίας", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Λίρα Τουρκίας", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Δολάριο Τρινιντάντ και Τομπάγκο", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Νέο δολάριο Ταϊβάν", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Σελίνι Τανζανίας", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Γρίβνα Ουκρανίας", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Καρμποβανέτς Ουκρανίας", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Σελίνι Ουγκάντας (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Σελίνι Ουγκάντας", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Δολάριο ΗΠΑ", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "Δολάριο ΗΠΑ (επόμενη ημέρα)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "Δολάριο ΗΠΑ (ίδια ημέρα)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Πέσο Ουρουγουάης (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Πέσο Ουρουγουάης", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Σομ Ουζμπεκιστάν", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Μπολιβάρ Βενεζουέλας (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Μπολιβάρ Βενεζουέλας (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Μπολιβάρ Βενεζουέλας", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Ντονγκ Βιετνάμ", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "Παλαιό ντονγκ Βιετνάμ", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "Βατού Βανουάτου", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Τάλα Σαμόα", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Φράγκο CFA Κεντρικής Αφρικής", Symbol: "FCFA"},
				currency.XBA: cldr.Currency{DisplayName: "Ευρωπαϊκή Σύνθετη Μονάδα", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Ευρωπαϊκή Νομισματική Μονάδα", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Ευρωπαϊκή μονάδα λογαριασμού (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Ευρωπαϊκή μονάδα λογαριασμού (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Δολάριο Ανατολικής Καραϊβικής", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Ειδικά Δικαιώματα Ανάληψης", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "Ευρωπαϊκή Συναλλαγματική Μονάδα", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Χρυσό Φράγκο Γαλλίας", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "UIC-Φράγκο Γαλλίας", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "Φράγκο CFA Δυτικής Αφρικής", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Φράγκο CFP", Symbol: "CFPF"},
				currency.XRE: cldr.Currency{DisplayName: "Ταμείο RINET", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Άγνωστο νόμισμα", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Δηνάριο Υεμένης", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Ριάλ Υεμένης", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Μεταλλικό Δηνάριο Γιουγκοσλαβίας", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Νέο Δηνάριο Γιουγκοσλαβίας", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Μετατρέψιμο Δηνάριο Γιουγκοσλαβίας", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "Αναμορφωμένο δηνάριο Γιουγκοσλαβίας", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Ραντ Νότιας Αφρικής (οικονομικό)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Ραντ Νότιας Αφρικής", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Κουάνζα Ζαΐρ (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Κουάτσα Ζάμπιας", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Νέο Ζαΐρ Ζαΐρ", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Ζαΐρ Ζαΐρ", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Δολάριο Ζιμπάμπουε", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Δολάριο Ζιμπάμπουε (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Δολάριο Ζιμπάμπουε (2008)", Symbol: ""},
			},
		},
	}
}
