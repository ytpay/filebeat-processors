## add_log_type

> add_log_type processor 用于提取文件名后缀并从 map 映射中确定日志文件类型。

### 如何使用?

将此 processor 添加到 filebeat 后，你可以在 filebeat processors 配置段中增加以下配置:

``` yaml
processors:
  - add_log_type:
      # 源字段，add_log_type processor 从此字段读取到一个文件名，然后提取文件名的扩展名
      # 此配置默认值为 "filename"
      source_field: "filename"
      # 目标字段，add_log_type processor 使用文件的扩展名(例如 ".error") 从映射 map 中
      # 获取该日志的类型，然后将其写入到目标字段中
      # 此配置默认值为 "log_type"
      target_field: "log_type"
      # 扩展名到日志类型的 `key <==> value` 映射
      type_map:
        .log: "log"
        .sqllog: "sql"
        .error: "error"
        .json: "json"
      # processor 标记位，add_log_type processor 处理成功后会将此字段设置为 true
      # 通常该字段用于标识作用，方便后面的 logstash 判断 event 是否被某个 processor 处理过
      # 此配置默认值为 "processors.add_log_type"
      processors_field: "processors.add_log_type"
      # 当无法找到 source_field 指定的字段时，如果该配置为 true，则忽略错误，继续处理 event
      # 此配置默认值为 false
      ignore_missing: true
      # 当出现一些错误时(例如上面的 source_field 找不到或者 source_field 不是个字符串等)忽略
      # 错误继续处理 event，可以将 ignore_failure 视为 ignore_missing 的更大范畴兼容
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

如果 add_log_type processor 处理成功后应该可以在 logstash 控制台看到 `target_filed` 字段

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
+     "log_type" => "log",
    "processors" => {
        "add_filename" => true,
          "add_prefix" => true,
+       "add_log_type" => true
    },
    "log_prefix" => "app"
}
```
