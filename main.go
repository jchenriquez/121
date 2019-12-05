package main

import "fmt"

func maxProfit(prices []int) (result int) {
	if len(prices) == 0 {
		return
	}
	smallest := prices[0]

	for _, price := range prices {
		if price < smallest {
			smallest = price
		} else {
			if price - smallest > result {
				result = price - smallest
			}
		}
	}

	return
}

func main() {
	fmt.Printf("max profit %d\n", maxProfit([]int{}))
}
