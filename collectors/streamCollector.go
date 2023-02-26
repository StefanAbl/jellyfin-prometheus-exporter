package collectors

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"jellyfin-prometheus-exporter/api"
	"jellyfin-prometheus-exporter/data"
)

type StreamCollector struct {
	api *api.JellyfinApi
}

var (
	numberOfStreams = prometheus.NewDesc(
		"jellyfin_active_streams_count", "The total number of streams",
		[]string{"user"}, nil,
	)
	bandwidthTotal = prometheus.NewDesc(
		"jellyfin_streams_bandwidth_bits", "The total bandwidth currently being streamed", nil,
		nil,
	)
	numberOfTranscodedStreams = prometheus.NewDesc(
		"jellyfin_active_streams_transcode_count",
		"The number of streams which are currently being transcoded", nil, nil,
	)
	numberOfDirectStreams = prometheus.NewDesc(
		"jellyfin_active_streams_direct_count",
		"The number of streams which are currently being direct streams", nil, nil,
	)
	items = prometheus.NewDesc("jellyfin_items_count", "Count of Media Items label denotes type", []string{"type"}, nil)
)

func NewStreamCollector(jellyfinApi *api.JellyfinApi) *StreamCollector {
	return &StreamCollector{
		api: jellyfinApi,
	}

}

func (s *StreamCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(s, ch)
}

func (s *StreamCollector) Collect(metrics chan<- prometheus.Metric) {
	sessions := s.api.GetSessions()
	s.getNumberOfActiveSessions(&sessions, metrics)
	metrics <- prometheus.MustNewConstMetric(
		bandwidthTotal, prometheus.GaugeValue, s.getBandwidthTotal(&sessions),
	)
	transcoded, direct := s.getTranscodedStreams(&sessions)
	metrics <- prometheus.MustNewConstMetric(
		numberOfTranscodedStreams, prometheus.GaugeValue,
		transcoded,
	)
	metrics <- prometheus.MustNewConstMetric(
		numberOfDirectStreams, prometheus.GaugeValue,
		direct,
	)
	itemCounts := s.api.GetItemCounts()
	for k, v := range itemCounts {
		value, _ := v.Float64()
		metrics <- prometheus.MustNewConstMetric(items, prometheus.GaugeValue, value, k)
	}
}

func (s StreamCollector) getTranscodedStreams(sessions *data.Sessions) (float64, float64) {
	transcoded, direct := 0, 0
	for _, session := range *sessions {
		if session.PlayState.IsPaused == false && session.NowPlayingItem.Name != "" {
			if session.TranscodingInfo.TranscodeReasons != nil && !session.TranscodingInfo.IsVideoDirect {
				transcoded++
			} else {
				direct++
			}

		}
	}
	return float64(transcoded), float64(direct)
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

func (s StreamCollector) getNumberOfActiveSessions(sessions *data.Sessions, metrics chan<- prometheus.Metric) {
	var sessionsPerUser = make(map[string]int)
	for _, session := range *sessions {
		if session.PlayState.IsPaused == false && session.NowPlayingItem.Name != "" {
			sessionsPerUser[session.UserName]++
		}
	}
	for user, i := range sessionsPerUser {
		metrics <- prometheus.MustNewConstMetric(
			numberOfStreams, prometheus.GaugeValue, float64(i), user,
		)
	}
}
