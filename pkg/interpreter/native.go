package interpreter

type NativeFunc struct {
	name  string
	arity int
	fn    func(in *Interpreter, args []any) any
}

func (n *NativeFunc) Arity() int { return n.arity }
func (n *NativeFunc) Call(in *Interpreter, args []any) any {
	return n.fn(in, args)
}

func (in *Interpreter) defineNatives() {
	// TODO: add native functions here
}