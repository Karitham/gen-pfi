package main

import (
	"flag"
	"io"
	"os"
	"strings"
	"time"

	"github.com/Karitham/pfigen/source"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	var t templater
	var verbose bool
	var tname string

	flag.BoolVar(&verbose, "v", false, "run the program using verbose mode")
	flag.StringVar(&tname, "t", "basic", "run the program using the named template")
	flag.Parse()

	start := time.Now()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
	log.Logger = log.Level(zerolog.InfoLevel)

	if verbose {
		log.Logger = log.Level(zerolog.TraceLevel)
	}

	switch tname {
	case "basic":
		t = Basic{}
	}

	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal().Err(err).Msg("invalid input, please use stdin (`cat urls.txt | pfigen > pfi.md`")
	}

	lines := strings.Split(string(b), "\n")

	var d []Definition
	for _, line := range lines {
		strs := strings.SplitN(line, " ", 2)
		url := strs[0]

		notes := ""
		if len(strs) > 1 {
			notes = strs[1]
		}

		ENPage, err := source.Lookup(url)
		if err != nil {
			log.Err(err).Str("word", line).Msg("error in getting word")
		}
		log.Trace().Str("url", url).Msg("looked up english page")

		var FRPage source.Page
		for _, l := range ENPage.Languages {
			if l.Language == "fr" {
				log.Trace().Str("url", l.Link).Msg("found french page")
				FRPage, err = source.Lookup(l.Link)
				if err != nil {
					log.Err(err).Str("word", line).Msg("error in getting word")
				}
			}
		}

		d = append(d, Definition{
			WordEN:    ENPage.Title,
			MeaningEN: strings.TrimSpace(ENPage.Abstract),
			WordFR:    FRPage.Title,
			MeaningFR: strings.TrimSpace(FRPage.Abstract),
			Notes:     notes,
			Sources:   []Source{{Link: ENPage.URL}, {Link: FRPage.URL}},
		})
		log.Debug().Dict("values", zerolog.Dict().Str("url", url).Str("notes", notes)).Msg("Found the required pages")
	}

	type D struct {
		Definitions []Definition
	}
	if err := t.Template().Execute(os.Stdout, &D{d}); err != nil {
		log.Panic().Err(err).Msg("error exc template")
	}

	log.Debug().Msgf("took %dms", time.Since(start).Milliseconds())
}
