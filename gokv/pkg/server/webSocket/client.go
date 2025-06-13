package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jnnkrdb/gokv/conf"
)

var WsHeader = make(http.Header)

type NodePool struct {
	Nodes []string `json:"nodes"`
}

func init() {
	WsHeader.Add("gokv-node", conf.SELF_NAME)
	WsHeader.Add("gokv-uid", conf.SELF_UID)
}

// get the current nodes from the defined api
func GetNodes() []string {
	var np = NodePool{}

	var u = url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s.%s.svc.%s:%d", conf.SELF_WEBSOCKET_SERVICE_NAME, conf.SELF_NAMESPACE, conf.CLUSTER_INTERNAL_DOMAIN, conf.GOSSIP_PORT),
		Path:   "/api/v1/connections",
	}

	if resp, err := http.Get(u.String()); err != nil {

		log.Printf("[ERR] couldn't connect to [%s], are there other hosts running?: %v\n", u.String(), err)

	} else {

		if err = json.NewDecoder(resp.Body).Decode(&np); err != nil {

			log.Printf("[ERR] couldn't decode the received response body: %v\n", err)
		}

	}

	return np.Nodes
}

// the init function tries to connect to the service inside the cluster, to
// gather the current nodes in the gokv-cluster. Then tries to open a connection
// to every node
func CreateWSConnections() {

	// range through the received nodes and creat the connections
	for _, node := range GetNodes() {

		if strings.Contains(node, conf.SELF_NAME) {
			continue
		}

		// create url
		u := url.URL{
			Scheme: "ws",
			Host:   fmt.Sprintf("%s.%s.%s.svc.%s:%d", node, conf.SELF_WEBSOCKET_HEADLESS_SERVICE_NAME, conf.SELF_NAMESPACE, conf.CLUSTER_INTERNAL_DOMAIN, conf.GOSSIP_PORT),
			Path:   WebsocketPath,
		}
		log.Printf("[INF] connecting to url: %s\n", u.String())

		// trying to connect to the node, if it does not work, then retry
		var currentRetry int = 1
		for {

			if c, _, err := websocket.DefaultDialer.Dial(u.String(), WsHeader); err != nil {

				log.Printf("[ERR][try-%d] couldn't connect to [%s]: %v\n", currentRetry, u.String(), err)

			} else {

				go HandleWebSocketConnection(node, c)

				return
			}

			currentRetry++
			time.Sleep(5 * time.Second)
		}
	}
}
