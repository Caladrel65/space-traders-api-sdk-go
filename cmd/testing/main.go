package main

import (
	"space-traders-api-sdk-go/pkg/account"
	"space-traders-api-sdk-go/pkg/agent"
	"space-traders-api-sdk-go/pkg/factions"
)

func main() {
	if account.Token == "" {
		account.Token = agent.RegisterAgent(agent.AgentSymbol, factions.Faction_COSMIC)
		println("Token not found, generated a new agent")
		println("Token:", account.Token)
	} else {
		println("Token found")
		println(account.Token)
	}

	println()
	println("Agent details:")
	println(agent.GetAgentDetails())

	// check agent data
	// view starting location

	// view contracts
	// accept contract

	// find shipyard
	// view available ships
	// purchase ship (mining drone)

	// find nearby engineered asteroid
	// send ship to orbit
	// fly to the asteroid
	// dock ship
	// refuel ship
	// orbit asteroid again
	// extract ores and minerals

	// view market data
	// TODO: If market does not buy your things, you need to go to another market.
	// list ship cargo
	// dock ship
	// sell goods (not needed for the contract)

	// navigate to delivery waypoint
	// deliver contract goods
	// fulfill contract (once all goods have been delivered, will take multiple trips)
}
