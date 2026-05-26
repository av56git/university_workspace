package main

import (
	"fmt"
	"os"
	"strconv"
)

// Rimuovi restituisce la lista degli interi ottenibili rimuovendo una cifra da n
func Rimuovi(n int) []int {

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

	risultati := []int{}
	potenza := 1

	for i := 0; i < cifre; i++ {
		parteDestra := n % potenza
		parteSinistra := n / (potenza * 10)
		nuovo := parteSinistra*potenza + parteDestra
		risultati = append(risultati, nuovo)
		potenza *= 10
	}

	return risultati
}

func main() {

	n, _ := strconv.Atoi(os.Args[1])
	d, _ := strconv.Atoi(os.Args[2])

	// Parto con la lista contenente solo n
	lista := []int{n}
	fmt.Println("Lista iniziale ", lista)

	// Applico Rimuovi d volte
	for i := 0; i < d; i++ {
		nuovaLista := []int{}
		for _, v := range lista {
			nuovaLista = append(nuovaLista, Rimuovi(v)...)
		}
		lista = nuovaLista
		fmt.Printf("Lista dopo %d rimozioni: %v\n", i+1, lista)
	}

	// Se la lista è vuota (tutte le cifre rimosse), il risultato è 0
	if len(lista) == 0 {
		fmt.Println(0)
		return
	}

	// Trovo il minimox
	minimo := lista[0]
	for _, v := range lista {
		if v < minimo {
			minimo = v
		}
	}

	fmt.Println(minimo)
}
