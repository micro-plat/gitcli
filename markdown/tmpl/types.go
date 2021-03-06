package tmpl

var tp2mysql = map[string]string{
	"^date$":                      "datetime",
	"^datetime$":                  "datetime",
	"^timestamp$":                 "datetime",
	"^decimal$":                   "decimal",
	"^float$":                     "float",
	"^int$":                       "int",
	"^number\\([1-2]{1}\\)$":      "tinyint",
	"^number\\([3-9]{1}\\)$":      "int",
	"^number\\(10\\)$":            "int",
	"^number\\(1[1-9]{1}\\)$":     "bigint",
	"^number\\(2[0-9]{1}\\)$":     "bigint",
	"^number\\((\\d+),(\\d+)\\)$": "decimal(*)",
	"^varchar\\((\\d+)\\)$":       "varchar(*)",
	"^varchar2\\((\\d+)\\)$":      "varchar(*)",
	"^string$":                    "tinytext",
	"^text$":                      "text",
	"^longtext$":                  "longtext",
}
var def2mysql = []map[string]string{
	{
		"^$":         "",
		"^-$":        "default '-'",
		"^seq$":      "",
		"^sysdate$":  "default current_timestamp",
		"^([0-9]+)$": "default *",
	},
	{
		"^(.+)$": "default '*'",
	},
}

var any2code = map[string]string{
	"^date$":                      "time.Time",
	"^datetime$":                  "time.Time",
	"^timestamp$":                 "time.Time",
	"^decimal$":                   "types.Decimal",
	"^float$":                     "types.Decimal",
	"^int$":                       "int",
	"^number\\([1-2]{1}\\)$":      "int",
	"^number\\([3-9]{1}\\)$":      "int",
	"^number\\(10\\)$":            "int",
	"^number\\(1[1-9]{1}\\)$":     "int64",
	"^number\\(2[0-9]{1}\\)$":     "int64",
	"^number\\((\\d+),(\\d+)\\)$": "types.Decimal",
	"^varchar\\(\\d+\\)$":         "string",
	"^varchar2\\(\\d+\\)$":        "string",
	"^string$":                    "string",
	"^text$":                      "string",
	"^longtext$":                  "string",
}

const idx = "[^idx\\((\\w+)[,]?[\\d]?\\)"

var keywordMatch = []string{"^\\w*%s\\w*$", ",\\w*%s\\w*,", "^\\w*%s\\w*,", ",\\w*%s\\w*$"}

var cons = map[string][]string{
	"pk":  {"[^pk]?[,]?pk[,]?[^pk]?"},
	"seq": {"[^seq]?[,]?seq[,]?[^seq]?"},
	"di":  {"[^di]?[,]?di[,]?[^di]?"},
	"dn":  {"[^dn]?[,]?dn[,]?[^dn]?"},
	"sl":  {"^sl\\(\\w+\\)$", ",sl\\(\\w+\\),", "^sl\\(\\w+\\),", ",sl\\(\\w+\\)$", "^%s$", ",%s,", "^%s,", ",%s$"},
	"slm": {"^slm\\(\\w+\\)$", ",slm\\(\\w+\\),", "^slm\\(\\w+\\),", ",slm\\(\\w+\\)$", "^%s$", ",%s,", "^%s,", ",%s$"},
	"rd":  {"^rd\\(\\w+\\)$", ",rd\\(\\w+\\),", "^rd\\(\\w+\\),", ",rd\\(\\w+\\)$", "^%s$", ",%s,", "^%s,", ",%s$"},
	"cb":  {"^cb\\(\\w+\\)$", ",cb\\(\\w+\\),", "^cb\\(\\w+\\),", ",cb\\(\\w+\\)$", "^%s$", ",%s,", "^%s,", ",%s$"},
	"ta":  {"^ta\\(\\w+\\)$", ",ta\\(\\w+\\),", "^ta\\(\\w+\\),", ",ta\\(\\w+\\)$", "^%s$", ",%s,", "^%s,", ",%s$"},
	"cc":  {"^cc\\(\\w+\\)$", ",cc\\(\\w+\\),", "^cc\\(\\w+\\),", ",cc\\(\\w+\\)$", "^%s$", ",%s,", "^%s,", ",%s$"},
	"idx": {"^idx\\((\\w+)[,]?([\\d]?)\\)$", ",idx\\((\\w+)[,]?([\\d]?)\\),", "^idx\\((\\w+)[,]?([\\d]?)\\),", ",idx\\((\\w+)[,]?([\\d]?)\\)$"},
	"unq": {"^unq\\((\\w+)[,]?([\\d]?)\\)$", ",unq\\((\\w+)[,]?([\\d]?)\\),", "^unq\\((\\w+)[,]?([\\d]?)\\),", ",unq\\((\\w+)[,]?([\\d]?)\\)$"},
	"d":   {"^d\\(\\w+(.+?)\\)$", ",d\\(\\w+(.+?)\\),", "^d\\(\\w+(.+?)\\),", ",d\\(\\w+(.+?)\\)$", "^%s$", ",%s,", "^%s,", ",%s$"},
	"c":   {"^c\\(\\w+(.+?)\\)$", ",c\\(\\w+(.+?)\\),", "^c\\(\\w+(.+?)\\),", ",c\\(\\w+(.+?)\\)$", "^%s$", ",%s,", "^%s,", ",%s$"},
	"u":   {"^u\\(\\w+(.+?)\\)$", ",u\\(\\w+(.+?)\\),", "^u\\(\\w+(.+?)\\),", ",u\\(\\w+(.+?)\\)$", "^%s$", ",%s,", "^%s,", ",%s$"},
	"r":   {"^r\\(\\w+(.+?)\\)$", ",r\\(\\w+(.+?)\\),", "^r\\(\\w+(.+?)\\),", ",r\\(\\w+(.+?)\\)$", "^%s$", ",%s,", "^%s,", ",%s$"},
	"l":   {"^l\\(\\w+(.+?)\\)$", ",l\\(\\w+(.+?)\\),", "^l\\(\\w+(.+?)\\),", ",l\\(\\w+(.+?)\\)$", "^%s$", ",%s,", "^%s,", ",%s$"},
	"q":   {"^q\\(\\w+(.+?)\\)$", ",q\\(\\w+(.+?)\\),", "^q\\(\\w+(.+?)\\),", ",q\\(\\w+(.+?)\\)$", "^%s$", ",%s,", "^%s,", ",%s$"},
	"*":   {"^%s$", ",%s,", "^%s,", ",%s$"},
}
var mysqlIsNull = map[string]string{
	"否":   "not null",
	"N":   "not null",
	"NO":  "not null",
	"是":   "",
	"":    "",
	"Y":   "",
	"YES": "",
}
