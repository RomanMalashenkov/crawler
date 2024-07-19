package app

import (
	"fmt"
	"os"

	"github.com/RomanMalashenkov/crawler.git/internal/utils"

	"github.com/RomanMalashenkov/crawler.git/internal/crawler"
)

const maxDepth = 3 // максимальная глубина

func Run() {
	utils.SetupLogging() // настройки логирования

	var startUrl, targetUrl string
	fmt.Print("Введите начальную ссылку: ")
	fmt.Fscan(os.Stdin, &startUrl)
	fmt.Print("Введите целевую ссылку: ")
	fmt.Fscan(os.Stdin, &targetUrl)
	fmt.Println()

	startNode := &crawler.Node{Url: startUrl, Depth: 0}
	targetNode := &crawler.Node{Url: targetUrl, Depth: maxDepth}

	c := crawler.NewSimpleCrawler(startNode, targetNode, maxDepth)
	c.Search()
}
