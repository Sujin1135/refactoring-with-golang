package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStatement(t *testing.T) {
	var sb strings.Builder
	customerName := "Chris"
	hamlet := &Play{Name: "Hamlet", Type: Comedy}
	romeo := &Play{Name: "Romeo and Juliet", Type: Tragedy}
	hamletAudience := 50
	romeoAudience := 20
	invoice := newInvoice(hamlet, hamletAudience, romeo, romeoAudience)

	sb.WriteString(fmt.Sprintf("청구 내역 (고객명: %s)\n", customerName))
	sb.WriteString(fmt.Sprintf("%s: %d (%d석)\n", hamlet.Name, 70000, hamletAudience))
	sb.WriteString(fmt.Sprintf("%s: %d (%d석)\n", romeo.Name, 40000, romeoAudience))
	sb.WriteString(fmt.Sprintf("총액: %d\n", 110000))
	sb.WriteString(fmt.Sprintf("적립 포인트: %f점\n", float64(30)))

	expected := sb.String()
	sut, _ := Statement(invoice)

	assert.Equal(t, expected, sut)
}

func newInvoice(hamlet *Play, hamletAudience int, romeo *Play, romeoAudience int) *Invoice {
	performs := &Performances{
		&Performance{Play: hamlet, Audience: hamletAudience},
		&Performance{Play: romeo, Audience: romeoAudience},
	}
	invoice := &Invoice{Customer: "Chris", Performances: performs}
	return invoice
}
