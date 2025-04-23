package utils

import (
	"image/color"
	"log"
)

func HTML2FyneRGB(R, G, B int) color.RGBA {
	return HTML2FyneRGBA(R, G, B, 1.0)
}

func HTML2FyneRGBA(R, G, B int, A float32) color.RGBA {
	// 判断 RGB 三个是不是在 0-255之间
	if !rgbNumIsValid(R) ||
		!rgbNumIsValid(G) ||
		!rgbNumIsValid(B) ||
		!aNumIsValid(A) {
		log.Println("RGBA格式错误")
		return color.RGBA{}
	}

	a := uint8(A * 255)

	return color.RGBA{
		R: uint8(R),
		G: uint8(G),
		B: uint8(B),
		A: a,
	}
}

// 判断rgb对应的数值是否有效
func rgbNumIsValid(num int) bool {
	return num >= 0 && num <= 255

}

// 判断透明度是否有效
func aNumIsValid(num float32) bool {

	return num >= 0 && num <= 1
}

// HTML2FyneRGBA 将 HTML rgba 字符串转换为 Fyne RGBA 结构体
//func HTML2FyneRGBA(s string) color.RGBA {
//	s = strings.ReplaceAll(s, " ", "") // 去除空格
//	parts := strings.Split(strings.TrimSuffix(s, ")"), "rgba(")[1]
//
//	// 解析参数
//	args := strings.Split(parts, ",")
//	if len(args) < 3 || len(args) > 4 {
//		log.Println("invalid rgba format")
//		return color.RGBA{}
//	}
//
//	var r, g, b, a uint64
//	var err error
//
//	// 必须参数 R, G, B
//	r, err = strconv.ParseUint(args[0], 10, 8)
//	if err != nil {
//		log.Println("invalid red value: ", err)
//		return color.RGBA{}
//	}
//	g, err = strconv.ParseUint(args[1], 10, 8)
//	if err != nil {
//		log.Println("invalid red value: ", err)
//		return color.RGBA{}
//	}
//	b, err = strconv.ParseUint(args[2], 10, 8)
//	if err != nil {
//		log.Println("invalid red value: ", err)
//		return color.RGBA{}
//	}
//
//	// 可选透明度参数
//	if len(args) >= 4 {
//		alphaStr := strings.TrimSuffix(args[3], ")")
//		alphaFloat, err := strconv.ParseFloat(alphaStr, 64)
//		if err != nil {
//			log.Println("invalid red value: ", err)
//			return color.RGBA{}
//		}
//		a = uint64(alphaFloat * 255)
//	} else {
//		a = 255 // 默认完全不透明
//	}
//
//	return color.RGBA{
//		R: uint8(r),
//		G: uint8(g),
//		B: uint8(b),
//		A: uint8(a),
//	}
//}
