package handler

import (
	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/middleware"
	"github.com/gorilla/websocket"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

// PriceReport check signature of message and return price info.
func PriceReport(conn *websocket.Conn, payload []byte) ([]byte, error) {
	err := middleware.IsConnectionAuthenticated(conn)
	if err != nil {
		return []byte{}, err
	}

	priceReport := new(model.PriceReportPacket).DeSia(&sia.Sia{Content: payload})
	priceInfoHash, err := priceReport.PriceInfo.Bls()
	if err != nil {
		return []byte{}, consts.ErrInternalError
	}

	signer, err := middleware.IsMessageValid(conn, priceInfoHash, priceReport.Signature)
	if err != nil {
		return []byte{}, err
	}

	priceInfo := model.BroadcastPricePacket{
		Info:      priceReport.PriceInfo,
		Signature: priceReport.Signature,
		Signer:    signer,
	}

	priceInfoByte := priceInfo.Sia().Content
	return priceInfoByte, nil
}
