package middlware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Printf("start request: timestamp:[%s] ip:%s method:%s path:%s protocol:%s user-agent:%s\n",
			start.Format(time.RFC1123),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Proto,
			c.Request.UserAgent(),
		)
		c.Next()
		fmt.Printf("response end : timestamp:[%s] status:%d error:%s elapsed:%0.4f\n",
			time.Now().Format(time.RFC1123),
			c.Writer.Status(),
			c.Errors,
			time.Since(start).Seconds(),
		)
	}
}
