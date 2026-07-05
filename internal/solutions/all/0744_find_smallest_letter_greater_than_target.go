package all

// Problem: 744
// Title: Find Smallest Letter Greater Than Target
// Category: all
// Tags: all


func nextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters)-1
	for low <= high {
		mid := low + (high-low)>>1
		if letters[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	find := letters[low%len(letters)]
	if find <= target {
		return letters[0]
	}
	return find
}
