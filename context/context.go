package context

import "context"

type KeyReference string

const (
	UserIDKey = KeyReference("UserID")
	ResourceIDKey = KeyReference("ResourceID")
	TokenKey = KeyReference("ResourceID")
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

	if token, ok := ctx.Value(ResourceIDKey).(string); ok {
		data.Token = token
	}

	return
}
