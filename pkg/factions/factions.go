package factions

// Factions is a struct containing all of the factions in the game.
type Faction string

const (
	Faction_COSMIC   Faction = "COSMIC"
	Faction_GALACTIC Faction = "GALACTIC"
	Faction_SOLAR    Faction = "SOLAR"
)

type FactionDetails struct {
	Symbol       Faction        `json:"symbol"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Headquarters string         `json:"headquarters"`
	Traits       []FactionTrait `json:"traits"`
	IsRecruiting bool           `json:"isRecruiting"`
}

type FactionTrait struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var (
	FactionTrait_INNOVATIVE = FactionTrait{
		Symbol:      "INNOVATIVE",
		Name:        "Innovative",
		Description: "Willing to try new and untested ideas. Sometimes able to come up with creative and original solutions to problems, and may be able to think outside the box. Sometimes at the forefront of technological or social change, and may be willing to take risks in order to advance the boundaries of human knowledge and understanding.",
	}
	FactionTrait_BOLD = FactionTrait{
		Symbol:      "BOLD",
		Name:        "Bold",
		Description: "Unafraid to take risks and challenge the status quo. Sometimes willing to do things that others would not dare, and may be able to overcome obstacles and challenges that would be insurmountable for others. Sometimes able to inspire and motivate others to take bold action as well.",
	}
	FactionTrait_VISIONARY = FactionTrait{
		Symbol:      "VISIONARY",
		Name:        "Visionary",
		Description: "Possessing a clear and compelling vision for the future. Sometimes able to see beyond the present and anticipate the needs and challenges of tomorrow. Sometimes able to inspire and guide others towards a better and brighter future, and may be willing to take bold and decisive action to make their vision a reality.",
	}
	FactionTrait_CURIOUS = FactionTrait{
		Symbol:      "CURIOUS",
		Name:        "Curious",
		Description: "Possessing a strong desire to learn and explore. Sometimes interested in a wide range of topics and may be willing to take risks in order to satisfy their curiosity. Sometimes able to think outside the box and come up with creative solutions to challenges.",
	}
)
