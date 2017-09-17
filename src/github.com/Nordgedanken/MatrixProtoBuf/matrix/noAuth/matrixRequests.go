package noAuth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Nordgedanken/MatrixProtoBuf/jsonTypes"
	"io/ioutil"
	"net/http"
)

func Login(address string, req *jsonTypes.ReqLogin, resp *jsonTypes.RespLogin) (err error, errCode int64) {
	var data []byte
	data, err = json.Marshal(req)
	if err != nil {
		errCode = 500
		return
	}

	reqresp, err := http.Post("https://"+address+"/_matrix/client/r0/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		errCode = 500
		return
	}

	defer reqresp.Body.Close()
	body, err := ioutil.ReadAll(reqresp.Body)

	if reqresp.StatusCode/100 != 2 { // not 2xx
		var wrap error
		var respErr jsonTypes.RespError
		if _ = json.Unmarshal(body, &respErr); respErr.ErrCode != "" {
			fmt.Errorf(respErr.Error())
			errCode = 3
			err = errors.New("failed to interpret JSON ERROR from Login Endpoint of the defined HS")
			return
		}

		// If we failed to decode as RespError, don't just drop the HTTP body, include it in the
		// HTTP error instead (e.g proxy errors which return HTML).
		msg := "Failed to POST JSON to " + reqresp.Request.URL.String()
		if wrap == nil {
			msg = msg + ": " + string(body)
		}

		errCode = int64(reqresp.StatusCode)
		err = errors.New(msg)
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		fmt.Errorf(err.Error())
		errCode = 500
		err = errors.New("failed to interpret JSON from Login Endpoint of the defined HS")
		return
	}
	return
}
