package services

import (
	"PlaylistsSynchronizer/pkg/api_services"
	"PlaylistsSynchronizer/pkg/models"
	"PlaylistsSynchronizer/pkg/repositories"
	"errors"
)

type TrackService struct {
	repoTrack    repositories.Track
	repoGroup    repositories.UserGroup
	repoToken    repositories.Token
	repoPlaylist repositories.PlayList
	apiService   api_services.ApiService
}

func NewTrackService(repoTrack repositories.Track, repoGroup repositories.UserGroup, repoToken repositories.Token,
	repoPlaylist repositories.PlayList) *TrackService {
	return &TrackService{
		repoTrack:    repoTrack,
		repoGroup:    repoGroup,
		repoPlaylist: repoPlaylist,
		repoToken:    repoToken,
		apiService:   *api_services.NewApiService(repoToken),
	}
}

func (s *TrackService) Add(input models.AddTrackInput) (int, error) {
	playListID, err := s.repoPlaylist.GetByGroupId(input.GroupID)
	if err != nil {
		return 0, nil
	}
	input.PlayListID = playListID.ID
	isTrackPlayListExist, err := s.repoTrack.GetByPlayListIDAndTrackApiID(playListID.ID, models.ApiTrackID{SpotifyUri: input.SpotifyUri,
		YouTubeMusic: input.YouTubeMusicID})
	if isTrackPlayListExist == nil {
		err = s.addSpotify(input)
		if err != nil {
			return 0, err
		}
		err = s.addYouTubeMusic(input)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, errors.New("track already exist in this playlist")
	}
	isTrackExist, err := s.repoTrack.GetByTrackApiID(models.ApiTrackID{SpotifyUri: input.SpotifyUri,
		YouTubeMusic: input.YouTubeMusicID})
	if err != nil {
		return 0, err
	}
	if isTrackExist == nil {
		track := models.Track{SpotifyUri: input.SpotifyUri, YouTubeMusicID: input.YouTubeMusicID}
		id, err := s.repoTrack.Create(input.PlayListID, track)
		if err != nil {
			return 0, nil
		}
		isTrackExist.ID = id
	}
	return isTrackExist.ID, nil
}

func (s *TrackService) DeleteFromPlayList(trackID string) error {
	//TODO implement delete music method
	return nil
}

func (s *TrackService) addSpotify(input models.AddTrackInput) error {
	usersSpotify, err := s.repoGroup.GetByGroupIdSpotifyUser(input.GroupID)
	if err != nil {
		return err
	}
	for _, value := range usersSpotify {
		spotify := s.apiService.GetSpotifyServiceApi()
		err := spotify.AddTrack(models.SpotifyData{Token: value.Token, SpotifyUri: value.SpotifyUri}, value.PlayListID, []models.Track{{SpotifyUri: input.SpotifyUri,
			YouTubeMusicID: input.YouTubeMusicID}})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *TrackService) addYouTubeMusic(input models.AddTrackInput) error {
	usersYouTubeMusic, err := s.repoGroup.GetByGroupIdYouTubeMusicUser(input.GroupID)
	if err != nil {
		return err
	}
	for _, value := range usersYouTubeMusic {
		youTubeMusic := s.apiService.GetYouTubeMusicApiServiceApi()
		err := youTubeMusic.AddTrack(value.Token, value.PlayListID, []models.Track{{SpotifyUri: input.SpotifyUri,
			YouTubeMusicID: input.YouTubeMusicID}})
		if err != nil {
			return err
		}
	}
	return nil
}
