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
	cContext *C.struct_sass_context
}

// The file context struct is used for file-based compilation.
type FileContext struct {
	cContext *C.struct_sass_file_context
}

// The folder context struct is used for full-folder multi-file compilation.
type FolderContext struct {
	cContext *C.struct_sass_folder_context
}

// NewContext creates a new context for string-in-string-out compilation.
func NewContext() (c Context, err error) {
	c.cContext, err = C.sass_new_context()
	return
}

// NewFileContext is used for creating a new context for file-based compilation.
func NewFileContext() (c FileContext, err error) {
	c.cContext, err = C.sass_new_file_context()
	return
}

// NewFolderContext creates a new context for full-folder multi-file
// compilation.
func NewFolderContext() (c FolderContext, err error) {
	c.cContext, err = C.sass_new_folder_context()
	return
}

// FreeContext manually frees the memory used by a context.
func FreeContext(c Context) (err error) {
	_, err = C.sass_free_context(c.cContext)
	return
}

// FreeFileContext manually frees the memory used by a file context.
func FreeFileContext(c FileContext) (err error) {
	_, err = C.sass_free_file_context(c.cContext)
	return
}

// FreeFolderContext manually frees the memory used by a folder context.
func FreeFolderContext(c FolderContext) (err error) {
	_, err = C.sass_free_folder_context(c.cContext)
	return
}

// Compile performs the actual compilation of the sass described by the
// context.
func Compile(c Context) (int, error) {
	i, err := C.sass_compile(c.cContext)
	return int(i), err
}

// CompileFile compiles the file described by the file context.
func CompileFile(c FileContext) (int, error) {
	i, err := C.sass_compile_file(c.cContext)
	return int(i), err
}

// CompileFolder compiles all sass files in the folder described by the folder
// context.
func CompileFolder(c FolderContext) (int, error) {
	i, err := C.sass_compile_folder(c.cContext)
	return int(i), err
}
