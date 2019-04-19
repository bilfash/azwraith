package rgxmatcher_test

import (
	"github.com/bilfash/azwraith/rgxmatcher"
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_matcher_IsMatch(t *testing.T) {
	type expected struct {
		matched bool
		err     bool
	}
	tests := []struct {
		name string
		exp  string
		args string
		want expected
	}{
		{
			name: "TestNegativeGivenErrorMatcher",
			exp:  "a\xc5z",
			args: "teststring",
			want: expected{
				matched: false,
				err:     true,
			},
		},
		{
			name: "TestPositiveGivenNoPattern",
			exp:  "github.com",
			args: "github.com/bilfash/azwraith",
			want: expected{
				matched: true,
				err:     false,
			},
		},
		{
			name: "TestPositiveGivenSimplePattern",
			exp:  "git*",
			args: "github.com/bilfash/azwraith",
			want: expected{
				matched: true,
				err:     false,
			},
		},
		{
			name: "TestPositiveGivenHttpsExpression",
			exp:  "github.com*",
			args: "https://github.com/bilfash/azwraith",
			want: expected{
				matched: true,
				err:     false,
			},
		},
		{
			name: "TestPositiveGivenGitExpression",
			exp:  "github.com*",
			args: "git@github.com:bilfash/azwraith",
			want: expected{
				matched: true,
				err:     false,
			},
		},
		{
			name: "TestPositiveGivenHttpsExpressionNotMatched",
			exp:  "github.com*",
			args: "https://gitlab.com/bilfash/azwraith",
			want: expected{
				matched: false,
				err:     false,
			},
		},
		{
			name: "TestPositiveGivenGitExpressionNotMatched",
			exp:  "github.com*",
			args: "git@gitlab.com:bilfash/azwraith",
			want: expected{
				matched: false,
				err:     false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := rgxmatcher.NewMatcher(tt.exp)
			assert.Equal(t, err != nil, tt.want.err)
			if m != nil {
				got := m.IsMatch(tt.args)
				assert.Equal(t, got, tt.want.matched)
			}
		})
	}
}
