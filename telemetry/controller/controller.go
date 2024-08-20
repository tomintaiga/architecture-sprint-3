package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	swagger "telemetry/go"
	"telemetry/repository"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	repo repository.Repository
}

func NewController(repo repository.Repository) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (c *Controller) Index() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "I'm ok")
	}
}

func (c *Controller) TelemetryRead() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		pageStr := ctx.DefaultQuery("page", "1")

		page, err := strconv.Atoi(pageStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Can't parse page query")
			log.Printf("Can't parse page query: %v", err)
			return
		}

		perPageStr := ctx.DefaultQuery("per_page", "10")

		perPage, err := strconv.Atoi(perPageStr)
		if err != nil {
			log.Printf("Can't parse per_page query: %v", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		serialStr, ok := ctx.GetQuery("serial")
		if !ok {
			log.Printf("Can't parse per_page query: %v", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		results, err := c.repo.GetValue(ctx, repository.RequestOptions{
			Page:    page,
			PerPage: perPage,
			Serial:  serialStr,
		})

		if errors.Is(err, repository.ErrDoesNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		if err != nil {
			log.Printf("Can't get results: %v", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, results)
	}
}

func (c *Controller) TelemetryWrite() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var record swagger.TelemetryRecord

		err := ctx.ShouldBindJSON(&record)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Can't read body")
			log.Printf("Can't parse body %v", err)
			return
		}

		err = c.repo.SetValue(ctx, record)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Can't set value")
			log.Panicf("Can't write telemetry: %v", err)
			return
		}

		ctx.JSON(http.StatusOK, "")
	}
}
