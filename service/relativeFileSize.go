package service

import (
	"encoding/json"
	"fmt"
	"jellyfin-prometheus-exporter/api"
	"jellyfin-prometheus-exporter/data"
	"math"
	"net/http"
	"sync"
)

type RelativeFileSizeService struct {
	Api *api.JellyfinApi
}

func (s *RelativeFileSizeService) GetRelativeFileSizes(w http.ResponseWriter, r *http.Request) {
	relSizes := s.computeRelativeFileSizes()
	array := make([]interface{}, len(relSizes))
	for i, v := range relSizes {
		array[i] = v
	}
	err := json.NewEncoder(w).Encode(array)
	if err != nil {
		w.WriteHeader(500)
	}

}

func (s *RelativeFileSizeService) computeRelativeFileSizes() []RelFileSize {
	relSizes := []RelFileSize{}
	iter := s.Api.GetMediaItems("Episode,Movie")

	var wg sync.WaitGroup
	ch := make(chan RelFileSize)
	go func() {
		for v := range ch {
			relSizes = append(relSizes, v)
		}
	}()
	for item, err := iter.Next(); item != nil; item, err = iter.Next() {
		if err != nil {
			fmt.Print(err)
		}
		wg.Add(1)
		go computeRelSizes(ch, &wg, item)
	}
	wg.Wait()
	close(ch)

	return relSizes
}

func computeRelSizes(ch chan RelFileSize, wg *sync.WaitGroup, item *data.MediaItem) {
	defer wg.Done()
	for _, source := range item.MediaSources {
		var width, height int
		var codec string
		for _, stream := range source.MediaStreams {
			if stream.Type == "Video" {
				width = stream.Width
				height = stream.Height
				codec = stream.Codec
			}
		}
		if codec == "" {
			break
		}
		length := source.RunTimeTicks / (10000 * 1000)
		compressionRatio := float64(source.Size) / float64(length*width*height)
		resolution := fmt.Sprintf("%dx%d", width, height)
		if math.IsNaN(compressionRatio) {
			compressionRatio = -1
		}
		ch <- RelFileSize{
			Name: item.Name, Path: source.Path, Resolution: resolution, SizeBytes: source.Size,
			LengthSeconds: length, RelSize: compressionRatio, Codec: codec,
		}

	}
}

type RelFileSize struct {
	Name          string  `json:"name"`
	Path          string  `json:"path"`
	Resolution    string  `json:"resolution"`
	SizeBytes     int     `json:"size_bytes"`
	LengthSeconds int     `json:"length_seconds"`
	RelSize       float64 `json:"rel_size"`
	Codec         string  `json:"codec"`
}
