package policy

import (
	"os"
	"path/filepath"
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

var cache = map[string]Policy{}

func Load(tenant string) error {

	path := filepath.Join("policies", tenant+".yaml")

	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var p Policy

	err = yaml.Unmarshal(b, &p)
	if err != nil {
		return err
	}

	cache[tenant] = p

	return nil
}

func Evaluate(tenant, text string) (string, int, string) {

	p, ok := cache[tenant]
	if !ok {
		return "", 0, "allow"
	}

	for _, r := range p.Rules {

		re := regexp.MustCompile(r.Pattern)

		if re.MatchString(text) {
			return r.ID, r.Score, r.Action
		}
	}

	return "", 0, "allow"
}
