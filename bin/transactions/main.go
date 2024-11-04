// package main generates a set of fake transaction data
package main

import "fmt"

type template struct {
	weight     int
	title      string
	subtitles  []string
	min        int
	max        int
	categories []string
}

var templates = []template{
	{6, "Tesco", []string{}, -0001, -12500, []string{"groceries"}},
	{4, "Sainsbury's", []string{"Shopping", "Dinner 🍽️"}, -0001, -15000, []string{"groceries"}},
	{4, "TfL", []string{}, -0001, -1250, []string{"travel"}},
	{2, "Uber", []string{}, -750, -1850, []string{"transport"}},
	{2, "Amazon", []string{}, -1000, -10000, []string{"shopping"}},
	{1, "Apple", []string{}, -99, -345, []string{"shopping"}},
	{1, "AirBnb", []string{}, -10000, -25000, []string{"holidays"}},
	{1, "Deliveroo", []string{"Thanks for lunch 🙏🏻"}, -795, -2350, []string{"eating_out"}},
	{1, "PayPal", []string{"F029479", "Netflix", "S1-2358"}, -999, -1299, []string{"bills"}},
	{1, "Monzo", []string{}, 100000, 200000, []string{"income"}},

	// Add some people
	// Add some more interesting text for merchants/references with overlaps
}

func main() {
	fmt.Println("Hello")
}
