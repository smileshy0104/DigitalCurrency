package logic

import (
	"common/db"
	"common/db/tran"
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
	"grpc-common/ucenter/types/asset"
	"grpc-common/ucenter/types/member"
)

// OrderLogic 用于处理汇率转换的逻辑
type OrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	orderDomain *domain.OrderDomain
	transaction tran.Transaction
	kafkaDomain *domain.KafkaDomain
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
		transaction: tran.NewTransaction(svcCtx.Db.Conn),
		kafkaDomain: domain.NewKafkaDomain(svcCtx.KafkaClient, domain.NewOrderDomain(svcCtx.Db)),
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
	// TODO 1、通过用户id查询用户信息，判断用户是否存在
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
	// TODO 2、通过交易币种名称获取币种信息
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
	// TODO 3、查询货币信息
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
	// TODO 4、查询用户钱包 BTC/USDT
	// 查询用户的钱包信息，包括基础货币（如BTC）和交易货币（如USDT）
	baseWalletInfo, err := l.svcCtx.AssetRpc.FindWalletBySymbol(l.ctx, &asset.AssetReq{
		UserId:   req.UserId, // 用户id
		CoinName: baseSymbol, // 货币名称 BTC/USDT
	})
	if err != nil {
		return nil, errors.New("用户钱包不存在该货币BTC！")
	}

	exCoinWalletInfo, err := l.svcCtx.AssetRpc.FindWalletBySymbol(l.ctx, &asset.AssetReq{
		UserId:   req.UserId, // 用户id
		CoinName: coinSymbol, // 货币名称 BTC/USDT
	})
	if err != nil {
		return nil, errors.New("用户钱包不存在该货币USDT！")
	}

	// 检查用户的钱包是否被锁定，如果任一钱包被锁定，则返回错误
	if baseWalletInfo.IsLock == 1 || exCoinWalletInfo.IsLock == 1 {
		return nil, errors.New("用户钱包被锁定！")
	}

	// 如果是卖出操作，检查是否低于最低限价
	if req.Direction == model.DirectionMap[model.SELL] && exchangeCoinInfo.GetMinSellPrice() > 0 {
		if req.Price < exchangeCoinInfo.GetMinSellPrice() || req.Type == model.TypeMap[model.MarketPrice] {
			return nil, errors.New("不能低于最低限价:" + fmt.Sprintf("%f", exchangeCoinInfo.GetMinSellPrice()))
		}
	}

	// 如果是买入操作，检查是否高于最高限价
	if req.Direction == model.DirectionMap[model.BUY] && exchangeCoinInfo.GetMaxBuyPrice() > 0 {
		if req.Price > exchangeCoinInfo.GetMaxBuyPrice() || req.Type == model.TypeMap[model.MarketPrice] {
			return nil, errors.New("不能低于最高限价:" + fmt.Sprintf("%f", exchangeCoinInfo.GetMaxBuyPrice()))
		}
	}

	// 检查是否启用了市价买卖，并根据买卖方向判断是否支持市价交易
	if req.Type == model.TypeMap[model.MarketPrice] {
		if req.Direction == model.DirectionMap[model.BUY] && exchangeCoinInfo.EnableMarketBuy == 0 {
			return nil, errors.New("不支持市价购买")
		} else if req.Direction == model.DirectionMap[model.SELL] && exchangeCoinInfo.EnableMarketSell == 0 {
			return nil, errors.New("不支持市价出售")
		}
	}

	// TODO 5、限制委托数量
	// FindCurrentTradingCount 查询当前用户在指定交易对下的订单数（交易中的订单）。
	count, err := l.orderDomain.FindCurrentTradingCount(l.ctx, req.UserId, req.Symbol, req.Direction)
	if err != nil {
		return nil, err
	}
	// 限制交易中的订单数
	if exchangeCoinInfo.GetMaxTradingOrder() > 0 && count >= exchangeCoinInfo.GetMaxTradingOrder() { // 最大允许同时交易的订单数，0表示不限制
		return nil, errors.New("超过最大挂单数量 " + fmt.Sprintf("%d", exchangeCoinInfo.GetMaxTradingOrder()))
	}

	// TODO 6、生成订单
	// 创建一个新的订单实例
	exchangeOrder := model.NewOrder()
	exchangeOrder.MemberId = req.UserId                     // 设置订单的用户ID
	exchangeOrder.Symbol = req.Symbol                       // 设置订单涉及的交易对符号
	exchangeOrder.BaseSymbol = baseSymbol                   // 设置基础货币符号
	exchangeOrder.CoinSymbol = coinSymbol                   // 设置计价货币符号
	typeCode := model.TypeMap.Code(req.Type)                // 根据请求类型获取订单类型代码
	exchangeOrder.Type = typeCode                           // 设置挂单类型 0 市场价 1 最低价
	directionCode := model.DirectionMap.Code(req.Direction) // 根据请求方向获取订单方向代码
	exchangeOrder.Direction = directionCode                 // 设置订单方向 0 买 1 卖
	// 根据订单类型设置价格：市价订单价格为0，其他类型订单价格为请求价格
	if exchangeOrder.Type == model.MarketPrice {
		exchangeOrder.Price = 0
	} else {
		exchangeOrder.Price = req.Price
	}
	exchangeOrder.UseDiscount = "0"   // 设置是否使用折扣：当前不使用折扣
	exchangeOrder.Amount = req.Amount // 设置买入或者卖出量

	// TODO 7、使用事务处理订单提交和消息发送的过程
	err = l.transaction.Action(func(conn db.DbConn) error {
		// 调用AddOrder方法，处理订单的创建和资金冻结
		money, err := l.orderDomain.AddOrder(l.ctx, conn, exchangeOrder, exchangeCoinInfo, baseWalletInfo, exCoinWalletInfo)
		if err != nil {
			// 如果订单提交失败，则返回错误
			return errors.New("订单提交失败")
		}
		fmt.Println("订单提交成功", money)
		// 通过kafka发消息，通知订单创建成功，此时钱包中的资金应该已被冻结
		err = l.kafkaDomain.SendOrderAdd(
			"add-exchange-order", req.UserId, exchangeOrder.OrderId, money,
			req.Symbol, exchangeOrder.Direction, baseSymbol, coinSymbol)
		if err != nil {
			// 如果消息发送失败，则返回错误
			return errors.New("发消息失败")
		}

		// 如果执行成功，返回nil表示事务可以提交
		return nil
	})
	if err != nil {
		// 如果事务执行失败，返回空响应体和错误
		return nil, err
	}

	// 如果事务执行成功，返回订单ID和空错误，表示订单创建成功
	return &order.AddOrderRes{
		OrderId: exchangeOrder.OrderId,
	}, nil
}

// FindByOrderId 根据订单ID查询订单详情。
// 该方法通过调用orderDomain中的FindOrderByOrderId方法来获取订单信息，
// 并将获取到的信息复制到ExchangeOrderOrigin对象中返回。
func (l *OrderLogic) FindByOrderId(req *order.OrderReq) (*order.ExchangeOrderOrigin, error) {
	// 调用orderDomain中的FindOrderByOrderId方法
	exchangeOrder, err := l.orderDomain.FindOrderByOrderId(l.ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	oo := &order.ExchangeOrderOrigin{}
	copier.Copy(oo, exchangeOrder)
	return oo, nil
}

// CancelOrder 取消指定的订单。
// 该方法通过调用orderDomain中的UpdateStatusCancel方法来更新订单的状态为取消。
func (l *OrderLogic) CancelOrder(req *order.OrderReq) (*order.CancelOrderRes, error) {
	// 调用orderDomain中的UpdateStatusCancel方法
	l.orderDomain.UpdateStatusCancel(l.ctx, req.OrderId)
	return nil, nil
}
