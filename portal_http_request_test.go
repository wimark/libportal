package libportal

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestPortalHTTPRequest(t *testing.T) {
	var testCases = []struct {
		r   PortalHTTPRequest
		err string
	}{
		{
			r: PortalHTTPRequest{
				Method: http.MethodGet,
				Url:    "https://localhost:3000?foo=foo&bar=1",
			},
			err: "",
		},
		{
			r: PortalHTTPRequest{
				Method: http.MethodPut,
				Url:    "https://localhost:3000?q=foo;q=bar&a=1",
			},
			err: ErrMethodNotSupported.Error(),
		},
		{
			r: PortalHTTPRequest{
				Method:      http.MethodPost,
				Url:         "https://localhost:3000/api/examples",
				ContentType: "application/json",
				Headers: map[string]string{
					"Authorization": "Bearer b1946ac92492d2347c6235b4d2611184",
				},
				Params: map[string]string{
					"x": "1",
					"y": "2",
				},
			},
			err: "",
		},
		{
			r: PortalHTTPRequest{
				Method:      http.MethodPost,
				Url:         "https://localhost:3000/api/examples",
				ContentType: "application/x-www-form-urlencoded",
				Headers: map[string]string{
					"Authorization": "Bearer b1946ac92492d2347c6235b4d2611184",
				},
				Params: map[string]string{
					"x": "1",
					"y": "2",
				},
			},
			err: "",
		},
		{
			r: PortalHTTPRequest{
				Method:      http.MethodPost,
				Url:         "https://localhost:3000/api/examples",
				ContentType: "multipart/form-data",
				Headers: map[string]string{
					"Authorization": "Bearer b1946ac92492d2347c6235b4d2611184",
				},
				Params: map[string]string{
					"x": "1",
					"y": "2",
				},
			},
			err: "",
		},
	}
	for i, test := range testCases {
		name := fmt.Sprintf("test %d method[%s]", i+1, test.r.Method)
		if test.r.ContentType != "" {
			name = fmt.Sprintf("%s content-type[%s]", name, test.r.ContentType)
		}
		t.Run(name, func(t *testing.T) {
			req, err := test.r.Build()
			if err != nil {
				if test.err == err.Error() {
					return
				}
				t.Errorf("invalid error ex: %s got: %s", test.err, err.Error())
				return
			}
			compare(t, test.r, req)
		})
	}

}

func compare(t *testing.T, r PortalHTTPRequest, req *http.Request) {
	t.Helper()

	// validate http method
	if r.Method != req.Method {
		t.Errorf("invalid method ex: %s got: %s", r.Method, req.Method)
	}

	// validate content-type
	ct := req.Header.Get("Content-Type")
	if r.ContentType != ct && !strings.HasPrefix(ct, "multipart/") {
		t.Errorf("invalid contenty-type ex: %s got: %s", r.ContentType, ct)
	}

	// validate headers
	for key, value := range r.Headers {
		header := req.Header.Get(key)
		if value != header {
			t.Errorf("invalid header ex: %s got: %s", value, header)
		}
	}

	// validate post form value
	if r.Method == http.MethodPost {
		getValidator(ct)(t, r.Params, req)
	}
}

type validator func(t *testing.T, params map[string]string, req *http.Request)

var contentTypeValidation = map[string]validator{
	applicationJSON:           validateJSONBody,
	applicationFormURLEncoded: validateFormURLEncoded,
	multipartFormData:         validateMultiPartForm,
}

func getValidator(ct string) validator {
	if strings.HasPrefix(ct, "multipart/") {
		return contentTypeValidation[multipartFormData]
	}
	return contentTypeValidation[ct]
}

func validateJSONBody(t *testing.T, params map[string]string, req *http.Request) {
	t.Helper()

	var testMap = make(map[string]string)
	if err := json.NewDecoder(req.Body).Decode(&testMap); err != nil {
		t.Error(err.Error())
	}

	if len(testMap) != len(params) {
		t.Errorf("bad count variables ex: %d got: %d", len(params), len(testMap))
	}

	for key, value := range params {
		v, ok := testMap[key]
		if !ok {
			t.Errorf("json params %s with value %s not found", key, value)
		}
		if value != v {
			t.Errorf("json params not valid ex: %s=%s got: %s=%s", key, value, key, v)
		}
	}

}
func validateFormURLEncoded(t *testing.T, params map[string]string, req *http.Request) {
	t.Helper()
	data := url.Values{}
	for key, value := range params {
		data.Set(key, value)
	}
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Errorf("read error %s", err.Error())
	}
	if string(b) != data.Encode() {
		t.Errorf("params failed ex: %s got: %s", string(b), data.Encode())
	}
}
func validateMultiPartForm(t *testing.T, params map[string]string, req *http.Request) {
	t.Helper()

	_, ctParams, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		t.Errorf("invalid parse media type %s", err.Error())
	}

	r := multipart.NewReader(req.Body, ctParams["boundary"])
	for {
		p, err := r.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err.Error())
		}
		slurp, err := io.ReadAll(p)
		if err != nil {
			t.Error(err.Error())
		}
		key := p.FormName()
		if v := params[key]; v != string(slurp) {
			t.Errorf("multipart/form-data params not valid ex: %s=%s got: %s=%s", key, v, key, string(slurp))
		}
		delete(params, key)
	}

	if len(params) != 0 {
		var str strings.Builder
		for key, value := range params {
			str.WriteString(key + "=" + value)
			str.WriteByte('\n')
		}

		t.Errorf("not found params %s", str.String()[:str.Len()-1])
	}
}
