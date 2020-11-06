// Copyright 2019 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/devtools/resultstore/v2/test_suite.proto

package resultstore

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The result of running a test case.
type TestCase_Result int32

const (
	// The implicit default enum value. Do not use.
	TestCase_RESULT_UNSPECIFIED TestCase_Result = 0
	// Test case ran to completion. Look for failures or errors to determine
	// whether it passed, failed, or errored.
	TestCase_COMPLETED TestCase_Result = 1
	// Test case started but did not complete because the test harness received
	// a signal and decided to stop running tests.
	TestCase_INTERRUPTED TestCase_Result = 2
	// Test case was not started because the test harness received a SIGINT or
	// timed out.
	TestCase_CANCELLED TestCase_Result = 3
	// Test case was not run because the user or process running the test
	// specified a filter that excluded this test case.
	TestCase_FILTERED TestCase_Result = 4
	// Test case was not run to completion because the test case decided it
	// should not be run (eg. due to a failed assumption in a JUnit4 test).
	// Per-test setup or tear-down may or may not have run.
	TestCase_SKIPPED TestCase_Result = 5
	// The test framework did not run the test case because it was labeled as
	// suppressed.  Eg. if someone temporarily disables a failing test.
	TestCase_SUPPRESSED TestCase_Result = 6
)

// Enum value maps for TestCase_Result.
var (
	TestCase_Result_name = map[int32]string{
		0: "RESULT_UNSPECIFIED",
		1: "COMPLETED",
		2: "INTERRUPTED",
		3: "CANCELLED",
		4: "FILTERED",
		5: "SKIPPED",
		6: "SUPPRESSED",
	}
	TestCase_Result_value = map[string]int32{
		"RESULT_UNSPECIFIED": 0,
		"COMPLETED":          1,
		"INTERRUPTED":        2,
		"CANCELLED":          3,
		"FILTERED":           4,
		"SKIPPED":            5,
		"SUPPRESSED":         6,
	}
)

func (x TestCase_Result) Enum() *TestCase_Result {
	p := new(TestCase_Result)
	*p = x
	return p
}

func (x TestCase_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestCase_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_resultstore_v2_test_suite_proto_enumTypes[0].Descriptor()
}

func (TestCase_Result) Type() protoreflect.EnumType {
	return &file_google_devtools_resultstore_v2_test_suite_proto_enumTypes[0]
}

func (x TestCase_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestCase_Result.Descriptor instead.
func (TestCase_Result) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_test_suite_proto_rawDescGZIP(), []int{2, 0}
}

// The result of running a test suite, as reported in a <testsuite> element of
// an XML log.
type TestSuite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The full name of this suite, as reported in the name attribute. For Java
	// tests, this is normally the fully qualified class name. Eg.
	// "com.google.common.hash.BloomFilterTest".
	SuiteName string `protobuf:"bytes,1,opt,name=suite_name,json=suiteName,proto3" json:"suite_name,omitempty"`
	// The results of the test cases and test suites contained in this suite,
	// as reported in the <testcase> and <testsuite> elements contained within
	// this <testsuite>.
	Tests []*Test `protobuf:"bytes,2,rep,name=tests,proto3" json:"tests,omitempty"`
	// Failures reported in <failure> elements within this <testsuite>.
	Failures []*TestFailure `protobuf:"bytes,3,rep,name=failures,proto3" json:"failures,omitempty"`
	// Errors reported in <error> elements within this <testsuite>.
	Errors []*TestError `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
	// The timing for the entire TestSuite, as reported by the time attribute.
	Timing *Timing `protobuf:"bytes,6,opt,name=timing,proto3" json:"timing,omitempty"`
	// Arbitrary name-value pairs, as reported in custom attributes or in a
	// <properties> element within this <testsuite>. Multiple properties are
	// allowed with the same key. Properties will be returned in lexicographical
	// order by key.
	Properties []*Property `protobuf:"bytes,7,rep,name=properties,proto3" json:"properties,omitempty"`
	// Files produced by this test suite, as reported by undeclared output
	// annotations.
	// The file IDs must be unique within this list. Duplicate file IDs will
	// result in an error. Files will be returned in lexicographical order by ID.
	Files []*File `protobuf:"bytes,8,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *TestSuite) Reset() {
	*x = TestSuite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestSuite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestSuite) ProtoMessage() {}

func (x *TestSuite) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestSuite.ProtoReflect.Descriptor instead.
func (*TestSuite) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_test_suite_proto_rawDescGZIP(), []int{0}
}

func (x *TestSuite) GetSuiteName() string {
	if x != nil {
		return x.SuiteName
	}
	return ""
}

func (x *TestSuite) GetTests() []*Test {
	if x != nil {
		return x.Tests
	}
	return nil
}

func (x *TestSuite) GetFailures() []*TestFailure {
	if x != nil {
		return x.Failures
	}
	return nil
}

func (x *TestSuite) GetErrors() []*TestError {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *TestSuite) GetTiming() *Timing {
	if x != nil {
		return x.Timing
	}
	return nil
}

func (x *TestSuite) GetProperties() []*Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *TestSuite) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

// The result of running a test case or test suite. JUnit3 TestDecorators are
// represented as a TestSuite with a single test.
type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Either a TestCase of a TestSuite
	//
	// Types that are assignable to TestType:
	//	*Test_TestCase
	//	*Test_TestSuite
	TestType isTest_TestType `protobuf_oneof:"test_type"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_test_suite_proto_rawDescGZIP(), []int{1}
}

func (m *Test) GetTestType() isTest_TestType {
	if m != nil {
		return m.TestType
	}
	return nil
}

func (x *Test) GetTestCase() *TestCase {
	if x, ok := x.GetTestType().(*Test_TestCase); ok {
		return x.TestCase
	}
	return nil
}

func (x *Test) GetTestSuite() *TestSuite {
	if x, ok := x.GetTestType().(*Test_TestSuite); ok {
		return x.TestSuite
	}
	return nil
}

type isTest_TestType interface {
	isTest_TestType()
}

type Test_TestCase struct {
	// When this contains just a single TestCase
	TestCase *TestCase `protobuf:"bytes,1,opt,name=test_case,json=testCase,proto3,oneof"`
}

type Test_TestSuite struct {
	// When this contains a TestSuite of test cases.
	TestSuite *TestSuite `protobuf:"bytes,2,opt,name=test_suite,json=testSuite,proto3,oneof"`
}

func (*Test_TestCase) isTest_TestType() {}

func (*Test_TestSuite) isTest_TestType() {}

// The result of running a test case, as reported in a <testcase> element of
// an XML log.
type TestCase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the test case, as reported in the name attribute. For Java,
	// this is normally the method name. Eg. "testBasic".
	CaseName string `protobuf:"bytes,1,opt,name=case_name,json=caseName,proto3" json:"case_name,omitempty"`
	// The name of the class in which the test case was defined, as reported in
	// the classname attribute. For Java, this is normally the fully qualified
	// class name. Eg. "com.google.common.hash.BloomFilterTest".
	ClassName string `protobuf:"bytes,2,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	// An enum reported in the result attribute that is used in conjunction with
	// failures and errors below to report the outcome.
	Result TestCase_Result `protobuf:"varint,3,opt,name=result,proto3,enum=google.devtools.resultstore.v2.TestCase_Result" json:"result,omitempty"`
	// Failures reported in <failure> elements within this <testcase>.
	Failures []*TestFailure `protobuf:"bytes,4,rep,name=failures,proto3" json:"failures,omitempty"`
	// Errors reported in <error> elements within this <testcase>.
	Errors []*TestError `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty"`
	// The timing for the TestCase, as reported by the time attribute.
	Timing *Timing `protobuf:"bytes,7,opt,name=timing,proto3" json:"timing,omitempty"`
	// Arbitrary name-value pairs, as reported in custom attributes or in a
	// <properties> element within this <testcase>. Multiple properties are
	// allowed with the same key. Properties will be returned in lexicographical
	// order by key.
	Properties []*Property `protobuf:"bytes,8,rep,name=properties,proto3" json:"properties,omitempty"`
	// Files produced by this test case, as reported by undeclared output
	// annotations.
	// The file IDs must be unique within this list. Duplicate file IDs will
	// result in an error. Files will be returned in lexicographical order by ID.
	Files []*File `protobuf:"bytes,9,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *TestCase) Reset() {
	*x = TestCase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestCase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCase) ProtoMessage() {}

func (x *TestCase) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCase.ProtoReflect.Descriptor instead.
func (*TestCase) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_test_suite_proto_rawDescGZIP(), []int{2}
}

func (x *TestCase) GetCaseName() string {
	if x != nil {
		return x.CaseName
	}
	return ""
}

func (x *TestCase) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *TestCase) GetResult() TestCase_Result {
	if x != nil {
		return x.Result
	}
	return TestCase_RESULT_UNSPECIFIED
}

func (x *TestCase) GetFailures() []*TestFailure {
	if x != nil {
		return x.Failures
	}
	return nil
}

func (x *TestCase) GetErrors() []*TestError {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *TestCase) GetTiming() *Timing {
	if x != nil {
		return x.Timing
	}
	return nil
}

func (x *TestCase) GetProperties() []*Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *TestCase) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

// Represents a violated assertion, as reported in a <failure> element within a
// <testcase>. Some languages allow assertions to be made without stopping the
// test case when they're violated, leading to multiple TestFailures. For Java,
// multiple TestFailures are used to represent a chained exception.
type TestFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The exception message reported in the message attribute. Typically short,
	// but may be multi-line. Eg. "Expected 'foo' but was 'bar'".
	FailureMessage string `protobuf:"bytes,1,opt,name=failure_message,json=failureMessage,proto3" json:"failure_message,omitempty"`
	// The type of the exception being thrown, reported in the type attribute.
	// Eg: "org.junit.ComparisonFailure"
	ExceptionType string `protobuf:"bytes,2,opt,name=exception_type,json=exceptionType,proto3" json:"exception_type,omitempty"`
	// The stack trace reported as the content of the <failure> element, often in
	// a CDATA block. This contains one line for each stack frame, each including
	// a method/function name, a class/file name, and a line number. Most recent
	// call is usually first, but not for Python stack traces. May contain the
	// exception_type and message.
	StackTrace string `protobuf:"bytes,3,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	// The expected values.
	//
	// These values can be diffed against the actual values. Often, there is just
	// one actual and one expected value. If there is more than one, they should
	// be compared as an unordered collection.
	Expected []string `protobuf:"bytes,4,rep,name=expected,proto3" json:"expected,omitempty"`
	// The actual values.
	//
	// These values can be diffed against the expected values. Often, there is
	// just one actual and one expected value. If there is more than one, they
	// should be compared as an unordered collection.
	Actual []string `protobuf:"bytes,5,rep,name=actual,proto3" json:"actual,omitempty"`
}

func (x *TestFailure) Reset() {
	*x = TestFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestFailure) ProtoMessage() {}

func (x *TestFailure) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestFailure.ProtoReflect.Descriptor instead.
func (*TestFailure) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_test_suite_proto_rawDescGZIP(), []int{3}
}

func (x *TestFailure) GetFailureMessage() string {
	if x != nil {
		return x.FailureMessage
	}
	return ""
}

func (x *TestFailure) GetExceptionType() string {
	if x != nil {
		return x.ExceptionType
	}
	return ""
}

func (x *TestFailure) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

func (x *TestFailure) GetExpected() []string {
	if x != nil {
		return x.Expected
	}
	return nil
}

func (x *TestFailure) GetActual() []string {
	if x != nil {
		return x.Actual
	}
	return nil
}

// Represents an exception that prevented a test case from completing, as
// reported in an <error> element within a <testcase>. For Java, multiple
// TestErrors are used to represent a chained exception.
type TestError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The exception message, as reported in the message attribute. Typically
	// short, but may be multi-line. Eg. "argument cannot be null".
	ErrorMessage string `protobuf:"bytes,1,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// The type of the exception being thrown, reported in the type attribute.
	// For Java, this is a fully qualified Throwable class name.
	// Eg: "java.lang.IllegalArgumentException"
	ExceptionType string `protobuf:"bytes,2,opt,name=exception_type,json=exceptionType,proto3" json:"exception_type,omitempty"`
	// The stack trace reported as the content of the <error> element, often in
	// a CDATA block. This contains one line for each stack frame, each including
	// a method/function name, a class/file name, and a line number. Most recent
	// call is usually first, but not for Python stack traces. May contain the
	// exception_type and message.
	StackTrace string `protobuf:"bytes,3,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
}

func (x *TestError) Reset() {
	*x = TestError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestError) ProtoMessage() {}

func (x *TestError) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestError.ProtoReflect.Descriptor instead.
func (*TestError) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_test_suite_proto_rawDescGZIP(), []int{4}
}

func (x *TestError) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *TestError) GetExceptionType() string {
	if x != nil {
		return x.ExceptionType
	}
	return ""
}

func (x *TestError) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

var File_google_devtools_resultstore_v2_test_suite_proto protoreflect.FileDescriptor

var file_google_devtools_resultstore_v2_test_suite_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x32, 0x1a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x03, 0x0a, 0x09, 0x54, 0x65,
	0x73, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x69, 0x74, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x69,
	0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x05, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x47, 0x0a, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65,
	0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x3e,
	0x0a, 0x06, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x54, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x48,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0a, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a,
	0x09, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x48, 0x00, 0x52, 0x08, 0x74, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73,
	0x75, 0x69, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x75, 0x69, 0x74, 0x65, 0x48, 0x00, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x53, 0x75, 0x69,
	0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22,
	0xdd, 0x04, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x47, 0x0a, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x3e, 0x0a,
	0x06, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54,
	0x69, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x48, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x22, 0x7a, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a,
	0x12, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x55, 0x50,
	0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x4b, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10, 0x05, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x55, 0x50, 0x50, 0x52, 0x45, 0x53, 0x53, 0x45, 0x44, 0x10, 0x06, 0x22,
	0xb2, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x63, 0x65,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x75, 0x61, 0x6c, 0x22, 0x78, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x42, 0x71,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65,
	0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x32, 0x3b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_resultstore_v2_test_suite_proto_rawDescOnce sync.Once
	file_google_devtools_resultstore_v2_test_suite_proto_rawDescData = file_google_devtools_resultstore_v2_test_suite_proto_rawDesc
)

func file_google_devtools_resultstore_v2_test_suite_proto_rawDescGZIP() []byte {
	file_google_devtools_resultstore_v2_test_suite_proto_rawDescOnce.Do(func() {
		file_google_devtools_resultstore_v2_test_suite_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_resultstore_v2_test_suite_proto_rawDescData)
	})
	return file_google_devtools_resultstore_v2_test_suite_proto_rawDescData
}

var file_google_devtools_resultstore_v2_test_suite_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_devtools_resultstore_v2_test_suite_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_devtools_resultstore_v2_test_suite_proto_goTypes = []interface{}{
	(TestCase_Result)(0), // 0: google.devtools.resultstore.v2.TestCase.Result
	(*TestSuite)(nil),    // 1: google.devtools.resultstore.v2.TestSuite
	(*Test)(nil),         // 2: google.devtools.resultstore.v2.Test
	(*TestCase)(nil),     // 3: google.devtools.resultstore.v2.TestCase
	(*TestFailure)(nil),  // 4: google.devtools.resultstore.v2.TestFailure
	(*TestError)(nil),    // 5: google.devtools.resultstore.v2.TestError
	(*Timing)(nil),       // 6: google.devtools.resultstore.v2.Timing
	(*Property)(nil),     // 7: google.devtools.resultstore.v2.Property
	(*File)(nil),         // 8: google.devtools.resultstore.v2.File
}
var file_google_devtools_resultstore_v2_test_suite_proto_depIdxs = []int32{
	2,  // 0: google.devtools.resultstore.v2.TestSuite.tests:type_name -> google.devtools.resultstore.v2.Test
	4,  // 1: google.devtools.resultstore.v2.TestSuite.failures:type_name -> google.devtools.resultstore.v2.TestFailure
	5,  // 2: google.devtools.resultstore.v2.TestSuite.errors:type_name -> google.devtools.resultstore.v2.TestError
	6,  // 3: google.devtools.resultstore.v2.TestSuite.timing:type_name -> google.devtools.resultstore.v2.Timing
	7,  // 4: google.devtools.resultstore.v2.TestSuite.properties:type_name -> google.devtools.resultstore.v2.Property
	8,  // 5: google.devtools.resultstore.v2.TestSuite.files:type_name -> google.devtools.resultstore.v2.File
	3,  // 6: google.devtools.resultstore.v2.Test.test_case:type_name -> google.devtools.resultstore.v2.TestCase
	1,  // 7: google.devtools.resultstore.v2.Test.test_suite:type_name -> google.devtools.resultstore.v2.TestSuite
	0,  // 8: google.devtools.resultstore.v2.TestCase.result:type_name -> google.devtools.resultstore.v2.TestCase.Result
	4,  // 9: google.devtools.resultstore.v2.TestCase.failures:type_name -> google.devtools.resultstore.v2.TestFailure
	5,  // 10: google.devtools.resultstore.v2.TestCase.errors:type_name -> google.devtools.resultstore.v2.TestError
	6,  // 11: google.devtools.resultstore.v2.TestCase.timing:type_name -> google.devtools.resultstore.v2.Timing
	7,  // 12: google.devtools.resultstore.v2.TestCase.properties:type_name -> google.devtools.resultstore.v2.Property
	8,  // 13: google.devtools.resultstore.v2.TestCase.files:type_name -> google.devtools.resultstore.v2.File
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_google_devtools_resultstore_v2_test_suite_proto_init() }
func file_google_devtools_resultstore_v2_test_suite_proto_init() {
	if File_google_devtools_resultstore_v2_test_suite_proto != nil {
		return
	}
	file_google_devtools_resultstore_v2_common_proto_init()
	file_google_devtools_resultstore_v2_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestSuite); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestCase); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestFailure); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_devtools_resultstore_v2_test_suite_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Test_TestCase)(nil),
		(*Test_TestSuite)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_devtools_resultstore_v2_test_suite_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_resultstore_v2_test_suite_proto_goTypes,
		DependencyIndexes: file_google_devtools_resultstore_v2_test_suite_proto_depIdxs,
		EnumInfos:         file_google_devtools_resultstore_v2_test_suite_proto_enumTypes,
		MessageInfos:      file_google_devtools_resultstore_v2_test_suite_proto_msgTypes,
	}.Build()
	File_google_devtools_resultstore_v2_test_suite_proto = out.File
	file_google_devtools_resultstore_v2_test_suite_proto_rawDesc = nil
	file_google_devtools_resultstore_v2_test_suite_proto_goTypes = nil
	file_google_devtools_resultstore_v2_test_suite_proto_depIdxs = nil
}
