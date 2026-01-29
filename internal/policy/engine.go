package policy

import (
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
)

type Rule struct {
	ID      string `yaml:"id"`
	Pattern string `yaml:"pattern"`
	Score   int    `yaml:"score"`
	Action  string `yaml:"action"`
}

type Policy struct {
	Rules []Rule `yaml:"rules"`
}

var active Policy

func Load(path string) error {

	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, &active)
}

func Evaluate(text string) (string, int, string) {

	for _, r := range active.Rules {

		re := regexp.MustCompile(r.Pattern)

		if re.MatchString(text) {
			return r.ID, r.Score, r.Action
		}
	}

	return "", 0, "allow"
}
