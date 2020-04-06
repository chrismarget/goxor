package decrypt

import (
	"bufio"
	"io"
)

func Decrypt(key []byte, inr io.Reader, owr io.Writer) error {
	in := bufio.NewReader(inr)
	out := bufio.NewWriter(owr)
	inChunk := make([]byte, len(key))
	outChunk := make([]byte, len(key))
	for {
		n, err := io.ReadFull(in, inChunk)
		if err == io.EOF { // at EOF (n==0) we're done
			err := out.Flush()
			return err
		}
		if err != nil {
			if err != io.ErrUnexpectedEOF { // short read EOF (n < len(inChunk)) is last read
				return err
			}
		}

		for i := 0; i < n; i++ {
			outChunk[i] = inChunk[i] ^ key[i]
		}
		_, err = out.Write(outChunk[:n])
		if err != nil {
			return err
		}
	}
}
