package main

import (
	"fmt"
)

const (
	from = 0
	to   = 1
	cost = 2

	firstLevel = 1
	maxLevel   = 5
)

// Items a slice of int
type Items []int

func (i *Items) len() int {
	return len(*i)
}

func (i *Items) push(it int) {
	*i = append(*i, it)
}

func (i *Items) pop() int {
	old := *i
	n := len(old)
	item := old[n-1]
	old[n-1] = -1 // avoid memory leak
	*i = old[0 : n-1]
	return item
}

type route struct {
	from int
	to   int
	cost int
}

// createRouteMap -> create a map with `pick` or `from` as key and `route` as value
// e.g. 1,3,5000; 1,4,3000 -> m[1] = {{1,3,5000},{1,4,3000}}
func createRouteMap(routes [][]int) map[int][]route {
	m := make(map[int][]route)

	for _, v := range routes {
		temp := route{
			from: v[from],
			to:   v[to],
			cost: v[cost],
		}
		m[v[from]] = append(m[v[from]], temp)
	}

	return m
}

// concatFromTo -> concat `from` and `to` to make a string
func concatFromTo(from, to int) string {
	return fmt.Sprintf("%d%d", from, to)
}

// createCostMap -> create a map with combination of `pick` & `dest` as key and `cost` as value
// e.g. 1,3,5000 -> m["13"] = 5000
func createCostMap(routes [][]int) map[string]int {
	costMap := make(map[string]int)

	for _, route := range routes {
		key := concatFromTo(route[from], route[to])
		costMap[key] = route[cost]
	}

	return costMap
}

func getCheapestPathFromStartToDest(routes [][]int, orders [][]int) []int {
	p := make(Items, 0)

	routeMap := createRouteMap(routes)
	costMap := createCostMap(routes)

	calcCost := func() int {
		totalCost := 0

		for i := len(p) - 1; i >= 1; i-- {
			key := concatFromTo(p[i-1], p[i])
			totalCost += costMap[key]
		}

		return totalCost
	}

	var getCheapestPathFromPickToDest func(pick, dest, level int, destination *[]int)

	minCost := 0

	getCheapestPathFromPickToDest = func(pick, dest, level int, destination *[]int) {
		if level == firstLevel {
			fmt.Println("pick:", pick, "dest:", dest)
			minCost = 0
		}

		if level >= maxLevel {
			return
		}

		p.push(pick)

		if d, ok := routeMap[pick]; ok {
			for _, v := range d {
				if dest == v.to {

					//a path is found
					p.push(v.to)

					totalCost := calcCost()

					if minCost == 0 || (minCost > 0 && totalCost < minCost) {
						minCost = totalCost
						*destination = append([]int{}, p...)
					}

					fmt.Println("cost of", p, ":", totalCost)

					p.pop()
					break
				}

				getCheapestPathFromPickToDest(v.to, dest, level+1, destination)
			}
		}

		p.pop()
	}

	level := firstLevel

	destinations := make([][]int, len(orders)+1)

	result := make([]int, 0)

	for i, o := range orders {
		if i == 0 {
			getCheapestPathFromPickToDest(1, o[from], level, &destinations[i])
			result = append(result, destinations[i]...)
		}

		getCheapestPathFromPickToDest(o[from], o[to], level, &destinations[i+1])
		//append start from index 1: since the first index will be the same with the last index from prev array
		result = append(result, destinations[i+1][1:]...)
	}

	return result
}

func main() {
	cases := []struct {
		routes [][]int
		orders [][]int
	}{
		{
			routes: [][]int{{1, 2, 3000}, {1, 5, 7000}, {2, 1, 3000}, {2, 3, 2000}, {3, 5, 5000}, {4, 1, 5000}, {5, 4, 5000}, {5, 3, 5000}},
			orders: [][]int{{4, 3}},
		},
		{
			routes: [][]int{{1, 2, 3000}, {1, 5, 7000}, {2, 1, 3000}, {2, 3, 2000}, {3, 5, 5000}, {4, 1, 5000}, {5, 4, 5000}, {5, 3, 5000}, {3, 2, 1000}},
			orders: [][]int{{2, 5}, {5, 1}},
		},
	}

	for i, c := range cases {
		fmt.Printf("--- case %d ---\n", i+1)
		fmt.Println(getCheapestPathFromStartToDest(c.routes, c.orders))
		fmt.Printf("--- end of case ---\n")
	}
}
