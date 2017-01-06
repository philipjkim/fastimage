package fastimage

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkCustomTimeout(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	images, err := readSampleFile("imgsrc.log")
	if err != nil {
		b.Errorf("failed to open image source file : %v", err)
	}

	type counter struct {
		success int
		failure int
	}
	c := counter{}
	for i := 0; i < b.N; i++ {
<<<<<<< HEAD
		url := images[r.Intn(len(images))]
		_, _, err := DetectImageTypeWithTimeout(url, 1000)
		// it, is, err := DetectImageTypeWithTimeout(url, 10000)
		// b.Logf("type:%v, size:%v, err:%v", it, is, err)
		if err == nil {
			c.success++
		} else {
			c.failure++
		}
=======
		it, is, err := DetectImageTypeWithTimeout(url, 1000)
		b.Logf("type:%v, size:%v, err:%v", it, is, err)
>>>>>>> e1a1877e509a0bcf0acc307d7349ee49da3c54be
	}
	b.Logf("success:%v, failure:%v", c.success, c.failure)
}
