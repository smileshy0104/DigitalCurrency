// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3
// Source: member.proto

package uc_client

import (
	"context"
	"grpc-common/ucenter/types/member"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	MemberInfo = member.MemberInfo
	MemberReq  = member.MemberReq

	Member interface {
		FindMemberById(ctx context.Context, in *MemberReq, opts ...grpc.CallOption) (*MemberInfo, error)
	}

	defaultMember struct {
		cli zrpc.Client
	}
)

func NewMember(cli zrpc.Client) Member {
	return &defaultMember{
		cli: cli,
	}
}

func (m *defaultMember) FindMemberById(ctx context.Context, in *MemberReq, opts ...grpc.CallOption) (*MemberInfo, error) {
	client := member.NewMemberClient(m.cli.Conn())
	return client.FindMemberById(ctx, in, opts...)
}
