package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

func HandleWebSocketConnection(uid string, c *websocket.Conn) {
	defer c.Close() // close if the function gets finished

	// if there already is an connection then skip this part
	if _, ok := Connections[uid]; ok {
		log.Printf("[WRN][%s] ws connection does already exist\n", uid)
		return
	}

	// add connection to pool
	Connections[uid] = c

	defer func() {
		delete(Connections, uid) // deleting conn from pool
		CreateWSConnections()    // trying to reconnect to closed butdesired connections
	}()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Printf("[WRN][%s] message received with error: %v\n", uid, err)
			return
		}

		// handle the incomming messages
		switch mt {
		case websocket.CloseAbnormalClosure: // --------------------------------------------- CloseAbnormalClosure
			log.Printf("[INF][%s][%d]CloseAbnormalClosure received\n", uid, mt)
			break
		case websocket.CloseGoingAway: // --------------------------------------------------- CloseGoingAway
			log.Printf("[INF][%s][%d]CloseGoingAway received\n", uid, mt)
			break
		case websocket.CloseInternalServerErr: // ------------------------------------------- CloseInternalServerErr
			log.Printf("[INF][%s][%d]CloseInternalServerErr received\n", uid, mt)
			break
		case websocket.CloseInvalidFramePayloadData: // ------------------------------------- CloseInvalidFramePayloadData
			log.Printf("[INF][%s][%d]CloseInvalidFramePayloadData received\n", uid, mt)
			break
		case websocket.CloseMandatoryExtension: // ------------------------------------------ CloseMandatoryExtension
			log.Printf("[INF][%s][%d]CloseMandatoryExtension received\n", uid, mt)
			break
		case websocket.CloseMessage: // ----------------------------------------------------- CloseMessage
			log.Printf("[INF][%s][%d]CloseMessage received\n", uid, mt)
			break
		case websocket.CloseMessageTooBig: // ----------------------------------------------- CloseMessageTooBig
			log.Printf("[INF][%s][%d]CloseMessageTooBig received\n", uid, mt)
			break
		case websocket.CloseNoStatusReceived: // -------------------------------------------- CloseNoStatusReceived
			log.Printf("[INF][%s][%d]CloseNoStatusReceived received\n", uid, mt)
			break
		case websocket.CloseNormalClosure: // ----------------------------------------------- CloseNormalClosure
			log.Printf("[INF][%s][%d]CloseNormalClosure received\n", uid, mt)
			break
		case websocket.ClosePolicyViolation: // --------------------------------------------- ClosePolicyViolation
			log.Printf("[INF][%s][%d]ClosePolicyViolation received\n", uid, mt)
			break
		case websocket.CloseProtocolError: // ----------------------------------------------- CloseProtocolError
			log.Printf("[INF][%s][%d]CloseProtocolError received\n", uid, mt)
			break
		case websocket.CloseServiceRestart: // ---------------------------------------------- CloseServiceRestart
			log.Printf("[INF][%s][%d]CloseServiceRestart received\n", uid, mt)
			break
		case websocket.CloseTLSHandshake: // ------------------------------------------------ CloseTLSHandshake
			log.Printf("[INF][%s][%d]CloseTLSHandshake received\n", uid, mt)
			break
		case websocket.CloseTryAgainLater: // ----------------------------------------------- CloseTryAgainLater
			log.Printf("[INF][%s][%d]CloseTryAgainLater received\n", uid, mt)
			break
		case websocket.CloseUnsupportedData: // --------------------------------------------- CloseUnsupportedData
			log.Printf("[INF][%s][%d]CloseUnsupportedData received\n", uid, mt)
			break
		case websocket.BinaryMessage: // ---------------------------------------------------- BinaryMessage
			log.Printf("[INF][%s][%d]BinaryMessage received\n", uid, mt)
			break
		case websocket.TextMessage: // ------------------------------------------------------ TextMessage
			log.Printf("[INF][%s][%d]TextMessage received\n", uid, mt)
			break
		case websocket.PingMessage: // ------------------------------------------------------ PingMessage
			log.Printf("[INF][%s][%d]PingMessage received\n", uid, mt)
			if err := c.WriteMessage(websocket.PongMessage, []byte("")); err != nil {
				log.Printf("[ERR][%s][%d]PingMessage error: %v\n", uid, mt, err)
			}
			break
		case websocket.PongMessage: // ------------------------------------------------------ PongMessage
			log.Printf("[INF][%s][%d]PongMessage received\n", uid, mt)
			break
		default:
			log.Printf("[WRN][%s][%d] unknown messagetype\n", uid, mt)
			break
		}

		log.Printf("[INF][%s][%d] content: %s\n", uid, mt, string(msg))
	}
}
