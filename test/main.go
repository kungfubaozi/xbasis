package main

import (
	"fmt"
	"konekko.me/gosion/commons/constants"
)

func main() {
	fmt.Println(gs_commons_constants.AuthTypeOfToken, gs_commons_constants.AuthTypeOfFace,
		gs_commons_constants.AuthTypeOfMobileConfirm, gs_commons_constants.AuthTypeOfValcode,
		gs_commons_constants.AuthTypeOfMiniProgramCodeConfirm)
}
