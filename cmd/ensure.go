package cmd

import (
	"fmt"
	"github.com/bilfash/azwraith/rgxmatcher"
	"github.com/bilfash/azwraith/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ensureConfigCmd)
}

var ensureConfigCmd = &cobra.Command{
	Use:   "ensure",
	Short: "Ensure azwraith config is working as expected",
	Long: "Ensure will match remote url from command argument to current azwraith config. This will help you" +
		"to make sure your azwraith config is working as expected",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conf := getConfig()
		entries := conf.GetEntry()
		for _, entry := range entries {
			matcher, err := rgxmatcher.NewMatcher(entry.Pattern)
			if err != nil {
				fmt.Println(err)
				return
			}
			if matcher.IsMatch(args[0]) {
				header := []string{"Key", "Value"}
				rows := [][]string{
					{
						"git user.name",
						entry.Name,
					},
					{
						"git user.email",
						entry.Email,
					},
				}
				util.PrintTable(header, rows)
				return
			}
		}
		fmt.Println("Config matches not found")
	},
}
