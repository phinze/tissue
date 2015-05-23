package tissue

import "fmt"

func Shell() {
	fmt.Println(`
function tj() {
	local dir=$(tissue "$@")
	test -d "${dir}" && cd $dir
}`)
}
