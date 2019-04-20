package config_test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bilfash/azwraith/config"
)

var yamlConfig = `entries:
  - name: username1
    email: user1@email.com
    pattern: gitlab1.com/*
  - name: username2
    email: user2@email.com
    pattern: gitlab2.com/*
  - name: username3
    email: user3@email.com
    pattern: gitlab3.com/*`

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
					Pattern: "gitlab1.com/*",
				},
				{
					Name:    "username2",
					Email:   "user2@email.com",
					Pattern: "gitlab2.com/*",
				},
				{
					Name:    "username3",
					Email:   "user3@email.com",
					Pattern: "gitlab3.com/*",
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

func Test_config_DeleteEntry(t *testing.T) {
	type args struct {
		index int
	}
	type fields struct {
		checkField bool
		name       string
		email      string
		pattern    string
		length     int
	}
	tests := []struct {
		name       string
		filename   string
		yamlString string
		args       args
		want       fields
	}{
		{
			name:       "TestPositiveSuccessRemoveIndex0",
			filename:   ".azwraith_test_delete",
			yamlString: yamlConfig,
			args: args{
				index: 0,
			},
			want: fields{
				checkField: true,
				name:       "username1",
				email:      "user1@email.com",
				pattern:    "gitlab.com/*",
				length:     2,
			},
		},
		{
			name:       "TestPositiveSuccessRemoveIndex1",
			filename:   ".azwraith",
			yamlString: yamlConfig,
			args: args{
				index: 1,
			},
			want: fields{
				checkField: true,
				name:       "username2",
				email:      "user2@email.com",
				pattern:    "gitlab.com/*",
				length:     2,
			},
		},
		{
			name:       "TestPositiveSuccessRemoveIndex2",
			filename:   ".azwraith",
			yamlString: yamlConfig,
			args: args{
				index: 1,
			},
			want: fields{
				checkField: false,
				length:     2,
			},
		},
		{
			name:       "TestNegativeIndexOutOfBound",
			filename:   ".azwraith",
			yamlString: yamlConfig,
			args: args{
				index: 99,
			},
			want: fields{
				checkField: false,
				length:     3,
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
			c.DeleteEntry(tt.args.index)
			entries := c.GetEntry()
			assert.Equal(t, tt.want.length, len(entries))
			if tt.want.checkField {
				assert.NotEqual(t, tt.want.email, entries[tt.args.index].Email)
				assert.NotEqual(t, tt.want.name, entries[tt.args.index].Name)
				assert.NotEqual(t, tt.want.pattern, entries[tt.args.index].Pattern)
			}
		})
	}
}
