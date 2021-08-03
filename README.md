# 使用阿里云函数计算搭建静态博客评论系统

给一个静态博客添加评论系统是一个比较麻烦的事情

- 使用第三方的系统，比如 Disqus、多说、GitHub Issue 等，可能会面临服务在国外比较慢、服务商停止运营等问题。
- 使用开源系统自己搭建，比如 Valine 等，是比较复杂的，而且为了评论再搞一台服务器也不值当。

在这种情况下，如果可以将服务搭建在按需付费的服务上，就方便很多了，本项目选择的就是阿里云函数计算，也就是大家常说的 serverless 服务。
它的计费项目包括外网流量、计算和存储，都是按量付费的。

本项目基本架构比较简单，没有考虑太多可配置型，如果有更复杂的需求，可以借鉴思路重新开发。

 - 运行环境：阿里云函数计算 https://www.aliyun.com/product/fc
 - 编程语言：Golang
 - 数据库：kv 数据库 https://github.com/etcd-io/bbolt
 - 存储：阿里云 NAS https://www.aliyun.com/product/nas

## 注意事项

 - 阿里云函数计算需要绑定已经备案的域名，如果没有，可以跳过本项目了。
 - `template.yml` 中有些配置和我的账号是绑定的，需要自行修改。
 - 依赖 fun 工具 https://help.aliyun.com/document_detail/64204.htm
 - 本项目只是 api，前端界面为非常简单的 html，见 https://github.com/virusdefender/strcpy.me/blob/master/comment.html