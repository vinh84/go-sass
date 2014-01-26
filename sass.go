// Copyright 2014 Sam Whited. All rights reserved.
// Use of this source code is governed by the MIT
// license which can be found in the LICENSE file.

package sass

// #cgo LDFLAGS: -lsass
/*
#include <sass_interface.h>
#include <stdlib.h>
void set_source(char* source_string, struct sass_context* ctx) {
	ctx->source_string = source_string;
}
void set_options(struct sass_options options, struct sass_context* ctx) {
	options.output_style = SASS_STYLE_NESTED;
	options.source_comments = 0;
	options.image_path = "images";
	options.include_paths = "";

	ctx->options = options;
}
char* get_output(struct sass_context* ctx) {
	return ctx->output_string;
}
*/
import "C"
import "unsafe"

const (
	SASS_STYLE_NESTED = iota
	SASS_STYLE_EXPANDED
	SASS_STYLE_COMPACT
	SASS_STYLE_COMPRESSED
)

const (
	SASS_SOURCE_COMMENTS_NONE = iota
	SASS_SOURCE_COMMENTS_DEFAULT
	SASS_SOURCE_COMMENTS_MAP
)

type Options struct {
	output_style    int
	source_comments int
	include_paths   string
	image_path      string
}

// Compile the given sass string.
func Compile(source string) (string, error) {
	var (
		ctx     *C.struct_sass_context
		options C.struct_sass_options
		ret     *C.char
	)

	ctx = C.sass_new_context()
	// TODO: Create a Go options struct and convert it
	C.set_options(options, ctx)
	C.set_source(C.CString(source), ctx)
	_, err := C.sass_compile(ctx)
	ret = C.get_output(ctx)
	out := C.GoString(ret)

	// Free memory used by C constructs
	C.sass_free_context(ctx)

	return out, err
}

// Sets the source for the given context.
func (ctx *_Ctype_struct_sass_context) setSource(source string) error {
	source_string := C.CString(source)
	_, err := C.set_source(source_string, ctx)
	C.free(unsafe.Pointer(source_string))
	return err
}
