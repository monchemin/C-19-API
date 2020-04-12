package context

import "context"

type KeyReference string

const (
	UserIDKey = KeyReference("UserID")
	ResourceIDKey = KeyReference("ResourceID")
	TokenKey = KeyReference("TokenKey")
)

type Values struct {
	UserID string
	ResourceID string
	Token string
}
func ContextKeys(ctx context.Context)(data Values) {
	if ctx == nil {
		return
	}

	if rawUserID, ok := ctx.Value(UserIDKey).(string); ok {
		data.UserID = rawUserID
	}

	if rawResourceID, ok := ctx.Value(ResourceIDKey).(string); ok {
		data.ResourceID = rawResourceID
	}

	if rawToken, ok := ctx.Value(TokenKey).(string); ok {
		data.Token = rawToken
	}

	return
}
