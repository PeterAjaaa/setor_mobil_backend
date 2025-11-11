package models

type Model interface {
	TableName() string
}

type JwtClaims interface {
	GetID() uint
	GetEmail() string
	GetName() string
}
