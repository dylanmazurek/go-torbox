package torrent

import (
	"io"
	"os"
	"time"

	"github.com/dylanmazurek/go-torbox/internal/crypto"
	"github.com/zeebo/bencode"
)

// Parse parses a torrent file from the given reader.
// r is the reader to read the torrent file from.
func Parse(r io.Reader) (*Torrent, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	metadata := &Metadata{}
	err = bencode.DecodeBytes(data, metadata)
	if err != nil {
		return nil, err
	}

	info := &InfoMetadata{}
	err = bencode.DecodeBytes(metadata.Info, info)
	if err != nil {
		return nil, err
	}

	if len(info.NameUtf8) != 0 {
		info.Name = info.NameUtf8
	}

	files := make([]*File, 0)

	if info.Length > 0 {
		files = append(files, &File{
			Path:   []string{info.Name},
			Length: info.Length,
		})
	} else {
		metadataFiles := make([]*FileMetadata, 0)
		err = bencode.DecodeBytes(info.Files, &metadataFiles)
		if err != nil {
			return nil, err
		}

		for _, f := range metadataFiles {
			if len(f.PathUtf8) != 0 {
				f.Path = f.PathUtf8
			}
			files = append(files, &File{
				Path:   append([]string{info.Name}, f.Path...),
				Length: f.Length,
			})
		}
	}

	announces := make([]string, 0)
	if len(metadata.AnnounceList) > 0 {
		for _, announceItem := range metadata.AnnounceList {
			announces = append(announces, announceItem...)
		}
	} else {
		announces = append(announces, metadata.Announce)
	}

	infoHash := crypto.ToSHA1(metadata.Info, nil)
	if err != nil {
		return nil, err
	}

	newTorrent := Torrent{
		Announce:  announces,
		Comment:   metadata.Comment,
		CreatedBy: metadata.CreatedBy,
		CreatedAt: time.Unix(metadata.CreatedAt, 0),
		InfoHash:  infoHash,
		Files:     files,
	}

	return &newTorrent, nil
}

// ParseFromFile parses a torrent file from the given path.
// p is the path to the torrent file.
func ParseFromFile(p string) (*Torrent, error) {
	file, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	newTorrent, err := Parse(file)
	if err != nil {
		return nil, err
	}

	return newTorrent, nil
}
