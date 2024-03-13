package util

type AccountIDGenerator interface {
	GenerateID() string
}
