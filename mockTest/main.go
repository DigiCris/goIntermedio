package main

import "fmt"
import "time"

type Person struct {
	DNI string
	Name string
	Age int
}

type Employee struct {
	Id int
	Position string
}

type FullTimeEmployee struct {
	Person
	Employee
}

/*
Simulacion de la devolucion en bases de datos
*/
var  GetPersonByDNI = func(dni string) (Person,error)  {
	time.Sleep(5 * time.Second);
	return Person{}, nil
}
var GetEmployById = func(id int) (Employee,error)  {
	time.Sleep(5 * time.Second);
	return Employee{}, nil
}
/////////////////////////////////////////////

func GetFullTimeEmployeeByIdandDNI(dni string, id int) (FullTimeEmployee,error) {
	ftEmployee := FullTimeEmployee{}

	e, err := GetEmployById(id);
	if err != nil{
		fmt.Printf("Error al obtener el employee: %v\n", err)
		return FullTimeEmployee{}, err

	}
	ftEmployee.Employee = e

	p, err := GetPersonByDNI(dni);
	if err != nil{
		fmt.Printf("Error al obtener el Person: %v\n", err)
		return FullTimeEmployee{}, err
	}
	ftEmployee.Person = p
	return ftEmployee,nil
}