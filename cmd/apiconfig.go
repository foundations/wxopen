// Copyright 2020 FastWeGo
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

package main

type Param struct {
	Name string
	Type string
}

type Api struct {
	Name        string
	Description string
	Request     string
	See         string
	FuncName    string
	GetParams   []Param
}

type ApiGroup struct {
	Name    string
	Apis    []Api
	Package string
}

var apiConfig = []ApiGroup{
	{
		Name:    `开放平台-授权`,
		Package: `auth`,
		Apis: []Api{
			{
				Name:        "获取 预授权码",
				Description: "预授权码（pre_auth_code）是第三方平台方实现授权托管的必备信息，每个预授权码有效期为 10 分钟。需要先获取令牌才能调用",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=COMPONENT_ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/pre_auth_code.html",
				FuncName:    "CreatePreauthCode",
			},
			{
				Name:        "方式一：授权注册页面扫码授权",
				Description: "第三方平台方可以在自己的网站中放置“微信公众号授权”或者“小程序授权”的入口，或生成授权链接放置在移动网页中，引导公众号和小程序管理员进入授权页。",
				Request:     "GET https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=xxxx&pre_auth_code=xxxxx&redirect_uri=xxxx&auth_type=xxx",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Authorization_Process_Technical_Description.html",
				FuncName:    "GetAuthorizationRedirectUri",
				GetParams: []Param{
					{Name: `component_appid`, Type: `string`},
					{Name: `pre_auth_code`, Type: `string`},
					{Name: `redirect_uri`, Type: `string`},
					{Name: `auth_type`, Type: `string`},
				},
			},
			{
				Name:        "方式二：点击移动端链接快速授权 ",
				Description: "第三方平台方可以生成授权链接，将链接通过移动端直接发给授权管理员，管理员确认后即授权成功",
				Request:     "GET https://mp.weixin.qq.com/safe/bindcomponent?action=bindcomponent&auth_type=3&no_scan=1&component_appid=xxxx&pre_auth_code=xxxxx&redirect_uri=xxxx&auth_type=xxx&biz_appid=xxxx#wechat_redirect",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Authorization_Process_Technical_Description.html",
				FuncName:    "GetAuthorizationRedirectUri2",
				GetParams: []Param{
					{Name: `component_appid`, Type: `string`},
					{Name: `pre_auth_code`, Type: `string`},
					{Name: `redirect_uri`, Type: `string`},
					{Name: `auth_type`, Type: `string`},
					{Name: `biz_appid`, Type: `string`},
				},
			},
			{
				Name:        "使用授权码获取授权信息",
				Description: "由当用户在第三方平台授权页中完成授权流程后，第三方平台开发者可以在回调 URI 中通过 URL 参数获取授权码。使用以下接口可以换取公众号/小程序的授权信息。建议保存授权信息中的刷新令牌（authorizer_refresh_token）",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=COMPONENT_ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/authorization_info.html",
				FuncName:    "",
			},
			{
				Name:        "获取/刷新接口调用令牌",
				Description: "在公众号/小程序接口调用令牌（authorizer_access_token）失效时，可以使用刷新令牌（authorizer_refresh_token）获取新的接口调用令牌",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=COMPONENT_ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/api_authorizer_token.html",
				FuncName:    "",
			},
			{
				Name:        "获取授权方的帐号基本信息",
				Description: "该 API 用于获取授权方的基本信息，包括头像、昵称、帐号类型、认证类型、微信号、原始ID和二维码图片URL",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=COMPONENT_ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/api_get_authorizer_info.html",
				FuncName:    "",
			},
			{
				Name:        "获取授权方选项信息",
				Description: "本 API 用于获取授权方的公众号/小程序的选项设置信息，如：地理位置上报，语音识别开关，多客服开关",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_option?component_access_token=COMPONENT_ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/api_get_authorizer_option.html",
				FuncName:    "",
			},
			{
				Name:        "设置授权方选项信息",
				Description: "本 API 用于设置授权方的公众号/小程序的选项信息，如：地理位置上报，语音识别开关，多客服开关",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/component/api_set_authorizer_option?component_access_token=COMPONENT_ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/api_set_authorizer_option.html",
				FuncName:    "",
			},
			{
				Name:        "拉取所有已授权的帐号信息",
				Description: "使用本 API 拉取当前所有已授权的帐号基本信息",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_list?component_access_token=COMPONENT_ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/api_get_authorizer_list.html",
				FuncName:    "",
			},
		},
	},
	{
		Name:    `开放平台-账号管理`,
		Package: `account`,
		Apis: []Api{
			{
				Name:        "创建开放平台帐号并绑定公众号/小程序",
				Description: "该 API 用于创建一个开放平台帐号，并将一个尚未绑定开放平台帐号的公众号/小程序绑定至该开放平台帐号上。新创建的开放平台帐号的主体信息将设置为与之绑定的公众号或小程序的主体",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/open/create?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/account/create.html",
				FuncName:    "",
			},
			{
				Name:        "将公众号/小程序绑定到开放平台帐号下",
				Description: "该 API 用于将一个尚未绑定开放平台帐号的公众号或小程序绑定至指定开放平台帐号上。二者须主体相同。",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/open/bind?access_token=xxxx",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/account/bind.html",
				FuncName:    "",
			},
			{
				Name:        "将公众号/小程序从开放平台帐号下解绑",
				Description: "该 API 用于将一个公众号或小程序与指定开放平台帐号解绑。开发者须确认所指定帐号与当前该公众号或小程序所绑定的开放平台帐号一致",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/open/unbind?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/account/unbind.html",
				FuncName:    "",
			},
			{
				Name:        "获取公众号/小程序所绑定的开放平台帐号",
				Description: "该 API 用于获取公众号或小程序所绑定的开放平台帐号",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/open/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/account/get.html",
				FuncName:    "",
			},
		},
	},
}
