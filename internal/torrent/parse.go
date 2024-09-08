package torrent

import "bufio"

func parseAnnounce(r *bufio.Reader) (string, error) {
	_, err := readUntilBytes(r, []byte(announceKey))
	if err != nil {
		return "", err
	}

	length, err := readLength(r)
	if err != nil {
		return "", err
	}

	announce, err := readString(r, length)
	if err != nil {
		return "", err
	}

	return announce, nil
}

func parsePieces(r *bufio.Reader) ([]Piece, error) {
	_, err := readUntilBytes(r, []byte(piecesKey))
	if err != nil {
		return nil, err
	}

	length, err := readLength(r)
	if err != nil {
		return nil, err
	}

	pieces := make([]Piece, length/20)
	for i := 0; i < length/20; i++ {
		_, err := r.Read(pieces[i][:])
		if err != nil {
			return nil, err
		}
	}

	return pieces, nil
}
