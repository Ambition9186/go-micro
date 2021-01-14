package sync

import (
	"context"
	"crypto/tls"
	"time"
)

// WithContext sets context for graceful exit
func WithContext(c context.Context) Option {
	return func(o *Options) {
		o.Cxt = c
	}
}

// Nodes sets the addresses to use
func Nodes(a ...string) Option {
	return func(o *Options) {
		o.Nodes = a
	}
}

// Prefix sets a prefix to any lock ids used
func Prefix(p string) Option {
	return func(o *Options) {
		o.Prefix = p
	}
}

// WithTLS sets the TLS config for the sync
func WithTLS(t *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = t
	}
}

// LeaderContext sets context for graceful exit
func LeaderContext(c context.Context) LeaderOption {
	return func(o *LeaderOptions) {
		o.Cxt = c
	}
}

// LockTTL sets the lock ttl
func LockTTL(t time.Duration) LockOption {
	return func(o *LockOptions) {
		o.TTL = t
	}
}

// LockWait sets the wait time
func LockWait(t time.Duration) LockOption {
	return func(o *LockOptions) {
		o.Wait = t
	}
}
