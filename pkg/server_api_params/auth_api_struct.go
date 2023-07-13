// Copyright © 2023 OpenIM SDK. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server_api_params

import "github.com/OpenIMSDK/Open-IM-Server/pkg/proto/sdkws"

//UserID               string   `protobuf:"bytes,1,opt,name=UserID" json:"UserID,omitempty"`
//	Nickname             string   `protobuf:"bytes,2,opt,name=Nickname" json:"Nickname,omitempty"`
//	FaceUrl              string   `protobuf:"bytes,3,opt,name=FaceUrl" json:"FaceUrl,omitempty"`
//	Gender               int32    `protobuf:"varint,4,opt,name=Gender" json:"Gender,omitempty"`
//	PhoneNumber          string   `protobuf:"bytes,5,opt,name=PhoneNumber" json:"PhoneNumber,omitempty"`
//	Birth                string   `protobuf:"bytes,6,opt,name=Birth" json:"Birth,omitempty"`
//	Email                string   `protobuf:"bytes,7,opt,name=Email" json:"Email,omitempty"`
//	Ex                   string   `protobuf:"bytes,8,opt,name=Ex" json:"Ex,omitempty"`

type UserRegisterReq struct {
	Secret   string `json:"secret" binding:"required,max=32"`
	Platform int32  `json:"platform" binding:"required,min=1,max=7"`
	sdkws.UserInfo
	OperationID string `json:"operationID" binding:"required"`
}

type UserTokenInfo struct {
	UserID      string `json:"userID"`
	Token       string `json:"token"`
	ExpiredTime int64  `json:"expiredTime"`
}
type UserRegisterResp struct {
	CommResp
	UserToken UserTokenInfo `json:"data"`
}

type UserTokenReq struct {
	Secret      string `json:"secret" binding:"required,max=32"`
	Platform    int32  `json:"platformID" binding:"required,min=1,max=8"`
	UserID      string `json:"userID" binding:"required,min=1,max=64"`
	OperationID string `json:"operationID" binding:"required"`
}

type UserTokenResp struct {
	CommResp
	UserToken UserTokenInfo `json:"data"`
}

type ParseTokenReq struct {
	OperationID string `json:"operationID" binding:"required"`
}

type ExpireTime struct {
	ExpireTimeSeconds uint32 `json:"expireTimeSeconds" `
}

type ParseTokenResp struct {
	CommResp
	ExpireTime ExpireTime `json:"expireTime"`
}
