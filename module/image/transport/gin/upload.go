package imagegin

import (
	"errors"
	"github.com/gin-gonic/gin"
	goservice "github.com/lequocbinh04/go-sdk"
	"github.com/lequocbinh04/go-sdk/plugin/aws"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"nckh-BE/appCommon"
	imagebiz "nckh-BE/module/image/biz"
	imagestorage "nckh-BE/module/image/storage"
	"net/http"
)

func UploadByFile(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		file, err := fileHeader.Open()
		if err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		defer file.Close()

		if fileHeader.Size > int64(1024*1024*15) {
			panic(appCommon.ErrInvalidRequest(errors.New("file size too large")))
		}
		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		db := sc.MustGet(appCommon.DBMain).(*mongo.Client)
		s3 := sc.MustGet(appCommon.PluginAWS).(aws.S3)

		imageStore := imagestorage.NewMgDBStore(db)
		imageBiz := imagebiz.NewUploadBiz(imageStore, s3)
		wc := writeconcern.New(writeconcern.WMajority())
		txnOptions := options.Transaction().SetWriteConcern(wc)

		session, err := db.StartSession()
		if err != nil {
			panic(err)
		}
		defer session.EndSession(c.Request.Context())
		res, err := session.WithTransaction(c.Request.Context(), func(ctx mongo.SessionContext) (interface{}, error) {
			return imageBiz.UploadImage(ctx, dataBytes, appCommon.S3Path, fileHeader.Filename)

		}, txnOptions)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
