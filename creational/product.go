package creational

type option func(*ComputerProduct)

// ComputerProduct 产品
type ComputerProduct struct {
	Type string
	CPU  string
	OS   string
}

// MakeProduct 组装赋值
func MakeProduct(cp *ComputerProduct, opts ...option) {
	for _, opt := range opts {
		opt(cp)
	}
}

// SetType 设置Type
func SetType(typ string) option {
	return func(cp *ComputerProduct) {
		cp.Type = typ
	}
}

// SetCPU 设置CPU
func SetCPU(cpu string) option {
	return func(cp *ComputerProduct) {
		cp.CPU = cpu
	}
}

// SetOS 设置操作系统
func SetOS(os string) option {
	return func(cp *ComputerProduct) {
		cp.OS = os
	}
}

// SetCPU 设置cpu
func (cp *ComputerProduct) SetCPU(cpu string) {
	cp.CPU = cpu
}

// SetOS 设置os
func (cp *ComputerProduct) SetOS(os string) {
	cp.OS = os
}

// SetType 设置type
func (cp *ComputerProduct) SetType(typ string) {
	cp.Type = typ
}
