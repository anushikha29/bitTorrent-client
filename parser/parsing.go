// parsing the .torrent file to get the metainfo from the torrent (step-1)

package main

import(
	"crypto/sha1"
	"fmt"
	"log"
	"os"

	"github.com/jackpal/bencode-go"
)

type TorrentFile struct {
	Announce string `bencode:"announce"`
	Info Info `bencode:"info"`
}

type Info struct {
	Name string `bencode:"name"`
	PieceLength int `bencode:"piece length"`
	Pieces string `bencode:"pieces"`
	Length int `bencode:"length"`
	Files []File `bencode:"files"`
}

type File struct {
	Length int `bencode:"length"`
	Path []string `bencode:"path"`
}

func main() {
	file, err := os.Open("e")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var torrent TorrentFile
	err = bencode.Unmarshal(file, &torrent)
	if err!= nil {
		log.Fatal(err)
	}

	fmt.Printf("Announce URL: %s\n", torrent.Announce)
    fmt.Printf("Name: %s\n", torrent.Info.Name)
    fmt.Printf("Piece Length: %d\n", torrent.Info.PieceLength)
    fmt.Printf("Length: %d\n", torrent.Info.Length)

	infoHash := sha1.New()
	err = bencode.Marshal(infoHash, torrent.Info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Info Hash: %x\n", infoHash.Sum(nil))

}




