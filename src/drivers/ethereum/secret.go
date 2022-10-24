package ethereum

import "os"

var EthPrivKey = os.Getenv("USER_PRIVATE_KEY")
var ethContractAddress = os.Getenv("CONTRACT_ADDRESS")
var ganaHost = os.Getenv("GANA_HOST")
