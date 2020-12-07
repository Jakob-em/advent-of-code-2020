package utils

import "regexp"

func ExtractGroups(regex string, line string) map[string]string {
	var inputPattern = regexp.MustCompile(regex)
	groups := map[string]string{}
	matches := inputPattern.FindStringSubmatch(line)
	for i, v := range inputPattern.SubexpNames() {
		groups[v] = matches[i]
	}
	return groups
}
