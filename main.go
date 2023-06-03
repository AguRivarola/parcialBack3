package main

import (
	"fmt"
	"sync"

	"github.com/AguRivarola/parcialBack3/internal/tickets"
)

func main() {
	wg := new(sync.WaitGroup)
	channel := make(chan string)

	for x := 1; x <= 4; x++ {
		wg.Add(3)
		go testPercentagePerDayDestination("Colombia", wg, channel)
		go testGetCountByPeriod("tarde", wg, channel)
		go testGetTotalTickets("Argentina", wg, channel)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	for msg := range channel {
		fmt.Println(msg)
	}
}

func testPercentagePerDayDestination(destino string, wg *sync.WaitGroup, c chan string) {
	defer wg.Done()
	porcentajeViajes, err := tickets.PercentagePerDayDestination(destino)
	if err != nil {
		c <- err.Error()
		close(c)
	}

	c <- fmt.Sprintf("\nResultado de PercentagePerDayDestination: \t%.2f", porcentajeViajes)

}
func testGetCountByPeriod(horario string, wg *sync.WaitGroup, c chan string) {
	defer wg.Done()
	cantidadPorPeriodo, err := tickets.GetCountByPeriod(horario)
	if err != nil {
		c <- err.Error()
		close(c)
	}

	c <- fmt.Sprintf("\nResultado de GetCountByPeriod\t%d", cantidadPorPeriodo)
}
func testGetTotalTickets(destino string, wg *sync.WaitGroup, c chan string) {
	defer wg.Done()
	total, err := tickets.GetTotalTickets(destino)
	if err != nil {
		c <- err.Error()
		close(c)
	}

	c <- fmt.Sprintf("\nResultado de GetTotalTickets\t%.2f", total)

}
