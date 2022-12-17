package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	id      int
	Destino string
	Horario string
	Nome    string
	Email   string
	Valor   float64
}

func GetTotalTickets(destination string) (int, error) {
	v, _ := os.Open("tickets.csv")
	q := csv.NewReader(v)

	calculation := 0

	for {
		record, err := q.Read()
		if err == io.EOF {
			break
		}

		id, _ := strconv.Atoi(record[0])
		Valor, _ := strconv.ParseFloat(record[1], 8)
		ticket := Ticket{
			id:      id,
			Valor:   Valor,
			Destino: record[2],
			Horario: record[3],
			Nome:    record[4],
			Email:   record[5],
		}

		if ticket.Destino == destination {
			calculation++
		}
	}
	memorandum := fmt.Sprintf("Total trips to country %s: %d", destination, calculation)
	return memorandum, nil
}

var ErrorBreak = errors.New("---> Please enter a valid range. Between 00:00 to 23:00 <---")

func GetCountByPeriod(time string) (string, error) {
	v, _ := os.Open("tickets.csv")
	q := csv.NewReader(v)

	calculation := 0

	for {
		record, err := q.Read()
		if err == io.EOF {
			break
		}

		id, _ := strconv.Atoi(record[0])
		Valor, _ := strconv.ParseFloat(record[1], 8)
		ticket := Ticket{
			id:      id,
			Valor:   Valor,
			Destino: record[2],
			Horario: record[3],
			Nome:    record[4],
			Email:   record[5],
		}

		period, err := strconv.Atoi(strings.Split(ticket.Horario, ":")[0])
		if err != nil {
			panic(err)
		}

		switch time {
		case "Night":
			if period >= 20 && period <= 23 {
				calculation++
			}
		case "Evening":
			if period >= 13 && period <= 19 {
				calculation++
			}
		case "Morning":
			if period >= 07 && period <= 12 {
				calculation++
			}
		case "Daybreak":
			if period >= 00 && period <= 06 {
				calculation++
			}
		}
	}

	memorandum := fmt.Sprintf("Total number of people traveling during the period of %s: %d", time, calculation)
	return memorandum, nil
}

func AverageDestination(destination string, total int) (int, error) {

	countries := make(map[string]int)

	v, _ := os.Open("tickets.csv")
	q := csv.NewReader(v)

	calculation := 0

	for {
		record, err := q.Read()
		if err == io.EOF {
			break
		}

		id, _ := strconv.Atoi(record[0])
		Valor, _ := strconv.ParseFloat(record[1], 8)
		ticket := Ticket{
			id:      id,
			Valor:   Valor,
			Destino: record[2],
			Horario: record[3],
			Nome:    record[4],
			Email:   record[5],
		}

		if countries[ticket.Destino] == 0 {
			countries[ticket.Destino]++
		} else {
			countries[ticket.Destino]++
		}
		calculation++
	}

	Average := calculation / len(countries)

	memorandum := fmt.Sprintf("Average trips: %d", Average)

	return memorandum, nil
}
