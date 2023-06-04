//go:generate stringer -type=Pill
package painkiller

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

//在 painkiller.go 文件所在的目录下运行 go generate 命令
//会在当前目录下面生成一个 pill_string.go 文件，文件中实现了我们需要的 String() 方法
