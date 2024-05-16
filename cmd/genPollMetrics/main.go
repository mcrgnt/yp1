//go:build ignore

package main

import (
	"os"
	"reflect"
	"text/template"

	"github.com/mcrgnt/yp1/internal/metrics"
)

var (
	tmpl = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package metrics

import (
	"github.com/mcrgnt/yp1/internal/storage"
	"github.com/mcrgnt/yp1/internal/common"
)

func pollMetrics(params *PollMetricsParams) {
	{{range .}}_ = params.Storage.MetricSet(&storage.StorageParams{
		Type: common.MetricTypeGauge,
		Name: "{{.Name}}",
		Value: {{.Value}},
	})
	{{end}}
}
`))
)

type data struct {
	Name  string
	Value string
}

func main() {
	f, err := os.Create("pollMetrics.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	datas := []data{}
	val := reflect.ValueOf(metrics.MemStats).Elem()
	for _, name := range metrics.PollMetricsFromMemStatsList {
		switch val.FieldByName(name).Interface().(type) {
		case uint32, uint64:
			datas = append(datas, data{
				Name:  name,
				Value: "float64(MemStats." + name + ")",
			})
		default:
			datas = append(datas, data{
				Name:  name,
				Value: "MemStats." + name,
			})

		}
	}
	tmpl.Execute(f, datas)
}
