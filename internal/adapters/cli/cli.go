package cli

import (
	"context"
	"fmt"
	"github.com/fmo/tm-players/internal/ports"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var teamId int
var season int

var log = logrus.New()

type Adapter struct {
	api ports.APIPorts
}

func NewAdapter(api ports.APIPorts) *Adapter {
	return &Adapter{
		api: api,
	}
}

func init() {
	Cmd.Flags().IntVarP(&teamId, "teamId", "t", 36, "Team Id")
	Cmd.Flags().IntVarP(&season, "season", "s", 2024, "Season")

	log.Out = os.Stdout
	log.Level = logrus.DebugLevel
}

var Cmd = &cobra.Command{
	Use:   "players",
	Short: "Get players from Transfermarkt",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func (a Adapter) Run(ctx context.Context) {
	rootCmd := &cobra.Command{
		Use:   "football-data-app",
		Short: "Football Data CLI Application",
	}

	Cmd.Run = func(cmd *cobra.Command, args []string) {
		log.Info("Starting player command with Adapter")

		err := a.api.SavePlayer(ctx, season, teamId)
		if err != nil {
			fmt.Println("error: ", err)
		}

	}

	rootCmd.AddCommand(Cmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
