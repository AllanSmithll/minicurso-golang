/*
Allan Alves Amâncio - Minicurso de Golang - IFTECH - IFPB
Conteúdo: Slices
Data: 16/10/2023
*/
package main

import "fmt"

func main() {
	var arrayNotas [3]float64;
	var sliceNotasBoas []float64;
	fmt.Println(arrayNotas)

	arrayNotas = [3]float64{10.0, 10.0, 7.0}
	sliceNotasBoas = arrayNotas[1:3] // 10.0, 10.0
	fmt.Println(sliceNotasBoas);
}