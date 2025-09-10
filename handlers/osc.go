package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hypebeast/go-osc/osc"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type AngleData struct {
	Alpha float64 `json:"alpha"`
	Beta  float64 `json:"beta"`
	Gamma float64 `json:"gamma"`
}

var oscClient *osc.Client

func init() {
	oscClient = osc.NewClient("127.0.0.1", 9000)
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket connection established")

	for {
		var angleData AngleData
		err := conn.ReadJSON(&angleData)
		if err != nil {
			log.Printf("Error reading JSON: %v", err)
			break
		}

		sendOSCData(angleData)
	}
}

func sendOSCData(data AngleData) {
	msg := osc.NewMessage("/usercamera/Pose")
	msg.Append(float32(0.0))
	msg.Append(float32(1.1))
	msg.Append(float32(0.0))
	msg.Append(float32(-data.Beta))
	msg.Append(float32(-data.Alpha))
	msg.Append(float32(-data.Gamma))

	err := oscClient.Send(msg)
	if err != nil {
		log.Printf("Error sending OSC message: %v", err)
	} else {
		log.Printf("OSC sent - Alpha: %.1f, Beta: %.1f, Gamma: %.1f",
			data.Alpha, data.Beta, data.Gamma)
	}
}
