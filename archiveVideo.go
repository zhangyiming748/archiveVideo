package archiveVideos

import (
	"fmt"
	"github.com/zhangyiming748/archiveVideos/archive"
)

func ArchiveVideos(root string) []error {
	files, err := archive.GetAllFiles(root)
	var report []error
	if err != nil {
		report=append(report, fmt.Errorf("归档视频时,获取视频文件列表发生错误:%v", err))
		return report
	}
	
	for i, file := range files {
		if e := archive.Convert(file); e != nil {
			report=append(report, fmt.Errorf("归档第%d个视频时,转换视频文件发生错误:%v",i+1, e))
			continue
		}
	}
	return report
}
