package ethereum

import "os"

var EthPrivKey = os.Getenv("USER_PRIVATE_KEY")
var ethContentAddress = os.Getenv("CONTENT_ADDRESS")
var ethAuditAddress = os.Getenv("AUDIT_ADDRESS")
var ethParamAddress = os.Getenv("PARAM_ADDRESS")
var ethPubKeyAddress = os.Getenv("PUBKEY_ADDRESS")
var ganaHost = os.Getenv("GANA_HOST")
