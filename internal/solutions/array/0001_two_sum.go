package array

// Problem: 1
// Title: Two Sum
// Category: array
// Tags: array, hash-table
//
// Description:
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

func TwoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        if j, ok := m[target-num]; ok {
            return []int{j, i}
        }
        m[num] = i
    }
    return nil
}
