package main
import "fmt"
import "math"

func Valuta(a, b, c float64, x float64) float64 {
	return a*x*x + b*x + c
}

func Deriva(a, b, c float64) (a1, b1, c1 float64) {
	return 0, 2*a, b
}

func Newton(a, b, c, x float64) float64 {
	a1, b1, c1 := Deriva(a, b, c)
	return x - Valuta(a, b, c, x) / Valuta(a1, b1, c1, x)
}

func CercaRadice(a, b, c, x0 float64) float64 {
	var x float64 = x0
	for i:=0; i<1000; i++ {
		if math.Abs(Valuta(a, b, c, x)) <= 1e-10 {
			break
		} else {
			x = Newton(a, b, c, x)
		}
	}
	return x
}

func main() {
	var a, b, c float64
	var x float64
	var inserito bool = false
	for {
		fmt.Println("1) inserisci coefficienti")
		fmt.Println("2) valuta P")
		fmt.Println("3) trova radici di P")
		fmt.Println("4) esci")
		var scelta int
		fmt.Scan(&scelta)
		switch scelta {
		
		case 1:
			fmt.Print("Inserisci i coefficienti per x^2, x, 1: ")
			fmt.Scan(&a, &b, &c)
			inserito = true		

		case 2:
			if !inserito {
				fmt.Println("Devi prima inserire i coefficienti.")
				break
			}
			fmt.Print("Inserisci x: ")
			fmt.Scan(&x)
			y := Valuta(a, b, c, x)
			fmt.Print("P(", x, ") = ", y, "\n")
		
		case 3:
			if !inserito {
				fmt.Println("Devi prima inserire i coefficienti.")
				break
			}
			var x, y float64

			x = CercaRadice(a, b, c, -1e10) // radice #1, da sx
			y = Valuta(a, b, c, x)
			if math.Abs(y) <= 1e-10 {
				fmt.Println("Ho trovato", x, "in cui P vale", y)
			} else {
				fmt.Println("Non sembrano esserci radici!")
				break // nessuna radice, usciamo
			}

			x = CercaRadice(a, b, c, 1e10) // radice #2, da dx
			y = Valuta(a, b, c, x)
			if math.Abs(y) <= 1e-10 {
				fmt.Println("Ho trovato", x, "in cui P vale", y)
			}

		case 4:
			return
		}
	}
}

