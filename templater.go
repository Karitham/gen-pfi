package main

import "text/template"

type templater interface {
	Template() *template.Template
}

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
