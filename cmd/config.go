package cmd

import (
	"bufio"
	"fmt"
	"github.com/bilfash/azwraith/util"
	"github.com/spf13/cobra"
	"os"
	"strconv"
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
	configCmd.AddCommand(deleteConfigCmd)
}

var getConfigCmd = &cobra.Command{
	Use:   "get",
	Short: "Get all config",
	Run: func(cmd *cobra.Command, args []string) {
		printConfig()
	},
}

var addConfigCmd = &cobra.Command{
	Use:   "add",
	Short: "Add config",
	Run: func(cmd *cobra.Command, args []string) {
		name := readLine("git user.name")
		email := readLine("git user.email")
		pattern := readLine("git remote pattern")

		conf := getConfig()
		conf.RegisterEntry(name, email, pattern)
	},
}

var deleteConfigCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete config given ID",
	Run: func(cmd *cobra.Command, args []string) {
		idxString := readLine("config ID")
		idx, _ := strconv.Atoi(idxString)

		conf := getConfig()
		conf.DeleteEntry(idx)
	},
}

func readLine(fieldName string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter %s: ", fieldName)
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}

func printConfig() {
	conf := getConfig()
	entries := conf.GetEntry()
	header := []string{"ID", "user.name", "user.email", "pattern"}
	rows := make([][]string, 0)
	for key, entry := range entries {
		rows = append(rows, []string{strconv.Itoa(key), entry.Name, entry.Email, entry.Pattern})
	}
	util.PrintTable(header, rows)
}
