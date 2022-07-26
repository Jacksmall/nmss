package remote

import (
	"github.com/gomodule/redigo/redis"
)

const LuaScript = `
	local ticket_key = KEYS[1]
	local ticket_total_key = ARGV[1]
	local ticket_sold_key = ARGV[2]
	local ticket_total_nums = tonumber(redis.call('HGET', ticket_key, ticket_total_key))
	local ticket_sold_nums = tonumber(redis.call('HGET', ticket_key, ticket_sold_key))
	-- 查看是否还有余票，增加订单数量,返回结果集
	if (ticket_total_nums >= ticket_sold_nums) then
		return redis.call('HINCRBY', ticket_key, ticket_sold_key, 1)
	end
	return 0
`

type RemoteSpikeKeys struct {
	SpikeOrderHashKey  string // redis中秒杀订单hash结构key
	TotalInventoryKey  string // hash结构中总订单库存key
	QuantityOfOrderKey string // hash结构中已有订单数据key
}

// NewPool redis连接池
func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   10000,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, nil
		},
	}
}

// 远程扣减库存
func (rs *RemoteSpikeKeys) RemoteDeductionStock(conn redis.Conn) bool {
	lua := redis.NewScript(1, LuaScript)
	result, err := redis.Int(lua.Do(conn, rs.SpikeOrderHashKey, rs.TotalInventoryKey, rs.QuantityOfOrderKey))
	if err != nil {
		return false
	}
	return result != 0
}
