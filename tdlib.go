// Package tdlib is an interface to tdlib's JSON client.
package tdlib

/*
#cgo LDFLAGS: -ltdjson
#include <stdlib.h>
#include <td/telegram/td_json_client.h>
*/
import "C"
import "unsafe"

// Client is an instance of tdlib's JSON client.
type Client struct {
	client unsafe.Pointer
}

// NewClient creates a new tdlib JSON client.
func NewClient() *Client {
	return &Client{
		client: C.td_json_client_create(),
	}
}

// Send sends a request to tdlib and does not get any response.
func (c *Client) Send(json []byte) {
	p := C.CString(string(json))
	defer C.free(unsafe.Pointer(p))

	C.td_json_client_send(c.client, p)
}

// Execute sends a request to tdlib and waits for the response.
//
// This method is synchronous and only a few may be executed at a time.
func (c *Client) Execute(json []byte) string {
	p := C.CString(string(json))
	defer C.free(unsafe.Pointer(p))

	str := C.td_json_client_execute(c.client, p)
	return C.GoString(str)
}

// Receive gets incoming updates and request responses from the client.
func (c *Client) Receive(timeout float32) string {
	str := C.td_json_client_receive(c.client, C.double(timeout))
	return C.GoString(str)
}

// Close destroys a tdlib client.
func (c *Client) Close() {
	C.td_json_client_destroy(c.client)
}
