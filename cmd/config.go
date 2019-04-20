package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Run azwraith config related command",
	}
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(getConfigCmd)
	configCmd.AddCommand(addConfigCmd)
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

var addConfigCmd = &cobra.Command{
	Use:   "add",
	Short: "Get all config",
	Run: func(cmd *cobra.Command, args []string) {
		name := readLine("user.name")
		email := readLine("user.email")
		pattern := readLine("pattern")

		conf := getConfig()
		conf.RegisterEntry(name, email, pattern)
	},
}

func readLine(fieldName string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter git %s: ", fieldName)
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}
