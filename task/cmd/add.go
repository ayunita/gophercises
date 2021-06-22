package cmd

import (
	"fmt"
	"gophercises/task/db"

	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		h := &db.Handler{}
		if err := h.OpenDB(); err != nil {
			fmt.Println("Cannot open the database. Error:", err)
		}
		if err := h.InitBucket(bucketName); err != nil {
			fmt.Println("Invalid storage. Error:", err)
		}
		if err := h.Write(bucketName, task); err != nil {
			fmt.Println("Cannot add a new task. Error:", err)
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)

		h.CloseDB()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
