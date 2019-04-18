package config_test

import (
	"github.com/bilfash/azwraith/config"
	"io/ioutil"
	"os"
	"testing"
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
	type fields struct {
		file    string
		Entries []entry
	}
	tests := []struct {
		name       string
		filename   string
		yamlString string
		fields     fields
		want       []entry
	}{
		{
			name:       "TestNegativeFileNotFound®",
			filename:   "",
			yamlString: "",
			fields: fields{
				Entries: []entry{},
			},
			want: []entry{},
		},
		{
			name:     "TestNegativeFileUnformattedProperly®",
			filename: ".azwraith_test",
			yamlString: `file:
	- gsg`,
			fields: fields{
				Entries: []entry{},
			},
			want: []entry{},
		},
		{
			name:       "TestPositiveInitialInstallation",
			filename:   ".azwraith_test",
			yamlString: yamlConfig,
			fields: fields{
				Entries: []entry{
					{
						Name:    "username",
						Email:   "useremail@mail.com",
						Pattern: "gitlab.com",
					},
				},
			},
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
				if tt.want[key].Email != expected.Email {
					t.Errorf("config.GetEntry() Email = \"%v\", want \"%v\"", expected.Email, tt.want[key].Email)
				}
				if tt.want[key].Name != expected.Name {
					t.Errorf("config.GetEntry() Name = \"%v\", want \"%v\"", expected.Name, tt.want[key].Name)
				}
				if tt.want[key].Pattern != expected.Pattern {
					t.Errorf("config.GetEntry() Pattern = \"%v\", want \"%v\"", expected.Pattern, tt.want[key].Pattern)
				}
			}
		})
	}
}
