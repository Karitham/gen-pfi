package main

import (
	"io"
	"os"
	"strings"

	"github.com/Karitham/pfigen/source"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	log.Logger.WithLevel(zerolog.TraceLevel)

	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal().Err(err).Msg("invalid input, please use stdin (`cat urls.txt | pfigen > pfi.md`")
	}

	lines := strings.Split(string(b), "\n")

	var d []Definition
	for _, line := range lines {
		ENPage, err := source.Lookup(line)
		if err != nil {
			log.Err(err).Str("word", line).Msg("error in getting word")
		}

		var FRPage source.Page
		for _, l := range ENPage.Languages {
			if l.Language == "fr" {
				FRPage, err = source.Lookup(l.Link)
				if err != nil {
					log.Err(err).Str("word", line).Msg("error in getting word")
				}
			}
		}

		d = append(d, Definition{
			WordEN:    ENPage.Title,
			MeaningEN: ENPage.Abstract,
			WordFR:    FRPage.Title,
			MeaningFR: FRPage.Abstract,
			Sources:   []Source{{Link: ENPage.URL}, {Link: FRPage.URL}},
		})
	}

	type D struct {
		Definitions []Definition
	}
	if err := t.Execute(os.Stdout, &D{d}); err != nil {
		log.Panic().Err(err).Msg("error exc template")
	}
}
