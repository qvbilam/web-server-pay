package enum

const PayStatusClosed = "CLOSED"     // 未付款交易超时关闭，或支付完成后全额退款
const PayStatusWait = "WAIT"         // 交易创建，等待买家付款
const PayStatusSuccess = "SUCCESS"   // 交易支付成功
const PayStatusFinished = "FINISHED" // 交易结束，不可退款
