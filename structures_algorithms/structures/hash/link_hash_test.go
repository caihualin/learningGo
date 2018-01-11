package hash

import (
	"testing"
	"fmt"
)

func Test_hashMap(t *testing.T) {
	hashMap := &linkHashMap{}
	hashMap.put(0, "one")
	hashMap.put(1, "two")
	hashMap.put(2, "three")
	fmt.Printf(
		"插入: %s,%s,%s,%s\n",
		hashMap.get(0),
		hashMap.get(1),
		hashMap.get(2),
		func(who int) string {
			w := hashMap.get(who)
			if w == nil {
				return "放不下了"
			}
			return w.(string)
		}(3),
	)
	hashMap.delete(2)
	fmt.Println("删除: three =>", hashMap.get(2))
	fmt.Println("当前:", hashMap.table[0], hashMap.table[1])
}
