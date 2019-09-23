package kata

import "math"

func Movie(card, ticket int, perc float64) int {
	totalSpentA := float64(0)
	totalSpentB := float64(card)
	ticketA := float64(ticket)
	ticketB := float64(ticket)
	tickets := 0

	for totalSpentA <= math.Ceil(totalSpentB) {
		tickets++
		totalSpentA += ticketA
		ticketB = ticketB * perc
		totalSpentB = totalSpentB + ticketB
	}
	return tickets

}
