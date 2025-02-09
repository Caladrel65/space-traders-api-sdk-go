package agent

import (
	"space-traders-api-sdk-go/pkg/account"
	"space-traders-api-sdk-go/pkg/errors"
	"space-traders-api-sdk-go/pkg/factions"
	"space-traders-api-sdk-go/pkg/ship"

	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// AgentSymbol contains the symbol for the single agent being used in the current context.
const AgentSymbol = "CALADREL"

// registerUrl is the URL of the registration endpoint.
var registerUrl = account.BaseUrl + "register"

// agentUrl is the URL to get your agent details.
var agentDetailsUrl = account.BaseUrl + "my/agent"

type RegisterResponse struct {
	Data RegisterData `json:"data"`
}

type RegisterData struct {
	Token    string                  `json:"token"`
	Agent    Agent                   `json:"agent"`
	Contract Contract                `json:"contract"`
	Faction  factions.FactionDetails `json:"faction"`
	Ship     ship.Ship               `json:"ship"`
}

type Agent struct {
	AccountId       string           `json:"accountId"`
	Symbol          string           `json:"symbol"`
	Headquarters    string           `json:"headquarters"`
	Credits         int              `json:"credits"`
	StartingFaction factions.Faction `json:"startingFaction"`
	ShipCount       int              `json:"shipCount"`
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
	Fulfilled        bool            `json:"fulfilled"`
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

func RegisterAgent(symbol string, faction factions.Faction) string {
	// Create the register request JSON.
	registerJSON, _ := json.Marshal(map[string]string{
		"symbol":  symbol,
		"faction": string(faction),
	})

	// Marshal the JSON body into bytes.
	registerBody := bytes.NewBuffer(registerJSON)

	// Build the register request.
	registerReq, _ := http.NewRequest(http.MethodPost, registerUrl, registerBody)

	// Add the content-type header.
	registerReq.Header.Add("content-type", "application/json")

	// Send ("Do") the request and receive the response.
	registerResp, err := http.DefaultClient.Do(registerReq)
	if err != nil {
		// This error would mean we had an issue in talking to the server, not that the server returned an error.
		log.Fatalf("An error occurred during agent registration: %s", err.Error())
	}

	defer registerResp.Body.Close()

	// Read the body from the response.
	registerRespBody, err := io.ReadAll(registerResp.Body)
	if err != nil {
		log.Fatalf("An error occurred reading the agent registration response body: %s", err.Error())
	}

	var resp *RegisterResponse

	err = json.Unmarshal(registerRespBody, resp)
	if err != nil {
		// TODO: Maybe we found an error here?
		log.Fatalf("An error occurred while unmarshalling the response body: %s", err.Error())
	}

	// Convert the body to a printable string.
	registerRespString := string(registerRespBody)

	return registerRespString
}

func GetAgentDetails() string {
	// Build the get details request.
	agentReq, _ := http.NewRequest(http.MethodGet, agentDetailsUrl, nil)

	// Add the headers.
	account.AddAuthorization(agentReq)

	// Send the request and receive the response.
	agentResp, err := http.DefaultClient.Do(agentReq)
	if err != nil {
		log.Fatalf("An error occurred during agent details: %s", err.Error())
	}

	defer agentResp.Body.Close()

	// Read the body from the response.
	agentRespBody, err := io.ReadAll(agentResp.Body)
	if err != nil {
		log.Fatalf("An error occurred reading the agent details response body: %s", err.Error())
	}

	println(agentRespBody)
	var resp *AgentDetailsResponse

	err = json.Unmarshal(agentRespBody, resp)
	if err != nil {
		log.Fatalf("An error occurred while unmarshalling the response body: %s", err.Error())
	}

	// Convert the body to a printable string.
	agentRespString := string(agentRespBody)

	return agentRespString
}

type AgentDetailsResponse struct {
	Data  AgentDetailsData `json:"data"`
	Error errors.APIError  `json:"error"`
}

type AgentDetailsData struct {
	AccountId       string           `json:"accountId"`
	Symbol          string           `json:"symbol"`
	Headquarters    string           `json:"headquarters"`
	Credits         int              `json:"credits"`
	StartingFaction factions.Faction `json:"startingFaction"`
	ShipCount       int              `json:"shipCount"`
}
