package main

import (
	"fmt"
	"math"
	"strconv"
)

func primeKeX(number int) (prime int) {

	var index int
	i := 2

	for i >= 2 {
		isPrime := true
		sqrtn := int(math.Sqrt(float64(i)))
		j := 2
		for j <= sqrtn {
			if i%j == 0 {
				isPrime = false
				j = i
			}
			j++
		}

		if isPrime {
			index++
		}

		if index == number {
			return i
		}

		i++
	}
	return i
}

func fibonacci(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		return fibonacci(number-1) + fibonacci(number-2)
	}

}

func mapOfPrime(high, wide, start int) map[int]int {

	var index int
	indexToPrime := make(map[int]int)

	i := start + 1

	for i >= start+1 {
		isPrime := true
		sqrtn := int(math.Sqrt(float64(i)))
		j := 2
		for j <= sqrtn {
			if i%j == 0 {
				isPrime = false
				j = i
			}
			j++
		}

		if isPrime {
			index++
			indexToPrime[index] = i
		}

		if index == high*wide {
			break
		}

		i++
	}

	return indexToPrime
}

func primaSegiEmpat(high, wide, start int) {
	fmt.Printf("\n\n")

	var sum int
	mapOfPrime := mapOfPrime(high, wide, start)
	j := 1
	for i := 1; i <= wide; i++ {
		for j <= i*high {
			fmt.Printf("%d\t", mapOfPrime[j])
			j++
		}
		fmt.Printf("\n\n\n")
	}

	for _, v := range mapOfPrime {
		sum += v
	}

	println(sum)

	fmt.Printf("\n")
}

func sumSubArray(arr []int, lindex int, rindex int) int {
	var sum int
	for i := lindex; i <= rindex; i++ {
		sum += arr[i]
	}
	return sum
}

func MaxSequence(arr []int) int {

	var maxSum int = arr[0]
	var subArray int

	for i := range arr {
		for j := len(arr) - 1; j > i; j-- {
			subArray = sumSubArray(arr, i, j)
			if maxSum < subArray {
				maxSum = subArray
			}
		}
	}

	return maxSum
}

func FindMinAndMax(arr []int) string {
	var result string
	min, indexMin := arr[0], 0
	max, indexMax := arr[0], 0

	for i, v := range arr {
		if v < min {
			min = v
			indexMin = i
		}
		if v > max {
			max = v
			indexMax = i
		}
	}

	result = "min: " + strconv.Itoa(min) + " index: " + strconv.Itoa(indexMin) + " max: " + strconv.Itoa(max) + " index: " + strconv.Itoa(indexMax)
	return result
}

func FindMax(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return max

}

func MaximumBuyProduct(money int, productPrice []int) {

	var maxProductBought []int

	for i := 0; i < len(productPrice); i++ {
		bought := productPrice[i]
		productBought := 1

		if bought > money {
			bought -= productPrice[i]
			productBought--
		} else {
			for j := i + 1; j < len(productPrice); j++ {
				bought += productPrice[j]
				productBought++
				if bought > money {
					bought -= productPrice[j]
					productBought--
				}
			}
			maxProductBought = append(maxProductBought, productBought)
		}
	}

	fmt.Println(strconv.Itoa(FindMax(maxProductBought)))
}

func sumEl(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}

	return sum
}

func canOut(hand, deck []int) bool {
	if len(hand) != len(deck) {
		return false
	}

	for _, v := range hand {
		for _, w := range deck {
			if v == w {
				return true
			}
		}
	}

	return false
}

func playingDomino(cards [][]int, deck []int) interface{} {

	var iPossOutCard []int

	for i, v := range cards {
		if canOut(v, deck) {
			iPossOutCard = append(iPossOutCard, i)
		}
	}

	if len(iPossOutCard) == 0 {
		return "tutup kartu"
	}

	outCard := cards[iPossOutCard[0]]

	for _, v := range iPossOutCard {
		if sumEl(cards[v]) > sumEl(outCard) {
			outCard = cards[v]
		}
	}

	return outCard
}

// Most Appear Item
type pair interface{}
type appear struct {
	item  string
	count int
}

func MostAppearItem(items []string) []pair {

	itemToCount := make(map[string]int)

	for _, v := range items {
		itemToCount[v]++
	}

	sortedResult := make([]appear, len(itemToCount))

	index := 0
	for k, v := range itemToCount {
		sortedResult[index].item = k
		sortedResult[index].count = v
		index++
	}

	for i := range sortedResult {
		for j := i + 1; j < len(sortedResult); j++ {
			if sortedResult[i].count > sortedResult[j].count {
				temp := sortedResult[i]
				sortedResult[i] = sortedResult[j]
				sortedResult[j] = temp
			}
		}
	}

	var output []pair

	for _, v := range sortedResult {
		output = append(output, v.item+"->"+strconv.Itoa(v.count))
	}

	return output
}

func main() {

	fmt.Println(primeKeX(1))  // 2
	fmt.Println(primeKeX(5))  // 11
	fmt.Println(primeKeX(8))  // 19
	fmt.Println(primeKeX(9))  // 23
	fmt.Println(primeKeX(10)) // 29

	// Fibonacci
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144

	// PrismaSegiEmpat
	primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)

	// TotalMaksimum deret
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
	fmt.Println(MaxSequence([]int{-2, -5, -3, -2, -3, -1, -6, -6}))

	// Find Minimum Maximum
	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))

	// MaximumBuyProduct
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})      // 3
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})   // 4
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})           // 1
	MaximumBuyProduct(0, []int{10000, 30000})                        // 0

	// Playing Domino
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))

	// Most Appear Item
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
