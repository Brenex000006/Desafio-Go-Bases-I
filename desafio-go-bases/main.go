package main

import (
	"fmt"

	"github.com/Brenex000006/Desafio-Go-Bases-I/tree/main/desafio-go-bases/internal/tickets"
)

func main() {

	TotalTickets, _ := tickets.GetTotalTickets("China")
	fmt.Println(TotalTickets)

	CountByPeriod, _ := tickets.GetCountByPeriod("Night")
	fmt.Println(CountByPeriod)

	MedianDestination, _ := tickets.AverageDestination()
	fmt.Println(MedianDestination)
}
