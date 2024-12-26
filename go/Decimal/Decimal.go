package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	transferAmount := "1000.000000000000000001"
	fmt.Printf("Transfer Amount (ETH): %q\n", transferAmount)

	dec, err := decimal.NewFromString(transferAmount)
	if err != nil {
		fmt.Println("Error converting string to decimal:", err)
		return
	}
	fmt.Printf("Decimal Value: %v\n", dec)

	// Multiply by 10^18 to convert to Wei (smallest unit)
	precision := decimal.NewFromInt(1).Shift(18) // 10^18 (for ETH)
	weiValue := dec.Mul(precision).BigInt()
	fmt.Printf("Wei Value (big.Int): %v\n", weiValue)

	revertedDec := decimal.NewFromBigInt(weiValue, -18)
	revertedStr := revertedDec.String()
	fmt.Printf("Reverted ETH Value: %q\n", revertedStr)
}
