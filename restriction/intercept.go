package restriction

import (
	"context"
	"github.com/XCPCBoard/common/dao"
	"github.com/XCPCBoard/common/logger"
	"strconv"
	"time"
)

// LimitAccess 限制服务的访问次数
//
//	@param	duration	存活时间
//	@param	times	限制访问次数
//	@return	isOK	为true时，资源可以访问
func LimitAccess(key string, duration time.Duration, times int) (bool, error) {

	backGround := context.Background()
	val, err := dao.RedisClient.Exists(backGround, key).Result()
	if err != nil {
		logger.L.Error("判断受限资源key是否存在时错误", err, 0, key)
		return false, err
	}

	if val == 0 { // no exists
		_, err := dao.RedisClient.Incr(backGround, key).Result()
		if err != nil {
			logger.L.Error("计数是出错", err, 0, key)
			return false, err
		}

		//设置键的过期时间
		_, err = dao.RedisClient.Expire(backGround, key, duration).Result()
		if err != nil {
			logger.L.Error("设置key过期时间出错", err, 0, key)
			return false, err
		}
		//设置成功，可以访问
		return true, nil
	}
	//exist
	count, err := dao.RedisClient.Get(backGround, key).Result()
	if err != nil {
		logger.L.Error("获取key具体数据时出错", err, 0, key)
		return false, err
	}

	temp, _err := strconv.Atoi(count)
	if _err != nil {
		logger.L.Error("获取key具体数据时出错", err, 0, key)
		return false, _err
	}
	//低于限制次数
	if temp < times {
		_, err = dao.RedisClient.Incr(backGround, key).Result()
		if err != nil {
			logger.L.Error("计数是出错", err, 0, key)
			return false, err
		}
		//成功访问
		return true, nil
	}
	//否则就是评率过快
	logger.L.Warn("资源访问频率过快", 0, key)
	return false, nil
}
