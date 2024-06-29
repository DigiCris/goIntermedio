package main

/*
Un worker es una función que se ejecuta concurrentemente para ir haciendo
diversas tareas.
Podemos poner tantos workers como queramos a escuchar y resolver tareas.
Se comunican con mi hilo principal mediante canales (Uno de ida con argumentos
otro de vuelta con los resultados a terminar).
El worker para su ejecución luego de lanzarlo espera tener en su canal de entrada alguna tarea a resolver.
El programa principal debe esperar a obtener los reultados de los workers lanzados escuchando los canales.
Los canales deben cerrarse al dejar de usarlos.

Usar para calcular fibonacci recursivos
*/

import "fmt"

func Fibonacci(n int) int {
	if n <= 1{
		return n;
	}
	n = Fibonacci(n-2) + Fibonacci(n-1);
	return n;
}

func Worker(id int, out chan<- int, ins <-chan int){
	for in := range ins {
		nout := Fibonacci(in);
		out <- nout
		fmt.Printf("worker %d => fib(%d)=%d",id,in,nout);
		fmt.Println("");		
	}

}

func main() {
	cantidadCanales := 8
	cantidadWorkers := 3
	cin := make(chan int, cantidadCanales)
	cout := make(chan int, cantidadCanales)

	for i := 0; i < cantidadWorkers; i++ {
		go Worker(i,cout,cin);
	}
	i := 0;

	i++
	cin <- 0
	i++
	cin <- 45
	i++
	cin <- 2
	i++
	cin <- 20
	i++
	cin <- 1
	i++
	cin <- 41
	i++
	cin <- 8

	for j:=0; j < i; j++ {
		<-cout
	}

	close(cin);
	close(cout);
}