package rgxmatcher

import "regexp"

type (
	Matcher interface {
		IsMatch(v string) bool
	}
	matcher struct {
		regexp *regexp.Regexp
	}
)

func NewMatcher(exprs string) (Matcher, error) {
	rgxp, err := regexp.Compile(exprs)
	if err != nil {
		return nil, err
	}
	return &matcher{
		regexp: rgxp,
	}, nil
}

func (m *matcher) IsMatch(v string) bool {
	return m.regexp.MatchString(v)
}
