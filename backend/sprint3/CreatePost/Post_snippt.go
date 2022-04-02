//unfinished code snippt
func CreatePost(c *Context, post *model.Post, triggerWebhooks bool) (*model.Post, *model.AppError) {
	var pchan store.StoreChannel
	if len(post.RootId) > 0 {
		pchan = Srv.Store.Post().Get(post.RootId)
	}

	// Verify the parent/child relationships are correct
	if pchan != nil {
		if presult := <-pchan; presult.Err != nil {
			return nil, model.NewLocAppError("createPost", "api.post.create_post.root_id.app_error", nil, "")
		} else {
			list := presult.Data.(*model.PostList)
			if len(list.Posts) == 0 || !list.IsChannelId(post.ChannelId) {
				return nil, model.NewLocAppError("createPost", "api.post.create_post.channel_root_id.app_error", nil, "")
			}

			if post.ParentId == "" {
				post.ParentId = post.RootId
			}

			if post.RootId != post.ParentId {
				parent := list.Posts[post.ParentId]
				if parent == nil {
					return nil, model.NewLocAppError("createPost", "api.post.create_post.parent_id.app_error", nil, "")
				}
			}
		}
	}

	if post.CreateAt != 0 && !HasPermissionToContext(c, model.PERMISSION_MANAGE_SYSTEM) {
		post.CreateAt = 0
		c.Err = nil
	}

	post.Hashtags, _ = model.ParseHashtags(post.Message)

	var rpost *model.Post
	if result := <-Srv.Store.Post().Save(post); result.Err != nil {
		return nil, result.Err
	} else {
		rpost = result.Data.(*model.Post)
	}

	if len(post.FileIds) > 0 {
		// There's a rare bug where the client sends up duplicate FileIds so protect against that
		post.FileIds = utils.RemoveDuplicatesFromStringArray(post.FileIds)

		for _, fileId := range post.FileIds {
			if result := <-Srv.Store.FileInfo().AttachToPost(fileId, post.Id); result.Err != nil {
				l4g.Error(utils.T("api.post.create_post.attach_files.error"), post.Id, post.FileIds, c.Session.UserId, result.Err)
			}
		}
	}

	handlePostEvents(c, rpost, triggerWebhooks)

	return rpost, nil
}
