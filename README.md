# filebeat-processors

> 一些适用于特定业务的 filebeat 自定义 processor

## 一、编译

- clone 官方 beats 仓库源码
- 编辑 `libbeat/cmd/instance/imports_common.go` 文件
- 将本仓库 processor 添加到包引入位置
- 进入 `beasts/filebeat` 目录执行 `make` 进行编译

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
+	_ "github.com/gozap/filebeat-processors/add_filename"
+	_ "github.com/gozap/filebeat-processors/split_message"
+	_ "github.com/gozap/filebeat-processors/regex_extract"
)
```

## 二、配置

目前所有 processor 的可用配置请参考每个目录下的 `config.go` 中的 `config` 结构体；例如:

``` go
type config struct {
	EnableTimestamp bool   `config:"enable_timestamp"`
	TimestampFormat string `config:"timestamp_format"`
	ProcessorsField string `config:"processors_field"`
	SourceField     string `config:"source_field"`
	TargetField     string `config:"target_field"`
	IgnoreMissing   bool   `config:"ignore_missing"`
	IgnoreFailure   bool   `config:"ignore_failure"`
}
```

其中 `defaultConfig` 方法为相关参数的默认值