package data

import (
	"encoding/json"
	"time"
)

type MediaItem struct {
	Name                         string      `json:"Name"`
	OriginalTitle                string      `json:"OriginalTitle"`
	ServerID                     string      `json:"ServerId"`
	ID                           string      `json:"Id"`
	Etag                         string      `json:"Etag"`
	SourceType                   string      `json:"SourceType"`
	PlaylistItemID               string      `json:"PlaylistItemId"`
	DateCreated                  time.Time   `json:"DateCreated"`
	DateLastMediaAdded           time.Time   `json:"DateLastMediaAdded"`
	ExtraType                    string      `json:"ExtraType"`
	AirsBeforeSeasonNumber       json.Number `json:"AirsBeforeSeasonNumber"`
	AirsAfterSeasonNumber        json.Number `json:"AirsAfterSeasonNumber"`
	AirsBeforeEpisodeNumber      json.Number `json:"AirsBeforeEpisodeNumber"`
	CanDelete                    bool        `json:"CanDelete"`
	CanDownload                  bool        `json:"CanDownload"`
	HasSubtitles                 bool        `json:"HasSubtitles"`
	PreferredMetadataLanguage    string      `json:"PreferredMetadataLanguage"`
	PreferredMetadataCountryCode string      `json:"PreferredMetadataCountryCode"`
	SupportsSync                 bool        `json:"SupportsSync"`
	Container                    string      `json:"Container"`
	SortName                     string      `json:"SortName"`
	ForcedSortName               string      `json:"ForcedSortName"`
	Video3DFormat                string      `json:"Video3DFormat"`
	PremiereDate                 time.Time   `json:"PremiereDate"`
	ExternalUrls                 []struct {
		Name string `json:"Name"`
		URL  string `json:"Url"`
	} `json:"ExternalUrls"`
	MediaSources []struct {
		Protocol              string      `json:"Protocol"`
		ID                    string      `json:"Id"`
		Path                  string      `json:"Path"`
		EncoderPath           string      `json:"EncoderPath"`
		EncoderProtocol       string      `json:"EncoderProtocol"`
		Type                  string      `json:"Type"`
		Container             string      `json:"Container"`
		Size                  int         `json:"Size"`
		Name                  string      `json:"Name"`
		IsRemote              bool        `json:"IsRemote"`
		ETag                  string      `json:"ETag"`
		RunTimeTicks          int         `json:"RunTimeTicks"`
		ReadAtNativeFramerate bool        `json:"ReadAtNativeFramerate"`
		IgnoreDts             bool        `json:"IgnoreDts"`
		IgnoreIndex           bool        `json:"IgnoreIndex"`
		GenPtsInput           bool        `json:"GenPtsInput"`
		SupportsTranscoding   bool        `json:"SupportsTranscoding"`
		SupportsDirectStream  bool        `json:"SupportsDirectStream"`
		SupportsDirectPlay    bool        `json:"SupportsDirectPlay"`
		IsInfiniteStream      bool        `json:"IsInfiniteStream"`
		RequiresOpening       bool        `json:"RequiresOpening"`
		OpenToken             string      `json:"OpenToken"`
		RequiresClosing       bool        `json:"RequiresClosing"`
		LiveStreamID          string      `json:"LiveStreamId"`
		BufferMs              json.Number `json:"BufferMs"`
		RequiresLooping       bool        `json:"RequiresLooping"`
		SupportsProbing       bool        `json:"SupportsProbing"`
		VideoType             string      `json:"VideoType"`
		IsoType               string      `json:"IsoType"`
		Video3DFormat         string      `json:"Video3DFormat"`
		MediaStreams          []struct {
			Codec                     string      `json:"Codec"`
			CodecTag                  string      `json:"CodecTag"`
			Language                  string      `json:"Language"`
			ColorRange                string      `json:"ColorRange"`
			ColorSpace                string      `json:"ColorSpace"`
			ColorTransfer             string      `json:"ColorTransfer"`
			ColorPrimaries            string      `json:"ColorPrimaries"`
			DvVersionMajor            json.Number `json:"DvVersionMajor"`
			DvVersionMinor            json.Number `json:"DvVersionMinor"`
			DvProfile                 json.Number `json:"DvProfile"`
			DvLevel                   json.Number `json:"DvLevel"`
			RpuPresentFlag            json.Number `json:"RpuPresentFlag"`
			ElPresentFlag             json.Number `json:"ElPresentFlag"`
			BlPresentFlag             json.Number `json:"BlPresentFlag"`
			DvBlSignalCompatibilityID json.Number `json:"DvBlSignalCompatibilityId"`
			Comment                   string      `json:"Comment"`
			TimeBase                  string      `json:"TimeBase"`
			CodecTimeBase             string      `json:"CodecTimeBase"`
			Title                     string      `json:"Title"`
			VideoRange                string      `json:"VideoRange"`
			VideoRangeType            string      `json:"VideoRangeType"`
			VideoDoViTitle            string      `json:"VideoDoViTitle"`
			LocalizedUndefined        string      `json:"LocalizedUndefined"`
			LocalizedDefault          string      `json:"LocalizedDefault"`
			LocalizedForced           string      `json:"LocalizedForced"`
			LocalizedExternal         string      `json:"LocalizedExternal"`
			DisplayTitle              string      `json:"DisplayTitle"`
			NalLengthSize             string      `json:"NalLengthSize"`
			IsInterlaced              bool        `json:"IsInterlaced"`
			IsAVC                     bool        `json:"IsAVC"`
			ChannelLayout             string      `json:"ChannelLayout"`
			BitRate                   json.Number `json:"BitRate"`
			BitDepth                  json.Number `json:"BitDepth"`
			RefFrames                 json.Number `json:"RefFrames"`
			PacketLength              json.Number `json:"PacketLength"`
			Channels                  json.Number `json:"Channels"`
			SampleRate                json.Number `json:"SampleRate"`
			IsDefault                 bool        `json:"IsDefault"`
			IsForced                  bool        `json:"IsForced"`
			Height                    int         `json:"Height"`
			Width                     int         `json:"Width"`
			AverageFrameRate          json.Number `json:"AverageFrameRate"`
			RealFrameRate             json.Number `json:"RealFrameRate"`
			Profile                   string      `json:"Profile"`
			Type                      string      `json:"Type"`
			AspectRatio               string      `json:"AspectRatio"`
			Index                     json.Number `json:"Index"`
			Score                     json.Number `json:"Score"`
			IsExternal                bool        `json:"IsExternal"`
			DeliveryMethod            string      `json:"DeliveryMethod"`
			DeliveryURL               string      `json:"DeliveryUrl"`
			IsExternalURL             bool        `json:"IsExternalUrl"`
			IsTextSubtitleStream      bool        `json:"IsTextSubtitleStream"`
			SupportsExternalStream    bool        `json:"SupportsExternalStream"`
			Path                      string      `json:"Path"`
			PixelFormat               string      `json:"PixelFormat"`
			Level                     json.Number `json:"Level"`
			IsAnamorphic              bool        `json:"IsAnamorphic"`
		} `json:"MediaStreams"`
		MediaAttachments []struct {
			Codec       string      `json:"Codec"`
			CodecTag    string      `json:"CodecTag"`
			Comment     string      `json:"Comment"`
			Index       json.Number `json:"Index"`
			FileName    string      `json:"FileName"`
			MimeType    string      `json:"MimeType"`
			DeliveryURL string      `json:"DeliveryUrl"`
		} `json:"MediaAttachments"`
		Formats             []string    `json:"Formats"`
		Bitrate             json.Number `json:"Bitrate"`
		Timestamp           string      `json:"Timestamp"`
		RequiredHTTPHeaders struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"RequiredHttpHeaders"`
		TranscodingURL             string      `json:"TranscodingUrl"`
		TranscodingSubProtocol     string      `json:"TranscodingSubProtocol"`
		TranscodingContainer       string      `json:"TranscodingContainer"`
		AnalyzeDurationMs          json.Number `json:"AnalyzeDurationMs"`
		DefaultAudioStreamIndex    json.Number `json:"DefaultAudioStreamIndex"`
		DefaultSubtitleStreamIndex json.Number `json:"DefaultSubtitleStreamIndex"`
	} `json:"MediaSources"`
	CriticRating             json.Number `json:"CriticRating"`
	ProductionLocations      []string    `json:"ProductionLocations"`
	Path                     string      `json:"Path"`
	EnableMediaSourceDisplay bool        `json:"EnableMediaSourceDisplay"`
	OfficialRating           string      `json:"OfficialRating"`
	CustomRating             string      `json:"CustomRating"`
	ChannelID                string      `json:"ChannelId"`
	ChannelName              string      `json:"ChannelName"`
	Overview                 string      `json:"Overview"`
	Taglines                 []string    `json:"Taglines"`
	Genres                   []string    `json:"Genres"`
	CommunityRating          float64     `json:"CommunityRating"`
	CumulativeRunTimeTicks   json.Number `json:"CumulativeRunTimeTicks"`
	RunTimeTicks             json.Number `json:"RunTimeTicks"`
	PlayAccess               string      `json:"PlayAccess"`
	AspectRatio              string      `json:"AspectRatio"`
	ProductionYear           json.Number `json:"ProductionYear"`
	IsPlaceHolder            bool        `json:"IsPlaceHolder"`
	Number                   string      `json:"Number"`
	ChannelNumber            string      `json:"ChannelNumber"`
	IndexNumber              json.Number `json:"IndexNumber"`
	IndexNumberEnd           json.Number `json:"IndexNumberEnd"`
	ParentIndexNumber        json.Number `json:"ParentIndexNumber"`
	RemoteTrailers           []struct {
		URL  string `json:"Url"`
		Name string `json:"Name"`
	} `json:"RemoteTrailers"`
	ProviderIds struct {
		Property1 string `json:"property1"`
		Property2 string `json:"property2"`
	} `json:"ProviderIds"`
	IsHD     bool   `json:"IsHD"`
	IsFolder bool   `json:"IsFolder"`
	ParentID string `json:"ParentId"`
	Type     string `json:"Type"`
	People   []struct {
		Name            string `json:"Name"`
		ID              string `json:"Id"`
		Role            string `json:"Role"`
		Type            string `json:"Type"`
		PrimaryImageTag string `json:"PrimaryImageTag"`
		ImageBlurHashes struct {
			Primary struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Primary"`
			Art struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Art"`
			Backdrop struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Backdrop"`
			Banner struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Banner"`
			Logo struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Logo"`
			Thumb struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Thumb"`
			Disc struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Disc"`
			Box struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Box"`
			Screenshot struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Screenshot"`
			Menu struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Menu"`
			Chapter struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Chapter"`
			BoxRear struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"BoxRear"`
			Profile struct {
				Property1 string `json:"property1"`
				Property2 string `json:"property2"`
			} `json:"Profile"`
		} `json:"ImageBlurHashes"`
	} `json:"People"`
	Studios []struct {
		Name string `json:"Name"`
		ID   string `json:"Id"`
	} `json:"Studios"`
	GenreItems []struct {
		Name string `json:"Name"`
		ID   string `json:"Id"`
	} `json:"GenreItems"`
	ParentLogoItemID        string      `json:"ParentLogoItemId"`
	ParentBackdropItemID    string      `json:"ParentBackdropItemId"`
	ParentBackdropImageTags []string    `json:"ParentBackdropImageTags"`
	LocalTrailerCount       json.Number `json:"LocalTrailerCount"`
	UserData                struct {
		Rating                json.Number `json:"Rating"`
		PlayedPercentage      json.Number `json:"PlayedPercentage"`
		UnplayedItemCount     json.Number `json:"UnplayedItemCount"`
		PlaybackPositionTicks json.Number `json:"PlaybackPositionTicks"`
		PlayCount             json.Number `json:"PlayCount"`
		IsFavorite            bool        `json:"IsFavorite"`
		Likes                 bool        `json:"Likes"`
		LastPlayedDate        time.Time   `json:"LastPlayedDate"`
		Played                bool        `json:"Played"`
		Key                   string      `json:"Key"`
		ItemID                string      `json:"ItemId"`
	} `json:"UserData"`
	RecursiveItemCount      json.Number `json:"RecursiveItemCount"`
	ChildCount              json.Number `json:"ChildCount"`
	SeriesName              string      `json:"SeriesName"`
	SeriesID                string      `json:"SeriesId"`
	SeasonID                string      `json:"SeasonId"`
	SpecialFeatureCount     json.Number `json:"SpecialFeatureCount"`
	DisplayPreferencesID    string      `json:"DisplayPreferencesId"`
	Status                  string      `json:"Status"`
	AirTime                 string      `json:"AirTime"`
	AirDays                 []string    `json:"AirDays"`
	Tags                    []string    `json:"Tags"`
	PrimaryImageAspectRatio json.Number `json:"PrimaryImageAspectRatio"`
	Artists                 []string    `json:"Artists"`
	ArtistItems             []struct {
		Name string `json:"Name"`
		ID   string `json:"Id"`
	} `json:"ArtistItems"`
	Album                 string `json:"Album"`
	CollectionType        string `json:"CollectionType"`
	DisplayOrder          string `json:"DisplayOrder"`
	AlbumID               string `json:"AlbumId"`
	AlbumPrimaryImageTag  string `json:"AlbumPrimaryImageTag"`
	SeriesPrimaryImageTag string `json:"SeriesPrimaryImageTag"`
	AlbumArtist           string `json:"AlbumArtist"`
	AlbumArtists          []struct {
		Name string `json:"Name"`
		ID   string `json:"Id"`
	} `json:"AlbumArtists"`
	SeasonName   string `json:"SeasonName"`
	MediaStreams []struct {
		Codec                     string      `json:"Codec"`
		CodecTag                  string      `json:"CodecTag"`
		Language                  string      `json:"Language"`
		ColorRange                string      `json:"ColorRange"`
		ColorSpace                string      `json:"ColorSpace"`
		ColorTransfer             string      `json:"ColorTransfer"`
		ColorPrimaries            string      `json:"ColorPrimaries"`
		DvVersionMajor            json.Number `json:"DvVersionMajor"`
		DvVersionMinor            json.Number `json:"DvVersionMinor"`
		DvProfile                 json.Number `json:"DvProfile"`
		DvLevel                   json.Number `json:"DvLevel"`
		RpuPresentFlag            json.Number `json:"RpuPresentFlag"`
		ElPresentFlag             json.Number `json:"ElPresentFlag"`
		BlPresentFlag             json.Number `json:"BlPresentFlag"`
		DvBlSignalCompatibilityID json.Number `json:"DvBlSignalCompatibilityId"`
		Comment                   string      `json:"Comment"`
		TimeBase                  string      `json:"TimeBase"`
		CodecTimeBase             string      `json:"CodecTimeBase"`
		Title                     string      `json:"Title"`
		VideoRange                string      `json:"VideoRange"`
		VideoRangeType            string      `json:"VideoRangeType"`
		VideoDoViTitle            string      `json:"VideoDoViTitle"`
		LocalizedUndefined        string      `json:"LocalizedUndefined"`
		LocalizedDefault          string      `json:"LocalizedDefault"`
		LocalizedForced           string      `json:"LocalizedForced"`
		LocalizedExternal         string      `json:"LocalizedExternal"`
		DisplayTitle              string      `json:"DisplayTitle"`
		NalLengthSize             string      `json:"NalLengthSize"`
		IsInterlaced              bool        `json:"IsInterlaced"`
		IsAVC                     bool        `json:"IsAVC"`
		ChannelLayout             string      `json:"ChannelLayout"`
		BitRate                   json.Number `json:"BitRate"`
		BitDepth                  json.Number `json:"BitDepth"`
		RefFrames                 json.Number `json:"RefFrames"`
		PacketLength              json.Number `json:"PacketLength"`
		Channels                  json.Number `json:"Channels"`
		SampleRate                json.Number `json:"SampleRate"`
		IsDefault                 bool        `json:"IsDefault"`
		IsForced                  bool        `json:"IsForced"`
		Height                    int         `json:"Height"`
		Width                     int         `json:"Width"`
		AverageFrameRate          float64     `json:"AverageFrameRate"`
		RealFrameRate             json.Number `json:"RealFrameRate"`
		Profile                   string      `json:"Profile"`
		Type                      string      `json:"Type"`
		AspectRatio               string      `json:"AspectRatio"`
		Index                     json.Number `json:"Index"`
		Score                     json.Number `json:"Score"`
		IsExternal                bool        `json:"IsExternal"`
		DeliveryMethod            string      `json:"DeliveryMethod"`
		DeliveryURL               string      `json:"DeliveryUrl"`
		IsExternalURL             bool        `json:"IsExternalUrl"`
		IsTextSubtitleStream      bool        `json:"IsTextSubtitleStream"`
		SupportsExternalStream    bool        `json:"SupportsExternalStream"`
		Path                      string      `json:"Path"`
		PixelFormat               string      `json:"PixelFormat"`
		Level                     json.Number `json:"Level"`
		IsAnamorphic              bool        `json:"IsAnamorphic"`
	} `json:"MediaStreams"`
	VideoType        string      `json:"VideoType"`
	PartCount        json.Number `json:"PartCount"`
	MediaSourceCount json.Number `json:"MediaSourceCount"`
	ImageTags        struct {
		Property1 string `json:"property1"`
		Property2 string `json:"property2"`
	} `json:"ImageTags"`
	BackdropImageTags   []string `json:"BackdropImageTags"`
	ScreenshotImageTags []string `json:"ScreenshotImageTags"`
	ParentLogoImageTag  string   `json:"ParentLogoImageTag"`
	ParentArtItemID     string   `json:"ParentArtItemId"`
	ParentArtImageTag   string   `json:"ParentArtImageTag"`
	SeriesThumbImageTag string   `json:"SeriesThumbImageTag"`
	ImageBlurHashes     struct {
		Primary struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"Primary"`
		Art struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"Art"`
		Backdrop struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"Backdrop"`
		Banner struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"Banner"`
		Logo struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"Logo"`
		Thumb struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"Thumb"`
		Disc struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"Disc"`
		Box struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"Box"`
		Screenshot struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"Screenshot"`
		Menu struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"Menu"`
		Chapter struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"Chapter"`
		BoxRear struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"BoxRear"`
		Profile struct {
			Property1 string `json:"property1"`
			Property2 string `json:"property2"`
		} `json:"Profile"`
	} `json:"ImageBlurHashes"`
	SeriesStudio             string `json:"SeriesStudio"`
	ParentThumbItemID        string `json:"ParentThumbItemId"`
	ParentThumbImageTag      string `json:"ParentThumbImageTag"`
	ParentPrimaryImageItemID string `json:"ParentPrimaryImageItemId"`
	ParentPrimaryImageTag    string `json:"ParentPrimaryImageTag"`
	Chapters                 []struct {
		StartPositionTicks json.Number `json:"StartPositionTicks"`
		Name               string      `json:"Name"`
		ImagePath          string      `json:"ImagePath"`
		ImageDateModified  time.Time   `json:"ImageDateModified"`
		ImageTag           string      `json:"ImageTag"`
	} `json:"Chapters"`
	LocationType           string      `json:"LocationType"`
	IsoType                string      `json:"IsoType"`
	MediaType              string      `json:"MediaType"`
	EndDate                time.Time   `json:"EndDate"`
	LockedFields           []string    `json:"LockedFields"`
	TrailerCount           json.Number `json:"TrailerCount"`
	MovieCount             json.Number `json:"MovieCount"`
	SeriesCount            json.Number `json:"SeriesCount"`
	ProgramCount           json.Number `json:"ProgramCount"`
	EpisodeCount           json.Number `json:"EpisodeCount"`
	SongCount              json.Number `json:"SongCount"`
	AlbumCount             json.Number `json:"AlbumCount"`
	ArtistCount            json.Number `json:"ArtistCount"`
	MusicVideoCount        json.Number `json:"MusicVideoCount"`
	LockData               bool        `json:"LockData"`
	Width                  json.Number `json:"Width"`
	Height                 json.Number `json:"Height"`
	CameraMake             string      `json:"CameraMake"`
	CameraModel            string      `json:"CameraModel"`
	Software               string      `json:"Software"`
	ExposureTime           json.Number `json:"ExposureTime"`
	FocalLength            json.Number `json:"FocalLength"`
	ImageOrientation       string      `json:"ImageOrientation"`
	Aperture               json.Number `json:"Aperture"`
	ShutterSpeed           json.Number `json:"ShutterSpeed"`
	Latitude               json.Number `json:"Latitude"`
	Longitude              json.Number `json:"Longitude"`
	Altitude               json.Number `json:"Altitude"`
	IsoSpeedRating         json.Number `json:"IsoSpeedRating"`
	SeriesTimerID          string      `json:"SeriesTimerId"`
	ProgramID              string      `json:"ProgramId"`
	ChannelPrimaryImageTag string      `json:"ChannelPrimaryImageTag"`
	StartDate              time.Time   `json:"StartDate"`
	CompletionPercentage   json.Number `json:"CompletionPercentage"`
	IsRepeat               bool        `json:"IsRepeat"`
	EpisodeTitle           string      `json:"EpisodeTitle"`
	ChannelType            string      `json:"ChannelType"`
	Audio                  string      `json:"Audio"`
	IsMovie                bool        `json:"IsMovie"`
	IsSports               bool        `json:"IsSports"`
	IsSeries               bool        `json:"IsSeries"`
	IsLive                 bool        `json:"IsLive"`
	IsNews                 bool        `json:"IsNews"`
	IsKids                 bool        `json:"IsKids"`
	IsPremiere             bool        `json:"IsPremiere"`
	TimerID                string      `json:"TimerId"`
	CurrentProgram         struct {
	} `json:"CurrentProgram"`
}
