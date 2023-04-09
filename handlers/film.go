package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	filmdto "week-02-task/dto/film"
	dto "week-02-task/dto/result"
	"week-02-task/models"
	"week-02-task/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_file = "http://localhost:5000/uploads/"

type handlerFilm struct {
	FilmRepository repositories.FilmRepository
}

func HandlerFilm(FilmRepository repositories.FilmRepository) *handlerFilm {
	return &handlerFilm{FilmRepository}
}

func (h *handlerFilm) FindFilms(c echo.Context) error {
	films, err := h.FilmRepository.FindFilms()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// image middleware
	for i, p := range films {
		films[i].ThumbnailFilm = path_file + p.ThumbnailFilm
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: films})
}

func (h *handlerFilm) GetFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	film.ThumbnailFilm = path_file + film.ThumbnailFilm

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: film})
}

func (h *handlerFilm) CreateFilm(c echo.Context) error {

	// request := new(filmdto.CreateFilmRequest)
	// if err := c.Bind(request); err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// }

	// validation := validator.New()
	// err := validation.Struct(request)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// }

	// // data form pattern submit to pattern entity db user
	// film := models.Films{
	// 	Title:         request.Title,
	// 	ThumbnailFilm: request.ThumbnailFilm,
	// 	Year:          request.Year,
	// 	Category:      request.Category,
	// 	CategoryID:    request.CategoryID,
	// 	Description:   request.Description,
	// }

	// data, err := h.FilmRepository.CreateFilm(film)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	// }

	// return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})

	// get the datafile here
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	// rubah string ke integer
	year, _ := strconv.Atoi(c.FormValue("year"))
	category_id, _ := strconv.Atoi(c.FormValue("category_id"))

	request := filmdto.FilmResponse{
		Title:         c.FormValue("title"),
		ThumbnailFilm: dataFile,
		Year:          year,
		CategoryID:    category_id,
		Description:   c.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	film := models.Films{
		Title:         request.Title,
		ThumbnailFilm: request.ThumbnailFilm,
		Year:          request.Year,
		CategoryID:    request.CategoryID,
		Description:   request.Description,
		// UserID:        int(userId),
	}

	film, err = h.FilmRepository.CreateFilm(film)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	film, _ = h.FilmRepository.GetFilm(film.Id)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(film)})

}

func (h *handlerFilm) UpdateFilm(c echo.Context) error {
	request := new(filmdto.UpdateFilmRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.GetFilm(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		film.Title = request.Title
	}

	if request.ThumbnailFilm != "" {
		film.ThumbnailFilm = request.ThumbnailFilm
	}

	if request.Year != 0 {
		film.Year = request.Year
	}

	if request.CategoryID != 0 {
		film.CategoryID = request.CategoryID
	}

	if request.Description != "" {
		film.Description = request.Description
	}

	data, err := h.FilmRepository.UpdateFilm(film)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(data)})
}

func (h *handlerFilm) DeleteFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.FilmRepository.DeleteFilm(film, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseDeleteFilm(data)})
}

func convertResponseFilm(f models.Films) filmdto.FilmResponse {
	return filmdto.FilmResponse{
		Id:            f.Id,
		Title:         f.Title,
		ThumbnailFilm: f.ThumbnailFilm,
		Year:          f.Year,
		Category:      f.Category,
		CategoryID:    f.CategoryID,
		Description:   f.Description,
	}
}

func convertResponseDeleteFilm(f models.Films) filmdto.FilmDeleteResponse {
	return filmdto.FilmDeleteResponse{
		ID: f.Id,
	}
}
