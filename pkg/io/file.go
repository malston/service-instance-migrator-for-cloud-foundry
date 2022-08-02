/*
 *  Copyright 2022 VMware, Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *  http://www.apache.org/licenses/LICENSE-2.0
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package io

import (
	"archive/tar"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileDescriptor struct {
	BaseDir   string
	Name      string
	Extension string
	Org       string
	Space     string
}

const pathSep = string(os.PathSeparator)

var supportedExtensions = []string{".yml", ".yaml"}

type InvalidFileExtensionError struct {
	Ext string
}

func (e *InvalidFileExtensionError) Error() string {
	return fmt.Sprintf("extension %q does not match yml or yaml", e.Ext)
}

func NewFileDescriptor(path string) (FileDescriptor, error) {
	dir, filename := filepath.Split(path)
	ext := filepath.Ext(path)
	if !isSupported(ext) {
		return FileDescriptor{}, &InvalidFileExtensionError{Ext: ext}
	}
	name := strings.TrimSuffix(filename, ext)
	org, space := GetOrgSpace(dir)
	orgSpacePath := filepath.Join(org, space)
	dirs := strings.TrimSuffix(dir, pathSep+orgSpacePath+pathSep)
	return FileDescriptor{
		BaseDir:   dirs,
		Extension: strings.TrimPrefix(filepath.Ext(path), "."),
		Name:      name,
		Org:       org,
		Space:     space,
	}, nil
}

func isSupported(ext string) bool {
	for _, s := range supportedExtensions {
		if s == ext {
			return true
		}
	}
	return false
}

func GetOrgSpace(path string) (string, string) {
	parts := strings.Split(path, pathSep)
	orgSpace := parts[len(parts)-3 : len(parts)-1]

	return orgSpace[0], orgSpace[1]
}

func CreateFileIfNotExist(f string) (*os.File, error) {
	var file *os.File
	var err error
	if _, err = os.Stat(f); os.IsNotExist(err) {
		dir := filepath.Dir(f)
		if err = os.MkdirAll(dir, 0755); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to create directory %s", dir))
		}
		if file, err = os.Create(f); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to create file %s", f))
		}
	}

	return file, err
}

func CopyToTempFile(src io.Reader) (*os.File, error) {
	var file *os.File
	var err error

	file, err = os.CreateTemp("", "")
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(file, src); err != nil {
		return nil, err
	}

	if err = file.Close(); err != nil {
		return nil, err
	}

	return file, nil
}

func Tar(src string, writers ...io.Writer) error {
	if _, err := os.Stat(src); err != nil {
		return fmt.Errorf("failed to tar files in %s, %v", src, err)
	}

	mw := io.MultiWriter(writers...)
	tw := tar.NewWriter(mw)
	defer func(tw *tar.Writer) {
		_ = tw.Close()
	}(tw)

	return filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}

		header.Name = fi.Name()

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		f, err := os.Open(file)
		if err != nil {
			return err
		}

		if _, err := io.Copy(tw, f); err != nil {
			return err
		}

		_ = f.Close()

		return nil
	})
}

func Untar(dst string, r io.Reader) error {
	tr := tar.NewReader(r)

	for {
		hdr, err := tr.Next()

		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		// if the header is nil, just skip it (not sure how this happens)
		case hdr == nil:
			continue
		}

		target := filepath.Join(dst, hdr.Name)

		switch hdr.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}

			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			_ = f.Close()
		}
	}
}
