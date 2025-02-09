package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// AgentSymbol contains the symbol for the single agent being used in this program.
const AgentSymbol = "CALADREL"

const RegistrationEndpoint = "https://api.spacetraders.io/v2/register"

var token string = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGlmaWVyIjoiQ0FMQURSRUwiLCJ2ZXJzaW9uIjoidjIuMi4wIiwicmVzZXRfZGF0ZSI6IjIwMjQtMTItMjMiLCJpYXQiOjE3MzY2MjE0NDgsInN1YiI6ImFnZW50LXRva2VuIn0.nYCFwnr9V0x7d-cToHQ1R5DN_kJvCfxsH-Fm_jrho0T0cP4CB0qqroXRL63NOJG_Qdz35cCLq1ibtYcxVCu9LGwjKQwK7CaWB8PeLTTx9yUYroIC-CqrBD8DRYcmiquj_-oDttaiyRM4n8uWFvKbxNKOVD6tf--64V0EEwKJ6LkQkY6jos_eCsdjG11j2_dAvM7zUFL5PRJhOOFqbIWQhDfUaxMNprxtZvQj_MepT12CoWqA3z2nEbaAZgGhXsVFMGh7lcQXuK5vVSaFwFeJHjvwvn1tg3yeh6VahNXaFeT6wz-ohGYuagRGdQH7-sQQ-2odaStSpCloD-P1lS_EeA"

// Factions is a struct containing all of the factions in the game.
var Factions = struct {
	Cosmic   string
	Galactic string
	Solar    string
}{
	Cosmic:   "COSMIC",
	Galactic: "GALACTIC",
	Solar:    "SOLAR",
}

type registerResponse struct {
	Data registerData `json:"data"`
	/*
			{
		  "data": {
		    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGlmaWVyIjoiQ0FMQURSRUwyIiwidmVyc2lvbiI6InYyLjIuMCIsInJlc2V0X2RhdGUiOiIyMDI0LTEyLTIzIiwiaWF0IjoxNzM2NzM5MDgyLCJzdWIiOiJhZ2VudC10b2tlbiJ9.DL2cGJf-bDo_68TpH7H3htubg_enXTxCv8Ex-FTYJP7Sp4vB54YhenEcNTpoLi_yD_hpSp0OlgX7w69LndiatbSDcBJc_PeR5qQ_phV67UQ8ZvOW3smvtmUVuTFMUrr9F9SeA2-kZI2uFqk9apihGNAELoFLntBP_frhxdJFYBl1GoZuYQwAl438QcHpSraJh1s7UYrqxj3mKfeezosI1ZRdujoq6-BD8zfdMyFZs3SLweYNb5fkqPBsdSRNeHnUE643H_Vk6qXgs2BM8ibgR5F8IOqeipyh1zTdMGOfNl3ZPSjuBFzVRWEAhZq2zV1kyQk7qKEKCRf4F5TBa79arQ",

		    "agent": {
		      "accountId": "cm5uhlwnzsb2ss60cyqgq57xn",
		      "symbol": "CALADREL2",
		      "headquarters": "X1-ST33-A1",
		      "credits": 175000,
		      "startingFaction": "COSMIC",
		      "shipCount": 0
		    },

		    "contract": {
		      "id": "cm5uhlwpnsb2us60csf26o4sh",
		      "factionSymbol": "COSMIC",
		      "type": "PROCUREMENT",
		      "terms": {
		        "deadline": "2025-01-20T03:31:22.688Z",
		        "payment": {
		          "onAccepted": 2223,
		          "onFulfilled": 13338
		        },
		        "deliver": [
		          {
		            "tradeSymbol": "COPPER_ORE",
		            "destinationSymbol": "X1-ST33-H51",
		            "unitsRequired": 78,
		            "unitsFulfilled": 0
		          }
		        ]
		      },
		      "accepted": false,
		      "fulfilled": false,
		      "expiration": "2025-01-14T03:31:22.688Z",
		      "deadlineToAccept": "2025-01-14T03:31:22.688Z"
		    },

		    "faction": {
		      "symbol": "COSMIC",
		      "name": "Cosmic Engineers",
		      "description": "The Cosmic Engineers are a group of highly advanced scientists and engineers who seek to terraform and colonize new worlds, pushing the boundaries of technology and exploration.",
		      "headquarters": "X1-MY38",
		      "traits": [
		        {
		          "symbol": "INNOVATIVE",
		          "name": "Innovative",
		          "description": "Willing to try new and untested ideas. Sometimes able to come up with creative and original solutions to problems, and may be able to think outside the box. Sometimes at the forefront of technological or social change, and may be willing to take risks in order to advance the boundaries of human knowledge and understanding."
		        },
		        {
		          "symbol": "BOLD",
		          "name": "Bold",
		          "description": "Unafraid to take risks and challenge the status quo. Sometimes willing to do things that others would not dare, and may be able to overcome obstacles and challenges that would be insurmountable for others. Sometimes able to inspire and motivate others to take bold action as well."
		        },
		        {
		          "symbol": "VISIONARY",
		          "name": "Visionary",
		          "description": "Possessing a clear and compelling vision for the future. Sometimes able to see beyond the present and anticipate the needs and challenges of tomorrow. Sometimes able to inspire and guide others towards a better and brighter future, and may be willing to take bold and decisive action to make their vision a reality."
		        },
		        {
		          "symbol": "CURIOUS",
		          "name": "Curious",
		          "description": "Possessing a strong desire to learn and explore. Sometimes interested in a wide range of topics and may be willing to take risks in order to satisfy their curiosity. Sometimes able to think outside the box and come up with creative solutions to challenges."
		        }
		      ],
		      "isRecruiting": true
		    },

		    "ship": {
		      "symbol": "CALADREL2-1",
		      "nav": {
		        "systemSymbol": "X1-ST33",
		        "waypointSymbol": "X1-ST33-A1",
		        "route": {
		          "origin": {
		            "symbol": "X1-ST33-A1",
		            "type": "PLANET",
		            "systemSymbol": "X1-ST33",
		            "x": 27,
		            "y": 5
		          },
		          "destination": {
		            "symbol": "X1-ST33-A1",
		            "type": "PLANET",
		            "systemSymbol": "X1-ST33",
		            "x": 27,
		            "y": 5
		          },
		          "arrival": "2025-01-13T03:31:22.752Z",
		          "departureTime": "2025-01-13T03:31:22.752Z"
		        },
		        "status": "DOCKED",
		        "flightMode": "CRUISE"
		      },
		      "crew": {
		        "current": 57,
		        "capacity": 80,
		        "required": 57,
		        "rotation": "STRICT",
		        "morale": 100,
		        "wages": 0
		      },
		      "fuel": {
		        "current": 400,
		        "capacity": 400,
		        "consumed": {
		          "amount": 0,
		          "timestamp": "2025-01-13T03:31:22.752Z"
		        }
		      },
		      "cooldown": {
		        "shipSymbol": "CALADREL2-1",
		        "totalSeconds": 0,
		        "remainingSeconds": 0
		      },
		      "frame": {
		        "symbol": "FRAME_FRIGATE",
		        "name": "Frigate",
		        "description": "A medium-sized, multi-purpose spacecraft, often used for combat, transport, or support operations.",
		        "moduleSlots": 8,
		        "mountingPoints": 5,
		        "fuelCapacity": 400,
		        "condition": 1,
		        "integrity": 1,
		        "requirements": {
		          "power": 8,
		          "crew": 25
		        }
		      },
		      "reactor": {
		        "symbol": "REACTOR_FISSION_I",
		        "name": "Fission Reactor I",
		        "description": "A basic fission power reactor, used to generate electricity from nuclear fission reactions.",
		        "condition": 1,
		        "integrity": 1,
		        "powerOutput": 31,
		        "requirements": {
		          "crew": 8
		        }
		      },
		      "engine": {
		        "symbol": "ENGINE_ION_DRIVE_II",
		        "name": "Ion Drive II",
		        "description": "An advanced propulsion system that uses ionized particles to generate high-speed, low-thrust acceleration, with improved efficiency and performance.",
		        "condition": 1,
		        "integrity": 1,
		        "speed": 30,
		        "requirements": {
		          "power": 6,
		          "crew": 8
		        }
		      },
		      "modules": [
		        {
		          "symbol": "MODULE_CARGO_HOLD_II",
		          "name": "Expanded Cargo Hold",
		          "description": "An expanded cargo hold module that provides more efficient storage space for a ship's cargo.",
		          "capacity": 40,
		          "requirements": {
		            "crew": 2,
		            "power": 2,
		            "slots": 2
		          }
		        },
		        {
		          "symbol": "MODULE_CREW_QUARTERS_I",
		          "name": "Crew Quarters",
		          "description": "A module that provides living space and amenities for the crew.",
		          "capacity": 40,
		          "requirements": {
		            "crew": 2,
		            "power": 1,
		            "slots": 1
		          }
		        },
		        {
		          "symbol": "MODULE_CREW_QUARTERS_I",
		          "name": "Crew Quarters",
		          "description": "A module that provides living space and amenities for the crew.",
		          "capacity": 40,
		          "requirements": {
		            "crew": 2,
		            "power": 1,
		            "slots": 1
		          }
		        },
		        {
		          "symbol": "MODULE_MINERAL_PROCESSOR_I",
		          "name": "Mineral Processor",
		          "description": "Crushes and processes extracted minerals and ores into their component parts, filters out impurities, and containerizes them into raw storage units.",
		          "requirements": {
		            "crew": 0,
		            "power": 1,
		            "slots": 2
		          }
		        },
		        {
		          "symbol": "MODULE_GAS_PROCESSOR_I",
		          "name": "Gas Processor",
		          "description": "Filters and processes extracted gases into their component parts, filters out impurities, and containerizes them into raw storage units.",
		          "requirements": {
		            "crew": 0,
		            "power": 1,
		            "slots": 2
		          }
		        }
		      ],
		      "mounts": [
		        {
		          "symbol": "MOUNT_SENSOR_ARRAY_II",
		          "name": "Sensor Array II",
		          "description": "An advanced sensor array that improves a ship's ability to detect and track other objects in space with greater accuracy and range.",
		          "strength": 4,
		          "requirements": {
		            "crew": 2,
		            "power": 2
		          }
		        },
		        {
		          "symbol": "MOUNT_GAS_SIPHON_II",
		          "name": "Gas Siphon II",
		          "description": "An advanced gas siphon that can extract gas from gas giants and other gas-rich bodies more efficiently and at a higher rate.",
		          "strength": 20,
		          "requirements": {
		            "crew": 2,
		            "power": 2
		          }
		        },
		        {
		          "symbol": "MOUNT_MINING_LASER_II",
		          "name": "Mining Laser II",
		          "description": "An advanced mining laser that is more efficient and effective at extracting valuable minerals from asteroids and other space objects.",
		          "strength": 5,
		          "requirements": {
		            "crew": 2,
		            "power": 2
		          }
		        },
		        {
		          "symbol": "MOUNT_SURVEYOR_II",
		          "name": "Surveyor II",
		          "description": "An advanced survey probe that can be used to gather information about a mineral deposit with greater accuracy.",
		          "strength": 2,
		          "deposits": [
		            "QUARTZ_SAND",
		            "SILICON_CRYSTALS",
		            "PRECIOUS_STONES",
		            "ICE_WATER",
		            "AMMONIA_ICE",
		            "IRON_ORE",
		            "COPPER_ORE",
		            "SILVER_ORE",
		            "ALUMINUM_ORE",
		            "GOLD_ORE",
		            "PLATINUM_ORE",
		            "DIAMONDS",
		            "URANITE_ORE"
		          ],
		          "requirements": {
		            "crew": 4,
		            "power": 3
		          }
		        }
		      ],
		      "registration": {
		        "name": "CALADREL2-1",
		        "factionSymbol": "COSMIC",
		        "role": "COMMAND"
		      },
		      "cargo": {
		        "capacity": 40,
		        "units": 0,
		        "inventory": []
		      }
		    }
		  }
		}

	*/
}

type registerData struct {
	Token string `json:"token"`
	Agent agent  `json:"agent"`
	// todo contract
	// todo faction
	Ship ship `json:"ship"`
}

type agent struct {
	AccountId       string `json:"accountId"`
	Symbol          string `json:"symbol"`
	Headquarters    string `json:"headquarters"`
	Credits         int    `json:"credits"`
	StartingFaction string `json:"startingFaction"`
	ShipCount       int    `json:"shipCount"`
}

type ship struct {
	Symbol string `json:"symbol"`
	Nav    nav    `json:"nav"`
	Crew   crew   `json:"crew"`
	/*
		"ship": {
			      "symbol": "CALADREL2-1",
			      "nav": {
			        "systemSymbol": "X1-ST33",
			        "waypointSymbol": "X1-ST33-A1",
			        "route": {
			          "origin": {
			            "symbol": "X1-ST33-A1",
			            "type": "PLANET",
			            "systemSymbol": "X1-ST33",
			            "x": 27,
			            "y": 5
			          },
			          "destination": {
			            "symbol": "X1-ST33-A1",
			            "type": "PLANET",
			            "systemSymbol": "X1-ST33",
			            "x": 27,
			            "y": 5
			          },
			          "arrival": "2025-01-13T03:31:22.752Z",
			          "departureTime": "2025-01-13T03:31:22.752Z"
			        },
			        "status": "DOCKED",
			        "flightMode": "CRUISE"
			      },
			      "crew": {
			        "current": 57,
			        "capacity": 80,
			        "required": 57,
			        "rotation": "STRICT",
			        "morale": 100,
			        "wages": 0
			      },
			      "fuel": {
			        "current": 400,
			        "capacity": 400,
			        "consumed": {
			          "amount": 0,
			          "timestamp": "2025-01-13T03:31:22.752Z"
			        }
			      },
			      "cooldown": {
			        "shipSymbol": "CALADREL2-1",
			        "totalSeconds": 0,
			        "remainingSeconds": 0
			      },
			      "frame": {
			        "symbol": "FRAME_FRIGATE",
			        "name": "Frigate",
			        "description": "A medium-sized, multi-purpose spacecraft, often used for combat, transport, or support operations.",
			        "moduleSlots": 8,
			        "mountingPoints": 5,
			        "fuelCapacity": 400,
			        "condition": 1,
			        "integrity": 1,
			        "requirements": {
			          "power": 8,
			          "crew": 25
			        }
			      },
			      "reactor": {
			        "symbol": "REACTOR_FISSION_I",
			        "name": "Fission Reactor I",
			        "description": "A basic fission power reactor, used to generate electricity from nuclear fission reactions.",
			        "condition": 1,
			        "integrity": 1,
			        "powerOutput": 31,
			        "requirements": {
			          "crew": 8
			        }
			      },
			      "engine": {
			        "symbol": "ENGINE_ION_DRIVE_II",
			        "name": "Ion Drive II",
			        "description": "An advanced propulsion system that uses ionized particles to generate high-speed, low-thrust acceleration, with improved efficiency and performance.",
			        "condition": 1,
			        "integrity": 1,
			        "speed": 30,
			        "requirements": {
			          "power": 6,
			          "crew": 8
			        }
			      },
			      "modules": [
			        {
			          "symbol": "MODULE_CARGO_HOLD_II",
			          "name": "Expanded Cargo Hold",
			          "description": "An expanded cargo hold module that provides more efficient storage space for a ship's cargo.",
			          "capacity": 40,
			          "requirements": {
			            "crew": 2,
			            "power": 2,
			            "slots": 2
			          }
			        },
			        {
			          "symbol": "MODULE_CREW_QUARTERS_I",
			          "name": "Crew Quarters",
			          "description": "A module that provides living space and amenities for the crew.",
			          "capacity": 40,
			          "requirements": {
			            "crew": 2,
			            "power": 1,
			            "slots": 1
			          }
			        },
			        {
			          "symbol": "MODULE_CREW_QUARTERS_I",
			          "name": "Crew Quarters",
			          "description": "A module that provides living space and amenities for the crew.",
			          "capacity": 40,
			          "requirements": {
			            "crew": 2,
			            "power": 1,
			            "slots": 1
			          }
			        },
			        {
			          "symbol": "MODULE_MINERAL_PROCESSOR_I",
			          "name": "Mineral Processor",
			          "description": "Crushes and processes extracted minerals and ores into their component parts, filters out impurities, and containerizes them into raw storage units.",
			          "requirements": {
			            "crew": 0,
			            "power": 1,
			            "slots": 2
			          }
			        },
			        {
			          "symbol": "MODULE_GAS_PROCESSOR_I",
			          "name": "Gas Processor",
			          "description": "Filters and processes extracted gases into their component parts, filters out impurities, and containerizes them into raw storage units.",
			          "requirements": {
			            "crew": 0,
			            "power": 1,
			            "slots": 2
			          }
			        }
			      ],
			      "mounts": [
			        {
			          "symbol": "MOUNT_SENSOR_ARRAY_II",
			          "name": "Sensor Array II",
			          "description": "An advanced sensor array that improves a ship's ability to detect and track other objects in space with greater accuracy and range.",
			          "strength": 4,
			          "requirements": {
			            "crew": 2,
			            "power": 2
			          }
			        },
			        {
			          "symbol": "MOUNT_GAS_SIPHON_II",
			          "name": "Gas Siphon II",
			          "description": "An advanced gas siphon that can extract gas from gas giants and other gas-rich bodies more efficiently and at a higher rate.",
			          "strength": 20,
			          "requirements": {
			            "crew": 2,
			            "power": 2
			          }
			        },
			        {
			          "symbol": "MOUNT_MINING_LASER_II",
			          "name": "Mining Laser II",
			          "description": "An advanced mining laser that is more efficient and effective at extracting valuable minerals from asteroids and other space objects.",
			          "strength": 5,
			          "requirements": {
			            "crew": 2,
			            "power": 2
			          }
			        },
			        {
			          "symbol": "MOUNT_SURVEYOR_II",
			          "name": "Surveyor II",
			          "description": "An advanced survey probe that can be used to gather information about a mineral deposit with greater accuracy.",
			          "strength": 2,
			          "deposits": [
			            "QUARTZ_SAND",
			            "SILICON_CRYSTALS",
			            "PRECIOUS_STONES",
			            "ICE_WATER",
			            "AMMONIA_ICE",
			            "IRON_ORE",
			            "COPPER_ORE",
			            "SILVER_ORE",
			            "ALUMINUM_ORE",
			            "GOLD_ORE",
			            "PLATINUM_ORE",
			            "DIAMONDS",
			            "URANITE_ORE"
			          ],
			          "requirements": {
			            "crew": 4,
			            "power": 3
			          }
			        }
			      ],
			      "registration": {
			        "name": "CALADREL2-1",
			        "factionSymbol": "COSMIC",
			        "role": "COMMAND"
			      },
			      "cargo": {
			        "capacity": 40,
			        "units": 0,
			        "inventory": []
			      }
			    }
			  }
			}
	*/
}

type nav struct {
	SystemSymbol   string `json:"systemSymbol"`
	WaypointSymbol string `json:"waypointSymbol"`
	Route          route  `json:"route"`
	Status         string `json:"status"`
	FlightMode     string `json:"flightMode"`
}

type route struct {
	Origin        origin      `json:"origin"`
	Destination   destination `json:"destination"`
	Arrival       string      `json:"arrival"`
	DepartureTime string      `json:"departureTime"`
}

type origin struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type destination struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type crew struct {
	Current  int    `json:"current"`
	Capacity int    `json:"capacity"`
	Required int    `json:"required"`
	Rotation string `json:"rotation"`
	Morale   int    `json:"morale"`
	Wages    int    `json:"wages"`
}

type errorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Data    map[string]any `json:"data"`
}

var AgentAlreadyClaimedError errorResponse = errorResponse{
	Error: Error{
		Message: fmt.Sprintf("Cannot register agent. Agent symbol %s has already been claimed.", AgentSymbol),
		Code:    4111,
		Data: map[string]any{
			"agentSymbol": AgentSymbol,
		},
	},
}

func registerAgent(symbol string, faction string) string {
	// Create the register request JSON.
	registerJSON, _ := json.Marshal(map[string]string{
		"symbol":  symbol,
		"faction": faction,
	})

	// Marshal the JSON body into bytes.
	registerBody := bytes.NewBuffer(registerJSON)

	// Build the register request.
	registerReq, err := http.NewRequest(http.MethodPost, RegistrationEndpoint, registerBody)

	// Add the content-type header.
	// "Add" is a method of the request's Header, which modifies the Header's data. That's why we don't have a return value here; it's all done internally.
	registerReq.Header.Add("content-type", "application/json")

	// Send ("Do") the request and receive the response.
	registerResp, err := http.DefaultClient.Do(registerReq)
	if err != nil {
		// If an error was returned, log it (fatally, crashing the program).
		// This error would mean we had an issue in actually talking to the server, not that the server returned an error
		log.Fatalf("An error occurred during agent registration: %v", err)
	}

	// If we made it this far, there was no error (we would have crashed out of the program).

	// Tell Go that we want to close the response body once we're done here preemptively ("defer" closing it).
	// This is good practice to avoid memory leaks.
	defer registerResp.Body.Close()

	// Now that we've said we'll clean up the body, let's use it!
	// Read the body from the response
	registerRespBody, err := io.ReadAll(registerResp.Body)
	if err != nil {
		log.Fatalf("An error occurred reading the agent registration response body: %v", err)
	}

	var resp *registerResponse

	// TODO: Marshal the body into the response struct. If that doesn't work, marshal it into the error struct.
	err = json.Unmarshal(registerRespBody, resp)
	if err != nil {
		// TODO: Maybe we found an error here?
	}

	// Convert the body to a printable string.
	registerRespString := string(registerRespBody)

	// Print out the response to your console/terminal.
	log.Println("test?")
	log.Println(registerRespString)

	return registerRespString
}

func main() {
	// if token == "" {
	// 	// TODO: Set up registerAgent to use token from response.
	// 	token = registerAgent()
	// }

	registerAgent(AgentSymbol, Factions.Cosmic)

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
