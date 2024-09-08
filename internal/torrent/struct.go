package torrent

import (
	"bufio"
	"bytes"
	"encoding/base32"
	"errors"
	"io"
	"os"
)

var ErrNotParsed = errors.New("torrent not parsed")

type Piece [20]byte
type InfoHash [20]byte

type Torrent struct {
	announce string
	pieces   []Piece
	infoHash InfoHash

	reader *bufio.Reader

	parsed bool
}

func (pi *Piece) String() string {
	return string(pi[:])
}

func (pi *Piece) Bytes() []byte {
	return pi[:]
}

func NewTorrentFromReader(r *bufio.Reader) *Torrent {
	return &Torrent{
		reader: r,
		parsed: false,
	}
}

func NewTorrentFromFile(filename string) (*Torrent, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return NewTorrentFromReader(bufio.NewReader(file)), nil
}

func (t *Torrent) Parse() error {
	buf, err := io.ReadAll(t.reader)
	if err != nil {
		return err
	}

	buf2 := make([]byte, len(buf))
	copy(buf2, buf)

	t.reader = bufio.NewReader(bytes.NewReader(buf))

	t.announce, err = parseAnnounce(t.reader)
	if err != nil {
		return err
	}

	t.pieces, err = parsePieces(t.reader)
	if err != nil {
		return err
	}

	t.reader = bufio.NewReader(bytes.NewReader(buf2))

	t.infoHash, err = calculateInfohash(t.reader)
	if err != nil {
		return err
	}

	t.parsed = true
	return nil
}

func (t *Torrent) Announce() string {
	if !t.parsed {
		panic(ErrNotParsed)
	}

	return t.announce
}

func (t *Torrent) Pieces() []Piece {
	if !t.parsed {
		panic(ErrNotParsed)
	}

	return t.pieces
}

func (t *Torrent) Infohash() InfoHash {
	if !t.parsed {
		panic(ErrNotParsed)
	}

	return t.infoHash
}

func (i InfoHash) Magnet() string {
	return base32.StdEncoding.EncodeToString(i[:])
}
