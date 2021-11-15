package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"../../utils"
)

type bagInfo struct {
	name  string
	count int
}

var info = make(map[string][]*bagInfo)

func main() {
	data, err := utils.GetInput(2020, 07)
	utils.CheckErr(err)

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.SplitN(line[:len(line)-1], " bags contain ", 2)
		key := split[0]
		val := strings.Split(split[1], ", ")

		if val[0] == "no other bags" {
			info[key] = nil
			continue
		}

		for _, v := range val {

			idx := strings.Index(v, " ")
			count, _ := strconv.Atoi(v[:idx])

			info[key] = append(info[key], &bagInfo{
				name:  v[idx+1 : strings.LastIndex(v, " ")],
				count: count,
			})
		}
	}

	total := 0
	for key := range info {
		if recExSearch("shiny gold", key) {
			total++
		}
	}
	fmt.Println()
	fmt.Println("total", total)

	fmt.Println("bag count", bagCount("shiny gold", 1)-1)
}

func bagCount(target string, level int) int {
	i := info[target]

	total := 1
	for _, v := range i {
		fmt.Println(strings.Repeat("--", level), target, "contains", v.count, v.name)

		add := v.count * bagCount(v.name, level+1)
		total += add + 0
	}

	return total
}

func recExSearch(target string, bag string) bool {
	i := info[bag]
	for _, v := range i {
		if v.name == target {
			return true
		}
		if v.name == "no other bags" {
			return false
		}

		nested := recExSearch(target, v.name)
		if nested {
			return true
		}
	}
	return false
}
