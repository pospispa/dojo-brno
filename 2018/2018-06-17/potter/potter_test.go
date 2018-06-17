package potter

import (
	"sort"
	"testing"
)

const (
	bookPrice = 8
)

var discounts = [5]float64{1.0, 0.95, 0.9, 0.8, 0.75}

func convertAndRemoveEmptyItemsAndSortAscending(basket [5]int) []int {
	sortedBasket := basket[:]
	sort.Ints(sortedBasket)
	booksInBasket := make([]int, 0, 5)
	for _, amount := range sortedBasket {
		if amount > 0 {
			booksInBasket = append(booksInBasket, amount)
		}
	}
	return booksInBasket
}

func removeAmount(basket []int, amount, times int) []int {
	if len(basket) == 0 {
		return basket
	}
	reducedBasket := make([]int, 0, 5)
	for i := len(basket) - 1; i >= 0; i-- {
		if times > 0 && basket[i] > amount {
			reducedBasket = append(reducedBasket, basket[i]-amount)
			times--
		} else if times > 0 {
			times--
		} else if basket[i] > 0 {
			reducedBasket = append(reducedBasket, basket[i])
		}
	}
	sort.Ints(reducedBasket)
	return reducedBasket
}

func calculatePriceForSeries(basket []int, seriesLength, numberOfSeries int) (float64, []int) {
	if seriesLength <= 0 || basket[0] < numberOfSeries {
		return 0.0, basket
	}
	priceForSeries := discounts[seriesLength-1] * float64(bookPrice*seriesLength*numberOfSeries)
	reducedBasket := removeAmount(basket, numberOfSeries, seriesLength)
	return priceForSeries, reducedBasket
}

func hasAtLeastTwoQuadruples(basket []int) bool {
	return len(removeAmount(basket, 1, 4)) >= 4
}

func Potter(basket [5]int) float64 {
	var totalPrice, priceForSeries float64
	booksInBasket := convertAndRemoveEmptyItemsAndSortAscending(basket)
	for len(booksInBasket) == 5 && hasAtLeastTwoQuadruples(booksInBasket) {
		priceForSeries, booksInBasket = calculatePriceForSeries(booksInBasket, 4, 1)
		totalPrice += priceForSeries
		priceForSeries, booksInBasket = calculatePriceForSeries(booksInBasket, 4, 1)
		totalPrice += priceForSeries
	}
	for len(booksInBasket) > 0 {
		priceForSeries, booksInBasket = calculatePriceForSeries(booksInBasket, len(booksInBasket), booksInBasket[0])
		totalPrice += priceForSeries
	}

	return totalPrice
}

func TestPotter(t *testing.T) {
	tests := []struct {
		basket [5]int
		want   float64
	}{
		{[5]int{0, 0, 0, 0, 0}, 0},
		{[5]int{0, 0, 1, 0, 0}, 8},
		{[5]int{0, 0, 1, 0, 1}, 15.20},
		{[5]int{0, 0, 1, 0, 2}, 23.20},
		{[5]int{0, 1, 1, 0, 1}, 21.60},
		{[5]int{1, 1, 1, 0, 1}, 25.60},
		{[5]int{1, 1, 1, 1, 1}, 30},
		{[5]int{1, 1, 2, 2, 2}, 51.20},
		{[5]int{2, 2, 2, 2, 2}, 66.40},
		{[5]int{4, 4, 4, 4, 4}, 128},
	}

	for _, tt := range tests {
		if got := Potter(tt.basket); got != tt.want {
			t.Errorf("Potter(%v) = %v WANT %v", tt.basket, got, tt.want)
		}
	}
}
