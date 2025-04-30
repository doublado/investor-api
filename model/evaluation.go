package model

type EvaluationRequest struct {
	Scenario string `json:"scenario"`
	Idea     string `json:"idea"`
}

type EvaluationResponse struct {
	Decision string `json:"investmentDecision"` // "Approved", "Rejected", etc.
	Amount   int    `json:"amount"`             // Investment amount in dkk
	Comment  string `json:"comment"`            // AI's rationale
}
