package main

import (
	"fmt"
	"os"
	"strconv"
)

// Rimuovi restituisce la lista degli interi ottenibili rimuovendo d cifre da n (ricorsiva)
func Rimuovi(n int, d int) []int {

	// Caso base 1: nessuna cifra da rimuovere
	if d == 0 {
		return []int{n}
	}

	// Conto le cifre di n
	cifre := 0
	temp := n
	if temp == 0 {
		cifre = 1
	}
	for temp > 0 {
		cifre++
		temp /= 10
	}

	// Caso base 2: cifre da rimuovere >= cifre disponibili
	if d >= cifre {
		return []int{0}
	}

	// Passo ricorsivo: rimuovo 1 cifra in ogni posizione, poi chiamo Rimuovi con d-1
	var risultati []int
	potenza := 1

	for i := 0; i < cifre; i++ {
		parteDestra := n % potenza
		parteSinistra := n / (potenza * 10)
		nuovo := parteSinistra*potenza + parteDestra

		sottoRisultati := Rimuovi(nuovo, d-1)
		risultati = append(risultati, sottoRisultati...)

		potenza *= 10
	}

	return risultati
}

func main() {

	n, _ := strconv.Atoi(os.Args[1])
	d, _ := strconv.Atoi(os.Args[2])

	lista := Rimuovi(n, d)

	// Se la lista è vuota, il risultato è 0
	if len(lista) == 0 {
		fmt.Println(0)
		return
	}

	// Trovo il minimo
	minimo := lista[0]
	for _, v := range lista {
		if v < minimo {
			minimo = v
		}
	}

	fmt.Println(minimo)
}
