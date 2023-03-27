package worker

const (
	GOLANG = "golang"
)

const (
	MEDIUM              = "medium"
	MEDIUM_ROOT_ADDRESS = "https://medium.com"
)

type redditTime string

const (
	REDDIT          = "reddit"
	REDDIT_POST_TOP = "https://www.reddit.com/r/%s/top/.json?limit=%d&t=%s"
	REDDIT_HOUR     = redditTime("hour")
	REDDIT_DAY      = redditTime("day")
	REDDIT_WEEK     = redditTime("week")
	REDDIT_MONTH    = redditTime("month")
	REDDIT_YEAR     = redditTime("year")
	REDDIT_ALL      = redditTime("all")
)
