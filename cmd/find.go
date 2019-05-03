package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jace-ys/taskar/pkg/manager"
	"github.com/jace-ys/taskar/pkg/printer"
)

var findFlags manager.FilterOptions

func init() {
	findCmd.Flags().StringVar(&findFlags.Name, "name", "", "Filter tasks by name")
	findCmd.Flags().StringVar(&findFlags.Category, "tag", "", "Filter tasks by category")
	findCmd.Flags().StringVar(&findFlags.DueDate, "due", "", "Filter tasks by due date")
	rootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Display a list of filtered tasks from your to-do list",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// Find all tasks
		tasks, err := taskManager.FindAll()
		if err != nil {
			exit(err)
		}
		// Filter tasks using provided flags
		filtered := taskManager.Filter(tasks, findFlags)
		printer.PrintTable(filtered)
	},
}
