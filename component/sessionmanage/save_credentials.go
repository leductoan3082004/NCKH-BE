package sessionmanage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"mindzone/appCommon"
	"mindzone/component/asyncjob"
	"time"
)

func SaveCredential(ctx context.Context, rdb *redis.Client, expire time.Duration, newCredential *CredentialsData) error {
	job := asyncjob.NewJob(func(ctx context.Context) error {
		_, err := appCommon.SetJsonRdb(ctx, rdb, newCredential.SessionId, newCredential, expire)
		if err != nil {
			return err
		}
		return nil
	})
	job.SetRetryDurations(time.Millisecond*200, time.Millisecond*200, time.Millisecond*200)
	if err := job.ExecuteWithRetry(ctx); err != nil {
		return err
	}
	return nil
}
