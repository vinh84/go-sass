// Copyright 2014 Sam Whited. All rights reserved.
// Use of this source code is governed by the MIT
// license which can be found in the LICENSE file.

// Package sass provides a cgo wrapper for libsass which can be used to compile
// SASS and SCSS to CSS.
package sass

// #cgo LDFLAGS: -lsass
/*
#include <sass_interface.h>
#include <stdlib.h>
void set_source(char* source_string, struct sass_context* ctx) {
	ctx->source_string = source_string;
}
void set_file_path(char* input_path, struct sass_file_context* ctx) {
	ctx->input_path = input_path;
}
void set_folder_paths(char* search_path, char* output_path, struct sass_folder_context* ctx) {
	ctx->search_path = search_path;
	ctx->output_path = output_path;
}
void set_options(struct sass_options options, struct sass_context* ctx) {
	ctx->options = options;
}
void set_file_options(struct sass_options options, struct sass_file_context* ctx) {
	ctx->options = options;
}
void set_folder_options(struct sass_options options, struct sass_folder_context* ctx) {
	ctx->options = options;
}
struct sass_options create_options(int output_style, int source_comments, char* image_path, char* include_paths) {
	struct sass_options options;
	options.output_style = output_style;
	options.source_comments = source_comments;
	//options.image_path = image_path;
	options.include_paths = include_paths;

	return options;
}

char* get_output(struct sass_context* ctx) {
	return ctx->output_string;
}

int get_error_status(struct sass_context* ctx) {
	return ctx->error_status;
}

char* get_error_message(struct sass_context* ctx) {
	return ctx->error_message;
}

char* get_file_output(struct sass_file_context* ctx) {
	return ctx->output_string;
}
*/
import "C"
import "unsafe"
import "errors"

// Output CSS styles for compiling SASS.
const (
	STYLE_NESTED = iota
	STYLE_EXPANDED
	STYLE_COMPACT
	STYLE_COMPRESSED
)

// Output style for comments.
const (
	SOURCE_COMMENTS_NONE = iota
	SOURCE_COMMENTS_DEFAULT
	SOURCE_COMMENTS_MAP
)

// A set of options to pass to the SASS compiler.
type options struct {
	output_style    int
	source_comments int
	include_paths   string
	image_path      string
}

// Returns a new options struct with the defaults initialized.
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
	errorstatus := C.get_error_status(ctx)
	errormessage := C.get_error_message(ctx)
	
	if errorstatus != 0 {
		err = errors.New(C.GoString(errormessage))
	}
	
	out := C.GoString(ret)

	return out, err
}

// Compile the given file.
func CompileFile(path string, opts options) (string, error) {
	var (
		ctx *C.struct_sass_file_context
		ret *C.char
	)

	ctx = C.sass_new_file_context()
	defer C.sass_free_file_context(ctx)
	defer C.free(unsafe.Pointer(ret))

	ctx.setOptions(opts)
	ctx.setPath(path)
	_, err := C.sass_compile_file(ctx)
	ret = C.get_file_output(ctx)
	out := C.GoString(ret)

	return out, err
}

// Compile all sass/scss files in the given directory.
func CompileDir(
	searchPath string,
	outPath string,
	opts options) error {

	ctx := C.sass_new_folder_context()
	defer C.sass_free_folder_context(ctx)

	ctx.setOptions(opts)
	ctx.setPaths(searchPath, outPath)
	_, err := C.sass_compile_folder(ctx)

	return err
}

// Sets the source for the given libsass context.
func (ctx *_Ctype_struct_sass_context) setSource(source string) error {
	csource := C.CString(source)
	_, err := C.set_source(csource, ctx)
	return err
}

// Sets the path for the given libsass file context.
func (ctx *_Ctype_struct_sass_file_context) setPath(path string) error {
	cpath := C.CString(path)
	_, err := C.set_file_path(cpath, ctx)
	return err
}

// Sets the search path and output path for the given libsass folder context.
func (ctx *_Ctype_struct_sass_folder_context) setPaths(
	searchPath string, outPath string) error {
	cspath := C.CString(searchPath)
	copath := C.CString(outPath)
	_, err := C.set_folder_paths(cspath, copath, ctx)
	return err
}

// Sets the libsass options for the given context.
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

// Sets the libsass options for the given file context.
func (ctx *_Ctype_struct_sass_file_context) setOptions(opts options) error {
	coptions, err := createCOptions(opts)
	if err != nil {
		return err
	}
	_, err = C.set_file_options(coptions, ctx)

	return err
}

// Sets the libsass options for the given folder context.
func (ctx *_Ctype_struct_sass_folder_context) setOptions(opts options) error {
	coptions, err := createCOptions(opts)
	if err != nil {
		return err
	}
	_, err = C.set_folder_options(coptions, ctx)

	return err
}

// Create a C options struct for libsass from some Go options.
func createCOptions(opts options) (C.struct_sass_options, error) {
	var (
		coptions C.struct_sass_options
		cim      = C.CString(opts.image_path)
		cin      = C.CString(opts.include_paths)
		cos      = C.int(opts.output_style)
		csc      = C.int(opts.source_comments)
	)

	coptions, err := C.create_options(cos, csc, cim, cin)
	return coptions, err
}
