package maildir

import (
	"github.com/infodancer/msgstore"
	"github.com/infodancer/msgstore/errors"
)

func init() {
	msgstore.Register("maildir", func(config msgstore.StoreConfig) (msgstore.MsgStore, error) {
		if config.BasePath == "" {
			return nil, errors.ErrStoreConfigInvalid
		}
		return NewStore(config.BasePath), nil
	})
}
