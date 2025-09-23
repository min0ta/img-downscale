package imgproccessor

import (
	"bytes"
	"fmt"
	"image"
	"image_thumb/internal/config"
	"net/http"
	"os"
	"time"

	"github.com/disintegration/imaging"
)

type ImageProcessor struct {
	cfg    *config.Config
	client *http.Client
}

func New(cfg *config.Config) *ImageProcessor {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	return &ImageProcessor{
		cfg:    cfg,
		client: client,
	}
}
func (proc *ImageProcessor) GetThumbnail(url string) (*image.NRGBA, error) {
	res, err := proc.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("bad status code %v", res.StatusCode)
	}
	if res.ContentLength > proc.cfg.DownloadLimit {
		return nil, fmt.Errorf("too large file (%v bytes)", res.ContentLength)
	}

	img, err := imaging.Decode(res.Body)
	if err != nil {
		return nil, err
	}
	thumb := imaging.Thumbnail(img, 100, 100, imaging.Lanczos)
	return thumb, nil
}

func (proc *ImageProcessor) SaveThumbnail(thumb *image.NRGBA, id string) error {
	storagePath := proc.cfg.StoragePath + id + ".jpg"
	var buf bytes.Buffer
	err := imaging.Encode(&buf, thumb, imaging.JPEG)
	if err != nil {
		return err
	}
	return os.WriteFile(storagePath, buf.Bytes(), 0644)
}
