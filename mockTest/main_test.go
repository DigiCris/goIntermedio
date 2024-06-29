package main

import "testing"

type response struct {
	id int
	dni string
	mockFunc func()
	response FullTimeEmployee
}

func TestGetFullTimeEmployeeByIdandDNI(t *testing.T) {
	table := []response{
		{
			1,
			"2",
			func() {
				// Cuando se declaran las funciones se pone var antes, acá que se reasigna sin la palabra var siendo así una variable global de tipo funcion
				GetEmployById = func(id int) (Employee,error)  {
					return Employee{1,"CEO"}, nil
				}
				GetPersonByDNI = func(dni string) (Person,error)  {
					return Person{"2","cris",35}, nil
				}
			},
			FullTimeEmployee{
				Person{"2","cris",35},
				Employee{1,"CEO"},
			},
		},
	}

	originalGetEmployById := GetEmployById // hago backup de funciones originales antes de mockearlas
	originalGetPersonByDNI := GetPersonByDNI
	for _, test := range table {
		test.mockFunc() // cambia las funciones por las del mock
		fte, err := GetFullTimeEmployeeByIdandDNI(test.dni,test.id)
		if err != nil {
			t.Errorf("erro getting employee")
		}

		if fte.Age != test.response.Age {
			t.Errorf("edades no coinciden")
		}
	}
	GetPersonByDNI = originalGetPersonByDNI // vuelvo a poner las funciones originales
	GetEmployById = originalGetEmployById

}