TM_PLAYERS_BINARY=tmPlayersApp

tm_players:
	@echo "Building binary..."
	go build -o ${TM_PLAYERS_BINARY} ./cmd/
	@echo "Done!"

tm_players_amd:
	@echo "Building binary..."
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${TM_PLAYERS_BINARY} ./cmd/
	@echo "Done!"
