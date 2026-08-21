package constants

// Redis key 常量，统一维护
const (
	// RedisKeyMktManagerSyncAdvertiser 广告主同步队列
	RedisKeyMktManagerSyncAdvertiser = "stack:mkt:manager:sync:advertiser"

	// RedisKeyMktManagerRefreshToken 管家 token 刷新分布式锁，%d 为 manager id
	RedisKeyMktManagerRefreshToken = "stack:mkt:manager:refresh:token:%d"
)
