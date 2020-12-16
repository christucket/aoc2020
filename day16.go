package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}
type Ticket struct {
	Values []int
	Valid  bool
}
type Field struct {
	Range1       Range
	Range2       Range
	IdentifiedAs int
	name         string
}

func IdentifyFieldWithValues(fields []Field, values []int, values_idx int) int {
	could_be := -1

	for i := range fields {
		field := fields[i]

		// if the field is identified as something already then dont check it
		if field.IdentifiedAs >= 0 {
			continue
		}

		good_values := true
		for _, value := range values {
			if !((value >= field.Range1.Start && value <= field.Range1.End) ||
				(value >= field.Range2.Start && value <= field.Range2.End)) {
				good_values = false
				break
			}
		}
		if good_values {
			if could_be > -1 {
				return -1
			}
			could_be = i
		}
	}

	return could_be
}

func GetRange(string_range []string) Range {
	start, _ := strconv.Atoi(string_range[1])
	end, _ := strconv.Atoi(string_range[2])

	return Range{start, end}
}

func day16() {
	inp, _ := ioutil.ReadFile("./inputs/day16.input")

	data := GetDoubleStringInput(inp)

	all_fields := []Field{}
	all_ranges := []Range{}
	all_tickets := []Ticket{}
	var my_ticket Ticket
	range_regexp := regexp.MustCompile("(\\d+)-(\\d+)")

	for _, l := range strings.Split(data[0], "\r\n") {
		rs := range_regexp.FindAllStringSubmatch(l, -1)
		name := strings.Split(l, ":")[0]

		range1 := GetRange(rs[0])
		range2 := GetRange(rs[1])
		all_ranges = append(all_ranges, range1)
		all_ranges = append(all_ranges, range2)
		all_fields = append(all_fields, Field{range1, range2, -1, name})
	}

	// MY ticket
	for _, l := range strings.Split(data[1], "\r\n") {
		if strings.Contains(l, "your") {
			continue
		}
		values := []int{}
		for _, is := range strings.Split(l, ",") {
			ii, _ := strconv.Atoi(is)
			values = append(values, ii)
		}
		my_ticket = Ticket{values, true}
	}

	// NEARBY tickets
	for _, l := range strings.Split(data[2], "\r\n") {
		if strings.Contains(l, "nearby") {
			continue
		}
		values := []int{}
		for _, is := range strings.Split(l, ",") {
			ii, _ := strconv.Atoi(is)
			values = append(values, ii)
		}
		all_tickets = append(all_tickets, Ticket{values, true})
	}

	bad_values := []int{}

	// part 1 check every value for every range
	// invalidate bad tickets
	for i := range all_tickets {
		ticket := &all_tickets[i]
		for _, ticket_value := range ticket.Values {
			bad := true
			for _, attr_range := range all_ranges {
				if ticket_value >= attr_range.Start && ticket_value <= attr_range.End {
					// fmt.Println(ticket_value, "is in ", attr_range.Start, "-", attr_range.End)
					bad = false
					break
				}
			}
			if bad {
				ticket.Valid = false
				bad_values = append(bad_values, ticket_value)
			}
		}
	}

	sum := 0
	for _, i := range bad_values {
		sum += i
	}

	fmt.Println(sum, bad_values)

	// part 2 try to identify the field with the values
	// try: loop through all fields with the set of values for that field
	//      if theres a unique field, set it to that one, else, fail

	found := make(map[int]bool)
	set_at_least_one := true
	tries := 0
	for set_at_least_one {
		set_at_least_one = false

		for i := range all_fields {
			if found[i] {
				continue
			}

			// need to transform the values into a different object based on every ticket
			// take value X from ticket 1, 2, .. n
			values := []int{}
			for _, ticket := range all_tickets {
				if ticket.Valid {
					values = append(values, ticket.Values[i])
				}
			}

			could_be := IdentifyFieldWithValues(all_fields, values, i)
			if could_be > -1 {
				all_fields[could_be].IdentifiedAs = i
				found[i] = true
				set_at_least_one = true
			}
		}

		// get all values for a particular field
		tries++
	}

	// part 2 asks for all the product of MY ticket's values identifiedas 0-5
	// guess 1: 1262754090593 too high
	//          1169409308827 too high
	//          i used logic ```if field.IdentifiedAs between 0-5, product```
	product := 1
	for i := range all_fields {
		field := all_fields[i]
		if strings.Contains(field.name, "departure") {
			product *= my_ticket.Values[field.IdentifiedAs]
			fmt.Println("my ticket's value for ", field.IdentifiedAs, "is", my_ticket.Values[field.IdentifiedAs])
			fmt.Println(field)
		}
	}
	fmt.Println("my ticket:", my_ticket)
	fmt.Println("product is:", product)
}
