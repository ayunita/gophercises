package cmd

import (
	"fmt"
	"gophercises/task/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		h := &db.Handler{}

		if err := h.OpenDB(); err != nil {
			fmt.Println("Cannot open the database. Error:", err)
		}

		tasks, err := h.List(bucketName)
		if err != nil {
			fmt.Println("Cannot open the tasks. Error:", err)
		}

		for i, t := range tasks {
			fmt.Println(i+1, t.Value)
		}

		h.CloseDB()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
