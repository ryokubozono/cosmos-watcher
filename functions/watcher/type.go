package watcher

type ProposalResponse struct {
	Description string     `json:"description"`
	Proposals   []Proposal `json:"proposals"`
}

type Proposal struct {
	Description      string           `json:"description"`
	ProposalID       string           `json:"proposal_id"`
	Content          Content          `json:"content"`
	Status           string           `json:"status"`
	SubmitTime       string           `json:"submit_time"`
	DepositEndTime   string           `json:"deposit_end_time"`
	FinalTallyResult FinalTallyResult `json:"final_tally_result"`
	VotingStartTime  string           `json:"voting_start_time"`
	VotingEndTime    string           `json:"voting_end_time"`
	Pagination       Pagination       `json:"pagination"`
}

type FinalTallyResult struct {
	Abstain    string `json:"abstain"`
	No         string `json:"no"`
	NoWithVeto string `json:"no_with_veto"`
	Yes        string `json:"yes"`
}

type Content struct {
	Type               string `json:"@type"`
	Description        string `json:"description"`
	SubjectClientID    string `json:"subject_client_id"`
	SubstituteClientID string `json:"substitute_client_id"`
	Title              string `json:"title"`
}

type Pagination struct {
	NextKey interface{} `json:"next_key"`
	Total   string      `json:"total"`
}
