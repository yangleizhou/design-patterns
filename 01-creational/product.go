package creational

import "fmt"

type option func(*ComputerProduct)
type phoneOption func(*PhoneProduct)

// ComputerProduct 电脑产品
type ComputerProduct struct {
	Type string
	CPU  string
	OS   string
}

// MakeComputerProduct 组装赋值
func MakeComputerProduct(cp *ComputerProduct, opts ...option) {
	for _, opt := range opts {
		opt(cp)
	}
}

// SetComputerType 设置电脑Type
func SetComputerType(typ string) option {
	return func(cp *ComputerProduct) {
		cp.Type = typ
	}
}

// SetComputerCPU 设置电脑CPU
func SetComputerCPU(cpu string) option {
	return func(cp *ComputerProduct) {
		cp.CPU = cpu
	}
}

// SetComputerOS 设置电脑操作系统
func SetComputerOS(os string) option {
	return func(cp *ComputerProduct) {
		cp.OS = os
	}
}

// SetComputerCPU 设置电脑cpu
func (cp *ComputerProduct) SetComputerCPU(cpu string) {
	cp.CPU = cpu
}

// SetComputerOS 设置电脑os
func (cp *ComputerProduct) SetComputerOS(os string) {
	cp.OS = os
}

// SetComputerType 设置电脑type
func (cp *ComputerProduct) SetComputerType(typ string) {
	cp.Type = typ
}

// GetProduct 产品
func (cp *ComputerProduct) GetProduct() {
	fmt.Println("mac 电脑")
}

// PhoneProduct 手机产品
type PhoneProduct struct {
	Type string
	CPU  string
	OS   string
}

// MakePhoneProduct 组装赋值
func MakePhoneProduct(cp *PhoneProduct, opts ...phoneOption) {
	for _, opt := range opts {
		opt(cp)
	}
}

// SetPhoneType 设置手机Type
func SetPhoneType(typ string) phoneOption {
	return func(cp *PhoneProduct) {
		cp.Type = typ
	}
}

// SetPhoneCPU 设置手机CPU
func SetPhoneCPU(cpu string) phoneOption {
	return func(cp *PhoneProduct) {
		cp.CPU = cpu
	}
}

// SetPhoneOS 设置手机操作系统
func SetPhoneOS(os string) phoneOption {
	return func(cp *PhoneProduct) {
		cp.OS = os
	}
}

// SetPhoneCPU 设置手机cpu
func (cp *PhoneProduct) SetPhoneCPU(cpu string) {
	cp.CPU = cpu
}

// SetPhoneOS 设置手机os
func (cp *PhoneProduct) SetPhoneOS(os string) {
	cp.OS = os
}

// SetPhoneType 设置手机type
func (cp *PhoneProduct) SetPhoneType(typ string) {
	cp.Type = typ
}

// GetProduct 产品
func (cp *PhoneProduct) GetProduct() {
	fmt.Println("ios 手机")
}
