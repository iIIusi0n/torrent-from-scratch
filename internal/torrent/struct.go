package parser

type Piece [20]byte

type Torrent struct {
	Announce string
	Pieces   []Piece
}

func (pi *Piece) String() string {
	return string(pi[:])
}

func (pi *Piece) Bytes() []byte {
	return pi[:]
}
