package main

type LRUCache struct {
	Capacity int
	Cache    map[int]int
	Keys     []int
	KMap     map[int]int
}

func Constructor(capacity int) LRUCache {
	m := make(map[int]int)
	keys := make([]int, capacity)
	kMap := make(map[int]int)
	cache := LRUCache{Capacity: capacity, Cache: m, Keys: keys, KMap: kMap}
	return cache
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.Cache[key]; ok {
		index := this.KMap[key]
		for i := index; i < this.Capacity-1; i++ {
			this.Keys[i] = this.Keys[i+1]
			this.KMap[this.Keys[i]] = i
		}
		this.KMap[key] = this.Capacity - 1
		this.Keys[this.Capacity-1] = key
		return v
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	this.Cache[key] = value
	exists := false
	if _, ok := this.Cache[key]; ok {
		exists = true
	}
	if _, ok := this.Cache[this.Keys[0]]; ok && len(this.Cache) > this.Capacity {
		delete(this.Cache, this.Keys[0])
		delete(this.KMap, this.Keys[0])
	}
	if exists {
		for i := this.KMap[key]; i < this.Capacity-1; i++ {
			this.Keys[i] = this.Keys[i+1]
			this.KMap[this.Keys[i]] = i
		}
	} else {
		for i := 0; i < this.Capacity-1; i++ {
			this.Keys[i] = this.Keys[i+1]
			this.KMap[this.Keys[i]] = i
		}
	}
	this.KMap[key] = this.Capacity - 1
	this.Keys[this.Capacity-1] = key
}

//func main() {
//	funcs := []string{"LRUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"};
//	params := [][]int{{10}, {10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26}}
//
//	cache := Constructor(0)
//	for i, v := range funcs {
//		if v == "LRUCache" {
//			cache = Constructor(params[i][0])
//		} else if v == "put" {
//			cache.Put(params[i][0], params[i][1])
//		} else if v == "get" {
//			cache.Get(params[i][0])
//			fmt.Println(cache.Keys)
//		}
//	}
//}
