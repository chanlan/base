package midware

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"gopkg.in/mgo.v2"
)

type Person struct {
	Name  string
	Phone string
}

//TestMog
func TestMog(t *testing.T) {
	Exec("people", func(collection *mgo.Collection) {
		count, _ := collection.Count()
		t.Log(count)
		fmt.Println(count)
	})
}

//BenchmarkMog
func BenchmarkMog(b *testing.B) {
	Exec("people", func(collection *mgo.Collection) {
		err := collection.Insert(&Person{"user1", "+86 134640044434"},
			&Person{"user2", "+86 15210071513"})
		if err != nil {
			b.Errorf("insert error")
		}
		result := Person{}
		err = collection.Find(bson.M{"name": "user1"}).One(&result)
		if err != nil {
			b.Errorf("query error")
		}
		b.Log(result)
	})
}

