package main

import (
	"fmt"
	"math"
	"strings"
)

type PlayType string

const (
	Tragedy = PlayType("T")
	Comedy  = PlayType("C")
)

type Play struct {
	Name string
	Type PlayType
}

type Performance struct {
	Play     *Play
	Audience int
}

type Performances []*Performance

type Invoice struct {
	Customer     string
	Performances *Performances
}

type statementData struct {
	Amount        int
	VolumeCredits float64
	Invoice       *Invoice
}

func newStatementData(invoice *Invoice) *statementData {
	return &statementData{
		Amount:        totalAmount(*invoice.Performances),
		VolumeCredits: totalVolumeCredits(*invoice.Performances),
		Invoice:       invoice,
	}
}

func Statement(invoice *Invoice) (string, error) {
	statementData := newStatementData(invoice)
	return renderText(statementData)
}

func renderText(data *statementData) (string, error) {
	var result strings.Builder

	result.WriteString(fmt.Sprintf("청구 내역 (고객명: %s)\n", data.Invoice.Customer))

	for _, item := range *data.Invoice.Performances {
		result.WriteString(fmt.Sprintf("%s: %d (%d석)\n", item.Play.Name, amountFor(item), item.Audience))
	}

	result.WriteString(fmt.Sprintf("총액: %d\n", data.Amount))
	result.WriteString(fmt.Sprintf("적립 포인트: %f점\n", data.VolumeCredits))

	return result.String(), nil
}

func totalVolumeCredits(aPerformances Performances) float64 {
	volumeCredits := float64(0)

	for _, item := range aPerformances {
		volumeCredits += volumeCreditsFor(item)
	}
	return volumeCredits
}

func totalAmount(aPerformance Performances) int {
	total := 0

	for _, item := range aPerformance {
		total += amountFor(item)
	}
	return total
}

func volumeCreditsFor(aPerformance *Performance) float64 {
	volumeCredit := float64(0)
	volumeCredit += math.Max(float64(aPerformance.Audience-30), 0)

	if Comedy == aPerformance.Play.Type {
		volumeCredit += math.Floor(float64(aPerformance.Audience / 5))
	}
	return volumeCredit
}

func amountFor(item *Performance) int {
	result := 0
	audience := item.Audience

	switch item.Play.Type {
	case Tragedy:
		result = amountTragedy(audience)
	case Comedy:
		result = amountComedy(audience)
	}
	return result
}

func amountComedy(audience int) int {
	result := 30000

	if audience > 20 {
		result += 10000 + 500*(audience-20)
	}
	result += 300 * audience
	return result
}

func amountTragedy(audience int) int {
	result := 40000
	if audience > 30 {
		result += 1000 * (audience - 30)
	}
	return result
}

func main() {
	hamlet := &Play{Name: "Hamlet", Type: Comedy}
	romeo := &Play{Name: "Romeo and Juliet", Type: Tragedy}
	performs := &Performances{
		&Performance{Play: hamlet, Audience: 50},
		&Performance{Play: romeo, Audience: 20},
	}

	result, err := Statement(&Invoice{Customer: "Chris", Performances: performs})
	if err != nil {
		fmt.Errorf("occurred an error when get a payment's statement")
	}

	fmt.Println(result)
}
