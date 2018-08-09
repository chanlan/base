/**
 * This is NOT a freeware, use is subject to license terms
 *
 * path   go-push/midware
 * date   2018/6/19 11:18
 * author chenjingxiu
 */
package midware

import (
	"time"
	"errors"
	"github.com/streadway/amqp"
	"github.com/jolestar/go-commons-pool"
)

type DefaultQueueChannel struct {
	QueueName    string
	RoutingKey   string
	ExchangeName string
	ExchangeType string
	Channel      *amqp.Channel
	Queue        *amqp.Queue
	Publishing   *amqp.Publishing
	CreateTime   time.Time
}

type AndroidQueueChannel struct {
	DefaultQueueChannel
}

type IOSQueueChannel struct {
	DefaultQueueChannel
}

type MIQueueChannel struct {
	DefaultQueueChannel
}

type HWQueueChannel struct {
	DefaultQueueChannel
}

type Consumer interface {
	Consumer([]byte)
}

var (
	defaultPool *pool.ObjectPool
	androidPool *pool.ObjectPool
	iosPool     *pool.ObjectPool
	miPool      *pool.ObjectPool
	hwPool      *pool.ObjectPool
)

func init() {
	NewAndroidChannel(QueueServer, amqpConfig)
	NewDefaultChannel(QueueServer, amqpConfig)
	NewHWChannel(QueueServer, amqpConfig)
	NewIOSChannel(QueueServer, amqpConfig)
	NewMIChannel(QueueServer, amqpConfig)
}

//NewDefaultChannel create a default channel to transport messages,
//by passed a parameter to url used to connect the rabbitMQ
func NewDefaultChannel(url string, config amqp.Config) (*pool.ObjectPool, error) {
	PoolConfig := pool.NewDefaultPoolConfig()
	PoolConfig.MaxTotal = 10
	WithAbandonedConfig := pool.NewDefaultAbandonedConfig()
	defaultPool = pool.NewObjectPoolWithAbandonedConfig(pool.NewPooledObjectFactorySimple(
		func() (interface{}, error) {
			conn, err := amqp.DialConfig(url, config)
			if err != nil {
				return nil, err
			}
			ch, err := conn.Channel()
			if err != nil {
				return nil, err
			}
			queue, err := ch.QueueDeclare(
				DefaultQueueName,
				true,
				false,
				false,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			err = ch.ExchangeDeclare(
				DefaultExchangeName,
				DefaultExchangeType,
				true,
				false,
				false,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			err = ch.QueueBind(
				DefaultQueueName,
				DefaultRoutingKey,
				DefaultExchangeName,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			return newDefaultQueue(ch, &queue), nil
		}), PoolConfig, WithAbandonedConfig)
	return defaultPool, nil
}

//NewAndroidChannel create a channel to transport messages,
//by passed a parameter to url used to connect the rabbitMQ
//For Android Device
func NewAndroidChannel(url string, config amqp.Config) (*pool.ObjectPool, error) {
	PoolConfig := pool.NewDefaultPoolConfig()
	PoolConfig.MaxTotal = 10
	WithAbandonedConfig := pool.NewDefaultAbandonedConfig()
	androidPool = pool.NewObjectPoolWithAbandonedConfig(pool.NewPooledObjectFactorySimple(
		func() (interface{}, error) {
			conn, err := amqp.DialConfig(url, config)
			if err != nil {
				return nil, err
			}
			ch, err := conn.Channel()
			if err != nil {
				return nil, err
			}
			queue, err := ch.QueueDeclare(
				AndroidQueueName,
				true,
				false,
				false,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			err = ch.ExchangeDeclare(
				AndroidExchangeName,
				DefaultExchangeType,
				true,
				false,
				false,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			err = ch.QueueBind(
				AndroidQueueName,
				AndroidRoutingKey,
				AndroidExchangeName,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			return newAndroidQueue(ch, &queue), nil
		}), PoolConfig, WithAbandonedConfig)
	return androidPool, nil
}

//NewIOSChannel create a channel to transport messages,
//by passed a parameter to url used to connect the rabbitMQ
//For IOS Device
func NewIOSChannel(url string, config amqp.Config) (*pool.ObjectPool, error) {
	PoolConfig := pool.NewDefaultPoolConfig()
	PoolConfig.MaxTotal = 10
	WithAbandonedConfig := pool.NewDefaultAbandonedConfig()
	iosPool = pool.NewObjectPoolWithAbandonedConfig(pool.NewPooledObjectFactorySimple(
		func() (interface{}, error) {
			conn, err := amqp.DialConfig(url, config)
			if err != nil {
				return nil, err
			}
			ch, err := conn.Channel()
			if err != nil {
				return nil, err
			}
			queue, err := ch.QueueDeclare(
				IOSQueueName,
				true,
				false,
				false,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			err = ch.ExchangeDeclare(
				IOSExchangeName,
				DefaultExchangeType,
				true,
				false,
				false,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			err = ch.QueueBind(
				IOSQueueName,
				IOSRoutingKey,
				IOSExchangeName,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			return newIOSQueue(ch, &queue), nil
		}), PoolConfig, WithAbandonedConfig)
	return iosPool, nil
}

//NewMIChannel create a channel to transport messages,
//by passed a parameter to url used to connect the rabbitMQ
//For MI Device
func NewMIChannel(url string, config amqp.Config) (*pool.ObjectPool, error) {
	PoolConfig := pool.NewDefaultPoolConfig()
	PoolConfig.MaxTotal = 10
	WithAbandonedConfig := pool.NewDefaultAbandonedConfig()
	miPool = pool.NewObjectPoolWithAbandonedConfig(pool.NewPooledObjectFactorySimple(
		func() (interface{}, error) {
			conn, err := amqp.DialConfig(url, config)
			if err != nil {
				return nil, err
			}
			ch, err := conn.Channel()
			if err != nil {
				return nil, err
			}
			queue, err := ch.QueueDeclare(
				MIQueueName,
				true,
				false,
				false,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			err = ch.ExchangeDeclare(
				MIExchangeName,
				DefaultExchangeType,
				true,
				false,
				false,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			err = ch.QueueBind(
				MIQueueName,
				MIRoutingKey,
				MIExchangeName,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			return newMIQueue(ch, &queue), nil
		}), PoolConfig, WithAbandonedConfig)
	return miPool, nil
}

//NewHWChannel create a channel to transport messages,
//by passed a parameter to url used to connect the rabbitMQ
//For HuaWei Device
func NewHWChannel(url string, config amqp.Config) (*pool.ObjectPool, error) {
	PoolConfig := pool.NewDefaultPoolConfig()
	PoolConfig.MaxTotal = 10
	WithAbandonedConfig := pool.NewDefaultAbandonedConfig()
	hwPool = pool.NewObjectPoolWithAbandonedConfig(pool.NewPooledObjectFactorySimple(
		func() (interface{}, error) {
			conn, err := amqp.DialConfig(url, config)
			if err != nil {
				return nil, err
			}
			ch, err := conn.Channel()
			if err != nil {
				return nil, err
			}
			queue, err := ch.QueueDeclare(
				HWQueueName,
				true,
				false,
				false,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			err = ch.ExchangeDeclare(
				HWExchangeName,
				DefaultExchangeType,
				true,
				false,
				false,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			err = ch.QueueBind(
				HWQueueName,
				HWRoutingKey,
				HWExchangeName,
				false,
				nil)
			if err != nil {
				return nil, err
			}
			return newHWQueue(ch, &queue), nil
		}), PoolConfig, WithAbandonedConfig)
	return hwPool, nil
}

func publish(ex_name, rt_key string, body []byte, ch *amqp.Channel) (error) {
	err := ch.Publish(
		ex_name,
		rt_key,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
		})
	if err != nil {
		return err
	}
	return nil
}

//Publish produce a message to rabbit mq to send
func Publish(ex_name string, body []byte) (error) {
	switch ex_name {
	case AndroidExchangeName:
		obj, _ := androidPool.BorrowObject()
		if obj != nil {
			ch := obj.(*AndroidQueueChannel)
			err := publish(ch.ExchangeName, ch.RoutingKey, body, ch.Channel)
			if err != nil {
				return err
			}
			androidPool.ReturnObject(ch)
		}
	case MIExchangeName:
		obj, _ := miPool.BorrowObject()
		if obj != nil {
			ch := obj.(*MIQueueChannel)
			err := publish(ch.ExchangeName, ch.RoutingKey, body, ch.Channel)
			if err != nil {
				return err
			}
			miPool.ReturnObject(ch)
		}
	case HWExchangeName:
		obj, _ := hwPool.BorrowObject()
		if obj != nil {
			ch := obj.(*HWQueueChannel)
			err := publish(ch.ExchangeName, ch.RoutingKey, body, ch.Channel)
			if err != nil {
				return err
			}
			hwPool.ReturnObject(ch)
		}
	case IOSExchangeName:
		obj, _ := iosPool.BorrowObject()
		if obj != nil {
			ch := obj.(*IOSQueueChannel)
			err := publish(ch.ExchangeName, ch.RoutingKey, body, ch.Channel)
			if err != nil {
				return err
			}
			iosPool.ReturnObject(ch)
		}
	case DefaultExchangeName:
		obj, _ := defaultPool.BorrowObject()
		if obj != nil {
			ch := obj.(*DefaultQueueChannel)
			err := publish(ch.ExchangeName, ch.RoutingKey, body, ch.Channel)
			if err != nil {
				return err
			}
			defaultPool.ReturnObject(ch)
		}
	default:
		return errors.New("not supported device type")
	}
	return nil
}

//Receive consume a message from a specified channel mq to parse
func Receive(ex_name string, consumer Consumer) (error) {
	switch ex_name {
	case AndroidExchangeName:
		obj, _ := androidPool.BorrowObject()
		if obj != nil {
			ch := obj.(*AndroidQueueChannel)
			err := ch.Channel.Qos(
				1,     // prefetch count
				0,     // prefetch size
				false, // global
			)
			msgs, err := ch.Channel.Consume(
				ch.QueueName,
				"",
				false,
				false,
				false,
				false,
				nil)
			if err != nil {
				return err
			}
			go func() {
				for d := range msgs {
					consumer.Consumer(d.Body)
					d.Ack(false)
				}
			}()
			androidPool.ReturnObject(ch)
		}
	case MIExchangeName:
		obj, _ := miPool.BorrowObject()
		if obj != nil {
			ch := obj.(*MIQueueChannel)
			err := ch.Channel.Qos(
				1,     // prefetch count
				0,     // prefetch size
				false, // global
			)
			msgs, err := ch.Channel.Consume(
				ch.QueueName,
				"",
				false,
				false,
				false,
				false,
				nil)
			if err != nil {
				return err
			}
			go func() {
				for d := range msgs {
					consumer.Consumer(d.Body)
					d.Ack(false)
				}
			}()
			miPool.ReturnObject(ch)
		}
	case HWExchangeName:
		obj, _ := hwPool.BorrowObject()
		if obj != nil {
			ch := obj.(*HWQueueChannel)
			err := ch.Channel.Qos(
				1,     // prefetch count
				0,     // prefetch size
				false, // global
			)
			msgs, err := ch.Channel.Consume(
				ch.QueueName,
				"",
				false,
				false,
				false,
				false,
				nil)
			if err != nil {
				return err
			}
			go func() {
				for d := range msgs {
					consumer.Consumer(d.Body)
					d.Ack(false)
				}
			}()
			hwPool.ReturnObject(ch)
		}
	case IOSExchangeName:
		obj, _ := iosPool.BorrowObject()
		if obj != nil {
			ch := obj.(*IOSQueueChannel)
			err := ch.Channel.Qos(
				1,     // prefetch count
				0,     // prefetch size
				false, // global
			)
			msgs, err := ch.Channel.Consume(
				ch.QueueName,
				"",
				false,
				false,
				false,
				false,
				nil)
			if err != nil {
				return err
			}
			go func() {
				for d := range msgs {
					consumer.Consumer(d.Body)
					d.Ack(false)
				}
			}()
			iosPool.ReturnObject(ch)
		}
	case DefaultExchangeName:
		obj, _ := defaultPool.BorrowObject()
		if obj != nil {
			ch := obj.(*DefaultQueueChannel)
			err := ch.Channel.Qos(
				1,     // prefetch count
				0,     // prefetch size
				false, // global
			)
			msgs, err := ch.Channel.Consume(
				ch.QueueName,
				"",
				false,
				false,
				false,
				false,
				nil)
			if err != nil {
				return err
			}
			go func() {
				for d := range msgs {
					consumer.Consumer(d.Body)
					d.Ack(false)
				}
			}()
			defaultPool.ReturnObject(ch)
		}
	}
	return errors.New("not supported device type")
}

func newDefaultQueue(ch *amqp.Channel, q *amqp.Queue) *DefaultQueueChannel {
	queue := new(DefaultQueueChannel)
	queue.QueueName = DefaultQueueName
	queue.RoutingKey = DefaultRoutingKey
	queue.ExchangeName = DefaultExchangeName
	queue.ExchangeType = DefaultExchangeType
	queue.Channel = ch
	queue.Queue = q
	queue.Publishing = &amqp.Publishing{
		DeliveryMode: amqp.Persistent,
	}
	queue.CreateTime = time.Now()
	return queue
}

func newAndroidQueue(ch *amqp.Channel, q *amqp.Queue) *AndroidQueueChannel {
	queue := new(AndroidQueueChannel)
	queue.QueueName = AndroidQueueName
	queue.RoutingKey = AndroidRoutingKey
	queue.ExchangeName = AndroidExchangeName
	queue.ExchangeType = DefaultExchangeType
	queue.Channel = ch
	queue.Queue = q
	queue.Publishing = &amqp.Publishing{
		DeliveryMode: amqp.Persistent,
	}
	queue.CreateTime = time.Now()
	return queue
}

func newIOSQueue(ch *amqp.Channel, q *amqp.Queue) *IOSQueueChannel {
	queue := new(IOSQueueChannel)
	queue.QueueName = IOSQueueName
	queue.RoutingKey = IOSRoutingKey
	queue.ExchangeName = IOSExchangeName
	queue.ExchangeType = DefaultExchangeType
	queue.Channel = ch
	queue.Queue = q
	queue.Publishing = &amqp.Publishing{
		DeliveryMode: amqp.Persistent,
	}
	queue.CreateTime = time.Now()
	return queue
}

func newMIQueue(ch *amqp.Channel, q *amqp.Queue) *MIQueueChannel {
	queue := new(MIQueueChannel)
	queue.QueueName = MIQueueName
	queue.RoutingKey = MIRoutingKey
	queue.ExchangeName = MIExchangeName
	queue.ExchangeType = DefaultExchangeType
	queue.Channel = ch
	queue.Queue = q
	queue.Publishing = &amqp.Publishing{
		DeliveryMode: amqp.Persistent,
	}
	queue.CreateTime = time.Now()
	return queue
}

func newHWQueue(ch *amqp.Channel, q *amqp.Queue) *HWQueueChannel {
	queue := new(HWQueueChannel)
	queue.QueueName = HWQueueName
	queue.RoutingKey = HWRoutingKey
	queue.ExchangeName = HWExchangeName
	queue.ExchangeType = DefaultExchangeType
	queue.Channel = ch
	queue.Queue = q
	queue.Publishing = &amqp.Publishing{
		DeliveryMode: amqp.Persistent,
	}
	queue.CreateTime = time.Now()
	return queue
}