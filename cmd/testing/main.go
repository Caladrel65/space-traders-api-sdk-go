package main

import (
	"space-traders-api-sdk-go/pkg/agent"
)

var token string = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGlmaWVyIjoiQ0FMQURSRUwiLCJ2ZXJzaW9uIjoidjIuMi4wIiwicmVzZXRfZGF0ZSI6IjIwMjQtMTItMjMiLCJpYXQiOjE3MzY2MjE0NDgsInN1YiI6ImFnZW50LXRva2VuIn0.nYCFwnr9V0x7d-cToHQ1R5DN_kJvCfxsH-Fm_jrho0T0cP4CB0qqroXRL63NOJG_Qdz35cCLq1ibtYcxVCu9LGwjKQwK7CaWB8PeLTTx9yUYroIC-CqrBD8DRYcmiquj_-oDttaiyRM4n8uWFvKbxNKOVD6tf--64V0EEwKJ6LkQkY6jos_eCsdjG11j2_dAvM7zUFL5PRJhOOFqbIWQhDfUaxMNprxtZvQj_MepT12CoWqA3z2nEbaAZgGhXsVFMGh7lcQXuK5vVSaFwFeJHjvwvn1tg3yeh6VahNXaFeT6wz-ohGYuagRGdQH7-sQQ-2odaStSpCloD-P1lS_EeA"

func main() {

	if token == "" {
		token = agent.RegisterAgent(agent.AgentSymbol, agent.Factions.Cosmic)
	}
	// if token == "" {
	// 	// TODO: Set up registerAgent to use token from response.
	// 	token = registerAgent()
	// }

	agent.RegisterAgent(agent.AgentSymbol, agent.Factions.Cosmic)

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
