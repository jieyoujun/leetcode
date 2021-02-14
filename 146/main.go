package lrucache

type page struct {
	key   int
	value int
	prev  *page
	next  *page
}

type LRUCache struct {
	capacity int
	size     int
	pages    map[int]*page
	front    *page
	rear     *page
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		pages:    make(map[int]*page),
	}
}

func (cache *LRUCache) Get(key int) int {
	p, ok := cache.pages[key]
	if !ok {
		return -1
	}

	// p就是最新页面或唯一页面
	if cache.front == p || (p.prev == nil && p.next == nil) {
		return p.value
	}

	cache.bring2Front(p)
	return p.value
}

func (cache *LRUCache) Put(key int, value int) {
	p, ok := cache.pages[key]
	if ok {
		p.value = value
		cache.bring2Front(p)
		return
	}

	p = &page{
		key:   key,
		value: value,
	}
	cache.pages[key] = p

	// cache未满
	if cache.size < cache.capacity {
		if cache.front == nil && cache.rear == nil {
			// cache空
			cache.front = p
			cache.rear = p
		} else {
			// 将p放入cache头
			cache.front.prev = p
			p.next = cache.front
			cache.front = p
		}
		cache.size++
		return
	}

	// cache满
	// 1.去掉最后一个
	lrup := cache.pages[cache.rear.key]
	delete(cache.pages, cache.rear.key)

	// 最后一个是唯一的页面
	if lrup.next == nil && lrup.prev == nil {
		cache.front = p
		cache.rear = p
		return
	}

	cache.rear = cache.rear.prev
	cache.rear.next = nil
	cache.front.prev = p
	p.next = cache.front
	cache.front = p
	return
}

func (cache *LRUCache) bring2Front(p *page) {
	// p已在最前
	if cache.front == p {
		return
	}

	// p抽出来
	if p.next == nil {
		// p在最后
		cache.rear = p.prev
	} else {
		// p不在最后
		p.next.prev = p.prev
	}
	p.prev.next = p.next

	// 重置p
	p.next = cache.front
	p.prev = nil
	// 重置cache.front
	cache.front.prev = p
	cache.front = p
}
