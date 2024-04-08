package environment

var envSchema EnvSchema

type EnvSchema struct {
	PORT         string `valid:"required"`
	DATABASE_URL string `valid:"required"`
}
