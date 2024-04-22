package model

import (
	"encoding/json"

	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"

	"github.com/KenshiTech/unchained/internal/ent/helpers"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type EventLogReportPacket struct {
	EventLog
	Signature [48]byte
}

func (e *EventLogReportPacket) Sia() *sia.Sia {
	return new(sia.Sia).
		EmbedSia(e.EventLog.Sia()).
		AddByteArray8(e.Signature[:])
}

func (e *EventLogReportPacket) DeSia(sia *sia.Sia) *EventLogReportPacket {
	e.EventLog.DeSia(sia)
	copy(e.Signature[:], sia.ReadByteArray8())

	return e
}

type BroadcastEventPacket struct {
	Info      EventLog
	Signature [48]byte
	Signer    Signer
}

func (b *BroadcastEventPacket) Sia() *sia.Sia {
	return new(sia.Sia).
		EmbedSia(b.Info.Sia()).
		AddByteArray8(b.Signature[:]).
		EmbedSia(b.Signer.Sia())
}

func (b *BroadcastEventPacket) DeSia(sia *sia.Sia) *BroadcastEventPacket {
	b.Info.DeSia(sia)
	copy(b.Signature[:], sia.ReadByteArray8())
	b.Signer.DeSia(sia)

	return b
}

type EventLogArg struct {
	Name  string `json:"Name"`
	Type  string `json:"Type"`
	Value any    `json:"Value"`
}

type EventLog struct {
	LogIndex uint64
	Block    uint64
	Address  string
	Event    string
	Chain    string
	TxHash   [32]byte
	Args     []EventLogArg

	Consensus    bool
	SignersCount uint64
	SignerIDs    []int
	Signature    []byte
	Voted        *helpers.BigInt
}

func (e *EventLog) Sia() *sia.Sia {
	argsEncoded, err := json.Marshal(e.Args)

	if err != nil {
		panic(err)
	}

	logSia := new(sia.Sia).
		AddUInt64(e.LogIndex).
		AddUInt64(e.Block).
		AddString8(e.Address).
		AddString8(e.Event).
		AddString8(e.Chain).
		AddByteArray8(e.TxHash[:]).
		AddByteArray16(argsEncoded)

	return logSia
}

func (e *EventLog) DeSia(sia *sia.Sia) *EventLog {
	e.LogIndex = sia.ReadUInt64()
	e.Block = sia.ReadUInt64()
	e.Address = sia.ReadString8()
	e.Event = sia.ReadString8()
	e.Chain = sia.ReadString8()
	copy(e.TxHash[:], sia.ReadByteArray8())

	argsEncoded := sia.ReadByteArray16()
	err := json.Unmarshal(argsEncoded, &e.Args)

	if err != nil {
		panic(err)
	}

	return e
}

func (e *EventLog) Bls() (bls12381.G1Affine, error) {
	hash, err := bls.Hash(e.Sia().Content)
	if err != nil {
		utils.Logger.Error("Can't hash bls: %v", err)
		return bls12381.G1Affine{}, err
	}

	return hash, err
}
