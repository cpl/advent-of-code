package aoc

import (
	"fmt"
	"io/ioutil"
)

func MetaGetInput(year, day int) ([]byte, error) {
	filename := fmt.Sprintf("./day%02d.txt", day)
	data, err := ioutil.ReadFile(filename)
	if err == nil {
		return data, nil
	}

	return data, nil
}
