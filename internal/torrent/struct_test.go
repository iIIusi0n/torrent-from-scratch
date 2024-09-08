package torrent_test

import (
	"testing"
	"torrent-from-scratch/internal/torrent"
)

func TestNewTorrentFromFile(t *testing.T) {
	tfile, err := torrent.NewTorrentFromFile("test/ubuntu-server.torrent")
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	err = tfile.Parse()
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if tfile.Announce() != "https://torrent.ubuntu.com/announce" {
		t.Fatalf("expected %s, got %s", "https://torrent.ubuntu.com/announce", tfile.Announce())
	}

	if len(tfile.Pieces()) != 10582 {
		t.Fatalf("expected %d, got %d", 10582, len(tfile.Pieces()))
	}

	if tfile.Infohash().Magnet() != "IHTM2UGM5RK42VYEYXR5C5XHWWJRPI73" {
		t.Fatalf("expected %s, got %s", "IHTM2UGM5RK42VYEYXR5C5XHWWJRPI73", tfile.Infohash().Magnet())
	}
}
