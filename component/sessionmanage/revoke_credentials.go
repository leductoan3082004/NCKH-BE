package sessionmanage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"mindzone/appCommon"
	"mindzone/component/asyncjob"
	"mindzone/plugin/tokenprovider"
	"time"
)

func RevokeCredential(ctx context.Context, rdb *redis.Client, sessionId string) error {
	job := asyncjob.NewJob(func(ctx context.Context) error {
		err := appCommon.DeleteJsonRdb(ctx, rdb, sessionId)
		if err != nil {
			return tokenprovider.ErrInvalidTokenWithErr(err)
		}
		return nil
	})
	job.SetRetryDurations(time.Millisecond*200, time.Millisecond*200, time.Millisecond*200)
	if err := job.ExecuteWithRetry(ctx); err != nil {
		return err
	}
	return nil
}
