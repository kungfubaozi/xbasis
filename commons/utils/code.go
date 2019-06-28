package xbasisutils

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandNumber() string {
	return fmt.Sprintf("%06v", rnd().Int31n(1000000))
}

func rnd() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
