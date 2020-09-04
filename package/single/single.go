package single

import "golang.org/x/sync/singleflight"

var Single *singleflight.Group

func Setup() {
	Single = new(singleflight.Group)
}
