package crawler

import (
	"fmt"

	"github.com/RomanMalashenkov/crawler.git/internal/parser"
)

// Node структура узла
type Node struct {
	Url    string
	Depth  int
	Text   string
	Parent *Node
}

// IsTargetNode проверяет, является ли узел целевым
func (n *Node) IsTargetNode(targetNode *Node) bool {
	return n.Url == targetNode.Url && n.Depth == targetNode.Depth
}

// Path функция для создания и вывода пути
func Path(node *Node) {
	path := []string{}
	sentences := []string{}

	for current := node; current != nil; current = current.Parent {
		path = append([]string{parser.DecodeURL(current.Url)}, path...)
		sentences = append([]string{current.Text}, sentences...)
	}

	for i := 1; i < len(path); i++ {
		fmt.Printf("%v------------------------\n", i)
		fmt.Println(sentences[i])
		fmt.Println(path[i])
	}
	fmt.Println()
}
