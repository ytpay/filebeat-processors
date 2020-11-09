## regex_extract

> regex_extract processor 根据给定的正则表达式提取日志信息中指定字符串，并放入 event 中。

### 如何使用?

将此 processor 添加到 filebeat 后，你可以在 filebeat processors 配置段中增加以下配置:

``` yaml
processors:
  - regex_extract:
      # 源字段，regex_extract processor 从此字段读取日志信息，然后按照给定正则提取字符串
      # 此配置默认值为 "message"
      source_field: "message"
      # 目标字段，regex_extract processor 正则提取成功后将其写入到目标字段中
      # 此配置默认值为 "log_prefix"
      target_field: "timestamp"
      # 提取正则，默认为 java logback 日志时间戳
      regex: "[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}(?:.\\d{3}\\b)?"
      # 当无法找到 source_field 指定的字段时，如果该配置为 true，则忽略错误，继续处理 event
      # 此配置默认值为 false
      ignore_missing: true
      # 当出现一些错误时(例如上面的 source_field 找不到或者 source_field 不是个字符串等)忽略
      # 错误继续处理 event，可以将 ignore_failure 视为 ignore_missing 的更大范畴兼容
      # 此配置默认值为 true
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

如果 regex_extract processor 处理成功后应该可以在 logstash 控制台看到 `target_filed` 字段

``` diff
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
+    "timestamp" => "2020-11-09 11:54:09.687",
    "log_prefix" => "app"
}
```
