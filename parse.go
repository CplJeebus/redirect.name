package main

import (
	"regexp"
	"strings"
)

type Config struct {
	From          string
	To            string
	RedirectState string
}

var configRE = regexp.MustCompile(`Redirects?(\s+.*)`)
var fromRE = regexp.MustCompile(`\s+from\s+(/\S*)`)
var toRE = regexp.MustCompile(`\s+to\s+(https?\://\S+|/\S*)`)
var stateRE = regexp.MustCompile(`\s+(permanently|temporarily)|\s+with\s+(301|302|307|308)`)

func Parse(record string) *Config {
	configMatches := configRE.FindStringSubmatch(record)
	if len(configMatches) == 0 {
		return nil
	}

	fromMatches := fromRE.FindStringSubmatch(configMatches[1])
	toMatches := toRE.FindStringSubmatch(configMatches[1])
	stateMatches := stateRE.FindStringSubmatch(configMatches[1])

	config := new(Config)
	if len(fromMatches) > 0 {
		config.From = fromMatches[1]
	}
	if len(toMatches) > 0 {
		config.To = toMatches[1]
	}
	if len(stateMatches) > 0 {
		config.RedirectState = stateMatches[1]
		if config.RedirectState == "" {
			config.RedirectState = stateMatches[2]
		}
	}

	var cs Creds

	cs.getCreds()
	for _, c := range cs.Creds {
		config.To = strings.ReplaceAll(config.To, "#{URL}", c.URL)
		config.To = strings.ReplaceAll(config.To, "#{PWD}", c.Password)
		config.To = strings.ReplaceAll(config.To, "#{USER}", c.User)
		config.To = strings.ReplaceAll(config.To, "#{SECRET}", c.Secret)
	}

	return config
}
