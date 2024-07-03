func accountBalanceAfterPurchase(purchaseAmount int) int {
    lastDigit := purchaseAmount%10
    multiplier := purchaseAmount/10
    if lastDigit >= 5 && lastDigit != 0{
        purchaseAmount = (multiplier + 1)*10
    } else {
        purchaseAmount = multiplier * 10
    }

    return 100 - purchaseAmount
}