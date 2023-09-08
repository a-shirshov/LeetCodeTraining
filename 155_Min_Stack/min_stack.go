package main

type Node struct {
    value int
    minValue int
    previousNode *Node
}

type MinStack struct {
    topPointer *Node
}


func Constructor() MinStack {
    return MinStack{
        topPointer: nil,
    }
}


func (this *MinStack) Push(val int)  {
    if this.topPointer == nil {
        this.topPointer = &Node{
            value: val,
            minValue: val,
            previousNode: nil,
        }
    } else {
        newMinValue := this.topPointer.minValue 
        if (newMinValue > val) {
            newMinValue = val
        }
        this.topPointer = &Node{
			value: val,
			minValue: newMinValue,
			previousNode: this.topPointer,
		}
    }
}


func (this *MinStack) Pop()  {
    if (this.topPointer != nil) {
		this.topPointer = this.topPointer.previousNode
	}
}


func (this *MinStack) Top() int {
    return this.topPointer.value
}


func (this *MinStack) GetMin() int {
    return this.topPointer.minValue
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

 func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	println(stack.GetMin())
	stack.Pop()
	println(stack.Top())
	println(stack.GetMin())
 }