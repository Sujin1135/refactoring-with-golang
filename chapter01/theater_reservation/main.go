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

func Statement(invoice *Invoice) (string, error) {
	var result strings.Builder

	totalAmount, volumeCredit := 0, float64(0)
	result.WriteString(fmt.Sprintf("청구 내역 (고객명: %s)\n", invoice.Customer))

	for _, item := range *invoice.Performances {
		volumeCredit += volumeCreditsFor(item)
		result.WriteString(fmt.Sprintf("%s: %d (%d석)\n", item.Play.Name, amountFor(item), item.Audience))
		totalAmount += amountFor(item)
	}

	result.WriteString(fmt.Sprintf("총액: %d\n", totalAmount))
	result.WriteString(fmt.Sprintf("적립 포인트: %f점\n", volumeCredit))

	return result.String(), nil
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
