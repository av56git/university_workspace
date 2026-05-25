// Lab 03 -- Esercizio Finale

package main
import "fmt"

func main() {
	
	var a, b, c, d int

	// leggo 4 interi nella forma a.b.c.d
	fmt.Scanf("%d.%d.%d.%d", &a, &b, &c, &d)

	if a < 0 || a > 255 || b < 0 || b > 255 || c < 0 || c > 255 || d < 0 || d > 255 {
		fmt.Println("Indirizzo non valido!")
		return
	}

	// stampa in binario, con riempimento di zeri a 
	// sinistra, fino ad avere 8 cifre
	fmt.Printf("%08b.%08b.%08b.%08b\n", a, b, c, d)

	// stampa la classe
	if a >= 0 && a < 128 {
		fmt.Println("Class A")
	}
	if a >= 128 && a < 192 {
		fmt.Println("Class B")
	}
	if a >= 192 && a < 224 {
		fmt.Println("Class C")
	}
	if a >= 224 && a < 240 {
		fmt.Println("Class D")
	}
	if a >= 241 && a < 256 {
		fmt.Println("Class E")
	}

	// se l'indirizzo è riservato, stampa info a riguardo...
	// (TODO)

}
