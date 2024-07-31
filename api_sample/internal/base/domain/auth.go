package domain

type SessionId struct {
	value string
}

func (v SessionId) ToString() string {
	return v.value
}

func NewSessionId(sessionId string) SessionId {
	return SessionId{value: sessionId}
}

type AuthenticationError struct {
}

func (v AuthenticationError) Error() string {
	return "authentication error"
}
