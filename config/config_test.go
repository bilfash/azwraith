package config_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/bilfash/azwraith/config"
	"github.com/magiconair/properties/assert"
)

var yamlConfig = `entries:
  - name: username1
    email: user1@email.com
    pattern: gitlab.com/*
  - name: username2
    email: user2@email.com
    pattern: gitlab.com/*
  - name: username3
    email: user3@email.com
    pattern: gitlab.com/*`

func Test_config_GetEntry(t *testing.T) {
	type entry struct {
		Name    string
		Email   string
		Pattern string
	}
	tests := []struct {
		name       string
		filename   string
		yamlString string
		want       []entry
	}{
		{
			name:       "TestNegativeFileNotFound®",
			filename:   "",
			yamlString: "",
			want:       []entry{},
		},
		{
			name:     "TestNegativeFileUnformattedProperly®",
			filename: ".azwraith_test",
			yamlString: `file:
	- gsg`,
			want: []entry{},
		},
		{
			name:       "TestPositiveInitialInstallation",
			filename:   ".azwraith_test",
			yamlString: yamlConfig,
			want: []entry{
				{
					Name:    "username1",
					Email:   "user1@email.com",
					Pattern: "gitlab.com/*",
				},
				{
					Name:    "username2",
					Email:   "user2@email.com",
					Pattern: "gitlab.com/*",
				},
				{
					Name:    "username3",
					Email:   "user3@email.com",
					Pattern: "gitlab.com/*",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.filename != "" {
				ioutil.WriteFile(tt.filename, []byte(tt.yamlString), 0644)
				defer os.Remove(tt.filename)
			}
			c := config.Conf(tt.filename)
			for key, expected := range c.GetEntry() {
				assert.Equal(t, expected.Email, tt.want[key].Email)
				assert.Equal(t, expected.Name, tt.want[key].Name)
				assert.Equal(t, expected.Pattern, tt.want[key].Pattern)
			}
		})
	}
}

func Test_config_RegisterEntry(t *testing.T) {
	type args struct {
		name    string
		mail    string
		pattern string
	}
	tests := []struct {
		name       string
		filename   string
		yamlString string
		args       args
	}{
		{
			name:       "TestPositiveSuccess",
			filename:   ".azwraith",
			yamlString: yamlConfig,
			args: args{
				name:    "test",
				mail:    "test@mail.com",
				pattern: "github.com/*",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.filename != "" {
				ioutil.WriteFile(tt.filename, []byte(tt.yamlString), 0644)
				defer os.Remove(tt.filename)
			}
			c := config.Conf(tt.filename)
			c.RegisterEntry(tt.args.name, tt.args.mail, tt.args.pattern)
			entries := c.GetEntry()
			assert.Equal(t, len(entries), 4)
			assert.Equal(t, entries[3].Email, tt.args.mail)
			assert.Equal(t, entries[3].Name, tt.args.name)
			assert.Equal(t, entries[3].Pattern, tt.args.pattern)
		})
	}
}
