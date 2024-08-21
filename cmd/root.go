package cmd

import (
	"fmt"
	"github.com/HoseinAsadolahi/github-user-activity/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

var pageNumber = 1

var rootCmd = &cobra.Command{
	Use:   "github-activity <username>",
	Short: "Fetch GitHub activity for a specific user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if pageNumber == 0 {
			pageNumber = 1
		}
		if pageNumber < 1 || pageNumber > 10 {
			fmt.Println("page number must be between 1 and 10")
		}
		username := args[0]
		fmt.Println(utils.InfoStyle.Render(fmt.Sprintf("Fetching data for user: %s, page: %d", username, pageNumber)))
		utils.DisplayInfo(username, pageNumber-1)
	},
}

func init() {
	rootCmd.Flags().IntVarP(&pageNumber, "page", "p", 0, "Page number to fetch (default is 0)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
