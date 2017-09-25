package htmlsafe

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

func Run(r io.Reader) (*bytes.Buffer, error) {
	tokenizer := html.NewTokenizer(r)
	state := &state{skipAt: -1}
	var buffer bytes.Buffer

outer:
	for {
		if tokenizer.Next() == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break outer
			} else if err == nil {
				panic("got html.ErrorToken without error")
			}
			return nil, err
		}

		t := tokenizer.Token()
		switch t.Type {
		case html.TextToken:
			if state.skipText() {
				break
			}
			switch state.topTag() {
			case "script", "style":
				// if allowed, don't escape CSS/JS
				buffer.WriteString(t.Data)
			default:
				buffer.WriteString(t.String())
			}

		case html.StartTagToken:
			if _, ok := voidMap[t.DataAtom]; !ok {
				if state.push(t.Data) {
					break // non-void, push says skip
				}
			} else if state.skipNode(t.Data) {
				break // void, we should skip
			}
			// TODO: strip attributes
			buffer.WriteString(t.String())

		case html.EndTagToken:
			renderEndTags(&buffer, state.popTo(t.Data))

		case html.SelfClosingTagToken:
			if state.skipNode(t.Data) {
				break
			}
			// TODO: strip attributes
			buffer.WriteString(t.String())

		case html.CommentToken:
			// ignore

		case html.DoctypeToken:
			// ignore
		}
	}

	renderEndTags(&buffer, state.popTo(""))
	return &buffer, nil
}

func renderEndTags(buffer *bytes.Buffer, tags []string) {
	for i := len(tags) - 1; i >= 0; i-- {
		buffer.WriteString("</")
		buffer.WriteString(tags[i])
		buffer.WriteByte('>')
	}
}
