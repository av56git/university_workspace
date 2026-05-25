package main
import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {
	n, E := LeggiGrafo()
	fmt.Println(n)
	for _, s := range E {
		fmt.Println(Arco(s))
	}
	fmt.Println("Trovati", Triangoli(n, E), "triangoli.")
}

func LeggiGrafo() (n int, E []string) {
	fmt.Scan(&n)
	E = make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		E = append(E, scanner.Text())
	}
	return
}

func Arco(s string) (u, v int) {
	tokens := strings.Fields(s)
	u, _ = strconv.Atoi(tokens[0])
	v, _ = strconv.Atoi(tokens[1])
	if v < u {
		u, v = v, u
	}
	return
}

func Triangolo(u, v, z int, n int, E[] string) bool {
	var uv, vz, uz bool = false, false, false
	for _, e := range E {
		\	x, y := Arco(e)
		uv = uv || (x == u && y == v)
		uz = uz || (x == u && y == z)
		vz = vz || (x == v && y == z)
	}
	return uv && uz && vz
}

func Triangoli(n int, E[] string) int {
	var T int = 0
	for u := 1; u <= n; u++ {
		for v := u+1; v <= n; v++ {
			for z := v+1; z <= n; z++ {
				if Triangolo(u, v, z, n, E) {
					T++
				}
			}
		}
	}
	return T
}


