/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package fs provides a file system abstraction layer.
package fs

import (
	"io"
	"os"
)

// FileSystem groups basic os filesystem methods.
type FileSystem interface {
	Create(name string) (File, error)
	Mkdir(name string) error
	Open(name string) (File, error)
	IsDir(name string) bool
	Exists(name string) bool
	ReadFile(name string) ([]byte, error)
	ReadFiles(name string) (map[string][]byte, error)
	WriteFile(name string, data []byte) error
}

// File groups the basic os.File methods.
type File interface {
	io.ReadWriteCloser
	Stat() (os.FileInfo, error)
}
