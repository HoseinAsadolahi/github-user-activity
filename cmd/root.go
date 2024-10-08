package cmd

import (
	"fmt"
	"github.com/HoseinAsadolahi/github-user-activity/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

var pageNumber int

var rootCmd = &cobra.Command{
	Use:   "github-activity <username>",
	Short: "Fetch GitHub activity for a specific user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if pageNumber < 1 || pageNumber > 10 {
			fmt.Println(utils.ErrorStyle.Render("page number must be between 1 and 10"))
			return
		}
		username := args[0]
		fmt.Println(utils.InfoStyle.Render("Fetching data ..."))
		utils.DisplayInfo(username, pageNumber-1)
	},
}

func init() {
	rootCmd.Flags().IntVarP(&pageNumber, "page", "p", 1, "Page number to fetch (default is 0)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
