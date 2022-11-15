package data

import (
	"context"
	"github.com/bighuangbee/account-svc/internal/conf"
	"github.com/bighuangbee/account-svc/internal/pkg/snowflakeId"
	"github.com/bighuangbee/gokit/storage/kitGorm"
	"github.com/bighuangbee/gokit/storage/kitRedis"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

var (
	ctx context.Context
	// 自动添加前缀
	autoRedisPrefix string
)

func init() {
	ctx = context.Background()
}

type (
	Prefix  struct{}
	MyRedis struct {
		Rdb    kitRedis.Client
		Prefix string
	}
)
type Data struct {
	dbInfo  *conf.Database
	db      *gorm.DB
	rdb     kitRedis.Client
	snowflakeId snowflakeId.ISnowflakeId
}

func NewData(bc *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	c := bc.Data

	autoRedisPrefix = c.Redis.AutoPrefix
	var db *gorm.DB
	db, err := kitGorm.New(&kitGorm.Options{
		Address:  c.Database.Address,
		UserName: c.Database.UserName,
		Password: c.Database.Password,
		DBName:   c.Database.DBName,
		// Tracer:   otel.GetTracerProvider(),
		Logger:  kitGorm.Logger{L: log.NewHelper(logger)},
		Charset: "utf8mb4",
	})
	if err != nil {
		return nil, nil, err
	}

	sfId, err := snowflakeId.New(bc.Server.NodeId)
	if err != nil {
		panic("snowflakeId fail" + err.Error())
	}

	rClient, err := kitRedis.New(&kitRedis.Options{
		Addr:     c.Redis.Address,
		Password: c.Redis.Password,
		DB:       int(c.Redis.DB),
	})
	if err != nil {
		return nil, nil, err
	}

	logger.Log(log.LevelDebug, "db connect:", c.Database.Address, ",driver:", c.Database.Driver)


	d := &Data{
		dbInfo:  c.Database,
		db:      db,
		rdb:     rClient,
		snowflakeId:      sfId,
	}

	return d, func() {
		d.rdb.Close()
	}, nil
}


func (d *Data) DB(ctx context.Context) *gorm.DB {
	return d.db.WithContext(ctx)
}

func (d *Data) Redis(prefix string) *MyRedis {
	if len(prefix) == 0 {
		prefix = autoRedisPrefix
	}
	myredis := MyRedis{Prefix: prefix, Rdb: d.rdb}
	return &myredis
}

func (d *Data) SnowflakeId() snowflakeId.ISnowflakeId {
	return d.snowflakeId
}

// key不需要加 prefix，自动加
func (t *MyRedis) RedisAdd(key, value string) error {
	err := t.Rdb.Set(ctx, t.Prefix+key, value, 0).Err()
	return err
}

// 手动添加过期时间的值
func (t *MyRedis) RedisAddAndExp(key, value string, exp time.Duration) error {
	err := t.Rdb.Set(ctx, t.Prefix+key, value, exp).Err()
	return err
}

// key不需要加 prefix，自动加
func (t *MyRedis) RedisGet(key string) (value string, err error) {
	return t.Rdb.Get(ctx, t.Prefix+key).Result()
}

// exist>0存在
func (t *MyRedis) RedisExist(key string) (exist int64, err error) {
	return t.Rdb.Exists(ctx, t.Prefix+key).Result()
}

// key不需要加 prefix，自动加
func (t *MyRedis) RedisHAdd(key, field, value string) error {
	err := t.Rdb.HSet(ctx, t.Prefix+key, field, value).Err()
	return err
}

// key不需要加 prefix，自动加
func (t *MyRedis) RedisHGet(key, field string) (value string, err error) {
	return t.Rdb.HGet(ctx, t.Prefix+key, field).Result()
}

//  key不需要加 prefix，自动加
func (t *MyRedis) RedisHExist(key, field string) (exist bool, err error) {
	return t.Rdb.HExists(ctx, t.Prefix+key, field).Result()
}

// key不需要加 prefix，自动加
func (t *MyRedis) RedisIncr(key string) (value int64, err error) {
	return t.Rdb.Incr(ctx, t.Prefix+key).Result()
}

// 删除key
//  key不需要加 prefix，自动加
func (t *MyRedis) RedisDel(key string) (value int64, err error) {
	return t.Rdb.Del(ctx, t.Prefix+key).Result()
}
