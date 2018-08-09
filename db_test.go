/**
 * This is NOT a freeware, use is subject to license terms
 *
 * path   go-push/midware
 * date   2018/6/19 14:00 
 * author chenjingxiu
 */
package midware

import (
	"testing"
)

//TestGetEngine
func TestGetEngine(t *testing.T) {
	db, err := GetDB(nil)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf(db.DataSourceName(), db.DriverName())
}

//BenchmarkGetEngine
func BenchmarkGetEngine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDB(nil)
	}
}
