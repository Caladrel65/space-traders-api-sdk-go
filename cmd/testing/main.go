package main

import (
	"fmt"
	"log"
	"space-traders-api-sdk-go/pkg/agent"
	"space-traders-api-sdk-go/pkg/client"
	"space-traders-api-sdk-go/pkg/game"
	"space-traders-api-sdk-go/pkg/navigation"
	"space-traders-api-sdk-go/pkg/ship"
	"strings"
)

func main() {
	var agentData *agent.RegisterData

	agentData, err := agent.LoadAgent()
	if err != nil {
		fmt.Println("No agent data found, registering a new agent...")
		client := client.NewClient("")
		agentData, err = agent.Register(client)
		if err != nil {
			log.Fatalf("Error registering agent: %s", err.Error())
		}
		err = agent.SaveAgent(agentData)
		if err != nil {
			log.Fatalf("Error saving agent data: %s", err.Error())
		}
	}

	client := client.NewClient(agentData.Token)

	if !agentData.Contract.Terms.Accepted {
		fmt.Println("Accepting the first contract...")
		acceptContractResponse, err := agent.AcceptContract(client, agentData.Contract.Id)
		if err != nil {
			log.Fatalf("Error accepting contract: %s", err.Error())
		}
		agentData.Agent = acceptContractResponse.Data.Agent
		agentData.Contract = acceptContractResponse.Data.Contract
		err = agent.SaveAgent(agentData)
		if err != nil {
			log.Fatalf("Error saving agent data: %s", err.Error())
		}
	}

	if agentData.Agent.ShipCount == 0 {
		fmt.Println("Purchasing a mining drone...")
		headquarters := agentData.Agent.Headquarters
		systemSymbol := strings.Split(headquarters, "-")[0] + "-" + strings.Split(headquarters, "-")[1]
		waypoints, err := navigation.GetSystemWaypoints(client, systemSymbol)
		if err != nil {
			log.Fatalf("Error getting system waypoints: %s", err.Error())
		}

		var shipyardWaypoint navigation.Waypoint
		for _, waypoint := range waypoints {
			for _, trait := range waypoint.Traits {
				if trait.Symbol == "SHIPYARD" {
					shipyardWaypoint = waypoint
					break
				}
			}
		}

		if shipyardWaypoint.Symbol == "" {
			log.Fatalf("No shipyard found in system %s", systemSymbol)
		}

		purchaseShipResponse, err := ship.PurchaseShip(client, "SHIP_MINING_DRONE", shipyardWaypoint.Symbol)
		if err != nil {
			log.Fatalf("Error purchasing ship: %s", err.Error())
		}

		agentData.Agent = purchaseShipResponse.Data.Agent
		agentData.Agent.Ships = append(agentData.Agent.Ships, purchaseShipResponse.Data.Ship)

		err = agent.SaveAgent(agentData)
		if err != nil {
			log.Fatalf("Error saving agent data: %s", err.Error())
		}
	}

	state := game.NewState(client, &agentData.Agent)
	state.Run()
}
