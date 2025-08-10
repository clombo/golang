package stringUtils

import "regexp"

//Regex package n parameter:
// < 0 means to return all matches
// > 0 means return at most n matches
// = 0 means return no matches

// Return all matches of the regex in the content
func FindAllByRegex(input string, regex string) ([][]string, int) {
	re := regexp.MustCompile(regex)
	matches := re.FindAllStringSubmatch(input, -1)
	return matches, len(matches)
}
