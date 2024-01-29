package appCommon

import (
	"os"
	"time"
)

type DbType int

const (
	CurrentUser             = "user"
	CurrentSpaceId          = "space_id"
	CurrentMaximumImageSize = "maximum_image_size"

	MainDBName = "mindzone"

	DBMain               = "mongodb"
	PasetoProvider       = "paseto"
	PluginRedis          = "redis"
	PluginLocker         = "locker"
	PluginGrpcServer     = "grpc-server"
	PluginGrpcUserClient = "grpc-user-client"
	PluginSocketIO       = "socketio"
	PluginAws            = "aws"
	PluginLocalLocker    = "local-locker"
	PluginES             = "es"
	PluginRabbitMQ       = "rabbitmq"
	PluginRedisCache     = "cache"
	PluginAwsRekognition = "rekognition"
	PluginTinybird       = "tinybird"
)

const (
	TopicAddTwitterCard           = "mindzone.twitter_card"
	TopicAddImageCard             = "mindzone.image_card"
	TopicAddWebCard               = "mindzone.web_card"
	TopicSummaryCard              = "mindzone.summary_card"
	TopicSmartFolder              = "mindzone.smart_folder"
	TopicDeleteSpaceImages        = "mindzone.delete_space_images"
	TopicEditCardSmartFolder      = "mindzone.edit_card_smart_folder"
	TopicUploadImagesFacebook     = "mindzone.upload_images_facebook"
	TopicCardError                = "mindzone.card_error"
	TopicCardSummaryError         = "mindzone.card_summary_error"
	TopicDeleteSpaceElasticSearch = "mindzone.delete_space_elastic_search"
)

var (
	S3Mindzone        = "mindzone_image"
	S3Domain          = "https://" + os.Getenv("AWS_DOMAIN")
	S3PathFolderImage = "folder_image"
	S3UrlFolderImage  = Join("/", S3Domain, S3PathFolderImage)

	S3PathSpaceImage = Join("/", "mindzone_image", "space")
)

const (
	UserCreateCardLock         = "USER_CREATE_CARD_"
	UserCreateLock             = "USER_CREATE_"
	CallbackTransactionLock    = "CALLBACK_TRANSACTION_LOCK_"
	CallbackCouponLock         = "CALLBACK_COUPON_LOCK_"
	SlugCreateLock             = "SLUG_CREATE_"
	LimitCardCreateLock        = "LIMIT_CARD_CREATE_"
	AccountCodeCreateLock      = "ACCOUNT_CODE_CREATE_"
	SummaryLock                = "SUMMARY_LOCK_"
	LimitSmartFolderCreateLock = "LIMIT_SMART_FOLDER_CREATE_LOCK_"
	LimitPinCard               = "LIMIT_PIN_CARD_"
)

const (
	UserKb2aIdCache = "USER_KB2A_ID_CACHE_"
	UserIdCache     = "USER_ID_CACHE_"
	LimitNameCache  = "LIMIT_NAME_CACHE_"
	LimitIdCache    = "LIMIT_ID_CACHE_"
	SpaceIdCache    = "SPACE_ID_CACHE_"
)

var (
	LimitCacheTTL = time.Minute * 30
)

const (
	CardLimitNormal  = 500
	CardLimitPremium = 1000
)

const (
	FreePackage    = "free"
	PremiumPackage = "premium"
)
const (
	DbTypeUser = iota
	DbTypeFingerprint
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

type TokenPayload struct {
	UId       int64  `json:"user_id"`
	URole     int    `json:"role"`
	Type      string `json:"type"`
	SessionId string `json:"session_id"`
}

func (p TokenPayload) UserId() int64 {
	return p.UId
}

func (p TokenPayload) Role() int {
	return p.URole
}

func (p TokenPayload) UType() string {
	return p.Type
}

func (p TokenPayload) USessionId() string {
	return p.SessionId
}
