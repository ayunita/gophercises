package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	bucketName = "MyTask"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
