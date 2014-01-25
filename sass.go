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

// Compile the given sass string.
func Compile(source string) (string, error) {
	var (
		ctx     *C.struct_sass_context
		options C.struct_sass_options
	)
	ctx = C.sass_new_context()
	C.set_options(options, ctx)
	_, err := C.sass_compile(ctx)
	out := C.GoString(C.get_output(ctx))
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
