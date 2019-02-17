package handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/mt-api/wingo/context"
	m "gitlab.com/mt-api/wingo/handlers/v1"
	"gitlab.com/mt-api/wingo/middleware"
)

const (
	ADMIN_PUB = `-----BEGIN PUBLIC KEY-----
	MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAugd+Lp+a6Si2/SyWLGPt
	NyCnxHkcmwNUaZkaWzDgJnCP8uwMrUai6i6E9a5ouwFF5kub8nNGL6FimfbdfSv1
	SNHv8mwreXkm/DHSSlINrBM7OVwNUw8OhQj5roSmOgON8s3BWeiMH/3xRd8v5+Qy
	sx98HNIMP5niZRmBH2oqN+baoFBuVU3uwgJVA9zRnb7e1n+TUl//gqGYdZI738gr
	/qWVGA91xub9NMgGeL0oajzTTz221eZnBnUmy/beTnSh6aCHlYw/1i6TE2h/9ifF
	qZknsrDsujJsYFdJ7C/DkGClDBL8tgWamfWh0wkQIuarsF/JY330e7u12tVLn6Ry
	MuJ67tNa3IGhf58133noq4CwWYinw/gpmYp5kAxmmPJVO7OYsf+CqMzvYF9iKWRe
	qX8n8z5xLAIg4lkGIgvGeSwQlgycnL9GmGleil3gnG6e++UMNzdP0s+KaiA/9/as
	kX2DByJijOuJB8aEny+E8Sjfjuato+rALuCMX1h8vMdw5gSLdpnEok02GWcHYlMs
	/ZZrcFOMBe4lAt4hiTdkT/zGk4ZmjnfWPRz0T2tSC7AckOvJOPHAaweI93j8Qojk
	Hmlihsbwd104ubzh7hG4hyAbrRjAFDUGSZBipLZ91v48rsi8WjXYhDR0OV3A/xh9
	XDG++pgio0hnw4GwySNxaGUCAwEAAQ==
	-----END PUBLIC KEY-----`
	USER_PUB = `-----BEGIN PUBLIC KEY-----
	MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAvx93jJtUbIq13D0lB51H
	82LIBNHtSlcRoxl0SxkxvXiD1AbtSwBvta3q9q65Nij7vhpwBdMYYrqXf4D5Yi6l
	4T0nDw0bo19fNvve0tSFW/bgg5G7/A2ajAeWmNZN1GSWcdXC53YX1Nk3NnTRJ3D/
	Mupv9T1rrVwuBTVg8cvHlosdikhVo20KheaH3UAuBIV2XwzdEd95tYoTpK1qkY2p
	8ITJwarOZFQW791sAO0CuhKOalzlo7V1sy0knE+7xAuAJmSgsXjB+YdyCoHOU1JD
	n2suSLN4iSALcZO2IVMggvJxfR35akfFA6/eIUBujx2zv+X7FXV1pPAzniS7MPpz
	9FXhhGtDxaYwYc+zi6VvK/0bUa18OCYV39zYVaPC/BDRxWK4szitAwM7ta73Z1bo
	OuuJMiDvrnwGs6Xv/zTv5WxmOuwJhroUhhFgYV9PajMw0EMiq6K+w/yTTPv0C675
	EUz9oIKch2X07t2Aop/zUW8cTiOpIMX3Ocz3sKJZVRItely3/1gI91hwN4KFm1pk
	UMXeFmOO6UuDglPmD1I356iBVdFQnlPs6Ebkc3+ZoZ2woi4ZVdJJ1NCaTixLHUn6
	HkQGlm4hPoD/rfGmF4EosRbHJSPGobUIwh33ss2RgDhhjU9FANYJXH2MfWi9wRQt
	cYuvY3Nse8zjzT7Lfw7UeqkCAwEAAQ==
	-----END PUBLIC KEY-----`
)

//Setup => Setup application handlers
func Setup(r *gin.Engine, appCtx *context.AppContext) {
	vh := m.Handlers{Context: appCtx}
	v1 := r.Group("/v1")

	// Define Groups
	admin := v1.Group("/admin")
	contest := v1.Group("/contest")

	// add auth middleware
	contest.Use(middleware.Auth(appCtx, appCtx.UserKey))
	// admin.Use(middleware.Auth(appCtx, appCtx.AdminKey))

	// admin routes
	admin.Use(middleware.IPCheck())
	admin.POST("contest/meta", vh.AddMetaContest)
	admin.POST("contest/question", vh.AttachQuestion)

	// contest routes
	contest.GET("meta", vh.FindContestMeta)
	contest.GET("config", vh.GetContestConfig)
	contest.POST("answer", vh.PostAnswer)
	contest.POST("store", vh.UpdateUserInfo)

	contest.POST("referral", middleware.IPCheck(), vh.AddReferral)
	contest.POST("user", middleware.IPCheck(), vh.NewUser)
}
