## add_filename

> `add_filename` processor 用于在 event 中添加文件名字段。

### 如何使用?

将此 processor 添加到 filebeat 后，你可以在 filebeat processors 配置段中增加以下配置:

``` yaml
processors:
  - add_filename:
      # 源字段，add_filename processor 从次字段读取到一个绝对路径，然后提取文件名
      # 此配置默认值为 "log.file.path"
      source_field: "log.file.path"
      # 目标字段，add_filename processor 将文件名提取成功后写入目标字段
      # 为保证不会覆盖其他 processor 或系统生成字段，推荐将其放入特定结构字段下
      # 此配置默认值为 "filename"，即 event 根路径下的 "filename"
      target_field: "custom.filename"
      # processor 标记位，add_filename processor 处理成功后会将此字段设置为 true
      # 通常该字段用于标识作用，方便后面的 logstash 判断 event 是否被某个 processor 处理过
      # 此配置默认值为 "processors.add_filename"
      processors_field: "processors.add_filename"
      # 文件名增加时间戳，如果设置为 true，add_filename processor 文件名提取成功后
      # 会在文件名前缀与扩展名之间增加时间戳，时间戳格式由 timestamp_format 提供
      # 例如原始文件名提取后为 "test-api.log"，开启时间戳后可能转换为 "test-api-20200102.log"
      # 此配置默认值为 false
      enable_timestamp: false
      # 文件名时间戳格式，用于在 enable_timestamp 开启时控制时间戳格式
      # 该格式请参考 golang 时间戳格式化相关文章(不要尝试使用 java 的 yyyyMMdd，根本不支持)
      # 此配置默认值为 "2006-01-02"
      timestamp_format: "2006-01-02"
      # 当无法找到 source_field 指定的字段时，如果该配置为 true，则忽略错误，继续处理 event
      # 此配置默认值为 false
      ignore_missing: true
      # 当出现一些错误时(例如上面的 source_field 找不到或者 source_field 不是个字符串等)忽略
      # 错误继续出 event，可以将 ignore_failure 视为 ignore_missing 的更大范畴兼容
      # 此配置默认值为 false
      ignore_failure: true
```

### 如何调试?

你可以为 logstash 开启终端输出来实时观察日志处理情况:

``` sh
output {
  stdout {
    codec => rubydebug
  }
}
```

如果 add_filename processor 处理成功后应该可以在 logstash 控制台看到 `target_filed` 字段

``` ruby
{
           "ecs" => {
        "version" => "1.5.0"
    },
          "tags" => [
        [0] "beats_input_codec_plain_applied"
    ],
       "message" => "2020-11-09 11:54:09.687 app-78b956cf7f-rtk7w [http-nio-8080-exec-9] INFO  c.y.m.i.AuthenticationInterceptor.preHandle - allow request"
      "@version" => "1",
      "filename" => "app.2020-11-09.micro-app-78b956cf7f-rtk7w.log",
          "file" => {
            "path" => "/data/logs/app/app.2020-11-09.micro-app-78b956cf7f-rtk7w.log"
        },
        "offset" => 204335
    },
         "input" => {
        "type" => "log"
    },
    "@timestamp" => 2020-11-09T03:54:13.272Z,
         "agent" => {
                  "id" => "4e67cd3c-a53c-48c1-b898-716539a083d3",
                "type" => "filebeat",
                "name" => "k8s23",
            "hostname" => "k8s23",
        "ephemeral_id" => "e5092345-6762-4410-9bf7-8ca84620764f",
             "version" => "7.9.3"
    },
      "log_type" => "log",
    "processors" => {
        "add_filename" => true,
          "add_prefix" => true,
        "add_log_type" => true
    },
    "log_prefix" => "app"
}
```
