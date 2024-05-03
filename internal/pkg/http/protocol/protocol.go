package protocol

import (
	"encoding/json"
)

//go:generate easyjson -output_filename protocol_easyjson.go -all

type Request struct {
	Data json.RawMessage `json:"data"`
	Meta Meta            `json:"meta"`
}

type Meta struct {
	RequestID string `json:"request_id"`
	TraceID   string `json:"trace_id"`
	From      string `json:"from"`
}

type Response struct {
	Success bool            `json:"success"`
	Meta    Meta            `json:"meta"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   *ErrResponse    `json:"error,omitempty"`
}

type ErrResponse struct {
	Message     string `json:"message"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
