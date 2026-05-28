package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	righe := Leggi()
	conteggi := Conta(righe)
	var paroleFiltrate []string
	for parola, conto := range conteggi {
		if conto >= k {
			paroleFiltrate = append(paroleFiltrate, parola)
		}
	}
	sort.Strings(paroleFiltrate)
	for _, parola := range paroleFiltrate {
		fmt.Printf("%s %d\n", parola, conteggi[parola])
	}
}

func Leggi() (righe [][]string) {
	fmt.Println("Inserire testo. Premi CTRL+D per terminare:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		righe = append(righe, words)
	}
	return
}

func Conta(righe [][]string) map[string]int {
	m := make(map[string]int)
	for _, riga := range righe {
		seen := make(map[string]bool)
		for _, parola := range riga {
			if !seen[parola] {
				seen[parola] = true
				m[parola]++
			}
		}
	}
	return m
}
