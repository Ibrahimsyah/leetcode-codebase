func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    s1 := len(nums1)
    s2 := len(nums2)
    p1, p2, m1, m2 := 0, 0, 0, 0;

    for i := 0; i <= (s1 + s2)/2; i++ {
        m2 = m1
        if p1 != s1 && p2 != s2 {
            if nums1[p1] < nums2[p2] {
                m1 = nums1[p1]
                p1++
            } else {
                m1 = nums2[p2]
                p2++
            }
        } else if p1 < s1 {
            m1 = nums1[p1]
            p1++
        } else {
            m1 = nums2[p2]
            p2++
        }
    }

    if (s1+s2)%2 == 1 {
        return float64(m1)
    }

    return float64(m1 + m2)/2
}