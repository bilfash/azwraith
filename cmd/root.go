package cmd

import (
	"fmt"
	"github.com/bilfash/azwraith/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
)

var configFile = ".azwraith"

var rootCmd = &cobra.Command{
	Use:   "azwraith",
	Short: "Azwraith is a cli command to manage credential when pushing your changes to version control system",
}

func getConfig() config.Config {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("Error when try to get home directory : %v", err)
		os.Exit(1)
	}
	return config.Conf(fmt.Sprintf("%s/%s", home, configFile))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
