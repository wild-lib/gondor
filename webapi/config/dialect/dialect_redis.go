package dialect

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

const REDIS_DEFAULT_PORT uint16 = 6379

type Redis struct {
	Options []redis.DialOption
}

func (Redis) Name() string {
	return "redis"
}

func (Redis) QuoteIdent(ident string) string {
	return WrapWith(ident, "'", "'")
}

func (r *Redis) GetDSN(params ConnParams) string {
	r.Options = make([]redis.DialOption, 0)
	dsn := params.GetAddr("127.0.0.1", REDIS_DEFAULT_PORT)
	if params.Password != "" {
		r.Options = append(r.Options, redis.DialPassword(params.Password))
	}
	if dbno, err := strconv.Atoi(params.Database); err == nil {
		r.Options = append(r.Options, redis.DialDatabase(dbno))
	}
	return dsn
}

func (r *Redis) Connect(params ConnParams) (redis.Conn, error) {
	dsn := r.GetDSN(params)
	return redis.Dial("tcp", dsn, r.Options...)
}
