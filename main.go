package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ornequarra/Desafio1_Go/internal/tickets"
)

const (
	filename = "./tickets.csv"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	storage := tickets.Storage{
		Tickets: tickets.ReadFile(filename),
	}

	//Crear Canales para comunicarnos con las GoRoutines
	canalTotalTickets := make(chan int)
	defer close(canalTotalTickets)

	canalViajantesPorHorario := make(chan string)
	defer close(canalViajantesPorHorario)

	canalErr := make(chan error)
	defer close(canalErr)

	//Requerimiento 1: Contar total de tickets
	var entrada string

	fmt.Print("Ingrese pais a buscar :")
	_, err := fmt.Scan(&entrada)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	go func(chan int, chan error) {
		totalTickets, err := storage.GetTotalTickets(entrada, storage.Tickets)
		if err != nil {
			canalErr <- err
			return
		}
		canalTotalTickets <- totalTickets
	}(canalTotalTickets, canalErr)

	time.Sleep(time.Millisecond * 100)

	//Requerimiento 2: Contar total de viajantes por rango horario
	var entradaRangoHorario string

	fmt.Print("Ingrese rango horario a buscar :")
	_, err2 := fmt.Scan(&entradaRangoHorario)

	if err2 != nil {
		log.Fatal(err2)
		os.Exit(1)
	}

	go func(chan string, chan error) {
		totalTickets, err := storage.GetCountByPeriod(entradaRangoHorario, storage.Tickets)
		if err != nil {
			canalErr <- err
			return
		}
		mensaje := fmt.Sprintf("La cantidad de viajantes en el rango %s es %d\n.", entradaRangoHorario, totalTickets)
		canalViajantesPorHorario <- mensaje
	}(canalViajantesPorHorario, canalErr)

	time.Sleep(time.Millisecond * 100)

	//Impresion de Canales
	select {
	case totalTicket := <-canalTotalTickets:
		fmt.Printf("El total de tickets para el país %s es: %d\n", entrada, totalTicket)
	case ticketPorHorario := <-canalViajantesPorHorario:
		fmt.Println(ticketPorHorario)
	case err := <-canalErr:
		fmt.Println(err)
		os.Exit(1)
	}
}
