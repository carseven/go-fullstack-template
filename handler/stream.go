package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/carseven/go-fullstack-template/view/cases/stream"
	layout "github.com/carseven/go-fullstack-template/view/layouts"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type StreamHandler struct {
	Environment string
}

func (h *StreamHandler) HandleStream(c echo.Context) error {
	streamId := c.Param("streamId")

	if streamId != "" {
		userLanguage := GetRequestHeaderKey(c, "Accept-Language")
		return render(c, layout.Base(stream.StreamView(userLanguage, streamId), h.Environment))
	} else {
		userLanguage := GetRequestHeaderKey(c, "Accept-Language")
		streamIds, err := getVideosIds()
		if err != nil {
			return err
		}
		return render(c, layout.Base(stream.DashboardView(userLanguage, streamIds), h.Environment))
	}

}

func (h *StreamHandler) HandleUpload(c echo.Context) error {
	file, err := c.FormFile("video")
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	src, err := file.Open()
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	defer src.Close()

	// Destination
	// Refactor to upload files to S3
	fileUUID := uuid.New().String()
	videoFileName := fileUUID + ".mp4"
	streamMasterFileName := fileUUID + "-master.m3u8"
	videoFolder := path.Join("assets/video", fileUUID)
	err = os.Mkdir(videoFolder, 0777)
	if err != nil && !os.IsExist(err) {
		c.Logger().Error(err)
		return err
	}

	filePath := path.Join(videoFolder, videoFileName)
	dst, err := os.Create(filePath)
	if err != nil {
		fmt.Println("CREATE")
		c.Logger().Error(err)
		return err
	}
	defer dst.Close()

	// Copy video to destination file
	if _, err = io.Copy(dst, src); err != nil {
		c.Logger().Error(err)
		err := os.Remove(videoFolder)
		return err
	}

	// Transform mp4 to hls format for HTTP streaming
	// Command replication ffmpeg -i video.mp4 -codec: copy -start_number 0 -hls_time 10 -hls_list_size 0 -f hls video.m3u8
	// TODO: Refactor to use https://github.com/u2takey/ffmpeg-go
	filePathHls := path.Join(videoFolder, streamMasterFileName)
	cmd := exec.Command("ffmpeg", "-i", filePath, "-codec:", "copy", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", filePathHls)
	err = cmd.Run()
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	videoStreamingURL := path.Join("/stream", fileUUID)
	return c.HTML(http.StatusOK, fmt.Sprintf(`<a href="%s">Link to start streaming the video uploaded</a>`, videoStreamingURL))
}

func getVideosIds() ([]string, error) {
	entries, err := os.ReadDir("assets/video")
	if err != nil {
		fmt.Println("Error listing streams")
		return nil, err
	}

	var streamIds []string
	for _, e := range entries {
		streamIds = append(streamIds, e.Name())
	}

	fmt.Println(streamIds)

	return streamIds, err
}
