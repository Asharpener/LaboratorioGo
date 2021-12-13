package main

import (
	"fmt"
	"strconv"
)

/*
legge da standard input un numero intero
e stampa a video ogni cifra diversa da quella precedente
*/
func main() {
	var numero int

	fmt.Scan(&numero)

	s := strconv.Itoa(numero)                 // s è il mio numero convertito in stringa (così possiamo prendere ogni posizione singolarmente
	fmt.Print(string(s[0]))                   // il valore alla prima posizione è sicuramente diverso dal precedente (che non esiste) quindi lo stampo
	for indie := 1; indie < len(s); indie++ { //parto da 1 perchè quello a 0 lo stampo sicuro.
		if s[indie] != s[indie-1] { //se il valore attuale è diverso dal valore precedente,
			fmt.Print(string(s[indie])) //allora me lo stampi
		}
	}
	fmt.Println() //a capo

}
