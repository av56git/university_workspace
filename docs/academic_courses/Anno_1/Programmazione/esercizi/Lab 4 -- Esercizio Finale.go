// Lab 4 -- Esercizio Finale 

package main
import ("fmt"; "math")

func main() {
	var op byte // simbolo dell'operazione (+,-,*,f,c,...)
	var risultato float64 // risultato
	var x float64 // operando attuale
	
	for true {
		fmt.Printf("%f\n> ", risultato) // prompt e risultato
		fmt.Scanf("%c%f", &op, &x) // lettura operazione (byte) e operando (float)
		
		if op == 'x' || op == 'q' {
			fmt.Println("BYE BYE!")
			break; // esci
		} else
		if op == '+' {
			risultato += x
		} else
		if op == '*' {
			risultato *= x
		} else
		if op == '/' {
			risultato /= x
		} else
		if op == '-' {
			risultato -= x
		} else
		if op == '^' {
			risultato = math.Pow(risultato, x)
		} else
		if op == 'c' {
			risultato = math.Ceil(risultato)
		} else
		if op == '=' {
			risultato = x
		} else
		if op == 'f' {
			risultato = math.Floor(risultato)
		} else {
			fmt.Println("operazione non valida")
		}
	}
}
