package common

import "fmt"

type OperatorPriority map[string]int

var Opt OperatorPriority

// -1 <
// 0 空
// 1 >
// 2 =
func init() {
	Opt = OperatorPriority{
		"++": 1, "+-": 1, "+*": -1, "+/": -1, "+(": -1, "+)": 1,"+#": 1,
		"-+": 1, "--": 1, "-*": -1, "-/": -1, "-(": -1, "-)": 1,"-#": 1,
		"*+": 1, "*-": 1, "**": 1, "*/": 1, "*(": -1, "*)": 1,"*#": 1,
		"/+": 1, "/-": 1, "/*": 1, "//": 1, "/(": -1, "/)": 1,"/#": 1,
		"(+": -1, "(-": -1, "(*": -1, "(/": -1, "((": -1, "()": 2,"(#": 0,
		")+": 1, ")-": 1, ")*": 1, ")/": 1, ")(": 0, "))": 1,")#": 1,
		"#+": -1, "#-": -1, "#*": -1, "#/": -1, "#(": -1, "#)": 0,"##": 2,
	}
}

func CompareS(s string) (int, bool) {
	if v, ok := Opt[s]; ok {
		return v, true
	} else {
		return 0, false
	}
}

func CompareR(r1, r2 rune) (int, bool) {
	s := fmt.Sprintf("%c%c", r1, r2)
	return CompareS(s)
}
