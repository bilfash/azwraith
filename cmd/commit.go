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
	cli     cliintrepeter.CliInterpreter
)

func init() {
	cli = cliintrepeter.NewCliInterpreter()
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringVarP(&remote, "remote", "r", "origin", "specify git remote to push")
}

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit your code",
	Args:  cobra.ExactArgs(1),
	Long: "Commit command will get remote url and match it with azwraith config, " +
		"after getting the right config azwraith will commit your code using credential from matched config",
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
			fmt.Println("Config matches not found\nCode not commited!!")
			return
		}
		fmt.Println("Remote          :", remote)
		fmt.Println("Set username to :", username)
		fmt.Println("Set email to    :", email)
		fmt.Print(setUsername(username))
		fmt.Print(setEmail(email))
		fmt.Print(push(args[0]))
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

func push(message string) string {
	output, _ := cli.Execute(command, "commit", "-m", message)
	return output
}
