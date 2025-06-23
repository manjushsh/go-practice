package linear

type entry struct {
	key   string
	value any
	next  *entry
}

type HashTable struct {
	buckets []*entry
	size    int
}

// hash function for strings
func hash(key string, bucketCount int) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = 31*hash + int(key[i])
	}
	if hash < 0 {
		hash = -hash
	}
	return hash % bucketCount
}

func NewHashTable() *HashTable {
	const defaultBuckets = 16
	return &HashTable{
		buckets: make([]*entry, defaultBuckets),
	}
}

func (ht *HashTable) Insert(key string, value any) {
	idx := hash(key, len(ht.buckets))
	for e := ht.buckets[idx]; e != nil; e = e.next {
		if e.key == key {
			e.value = value
			return
		}
	}
	ht.buckets[idx] = &entry{key, value, ht.buckets[idx]}
	ht.size++
}

func (ht *HashTable) Get(key string) (any, bool) {
	idx := hash(key, len(ht.buckets))
	for e := ht.buckets[idx]; e != nil; e = e.next {
		if e.key == key {
			return e.value, true
		}
	}
	return nil, false
}

func (ht *HashTable) Remove(key string) {
	idx := hash(key, len(ht.buckets))
	var prev *entry
	for e := ht.buckets[idx]; e != nil; e = e.next {
		if e.key == key {
			if prev == nil {
				ht.buckets[idx] = e.next
			} else {
				prev.next = e.next
			}
			ht.size--
			return
		}
		prev = e
	}
}

func (ht *HashTable) Exists(key string) bool {
	_, exists := ht.Get(key)
	return exists
}

func (ht *HashTable) Keys() []string {
	keys := make([]string, 0, ht.size)
	for _, bucket := range ht.buckets {
		for e := bucket; e != nil; e = e.next {
			keys = append(keys, e.key)
		}
	}
	return keys
}

func (ht *HashTable) Values() []any {
	values := make([]any, 0, ht.size)
	for _, bucket := range ht.buckets {
		for e := bucket; e != nil; e = e.next {
			values = append(values, e.value)
		}
	}
	return values
}

func (ht *HashTable) Clear() {
	ht.buckets = make([]*entry, len(ht.buckets))
	ht.size = 0
}

func (ht *HashTable) Size() int {
	return ht.size
}
