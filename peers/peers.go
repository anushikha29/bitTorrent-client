//step 3 - decoding the ip address info we have received from the tracker (from binary to human readable)
package peers

import (
	"fmt"
	"net"
	"strconv"
	"encoding/binary"
	"strings"
)

//structure that holds the IP address of the peer 
type Peer struct {
	IP net.IP //IP address
	Port uint16 //port number
}

//ip address consists of 4 bytes for IP address and 2 bytes for port number
func Unmarshal(peersGotten []byte)([]Peer, error){
	const peerSize = 6;
	numberOfPeers := len(peersGotten)/peerSize;
	if len(peersGotten)%peerSize !=0{
		err := fmt.Errorf("Didn't receive good peers")
		return nil, err
	}

	peers := make([]Peer, numberOfPeers)
	for i := 0; i <numberOfPeers; i ++ {
		
	}

}

