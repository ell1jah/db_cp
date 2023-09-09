package context

import (
	"context"

	"github.com/pkg/errors"
)

type contextKeyType string

const contextUserKey contextKeyType = "contextUserKey"

type ContextManager struct{}

func (cu ContextManager) ContextWithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, contextUserKey, userID)
}

func (cu *ContextManager) UserIDFromContext(ctx context.Context) (int, error) {
	user, ok := ctx.Value(contextUserKey).(int)
	if !ok {
		return -1, errors.Errorf("can`t get user from context")
	}

	return user, nil
}
