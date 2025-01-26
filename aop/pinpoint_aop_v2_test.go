/*
 * Copyright 2025 NAVER Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// [2025] Happy Spring Festival
package aop

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"golang.org/x/arch/x86/x86asm"
)

func TestEmptyBody(t *testing.T) {
	t.Error("xxx")
}

func TestShowAsm(t *testing.T) {
	dst := reflect.ValueOf(TestEmptyBody)
	// unsafe.Pointer(src.Pointer())

	ptr := (*byte)(unsafe.Pointer(dst.Pointer()))
	t.Log(ptr)
	data := unsafe.Slice(ptr, 64)
	// for {
	decoder, _ := x86asm.Decode(data, 64)
	// 打印汇编指令
	fmt.Println()
	t.Logf("%s", x86asm.GNUSyntax(decoder, 0, nil))
	t.Error("")
}
