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

// PutWifiNetworkIDGatewaysGatewayIDWifiReader is a Reader for the PutWifiNetworkIDGatewaysGatewayIDWifi structure.
type PutWifiNetworkIDGatewaysGatewayIDWifiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWifiNetworkIDGatewaysGatewayIDWifiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutWifiNetworkIDGatewaysGatewayIDWifiNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutWifiNetworkIDGatewaysGatewayIDWifiDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutWifiNetworkIDGatewaysGatewayIDWifiNoContent creates a PutWifiNetworkIDGatewaysGatewayIDWifiNoContent with default headers values
func NewPutWifiNetworkIDGatewaysGatewayIDWifiNoContent() *PutWifiNetworkIDGatewaysGatewayIDWifiNoContent {
	return &PutWifiNetworkIDGatewaysGatewayIDWifiNoContent{}
}

/*PutWifiNetworkIDGatewaysGatewayIDWifiNoContent handles this case with default header values.

Success
*/
type PutWifiNetworkIDGatewaysGatewayIDWifiNoContent struct {
}

func (o *PutWifiNetworkIDGatewaysGatewayIDWifiNoContent) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/gateways/{gateway_id}/wifi][%d] putWifiNetworkIdGatewaysGatewayIdWifiNoContent ", 204)
}

func (o *PutWifiNetworkIDGatewaysGatewayIDWifiNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutWifiNetworkIDGatewaysGatewayIDWifiDefault creates a PutWifiNetworkIDGatewaysGatewayIDWifiDefault with default headers values
func NewPutWifiNetworkIDGatewaysGatewayIDWifiDefault(code int) *PutWifiNetworkIDGatewaysGatewayIDWifiDefault {
	return &PutWifiNetworkIDGatewaysGatewayIDWifiDefault{
		_statusCode: code,
	}
}

/*PutWifiNetworkIDGatewaysGatewayIDWifiDefault handles this case with default header values.

Unexpected Error
*/
type PutWifiNetworkIDGatewaysGatewayIDWifiDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put wifi network ID gateways gateway ID wifi default response
func (o *PutWifiNetworkIDGatewaysGatewayIDWifiDefault) Code() int {
	return o._statusCode
}

func (o *PutWifiNetworkIDGatewaysGatewayIDWifiDefault) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/gateways/{gateway_id}/wifi][%d] PutWifiNetworkIDGatewaysGatewayIDWifi default  %+v", o._statusCode, o.Payload)
}

func (o *PutWifiNetworkIDGatewaysGatewayIDWifiDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutWifiNetworkIDGatewaysGatewayIDWifiDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
