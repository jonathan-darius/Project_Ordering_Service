package Utils

import (
	"ProductService/ProductService/Config"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

func SaveToStorage(userID string, imageData bytes.Buffer) (*ImageInfo, error) {
	APITOKEN := Config.GetEnv("IMG_ASSET")
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
