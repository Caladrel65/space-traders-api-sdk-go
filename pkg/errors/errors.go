package errors

type APIError struct {
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Data    map[string]any `json:"data"`
}

// var AgentAlreadyClaimedError APIError = APIError{
// 	Message: fmt.Sprintf("Cannot register agent. Agent symbol %s has already been claimed.", agent.AgentSymbol),
// 	Code:    4111,
// 	Data: map[string]any{
// 		"agentSymbol": agent.AgentSymbol,
// 	},
// }
