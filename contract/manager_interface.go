package contract

import (
	"context"
	"fmt"
)

type Manager interface {
	Add(...Runnable)
	Start(context.Context) error
	Shutdown(ctx context.Context) error
}

type Runnable interface {
	fmt.Stringer
	// Start 启动
	Start(ctx context.Context) error
	// Attempt 是否允许启动
	Attempt() bool
}
