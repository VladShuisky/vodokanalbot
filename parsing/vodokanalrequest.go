package parsing

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
)

func GetHtmlDataFromVodokanal() *html.Node {
	resp, err := http.Get(os.Getenv("VODOKANAL_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 2. Автоматически определяем кодировку и преобразуем в UTF-8
	utf8Reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}

	// 3. Парсим HTML
	doc, err := html.Parse(utf8Reader)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}