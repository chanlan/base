/**
 * This is NOT a freeware, use is subject to license terms
 *
 * path   go-push/midware
 * date   2018/6/19 17:37 
 * author chenjingxiu
 */
package midware

import (
	"gopkg.in/mgo.v2"
	"strings"
	"time"
)

const (
	poolSize = 10
	timeout  = 30 * time.Second
)

var (
	session    *mgo.Session
	newSession *mgo.Session
	err        error
)

func init() {
	newSession, err = MogConn()
	if err != nil {
		panic(err)
	}
	session = newSession.Clone()
}
//MogConn return a copy of *mgo.Session
func MogConn() (*mgo.Session, error) {
	dialInfo := &mgo.DialInfo{
		Addrs:    strings.Split(MogServers, ";"),
		Direct:   false,
		Timeout:  timeout,
		Database: MogDBName,
		Username: MogUser,
		Password: MogPwd,
	}
	newSession, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}
	newSession.SetMode(mgo.Eventual, true)
	newSession.SetPoolLimit(poolSize)
	return newSession, nil
}

//GetIns return a mgo.Collection instance used to operate mongodb
func Exec(collection string, f func(*mgo.Collection)) (error) {
	if err = session.Ping(); err != nil {
		if err = newSession.Ping(); err != nil {
			newSession, err = MogConn()
		}
		session = newSession.Clone()
	}
	defer func() {
		session.Close()
	}()
	c := session.DB(MogDBName).C(collection)
	f(c)
	return nil
}
