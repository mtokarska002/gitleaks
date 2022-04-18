package rules

import (
	"regexp"

	"github.com/zricethezav/gitleaks/v8/config"
)

func Clojars() *config.Rule {
	// define rule
	r := config.Rule{
		Description: "Clojars API token",
		RuleID:      "clojars-api-token",
		Regex:       regexp.MustCompile(`(CLOJARS_)(?i)[a-z0-9]{60}`),
		Keywords:    []string{"clojars"},
	}

	// validate
	tps := []string{
		"clojarsAPIToken := \"CLOJARS_" + sampleAlphaNumeric60Token + "\"",
	}
	return validate(r, tps)
}