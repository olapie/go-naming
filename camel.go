package naming

import "unicode"

func ToCamel(s string, options ...Option) string {
	input := []rune(s)
	type State int
	const (
		Start State = iota
		FirstUpper
		SecondUpper
		Lower
		FirstWordEnd
		UpperWord
		UpperWordEnd
	)

	ignoreAcronym := false
	if options != nil {
		var opts Options
		for _, fn := range options {
			fn(&opts)
		}
		ignoreAcronym = opts.ignoreAcronym
	}

	state := Start
	wordStart := -1
	for i, r := range input {
		switch state {
		case Start:
			if isDelimiter(r) {
				break
			}
			wordStart = i
			if unicode.IsUpper(r) {
				state = FirstUpper
			} else {
				state = Lower
			}
		case FirstUpper:
			if isDelimiter(r) {
				state = FirstWordEnd
				return string(toCamel(input, wordStart, i))
			}
			if unicode.IsUpper(r) {
				state = SecondUpper
			} else {
				state = Lower
			}
		case SecondUpper:
			if isDelimiter(r) {
				state = FirstWordEnd
				return string(toCamel(input, wordStart, i))
			}
			if unicode.IsUpper(r) {
				state = Lower
			} else {
				state = UpperWord
			}
		case Lower:
			if isDelimiter(r) || unicode.IsUpper(r) {
				state = FirstWordEnd
				return string(toCamel(input, wordStart, i))
			}
		case UpperWord:
			if isDelimiter(r) {
				state = FirstWordEnd
				return string(toCamel(input, wordStart, i))
			}
			if unicode.IsUpper(r) {
				break
			}

			state = UpperWordEnd
			if ignoreAcronym {
				state = Lower
				break
			}

			if acronym := acronymsMap[string(input[wordStart:i-1])]; acronym != "" {
				state = FirstWordEnd
				return string(toCamel(input, wordStart, i-1))
			}
			state = Lower
		}
	}

	if wordStart < 0 {
		return ""
	}
	n := len(input)
	for i := wordStart; i < n; i++ {
		input[i] = unicode.ToLower(input[i])
	}
	return string(input)
}

func toCamel(input []rune, firstWordStart, firstWordEnd int) []rune {
	word := input[firstWordStart:firstWordEnd]
	for i := range word {
		word[i] = unicode.ToLower(word[i])
	}
	pascal := toPascal(input[firstWordEnd:])
	return append(word, pascal...)
}
