package utils

import (
	"awesomeProject2/types"
	"encoding/base64"
	"encoding/json"
)

// DecodeToken decodes a base64 URL-encoded token into a PageToken.
func DecodeToken(token string) (*types.PageToken, error) {
	var result types.PageToken
	bytes, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// EncodeToken encodes a PageToken into a base64 URL-safe string.
func EncodeToken(request *types.PageToken) (string, error) {
	bytes, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
