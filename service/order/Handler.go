package order

import (
	"github.com/BideWong/iStock/model"
	"github.com/BideWong/iStock/db"
)

/***
	userid: 用户id
	trade_type: 交易类型
	stockc_code: 股票代码
	stock_name：股票名
	stock_price：委托价
	stock_count：委托量（股）
	amount： 总金额
	stamp_tax：印花税
	transfer_tax：过户费
	brokerage：交易佣金

添加一条新的订单记录。
在数据库中将产生两个recode， Tb_order 存储的记录不会被更改
	Tb_order_real 存储的记录会在每次的部分成交中修改其 数量，这个记录也是定序系统（sequence）中排序依据，存储依据。
 */
func NewOrder(userid int, trade_type int, stock_code, stock_name string, stock_price float64, stock_count int,
	amount, stamp_tax, transfer_tax, brokerage float64) (model.Tb_order_real, error) {

	order := model.Tb_order{
		User_id : userid,
		Stock_name : stock_name,
		Stock_code : stock_code,
		Stock_price : stock_price,
		Stock_count : stock_count,
		Transfer_fee : transfer_tax,
		Brokerage : brokerage,
		Trade_type : trade_type,
		Trade_type_desc: "",   // 待填写
		Order_status : 0,
	}
	if order.Trade_type == model.TRADE_TYPE_BUY {
		order.Trade_type_desc = "买入"
		order.Freeze_amount = amount
	}else if order.Trade_type == model.TRADE_TYPE_SALE {
		order.Trade_type_desc = "卖出"
	}

	db.DBSession.Save(&order)

	order_real := model.Tb_order_real{
		Order_id : (int)(order.ID),
		User_id : userid,
		Stock_name : stock_name,
		Stock_code : stock_code,
		Stock_price : stock_price,
		Stock_count : stock_count,
		Order_status : 0,

		Trade_type : trade_type,
		Trade_type_desc: "",   // 待填写
	}
	if order_real.Trade_type == model.TRADE_TYPE_BUY {
		order_real.Trade_type_desc = "买入"
		order_real.Freeze_amount = amount
	}else if order_real.Trade_type == model.TRADE_TYPE_SALE {
		order_real.Trade_type_desc = "卖出"
	}

	db.DBSession.Save(&order_real)

	return order_real, nil
}