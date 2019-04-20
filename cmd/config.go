package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/spf13/cobra"
)

func init() {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Run azwraith config related command",
	}

	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(getConfigCmd)
}

var getConfigCmd = &cobra.Command{
	Use:   "get",
	Short: "Get all config",
	Run: func(cmd *cobra.Command, args []string) {
		conf := getConfig()
		entries := conf.GetEntry()
		for _, entry := range entries {
			mapEntry := structs.Map(entry)
			formattedValue, err := json.MarshalIndent(mapEntry, "", "  ")
			if err != nil {
				fmt.Println("Config is not well formatted: ", err)
			}
			fmt.Println(string(formattedValue))
		}
	},
}
