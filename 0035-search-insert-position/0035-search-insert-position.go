func searchInsert(nums []int, target int) int {
    for i, num := range nums {
        if i == 0 && target <= num{
            return 0
        }

        if i + 1 >= len(nums){
            return len(nums)
        }

        if target > num && target <= nums[i+1] {
            return i + 1
        }
    }

    return len(nums)
}