package adapters

import (
	"math/rand"
	"time"
)

func NewRandomIntPicker() RandomIntPicker {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return RandomIntPicker{r}
}

type RandomIntPicker struct {
	rand *rand.Rand
}

func (r RandomIntPicker) PickRandomInt(upto int) int {
	return r.rand.Intn(upto)
}
