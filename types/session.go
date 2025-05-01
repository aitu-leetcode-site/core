package types

import "context"

type Session struct {
	userID UserID
}

func NewSession(userID UserID) *Session {
	return &Session{
		userID: userID,
	}
}

func (s *Session) UserID() UserID {
	return s.userID
}

func (s *Session) Validate() error {
	return s.userID.Validate()
}

const SessionKey = "_sess"

func (s *Session) SetInContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, SessionKey, s)
}

func GetSession(ctx context.Context) *Session {
	sess, ok := ctx.Value(SessionKey).(*Session)
	if !ok {
		return nil
	}
	return sess
}
