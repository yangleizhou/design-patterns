package structural

import "fmt"

//为多个复杂的子系统提供一个一致的接口，使这些子系统更加容易被访问

var _ FacadeService = (*OrderService)(nil)
var _ FacadeService = (*PayService)(nil)
var _ FacadeService = (*LogisticsService)(nil)

// FacadeService 外观角色，为子系统提供一致的接口
type FacadeService interface {
	Service()
}

// OrderService 订单服务-子系统角色对象
type OrderService struct {
}

//Service 一致接口实现
func (s *OrderService) Service() {
	fmt.Println("正在下单,准备跳转到支付流程...")
}

//PayService 支付服务-子系统角色对象
type PayService struct {
}

//Service 一致接口实现
func (s *PayService) Service() {
	fmt.Println("正在支付...支付成功！")
}

//LogisticsService 物流-子系统角色对象
type LogisticsService struct {
}

//Service 一致接口实现
func (s *LogisticsService) Service() {
	fmt.Println("进入物流系统，准备出库...")
}

//BuyService 购买服务 - 客户角色对象
type BuyService struct {
	os *OrderService
	ps *PayService
	ls *LogisticsService
}

//Service 接口实现
func (bs *BuyService) Service() {
	bs.os.Service()
	bs.ps.Service()
	bs.ls.Service()
}

// NewBuyService 购买服务对象
func NewBuyService() FacadeService {
	return &BuyService{
		os: &OrderService{},
		ps: &PayService{},
		ls: &LogisticsService{},
	}
}
