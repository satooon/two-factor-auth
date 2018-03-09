//go:generate go-bindata -pkg template -ignore=\\*.go -o template_gen.go ./...
package template

import (
	"net/http"
	tmp "text/template"
	"github.com/gin-gonic/gin/render"
	"fmt"
)

var htmlContentType = []string{"text/html; charset=utf-8"}

type htmlRender struct {
	Layout []string
}

func NewHtmlRender(layout ...string) *htmlRender {
	return &htmlRender{
		Layout: layout,
	}
}

func (r *htmlRender) Instance(name string, obj interface{}) render.Render {
	t := tmp.New("complex")
	arr := make([]string, len(r.Layout) + 1)
	copy(arr, r.Layout)
	arr[len(r.Layout)] = name

	for _, v := range arr {
		b, err := Asset(v)
		if err != nil {
			panic(fmt.Errorf("template.Asset name:%s error: %v\n", v, err))
		}
		if _, err := t.Parse(string(b)); err != nil {
			panic(fmt.Errorf("template Parse name:%s error: %v\n", v, err))
		}
	}

	return html{
		Template: t,
		Data: obj,
	}
}

type html struct {
	Template *tmp.Template
	Data interface{}
}

func (h html) Render(w http.ResponseWriter) error {
	h.WriteContentType(w)
	return h.Template.Execute(w, h.Data)
}

func (h html) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = htmlContentType
	}
}