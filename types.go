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

/*
#include <openhmd/openhmd.h>
#cgo LDFLAGS: -L. -lopenhmd
*/
import "C"

// ArraySize defines all used Integer and Float array sizes.
type ArraySize int

// StatusCode - Response status codes.
type StatusCode C.ohmd_status

// StringValue - String values to fetch information.
type StringValue C.ohmd_string_value

// StringDescription - String descriptions for StringValues.
type StringDescription C.ohmd_string_description

// ControlHint - Control hints.
type ControlHint C.ohmd_control_hint

// ControlType - Control types.
type ControlType C.ohmd_control_type

// FloatValue - Floating point values.
type FloatValue C.ohmd_float_value

// IntValue - Integer values.
type IntValue C.ohmd_int_value

// DataValue - Data values.
type DataValue C.ohmd_data_value

// IntSettings - Integer-based settings.
type IntSettings C.ohmd_int_settings

// DeviceClass - Device classes for OpenHMD devices.
type DeviceClass C.ohmd_device_class

// DeviceFlags - Device flags for OpenHMD devices.
type DeviceFlags C.ohmd_device_flags

// Context holds contextial data.
type Context struct {
	c *C.struct_ohmd_context
}

// Device defines a OpenHMD device, like a HMD.
type Device struct {
	c *C.struct_ohmd_device
}

// DeviceSettings represents arguments for a Device.
type DeviceSettings struct {
	c *C.struct_ohmd_device_settings
}
