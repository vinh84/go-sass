// Copyright 2014 Sam Whited. All rights reserved.
// Use of this source code is governed by the MIT
// license which can be found in the LICENSE file.

package sass

// #cgo LDFLAGS: -lsass
/*
#include <sass_interface.h>
*/
import "C"

// Libsass provides three sass context structs which are used to define
// different execution parameters for sass. The Context struct is used for
// string-in-string-out compilation.
type Context struct {
	ctx *C.struct_sass_context
}

// The file context struct is used for file-based compilation.
type FileContext struct {
	ctx *C.struct_sass_file_context
}

// The folder context struct is used for full-folder multi-file compilation.
type FolderContext struct {
	ctx *C.struct_sass_folder_context
}

// NewContext creates a new context for string-in-string-out compilation.
func NewContext() (c Context, err error) {
	c.ctx, err = C.sass_new_context()
	return
}

// NewFileContext is used for creating a new context for file-based compilation.
func NewFileContext() (c FileContext, err error) {
	c.ctx, err = C.sass_new_file_context()
	return
}

// NewFolderContext creates a new context for full-folder multi-file
// compilation.
func NewFolderContext() (c FolderContext, err error) {
	c.ctx, err = C.sass_new_folder_context()
	return
}

// Free manually frees the memory used by a context.
func (c Context) Free() (err error) {
	_, err = C.sass_free_context(c.ctx)
	return
}

func (c FileContext) Free() (err error) {
	_, err = C.sass_free_file_context(c.ctx)
	return
}

func (c FolderContext) Free() (err error) {
	_, err = C.sass_free_folder_context(c.ctx)
	return
}

// Compile performs the actual compilation of the sass described by the
// context.
func (c Context) Compile() (int, error) {
	i, err := C.sass_compile(c.ctx)
	return int(i), err
}

func (c FileContext) Compile() (int, error) {
	i, err := C.sass_compile_file(c.ctx)
	return int(i), err
}

func (c FolderContext) Compile() (int, error) {
	i, err := C.sass_compile_folder(c.ctx)
	return int(i), err
}
