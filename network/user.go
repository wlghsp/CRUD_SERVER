package network

import (
	"CRUD_SERVER/service"
	"CRUD_SERVER/types"
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

	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}

		router.registerGET("/", userRouterInstance.get)
		router.registerPOST("/", userRouterInstance.create)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)

	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create 입니다. ")

	u.router.okResponse(c, &types.CreateUserResponse{
		ApiResponse: types.NewApiResponse("생성입니다", 1),
	})
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get 입니다. ")

	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("조회입니다", 1),
		User:        nil,
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update 입니다. ")

	u.router.okResponse(c, &types.UpdateUserResponse{
		ApiResponse: types.NewApiResponse("수정입니다", 1),
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete 입니다. ")

	u.router.okResponse(c, &types.DeleteUserResponse{
		ApiResponse: types.NewApiResponse("삭제입니다", 1),
	})
}
