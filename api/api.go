package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jellyfin-prometheus-exporter/data"
	"net/http"
	"time"
)

type JellyfinApi struct {
	url   string
	token string
}

func NewJellyfinApi(url string, token string) *JellyfinApi {
	return &JellyfinApi{
		url:   url,
		token: token,
	}

}

type MediaItemsIterator struct {
	urlTemplate  string
	counter      int
	currentItems []data.MediaItem
	step         int
	api          *JellyfinApi
}

func newMediaItemsIterator(urlTemplate string, api *JellyfinApi) *MediaItemsIterator {
	return &MediaItemsIterator{
		urlTemplate:  urlTemplate,
		counter:      0,
		currentItems: nil,
		step:         250,
		api:          api,
	}
}

func (i *MediaItemsIterator) Next() (*data.MediaItem, error) {
	if i.counter%i.step == 0 {
		url := fmt.Sprintf(i.urlTemplate, i.counter, i.step)
		res, err := i.api.call(url, "GET")

		if err != nil {
			fmt.Println(fmt.Errorf("got error %s", err.Error()))
		}
		var ir data.ItemResponse
		if err := json.Unmarshal(res, &ir); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON", err.Error())
			if err.Error() != "unexpected end of JSON input" {
				return nil, err
			}
		} else {
			if ir.Items == nil {
				return nil, nil
			}
			i.currentItems = ir.Items
		}
	}
	i.counter = i.counter + 1
	if (i.counter-1)%i.step >= len(i.currentItems) {
		return nil, nil
	}
	return &i.currentItems[(i.counter-1)%i.step], nil

}

func (s *JellyfinApi) GetMediaItems(includeItemTypes string) *MediaItemsIterator {
	return newMediaItemsIterator(
		"Items?"+
			"StartIndex=%d"+
			"&Limit=%d"+
			"&Fields=Path,MediaSources"+
			"&ImageTypeLimit=0"+
			"&SortBy=SortName"+
			"&SortOrder=Ascending"+
			"&includeItemTypes="+includeItemTypes+
			"&Recursive=true", s,
	)
}

func (s *JellyfinApi) GetMediaItemsByParent(parentId string) ([]data.MediaItem, error) {
	res, err := s.call(
		"Items?StartIndex=0&Limit=100&Fields=PrimaryImageAspectRatio,SortName,Path,SongCount,"+
			"ChildCount,MediaSources&ImageTypeLimit=0&ParentId="+parentId+"&SortBy=SortName&SortOrder=Ascending", "GET",
	)
	if err != nil {
		return nil, err
	}
	var ir data.ItemResponse
	if err := json.Unmarshal(res, &ir); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON", err.Error())
	}
	return ir.Items, nil
}

func (s *JellyfinApi) GetVirtualFolders() ([]data.VirtualFolder, error) {
	res, err := s.call("Library/VirtualFolders", "GET")
	if err != nil {
		return nil, err
	}
	var vf []data.VirtualFolder
	if err := json.Unmarshal(res, &vf); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON", err.Error())
	}
	return vf, nil
}

func (s *JellyfinApi) GetItemCounts() data.ItemCounts {
	res, err := s.call("Items/Counts", "GET")
	if err != nil {
		fmt.Println(err)
	}
	var itemCounts data.ItemCounts
	if err := json.Unmarshal(res, &itemCounts); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON", err.Error())
	}
	return itemCounts
}

func (s *JellyfinApi) GetSessions() data.Sessions {
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

func (s JellyfinApi) call(path, method string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 60,
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
