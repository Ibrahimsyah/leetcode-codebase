func searchRange(nums []int, target int) []int {
    first, last := -1, -1
    l, r := 0, len(nums) - 1
    for l <= r {
        if first != -1 && last != -1 {
            break
        }
        
        if nums[l] == target && first == -1 {
            first = l
        }

        if nums[r] == target && last == -1 {
            last = r
        }

        if nums[l] < target {
            l++
        }

        if nums[r] > target {
            r-- 
        }
    }

    return []int{first, last}
}