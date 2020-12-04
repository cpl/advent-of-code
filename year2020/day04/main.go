package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"

	"../../utils"
)

const ()

type Passport struct {
	BirthYear      string // byr
	IssueYear      string // iyr
	ExpirationYear string // eyr
	Height         string // hgt
	HairColor      string // hcl
	EyeColor       string // ecl
	PassportID     string // pid
	CountryID      string // cid
}

func (p *Passport) isValid() bool {
	return p.BirthYear != "" &&
		p.IssueYear != "" &&
		p.ExpirationYear != "" &&
		p.Height != "" &&
		p.HairColor != "" &&
		p.EyeColor != "" &&
		p.PassportID != ""
}

func (p *Passport) populate(line string) {
	split := strings.Split(line, " ")
	for _, kv := range split {
		keyValue := strings.Split(kv, ":")
		if len(keyValue) != 2 {
			log.Println(line)
			continue
		}

		key := strings.TrimSpace(keyValue[0])
		val := strings.TrimSpace(keyValue[1])

		// all this can be avoided with a map[string]string where the key is the actual key
		// but who doesn't like a well typed language :)
		switch key {
		case "byr":
			year, err := strconv.Atoi(val)
			if err != nil || year > 2002 || year < 1920 {
				continue
			}
			p.BirthYear = val
		case "iyr":
			year, err := strconv.Atoi(val)
			if err != nil || year < 2010 || year > 2020 {
				continue
			}
			p.IssueYear = val
		case "eyr":
			year, err := strconv.Atoi(val)
			if err != nil || year < 2020 || year > 2030 {
				continue
			}
			p.ExpirationYear = val
		case "hgt":
			switch len(val) {
			case 4:
				if val[2:4] != "in" {
					continue
				}
				num, err := strconv.Atoi(val[:2])
				if err != nil || num < 59 || num > 76 {
					continue
				}
			case 5:
				if val[3:5] != "cm" {
					continue
				}
				num, err := strconv.Atoi(val[:3])
				if err != nil || num < 150 || num > 193 {
					continue
				}
			default:
				continue
			}
			p.Height = val
		case "hcl":
			if len(val) != 7 || val[0] != '#' {
				continue
			}
			color := make([]byte, 3)
			_, err := hex.Decode(color, []byte(val[1:]))
			if err != nil {
				continue
			}
			p.HairColor = val
		case "ecl":
			switch val {
			case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			default:
				continue
			}
			p.EyeColor = val
		case "pid":
			if len(val) != 9 {
				continue
			}
			_, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			p.PassportID = val
		case "cid":
			p.CountryID = val
		default:
			log.Println("unknown key", key)
		}
	}
}

func main() {
	data, err := utils.GetInput(2020, 04)
	utils.CheckErr(err)

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	scanner.Split(bufio.ScanLines)

	passports := make([]*Passport, 0)
	passport := new(Passport)

	countValid := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			fmt.Printf("%#v\n", passport)
			if passport.isValid() {
				countValid++
			}

			passports = append(passports, passport)
			passport = new(Passport)

			continue
		}

		passport.populate(line)
	}

	fmt.Printf("%#v\n", passport)
	if passport.isValid() {
		countValid++
	}
	passports = append(passports, passport)

	log.Println("valid passports", countValid, "out of", len(passports))
}
