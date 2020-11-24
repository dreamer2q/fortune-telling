package fortune_telling

import (
	"errors"
	"math/rand"
)

type Telling struct {
	Level   string
	Content string
	Detail1 string
	Detail2 string
}

func Ask(key string) (Telling, error) {
	defer func() { signedMap[key]++ }()
	if !HasAsked(key) {
		return genTelling(), nil
	}
	return Telling{}, errors.New("already asked for a fortune-telling")
}

func HasAsked(key string) bool {
	return signedMap[key] > 0
}

func genTelling() Telling {
	var index = rand.Intn(len(signs))
	return signs[index]
}

func (t Telling) String(Ji string) string {
	var count = levelMap[Ji]
	var ret string
	for i := 0; i < 6; i++ {
		var cell string
		if i < count {
			cell = "★"
		} else {
			cell = "☆"
		}
		ret += cell
	}
	return ret
}
