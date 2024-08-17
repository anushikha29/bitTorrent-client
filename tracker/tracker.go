// establishing the connection with the tracker to get the list of peers using the announce URL(step-2)

package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func getPeers(announceURL string, infoHash [20]byte, peerID string){
	params := url.Values{}
	params.Add("info_hash", string(infoHash[:]))
	params.Add("peer_id", peerID)
	params.Add("port", "6881")
	params.Add("uploaded", "0")
	params.Add("downloaded", "0")
	params.Add("left", "1000")
	params.Add("compact", "1")

	//constructing the Tracker URL with the announceURL and the query parameters 
	trackerURL := fmt.Sprintf("%s?%s", announceURL, params.Encode())

	resp, err := http.Get(trackerURL)
	if err != nil {
		fmt.Println("Error Connecting to tracker:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Tracker response error: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading tracker response: %v", err)
	}
	fmt.Println("Tracker response:", string(body))
}

func generatePeerID() string {
	// Generate a 20-byte peer ID with a -XY0001- prefix where XY are random bytes
	var peerID [20]byte
	copy(peerID[:8], "-XY0001-")
	binary.LittleEndian.PutUint64(peerID[8:], uint64(time.Now().UnixNano()))
	return string(peerID[:])
}

func main() {
	// Example usage
	announceURL := "http://example.com/announce"
	infoHash := sha1.Sum([]byte("example_info_hash"))

	peerID := generatePeerID()
	
	getPeers(announceURL, infoHash, peerID)
}