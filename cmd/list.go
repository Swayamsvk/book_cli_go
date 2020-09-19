package cmd

import (
	"fmt"
	"os"

	"github.com/Swayamsvk/book-module/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows the list of books",

	Run: func(cmd *cobra.Command, args []string) {
		books, err := db.AllBooks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if len(books) == 0 {
			fmt.Println("You have no books please add some")
			return
		}
		fmt.Println("You have these books:")
		for i, book := range books {
			fmt.Printf("%d.%s\n", i+1, book.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
