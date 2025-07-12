package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/ucenter/types/asset"
	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

// Asset 结构体用于封装资产相关的业务逻辑
// 包含日志记录器、上下文和服务上下文
type Asset struct {
	logx.Logger                     // 日志记录器，支持日志打印功能
	ctx         context.Context     // 请求上下文，用于传递请求级数据
	svcCtx      *svc.ServiceContext // 服务上下文，包含服务配置和依赖项
}

// NewAssetLogic 创建并初始化Asset逻辑处理器实例
func NewAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Asset {
	// 创建Asset实例并注入上下文依赖
	return &Asset{
		Logger: logx.WithContext(ctx), // 创建带上下文信息的日志记录器
		ctx:    ctx,                   // 保存请求上下文
		svcCtx: svcCtx,                // 保存服务配置上下文
	}
}

// FindWalletBySymbol 根据货币符号查找用户的钱包信息。
// 该方法从当前上下文中提取用户ID，并结合请求中提供的货币名称，
// 通过RPC调用资产服务来获取用户对应货币的钱包信息。
// 主要解决了如何在大量用户和货币种类中高效查找特定用户和货币的钱包信息问题。
func (l *Asset) FindWalletBySymbol(req *types.AssetReq) (*types.MemberWallet, error) {
	// 创建一个带有超时的上下文，以确保请求不会无限期地等待。
	// 这里设置的超时时间是5秒，旨在防止在服务调用响应缓慢时导致资源浪费或潜在的死锁情况。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 在函数返回前取消创建的上下文，以释放相关资源。
	defer cancel()

	// 从当前上下文中提取用户ID。
	// 注意：这里假设了上下文中已经设置了"userId"，且其能成功转换为int64类型。
	value := l.ctx.Value("userId").(int64)

	// 通过RPC调用资产服务，查找用户指定货币的钱包信息。
	// 这里将用户ID和请求的货币名称作为参数传递给服务。
	memberWallet, err := l.svcCtx.UCAssetRpc.FindWalletBySymbol(ctx, &asset.AssetReq{
		CoinName: req.CoinName,
		UserId:   value,
	})
	if err != nil {
		// 如果发生错误，返回nil和错误信息。
		return nil, err
	}

	// 创建一个MemberWallet的响应对象。
	resp := &types.MemberWallet{}
	// 使用copier库将找到的钱包信息复制到响应对象中。
	// 这里使用copier是为了简化对象间的字段复制，提高代码的可读性和维护性。
	if err := copier.Copy(resp, memberWallet); err != nil {
		// 如果复制过程中发生错误，返回nil和错误信息。
		return nil, err
	}

	// 返回填充好的响应对象和nil错误，表示操作成功。
	return resp, nil
}

// FindWallet 根据用户ID查找钱包信息。
// 该方法从当前上下文中提取用户ID，并通过RPC调用获取用户的钱包信息。
func (l *Asset) FindWallet(req *types.AssetReq) ([]*types.MemberWallet, error) {
	// 创建一个带有超时的上下文，以确保请求不会无限期地等待。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 在函数返回时取消上下文，释放相关资源。
	defer cancel()

	// 从上下文中获取userId值，转换为int64类型。
	value := l.ctx.Value("userId").(int64)

	// 调用RPC服务，根据userId查找钱包信息。
	memberWalletResp, err := l.svcCtx.UCAssetRpc.FindWallet(ctx, &asset.AssetReq{
		UserId: value,
	})
	if err != nil {
		// 如果发生错误，返回nil和错误信息。
		return nil, err
	}

	// 初始化响应切片。
	var resp []*types.MemberWallet

	// 使用copier库复制钱包信息列表到响应切片。
	if err := copier.Copy(&resp, memberWalletResp.List); err != nil {
		// 如果复制过程中发生错误，返回nil和错误信息。
		return nil, err
	}

	// 返回钱包信息切片和nil错误，表示操作成功。
	return resp, nil
}

// ResetAddress 重置用户钱包地址
func (l *Asset) ResetAddress(req *types.AssetReq) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	value := l.ctx.Value("userId").(int64)
	// 调用RPC服务，将用户ID和币种名称作为参数传递给服务。
	_, err := l.svcCtx.UCAssetRpc.ResetAddress(ctx, &asset.AssetReq{
		UserId:   value,
		CoinName: req.Unit,
	})
	if err != nil {
		return "", err
	}
	return "", nil
}

//
//func (l *Asset) FindTransaction(req *types.AssetReq) (*pages.PageResult, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	value := l.ctx.Value("userId").(int64)
//	resp, err := l.svcCtx.UCAssetRpc.FindTransaction(ctx, &asset.AssetReq{
//		UserId:    value,
//		StartTime: req.StartTime,
//		EndTime:   req.EndTime,
//		PageNo:    int64(req.PageNo),
//		PageSize:  int64(req.PageSize),
//		Symbol:    req.Symbol,
//		Type:      req.Type,
//	})
//	if err != nil {
//		return nil, err
//	}
//	total := resp.Total
//	b := make([]any, len(resp.List))
//	for i, v := range resp.List {
//		b[i] = v
//	}
//	return pages.New(b, int64(req.PageNo), int64(req.PageSize), total), nil
//}
