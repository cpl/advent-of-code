package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fp, err := os.Open("year2019/day1/input.txt")
	checkErr(err)
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	var totalFuelReq int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		checkErr(err)
		totalFuelReq += calculateFuelReq(float64(num))
	}

	fmt.Printf("result: %d\n", totalFuelReq)
}

// 4882337
func calculateFuelReq(mass float64) int {
	mass = mass / 3
	mass = math.Floor(mass)
	fuel := int(mass) - 2

	if fuel <= 0 {
		return 0
	}

	fuel += calculateFuelReq(float64(fuel))

	return fuel
}
