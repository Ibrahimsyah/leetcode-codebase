func triangularSum(nums []int) int {
    for {
        if len(nums) == 1 {
            return nums[0]
        }

        temp := make([]int, len(nums) - 1)
        for i := range temp {
            temp[i] = (nums[i] + nums[i+1]) % 10
        }
        nums = temp
    }
}