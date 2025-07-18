package navigation

import (
	"fmt"
	"space-traders-api-sdk-go/pkg/client"
)

type Status string

const (
	Status_DOCKED = "DOCKED"
	// TODO: Other statuses
)

type FlightMode string

const (
	// FlightMode_CRUISE - Cruise flight mode is the default mode for all ships. It consumes fuel at a normal rate and travels at a normal speed.
	FlightMode_CRUISE FlightMode = "CRUISE"
	// FlightMode_BURN - Burn flight mode consumes fuel at a faster rate and travels at a faster speed.
	FlightMode_BURN FlightMode = "BURN"
	// FlightMode_DRIFT - Drift flight mode consumes the least fuel and travels at a much slower speed. Drift mode is useful when your ship has run out of fuel and you need to conserve what little fuel you have left.
	FlightMode_DRIFT FlightMode = "DRIFT"
	// FlightMode_STEALTH - Stealth flight mode runs with systems at a minimum, making it difficult to detect. It consumes fuel at a normal rate but travels at a reduced speed.
	FlightMode_STEALTH FlightMode = "STEALTH"
)

type Nav struct {
	SystemSymbol   string `json:"systemSymbol"`
	WaypointSymbol string `json:"waypointSymbol"`
	Route          Route  `json:"route"`
	Status         string `json:"status"`
	FlightMode     string `json:"flightMode"`
}

type Route struct {
	Origin        Location `json:"origin"`
	Destination   Location `json:"destination"`
	Arrival       string   `json:"arrival"` // TODO: These are timestamps, can we maybe parse them a certain way? Idk. "2025-01-13T03:31:22.752Z" ex.
	DepartureTime string   `json:"departureTime"`
}

type Location struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type Waypoint struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
	Orbitals     []struct {
		Symbol string `json:"symbol"`
	} `json:"orbitals"`
	Traits []struct {
		Symbol      string `json:"symbol"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"traits"`
}

func GetSystemWaypoints(client *client.Client, systemSymbol string) ([]Waypoint, error) {
	var resp struct {
		Data []Waypoint `json:"data"`
	}
	err := client.Get(fmt.Sprintf("/systems/%s/waypoints", systemSymbol), &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func FindWaypointByType(client *client.Client, systemSymbol string, waypointType string) (*Waypoint, error) {
	waypoints, err := GetSystemWaypoints(client, systemSymbol)
	if err != nil {
		return nil, err
	}

	for i := range waypoints {
		if waypoints[i].Type == waypointType {
			return &waypoints[i], nil
		}
	}

	return nil, fmt.Errorf("no waypoint with type %s found in system %s", waypointType, systemSymbol)
}
