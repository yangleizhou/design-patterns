package creational

import (
	"fmt"
	"time"
)

//根据需求将预测的对象保存到channel中， 用于对象的生成成本大于维持成本
//缓存通道实现可复用对象池

//channel使用注意
// 1.不要在读取端关闭 channel ，因为写入端无法知道 channel 是否已经关闭，往已关闭的 channel 写数据会 panic
// 2.有多个写入端时，不要在写入端关闭 channle ，因为其他写入端无法知道 channel 是否已经关闭，关闭已经关闭的 channel 会发生 panic
// 3.如果只有一个写入端，可以在这个写入端放心关闭 channel 。
//   关闭 channel 粗暴一点的做法是随意关闭，如果产生了 panic 就用 recover 避免进程挂掉。
//   稍好一点的方案是使用标准库的 sync 包来做关闭 channel 时的协程同步，不过使用起来也稍微复杂些

var _ ReUsable = (*reUsableObj)(nil)

// ReUsable 定义对象公共接口
type ReUsable interface {
	GetName() string
}

// NewPool 创建对象池
func NewPool(count int) *PoolManager {
	pm := &PoolManager{}
	pm.pool = make(chan ReUsable, count)
	for i := 0; i < count; i++ {
		pm.pool <- &reUsableObj{name: fmt.Sprintf("%d", i)}
	}
	return pm
}

// reUsableObj 对象
type reUsableObj struct {
	name string
}

func (r *reUsableObj) GetName() string {
	return r.name
}

//PoolManager 对象池
type PoolManager struct {
	pool chan ReUsable
}

//GetObj 从池子获取对象
func (p *PoolManager) GetObj(timeout time.Duration) (ReUsable, error) {
	select {
	case ret, ok := <-p.pool:
		{
			if !ok {
				break
			}
			return ret, nil
		}
	case <-time.After(timeout):
		{
			return nil, fmt.Errorf("time out")
		}
	}
	return nil, nil
}

// ReleaseObj 放对象到池子
func (p *PoolManager) ReleaseObj(ru ReUsable) error {
	select {
	case p.pool <- ru:
		return nil
	default:
		return fmt.Errorf("overflow")
	}
}
