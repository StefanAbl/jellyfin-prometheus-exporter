package data

type VirtualFolder struct {
	Name           string   `json:"Name"`
	Locations      []string `json:"Locations"`
	CollectionType string   `json:"CollectionType"`
	LibraryOptions struct {
		EnablePhotos                          bool `json:"EnablePhotos"`
		EnableRealtimeMonitor                 bool `json:"EnableRealtimeMonitor"`
		EnableChapterImageExtraction          bool `json:"EnableChapterImageExtraction"`
		ExtractChapterImagesDuringLibraryScan bool `json:"ExtractChapterImagesDuringLibraryScan"`
		PathInfos                             []struct {
			Path string `json:"Path"`
		} `json:"PathInfos"`
		SaveLocalMetadata                       bool          `json:"SaveLocalMetadata"`
		EnableInternetProviders                 bool          `json:"EnableInternetProviders"`
		EnableAutomaticSeriesGrouping           bool          `json:"EnableAutomaticSeriesGrouping"`
		EnableEmbeddedTitles                    bool          `json:"EnableEmbeddedTitles"`
		EnableEmbeddedEpisodeInfos              bool          `json:"EnableEmbeddedEpisodeInfos"`
		AutomaticRefreshIntervalDays            int           `json:"AutomaticRefreshIntervalDays"`
		PreferredMetadataLanguage               string        `json:"PreferredMetadataLanguage"`
		MetadataCountryCode                     string        `json:"MetadataCountryCode"`
		SeasonZeroDisplayName                   string        `json:"SeasonZeroDisplayName"`
		MetadataSavers                          []interface{} `json:"MetadataSavers"`
		DisabledLocalMetadataReaders            []interface{} `json:"DisabledLocalMetadataReaders"`
		LocalMetadataReaderOrder                []string      `json:"LocalMetadataReaderOrder"`
		DisabledSubtitleFetchers                []interface{} `json:"DisabledSubtitleFetchers"`
		SubtitleFetcherOrder                    []interface{} `json:"SubtitleFetcherOrder"`
		SkipSubtitlesIfEmbeddedSubtitlesPresent bool          `json:"SkipSubtitlesIfEmbeddedSubtitlesPresent"`
		SkipSubtitlesIfAudioTrackMatches        bool          `json:"SkipSubtitlesIfAudioTrackMatches"`
		SubtitleDownloadLanguages               []interface{} `json:"SubtitleDownloadLanguages"`
		RequirePerfectSubtitleMatch             bool          `json:"RequirePerfectSubtitleMatch"`
		SaveSubtitlesWithMedia                  bool          `json:"SaveSubtitlesWithMedia"`
		AutomaticallyAddToCollection            bool          `json:"AutomaticallyAddToCollection"`
		AllowEmbeddedSubtitles                  string        `json:"AllowEmbeddedSubtitles"`
		TypeOptions                             []struct {
			Type                 string   `json:"Type"`
			MetadataFetchers     []string `json:"MetadataFetchers"`
			MetadataFetcherOrder []string `json:"MetadataFetcherOrder"`
			ImageFetchers        []string `json:"ImageFetchers"`
			ImageFetcherOrder    []string `json:"ImageFetcherOrder"`
			ImageOptions         []struct {
				Type     string `json:"Type"`
				Limit    int    `json:"Limit"`
				MinWidth int    `json:"MinWidth"`
			} `json:"ImageOptions"`
		} `json:"TypeOptions"`
	} `json:"LibraryOptions"`
	ItemID             string `json:"ItemId"`
	PrimaryImageItemID string `json:"PrimaryImageItemId"`
	RefreshStatus      string `json:"RefreshStatus"`
}
