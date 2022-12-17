package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	TotalTickets, _ := tickets.GetTotalTickets("China")
	fmt.Println(TotalTickets)

	CountByPeriod, _ := tickets.GetCountByPeriod("Night")
	fmt.Println(CountByPeriod)

	MedianDestination, _ := tickets.AverageDestination()
	fmt.Println(MedianDestination)
}
