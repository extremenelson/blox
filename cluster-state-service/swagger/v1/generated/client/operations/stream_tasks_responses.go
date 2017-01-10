// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StreamTasksReader is a Reader for the StreamTasks structure.
type StreamTasksReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *StreamTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStreamTasksOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewStreamTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStreamTasksOK creates a StreamTasksOK with default headers values
func NewStreamTasksOK(writer io.Writer) *StreamTasksOK {
	return &StreamTasksOK{
		Payload: writer,
	}
}

/*StreamTasksOK handles this case with default header values.

Stream tasks - success
*/
type StreamTasksOK struct {
	Payload io.Writer
}

func (o *StreamTasksOK) Error() string {
	return fmt.Sprintf("[GET /stream/tasks][%d] streamTasksOK  %+v", 200, o.Payload)
}

func (o *StreamTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamTasksInternalServerError creates a StreamTasksInternalServerError with default headers values
func NewStreamTasksInternalServerError() *StreamTasksInternalServerError {
	return &StreamTasksInternalServerError{}
}

/*StreamTasksInternalServerError handles this case with default header values.

Stream tasks - unexpected error
*/
type StreamTasksInternalServerError struct {
	Payload string
}

func (o *StreamTasksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /stream/tasks][%d] streamTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *StreamTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}