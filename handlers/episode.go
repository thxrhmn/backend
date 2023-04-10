package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	episodedto "week-02-task/dto/episode"
	dto "week-02-task/dto/result"
	"week-02-task/models"
	"week-02-task/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerEpisode struct {
	EpisodeRepository repositories.EpisodeRepository
}

func HandlerEpisode(EpisodeRepository repositories.EpisodeRepository) *handlerEpisode {
	return &handlerEpisode{EpisodeRepository}
}

func (h *handlerEpisode) FindEpisodesByFilm(c echo.Context) error {
	film_id, _ := strconv.Atoi(c.Param("film_id"))
	episodes, err := h.EpisodeRepository.FindEpisodesByFilm(film_id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	for i, m := range episodes {
		episodes[i].ThumbnailFilm = path_file + m.ThumbnailFilm
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: episodes})
}

func (h *handlerEpisode) GetEpisodeByFilm(c echo.Context) error {
	film_id, _ := strconv.Atoi(c.Param("film_id"))
	episode_id, _ := strconv.Atoi(c.Param("episode_id"))

	episode, err := h.EpisodeRepository.GetEpisodeByFilm(film_id, episode_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	episode.ThumbnailFilm = path_file + episode.ThumbnailFilm

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: episode})
}

func (h *handlerEpisode) FindEpisodes(c echo.Context) error {
	episodes, err := h.EpisodeRepository.FindEpisodes()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// image middleware
	for i, p := range episodes {
		episodes[i].ThumbnailFilm = path_file + p.ThumbnailFilm
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: episodes})
}

func (h *handlerEpisode) GetEpisode(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	episode.ThumbnailFilm = path_file + episode.ThumbnailFilm

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: episode})
}

func (h *handlerEpisode) CreateEpisode(c echo.Context) error {
	// request := new(episodedto.CreateEpisodeRequest)
	// if err := c.Bind(request); err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// }

	// validation := validator.New()
	// err := validation.Struct(request)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// }

	// // data form pattern submit to pattern entity db user
	// episode := models.Episode{
	// 	Title:         request.Title,
	// 	ThumbnailFilm: request.ThumbnailFilm,
	// 	LinkFilm:      request.LinkFilm,
	// 	Film:          request.Film,
	// 	FilmID:        request.FilmID,
	// }

	// data, err := h.EpisodeRepository.CreateEpisode(episode)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	// }

	// return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
	// get the datafile here
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	// rubah string ke integer
	film_id, _ := strconv.Atoi(c.FormValue("film_id"))

	request := episodedto.EpisodeResponse{
		Title:         c.FormValue("title"),
		ThumbnailFilm: dataFile,
		LinkFilm:      c.FormValue("linkfilm"),
		FilmID:        film_id,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	episode := models.Episode{
		Title:         request.Title,
		ThumbnailFilm: request.ThumbnailFilm,
		LinkFilm:      request.LinkFilm,
		FilmID:        request.FilmID,
	}

	episode, err = h.EpisodeRepository.CreateEpisode(episode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	episode, _ = h.EpisodeRepository.GetEpisode(episode.Id)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(episode)})
}

func (h *handlerEpisode) UpdateEpisode(c echo.Context) error {
	request := new(episodedto.UpdateEpisodeRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	episode, err := h.EpisodeRepository.GetEpisode(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		episode.Title = request.Title
	}

	if request.ThumbnailFilm != "" {
		episode.ThumbnailFilm = request.ThumbnailFilm
	}

	if request.LinkFilm != "" {
		episode.LinkFilm = request.LinkFilm
	}

	data, err := h.EpisodeRepository.UpdateEpisode(episode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerEpisode) DeleteEpisode(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.EpisodeRepository.DeleteEpisode(episode, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDeleteEpisode(data)})
}

func convertResponseEpisode(e models.Episode) episodedto.EpisodeResponse {
	return episodedto.EpisodeResponse{
		Title:         e.Title,
		ThumbnailFilm: e.ThumbnailFilm,
		LinkFilm:      e.LinkFilm,
		FilmID:        e.FilmID,
	}
}

func convertResponseDeleteEpisode(e models.Episode) episodedto.EpisodeDeleteResponse {
	return episodedto.EpisodeDeleteResponse{
		ID: e.Id,
	}
}
