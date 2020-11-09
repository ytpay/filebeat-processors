# filebeat-processors

> 一些适用于特定业务的 filebeat 自定义 processor

## 一、编译

### 1.1、Docker 编译

在安装好 docker 的 Linux 机器上，执行目录下的 `build.sh` 既可完成编译，编译完成二进制文件存放在 `build` 目录中。

### 1.2、手动编译

- 确保配置好相关 go 开发环境
- clone 官方 beats 仓库源码
- 编辑 `libbeat/cmd/instance/imports_common.go` 文件
- 将本仓库 processor 添加到包引入位置
- 进入 `beasts/filebeat` 目录执行 `make crosscompile` 进行编译

修改样例:

``` diff
package instance

import (
	_ "github.com/elastic/beats/v7/libbeat/autodiscover/appenders/config" // Register autodiscover appenders
	_ "github.com/elastic/beats/v7/libbeat/autodiscover/providers/jolokia"
	_ "github.com/elastic/beats/v7/libbeat/monitoring/report/elasticsearch" // Register default monitoring reporting
	_ "github.com/elastic/beats/v7/libbeat/processors/actions"              // Register default processors.
	_ "github.com/elastic/beats/v7/libbeat/processors/add_cloud_metadata"
	_ "github.com/elastic/beats/v7/libbeat/processors/add_host_metadata"
	_ "github.com/elastic/beats/v7/libbeat/processors/add_id"
	_ "github.com/elastic/beats/v7/libbeat/processors/add_locale"
	_ "github.com/elastic/beats/v7/libbeat/processors/add_observer_metadata"
	_ "github.com/elastic/beats/v7/libbeat/processors/add_process_metadata"
	_ "github.com/elastic/beats/v7/libbeat/processors/communityid"
	_ "github.com/elastic/beats/v7/libbeat/processors/convert"
	_ "github.com/elastic/beats/v7/libbeat/processors/dissect"
	_ "github.com/elastic/beats/v7/libbeat/processors/dns"
	_ "github.com/elastic/beats/v7/libbeat/processors/extract_array"
	_ "github.com/elastic/beats/v7/libbeat/processors/fingerprint"
	_ "github.com/elastic/beats/v7/libbeat/processors/registered_domain"
	_ "github.com/elastic/beats/v7/libbeat/processors/translate_sid"
	_ "github.com/elastic/beats/v7/libbeat/publisher/includes" // Register publisher pipeline modules
+	_ "github.com/ytpay/filebeat-processors/add_filename"
+	_ "github.com/ytpay/filebeat-processors/add_prefix"
+	_ "github.com/ytpay/filebeat-processors/add_log_type"
+	_ "github.com/ytpay/filebeat-processors/split_message"
+	_ "github.com/ytpay/filebeat-processors/regex_extract"
)
```

## 二、配置

关于每个 processor 的具体配置请查看 processor 目录下的 README.md 文档，以下为 processor 列表:

- [add_filename](https://github.com/ytpay/filebeat-processors/blob/master/add_filename/README.md)
- [add_log_type](https://github.com/ytpay/filebeat-processors/blob/master/add_log_type/README.md)
- [add_prefix](https://github.com/ytpay/filebeat-processors/blob/master/add_prefix/README.md)
- [regex_extract](https://github.com/ytpay/filebeat-processors/blob/master/regex_extract/README.md)
- [split_message](https://github.com/ytpay/filebeat-processors/blob/master/split_message/README.md)
