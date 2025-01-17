// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package go_sms_sender

const (
	Aliyun       = "Aliyun SMS"
	TencentCloud = "Tencent Cloud SMS"
	VolcEngine   = "Volc Engine SMS"
)

type SmsClient interface {
	SendMessage(param map[string]string, targetPhoneNumber ...string)
}

func NewSmsClient(provider, accessId, accessKey, sign, region, template string, other ...string) SmsClient {
	switch provider {
	case Aliyun:
		return GetAliyunClient(accessId, accessKey, sign, region, template)
	case TencentCloud:
		return GetTencentClient(accessId, accessKey, sign, region, template, other)
	case VolcEngine:
		return GetVolcClient(accessId, accessKey, sign, region, template, other)
	default:
		return nil
	}
}

type SmsError struct {
	errorText string
}

func (e SmsError) Error() string {
	return e.errorText
}
