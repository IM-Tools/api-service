package cloud

import (
	"github.com/gin-gonic/gin"
	"im-services/internal/api/services"
	"im-services/internal/config"
	"im-services/internal/enum"
	"im-services/pkg/response"
)

type QiNiuHandler struct {
}

var (
	Service services.QiNiuService
)

// @BasePath /api

// PingExample godoc
// @Summary upload/file 文件上传接口
// @Schemes
// @Description 文件上传接口
// @Tags 文件
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer "
// @Param file formData file true "文件"
// @Produce json
// @Success 200 {object} response.JsonResponse{data=Response} "ok"
// @Router /upload/file [post]
func (qiniu *QiNiuHandler) UploadFile(cxt *gin.Context) {
	file, err := cxt.FormFile("file")
	if err != nil {
		response.FailResponse(enum.ParamError, err.Error()).ToJson(cxt)
		return
	}
	filePath := config.Conf.Server.FilePath + "/" + file.Filename
	err = cxt.SaveUploadedFile(file, filePath)

	if err != nil {
		response.FailResponse(enum.ParamError, err.Error()).ToJson(cxt)
		return
	}
	var res Response
	fileUrl, _ := Service.UploadFile(filePath, file.Filename)
	res.FileUrl = fileUrl
	response.SuccessResponse(res).ToJson(cxt)
	return
}

type Response struct {
	FileUrl string `json:"file_url"`
}
