package main

import "text/template"

var t, _ = template.New("templatePFI").Parse(templatePFI)

const templatePFI = `
# PFI
{{ range .Definitions }}
## {{ .WordEN }}
{{ .MeaningEN }}
{{if ne .WordFR ""}}### {{ .WordFR }}

{{ .MeaningFR }}{{end}}
{{if ne .Notes ""}}### Notes

{{ .Notes }}
{{ end }}### Sources
{{range .Sources}}
{{if ne .Link ""}}- [{{.Link}}]({{.Link}}){{end}}{{end}}
{{end}}`
