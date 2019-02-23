// Boost Software License - Version 1.0 - August 17th, 2003
//
// Permission is hereby granted, free of charge, to any person or organization
// obtaining a copy of the software and accompanying documentation covered by
// this license (the "Software") to use, reproduce, display, distribute,
// execute, and transmit the Software, and to prepare derivative works of the
// Software, and to permit third-parties to whom the Software is furnished to
// do so, all subject to the following:
//
// The copyright notices in the Software and this entire statement, including
// the above license grant, this restriction and the following disclaimer,
// must be included in all copies of the Software, in whole or in part, and
// all derivative works of the Software, unless such copies or derivative
// works are solely in the form of machine-executable object code generated by
// a source language processor.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE, TITLE AND NON-INFRINGEMENT. IN NO EVENT
// SHALL THE COPYRIGHT HOLDERS OR ANYONE DISTRIBUTING THE SOFTWARE BE LIABLE
// FOR ANY DAMAGES OR OTHER LIABILITY, WHETHER IN CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package openhmd

//#cgo LDFLAGS: -L. -lopenhmd
//#include "OpenHMD/include/openhmd.h"
import "C"

// Create makes an OpenHMD context.
// Returns nil if the context can't allocate memory.
func Create() *Context {
	ctx := C.ohmd_ctx_create()
	if ctx == nil {
		return nil
	}

	return &Context{c: ctx}
}

// Destroy removes the current OpenHMD context.
// Note: Your context will be nulled and all devices associated with the context are automatically closed.
func (c *Context) Destroy() {
	C.ohmd_ctx_destroy(c.c)
}

// GetError gets the last error as a human readable string.
func (c *Context) GetError() string {
	return C.GoString(C.ohmd_ctx_get_error(c.c))
}

// Update updates the current context
// to fetch the values for the devices handled by a context.
func (c *Context) Update() {
	C.ohmd_ctx_update(c.c)
}

// Probe for devices.
// Returns the number of devices found on the system.
func (c *Context) Probe() int {
	return int(C.ohmd_ctx_probe(c.c))
}

// ListGetString gets device description from enumeration list index.
func (c *Context) ListGetString(deviceIndex int, value StringValue) string {
	return C.GoString(C.ohmd_list_gets(c.c, C.int(deviceIndex), C.ohmd_string_value(value)))
}

// ListOpenDevice opens a device.
func (c *Context) ListOpenDevice(index int) *Device {
	return &Device{c: C.ohmd_list_open_device(c.c, C.int(index))}
}

// GetFloat gets a floating point value from a device.
func (d *Device) GetFloat(value FloatValue, out int) StatusCode {
	val := C.float(out)
	return StatusCode(C.ohmd_device_getf(d.c, C.ohmd_float_value(value), &val))
}

// SetFloat sets a floating point value for a device.
func (d *Device) SetFloat(value FloatValue, input int) StatusCode {
	val := C.float(input)
	return StatusCode(C.ohmd_device_setf(d.c, C.ohmd_float_value(value), &val))
}

// GetInt gets an integer value from a device.
func (d *Device) GetInt(value IntValue, out int) StatusCode {
	val := C.int(out)
	return StatusCode(C.ohmd_device_geti(d.c, C.ohmd_int_value(value), &val))
}
