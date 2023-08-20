package tickets

import (
	"os"
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
func (s Storage) GetTotalTickets(destination string) (int, error) {}

// Requerimiento 2
func GetMornings(time string) (int, error) {}

// Requerimiento 3
func AverageDestination(destination string, total int) (int, error) {}