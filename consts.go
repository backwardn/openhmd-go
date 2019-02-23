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

// StringSize defines the maximum length of a string, including termination, in OpenHMD.
const StringSize = C.OHMD_STR_SIZE

// All known OpenHMD StatusCodes.
const (
	StatusCodeOkay             StatusCode = 0
	StatusCodeUnknownError     StatusCode = -1
	StatusCodeInvalidParameter StatusCode = -2
	StatusCodeUnsupported      StatusCode = -3
	StatusCodeInvalidOperation StatusCode = -4

	StatusCodeUserReserved StatusCode = -16384
)

// All known StringValues.
const (
	StringValueVendor C.ohmd_string_value = iota
	StringValueProduct
	StringValuePath
)

// All known StringDescriptions.
const (
	StringDescriptionGlslDisortionVertSrc StringDescription = iota
	StringDescriptionGlslDisortionFragSRC
	StringDescriptionGsls330DisortionVertSrc
	StringDescriptionGsls330DisortionFragSrc
	StringDescriptionGslsEsDisortionVertSrc
	StringDescriptionGslsEsDisortionFragSrc
)

// All known ControlHints.
const (
	ControlHintGeneric ControlHint = iota
	ControlHintTrigger
	ControlHintTriggerClick
	ControlHintSqueeze
	ControlHintMenu
	ControlHintHome
	ControlHintAnalogX
	ControlHintAnalogY
	ControlHintAnalogPress
	ControlHintButtonA
	ControlHintButtonB
	ControlHintButtonX
	ControlHintButtonY
	ControlHintVolumePlus
	ControlHintVolumeMinus
	ControlHintMicMute
)

// All known ControlTypes.
const (
	ControlTypeDigital ControlType = iota
	ControlTypeAnalog
)

// All known FloatValues.
const (
	FloatValueRotationQuat FloatValue = iota + 1
	FloatValueLeftEyeGlModelViewMatrix
	FloatValueRightEyeGlModelViewMatrix
	FloatValueLeftEyeGlProjectionMatrix
	FloatValueRightEyeGlProjectionMatrix
	FloatValuePositionVector
	FloatValueScreenHorizontalSize
	FloatValueScreenVerticalSize
	FloatValueLensHorizontalSeparation
	FloatValueLensVerticalPosition
	FloatValueLeftEyeFOV
	FloatValueLeftEyeAspectRatio
	FloatValueRightEyeFov
	FloatValueRightEyeAspectRatio
	FloatValueEyeIpd
	FloatValueProjectionZFar
	FloatValueProjectionZNear
	FloatValueDistortionK
	FloatValueExternalSensorFusion
	FloatValueUniversalDisortionK
	FloatValueUniversalAberrationK
	FloatValueControlsState
)

// All known IntValues.
const (
	IntValueScreenHorizontalResolution IntValue = iota
	IntValueScreenVerticalResolution
	IntValueDeviceClass
	IntValueDeviceFlags
	IntValueControlsCount
	IntValueControlsHints
	IntValueControlsTypes
)

// All known DeviceClasses.
const (
	DeviceClassHMD DeviceClass = iota
	DeviceClassController
	DeviceClassGenericTracker
)

// All known IntSettings.
const (
	IntSettingsIdsAutomaticUpdate IntSettings = iota
)

// All known DataValues.
const (
	DataValueDriverData DataValue = iota
	DataValueDriverProperties
)

// All known DeviceFlags.
const (
	DeviceFlagsNullDevice         DeviceFlags = 1
	DeviceFlagsPositionalTracking DeviceFlags = 2
	DeviceFlagsRotationalTracking DeviceFlags = 4
	DeviceFlagsLeftController     DeviceFlags = 8
	DeviceFlagsRightController    DeviceFlags = 16
)
