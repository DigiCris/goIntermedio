/*
Testing

Correr pruebas y ver si pasan o no:
go test

Ver que porcentaje de testing está testeado
go test -cover 

Generamos archivos para ver que partes fueron testeadas y cuales no
go test -coverprofile=coverage.out

coverage.out será ilegible. Para volverlo entendible usamos:
go tool cover -func=coverage.out
Esto me dirá el porcentaje de covertura de cada función

Si quiero ver gráficamente por donde si y por donde no se está testeando uso:
go tool cover -html=coverage.out -o coverage.html
y luego lo veré en el archivo html






Optimizacion

Hace test generando un archivo con detalles de tiempos en cada parte del codigo
go test -cpuprofile=cpu.out

Analizar el archivo anterior anterior con pprof. abrirá pprof y cargando cpu.ou para analizar
go tool pprof cpu.out

Para ver el consumo de cada función en macro uso:
top

Luego al detectar que fibonacci fue la peor profundizo en donde esta con:
list fibonacci

ahí puedo ver linea por linea donde está el problema y detecar que "return Fibonacci(n-2)+Fibonacci(n-1);"
es la peor y así poder mejorarla.

exportar gráfico de desarrollo de tiempos
pdf
png
web => no me funciona en mi wsl
*/

package main

import "testing"

type operation struct {
	op string
	a int
	b int
	c int
}

func TestOp(t *testing.T) {
	cases := []operation{
		{"sum",5,5,10},
		{"sum",6,5,11},
		{"sum",7,5,12},
		{"mul",7,5,35},
		{"div",7,5,0},
	}
	for _, item := range cases {
		total := Operacion(item.op,item.a,item.b)
		if total != item.c {
			t.Errorf("Sum was incorrect, got %d instead of %d",total,item.c)
		}		
	}

}

func TestGetMax(t *testing.T) {
	cases := []struct{
		a int
		b int
		r int
	}{
		{1,2,2},
		{1,3,3},
		{3,2,3},
		{3,1,3},
		{1,1,1},
	}
	for _, item := range cases {
		max := GetMax(item.a,item.b)
		if max != item.r {
			t.Errorf("Max should be %d instead of %d", item.r,max);
		}
	}
}

func TestFibonacci(t *testing.T) {
	cases := []struct{
		a int
		r int
	}{
		{0,0},
		{1,1},
		{5,5},
		{10,55},
		{45,1134903170},
	}

	for _, item := range cases {
		res := Fibonacci(item.a)
		if item.r != res {
			t.Errorf("fibonacci value should be %d instead of %d",item.r,res);
		}
	}
}