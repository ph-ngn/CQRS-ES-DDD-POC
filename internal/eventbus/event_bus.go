package eventbus

import "github.com/andyj29/wannabet/internal/domain/common"

type Interface interface {
	Publish(common.Event)
}
