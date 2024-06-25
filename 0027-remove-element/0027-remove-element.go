func removeElement(nums []int, val int) int {
    index := 0
    for _, n := range nums {
        if n != val {
            nums[index] = n
            index++
        }
    }

    return index
}