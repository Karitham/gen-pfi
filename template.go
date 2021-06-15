package main

import "text/template"

var t, _ = template.New("templatePFI").Parse(templatePFI)

// Definition are PFI definitions
type Definition struct {
	WordEN    string
	WordFR    string
	MeaningEN string
	MeaningFR string
	Notes     string
	Sources   []Source
}

// Source are the source links for the definitions
type Source struct {
	Link string
}

const templatePFI = `# PFI
{{ range .Definitions }}
## {{ .WordEN }}
{{ .MeaningEN }}
{{if ne .WordFR ""}}### {{ .WordFR }}

{{ .MeaningFR }}{{end}}
{{if ne .Notes ""}}### Notes

{{ .Notes }}
{{ end }}
### Sources
{{range .Sources}}
{{if ne .Link ""}}- [{{.Link}}]({{.Link}}){{end}}{{end}}
{{end}}`
