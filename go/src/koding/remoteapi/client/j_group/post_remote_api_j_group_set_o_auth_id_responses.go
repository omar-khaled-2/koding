package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJGroupSetOAuthIDReader is a Reader for the PostRemoteAPIJGroupSetOAuthID structure.
type PostRemoteAPIJGroupSetOAuthIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJGroupSetOAuthIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJGroupSetOAuthIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJGroupSetOAuthIDOK creates a PostRemoteAPIJGroupSetOAuthIDOK with default headers values
func NewPostRemoteAPIJGroupSetOAuthIDOK() *PostRemoteAPIJGroupSetOAuthIDOK {
	return &PostRemoteAPIJGroupSetOAuthIDOK{}
}

/*PostRemoteAPIJGroupSetOAuthIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJGroupSetOAuthIDOK struct {
	Payload *models.JGroup
}

func (o *PostRemoteAPIJGroupSetOAuthIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.setOAuth/{id}][%d] postRemoteApiJGroupSetOAuthIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJGroupSetOAuthIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
