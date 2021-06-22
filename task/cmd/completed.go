package cmd

import (
	"fmt"
	"gophercises/task/db"
	"time"

	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List all completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		h := &db.Handler{}

		if err := h.OpenDB(); err != nil {
			fmt.Println("Cannot open the database. Error:", err)
		}

		tasks, err := h.List(bucketName, true)
		if err != nil {
			fmt.Println("Cannot open the tasks. Error:", err)
		}

		fmt.Println("You have finished the following tasks today:")
		for _, t := range tasks {
			ts := t.Timestamp
			now := time.Now()
			if ts.Day() == now.Day() && ts.Month() == now.Month() && ts.Year() == now.Year() {
				// Use the layout string "2006-01-02"
				fmt.Println("[", ts.Format("02-Jan-2006"), "]", t.Value)
			}
		}

		h.CloseDB()
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)
}
