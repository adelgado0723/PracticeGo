package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotReader *rot13Reader) Read(in []byte) (int, error) {

	if bytesRead, e := rotReader.r.Read(in); e != nil {
		return 0, e
	} else {
		// Apply Rot 13

		for i := 0; i < len(in); i++ {

			if in[i] > 64 && in[i] <= 90 {
				if in[i]+13 > 90 {
					in[i] = ((in[i] + 13) % 90) + 64
				} else {
					in[i] = in[i] + 13

				}
			} else if in[i] > 96 && in[i] <= 122 {
				if in[i]+13 > 122 {
					in[i] = ((in[i] + 13) % 122) + 96
				} else {
					in[i] = in[i] + 13

				}
			}

		}
		return bytesRead, nil
	}

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
