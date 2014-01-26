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
	ctx->options = options;
}
struct sass_options create_options(int output_style, int source_comments, char* image_path, char* include_paths) {
	struct sass_options options;
	options.output_style = output_style;
	options.source_comments = source_comments;
	options.image_path = image_path;
	options.include_paths = include_paths;

	return options;
}
char* get_output(struct sass_context* ctx) {
	return ctx->output_string;
}
*/
import "C"
import "unsafe"

const (
	STYLE_NESTED = iota
	STYLE_EXPANDED
	STYLE_COMPACT
	STYLE_COMPRESSED
)

const (
	SOURCE_COMMENTS_NONE = iota
	SOURCE_COMMENTS_DEFAULT
	SOURCE_COMMENTS_MAP
)

type options struct {
	output_style    int
	source_comments int
	include_paths   string
	image_path      string
}

// Returns a new options struct with the defaults initialized
func NewOptions() options {
	return options{
		output_style:    STYLE_NESTED,
		source_comments: SOURCE_COMMENTS_NONE,
		include_paths:   "",
		image_path:      "images",
	}
}

// Compile the given sass string.
func Compile(source string, opts options) (string, error) {
	var (
		ctx *C.struct_sass_context
		ret *C.char
	)

	ctx = C.sass_new_context()
	defer C.sass_free_context(ctx)
	defer C.free(unsafe.Pointer(ret))

	ctx.setOptions(opts)
	ctx.setSource(source)
	_, err := C.sass_compile(ctx)
	ret = C.get_output(ctx)
	out := C.GoString(ret)

	return out, err
}

// Sets the source for the given context.
func (ctx *_Ctype_struct_sass_context) setSource(source string) error {
	source_string := C.CString(source)
	_, err := C.set_source(source_string, ctx)
	return err
}

// Sets the options for the given context
func (ctx *_Ctype_struct_sass_context) setOptions(opts options) error {
	var (
		coptions C.struct_sass_options
		cim      = C.CString(opts.image_path)
		cin      = C.CString(opts.include_paths)
		cos      = C.int(opts.output_style)
		csc      = C.int(opts.source_comments)
	)

	coptions, err := C.create_options(cos, csc, cim, cin)
	if err != nil {
		return err
	}
	_, err = C.set_options(coptions, ctx)

	return err
}
