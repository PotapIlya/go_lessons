package main

import "fmt"

type Cache []int

func (c *Cache) add(val int) {
	*c = append(*c, val)
}

func (c *Cache) delete(val int) {
	for i, v := range *c {
		if v == val {
			*c = append((*c)[:i], (*c)[i+1:]...)
			return
		}
	}
}

func (c *Cache) count() int {
	return len(*c)
}

func main() {

	cache := Cache{}

	cache.add(1)
	cache.add(2)
	cache.add(3)
	cache.add(4)

	cache.delete(2)

	fmt.Println(cache.count())

}
