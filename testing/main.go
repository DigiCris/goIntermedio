package main

func Operacion(op string, a, b int) int {
	switch op {
	case "sum":
		return a + b
	case "mul":
		return a * b
	default:
		return 0
	}
}

func GetMax(x,y int) int {
	if x>y {
		return x
	}else {
		return y
	}
}
/*
func Fibonacci(n int) int {
	if n>1 {
		return Fibonacci(n-2)+Fibonacci(n-1);
	} else {
		return n;
	}
}
*/
func Fibonacci(n int) int {
	if n==0 {
		return 0
	}
	if n==1{
		return 1;
	}
	nmenos1 := 0
	nmenos2 := 1
	r := 0

	for i := 0; i < n; i++ {
		r = nmenos1 + nmenos2
		nmenos2 = nmenos1
		nmenos1 = r
	}
	return r
}
