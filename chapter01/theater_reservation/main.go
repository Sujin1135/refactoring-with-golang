package main

import (
	"errors"
	"fmt"
	"math"
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

type Performances []Performance

type Invoice struct {
	Customer     string
	Performances *Performances
}

func statement(invoice *Invoice) (string, error) {
	totalAmount, volumeCredit := 0, float64(0)
	paymentStr := fmt.Sprintf("청구 내역 (고객명: %s)\n", invoice.Customer)

	for _, item := range *invoice.Performances {
		thisAmount := 0

		switch item.Play.Type {
		case Tragedy:
			thisAmount = 40000
			if item.Audience > 30 {
				thisAmount += 1000 * (item.Audience - 30)
			}
		case Comedy:
			thisAmount = 30000
			if item.Audience > 20 {
				thisAmount += 10000 + 500*(item.Audience-20)
			}
			thisAmount += 300 * item.Audience
		default:
			return "", errors.New(fmt.Sprintf("Not found this play type:%s\n", item.Play.Type))
		}

		volumeCredit += math.Max(float64(item.Audience-30), 0)

		if Comedy == item.Play.Type {
			volumeCredit += math.Floor(float64(item.Audience / 5))
		}

		paymentStr += fmt.Sprintf("%s: %d (%d석)\n", item.Play.Name, thisAmount, item.Audience)
		totalAmount += thisAmount
	}

	paymentStr += fmt.Sprintf("총액: %d\n", totalAmount)
	paymentStr += fmt.Sprintf("적립 포인트: %f점\n", volumeCredit)

	return paymentStr, nil
}

func main() {
	otti := &Play{Name: "Otti", Type: Comedy}
	performs := &Performances{
		Performance{Play: otti, Audience: 50},
	}

	result, err := statement(&Invoice{Customer: "Chris", Performances: performs})
	if err != nil {
		fmt.Errorf("occurred an error when get a payment's statement")
	}

	fmt.Println(result)
}
