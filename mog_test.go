/**
 * This is NOT a freeware, use is subject to license terms
 *
 * path   go-push/midware
 * date   2018/6/21 15:09 
 * author chenjingxiu
 */
package midware

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"testing"
	"fmt"
	"gopkg.in/mgo.v2"
)

type UserMsgLog struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	U    int           `bson:"u"`
	Fmid string        `bson:"fmid"`
	Mid  int64         `bson:"mid"`
	S    string        `bson:"s"`
	P    string        `bson:"p"`
	T    int64         `bson:"t"`
	Mut  float64       `bson:"mut"`
	Sut  float64       `bson:"sut"`
	Rut  float64       `bson:"rut"`
	Aut  float64       `bson:"aut"`
	Out  float64       `bson:"out"`
	Cut  time.Time     `bson:"cut"`
}

func TestQuery(t *testing.T) {
	var group []UserMsgLog
	Exec("user_msglog", func(collection *mgo.Collection) {
		query := collection.Find(bson.M{"p": "android"})
		err := query.All(&group)
		if err != nil {
			panic(err)
		}
		for _, v := range group {
			fmt.Println(v.P, v.S)
		}
	})
}

func TestRun(t *testing.T) {
	session, err := MogConn()
	if err != nil {
		t.Errorf(err.Error())
	}
	var group bson.M
	database := session.DB("test")
	o := bson.M{
		"group": bson.M{
			"ns": "user_msglog",
			"key": bson.M{
				"p": true,
				"s": true,
			},
			"cond": bson.M{
				"cut": bson.M{"$lt": time.Now(), "$gte": time.Now().AddDate(0, 0, -90)},
			},
			"$reduce": `function(doc,prev){
				prev.total++;
				if(doc.sut>0){
					prev.sendTotal++;
				}
				if(doc.rut>0){
					prev.arrivedTotal++;
				}
				if(doc.out > 0){
					prev.openTotal++;
				}
			}`,
			"initial": bson.M{
				"total":        0,
				"sendTotal":    0,
				"arrivedTotal": 0,
				"openTotal":    0,
			},
		},
	}

	lastErr := database.Run(o, &group)
	if lastErr != nil {
		panic(lastErr)
	}

	for _, v := range group["retval"].([]interface{}) {
		fmt.Println(v.(bson.M))
	}
}
