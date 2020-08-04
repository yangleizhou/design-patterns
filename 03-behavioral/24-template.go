package behavioral

import "fmt"

//定义一个操作中的算法骨架，而将算法的一些步骤延迟到子类中，使得子类可以不改变该算法结构的情况下重定义该算法的某些特定步骤

var _ Downloader = (*template)(nil)     //实现了Downloader模板方法
var _ Template = (*template)(nil)       //实现了Template默认基本方法
var _ Template = (*HTTPDownloader)(nil) //匿名组合template
var _ Template = (*FTPDownloader)(nil)
var _ Template = (*P2PDownloader)(nil)

//Downloader 下载器,基本流程包括下载和保存，模板提取公共方法
type Downloader interface {
	TemplateMethod(uri string) //模板方法
}

// Template 抽象类，负责给出一个算法的轮廓和骨架。它由一个模板方法和若干个基本方法构成
type Template interface {
	Download() //基本方法
	Save()     //基本方法
}

type template struct {
	Template //父类持有子类的引用
	uri      string
}

//TemplateMethod 一个模板方法，定义了算法的骨架，按某种顺序调用其包含的基本方法
func (t *template) TemplateMethod(uri string) {
	t.uri = uri
	fmt.Printf("prepare downloading ...\n")
	fmt.Println(t.uri)
	t.Template.Download()
	t.Template.Save()
	fmt.Printf("finish download !\n")
}

// newTemplate 获取模板对象
func newTemplate(tmp Template) *template {
	return &template{Template: tmp}
}

// Download 模板提供默认Download方法
func (t *template) Download() {
	fmt.Printf("default download uri = %s\n", t.uri)
}

// Save 模板提供默认Save方法
func (t *template) Save() {
	fmt.Printf("defalut save!")
}

// HTTPDownloader http下载
// golang 没有继承，采用匿名组合
type HTTPDownloader struct {
	*template
}

// newHTTPDownloader http下载器
func newHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	downloader.template = newTemplate(downloader)
	return downloader
}

// FTPDownloader ftp下载
type FTPDownloader struct {
	*template
}

// newFTPDownloader ftp下载器
func newFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	downloader.template = newTemplate(downloader)
	return downloader
}

// Download 实现Download接口
func (d *FTPDownloader) Download() {
	fmt.Printf("ftp download !\n")
}

// P2PDownloader p2p下载
type P2PDownloader struct {
	*template
}

// newP2PDownloader p2p下载器
func newP2PDownloader() Downloader {
	downloader := &P2PDownloader{}
	downloader.template = newTemplate(downloader)
	return downloader
}

// Download 实现Download接口
func (d *P2PDownloader) Download() {
	fmt.Printf("P2P download !\n")
}

// Save 实现Save接口
func (d *P2PDownloader) Save() {
	fmt.Printf("p2p Save !\n")
}
