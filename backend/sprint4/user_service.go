//service/user_service.go
func (u *userService) AddFriend(userFriendRequest *request.FriendRequest) error {
    var queryUser *model.User
    db := pool.GetDB()//connect to database
    db.First(&queryUser, "uuid = ?", userFriendRequest.Uuid)
    log.Logger.Debug("queryUser", log.Any("queryUser", queryUser))
    var nullId int32 = 0
    if nullId == queryUser.Id {
        return errors.New("用户不存在")
    }

    var friend *model.User
    db.First(&friend, "username = ?", userFriendRequest.FriendUsername)
    if nullId == friend.Id {
        return errors.New("已添加该好友")
    }

    userFriend := model.UserFriend{
        UserId:   queryUser.Id,
        FriendId: friend.Id,
    }

    var userFriendQuery *model.UserFriend
    db.First(&userFriendQuery, "user_id = ? and friend_id = ?", queryUser.Id, friend.Id)
    if userFriendQuery.ID != nullId {
        return errors.New("该用户已经是你好友")
    }

    db.AutoMigrate(&userFriend)
    db.Save(&userFriend)
    log.Logger.Debug("userFriend", log.Any("userFriend", userFriend))

    return nil
}
