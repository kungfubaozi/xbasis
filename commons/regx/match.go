package gs_commons_regx

import (
	"regexp"
)

func Phone(str string) bool {
	reg := `^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(str)
}

func Email(str string) bool {
	reg := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(str)
}
