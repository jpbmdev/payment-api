package models

// -----------------------------------------------
// -- Global return models
// -----------------------------------------------

type SucessfullOperation struct {
	Message string `json:"message"`
}

type FailedOperation struct {
	InternalCode string `json:"internalCode"`
	Message      string `json:"message"`
}
