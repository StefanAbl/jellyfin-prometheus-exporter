package service

import (
	"encoding/json"
	"fmt"
	"jellyfin-prometheus-exporter/api"
	"math"
	"net/http"
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
	for item, _ := iter.Next(); item != nil; item, _ = iter.Next() {
		for _, source := range item.MediaSources {
			var width, height int
			for _, stream := range source.MediaStreams {
				if stream.Type == "Video" {
					width = stream.Width
					height = stream.Height
				}
			}
			length := source.RunTimeTicks / (10000 * 1000)
			compressionRatio := float64(source.Size) / float64(length*width*height)
			resolution := fmt.Sprintf("%dx%d", width, height)
			if math.IsNaN(compressionRatio) {
				compressionRatio = -1
			}
			relSizes = append(
				relSizes, RelFileSize{
					Name: item.Name, Path: source.Path, Resolution: resolution, SizeBytes: source.Size,
					LengthSeconds: length, RelSize: compressionRatio,
				},
			)
		}
	}
	return relSizes
}

type RelFileSize struct {
	Name          string  `json:"name"`
	Path          string  `json:"path"`
	Resolution    string  `json:"resolution"`
	SizeBytes     int     `json:"size_bytes"`
	LengthSeconds int     `json:"length_seconds"`
	RelSize       float64 `json:"rel_size"`
}
