/*
 * MIT License
 *
 * Copyright (c) 2018
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package trie

import (
	"errors"
)

var ErrNotImplemented = errors.New("not implemented")

type Trie interface {
	// Try to add the key/val to the trie,
	// can fail if there isn't enough room.
	//
	// If already exists, will overwrite old 
	// value.
	TryWrite(key string, value int64) error

	// Try to find the key/val in the trie,
	// returns error if could not find.
	TryRead(key string) (int64, error)

	// Remove the key from the trie, noop
	// if the key does not exist.
	Delete(key string)

	// Save the internal array to a file.
	Save(filename string) error

	// Load the internal array from a file.
	Load(filename string) error
}

type trie struct{}

func (t *trie) TryWrite(key string, value int64) error {
	return ErrNotImplemented
}

func (t *trie) TryRead(key string) (int64, error) {
	return 0, ErrNotImplemented
}

func (t *trie) Delete(key string) {}

func (t *trie) Save(filename string) error {
	return ErrNotImplemented
}

func (t *trie) Load(filename string) error {
	return ErrNotImplemented
}
