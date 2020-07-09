package regexextract

import (
	"testing"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/stretchr/testify/assert"
)

func newTestRegexExtract(t testing.TB, regex, field string, target string, ignoreMissing bool) *RegexExtract {
	c, err := common.NewConfigFrom(map[string]interface{}{
		"regex":         regex,
		"field":         field,
		"target":        target,
		"ignoreMissing": ignoreMissing,
	})
	if err != nil {
		t.Fatal(err)
	}

	f, err := New(c)
	if err != nil {
		t.Fatal(err)
	}
	return f.(*RegexExtract)
}

func TestRegexExtract(t *testing.T) {
	var tests = []struct {
		Value, Regex, Field, Target, Result string
		Error                               bool
		IngoreMissing                       bool
	}{
		{
			Value:         "2020-07-07 18:10:42.909 [6f8574460f0f16cb/6f8574460f0f16cb] [http-nio-8080-exec-4] [INFO ] c.y.service.AppUserTOTPKeyService [http-nio-8080-exec-4] [com.ytpay.service.AppUserTOTPKeyService.getPayCode(AppUserTOTPKeyService.java:52)] - 生成支付码[userId=5,cardNo=1001005020000004713]",
			Regex:         "[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}(?:.\\d{3}\\b)?",
			Field:         "message",
			Result:        "2020-07-07 18:10:42.909",
			Target:        "date",
			IngoreMissing: true,
			Error:         false,
		},
		{
			Value:         "2020-07-07 18:10:42 [6f8574460f0f16cb/6f8574460f0f16cb] [http-nio-8080-exec-4] [INFO ] c.y.service.AppUserTOTPKeyService [http-nio-8080-exec-4] [com.ytpay.service.AppUserTOTPKeyService.getPayCode(AppUserTOTPKeyService.java:52)] - 生成支付码[userId=5,cardNo=1001005020000004713]",
			Regex:         "balabala",
			Field:         "message",
			Result:        "2020-07-07 18:10:42",
			Target:        "date",
			IngoreMissing: false,
			Error:         true,
		},
	}

	for _, test := range tests {
		f := newTestRegexExtract(t, test.Regex, test.Field, test.Target, test.IngoreMissing)
		event := &beat.Event{Fields: common.MapStr{"message": test.Value}}
		event, err := f.Run(event)
		if test.Error {
			assert.NotNil(t, err)
		} else {
			if assert.NoError(t, err) {
				assert.Equal(t,
					test.Result,
					event.Fields[test.Target])
			}
		}
	}
}