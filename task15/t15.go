package main

import (
	"fmt"
)

/*
1. Использование глобальной переменной justString. Недостатками использования глобальных переменных являются:
- неконтролируемый жизненный цикл
- невозможность иметь два экземпляра без изменения кода, использующего их
- усложнение читаемости кода из-за неявных связей
- неконтролируемый доступ

2. Когда в исходном коде мы присваиваем глобальной переменной justString значение v[:100], тем самым мы ссылаемся на область внутри SomeFunc из другого неймспейса.
В результате этого есть вероятность, что значение v не будет очищено сборщиком мусора и мы долгое время будем хранить пойнтер на память (вероятнее всего, аллоцированную в куче),
которая нам уже не требуется.

*/

func main() {
	var justString string = SomeFunc()
	fmt.Println(justString)

}

func SomeFunc() string {
	v := createHugeString(1 << 10)
	fmt.Printf("%T", v[:100])
	return v[:100]
}

func createHugeString(size int) string {
	return string(make([]rune, size))
}
