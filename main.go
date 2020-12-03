package main

import "fmt"

const (
	//kb = 1024
	kb = 1 << 10
	mb = 1024 * kb
	gb = 1024 * mb
)

func main() {

	imprimir(kb)
	imprimir(mb)
	imprimir(gb)

}

func imprimir(n int) {

	fmt.Printf("%d\t%b\n", n, n)

}
