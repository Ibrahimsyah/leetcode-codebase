func reverseBits(num uint32) uint32 {
    var res uint32 = 0
    for i := 0; i < 32; i++{
        res = res*2 + num%2
        num /=2
    }
    return res
}