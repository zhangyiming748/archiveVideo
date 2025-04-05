package archiveVideos

import (
	"testing"
	"path/filepath"
)

func TestArchiveVideos(t *testing.T) {
	tests := []struct {
		name    string
		root    string
		wantErr bool
	}{
		{
			name:    "正常测试",
			root:    filepath.Join("testdata", "videos"),
			wantErr: false,
		},
		{
			name:    "空目录测试",
			root:    filepath.Join("testdata", "empty"),
			wantErr: true,
		},
		{
			name:    "不存在目录测试",
			root:    filepath.Join("testdata", "nonexistent"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := ArchiveVideos(tt.root)
			if (len(errors) > 0) != tt.wantErr {
				t.Errorf("ArchiveVideos() error = %v, wantErr %v", errors, tt.wantErr)
			}
		})
	}
}