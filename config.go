package midware

// ------------------------------------------------------------------------------------------------
//Database default config
const (
	DefaultDBDriver  = "mysql"
	DefaultDBPrefix  = ""
	DefaultDBShowSQL = true
	DefaultDBMaxOpen = 10
	DefaultDBMaxIdle = 10
	DefaultDBDSN     = "root:123456@(127.0.0.1:3306)/test"
)

//--------------------------------------------------------------------------------------------------
//RabbitMQ
//ExchangeMapping　the mapping of device and exchange
const (
	EXCHANGE_DEFAULT = iota
	EXCHANGE_ANDROID
	EXCHANGE_XIAOMI
	EXCHANGE_HUAWEI
	EXCHANGE_IOS
)

var ExchangeMapping = map[int]string{
	EXCHANGE_ANDROID: AndroidExchangeName,
	EXCHANGE_XIAOMI:  MIExchangeName,
	EXCHANGE_HUAWEI:  HWExchangeName,
	EXCHANGE_IOS:     IOSExchangeName,
	EXCHANGE_DEFAULT: DefaultExchangeName,
}
//Supported Exchange List
const (
	DefaultQueueName    = "default"
	DefaultExchangeName = "default"
	DefaultRoutingKey   = "default"
	IOSQueueName        = "ios"
	IOSExchangeName     = "ios"
	IOSRoutingKey       = "ios"
	AndroidQueueName    = "android"
	AndroidExchangeName = "android"
	AndroidRoutingKey   = "android"
	MIQueueName         = "xiaomi"
	MIExchangeName      = "xiaomi"
	MIRoutingKey        = "xiaomi"
	HWQueueName         = "huawei"
	HWExchangeName      = "huawei"
	HWRoutingKey        = "huawei"
)
const DefaultExchangeType = "direct"

//Queue Server
const QueueServer = "amqp://guest:guest@127.0.0.1:5672/"

//-----------------------------------------------------------------------------------------------------
//Mongodb
const (
	MogServers = "127.0.0.1:27017"
	MogUser    = ""
	MogPwd     = ""
	MogDBName  = "test"
)

//------------------------------------------------------------------------------------------------------
//redis
const RedisServers = "127.0.0.1:6709;127.0.0.1:6707;127.0.0.1:6708"
