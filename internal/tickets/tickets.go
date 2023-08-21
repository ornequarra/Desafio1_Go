package tickets

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// Ticket es una estructura que define todos los tickets de la aerolinea
type Ticket struct {
	id          string
	nombre      string
	email       string
	paisDestino string
	horaVuelo   string
	precio      string
}

// Storage es una estructura que almacena todos los datos de tickets
type Storage struct {
	Tickets []Ticket
}

// readFile es una funcion que lee el archivo tickets.csv
func ReadFile(filename string) []Ticket {

	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\n")

	var resultado []Ticket
	for i := 0; i < len(data); i++ {
		if len(data[i]) > 0 {
			file := strings.Split(string(data[i]), ",")
			tickets := Ticket{
				id:          file[0],
				nombre:      file[1],
				email:       file[2],
				paisDestino: file[3],
				horaVuelo:   file[4],
				precio:      file[5],
			}
			resultado = append(resultado, tickets)
		}
	}
	return resultado
}

// Requerimiento 1
func (s Storage) GetTotalTickets(destination string, tickets []Ticket) (int, error) {
	var contador int = 0
	for _, ticket := range tickets {
		if ticket.paisDestino == destination {
			contador++
		}
	}
	if contador == 0 {
		return contador, errors.New("No se encontraron viajes en ese país")
	}
	return contador, nil
}

// Requerimiento 2
func (s Storage) GetCountByPeriod(time string, tickets []Ticket) (int, error) {
	var desde int64
	var hasta int64
	var contador int
	switch time {
	case "madrugada":
		desde = 0
		hasta = 6
	case "mañana":
		desde = 7
		hasta = 12
	case "tarde":
		desde = 13
		hasta = 19
	case "noche":
		desde = 20
		hasta = 23
	default:
		return 0, errors.New("Usted ingresó un período incorrecto")
	}

	for _, ticket := range tickets {
		horario := strings.Split(ticket.horaVuelo, ":")
		hora, _ := strconv.ParseInt(horario[0], 10, 64)
		if hora > desde && hora < hasta {
			contador++
		}
	}
	return contador, nil
}

// Requerimiento 3
func AverageDestination(destination string, total int) (int, error) { return 0, nil }
