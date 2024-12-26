package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	nodeURL      = "https://sepolia.infura.io/v3/67feff665bfb44ba8a238eb13314f81b"
	privateKey   = "xxxxxxxx"
	recipient    = "0xD54757c19d677b867cFc6A9bb86152B59e61550A"
	providedAddr = "0x222F753BEB642DDb7572A0FaD3d73afA57de9856"
)

func main() {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Fetch native currency balance and transfer
	balance, err := client.BalanceAt(ctx, common.HexToAddress(providedAddr), nil)
	if err != nil {
		log.Fatalf("Failed to fetch balance: %v", err)
	}
	fmt.Printf("Native currency balance: %s\n", balance.String())
	if balance.Cmp(big.NewInt(0)) > 0 {
		err = transferNativeCurrency(client, recipient, balance)
		if err != nil {
			log.Fatalf("Failed to transfer native currency: %v", err)
		}
		fmt.Println("Native currency transferred!")
	}
}

// Transfer native currency (ETH)
func transferNativeCurrency(client *ethclient.Client, recipient string, amount *big.Int) error {
	privateKeyECDSA, err := keystore.DecryptKey([]byte(privateKey), "")
	if err != nil {
		return fmt.Errorf("failed to load private key: %w", err)
	}

	fromAddr := privateKeyECDSA.Address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		return fmt.Errorf("failed to fetch nonce: %w", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch gas price: %w", err)
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(recipient), amount, 21000, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch chain ID: %w", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKeyECDSA.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	fmt.Printf("Native currency transfer tx sent: %s\n", signedTx.Hash().Hex())
	return nil
}
