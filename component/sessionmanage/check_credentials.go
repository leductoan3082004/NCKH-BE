package sessionmanage

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"mindzone/appCommon"
	"mindzone/component/asyncjob"
	"mindzone/plugin/tokenprovider"
	"time"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func levenshteinDistance(a, b string) int {
	lenA, lenB := len(a), len(b)
	if lenA == 0 {
		return lenB
	}
	if lenB == 0 {
		return lenA
	}

	dp := make([][]int, lenA+1)
	for i := range dp {
		dp[i] = make([]int, lenB+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i <= lenA; i++ {
		for j := 1; j <= lenB; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+cost)
		}
	}

	return dp[lenA][lenB]
}
func CheckCredential(ctx context.Context, rdb *redis.Client, sessionId, userAgent, ip string) error {
	job := asyncjob.NewJob(func(ctx context.Context) error {
		var credentialData *CredentialsData
		err := appCommon.GetJsonRdb(ctx, rdb, sessionId, &credentialData)
		if err != nil {
			return tokenprovider.ErrInvalidTokenWithErr(err)
		}
		if float64(levenshteinDistance(credentialData.UserAgent, userAgent))/float64((len(credentialData.UserAgent)+len(userAgent))/2)*100 >= 5 {
			return tokenprovider.ErrInvalidTokenWithErr(errors.New("diff ua"))
		}
		if credentialData.Ip != ip {
			return tokenprovider.ErrInvalidTokenWithErr(errors.New("diff ip: " + credentialData.Ip + "!=" + ip))
		}
		return nil
	})
	job.SetRetryDurations(time.Millisecond*200, time.Millisecond*200, time.Millisecond*200)
	if err := job.ExecuteWithRetry(ctx); err != nil {
		return err
	}
	return nil
}
