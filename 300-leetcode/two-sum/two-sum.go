
package main

import "fmt"

func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if j, exists := hashMap[complement]; exists {
            return []int{j, i}
        }
        hashMap[num] = i
    }
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    fmt.Println(twoSum(nums, target)) // Выведет: [0 1]
}

