package ship

import (
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
	Cooldown     ShipCooldown   `json:"cooldown"`
	Frame        Frame          `json:"frame"`
	Reactor      Reactor        `json:"reactor"`
	Engine       Engine         `json:"engine"`
	Modules      []Module       `json:"modules"`
	Mounts       []Mount        `json:"mounts"`
	Registration Registration   `json:"registration"`
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
