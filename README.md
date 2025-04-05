# archiveVideo

一个用于视频文件归档和转换的Go语言工具。

## 功能特点

- 批量扫描指定目录下的视频文件
- 自动转换视频文件格式
- 错误处理和报告机制
- 支持多种视频格式

## 安装

确保您的系统已安装Go 1.24.1或更高版本，然后执行：

```bash
go get github.com/zhangyiming748/archiveVideos
```

## 使用方法

```go
package main

import "github.com/zhangyiming748/archiveVideos"

func main() {
    // 指定要处理的视频文件目录
    root := "path/to/your/videos"
    
    // 执行视频归档
    if errors := archiveVideos.ArchiveVideos(root); len(errors) > 0 {
        // 处理错误
        for _, err := range errors {
            fmt.Println(err)
        }
    }
}
```

## 依赖项

- github.com/h2non/filetype v1.1.3
- github.com/schollz/progressbar/v3 v3.18.0
- github.com/zhangyiming748/FastMediaInfo v0.0.7
- github.com/zhangyiming748/lumberjack v0.0.1
- github.com/zhangyiming748/pretty v0.0.3

## 测试

项目包含完整的测试用例，可以通过以下命令运行测试：

```bash
go test ./...
```

## 许可证

本项目采用MIT许可证。详情请参见[LICENSE](LICENSE)文件。