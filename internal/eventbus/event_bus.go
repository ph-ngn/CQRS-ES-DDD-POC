package eventbus

import "github.com/andyj29/wannabet/internal/domain"

type Interface interface {
	Publish(domain.Event)
}
