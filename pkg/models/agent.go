package models

import (
	"encoding/json"
	"fmt"
	"os"
	"space-traders-api-sdk-go/pkg/client"
	"space-traders-api-sdk-go/pkg/factions"
)

const AgentSymbol = "CALADREL"

type RegisterResponse struct {
	Data RegisterData `json:"data"`
}

type RegisterData struct {
	Token    string                  `json:"token"`
	Agent    Agent                   `json:"agent"`
	Contract Contract                `json:"contract"`
	Faction  factions.FactionDetails `json:"faction"`
	Ship     Ship               `json:"ship"`
}

type Agent struct {
	AccountId       string           `json:"accountId"`
	Symbol          string           `json:"symbol"`
	Headquarters    string           `json:"headquarters"`
	Credits         int              `json:"credits"`
	StartingFaction factions.Faction `json:"startingFaction"`
	ShipCount       int              `json:"shipCount"`
	Ships           []Ship      `json:"ships"`
}

type Contract struct {
	Id            string           `json:"id"`
	FactionSymbol factions.Faction `json:"factionSymbol"`
	Type          ContractType     `json:"type"`
	Terms         ContractTerms    `json:"terms"`
}

type ContractType string

const (
	ContractType_PROCUREMENT ContractType = "PROCUREMENT"
)

type ContractTerms struct {
	Deadling         string          `json:"deadline"`
	Payment          ContractPayment `json:"payment"`
	Deliver          []Delivery      `json:"deliver"`
	Accepted         bool            `json:"accepted"`
	Fulfilled        bool            `json:tfulfilled"`
	Expiration       string          `json:"expiration"`
	DeadlineToAccept string          `json:"deadlineToAccept"`
}

type ContractPayment struct {
	OnAccepted  int `json:"onAccepted"`
	OnFulfilled int `json:"onFulfilled"`
}

type Delivery struct {
	TradeSymbol       string `json:"tradeSymbol"`
	DestinationSymbol string `json:"destinationSymbol"`
	UnitsRequired     int    `json:"unitsRequired"`
	UnitsFulfilled    int    `json:"unitsFulfilled"`
}

func Register(client *client.Client) (*RegisterData, error) {
	registerJSON := map[string]string{
		"symbol":  AgentSymbol,
		"faction": string(factions.Faction_COSMIC),
	}

	var resp RegisterResponse
	err := client.Post("/register", registerJSON, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func GetMyAgent(client *client.Client) (*Agent, error) {
	var resp struct {
		Data Agent `json:"data"`
	}
	err := client.Get("/my/agent", &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func SaveAgent(agent *RegisterData) error {
	b, err := json.MarshalIndent(agent, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("agent.json", b, 0644)
}

func LoadAgent() (*RegisterData, error) {
	b, err := os.ReadFile("agent.json")
	if err != nil {
		return nil, err
	}

	var agent RegisterData
	err = json.Unmarshal(b, &agent)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

type AcceptContractResponse struct {
	Data struct {
		Agent    Agent    `json:"agent"`
		Contract Contract `json:"contract"`
	} `json:"data"`
}

func AcceptContract(client *client.Client, contractId string) (*AcceptContractResponse, error) {
	var resp AcceptContractResponse
	err := client.Post(fmt.Sprintf("/my/contracts/%s/accept", contractId), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

