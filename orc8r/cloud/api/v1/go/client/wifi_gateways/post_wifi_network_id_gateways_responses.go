// Code generated by go-swagger; DO NOT EDIT.

package wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PostWifiNetworkIDGatewaysReader is a Reader for the PostWifiNetworkIDGateways structure.
type PostWifiNetworkIDGatewaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWifiNetworkIDGatewaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostWifiNetworkIDGatewaysCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostWifiNetworkIDGatewaysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWifiNetworkIDGatewaysCreated creates a PostWifiNetworkIDGatewaysCreated with default headers values
func NewPostWifiNetworkIDGatewaysCreated() *PostWifiNetworkIDGatewaysCreated {
	return &PostWifiNetworkIDGatewaysCreated{}
}

/*PostWifiNetworkIDGatewaysCreated handles this case with default header values.

Success
*/
type PostWifiNetworkIDGatewaysCreated struct {
}

func (o *PostWifiNetworkIDGatewaysCreated) Error() string {
	return fmt.Sprintf("[POST /wifi/{network_id}/gateways][%d] postWifiNetworkIdGatewaysCreated ", 201)
}

func (o *PostWifiNetworkIDGatewaysCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWifiNetworkIDGatewaysDefault creates a PostWifiNetworkIDGatewaysDefault with default headers values
func NewPostWifiNetworkIDGatewaysDefault(code int) *PostWifiNetworkIDGatewaysDefault {
	return &PostWifiNetworkIDGatewaysDefault{
		_statusCode: code,
	}
}

/*PostWifiNetworkIDGatewaysDefault handles this case with default header values.

Unexpected Error
*/
type PostWifiNetworkIDGatewaysDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post wifi network ID gateways default response
func (o *PostWifiNetworkIDGatewaysDefault) Code() int {
	return o._statusCode
}

func (o *PostWifiNetworkIDGatewaysDefault) Error() string {
	return fmt.Sprintf("[POST /wifi/{network_id}/gateways][%d] PostWifiNetworkIDGateways default  %+v", o._statusCode, o.Payload)
}

func (o *PostWifiNetworkIDGatewaysDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostWifiNetworkIDGatewaysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
