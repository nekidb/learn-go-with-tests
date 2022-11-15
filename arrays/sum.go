package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range(numbers) {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
//	var result []int
//	
//	for _, numbers := range(numbersToSum) {
//		result = append(result, Sum(numbers))
//	}
//	return result

//	lenghtOfNumbers := len(numbersToSum)
	sums := make([]int, len(numbersToSum))
	
	for i, numbers := range(numbersToSum) {
		sums[i] = Sum(numbers)
	}
	
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range(numbersToSum) {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	
	return sums
}