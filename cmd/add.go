package cmd

import (
	"fmt"
	"strings"

	"github.com/Swayamsvk/book-module/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a book to the main list of books",

	Run: func(cmd *cobra.Command, args []string) {
		book := strings.Join(args, " ")
		_, err := db.CreateBook(book)
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		fmt.Printf("Added %s to your books list\n", book)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
