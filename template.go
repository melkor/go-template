package main

import (
	"os"
	"text/template"
)

type InnerStruct struct {
	Name string
}

type Data struct {
	TestBool        bool
	TestString      string
	TestInt         int
	TestSlice       []string
	TestMap         map[string]string
	TestSliceStruct []InnerStruct
	TestMapStruct   map[string]InnerStruct
}

func main() {

	data := Data{
		TestBool:   true,
		TestString: "Hello, I'm String",
		TestInt:    12,
		TestSlice:  []string{"I", "'m", "a", "slice"},
		TestMap: map[string]string{
			"coucou": "salut",
			"ca":     "va",
		},
		TestSliceStruct: []InnerStruct{
			{
				Name: "one",
			},
			{
				Name: "two",
			},
		},
		TestMapStruct: map[string]InnerStruct{
			"lalaun": InnerStruct{
				Name: "coucouun",
			},
		},
	}

	t := template.Must(
		template.New("todos").Parse(`
      Content:{{ block "content" .}}
        - Bool: {{ .TestBool }}
        - String: {{ .TestString }}
        - Integer: {{ .TestInt }}
        - Slice: {{ range .TestSlice }}
            * {{ .}}{{ end }}
        - Map: {{ with .TestMap }}
            * {{ index . "coucou" }}
            * {{ index . "ca" }}{{ end }}
            * {{ index .TestMap "coucou"}}
				- Slice Struct: {{ range .TestSliceStruct }}
				    * {{ .Name}}{{ end }}
				- Map Struct:
				    * {{  .TestMapStruct.lalaun.Name }}
				- Map test: {{ range $key, $value := .TestMap}}
				    * {{ $key }}: {{ $value }}{{end}}
				{{end}}

`))

	err := t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	t2, err := template.Must(
		t.Clone()).Parse(`{{define "content"}} coucou
{{end}}`)

	err = t2.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
