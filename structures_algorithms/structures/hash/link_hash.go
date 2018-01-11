package hash




const TABLE_SIZE = 2

type entry struct {
	key   int
	value interface{}
	next  *entry
}

type linkHashMap struct {
	table []*entry
}

// hash表初始化
func (hashMap *linkHashMap) checkTable() {
	if hashMap.table == nil {
		hashMap.table = make([]*entry, 0)
		for i := 0; i < TABLE_SIZE; i++ {
			hashMap.table = append(hashMap.table, nil)
		}
	}
}

func getHashValue(key int) int {
	return key % TABLE_SIZE
}

func (hashMap *linkHashMap) put(key int, value interface{}) {
	hashMap.checkTable()
	hash := getHashValue(key)
	e := hashMap.table[hash]

	// 当前下标为空时直接赋值
	if e == nil {
		hashMap.table[hash] = &entry{key, value, nil}
		return
	}

	for e.next != nil {
		if e.key == key {
			e.value = value
			return
		}
		e = e.next
	}

	if e.key == key {
		e.value = value
	} else {
		e.next = &entry{key, value, nil}
	}
}

func (hashMap *linkHashMap) get(key int) interface{} {
	hashMap.checkTable()
	hash := getHashValue(key)
	e := hashMap.table[hash]
	if e != nil {
		for e != nil {
			if e.key == key {
				return e.value
			}
			e = e.next
		}
	}
	return nil
}

func (hashMap *linkHashMap) delete(key int) {
	hashMap.checkTable()
	hash := getHashValue(key)
	e := hashMap.table[hash]
	ep := e
	if e != nil {
		if e.key == key {
			hashMap.table[hash] = nil
			return
		}
		for e.next != nil {
			if e.key == key {
				e.key = e.next.key
				e.value = e.next.value
				e.next = e.next.next
				return
			}
			ep, e = e, e.next
		}
		if e.key == key {
			ep.next = nil
		}
	}
}
