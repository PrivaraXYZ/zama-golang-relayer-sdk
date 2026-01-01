package network

// PublicKeyResponse represents the response from the public key endpoint.
type PublicKeyResponse struct {
	PublicKey string `json:"publicKey"`
	ChainID   uint64 `json:"chainId"`
}

// EncryptRequest represents an encryption request to the relayer.
type EncryptRequest struct {
	ContractAddress string   `json:"contractAddress"`
	UserAddress     string   `json:"userAddress"`
	Values          [][]byte `json:"values"`
	Types           []string `json:"types"`
}

// EncryptResponse represents the response from the encrypt endpoint.
type EncryptResponse struct {
	Handles    []string `json:"handles"`
	InputProof string   `json:"inputProof"`
}

// ErrorResponse represents an error response from the relayer.
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
