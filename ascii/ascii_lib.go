package ascii

import "fmt"

func isWhitespace(char byte) bool {
	return char == '\n' || char == ' ' || char == '\t'
}

func TrimWhitespaces(data []byte) []byte {
	var i, j int
	for i = len(data) - 1; isWhitespace(data[i]); i-- {
	}
	for j = 0; isWhitespace(data[j]); j++ {
	}
	return data[j : i+1]
}

func IsDigit(data byte) bool {
	return data >= '0' && data <= '9'
}

func ToDigit(data byte) (digit int, err error) {
	if !IsDigit(data) {
		err = fmt.Errorf("ascii_lib: invalid digit")
		return
	}
	digit = int(data - '0')
	return
}

func ParseInt(data []byte) (int, error) {
    var result int
	for i, baseExp := len(data)-1, 1; i >= 0; i, baseExp = i-1, baseExp*10 {
		var value int
        value, err := ToDigit(data[i])
		if err != nil {
			return result, err
		}
		result += int(value) * baseExp
	}
    return result, nil
}
