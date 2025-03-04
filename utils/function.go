package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math"
	"reflect"
)

func Md5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	hashedBytes := hash.Sum(nil)
	hashedString := fmt.Sprintf("%x", hashedBytes)
	return hashedString
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println("解码出错:", err)
		return "", err
	}
	decodedString := string(decodedBytes)
	return decodedString, nil
}

// Divide 计算两个数值的商，并根据需要返回小数或百分比
func Divide(a, b interface{}, returnPercentage bool, decimalPlaces int) (interface{}, error) {
	// 将参数转换为 float64 类型
	aVal, err := toFloat64(a)
	if err != nil {
		return nil, fmt.Errorf("第一个参数无效: %v", err)
	}

	bVal, err := toFloat64(b)
	if err != nil {
		return nil, fmt.Errorf("第二个参数无效: %v", err)
	}

	// 检查除数是否为零
	if bVal == 0 {
		return nil, fmt.Errorf("除数不能为零")
	}

	// 计算商
	result := aVal / bVal

	// 如果需要返回百分比，则将结果乘以 100
	if returnPercentage {
		result *= 100
	}

	// 格式化结果，保留指定小数位数
	roundedResult := round(result, decimalPlaces)

	return roundedResult, nil
}

// toFloat64 将任意数值类型转换为 float64
func toFloat64(value interface{}) (float64, error) {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	default:
		return 0, fmt.Errorf("不支持的类型: %v", v.Kind())
	}
}

// round 对浮点数进行四舍五入，保留指定小数位数
func round(value float64, decimalPlaces int) float64 {
	shift := math.Pow(10, float64(decimalPlaces))
	return math.Round(value*shift) / shift
}
