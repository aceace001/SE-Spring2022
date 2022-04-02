func GetUserDetails(c *gin.Context) {
	uuid := c.Param("uuid")

	c.JSON(http.StatusOK, response.SuccessMsg(service.UserService.GetUserDetails(uuid)))
}

// 通过用户名获取用户信息
func GetUserOrGroupByName(c *gin.Context) {
	name := c.Query("name")

	c.JSON(http.StatusOK, response.SuccessMsg(service.UserService.GetUserOrGroupByName(name)))
}

func GetUserList(c *gin.Context) {
	uuid := c.Query("uuid")
	c.JSON(http.StatusOK, response.SuccessMsg(service.UserService.GetUserList(uuid)))
}

func AddFriend(c *gin.Context) {
	var userFriendRequest request.FriendRequest
	c.ShouldBindJSON(&userFriendRequest)

	err := service.UserService.AddFriend(&userFriendRequest)
	if nil != err {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessMsg(nil))
}
