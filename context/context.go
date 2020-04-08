package context

import "context"

type KeyReference string

const (
	UserIDKey = KeyReference("UserID")
	ResourceIDKey = KeyReference("ResourceID")
)

func ContextKeys(ctx context.Context)(userID, resourceID string) {
	if ctx == nil {
		return
	}

	if rawUserID, ok := ctx.Value(UserIDKey).(string); ok {
		userID = rawUserID
	}

	if rawResourceID, ok := ctx.Value(ResourceIDKey).(string); ok {
		resourceID = rawResourceID
	}

	return
}
