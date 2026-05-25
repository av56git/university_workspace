package main
import "fmt"

func main() {
	var c1, c2, C int
	fmt.Print("Inserisci C, c1, c2: ")
	fmt.Scan(&C, &c1, &c2)
	n := 0
	trovata := false
	for {
		n1, n2, esiste := Resto(C, c1, c2, n)
		if esiste {
			trovata = true
			fmt.Printf("%d€ = %d * %d€ + %d * %d€\n", C, n1, c1, n2, c2)
			n = n1+1
		} else {
			break
		}
	}
	if !trovata {
		fmt.Println("Nessuna soluzione")
	}
}

func Resto(C, c1, c2, n int) (n1, n2 int, esiste bool) {
	for D := n*c1; D <= C; D += c1 {
		if (C-D) % c2 == 0 {
			return D/c1, (C-D)/c2, true
		}
	}
	return 0, 0, false
}