package serv

//import "fmt"

type Cache struct {
	Size int
	CacheContainer map[int]int
	History map[int]int
}

func NewCache(size int) *Cache {
	return &Cache{
		Size: size,
		CacheContainer: make(map[int]int),
		History: make(map[int]int),
	}
}

func (cache Cache) UpdateHistory(key int) {
	cache.History[key]++;
}

func (cache Cache) CheckValueInCache(key int) bool {
	if len(cache.CacheContainer) == 0 {
		return false
	} else {
		for mapKey := range cache.CacheContainer {
			if mapKey == key { return true }
		}
	}
	return false
}

func (cache Cache) GetFibonValue(key int) int {
	return cache.CacheContainer[key]
}

func (cache Cache) AddNewValue(key int, value int) {
	if len(cache.CacheContainer) >= cache.Size {
		arr := make(map[int]int)
		var buff []int;
		for cacheKey := range  cache.CacheContainer {
			buff = append(buff, cache.History[cacheKey])
		}
		minValue := Min(buff)
		var buff2 []int
		for cacheKey := range  cache.CacheContainer {
			arr[cacheKey] = cache.History[cacheKey]
		}
		for cacheKey, value := range arr {
			if value == minValue {
				buff2 = append(buff2, cacheKey)
			}
		}
		valueForDrop := Min(buff2)
		delete(cache.CacheContainer, valueForDrop)
	}
	cache.CacheContainer[key] = value
}

func Min(arr []int) int {
	smallest := arr[0]
	for _, v := range arr {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}

