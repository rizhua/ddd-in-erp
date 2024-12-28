package constant

type RedisPrefix string

const (
	TOKEN   RedisPrefix = "token"
	USER    RedisPrefix = "user"
	CAPTCHA RedisPrefix = "captcha"
)
