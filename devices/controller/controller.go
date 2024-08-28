package controller

import (
	"devices/models"
	"devices/repository"
	"errors"
	"log"
	"net/http"
	"strconv"

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

func (c *Controller) AddDevice() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var device models.Device

		err := ctx.BindJSON(&device)
		if err != nil {
			log.Printf("Can't parse device body: %v", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		newDevice, err := c.repo.AddDevice(ctx, device)
		if err != nil {
			log.Printf("Can't add device: %v", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, newDevice)
	}
}

func (c *Controller) findDeviceById(ctx *gin.Context) *models.Device {
	idStr := ctx.Param("id")

	if idStr == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		log.Printf("No device ID")
		return nil
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Can't parse id param: %v", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return nil
	}

	device, err := c.repo.GetDeviceById(ctx, int32(id))
	if errors.Is(err, repository.ErrDeviceNotFound) {
		log.Printf("Device %v not found", id)
		ctx.AbortWithStatus(http.StatusNotFound)
		return nil
	}

	return &device
}

func (c *Controller) GetDevice() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		device := c.findDeviceById(ctx)

		if device != nil {
			ctx.JSON(http.StatusOK, device)
		}
	}
}

func (c *Controller) CallCommand() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		device := c.findDeviceById(ctx)

		if device == nil {
			return
		}

		var cmd models.Command
		err := ctx.BindJSON(&cmd)
		if err != nil {
			log.Printf("Can't parse cmd body: %v", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if cmd.Name != models.PowerCmdName {
			log.Printf("Bad CMD name: %v", cmd.Name)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if cmd.Args != models.PowerOn && cmd.Args != models.PowerOff {
			log.Printf("Bad CMD args: %v", cmd.Args)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if cmd.Args == models.PowerOn {
			device.Status = true
		} else {
			device.Status = false
		}

		err = c.repo.UpdateDevice(ctx, *device)
		if err != nil {
			log.Printf("Can't update device: %v", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, device)
	}
}
