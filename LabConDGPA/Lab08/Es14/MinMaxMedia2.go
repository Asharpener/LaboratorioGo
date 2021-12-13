package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var slInt []int = LeggiNumeri()
	fmt.Println("Minimo: ", Minimo(slInt)) //numeri è un parametro effettivo. Quando va nel sottoprogramma è un parametro formale (sl nel mio caso)
	fmt.Println("Massimo: ", Massimo(slInt))
	fmt.Println("Media: ", Media(slInt))

}

func LeggiNumeri() (numeri []int) {
	for _, v := range os.Args[1:] {
		n, err := strconv.Atoi(v) //n è il valore convertito, err è nil se non ci sono errori
		if err == nil {
			numeri = append(numeri, n)
		}
	}
	return
}

func Minimo(sl []int) int { //sl è un parametro formale. NOTA: viene passato per copia
	var minimo int
	minimo = sl[0] //prendo il valore alla prima posizione dello slice di interi di sl e lo assegna a minimo
	// potevo direttamente fare var minimo int = sl[0] oppure minimo := sl[0]

	for i := 1; i < len(sl); i++ { //i deve scorrere tutta la LUNGHEZZA della slice!!
		if sl[i] < minimo {
			minimo = sl[i]
		}

	}
	return minimo

}

func Massimo(sl []int) int {
	var massimo int
	massimo = sl[0]

	for i := 1; i < len(sl); i++ {
		if sl[i] > massimo {
			massimo = sl[i]
		}
	}
	return massimo

}

func Media(sl []int) float64 {
	var somma, i, v int
	for i, v = range sl {
		somma += v //sommma = sommma + v
	}
	return float64(somma) / float64(i+1)

}
