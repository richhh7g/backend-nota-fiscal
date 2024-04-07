package config

type ConfigBase[T any] interface {
	Configure() (T, error)
}
