package websocket

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/KenshiTech/unchained/internal/constants"
	"github.com/KenshiTech/unchained/internal/constants/opcodes"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/handler"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/store"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func WithWebsocket() func() {
	return func() {
		log.Logger.Info("Websocket is activated")

		versionedRoot := fmt.Sprintf("/%s", constants.ProtocolVersion)
		http.HandleFunc(versionedRoot, multiplexer)
	}
}

func multiplexer(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Logger.Error("Can't upgrade connection: %v", err)
		return
	}

	defer store.Signers.Delete(conn)
	defer store.Challenges.Delete(conn)
	defer store.Consumers.Delete(conn)
	defer store.BroadcastMutex.Delete(conn)

	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			log.Logger.Error("Can't read message: %v", err)

			err := conn.Close()
			if err != nil {
				log.Logger.Error("Can't close connection: %v", err)
			}

			break
		}

		if len(payload) == 0 {
			continue
		}

		switch opcodes.OpCode(payload[0]) {
		case opcodes.Hello:
			log.Logger.With("IP", conn.RemoteAddr().String()).Info("New Client Registered")
			result, err := handler.Hello(conn, payload[1:])
			if err != nil {
				handler.SendError(conn, messageType, opcodes.Error, err)
				continue
			}

			handler.SendMessage(conn, messageType, opcodes.Feedback, "conf.ok")
			handler.Send(conn, messageType, opcodes.KoskChallenge, result)
		case opcodes.PriceReport:
			result, err := handler.PriceReport(conn, payload[1:])
			if err != nil {
				handler.SendError(conn, messageType, opcodes.Error, err)
				continue
			}

			handler.BroadcastPayload(opcodes.PriceReportBroadcast, result)
			handler.SendMessage(conn, messageType, opcodes.Feedback, "signature.accepted")
		case opcodes.EventLog:
			result, err := handler.EventLog(conn, payload[1:])
			if err != nil {
				handler.SendError(conn, messageType, opcodes.Error, err)
				continue
			}

			handler.BroadcastPayload(opcodes.EventLogBroadcast, result)
			handler.SendMessage(conn, messageType, opcodes.Feedback, "signature.accepted")

		case opcodes.CorrectnessReport:
			result, err := handler.CorrectnessRecord(conn, payload[1:])
			if err != nil {
				handler.SendError(conn, messageType, opcodes.Error, err)
				continue
			}

			handler.BroadcastPayload(opcodes.CorrectnessReportBroadcast, result)
			handler.SendMessage(conn, messageType, opcodes.Feedback, "signature.accepted")

		case opcodes.KoskResult:
			err := handler.Kosk(conn, payload[1:])
			if err != nil {
				handler.SendError(conn, messageType, opcodes.Error, err)
				continue
			}

			handler.SendMessage(conn, messageType, opcodes.Feedback, "kosk.ok")

		case opcodes.RegisterConsumer:
			log.Logger.With("IP", conn.RemoteAddr().String()).Info("New Consumer registered")

			// TODO: Consumers must specify what they're subscribing to
			store.Consumers.Store(conn, true)
			store.BroadcastMutex.Store(conn, new(sync.Mutex))
		default:
			handler.SendError(conn, messageType, opcodes.Error, constants.ErrNotSupportedInstruction)
		}
	}
}
