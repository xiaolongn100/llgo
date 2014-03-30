// Copyright 2011 The llgo Authors.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package llgo

import (
	"code.google.com/p/go.tools/go/types"
	"github.com/go-llvm/llvm"
)

// createCall emits the code for a function call,
// taking into account receivers, and panic/defer.
func (fr *frame) createCall(fn *govalue, argValues []*govalue) []*govalue {
	fntyp := fn.Type().Underlying().(*types.Signature)
	typinfo := fr.types.getSignatureInfo(fntyp)

	args := make([]llvm.Value, len(argValues))
	for i, arg := range argValues {
		args[i] = arg.value
	}
	results := typinfo.call(fr.types.ctx, fr.allocaBuilder, fr.builder, fn.value, args)

	resultValues := make([]*govalue, len(results))
	for i, res := range results {
		resultValues[i] = newValue(res, fntyp.Results().At(i).Type())
	}
	return resultValues
}
