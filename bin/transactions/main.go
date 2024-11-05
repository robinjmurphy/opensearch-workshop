// package main generates a set of fake transaction data
package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"time"
)

type transaction struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle,omitempty"`
	Category  string    `json:"category"`
	Amount    int       `json:"amount"`
	Currency  string    `json:"currency"`
	AccountID string    `json:"account_id"`
}

type template struct {
	weight     int
	title      string
	subtitles  []string
	minAmount  int
	maxAmount  int
	categories []string
}

var templates = []template{
	{6, "Tesco", []string{}, 0001, 12500, []string{"groceries"}},
	{4, "Sainsbury's", []string{"", "Shopping", "Dinner 🍽️"}, 0001, 15000, []string{"groceries"}},
	{4, "TfL", []string{}, 0001, 1250, []string{"travel"}},
	{2, "Amazon", []string{}, 1000, 10000, []string{"shopping"}},
	{1, "Uber", []string{}, 750, 1850, []string{"transport"}},
	{1, "Black Sheep Coffee", []string{}, 400, 1200, []string{"eating_out"}},
	{1, "Apple", []string{}, 99, 345, []string{"shopping"}},
	{1, "Spotify", []string{}, 999, 1399, []string{"bills"}},
	{1, "Netflix", []string{}, 999, 1399, []string{"bills"}},
	{1, "Apple", []string{}, 99, 345, []string{"shopping"}},
	{1, "Humble Crumble", []string{}, 400, 500, []string{"eating_out"}},
	{1, "AirBnb", []string{}, 10000, 25000, []string{"holidays"}},
	{1, "Deliveroo", []string{"", "McDonalds 🍔", "Thanks for lunch 🙏🏻"}, 795, 2350, []string{"eating_out"}},
	{1, "PayPal", []string{"F029479", "Netflix", "S1-2358", ""}, 999, 1299, []string{"bills"}},
	{1, "Monzo", []string{}, 100000, 200000, []string{"income"}},

	{1, "Homer Simpson", []string{"Donuts"}, 100, 700, []string{"eating_out"}},
	{1, "George Bluth", []string{"Netflix account"}, 300, 800, []string{"general"}},
	{1, "Philomena Cunk", []string{"Money for books"}, 4000, 8000, []string{"general"}},
}

const (
	numDaysToGenerate     = 90
	minTransactionsPerDay = 0
	maxTransactionsPerDay = 6
)

func main() {
	var transactions []*transaction

	var weightedTemplates []template
	for _, t := range templates {
		for i := 0; i < t.weight; i++ {
			weightedTemplates = append(weightedTemplates, t)
		}
	}

	accountIDs := []string{
		fmt.Sprintf("acc_%s", uuid.New()),
		fmt.Sprintf("acc_%s", uuid.New()),
	}

	for i := 0; i < numDaysToGenerate; i++ {
		date := time.Now().UTC().AddDate(0, 0, i*-1).Truncate(time.Hour * 24)
		numTransactionsToGenerate := rand.Intn(maxTransactionsPerDay-minTransactionsPerDay) + minTransactionsPerDay

		for j := 0; j < numTransactionsToGenerate; j++ {
			t := weightedTemplates[rand.Intn(len(weightedTemplates))]

			subtitle := ""
			if len(t.subtitles) > 0 {
				subtitle = t.subtitles[rand.Intn(len(t.subtitles))]
			}

			transactions = append(transactions, &transaction{
				ID:        fmt.Sprintf("tx_%s", uuid.New()),
				Title:     t.title,
				Subtitle:  subtitle,
				Category:  t.categories[rand.Intn(len(t.categories))],
				Amount:    rand.Intn(t.maxAmount-t.minAmount) + t.minAmount,
				Currency:  "GBP",
				Timestamp: date.Add(time.Duration(rand.Intn(24)) * time.Hour),
				AccountID: accountIDs[rand.Intn(len(accountIDs))],
			})
		}
	}

	out, err := json.Marshal(transactions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
