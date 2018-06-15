package ids

import (
	"regexp"
	"strings"
)

func ReadIDUsingStrings(mqttID string) (id string, v3 bool) {
	v3ID := strings.TrimPrefix(mqttID, "v3/")
	if v3ID == mqttID {
		return mqttID, false
	}
	return v3ID, true
}

var idRegex = regexp.MustCompile("v3\\/(.*)")

func ReadIDUsingRegexp(mqttID string) (id string, v3 bool) {
	results := idRegex.FindAllStringSubmatch(mqttID, -1)
	if len(results) == 0 {
		return mqttID, false
	}

	return results[0][1], true
}
