package main

import (
	"regexp"
	"strings"

	bblfsh "gopkg.in/bblfsh/client-go.v2"
	"gopkg.in/bblfsh/client-go.v2/tools"
)

var bblfshClient *bblfsh.Client
var naturalLangage = regexp.MustCompile(`(\w|\d|\.|\!|\,\?)*`) // sorry to non english languages

func init() {
	var err error
	bblfshClient, err = bblfsh.NewClient("0.0.0.0:9432")
	if err != nil {
		panic(err)
	}
}

func fetchIdentifiers(language, file string) ([]string, error) {
	out := []string{}

	res, err := bblfshClient.NewParseRequest().Language(language).Content(file).Do()
	if err != nil {
		return nil, err
	}

	nodes, _ := tools.Filter(res.UAST, "//*[@roleIdentifier]")
	for _, n := range nodes {
		if text := cleanIdentifer(n.Token); text != "" {
			out = append(out, text)
		}
	}

	return out, nil
}

func cleanIdentifer(comment string) string {
	submatches := naturalLangage.FindAllString(comment, -1)
	comment = strings.Join(submatches, " ")

	// filter whitespace
	comment = strings.Replace(comment, "\n", " ", -1)
	comment = strings.Trim(comment, " ")

	return comment
}
