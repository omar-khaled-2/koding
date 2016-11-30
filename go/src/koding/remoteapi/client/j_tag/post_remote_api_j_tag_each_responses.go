package j_tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJTagEachReader is a Reader for the PostRemoteAPIJTagEach structure.
type PostRemoteAPIJTagEachReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJTagEachReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJTagEachOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJTagEachUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJTagEachOK creates a PostRemoteAPIJTagEachOK with default headers values
func NewPostRemoteAPIJTagEachOK() *PostRemoteAPIJTagEachOK {
	return &PostRemoteAPIJTagEachOK{}
}

/*PostRemoteAPIJTagEachOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPIJTagEachOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJTagEachOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTag.each][%d] postRemoteApiJTagEachOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJTagEachOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJTagEachUnauthorized creates a PostRemoteAPIJTagEachUnauthorized with default headers values
func NewPostRemoteAPIJTagEachUnauthorized() *PostRemoteAPIJTagEachUnauthorized {
	return &PostRemoteAPIJTagEachUnauthorized{}
}

/*PostRemoteAPIJTagEachUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJTagEachUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJTagEachUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTag.each][%d] postRemoteApiJTagEachUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJTagEachUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
