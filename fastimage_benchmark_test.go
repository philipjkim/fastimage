package fastimage

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func BenchmarkCustomTimeout(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	images, err := readSampleFile("imgsrc.log")
	if err != nil {
		b.Errorf("failed to open image source file : %v", err)
	}

	logfile, err := os.Create("imgfetch.log")
	defer logfile.Close()
	if err != nil {
		b.Errorf("failed to create log file : %v", err)
	}
	type counter struct {
		success int
		failure int
	}
	c := counter{}
	for i := 0; i < b.N; i++ {
		url := images[r.Intn(len(images))]
		// _, _, err := DetectImageTypeWithTimeout(url, 1000)
		it, is, err := DetectImageTypeWithTimeout(url, 1000)
		logfile.WriteString(fmt.Sprintf("url:%v, type:%v, size:%v, err:%v\n", url, it, is, err))
		if err == nil {
			c.success++
		} else {
			c.failure++
		}
	}
	b.Logf("success:%v, failure:%v", c.success, c.failure)
}
