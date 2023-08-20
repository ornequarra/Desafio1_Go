package main

import (
	"fmt"
	"log"
	"os"

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

	canalTicket := make(chan tickets.Ticket)
	defer close(canalTicket)

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
		totalTickets, err := storage.GetTotalTickets(entrada)
		if err != nil {
			canalErr <- err
			return
		}
		canalTotalTickets <- totalTickets
	}(canalTotalTickets, canalErr)

	//Impresion de Canales
	select {
	case totalTicket := <-canalTotalTickets:
		fmt.Println("El total de tickets para el pais %s es: %d", entrada, totalTicket)
	case ticket := <-canalTicket:
		fmt.Println(ticket)
	case err := <-canalErr:
		fmt.Println(err)
		os.Exit(1)
	}
}
