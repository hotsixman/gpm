package util

import (
	"fmt"
	"geep/module/types"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func ListTable(listResultMessage types.ListResultMessage) *table.Table {
	// Define styles
	re := lipgloss.NewRenderer(os.Stdout)
	headerStyle := re.NewStyle().Foreground(lipgloss.Color("5")).Bold(true).Align(lipgloss.Center)
	cellStyle := re.NewStyle().Padding(0, 1)
	borderStyle := re.NewStyle().Foreground(lipgloss.Color("240"))

	// Prepare rows
	var rows [][]string
	for _, elem := range listResultMessage.List {
		rows = append(rows, []string{
			elem.Name,
			elem.Status,
			fmt.Sprintf("%d", elem.Recovered),
			fmt.Sprintf("%.2f%%", elem.CPUPercent),
			fmt.Sprintf("%.2f MB", elem.Mem),
			elem.Pid,
		})
	}

	// Create and render table
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(borderStyle).
		Headers("Name", "Status", "Recovered", "CPU", "Memory", "Pid").
		Rows(rows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == -1 {
				return headerStyle
			}
			return cellStyle
		})

	return t
}
