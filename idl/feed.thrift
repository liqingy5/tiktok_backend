namespace go feed
struct douyinFeedRequest {
    1: optional i64 latestTime;
    2: optional string token;
}

struct douyinFeedResponse {
    1: optional i32 statusCode;
    2: optional string statusMsg;
    3: list<video> videoList;
    4: optional i64 nextTime;
}

struct video {
    1: i64 id;
    2: user author;
    3: string playUrl;
    4: string coverUrl;
    5: i64 favoriteCount;
    6: i64 commentCount;
    7: bool isFavorite;
    8: string title;
}

struct user {
    1: i64 id;
    2: string name;
    3: optional i64 followCount;
    4: optional i64 followerCount;
    5: bool isFollow;
}

service feedService {
    douyinFeedResponse getFeeds(1: douyinFeedRequest req) (api.get="/douyin/feed");
}