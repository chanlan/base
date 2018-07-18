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
var ExchangeMapping = map[int]string{
	0:  AndroidExchangeName,
	1:  MIExchangeName,
	2:  HWExchangeName,
	9:  IOSExchangeName,
	-1: DefaultExchangeName,
}
//Supported Exchange List
const (
	DefaultQueueName    = "android_push_sell"
	DefaultExchangeName = "android_push_sell"
	DefaultRoutingKey   = "android_push_sell"
	IOSQueueName        = "ios_push_sell"
	IOSExchangeName     = "ios_push_sell"
	IOSRoutingKey       = "ios_push_sell"
	AndroidQueueName    = "android_push_event"
	AndroidExchangeName = "android_push_event"
	AndroidRoutingKey   = "android_push_event"
	MIQueueName         = "xiaomi_push_event"
	MIExchangeName      = "xiaomi_push_event"
	MIRoutingKey        = "xiaomi_push_event"
	HWQueueName         = "huawei_push_event"
	HWExchangeName      = "huawei_push_event"
	HWRoutingKey        = "huawei_push_event"
	DefaultExchangeType = "direct"
)

//Queue Server
const QueueServer = "amqp://guest:guest@127.0.0.1:5672/"

//-----------------------------------------------------------------------------------------------------
//Mongodb
var (
	MogServers = "127.0.0.1:27017"
	MogUser    = ""
	MogPwd     = ""
	MogDBName  = "test"
)

//------------------------------------------------------------------------------------------------------
//redis
var RedisServers = "127.0.0.1:6709;127.0.0.1:6707;127.0.0.1:6708"
