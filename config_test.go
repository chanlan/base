/**
 * This is NOT a freeware, use is subject to license terms
 *
 * path   go-push/main
 * date   2018/9/10 17:03 
 * author chenjingxiu
 */
package config

import (
	"testing"
	"fmt"
)

func BenchmarkGetConfig(t *testing.B) {
	for i := 0; i < t.N; i++ {
		c := GetConfig()
		fmt.Println(c.Database.Default.Driver)
	}
}
