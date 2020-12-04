package main

// i think this is all wrong, i think i need a passport package or something
import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         int
	HeightUnit     string
	HairColor      string
	EyeColor       string
	PassportId     string
}

func New(raw_passport string) (Passport, error) {
	var passport Passport
	raw_passport = strings.ReplaceAll(raw_passport, "\r\n", " ") + " "

	birth_year, err1 := ParseBirthYear(raw_passport)
	issue_year, err2 := ParseIssueYear(raw_passport)
	exp_year, err3 := ParseExpirationYear(raw_passport)
	height, height_unit, err4 := ParseHeight(raw_passport)
	hair_color, err5 := ParseHairColor(raw_passport)
	eye_color, err6 := ParseEyeColor(raw_passport)
	passport_id, err7 := ParsePassportId(raw_passport)

	if err1 != nil ||
		err2 != nil ||
		err3 != nil ||
		err4 != nil ||
		err5 != nil ||
		err6 != nil ||
		err7 != nil {
		return passport, errors.New("Passport invalid")
	}

	return Passport{
		BirthYear:      birth_year,
		IssueYear:      issue_year,
		ExpirationYear: exp_year,
		Height:         height,
		HeightUnit:     height_unit,
		HairColor:      hair_color,
		EyeColor:       eye_color,
		PassportId:     passport_id,
	}, nil
}

func ParseBirthYear(raw_passport string) (int, error) {
	birth_year_data := regexp.MustCompile("byr:(.*?) ").FindStringSubmatch(raw_passport)
	birth_year := 0

	if len(birth_year_data) > 1 {
		birth_year, _ = strconv.Atoi(birth_year_data[1])
	} else {
		return -1, errors.New("No birth year found")
	}

	if birth_year >= 1920 && birth_year <= 2002 {
		return birth_year, nil
	} else {
		return -1, errors.New("Birth year was out of range")
	}
}

func ParseIssueYear(raw_passport string) (int, error) {
	issue_year_data := regexp.MustCompile("iyr:(.*?) ").FindStringSubmatch(raw_passport)
	issue_year := 0

	if len(issue_year_data) > 1 {
		issue_year, _ = strconv.Atoi(issue_year_data[1])
	} else {
		return -1, errors.New("No issue year found")
	}

	if issue_year >= 2010 && issue_year <= 2020 {
		return issue_year, nil
	} else {
		return -1, errors.New("Issue year was out of range")
	}
}

func ParseExpirationYear(raw_passport string) (int, error) {
	exp_year_data := regexp.MustCompile("eyr:(.*?) ").FindStringSubmatch(raw_passport)
	exp_year := 0

	if len(exp_year_data) > 1 {
		exp_year, _ = strconv.Atoi(exp_year_data[1])
	} else {
		return -1, errors.New("No expiration year found")
	}

	if exp_year >= 2020 && exp_year <= 2030 {
		return exp_year, nil
	} else {
		return -1, errors.New("Expiration year was out of range")
	}
}

func ParseHeight(raw_passport string) (int, string, error) {
	height_data := regexp.MustCompile("hgt:(.*?)(cm|in) ").FindStringSubmatch(raw_passport)
	height := 0
	unit := ""

	if len(height_data) > 2 {
		height, _ = strconv.Atoi(height_data[1])
		unit = height_data[2]
	} else {
		return 0, "", errors.New("No height found")
	}

	if (unit == "in" && height >= 59 && height <= 76) || (unit == "cm" && height >= 150 && height <= 193) {
		return height, unit, nil
	} else {
		return 0, "", errors.New("height was out of range")
	}
}

func ParseHairColor(raw_passport string) (string, error) {
	hair_data := regexp.MustCompile("hcl:(#[0-9a-f]{6}) ").FindStringSubmatch(raw_passport)
	hair_color := ""

	if len(hair_data) > 1 {
		hair_color = hair_data[1]
	} else {
		return "", errors.New("No hair color found")
	}

	return hair_color, nil
}

func ParseEyeColor(raw_passport string) (string, error) {
	eye_data := regexp.MustCompile("ecl:(.*?) ").FindStringSubmatch(raw_passport)
	eye_color := ""
	valid_eye_colors := "amb,blu,brn,gry,grn,hzl,oth"

	if len(eye_data) > 1 {
		eye_color = eye_data[1]
	} else {
		return "", errors.New("No eye color found")
	}

	if strings.Contains(valid_eye_colors, eye_color) {
		return eye_color, nil
	} else {
		return "", errors.New("Eye color did not match valid list")
	}
}

func ParsePassportId(raw_passport string) (string, error) {
	passport_data := regexp.MustCompile("pid:([0-9]{9}) ").FindStringSubmatch(raw_passport)
	passport := ""

	if len(passport_data) > 1 {
		passport = passport_data[1]
	} else {
		return "", errors.New("No passport found")
	}

	return passport, nil
}
func IsValid(passport string) bool {
	// cid missing
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, required_field := range required {
		if !strings.Contains(passport, required_field+":") {
			return false
		}
	}

	return true
}
