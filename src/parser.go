package src

import "strconv"

func ParseInt(str string) int {
	integer, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return integer
}

func ParseStr(integer int64) string {
	return strconv.FormatInt(integer, 10)
}
