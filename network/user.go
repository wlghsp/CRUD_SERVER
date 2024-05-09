package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	// service
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
		}

		userRouterInstance.router.engin.POST("/", userRouterInstance.create)
		userRouterInstance.router.engin.PUT("/", userRouterInstance.update)
		userRouterInstance.router.engin.DELETE("/", userRouterInstance.delete)
		userRouterInstance.router.engin.GET("/", userRouterInstance.get)
	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create 입니다. ")
}

func (u *userRouter) get(c *gin.Context) {
}

func (u *userRouter) update(c *gin.Context) {
}

func (u *userRouter) delete(c *gin.Context) {
}
