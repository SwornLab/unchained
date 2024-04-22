package consts

import "errors"

var (
	ErrInvalidKosk             = errors.New("kosk.invalid")
	ErrInvalidConfig           = errors.New("conf.invalid")
	ErrKosk                    = errors.New("kosk.error")
	ErrMissingHello            = errors.New("hello.missing")
	ErrMissingKosk             = errors.New("kosk.missing")
	ErrInternalError           = errors.New("internal_error")
	ErrCantVerifyBls           = errors.New("cant_verify_bls")
	ErrInvalidSignature        = errors.New("signature.invalid")
	ErrNotSupportedDataset     = errors.New("dataset not supported")
	ErrNotSupportedInstruction = errors.New("instruction not supported")
	ErrCantLoadSecret          = errors.New("can't load secrets")
	ErrCantLoadConfig          = errors.New("can't load config")
	ErrCantWriteSecret         = errors.New("can't write secrets")
	ErrTokenNotSupported       = errors.New("token not supported")
	ErrEventNotSupported       = errors.New("event not supported")
	ErrTopicNotSupported       = errors.New("topic not supported")
	ErrDataTooOld              = errors.New("data too old")
	ErrCantAggregateSignatures = errors.New("can't aggregate signatures")
	ErrCantRecoverSignature    = errors.New("can't recover signature")
	ErrClientNotFound          = errors.New("client not found")
	ErrSignatureNotfound       = errors.New("signature not found")
	ErrCantLoadLastBlock       = errors.New("can't load last block")
	ErrDuplicateSignature      = errors.New("duplicate signature")
	ErrCrossPriceIsNotZero     = errors.New("cross price is not zero")
	ErrAlreadySynced           = errors.New("already synced")
)
