package cmd

import (
	"fmt"
	"github.com/bilfash/azwraith/cliintrepeter"
	"github.com/bilfash/azwraith/rgxmatcher"
	"github.com/spf13/cobra"
)

var (
	command = "git"
	remote  string
	branch  string
	cli     cliintrepeter.CliInterpreter
)

func init() {
	cli = cliintrepeter.NewCliInterpreter()
	rootCmd.AddCommand(pushConfigCmd)
	pushConfigCmd.Flags().StringVarP(&remote, "remote", "r", "origin", "specify git remote to push")
	pushConfigCmd.Flags().StringVarP(&branch, "branch", "b", "", "specify git branch to push")
}

var pushConfigCmd = &cobra.Command{
	Use:   "push",
	Short: "Push your code",
	Long: "Push command will get remote url and match it with azwraith config, " +
		"after getting the right config azwraith will push your code",
	Run: func(cmd *cobra.Command, args []string) {
		remoteUrl := getRemoteUrl(remote)
		username := ""
		email := ""
		conf := getConfig()
		entries := conf.GetEntry()
		for _, entry := range entries {
			matcher, err := rgxmatcher.NewMatcher(entry.Pattern)
			if err != nil {
				fmt.Println(err)
				return
			}
			if matcher.IsMatch(remoteUrl) {
				username = entry.Name
				email = entry.Email
				break
			}
		}
		if username == "" || email == "" {
			fmt.Println("Config matches not found\nCode not pushed!!")
			return
		}

		fmt.Println(setUsername(username))
		fmt.Println(setEmail(email))
		fmt.Println(push())
	},
}

func getRemoteUrl(remote string) string {
	output, _ := cli.Execute(command, "config", "--get", fmt.Sprintf("remote.%s.url", remote))
	return output
}

func setUsername(username string) string {
	output, _ := cli.Execute(command, "config", "--global", "user.name", username)
	return output
}

func setEmail(email string) string {
	output, _ := cli.Execute(command, "config", "--global", "user.email", email)
	return output
}

func push() string {
	output := ""
	if branch == "" {
		output, _ = cli.Execute(command, "push", remote)
	} else {
		output, _ = cli.Execute(command, "push", remote, branch)
	}
	return output
}
