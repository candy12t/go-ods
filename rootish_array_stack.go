package ods

type rootishArrayStack struct {
	len    int
	blocks []*arrayStack
}
