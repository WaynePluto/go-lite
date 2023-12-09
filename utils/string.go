package utils

import "strings"

// 将path路径分割为一个节点名称切片
func SplitPathToParts(path string) []string {
	parts := []string{"/"}
	arr := strings.Split(path, "/")
	for _, v := range arr {
		if v != "" {
			parts = append(parts, v)
		}
	}
	return parts
}

// 将路径分割为多个子路径
func SplitPathToPaths(path string) []string {
	parts := SplitPathToParts(path)
	for i := 1; i < len(parts); i++ {
		if i == 1 {
			parts[i] = "/" + parts[i]
		} else {
			parts[i] = parts[i-1] + "/" + parts[i]
		}
	}
	return parts
}
