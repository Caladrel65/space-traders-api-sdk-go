package models

import (
	"fmt"
	"space-traders-api-sdk-go/pkg/client"
	"space-traders-api-sdk-go/pkg/factions"
	"space-traders-api-sdk-go/pkg/navigation"
)

type Fuel struct {
	Current  int          `json:"current"`
	Capacity int          `json:"capacity"`
	Consumed ConsumedFuel `json:"consumed"`
}

type ConsumedFuel struct {
	Amount    int    `json:"amount"`
	Timestamp string `json:"timestamp"` // TODO: What is this for - time we checked? Time we'll run out?
}

type ShipCooldown struct {
	ShipSymbol       string `json:"shipSymbol"`
	TotalSeconds     int    `json:"totalSeconds"`
	RemainingSeconds int    `json:"remainingSeconds"`
}

type Frame struct {
	Symbol         FrameSymbol       `json:"symbol"`      // TODO: Customizeable?
	Name           string            `json:"name"`        // TODO: Customizeable?
	Description    string            `json:"description"` // TODO: Customizeable?
	ModuleSlots    int               `json:"moduleSlots"`
	MountingPoints int               `json:"mountaingPoints"`
	FuelCapacity   int               `json:"fuelCapacity"`
	Condition      int               `json:"condition"`
	Integrity      int               `json:"integrity"`
	Requirements   FrameRequirements `json:"requirements"`
}

type FrameSymbol string

const (
	FrameSymbol_FRAME_FRIGATE FrameSymbol = "FRAME_FRIGATE"
)

// TODO: Can probably collapse into a single requirements type?
type FrameRequirements struct {
	Power int `json:"power"`
	Crew  int `json:"crew"`
}

type Reactor struct {
	Symbol       ReactorSymbol       `json:"symbol"`
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Condition    int                 `json:"condition"`
	Integrity    int                 `json:"integrity"`
	PowerOutput  int                 `json:"powerOutput"`
	Requirements ReactorRequirements `json:"requirements"`
}

type ReactorSymbol string

const ReactorSymbol_REACTOR_FISSION_I ReactorSymbol = "REACTOR_FISSION_I"

type ReactorRequirements struct {
	Crew int `json:"crew"`
}

type Engine struct {
	Symbol       EngineSymbol      `json:"symbol"`      // TODO: Customizeable?
	Name         string            `json:"name"`        // TODO: Customizeable?
	Description  string            `json:"description"` // TODO: Customizeable?
	Condition    int               `json:"condition"`
	Integrity    int               `json:"integrity"`
	Speed        int               `json:"speed"`
	Requirements FrameRequirements `json:"requirements"`
}

type EngineSymbol string

const (
	EngineSymbol_ENGINE_ION_DRIVE_II EngineSymbol = "ENGINE_ION_DRIVE_II"
)

type EngineRequirements struct {
	Power int `json:"power"`
	Crew  int `json:"crew"`
}

type Module struct {
	Symbol       ModuleSymbol       `json:"symbol"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Capacity     int                `json:"capacity"`
	Requirements ModuleRequirements `json:"requirements"`
}

type ModuleSymbol string

const (
	ModuleSymbol_MODULE_CARGO_HOLD_II       ModuleSymbol = "MODULE_CARGO_HOLD_II"
	ModuleSymbol_MODULE_CREW_QUARTERS_I     ModuleSymbol = "MODULE_CREW_QUARTERS_I"
	ModuleSymbol_MODULE_MINERAL_PROCESSOR_I ModuleSymbol = "MODULE_MINERAL_PROCESSOR_I"
	ModuleSymbol_MODULE_GAS_PROCESSOR_I     ModuleSymbol = "MODULE_GAS_PROCESSOR_I"
)

type ModuleRequirements struct {
	Crew  int `json:"crew"`
	Power int `json:"power"`
	Slots int `json:"slots"`
}

type Mount struct {
	Symbol       MountSymbol       `json:"symbol"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Strength     int               `json:"strength"`
	Deposits     []Deposit         `json:"deposits"`
	Requirements MountRequirements `json:"requirements"`
}

type MountSymbol string

const (
	MountSymbol_MOUNT_SENSOR_ARRAY_II MountSymbol = "MOUNT_SENSOR_ARRAY_II"
	MountSymbol_MOUNT_GAS_SIPHON_II   MountSymbol = "MOUNT_GAS_SIPHON_II"
	MountSymbol_MOUNT_MINING_LASER_II MountSymbol = "MOUNT_MINING_LASER_II"
	MountSymbol_MOUNT_SURVEYOR_II     MountSymbol = "MOUNT_SURVEYOR_II"
)

type Deposit string

const (
	Deposit_QUARTZ_SAND      Deposit = "QUARTZ_SAND"
	Deposit_SILICON_CRYSTALS Deposit = "SILICON_CRYSTALS"
	Deposit_PRECIOUS_STONES  Deposit = "PRECIOUS_STONES"
	Deposit_ICE_WATER        Deposit = "ICE_WATER"
	Deposit_AMMONIA_ICE      Deposit = "AMMONIA_ICE"
	Deposit_IRON_ORE         Deposit = "IRON_ORE"
	Deposit_COPPER_ORE       Deposit = "COPPER_ORE"
	Deposit_SILVER_ORE       Deposit = "SILVER_ORE"
	Deposit_ALUMINUM_ORE     Deposit = "ALUMINUM_ORE"
	Deposit_GOLD_ORE         Deposit = "GOLD_ORE"
	Deposit_PLATINUM_ORE     Deposit = "PLATINUM_ORE"
	Deposit_DIAMONDS         Deposit = "DIAMONDS"
	Deposit_URANITE_ORE      Deposit = "URANITE_ORE"
)

type MountRequirements struct {
	Crew  int `json:"crew"`
	Power int `json:"power"`
}

type Registration struct {
	Name          string           `json:"name"`
	FactionSymbol factions.Faction `json:"factionSymbol"`
	Role          RegistrationRole `json:"role"`
}

type RegistrationRole string

const RegistrationRole_COMMAND RegistrationRole = "COMMAND"

type Cargo struct {
	Capacity  int         `json:"capacity"`
	Units     int         `json:"units"`
	Inventory []Inventory `json:"inventory"`
}

// TODO: What can be in the inventory? We may never know.
type Inventory string

type Ship struct {
	Symbol       string         `json:"symbol"` // ex "CALADREL2-1"
	Nav          navigation.Nav `json:"nav"`
	Crew         Crew           `json:"crew"`
	Fuel         Fuel           `json:"fuel"`
	Cargo        Cargo          `json:"cargo"`
	Cooldown     ShipCooldown   `json:"cooldown"`
	Frame        Frame          `json:"frame"`
	Reactor      Reactor        `json:"reactor"`
	Engine       Engine         `json:"engine"`
	Modules      []Module       `json:"modules"`
	Mounts       []Mount        `json:"mounts"`
	Registration Registration   `json:"registration"`
	Role         string         `json:"role"`
}

type Crew struct {
	Current  int          `json:"current"`
	Capacity int          `json:"capacity"`
	Required int          `json:"required"`
	Rotation CrewRotation `json:"rotation"`
	Morale   int          `json:"morale"`
	Wages    int          `json:"wages"`
}

type CrewRotation string

const (
	CrewRotation_STRICT = "STRICT"
	// TODO: Others
)

type Shipyard struct {
	Symbol       string         `json:"symbol"`
	ShipTypes    []ShipType     `json:"shipTypes"`
	Transactions []Transaction  `json:"transactions"`
	Ships        []ShipyardShip `json:"ships"`
}

type ShipType struct {
	Type string `json:"type"`
}

type Transaction struct {
	WaypointSymbol string `json:"waypointSymbol"`
	ShipSymbol     string `json:"shipSymbol"`
	Price          int    `json:"price"`
	AgentSymbol    string `json:"agentSymbol"`
	Timestamp      string `json:"timestamp"`
}

type ShipyardShip struct {
	Type           string         `json:"type"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	PurchasePrice  int            `json:"purchasePrice"`
	Frame          Frame          `json:"frame"`
	Reactor        Reactor        `json:"reactor"`
	Engine         Engine         `json:"engine"`
	Modules        []Module       `json:"modules"`
	Mounts         []Mount        `json:"mounts"`
	Crew           Crew           `json:"crew"`
	Fuel           Fuel           `json:"fuel"`
	Nav            navigation.Nav `json:"nav"`
}

func GetShipyard(client *client.Client, systemSymbol string, waypointSymbol string) (*Shipyard, error) {
	var resp struct {
		Data Shipyard `json:"data"`
	}
	err := client.Get(fmt.Sprintf("/systems/%s/waypoints/%s/shipyard", systemSymbol, waypointSymbol), &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

type PurchaseShipResponse struct {
	Data struct {
		Agent       Agent       `json:"agent"`
		Ship        Ship        `json:"ship"`
		Transaction Transaction `json:"transaction"`
	} `json:"data"`
}

func PurchaseShip(client *client.Client, shipType string, waypointSymbol string) (*PurchaseShipResponse, error) {
	body := map[string]string{
		"shipType":       shipType,
		"waypointSymbol": waypointSymbol,
	}

	var resp PurchaseShipResponse
	err := client.Post("/my/ships", body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type NavigateShipResponse struct {
	Data struct {
		Fuel Fuel           `json:"fuel"`
		Nav  navigation.Nav `json:"nav"`
	} `json:"data"`
}

func Navigate(client *client.Client, shipSymbol string, waypointSymbol string) (*NavigateShipResponse, error) {
	body := map[string]string{
		"waypointSymbol": waypointSymbol,
	}

	var resp NavigateShipResponse
	err := client.Post(fmt.Sprintf("/my/ships/%s/navigate", shipSymbol), body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type OrbitShipResponse struct {
	Data struct {
		Nav navigation.Nav `json:"nav"`
	} `json:"data"`
}

func Orbit(client *client.Client, shipSymbol string) (*OrbitShipResponse, error) {
	var resp OrbitShipResponse
	err := client.Post(fmt.Sprintf("/my/ships/%s/orbit", shipSymbol), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DockShipResponse struct {
	Data struct {
		Nav navigation.Nav `json:"nav"`
	} `json:"data"`
}

func Dock(client *client.Client, shipSymbol string) (*DockShipResponse, error) {
	var resp DockShipResponse
	err := client.Post(fmt.Sprintf("/my/ships/%s/dock", shipSymbol), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RefuelShipResponse struct {
	Data struct {
		Agent Agent `json:"agent"`
		Fuel  Fuel  `json:"fuel"`
	} `json:"data"`
}

func Refuel(client *client.Client, shipSymbol string) (*RefuelShipResponse, error) {
	var resp RefuelShipResponse
	err := client.Post(fmt.Sprintf("/my/ships/%s/refuel", shipSymbol), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ExtractResourcesResponse struct {
	Data struct {
		Cooldown   ShipCooldown `json:"cooldown"`
		Extraction Extraction   `json:"extraction"`
		Cargo      Cargo        `json:"cargo"`
	} `json:"data"`
}

type Extraction struct {
	ShipSymbol string `json:"shipSymbol"`
	Yield      Yield  `json:"yield"`
}

type Yield struct {
	Symbol string `json:"symbol"`
	Units  int    `json:"units"`
}

func Extract(client *client.Client, shipSymbol string) (*ExtractResourcesResponse, error) {
	var resp ExtractResourcesResponse
	err := client.Post(fmt.Sprintf("/my/ships/%s/extract", shipSymbol), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type TransferCargoResponse struct {
	Data struct {
		Cargo Cargo `json:"cargo"`
	} `json:"data"`
}

func TransferCargo(client *client.Client, shipSymbol string, tradeSymbol string, units int, toShipSymbol string) (*TransferCargoResponse, error) {
	body := map[string]interface{}{
		"tradeSymbol":  tradeSymbol,
		"units":        units,
		"shipSymbol": toShipSymbol,
	}

	var resp TransferCargoResponse
	err := client.Post(fmt.Sprintf("/my/ships/%s/transfer", shipSymbol), body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetMyShipsResponse struct {
	Data []Ship `json:"data"`
}

func GetMyShips(client *client.Client) ([]Ship, error) {
	var resp GetMyShipsResponse
	err := client.Get("/my/ships", &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
