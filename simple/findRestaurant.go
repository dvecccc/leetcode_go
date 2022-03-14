package main

import (
	"fmt"
	"math"
)

/*func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) > len(list2) {
		return findRestaurant(list2, list1)
	}
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	m3 := make(map[string]int)
	for index, value := range list1 {
		m1[value] = index
	}
	for index, value := range list2 {
		m2[value] = index
	}
	minIndex := math.MaxInt32
	for i := 0; i < len(list1); i++ {
		v1, ok1 := m1[list1[i]]
		v2, ok2 := m2[list1[i]]
		if ok1 && ok2 {
			m3[list1[i]] = v1 + v2
			if minIndex > v1 + v2 {
				minIndex = v1 + v2
			}
		}
	}
	if len(m3) == 0 {
		return nil
	}
	res := make([]string, 0)
	for k, v := range m3 {
		if v == minIndex {
			res = append(res, k)
		}
	}
	return res
}*/

func findRestaurant(list1, list2 []string) (res []string) {
	m := make(map[string]int, len(list1))
	for idx, v := range list1 {
		m[v] = idx
	}
	minIndex := math.MaxInt32
	for idx, v := range list2 {
		if i, ok := m[v]; ok {
			if idx + i < minIndex {
				minIndex = i + idx
				res = []string{v}
			} else if idx + i == minIndex {
				res = append(res, v)
			}
		}
	}
	return
}

func main()  {
	list1 := []string{"Shogun","Tapioca Express","Burger King","KFC"}
	list2 := []string{"KFC","Burger King","Tapioca Express","Shogun"}
	res := findRestaurant(list1, list2)
	fmt.Println(res)
}
