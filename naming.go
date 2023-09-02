package naming

type Options struct {
	ignoreAcronym bool
}

type Option func(options *Options)

func WithoutAcronym() Option {
	return func(options *Options) {
		options.ignoreAcronym = true
	}
}

func trimLeadingDelimiters(a []rune) []rune {
LOOP:
	for i := range a {
		switch a[i] {
		case '.', ' ', '-', '\t', '_':
			continue LOOP
		default:
			return a[i:]
		}
	}
	return nil
}

func trimTrailingDelimiters(a []rune) []rune {
	n := len(a)
LOOP:
	for i := n - 1; i >= 0; i-- {
		switch a[i] {
		case '.', ' ', '-', '\t', '_':
			continue LOOP
		default:
			return a[:i+1]
		}
	}
	return nil
}

func contractDelimiters(a []rune) []rune {
	flag := false
	n := len(a)
	for i := 0; i < n; i++ {
		switch a[i] {
		case '.', ' ', '-', '\t', '_':
			if flag {
				for j := i; j < n-1; j++ {
					a[j] = a[j+1]
				}
				n--
				i--
				a = a[:n]
			} else {
				a[i] = '_'
			}
			flag = true
		default:
			flag = false
		}
	}
	return a
}

func normalize(s string) []rune {
	a := []rune(s)
	a = trimLeadingDelimiters(a)
	a = trimTrailingDelimiters(a)
	a = contractDelimiters(a)
	return a
}

func isDelimiter(r rune) bool {
	switch r {
	case '_', '.', '-', ' ', '\t':
		return true
	default:
		return false
	}
}
