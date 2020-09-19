package cmd

import (
	"fmt"
	"strconv"

	"github.com/Swayamsvk/book-module/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "completed",
	Short: "Marks books as done",

	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)

			} else {
				ids = append(ids, id)
			}

		}
		books, err := db.AllBooks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(books) {
				fmt.Println("Invalid book number:", id)
				continue
			}
			book := books[id-1]
			err := db.DeleteBook(book.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed.Error:%s", id, err)

			} else {
				fmt.Printf("Markes \"%d\" as completed.\n", id)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
