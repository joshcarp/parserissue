package main

import (
	"fmt"

	"github.com/anz-bank/sysl/pkg/parse"
	"github.com/spf13/afero"
)

func main() {
	fs := afero.NewOsFs()
	m, _ := parse.NewParser().Parse("example.sysl", fs)
	for _, app := range m.GetApps() {
		for _, t := range app.GetTypes() {
			tuple := t.GetTuple()
			for _, field := range tuple.GetAttrDefs() {
				// Should only print out " @json_tag = "pwd"" but also prints out the attribute of the type decl
				// map[description:s:"blah nah brah" json_tag:s:"pwd"]
				fmt.Println(field.GetAttrs())
			}
		}
	}
}
