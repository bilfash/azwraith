package util

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func PrintTable(header []string, rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, v := range rows {
		table.Append(v)
	}
	table.Render()
}
