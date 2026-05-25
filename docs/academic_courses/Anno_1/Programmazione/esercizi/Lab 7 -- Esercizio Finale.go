package main
import "fmt"
import "strings"
import "math"

func main() {
	var s1, s2 string
	fmt.Print("Prima sequenza: ")
	fmt.Scan(&s1)
	fmt.Print("Seconda sequenza: ")
	fmt.Scan(&s2)
	best_i, best_score := 0, 0.0
	for i := 0; i < len(s1); i++ {
		// calcola score dell'allineamento tra s1[i:] ed s2
		score := Score(s1, s2, i)
		if score > best_score {
			best_score = score
			best_i = i
		}
	}
	StampaAllineamento(s1, s2, best_i)
	fmt.Println("Punteggio:", best_score)
}

func StampaAllineamento(s1, s2 string, i int) {
	/* stampa l'allineamento di s1 ed s2, con s2
	spostata di i posizioni a destra
	*/
	fmt.Println(s1)
	fmt.Print(strings.Repeat(" ", i))
	fmt.Println(s2)
}

func Score(s1, s2 string, k int) float64 {
	/* calcola lo score dell'allineamento tra s1 ed s2 */
	count := 0
	s1 = s1[k:]
	n := int(math.Min(float64(len(s1)), float64(len(s2))))
	for i := 0; i < n; i++ {
		if s1[i] == s2[i] {
			count++
		}
	}
	return float64(count)/float64(n)
}