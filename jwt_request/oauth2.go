package jwt_request

import (
	"strings"
)

// Strips 'Bearer ' prefix from bearer token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}

// Extract bearer token from Authorization header
// Uses PostExtractionFilter to strip "Bearer " prefix from header
var AuthorizationHeaderExtractor = &PostExtractionFilter{
	HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}
