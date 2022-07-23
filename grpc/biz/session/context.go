package session

import "context"

type key struct{}

var sessionKey key

func Load(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, sessionKey, *session)
}

func Extract(ctx context.Context) *Session {
	session, ok := ctx.Value(sessionKey).(Session)
	if !ok {
		return Empty()
	}

	return &session
}
