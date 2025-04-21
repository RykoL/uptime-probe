package testutils

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"
	"io"
)

func RenderComponent(component templ.Component) *goquery.Document {
	r, w := io.Pipe()
	go func() {
		component.Render(context.Background(), w)
		w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)

	if err != nil {
		panic(err)
	}

	return doc
}
