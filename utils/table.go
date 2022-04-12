package utils

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func BuildTable(header []string, rows [][]string) {
	if len(rows) == 0 {
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.SetBorder(false)
	table.AppendBulk(rows)
	table.Render()
}
