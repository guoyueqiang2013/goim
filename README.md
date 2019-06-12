# goim


关键问题一、服务器负载分析<br/>
1.A发送一张图片512K<br/>
2.100人在线群人员同时接收到512kb*100=1024kb*50=50M<br/>
3.1024个群50M*1024=50G<br/>
解决方案：<br/>
1.使用缩略图(51.2K)提高单图下载和渲染速度<br/>
2.提高资源服务并发能力使用云服务(qos/alioss),100ms以内<br/>
3.压缩消息体，发送文件路径而不是整个文件<br/>


关键问题二、高并发<br/>
1.单机并发性性能最优<br/>
2.海量用户分布式部署<br/>
3.应对突发事件弹性扩容<br/>



//消息体核心
type Message struct {
	Id      int64  `json:"id,omitempty" form:"id"`           //消息id
	Userid  int64  `json:"userid,omitempty" form:"userid"`   //谁发的
	Cmd     int    `json:"cmd,omitempty" form:"cmd"`         //群聊还是私聊
	Dstid   int64  `json:"dstid,omitempty" form:"dstid"`     //对端ID 或 群ID
	Media   int    `json:"media,omitempty" form:"media"`     //消息样式
	Content string `json:"content,omitempty" form:"content"` //消息内容
	Pic     string `json:"pic,omitempty" form:"pic"`         //预览图片
	Url     string `json:"url,omitempty" form:"url"`         //服务的URL
	Memo    string `json:"memo,omitempty" form:"memo"`       //简单描述
	Amount  int    `json:"amount,omitempty" form:"amount"`   //和数字相关的
}



###IM系统的一般架构描述
前端： iOS , Android , Webapp , SDK , API , websocket<br/>
接入层：TCP , HTTPS , HTTP2 , websocket <br/>
逻辑层：鉴权,登录...,关系管理,群聊,单聊,消息上报,消息发下<br/>
存储层：Mysql,Redis,Mogondb ... , Hbase,Hive , 文件服务器<br/>

本系统已经实现： Webapp,websocket(前端),websocket(后端),鉴权,登录,关系管理,群聊,单聊,消息发下,Mysql,文件服务器<br/>


###网络结构
Hybrid App  ----|   websocket / http(s)               R/W            <br/>
浏览器 ----------|-----------------------> 应用服务器 <----->  数据库   <br/>
微信环境 --------|                                                    <br/>

http:提供api服务， websocket 提供长连接推送服务

###websocket心跳机制
距离最近一次发送30s后发送一次


###流程：A如何发送消息给B？



#优化方案：
###代码优化Map
1.使用读写锁
2.map不要太大

###突破系统瓶颈优化连接数
1.linux系统还是windows系统？
2.linux有最大文件数限制，需要解除

###降低对CPU资源的使用
1.降低JSON编码频次
2.一次编码多次使用

###降低对IO资源的使用
1.合并写数据库的次数
2.优化对数据库的读操作
3.能缓存的就缓存

###应用/资源服务分离
1.应用服务：系统提供动态服务
2.资源文件服务迁移到oss