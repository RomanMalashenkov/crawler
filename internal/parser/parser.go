package parser

import (
	"log"
	"net/url"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
)

// DecodeURL декодирует ссылку
func DecodeURL(nurl string) string {
	decURL, _ := url.QueryUnescape(nurl)
	return decURL
}

// SplitIntoSentences разделяет параграфы на предложения (в HTML формате)
func SplitIntoSentences(html string) []string {
	var sentences []string
	var sentence []rune
	openBrackets := 0

	for i := 0; i < len(html); {
		char, size := utf8.DecodeRuneInString(html[i:])
		sentence = append(sentence, char)

		switch char {
		case '(', '[', '{':
			openBrackets++
		case ')', ']', '}':
			openBrackets--
		}

		if (char == '.' || char == '!' || char == '?') && openBrackets == 0 {
			if i+size < len(html) {
				nextChar, nextSize := utf8.DecodeRuneInString(html[i+size:])
				if unicode.IsSpace(nextChar) {
					if i+size+nextSize < len(html) {
						nextNextChar, _ := utf8.DecodeRuneInString(html[i+size+nextSize:])
						if unicode.IsUpper(nextNextChar) {
							trimmedSentence := strings.TrimSpace(string(sentence))
							if trimmedSentence != "" {
								sentences = append(sentences, trimmedSentence)
							}
							sentence = []rune{}
						}
					}
				}
			}
		}

		i += size
	}

	if len(sentence) > 0 {
		trimmedSentence := strings.TrimSpace(string(sentence))
		if trimmedSentence != "" {
			sentences = append(sentences, trimmedSentence)
		}
	}

	return sentences
}

// FindLinksInSentence находит ссылки в предложениях формата HTML
func FindLinksInSentence(sentenceHTML string) []string {
	var links []string
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(sentenceHTML))
	if err != nil {
		log.Println("Ошибка:", err)
	}

	doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		link, _ := item.Attr("href")
		if strings.HasPrefix(link, "/wiki/") {
			links = append(links, link)
		}
	})

	return links
}

// HtmlToText конвертирует HTML в текст
func HtmlToText(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}
	return doc.Text()
}
