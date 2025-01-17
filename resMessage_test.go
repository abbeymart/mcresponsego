// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: mConnect standard transaction response testing | v0.9.1

package mcresponse

import (
	"github.com/abbeymart/mcresponse/messagecodes"
	"strconv"
	"strings"
	"testing"
)
import "github.com/abbeymart/mctest"

func TestResMessage(t *testing.T) {
	// test-data
	msgType := messagecodes.Success
	msgType2 := messagecodes.CheckError
	msgType3 := "custom"
	msg3 := "Custom Message"
	options := ResponseMessageOptions{
		Message: "",
		Value:   []string{"a", "b", "c"},
	}
	options2 := ResponseMessageOptions{
		Message: "",
		Value:   "",
	}
	options3 := ResponseMessageOptions{
		Message: msg3,
		Value:   "Custom",
	}
	res := ResponseMessage{
		Code:       messagecodes.Success,
		ResCode:    200,
		ResMessage: "OK",
		Value:      "",
		Message:    "Request completed successfully",
	}
	res2 := ResponseMessage{
		Code:       messagecodes.ParamsError,
		ResCode:    406,
		ResMessage: "Not Acceptable",
		Value:      "",
		Message:    "Parameters checking error",
	}
	res3 := ResponseMessage{
		Code:       "custom",
		ResCode:    200,
		ResMessage: "OK",
		Value:      "Custom",
		Message:    "Custom Message",
	}
	// Test cases
	mctest.McTest(mctest.OptionValue{
		Name: "should return success code for success-message",
		TestFunc: func() {
			req := GetResMessage(msgType, options)
			mctest.AssertEquals(t, req.Code, res.Code, "response-code should be: "+res.Code)
			mctest.AssertEquals(t, req.Message, res.Message, "response-message should be: "+res.Message)
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should return ok/200 resCode for success-message",
		TestFunc: func() {
			req := GetResMessage(msgType, ResponseMessageOptions{})
			mctest.AssertEquals(t, req.ResCode, res.ResCode, "response-code should be: "+strconv.Itoa(res.ResCode))
			mctest.AssertEquals(t, req.Message, res.Message, "response-message should be: "+res.Message)
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should return Completed successfully message for success-message",
		TestFunc: func() {
			req := GetResMessage(msgType, options)
			mctest.AssertEquals(t, req.Message, res.Message, "response-message should be: "+res.Message)
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should return correct default message",
		TestFunc: func() {
			options := ResponseMessageOptions{
				Value:   []string{"a", "b", "c"},
				Message: "Successful",
			}
			req := GetResMessage(msgType, options)
			mctest.AssertEquals(t, strings.Contains(req.Message, options.Message), true, "response-message should contains: "+options.Message)
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should return correct custom message",
		TestFunc: func() {
			req := GetResMessage(msgType3, options3)
			mctest.AssertEquals(t, req.Code, res3.Code, "response-code should be: "+res3.Code)
			mctest.AssertEquals(t, strings.Contains(req.Message, options3.Message), true, "response-message should contains: "+options3.Message)
		},
	})

	// check-error test-cases
	mctest.McTest(mctest.OptionValue{
		Name: "should return paramsError code for checkError-message",
		TestFunc: func() {
			req := GetResMessage(msgType2, options2)
			mctest.AssertEquals(t, req.Code, res2.Code, "response-code should be: "+res2.Code)
			mctest.AssertNotEquals(t, req.Code, "unAuthorized", "response-code"+req.Code+"should not be: unAuthorized")
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should return NOT_ACCEPTABLE/406 resCode",
		TestFunc: func() {
			req := GetResMessage(msgType2, ResponseMessageOptions{})
			mctest.AssertEquals(t, req.ResCode, res2.ResCode, "response-code should be: "+strconv.Itoa(res2.ResCode))
			mctest.AssertEquals(t, req.ResMessage, res2.ResMessage, "response-message should be: "+res2.ResMessage)
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should return Parameters checking error message",
		TestFunc: func() {
			req := GetResMessage(msgType2, options2)
			mctest.AssertEquals(t, req.Message, res2.Message, "response-code should be: "+res2.Message)
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should return authCode and auth-code message",
		TestFunc: func() {
			req := GetResMessage("authCode", ResponseMessageOptions{
				Message: "auth-code",
				Value:   nil,
			})
			mctest.AssertEquals(t, req.Code, "authCode", "response-code should be: authCode")
			mctest.AssertEquals(t, req.Message, "auth-code", "response-message should be: auth-code")
		},
	})

	mctest.PostTestResult()
}
