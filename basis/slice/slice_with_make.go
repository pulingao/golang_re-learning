package slice

import "fmt"

/**
 * 字符分割线（character dividing line）
 */
func NewLineWithCDL() {
	fmt.Println()
	fmt.Println("// ---------------------------------------------------------------------------------------------------------------------")
	fmt.Println("//     _____             __   _____       _")
	fmt.Println("//    / ____|           / /  / ____|     | |")
	fmt.Println("//   | |  __  ___      / /  | |  __  ___ | | __ _ _ __   __ _")
	fmt.Println("//   | | |_ |/ _ \\    / /   | | |_ |/ _ \\| |/ _` | '_ \\ / _` |")
	fmt.Println("//   | |__| | (_) |  / /    | |__| | (_) | | (_| | | | | (_| |")
	fmt.Println("//    \\_____|\\___/  /_/      \\_____|\\___/|_|\\__,_|_| |_|\\__, |")
	fmt.Println("//                                                       __/ |")
	fmt.Println("//                                                      |___/")
	fmt.Println("// ---------------------------------------------------------------------------------------------------------------------")
	fmt.Println()
}

/**
 * make初始化slice时，make([]int, size, cap)
 *
 * @see
 * @param
 */
func SliceWithMakeInit() {
	s1 := make([]int, 0, 3)
	s1 = append(s1, 1, 2)
	for i, v := range s1 {
		fmt.Printf("i(%v) = \"%v\"\n", i, v)
	}

	NewLineWithCDL()

	s2 := make([]int, 2, 5)
	s2 = append(s2, 1, 2)
	for i, v := range s2 {
		fmt.Printf("i(%v) = \"%v\"\n", i, v)
	}

	NewLineWithCDL()

	s3 := make([]string, 3, 9)
	s3 = append(s3, "中国", "人嘛", "就是")
	for i, v := range s3 {
		fmt.Printf("i(%v) = \"%v\"\n", i, v)
	}
}
