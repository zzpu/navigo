package data

import (
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/xormplus/xorm"
	"ucenter/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMemberRepo)

// Data .
type Data struct {
	db  *xorm.EngineGroup
	rdb *redis.Client
	// TODO warpped database client
}

// NewData .
func NewData(conf *conf.Data) (data *Data, err error) {
	//主库
	master, err := xorm.NewEngine(conf.Database.Driver, conf.Database.Source)
	if err != nil {
		return
	}

	//从库
	slave, err := xorm.NewEngine(conf.Database.Driver, conf.Database.Source)
	if err != nil {
		return
	}

	eg, err := xorm.NewEngineGroup(master, slave)

	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})

	return &Data{
		db:  eg,
		rdb: rdb,
	}, nil
}
