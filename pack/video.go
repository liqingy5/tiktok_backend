/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pack

import (
	feed "tiktok_backend/biz/model/feed"
	"tiktok_backend/model"
)

// Videos Convert model.Video list to user_gorm.Video list
func Videos(models []*model.Video) []*feed.Video {
	videoList := make([]*feed.Video, 0, len(models))
	for _, m := range models {
		if u := Video(m); u != nil {
			videoList = append(videoList, u)
		}
	}
	return videoList
}

// Video Convert model.video to user_gorm.Video
func Video(model *model.Video) *feed.Video {
	if model == nil {
		return nil
	}
	return &feed.Video{
		ID:            int64(model.ID),
		Author:        &feed.User{ID: int64(model.UserId)},
		PlayUrl:       model.VideoUrl,
		CoverUrl:      model.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         model.Title,
	}
}

func User(model *model.User) *feed.User {
	var user *feed.User = new(feed.User)
	user.ID = int64(model.ID)
	user.Name = model.Username
	user.FollowCount = new(int64)
	user.FollowerCount = new(int64)
	*user.FollowCount = int64(model.FollowingCount)
	*user.FollowerCount = int64(model.FollowerCount)
	return user
}
