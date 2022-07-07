package ethereum

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"os"
	"user/src/domains/entities"
	audit "user/src/drivers/ethereum/audit"
	content "user/src/drivers/ethereum/content"
	param "user/src/drivers/ethereum/param"
	pubkey "user/src/drivers/ethereum/pubkey"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnectContentNetWork() (*content.Contracts, *ethclient.Client) {
	contractAddress := os.Getenv("CONTENT_ADDRESS")
	ganaHost := os.Getenv("GANA_HOST")
	client, err := ethclient.Dial(ganaHost)
	if err != nil {
		panic(err)
	}
	conn, err := content.NewContracts(common.HexToAddress(contractAddress), client)
	if err != nil {
		panic(err)
	}
	return conn, client
}

func ConnectAuditNetWork() (*audit.Contracts, *ethclient.Client) {
	contractAddress := os.Getenv("AUDIT_ADDRESS")
	ganaHost := os.Getenv("GANA_HOST")
	client, err := ethclient.Dial(ganaHost)
	if err != nil {
		panic(err)
	}
	conn, err := audit.NewContracts(common.HexToAddress(contractAddress), client)
	if err != nil {
		panic(err)
	}
	return conn, client
}

func ConnectPubkeyNetWork() (*pubkey.Contracts, *ethclient.Client) {
	contractAddress := os.Getenv("PUBKEY_ADDRESS")
	ganaHost := os.Getenv("GANA_HOST")
	client, err := ethclient.Dial(ganaHost)
	if err != nil {
		panic(err)
	}
	conn, err := pubkey.NewContracts(common.HexToAddress(contractAddress), client)
	if err != nil {
		panic(err)
	}
	return conn, client
}

func ConnectParamNetWork() (*param.Contracts, *ethclient.Client) {
	contractAddress := os.Getenv("PARAM_ADDRESS")
	ganaHost := os.Getenv("GANA_HOST")
	client, err := ethclient.Dial(ganaHost)
	if err != nil {
		panic(err)
	}
	conn, err := param.NewContracts(common.HexToAddress(contractAddress), client)
	if err != nil {
		panic(err)
	}
	return conn, client
}

func GetUserAddress(privKey string) common.Address {
	userPrivKeyStr := os.Getenv("SP_PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(userPrivKeyStr)
	if err != nil {
		log.Fatal("error!!")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	Address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return Address
}

func AuthUser(client *ethclient.Client, privKey string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(30000000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

func GetParam() (*entities.Param, error) {
	conn, _ := ConnectParamNetWork()
	p, err := conn.GetParam(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return &entities.Param{
		Pairing: p.Pairing,
		G:       []byte(p.G),
		U:       []byte(p.U),
	}, nil
}
