func minMaxGame(nums []int) int {
    for {
        if len(nums) == 1  {
            return nums[0]
        }

        temp := make([]int, len(nums) / 2)
        for i := range temp {
            f := max
            if i % 2 == 0 {
                f = min
            }
            temp[i] = f(nums[2*i], nums[2*i+1])
        }
        nums = temp
    }   
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}

func min(i, j int) int{
    if i < j {
        return i
    }
    return j
}