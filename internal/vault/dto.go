package vault

type ErrorResponse struct {
	Error string `json:"error"`
}

type VaultDTO struct {
	Name        string `json:name`
	Description string `json:description`
}
