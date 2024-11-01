package cmd

import (
	"fmt"

	"github.com/not7cd/lbrd/models"
	"github.com/spf13/cobra"
)

var listLeaderBoardCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Leaderboard 1")

		db, err := models.ConnectToSQLiteDB()
		if err != nil {
			fmt.Println("Error connecting to SQLite:", err)
			return
		}
		defer db.Close()

		scores, err := models.GetScores(db, 1)
		if err != nil {
			fmt.Println("Failed getting score for LB 1")
			return
		}

		for i, score := range scores {
			fmt.Printf("%d\t%s\t%d\n", i+1, score.Player, score.Value)
		}
	},
}
