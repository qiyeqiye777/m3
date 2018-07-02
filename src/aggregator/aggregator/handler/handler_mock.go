// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3aggregator/aggregator/handler (interfaces: Handler)

package handler

import (
	"github.com/m3db/m3aggregator/aggregator/handler/writer"

	"github.com/golang/mock/gomock"
	"github.com/uber-go/tally"
)

// Mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *_MockHandlerRecorder
}

// Recorder for MockHandler (not exported)
type _MockHandlerRecorder struct {
	mock *MockHandler
}

func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &_MockHandlerRecorder{mock}
	return mock
}

func (_m *MockHandler) EXPECT() *_MockHandlerRecorder {
	return _m.recorder
}

func (_m *MockHandler) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockHandlerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockHandler) NewWriter(_param0 tally.Scope) (writer.Writer, error) {
	ret := _m.ctrl.Call(_m, "NewWriter", _param0)
	ret0, _ := ret[0].(writer.Writer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHandlerRecorder) NewWriter(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewWriter", arg0)
}