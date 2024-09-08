package torrent

import (
	"bufio"
	"crypto/sha1"
	"io"
)

func calculateInfohash(r *bufio.Reader) (InfoHash, error) {
	_, err := readUntilBytes(r, []byte(infoKey))
	if err != nil {
		return InfoHash{}, err
	}

	content, err := io.ReadAll(r)
	if err != nil && err != io.EOF {
		return InfoHash{}, err
	}

	content = content[:len(content)-1]

	hash := sha1.New()
	hash.Write(content)
	sum := hash.Sum(nil)

	return InfoHash(sum), nil
}
