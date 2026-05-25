package main
import "fmt"

type Node struct {
	val int
	next *Node
}

func main() {
	var testa Node = Node{-1, nil}
	var lista *Node = &testa
	Add(lista, 10)
	Add(lista, 20)
	Stampa(lista)
}

func Add(lista *Node, val int) {
	// creo nuovo (ultimo) nodo
	var node Node
	node.val = val
	node.next = nil
	// lo attacco all'ultimo nodo esistente
	var ptr *Node = lista
	for (*ptr).next != nil {
		ptr = (*ptr).next
	}
	(*ptr).next = &node
}

func Stampa(lista *Node) {
	var ptr *Node = lista
	for ptr != nil {
		var nodo Node = *ptr
		fmt.Println(nodo.val)
		ptr = nodo.next
	}
}

