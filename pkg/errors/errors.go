package errors

import (
	"fmt"
	"space-traders-api-sdk-go/pkg/agent"
)

// TODO: Is this nesting really needed?
type APIError struct {
	Error Error `json:"error"`
}

type Error struct {
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Data    map[string]any `json:"data"`
}

var AgentAlreadyClaimedError APIError = APIError{
	Error: Error{
		Message: fmt.Sprintf("Cannot register agent. Agent symbol %s has already been claimed.", agent.AgentSymbol),
		Code:    4111,
		Data: map[string]any{
			"agentSymbol": agent.AgentSymbol,
		},
	},
}
