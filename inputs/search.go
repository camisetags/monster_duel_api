package inputs

// CardSearchFieldInput struct
type CardSearchFieldInput struct {
	Name *string `json:"name"`
}

// CardSearchInput struct
type CardSearchInput struct {
	By     *CardSearchFieldInput `json:"by"`
	Limit  *int                  `json:"limit"`
	Offset *int                  `json:"offset"`
}
