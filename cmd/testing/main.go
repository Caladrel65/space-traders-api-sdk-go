package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"

	"space-traders-api-sdk-go/pkg/client"
	"space-traders-api-sdk-go/pkg/game"
	"space-traders-api-sdk-go/pkg/models"
	"space-traders-api-sdk-go/pkg/navigation"
)

type AccountToken struct {
	Token string `json:"token"`
}

func main() {
	const archiveDir = "archived_agents"

	// Read the account token
	tokenFile, err := ioutil.ReadFile("account_token.json")
	if err != nil {
		log.Fatalf("Error reading account_token.json: %s", err.Error())
	}

	var accountToken AccountToken
	err = json.Unmarshal(tokenFile, &accountToken)
	if err != nil {
		log.Fatalf("Error unmarshaling account_token.json: %s", err.Error())
	}

	var agentData *models.RegisterData
	newAgentRegistered := false

	agentData, err = models.LoadAgent()
	if err != nil || agentData.Agent.Symbol == "" {
		if err == nil && agentData.Agent.Symbol == "" {
			// Agent data exists but is invalid, archive it
			fmt.Println("Invalid agent data found, archiving existing agent.json...")

			// Create archive directory if it doesn't exist
			if _, err := os.Stat(archiveDir); os.IsNotExist(err) {
				err = os.Mkdir(archiveDir, 0755)
				if err != nil {
					log.Fatalf("Error creating archive directory: %s", err.Error())
				}
			}

			archiveFileName := fmt.Sprintf("%s/agent_%s.json", archiveDir, uuid.New().String())
			err = os.Rename("agent.json", archiveFileName)
			if err != nil {
				log.Fatalf("Error archiving agent.json: %s", err.Error())
			}
		}

		fmt.Println("No valid agent data found, registering a new agent...")
		client := client.NewClient(accountToken.Token)
		agentData, err = models.Register(client)
		if err != nil {
			log.Fatalf("Error registering agent: %s", err.Error())
		}
		client.Token = agentData.Token
		err = models.SaveAgent(agentData)
		if err != nil {
			log.Fatalf("Error saving agent data: %s", err.Error())
		}
		newAgentRegistered = true
	}
	os.Exit(0)

	client := client.NewClient(agentData.Token)

	if newAgentRegistered {
		if !agentData.Contract.Terms.Accepted {
			fmt.Println("Accepting the first contract...")
			acceptContractResponse, err := models.AcceptContract(client, agentData.Contract.Id)
			if err != nil {
				log.Fatalf("Error accepting contract: %s", err.Error())
			}
			agentData.Agent = acceptContractResponse.Data.Agent
			agentData.Contract = acceptContractResponse.Data.Contract
			err = models.SaveAgent(agentData)
			if err != nil {
				log.Fatalf("Error saving agent data: %s", err.Error())
			}
		}

		if agentData.Agent.ShipCount == 0 {
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

			purchaseShipResponse, err := models.PurchaseShip(client, "SHIP_MINING_DRONE", shipyardWaypoint.Symbol)
			if err != nil {
				log.Fatalf("Error purchasing ship: %s", err.Error())
			}

			agentData.Agent = purchaseShipResponse.Data.Agent
			agentData.Agent.Ships = append(agentData.Agent.Ships, purchaseShipResponse.Data.Ship)

			err = models.SaveAgent(agentData)
			if err != nil {
				log.Fatalf("Error saving agent data: %s", err.Error())
			}
		}
	}

	state := game.NewState(client, &agentData.Agent)
	state.Run()
}
