package runtime

import "strings"

func GenerateSlug(name string) (result string) {
	allowSymbols := "ABCDEFGHJKLMNOPRSTUVZXQWabcdefghijklmnoprstuvzxqw0123456789- "
	buff := new(strings.Builder)
	for _, v := range name {
		if strings.Contains(allowSymbols, string(v)) {
			buff.WriteRune(v)
		}
	}
	result = strings.ToLower(strings.ReplaceAll(buff.String(), " ", "-"))
	return
}
