//service/user_service.go
func (u *userService) AddFriend(userFriendRequest *request.FriendRequest) error {
    var queryUser *model.User
    db := pool.GetDB()//connect to database
    db.First(&queryUser, "uuid = ?", userFriendRequest.Uuid)
    log.Logger.Debug("queryUser", log.Any("queryUser", queryUser))
    var nullId int32 = 0
    if nullId == queryUser.Id {
        return errors.New("User does not exist")
    }

    var friend *model.User
    db.First(&friend, "username = ?", userFriendRequest.FriendUsername)
    if nullId == friend.Id {
        return errors.New("this friend has been added")
    }

    userFriend := model.UserFriend{
        UserId:   queryUser.Id,
        FriendId: friend.Id,
    }

    var userFriendQuery *model.UserFriend
    db.First(&userFriendQuery, "user_id = ? and friend_id = ?", queryUser.Id, friend.Id)
    if userFriendQuery.ID != nullId {
        return errors.New("This user is already your friend")
    }

    db.AutoMigrate(&userFriend)
    db.Save(&userFriend)
    log.Logger.Debug("userFriend", log.Any("userFriend", userFriend))

    return nil
}
