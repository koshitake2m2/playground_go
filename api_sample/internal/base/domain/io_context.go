package domain

type IOContext interface {
	Transaction() Transaction
}
