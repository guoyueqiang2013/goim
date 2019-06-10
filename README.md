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
