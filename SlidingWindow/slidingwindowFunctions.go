package SlidingWindow

func GetBTTBASS(prices []int) int {
	low := prices[0]
	high := prices[0]
	sold := 0

	for _, price := range prices {
		if price > high {
			high = price

			if (high - low) > sold {
				sold = high - low
			}
		}

		if price < low {
			low = price
			high = price
		}
	}

	return sold
}
