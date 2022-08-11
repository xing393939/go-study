package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
	"github.com/xing393939/jsonobject"
	gst "go-study/go-demo/gstreamer-src"
	"sync"
)

func main() {
	// Prepare the configuration
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:192.168.2.119:56842"},
			},
		},
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	// Set the handler when the peer has connected/disconnected
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("Connection State has changed %s \n", connectionState.String())
	})

	// Create a video track
	firstVideoTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: "video/vp8"}, "video", "pion1")
	if err != nil {
		panic(err)
	}
	_, err = peerConnection.AddTrack(firstVideoTrack)
	if err != nil {
		panic(err)
	}

	// Create a second video track
	secondVideoTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: "video/vp8"}, "video", "pion2")
	if err != nil {
		panic(err)
	}
	_, err = peerConnection.AddTrack(secondVideoTrack)
	if err != nil {
		panic(err)
	}

	// Create an offer to send to the other process
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}
	if err = peerConnection.SetLocalDescription(offer); err != nil {
		panic(err)
	}

	c, _, err := websocket.DefaultDialer.Dial("ws://192.168.2.119:8004", nil)
	if err != nil {
		panic(err)
	}

	// send ice
	var candidatesMux sync.Mutex
	pendingCandidates := make([]*webrtc.ICECandidate, 0)
	peerConnection.OnICECandidate(func(ice *webrtc.ICECandidate) {
		fmt.Println(ice)
		if ice == nil {
			return
		}
		candidatesMux.Lock()
		defer candidatesMux.Unlock()
		pendingCandidates = append(pendingCandidates, ice)
	})

	go func() {
		for {
			_, data, _ := c.ReadMessage()
			jo := jsonobject.NewJsonObject(string(data))
			if jo.GetString("type") == "playerConnected" {
				myOffer := map[string]string{
					"type":     "offer",
					"sdp":      offer.SDP,
					"playerId": jo.GetString("playerId"),
				}
				_ = c.WriteJSON(myOffer)
			} else if jo.GetString("type") == "iceCandidate" {
				ice := jo.GetJsonObject("candidate").GetString("candidate")
				iceErr := peerConnection.AddICECandidate(webrtc.ICECandidateInit{Candidate: ice})
				if iceErr != nil {
					println("iceCandidate", iceErr.Error())
				}
			} else if jo.GetString("type") == "answer" {
				sdp := webrtc.SessionDescription{}
				_ = json.Unmarshal(data, &sdp)
				if sdpErr := peerConnection.SetRemoteDescription(sdp); sdpErr != nil {
					println("answer", sdpErr.Error())
				}
				candidatesMux.Lock()
				for _, ice2 := range pendingCandidates {
					fmt.Println(ice2.ToJSON())
					myOffer := map[string]interface{}{
						"type":      "iceCandidate",
						"playerId":  jo.GetString("playerId"),
						"candidate": ice2.ToJSON(),
					}
					_ = c.WriteJSON(myOffer)
				}
				candidatesMux.Unlock()
			}
		}
	}()

	// Create channel that is blocked until ICE Gathering is complete
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)
	<-gatherComplete

	// Start pushing buffers on these tracks
	// gst.CreatePipeline("opus", []*webrtc.TrackLocalStaticSample{audioTrack}, *audioSrc).Start()
	src := "uridecodebin uri=file:///mnt/c/addons/goweb/go-study/go-demo/gstreamer-send/ma2fankong.webm ! videoscale ! video/x-raw, width=320, height=240 ! queue "
	gst.CreatePipeline("vp8", []*webrtc.TrackLocalStaticSample{firstVideoTrack, secondVideoTrack}, src).Start()
	select {}
}
