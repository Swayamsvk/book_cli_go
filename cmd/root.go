package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "book",
	Short: "Book is a book listing cli manager",
}
