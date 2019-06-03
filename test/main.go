package main

import "fmt"

func main() {
	//fmt.Println(gs_commons_constants.AuthTypeOfToken, gs_commons_constants.AuthTypeOfFace,
	//	gs_commons_constants.AuthTypeOfMobileConfirm, gs_commons_constants.AuthTypeOfValcode,
	//	gs_commons_constants.AuthTypeOfMiniProgramCodeConfirm)
	a := []int64{2, 3, 4, 6, 9}
	size := len(a)
	s := size / 2
	fmt.Println(s, size-s)
	fmt.Println(a[:s], a[s:])
}
