package torrent

import (
	"bufio"
	"strconv"
)

const (
	announceKey = "8:announce"
	piecesKey   = "6:pieces"
)

func readUntilBytes(r *bufio.Reader, target []byte) ([]byte, error) {
	res := make([]byte, 0)

	pos := 0
	for {
		b, err := r.ReadByte()
		if err != nil {
			return res, err
		}
		if b == target[pos] {
			pos++
			if pos == len(target) {
				return res, nil
			}
		} else {
			pos = 0
		}

		res = append(res, b)
	}
}

func readUntilByte(r *bufio.Reader, target byte) ([]byte, error) {
	res := make([]byte, 0)

	for {
		b, err := r.ReadByte()
		if err != nil {
			return res, err
		}
		if b == target {
			return res, nil
		}

		res = append(res, b)
	}
}

func readLength(r *bufio.Reader) (int, error) {
	lengthStr, err := readUntilByte(r, ':')
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(string(lengthStr))
}

func readString(r *bufio.Reader, length int) (string, error) {
	str := make([]byte, length)
	_, err := r.Read(str)
	return string(str), err
}

func readBlob(r *bufio.Reader, length int) ([]byte, error) {
	blob := make([]byte, length)
	_, err := r.Read(blob)
	return blob, err
}
