package main

/*
#include <stdio.h>

// 方法1: 直接在注释块中定义C函数
void cHello() {
    printf("Hello World from C!\n");
}

int cAdd(int a, int b) {
    return a + b;
}
*/
import "C"
import (
	"fmt"
)

func main() {
	// 需要安装gcc并且配置环境变量 go env -w CGO_ENABLED=1
	// 方法1: 调用在注释块中定义的C函数
	C.cHello()
	// 调用C函数并获取返回值
	a := C.int(5)
	b := C.int(3)
	result := C.cAdd(a, b)
	fmt.Printf("cAdd(5, 3) = %d\n", result)

	// 方法2: 直接调用C标准库函数
	cstring := C.CString("hello world")
	C.printf(cstring)

	// 方法3: 传递参数给C函数
	//name := C.CString("Go Developer")
	//defer C.free(unsafe.Pointer(name)) // 记得释放C分配的内存
	//C.printf(C.CString("Hello, %s!\n"), name)

	fmt.Println("Hello from Go!")
}
