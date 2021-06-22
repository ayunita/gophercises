package cmd

import (
	"fmt"
	"gophercises/task/db"
	"strconv"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete an existing task",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the arguments:", arg)
			} else {
				ids = append(ids, id)
			}
		}

		h := &db.Handler{}
		if err := h.OpenDB(); err != nil {
			fmt.Println("Cannot open the database. Error:", err)
		}

		tasks, err := h.List(bucketName, false)
		if err != nil {
			fmt.Println("Cannot open the tasks. Error:", err)
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number: ", id)
				continue
			}
			task := tasks[id-1]
			err := h.Delete(bucketName, task.Key)
			if err != nil {
				fmt.Println("Cannot mark the task as completed. Error:", err)
			} else {
				fmt.Printf("You have deleted the \"%s\" task.\n", task.Value)
			}
		}

		h.CloseDB()
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
