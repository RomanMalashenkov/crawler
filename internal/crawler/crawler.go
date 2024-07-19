package crawler

import (
	"log"

	"github.com/RomanMalashenkov/crawler.git/internal/parser"

	"github.com/gocolly/colly"
)

// SimpleCrawler структура для простого поиска и обработки узлов
type SimpleCrawler struct {
	MaxDepth    int
	Stack       []*Node
	Target      *Node
	TargetFound bool // флаг указания о нахождении целевой ссылки
}

// NewSimpleCrawler создает новый экземпляр SimpleCrawler
func NewSimpleCrawler(startNode, targetNode *Node, maxDepth int) *SimpleCrawler {
	return &SimpleCrawler{
		MaxDepth: maxDepth,
		Stack:    []*Node{startNode},
		Target:   targetNode,
	}
}

// Search выполняет поиск целевого узла
func (c *SimpleCrawler) Search() {
	for len(c.Stack) > 0 {
		if c.TargetFound {
			return
		}

		currentNode := c.popStack()
		if currentNode.Depth >= c.MaxDepth {
			continue
		}

		log.Println(parser.DecodeURL(currentNode.Url))

		collector := c.createCollector(currentNode)
		if err := collector.Visit(currentNode.Url); err != nil {
			log.Println("Не удалось посетить ссылку:", err)
		}
	}
}

// popStack возвращает элемент, находящийся наверху стека, а затем удаляет его.
func (c *SimpleCrawler) popStack() *Node {
	n := c.Stack[len(c.Stack)-1]
	c.Stack = c.Stack[:len(c.Stack)-1]
	return n
}

// createCollector создает новый коллектор и ищет параграфы
func (c *SimpleCrawler) createCollector(currentNode *Node) *colly.Collector {
	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, как Gecko) Chrome/99.0.9999.999 Safari/537.36"),
	)
	collector.OnHTML("p", func(e *colly.HTMLElement) {
		c.handleParagraphs(e, currentNode)
	})
	return collector
}

// метод handleParagraphs для обработки параграфа
func (c *SimpleCrawler) handleParagraphs(e *colly.HTMLElement, currentNode *Node) {
	if c.TargetFound {
		return
	}

	paragraphHTML, err := e.DOM.Html() // получаем HTML-код параграфа
	if err != nil {
		log.Println("Не удалось извлечь HTML-код параграфа:", err)
		return
	}

	sentences := parser.SplitIntoSentences(paragraphHTML)

	// обработка предложений (в HTML формате)
	for _, sentenceHTML := range sentences {
		links := parser.FindLinksInSentence(sentenceHTML)
		if len(links) == 0 {
			continue
		}

		sentenceText := parser.HtmlToText(sentenceHTML)

		// обработка ссылок
		for _, alink := range links {
			linkNode := &Node{
				Url:    "https://ru.wikipedia.org" + alink,
				Depth:  currentNode.Depth + 1,
				Text:   sentenceText,
				Parent: currentNode,
			}

			c.Stack = append(c.Stack, linkNode)

			if linkNode.IsTargetNode(c.Target) {
				c.TargetFound = true
				Path(linkNode)
				return
			}
		}
	}
}
