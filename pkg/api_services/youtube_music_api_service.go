package api_services

import (
	"PlaylistsSynchronizer/pkg/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type YouTubeMusicApiService struct {
}

func NewYouTubeMusicApiService() *YouTubeMusicApiService {
	return &YouTubeMusicApiService{}
}

func (s *YouTubeMusicApiService) CreatePlayList(token string, playList models.PlayList) (string, error) {
	query := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/playlists?part=snippet%%2Cstatus&key=%s", os.Getenv("API_KEY"))
	jsonBody := []byte(fmt.Sprintf(`{"snippet": { "title": "%s", "description": "%s", "tags": ["tag"], "defaultLanguage": "uk" }}`, playList.Name,
		playList.Description))
	bodyReader := bytes.NewReader(jsonBody)
	client := http.Client{}
	request, err := http.NewRequest("POST", query, bodyReader)
	if err != nil {
		return "", err
	}
	request.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + token},
	}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	var oauthResponse map[string]interface{}
	err = json.Unmarshal(contents, &oauthResponse)
	if err != nil {
		return "", err
	}
	playListID := oauthResponse["id"].(string)
	return playListID, nil
}

func (s *YouTubeMusicApiService) AddTrack(token string, playListId string, tracks []models.Track) error {
	for _, value := range tracks {
		query := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/playlistItems?part=snippet&key=%s",
			os.Getenv("API_KEY"))
		jsonBody := []byte(fmt.Sprintf(`{"snippet": { "playlistId": "%s", "resourceId":{ "kind": "youtube#video", "videoId": "%s"} }}`,
			playListId, value.YouTubeMusicID))
		bodyReader := bytes.NewReader(jsonBody)
		client := http.Client{}
		request, err := http.NewRequest("POST", query, bodyReader)
		if err != nil {
			return err
		}
		request.Header = http.Header{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer " + token},
		}
		response, err := client.Do(request)
		body, _ := io.ReadAll(response.Body)
		jsonMap := make(map[string]interface{})
		err = json.Unmarshal(body, &jsonMap)
		if err != nil {
			return err
		}
		if _, ok := jsonMap["error"]; ok {
			someErrorMap := jsonMap["error"].(map[string]interface{})
			if _, ok := someErrorMap["message"]; ok {
				msg := someErrorMap["message"].(string)
				if strings.Contains(msg, "Video not found") {
					return errors.New("invalid youtube music id")
				}
			}
		}
	}
	return nil
}
