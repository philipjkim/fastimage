package fastimage

import (
	"bufio"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestPNGImage(t *testing.T) {
	t.Parallel()

	url := "http://fc08.deviantart.net/fs71/f/2012/214/7/c/futurama__bender_by_suzura-d59kq1p.png"

	imagetype, size, err := DetectImageType(url)
	if err != nil {
		t.Error("Failed to detect image type")
	}

	if size == nil {
		t.Error("Failed to detect image size")
	}

	if imagetype != PNG {
		t.Error("Image is not PNG")
	}

	if size.Width != 988 {
		t.Error("Image width is wrong")
	}

	if size.Height != 1240 {
		t.Error("Image height is wrong")
	}
}

func TestJPEGImage(t *testing.T) {
	t.Parallel()

	url := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"

	imagetype, size, err := DetectImageType(url)
	if err != nil {
		t.Error("Failed to detect image type")
	}

	if size == nil {
		t.Error("Failed to detect image size")
	}

	if imagetype != JPEG {
		t.Error("Image is not JPEG")
	}

	if size.Width != 5000 {
		t.Error("Image width is wrong")
	}

	if size.Height != 2813 {
		t.Error("Image height is wrong")
	}
}

func TestGIFImage(t *testing.T) {
	t.Parallel()

	url := "http://media.giphy.com/media/gXcIuJBbRi2Va/giphy.gif"

	imagetype, size, err := DetectImageType(url)
	if err != nil {
		t.Error("Failed to detect image type")
	}

	if size == nil {
		t.Error("Failed to detect image size")
	}

	if imagetype != GIF {
		t.Error("Image is not GIF")
	}

	if size.Width != 500 {
		t.Error("Image width is wrong")
	}

	if size.Height != 286 {
		t.Error("Image height is wrong")
	}
}

func TestBMPImage(t *testing.T) {
	t.Parallel()

	url := "http://www.ac-grenoble.fr/ien.vienne1-2/spip/IMG/bmp_Image004.bmp"

	imagetype, size, err := DetectImageType(url)
	if err != nil {
		t.Error("Failed to detect image type")
	}

	if imagetype != BMP {
		t.Error("Image is not BMP")
	}

	if size != nil {
		t.Error("We can't detect BMP size yet")
	}
}

func TestTIFFImage(t *testing.T) {
	t.Parallel()

	url := "http://www.fileformat.info/format/tiff/sample/c44cf1326c2240d38e9fca073bd7a805/download"

	imagetype, size, err := DetectImageType(url)
	if err != nil {
		t.Error("Failed to detect image type")
	}

	if imagetype != TIFF {
		t.Error("Image is not TIFF")
	}

	if size != nil {
		t.Error("We can't detect TIFF size yet")
	}

}

func TestReadFile(t *testing.T) {
	f, err := os.Open("imgsrc.log")
	if err != nil {
		t.Errorf("failed to open image source file : %v", err)
	}
	fs := bufio.NewScanner(f)
	var images []string
	for fs.Scan() {
		images = append(images, fs.Text())
	}

	t.Logf("total entries:%v", len(images))
}

func TestCustomTimeout(t *testing.T) {
	t.Parallel()

	images, err := readSampleFile("imgsrc.log")
	if err != nil {
		t.Errorf("failed to open image source file : %v", err)
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	url := images[r.Intn(len(images))]
	t.Logf("url: %v", url)

	imagetype, size, err := DetectImageTypeWithTimeout(url, 500)
	t.Logf("imageType: %v", imagetype)
	t.Logf("size: %v", size)
	t.Logf("error: %v", err)
}
