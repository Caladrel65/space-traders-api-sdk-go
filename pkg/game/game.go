package game

import (
	"space-traders-api-sdk-go/pkg/agent"
	"space-traders-api-sdk-go/pkg/client"
)

type State struct {
	Client *client.Client
	Agent  *agent.Agent
}

func NewState(client *client.Client, agent *agent.Agent) *State {
	return &State{
		Client: client,
		Agent:  agent,
	}
}

func (s *State) Run() {
	for i := range s.Agent.Ships {
		if i == 0 {
			s.Agent.Ships[i].Role = "hauler"
		} else {
			s.Agent.Ships[i].Role = "drone"
		}
	}

	hauler := &s.Agent.Ships[0]
	miningWaypoint, err := navigation.FindWaypointByType(s.Client, hauler.Nav.SystemSymbol, "ASTEROID_FIELD")
	if err != nil {
		log.Fatalf("Error finding mining waypoint: %s", err.Error())
	}

	if hauler.Nav.WaypointSymbol != miningWaypoint.Symbol {
		fmt.Printf("Navigating hauler to %s...\n", miningWaypoint.Symbol)
		navigateResp, err := ship.Navigate(s.Client, hauler.Symbol, miningWaypoint.Symbol)
		if err != nil {
			log.Fatalf("Error navigating hauler: %s", err.Error())
		}
		hauler.Nav = navigateResp.Data.Nav
	}

	for i := range s.Agent.Ships {
		if s.Agent.Ships[i].Role == "drone" {
			drone := &s.Agent.Ships[i]
			if drone.Nav.WaypointSymbol != miningWaypoint.Symbol {
				fmt.Printf("Navigating drone %s to %s...\n", drone.Symbol, miningWaypoint.Symbol)
				navigateResp, err := ship.Navigate(s.Client, drone.Symbol, miningWaypoint.Symbol)
				if err != nil {
					log.Printf("Error navigating drone %s: %s", drone.Symbol, err.Error())
				}
				drone.Nav = navigateResp.Data.Nav
			}
		}
	}

	for {
		for i := range s.Agent.Ships {
			if s.Agent.Ships[i].Role == "drone" {
				drone := &s.Agent.Ships[i]

				if drone.Nav.Status != navigation.Status_IN_ORBIT {
					fmt.Printf("Putting drone %s into orbit...\n", drone.Symbol)
					orbitResp, err := ship.Orbit(s.Client, drone.Symbol)
					if err != nil {
						log.Printf("Error putting drone %s into orbit: %s", drone.Symbol, err.Error())
					} else {
						drone.Nav = orbitResp.Data.Nav
					}
				}

				fmt.Printf("Extracting resources with drone %s...\n", drone.Symbol)
				extractResp, err := ship.Extract(s.Client, drone.Symbol)
				if err != nil {
					log.Printf("Error extracting resources with drone %s: %s", drone.Symbol, err.Error())
				} else {
					drone.Cargo = extractResp.Data.Cargo
				}

				if drone.Nav.Status != navigation.Status_DOCKED {
					fmt.Printf("Docking drone %s...\n", drone.Symbol)
					dockResp, err := ship.Dock(s.Client, drone.Symbol)
					if err != nil {
						log.Printf("Error docking drone %s: %s", drone.Symbol, err.Error())
					} else {
						drone.Nav = dockResp.Data.Nav
					}
				}

				if drone.Fuel.Current < drone.Fuel.Capacity/4 {
					fmt.Printf("Refueling drone %s...\n", drone.Symbol)
					refuelResp, err := ship.Refuel(s.Client, drone.Symbol)
					if err != nil {
						log.Printf("Error refueling drone %s: %s", drone.Symbol, err.Error())
					} else {
						drone.Fuel = refuelResp.Data.Fuel
						s.Agent.Credits = refuelResp.Data.Agent.Credits
					}
				}
			}
		}
	}
}
