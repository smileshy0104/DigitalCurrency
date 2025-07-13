package logic

import (
	"context"
	"errors"
	"exchange/internal/domain"
	"exchange/internal/model"
	"exchange/internal/svc"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/exchange/types/order"
	"grpc-common/market/types/market"
	"grpc-common/ucenter/types/member"
)

// OrderLogic 用于处理汇率转换的逻辑
type OrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	orderDomain *domain.OrderDomain
}

// NewOrderLogic 创建一个新的 OrderLogic 实例
// 参数 ctx 是上下文环境信息
// 参数 svcCtx 是服务的上下文信息，包含了服务所需的各种配置和初始化信息
func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		orderDomain: domain.NewOrderDomain(svcCtx.Db),
	}
}

// FindOrderHistory 查询指定用户的历史订单
func (l *OrderLogic) FindOrderHistory(req *order.OrderReq) (*order.OrderRes, error) {
	// 查询对应用户的历史订单
	orderList, total, err := l.orderDomain.FindOrderHistory(l.ctx, req.Symbol, req.Page, req.PageSize, req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*order.ExchangeOrder
	copier.Copy(&list, orderList)
	return &order.OrderRes{
		List:  list,
		Total: total,
	}, nil
}

// FindOrderCurrent 查询当前用户订单（交易中）
func (l *OrderLogic) FindOrderCurrent(req *order.OrderReq) (*order.OrderRes, error) {
	// 查询对应用户的历史订单
	orderList, total, err := l.orderDomain.FindOrderCurrent(l.ctx, req.Symbol, req.Page, req.PageSize, req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*order.ExchangeOrder
	copier.Copy(&list, orderList)
	return &order.OrderRes{
		List:  list,
		Total: total,
	}, nil
}

// Add 添加订单
func (l *OrderLogic) Add(req *order.OrderReq) (*order.AddOrderRes, error) {
	// 1、通过用户id查询用户信息，判断用户是否存在
	memberInfo, err := l.svcCtx.MemberRpc.FindMemberById(l.ctx, &member.MemberReq{
		MemberId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	if memberInfo.TransactionStatus == 0 {
		return nil, errors.New("此用户已经被禁止交易！")
	}
	if req.Type == model.TypeMap[model.LimitPrice] && req.Price <= 0 {
		return nil, errors.New("限价模式下价格不能小于等于0！")
	}
	if req.Amount <= 0 {
		return nil, errors.New("购买数量不能小于等于0！")
	}
	// 2、通过交易币种名称获取币种信息
	exchangeCoinInfo, err := l.svcCtx.MarketRpc.FindSymbolInfo(l.ctx, &market.MarketReq{
		Symbol: req.Symbol, // 交易币种名称，格式：BTC/USDT
	})
	if err != nil {
		return nil, err
	}
	if exchangeCoinInfo.Exchangeable != 1 && exchangeCoinInfo.Enable != 1 {
		return nil, errors.New("该币种被禁用！")
	}
	// 获取基准币种，如BTC/USDT中的BTC
	baseSymbol := exchangeCoinInfo.GetBaseSymbol()
	// 获取交易币种，如BTC/USDT中的USDT
	coinSymbol := exchangeCoinInfo.GetCoinSymbol()
	unit := baseSymbol
	if req.Direction == model.DirectionMap[model.SELL] {
		//根据交易币查询
		unit = coinSymbol
	}
	// 查询货币信息
	coinInfo, err := l.svcCtx.MarketRpc.FindCoinInfo(l.ctx, &market.MarketReq{
		Unit: unit,
	})
	// 如果查询出错或货币信息不存在，返回错误
	if err != nil || coinInfo == nil {
		return nil, errors.New("该货币不存在！")
	}

	// 根据请求类型和方向检查交易限制
	if req.Type == model.TypeMap[model.MarketPrice] && req.Direction == model.DirectionMap[model.BUY] {
		// 对于买入操作，检查最小成交额限制
		if exchangeCoinInfo.GetMinTurnover() > 0 && req.Amount < float64(exchangeCoinInfo.GetMinTurnover()) {
			return nil, errors.New("成交额至少是" + fmt.Sprintf("%d", exchangeCoinInfo.GetMinTurnover()))
		}
	} else {
		// 对于其他操作，检查最大和最小交易量限制
		if exchangeCoinInfo.GetMaxVolume() > 0 && exchangeCoinInfo.GetMaxVolume() < req.Amount {
			return nil, errors.New("数量超出" + fmt.Sprintf("%f", exchangeCoinInfo.GetMaxVolume()))
		}
		if exchangeCoinInfo.GetMinVolume() > 0 && exchangeCoinInfo.GetMinVolume() > req.Amount {
			return nil, errors.New("数量不能低于" + fmt.Sprintf("%f", exchangeCoinInfo.GetMinVolume()))
		}
	}
	return &order.AddOrderRes{
		OrderId: "hhhhh",
	}, nil
}

func (l *OrderLogic) FindByOrderId(req *order.OrderReq) (*order.ExchangeOrderOrigin, error) {
	return nil, nil
}

func (l *OrderLogic) CancelOrder(req *order.OrderReq) (*order.CancelOrderRes, error) {
	return nil, nil
}
