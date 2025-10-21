// Package utils provides utility functions for simplejsonx generic operations.
//
// utils 包为 simplejsonx 泛型操作提供实用工具函数
package utils

// Zero returns the zero value for any type T.
// Used to initialize return values in error cases for generic functions.
//
// Zero 返回任意类型 T 的零值
// 用于在泛型函数的错误情况下初始化返回值
func Zero[T any]() (x T) {
	return x
}
