package tickets

import (
	"encoding/csv"
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID      int
	Nombre  string
	Email   string
	Destino string
	Hora    string
	Precio  int
}

func ObtenerTicketsTotales() ([]Ticket, error) {
	var ticketsArr []Ticket
	fd, error := os.Open("tickets.csv")
	if error != nil {
		return []Ticket{}, error
	}
	defer fd.Close()

	// read CSV file
	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()
	if error != nil {
		return []Ticket{}, error
	}

	for _, dato := range records {
		Id, err := strconv.Atoi(dato[0])
		if err != nil {
			return []Ticket{}, errors.New("Error al parsear ID")
		}
		Precio, err := strconv.Atoi(dato[5])
		if err != nil {
			return []Ticket{}, errors.New("Error al parsear precio")
		}
		ticket := Ticket{
			ID:      Id,
			Nombre:  dato[1],
			Email:   dato[2],
			Destino: dato[3],
			Hora:    dato[4],
			Precio:  Precio,
		}
		ticketsArr = append(ticketsArr, ticket)
	}
	return ticketsArr, nil
}

var ticketsTotales, er = ObtenerTicketsTotales()

// ejemplo 1
func GetTotalTickets(destination string) (float64, error) {
	var count float64
	for _, dato := range ticketsTotales {
		if dato.Destino == destination {
			count++
		}
	}
	if count == 0 {
		return 0, errors.New("Destino sin pasajeros")
	}
	return count, nil
}

// ejemplo 2
func GetCountByPeriod(time string) (int, error) {
	var cantidadViajes int
	switch time {
	case "madrugada":
		cantidadViajes = calcularHorario(0, 6)
	case "manana":
		cantidadViajes = calcularHorario(7, 12)
	case "tarde":
		cantidadViajes = calcularHorario(13, 19)
	case "noche":
		cantidadViajes = calcularHorario(20, 23)
	default:
		return 0, errors.New("Tiempo no definido")
	}
	return cantidadViajes, nil
}

func calcularHorario(comienzo, fin int) int {
	count := 0
	for _, t := range ticketsTotales {
		horaString := strings.Split(t.Hora, ":")
		hora, err := strconv.Atoi(horaString[0])
		if err != nil {
			return 0
		}
		if hora >= comienzo && hora <= fin {
			count++
		}
	}
	return count
}

// ejemplo 3
func PercentagePerDayDestination(destination string) (float64, error) {

	totalTicketsDestination, err := GetTotalTickets(destination)
	if err != nil {
		return 0, errors.New("No hay viajes a ese destino")
	}
	totalViajes := float64(len(ticketsTotales))

	porcentajeViajes := (totalTicketsDestination / totalViajes) * 100
	// fmt.Println(porcentajeViajes)
	// fmt.Printf("Total a destinacion: %.2f\nTotal viajes: %.2f\nPorcentaje: %.2f", totalTicketsDestination, totalViajes, porcentajeViajes)

	return roundFloat(porcentajeViajes, 2), nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
