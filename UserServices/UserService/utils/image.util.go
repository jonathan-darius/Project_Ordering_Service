package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
)

type DiskImageStore struct {
	mutex       sync.RWMutex
	imageFolder string
	images      map[string]*ImageInfo
}
type ImageInfo struct {
	UserId string
	Type   string
	Path   string
}
type data struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Display_url string `json:"display_url"`
}
type resImage struct {
	Data data `json:"data"`
}

func Save(userID string, imageType string, imageData bytes.Buffer) (*ImageInfo, error) {
	imagePath := fmt.Sprintf("%s/%s%s", "Image_TMP", userID, imageType)
	mydir, err := os.Getwd()
	file, err := os.Create(imagePath)
	if err != nil {
		return nil, fmt.Errorf("cannot create image file: %w , %s", err, mydir)
	}

	_, err = imageData.WriteTo(file)
	if err != nil {
		return nil, fmt.Errorf("cannot write image to file: %w", err)
	}

	detailImage := &ImageInfo{
		UserId: userID,
		Type:   imageType,
		Path:   imagePath,
	}

	return detailImage, nil
}

func SaveToStorage(userID string, imageData bytes.Buffer) (*ImageInfo, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Load .env File")
	}
	APITOKEN := os.Getenv("IMG_ASSET")
	encoded := base64.StdEncoding.EncodeToString(imageData.Bytes())
	data := url.Values{
		"image": {encoded},
		"name":  {userID},
	}
	r, _ := http.PostForm(fmt.Sprintf("https://api.imgbb.com/1/upload?key=%s", APITOKEN), data)
	defer r.Body.Close()

	buf := &bytes.Buffer{}

	var result resImage
	if _, err := buf.ReadFrom(r.Body); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		return nil, err
	}
	return &ImageInfo{
		UserId: userID,
		Path:   result.Data.Display_url,
	}, nil
}
