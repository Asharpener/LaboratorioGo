package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var stringhe []string = os.Args[1:]
	var numeri []int = make([]int, 0, len(stringhe)) //creo dinamicamente una slice che ha dentro il tipo (int), lunghezza 0 (perchè siamo furbi) e capacità la metto come la lunghezza di "stringhe"
	for _, v := range stringhe {
		n, err := strconv.Atoi(v) //n è il valore convertito, err è nil se non ci sono errori
		if err == nil {
			numeri = append(numeri, n)
		}
	}
	fmt.Println("Minimo: ", Minimo(numeri)) //numeri è un parametro effettivo. Quando va nel sottoprogramma è un parametro formale (sl nel mio caso)
	fmt.Println("Massimo: ", Massimo(numeri))
	fmt.Println("Media: ", Media(numeri))

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
