package utils

import (
	"encoding/binary"

	"github.com/klauspost/compress/zstd"
)

// https://github.com/facebook/zstd/blob/dev/doc/zstd_compression_format.md#zstandard-frames
func isZSTDCompressed(content []byte) bool {
	return len(content) > 4 && binary.LittleEndian.Uint32(content) == 0xFD2FB528
}
func Compress(content []byte) ([]byte, error) {

	compressor, err := zstd.NewWriter(nil)
	if err != nil {
		return nil, err
	}
	defer compressor.Close()
	return compressor.EncodeAll(content, nil), nil
}

func Decompress(content []byte) ([]byte, error) {
	if !isZSTDCompressed(content) {
		return content, nil
	}
	decoder, err := zstd.NewReader(nil)
	if err != nil {
		return nil, err
	}
	defer decoder.Close()
	return decoder.DecodeAll(content, nil)
}
