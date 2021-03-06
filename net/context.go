// Copyright 2020 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "context"

type withoutCancelContext struct {
	context.Context
}

func (*withoutCancelContext) Err() error {
	return nil
}

func (*withoutCancelContext) Done() <-chan struct{} {
	return nil
}

func WithoutCancel(ctx context.Context) context.Context {
	return &withoutCancelContext{Context: ctx}
}
