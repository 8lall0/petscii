package petscii

import (
	"errors"
	"unicode"
)

/*
	If the PETSCII character is A-Z, make it a-z (PETSCII 97-122, subtract 32)
	If the PETSCII character is a-z, make it A-Z (PETSCII 65-90, add 32)
	If the PETSCII character is 192-223, subtract 96. Then subtract 32 if the resultant value is 97-122.
*/

func AsciiToPetscii(str string) ([]byte, error) {
	s := []rune(str)
	output := make([]byte, 0)

	for _, c := range s {
		if c > unicode.MaxASCII {
			return nil, errors.New("invalid character in ASCII")
		}

		if c >= 'A' && c <= 'Z' {
			c += 32
		} else if c >= 'a' && c <= 'z' {
			c -= 32
		}

		output = append(output, byte(c))
	}

	return output, nil
}

func PetsciiToAscii(str []byte) (string, error) {
	for i, b := range str {
		if b >= 192 && b <= 223 {
			b -= 96
		}

		if b >= 65 && b <= 90 {
			b += 32
		} else if b >= 97 && b <= 122 {
			b -= 32
		}

		str[i] = b
	}

	return string(str), nil
}
