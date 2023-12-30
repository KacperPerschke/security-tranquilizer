package img

import (
	"math/rand"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	in := `Katz`
	img, _ := PackToImg([]byte(in))
	b, _ := UnpackFromImg(img)
	out := string(b)
	if out != in {
		t.Fatalf(`TestBasic: want %#q and got %#q`, in, out)
	}
}

const (
	sizeRange = 1200
	charset   = "\t\n " +
		`!"#$%&'()*+,-./` +
		`0123456789` +
		`:;<=>?@` +
		`ABCDEFGHIJKLMNOPQRSTUVWXYZ` +
		`[\]^_` +
		"`" +
		`abcdefghijklmnopqrstuvwxyz` +
		`{|}~`
)

func TestIncreasingSize(t *testing.T) {
	seededRand := rand.New(
		rand.NewSource(
			time.Now().UnixNano(),
		),
	)
	for size := 1; size <= sizeRange; size++ {
		in := make([]byte, size)
		for i := range in {
			in[i] = charset[seededRand.Intn(len(charset))]
		}
		img, err := PackToImg(in)
		if err != nil {
			t.Fatalf(`TestIncreasingSize got err '%q' from PackToImg at size=%5d`, err, size)
		}
		out, err := UnpackFromImg(img)
		if err != nil {
			t.Fatalf(`TestIncreasingSize got err '%q' from UnpackFromImg at size=%5d`, err, size)
		}
		inS := string(in)
		outS := string(out)
		if outS != inS {
			t.Fatalf(`TestIncreasingSize at size=%5d: want %#q and got %#q`, size, inS, outS)
		}
	}
}
