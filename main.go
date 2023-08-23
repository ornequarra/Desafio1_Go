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
	canalTotalTickets := make(chan string)
	defer close(canalTotalTickets)

	canalViajantesPorHorario := make(chan string)
	defer close(canalViajantesPorHorario)

	canalPorcentajePorDestino := make(chan string)
	defer close(canalPorcentajePorDestino)

	canalErr := make(chan error)
	defer close(canalErr)

	//Requerimiento 1: Contar total de tickets
	var entrada string

	fmt.Print("Ingrese pais a buscar: ")
	_, err := fmt.Scan(&entrada)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	go func(chan string, chan error) {
		totalTickets, err := storage.GetTotalTickets(entrada, storage.Tickets)

		if err != nil {
			canalErr <- err
		}
		mensaje := fmt.Sprintf("El total de tickets para el pais %s es: %d.", entrada, totalTickets)
		canalTotalTickets <- mensaje
	}(canalTotalTickets, canalErr)

	//Requerimiento 2: Contar total de viajantes por rango horario
	var entradaRangoHorario string

	fmt.Print("Ingrese rango horario a buscar: ")
	_, err2 := fmt.Scan(&entradaRangoHorario)

	if err2 != nil {
		log.Fatal(err2)
		os.Exit(1)
	}

	go func(chan string, chan error) {
		totalTickets, err := storage.GetCountByPeriod(entradaRangoHorario, storage.Tickets)
		if err != nil {
			canalErr <- err
		}
		mensaje := fmt.Sprintf("La cantidad de viajantes en el rango %s es %d.", entradaRangoHorario, totalTickets)
		canalViajantesPorHorario <- mensaje
	}(canalViajantesPorHorario, canalErr)
	//Requerimiento 3: Contar total de viajantes por rango horario
	var entradaPorcentaje string

	fmt.Print("Ingrese pais elegido para calcular el porcentaje: ")
	_, err3 := fmt.Scan(&entradaPorcentaje)

	if err3 != nil {
		log.Fatal(err3)
		os.Exit(1)
	}

	go func(chan string, chan error) {
		totalTickets := 0
		for i := 0; i < len(storage.Tickets); i++ {
			totalTickets++
		}

		porcentaje, err := tickets.AverageDestination(entradaPorcentaje, totalTickets)
		if err != nil {
			canalErr <- err
		}
		mensaje := fmt.Sprintf("El porcentaje de personas que viajan al destino %s es %f.", entradaPorcentaje, porcentaje)
		canalPorcentajePorDestino <- mensaje
	}(canalPorcentajePorDestino, canalErr)

	//Impresion de Canales
	for i := 0; i < 3; i++ {
		select {
		case totalTicket := <-canalTotalTickets:
			fmt.Println(totalTicket)
		case ticketPorHorario := <-canalViajantesPorHorario:
			fmt.Println(ticketPorHorario)
		case porcentajePorDestino := <-canalPorcentajePorDestino:
			fmt.Println(porcentajePorDestino)
		case err := <-canalErr:
			fmt.Println("Error:", err)
		}
	}
}
