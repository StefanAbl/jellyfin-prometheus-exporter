package collectors

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"jellyfin-prometheus-exporter/data"
	"net/http"
	"time"
)

type StreamCollector struct {
	url   string
	token string
}

var (
	numberOfStreams = prometheus.NewDesc("streams_total", "The total number of streams", nil, nil)
	bandwidthTotal  = prometheus.NewDesc(
		"streams_bandwidth_bits", "The total bandwidth currently being streamed", nil,
		nil,
	)
)

func NewStreamCollector(url string, token string) *StreamCollector {
	return &StreamCollector{
		url:   url,
		token: token,
	}

}

func (s *StreamCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(s, ch)
}

func (s *StreamCollector) Collect(metrics chan<- prometheus.Metric) {
	sessions := s.getSessions()
	metrics <- prometheus.MustNewConstMetric(
		numberOfStreams, prometheus.GaugeValue, s.getNumberOfActiveSessions(&sessions),
	)
	metrics <- prometheus.MustNewConstMetric(
		bandwidthTotal, prometheus.GaugeValue, s.getBandwidthTotal(&sessions),
	)
}

func (s StreamCollector) getBandwidthTotal(sessions *data.Sessions) float64 {
	total := 0.0
	for _, session := range *sessions {
		if session.PlayState.IsPaused == false && session.NowPlayingItem.Name != "" {
			for _, mediaStream := range session.NowPlayingItem.MediaStreams {
				if bitrate, err := mediaStream.BitRate.Float64(); err != nil {
					_ = fmt.Errorf("got error %s", err.Error())
				} else {
					total += bitrate
				}
			}
		}
	}
	return total
}

func (s StreamCollector) getNumberOfActiveSessions(sessions *data.Sessions) float64 {
	number := 0
	for _, session := range *sessions {
		if session.PlayState.IsPaused == false && session.NowPlayingItem.Name != "" {
			number++
		}
	}
	return float64(number)
}

func (s *StreamCollector) getSessions() data.Sessions {
	var body, err = s.call("Sessions?ActiveWithinSeconds=30", "GET")
	if err != nil {
		fmt.Println(err)
	}

	var sessions data.Sessions
	if err := json.Unmarshal(body, &sessions); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON", err.Error())
	}
	return sessions
}

func (s StreamCollector) call(path, method string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	req, err := http.NewRequest(method, s.url+path, nil)
	if err != nil {
		return nil, fmt.Errorf("got error %s", err.Error())
	}
	req.Header.Set("X-Emby-Token", s.token)
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("got error %s", err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return body, err
}
