package entity

type Key = string

func NewKey(s string) Key {
	return Key(s)
}

func ValidateKey(s string) bool {
	return len(s) == 44
}
