package handler

import (
	"exchange-api/internal/svc"
	"net/http"
)

// OrderHandler 订单Handler
type OrderHandler struct {
	svcCtx *svc.ServiceContext
}

// NewOrderHandler 创建OrderHandler
func NewOrderHandler(svcCtx *svc.ServiceContext) *OrderHandler {
	return &OrderHandler{
		svcCtx: svcCtx, // 添加ServiceContext
	}
}

func (h *OrderHandler) History(w http.ResponseWriter, r *http.Request) {
}
