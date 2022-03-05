package main

import (
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.GET("/bucket/:bucketName", getGCSBucket())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func getGCSBucket() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Read credenatil json from default environment value, GOOGLE_APPLICATION_CREDENTIALS
		client, err := storage.NewClient(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		bucketName := c.Param("bucketName")
		if bucketName == "" {
			return c.JSON(http.StatusBadRequest, fmt.Sprintln("failed to get bucket name"))
		}

		bucket := client.Bucket(bucketName)
		attr, err := bucket.Attrs(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, fmt.Sprintf("%#v", attr))
	}
}
