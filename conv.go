package wolf

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (l *Lexicon) Unwolf(wolf string) (string, error) {
	if !IsValidWolf(wolf) {
		return wolf, errors.New("not a wolf name")
	}

	var output string

	for _, fragment := range strings.Split(wolf, "-") {
		if value, ok := l.index[fragment]; ok {
			output += fmt.Sprintf("%04x", value)
		} else {
			return fragment, errors.New("fragment not found in lexicon")
		}
	}

	return strings.TrimLeft(output, "0"), nil
}

func (l *Lexicon) Wolf(hex string) (string, error) {
	if !IsValidHex(hex) {
		return hex, errors.New("not a hex string")
	}

	var output []string

	for len(hex) > 0 {
		boundary := len(hex) - 4

		if boundary < 0 {
			boundary = 0
		}

		if parsed, err := strconv.ParseUint(hex[boundary:], 16, 16); err == nil {
			output = append([]string{l.words[parsed]}, output...)
		} else {
			return hex, err
		}

		hex = hex[0:boundary]
	}

	return strings.Join(output, "-"), nil
}
