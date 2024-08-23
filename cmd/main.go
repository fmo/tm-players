package cmd

import (
	pb "github.com/fmo/football-proto/golang/player"
	"github.com/fmo/tm-players/internal/kafka"
	"github.com/fmo/tm-players/internal/maps"
	"github.com/fmo/tm-players/internal/rapidapi"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var teamId int
var season int

var log = logrus.New()

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
		var players []*pb.Player

		log.Info("starting player command")

		r := rapidapi.NewPlayersApi(log)
		rapidPlayers := r.GetPlayers(season, teamId)

		m := maps.NewMapPlayers(log)
		m.MapPlayers(rapidPlayers, &players, teamId)

		publisher := kafka.NewPublisher(log, os.Getenv("KAFKA_TOPIC_PLAYERS"))
		publisher.Publish(players)
	},
}
