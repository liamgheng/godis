package protocol

const CLCR = "\r\n"

type MultiBulk struct {
	items [][]byte
}

// NOTE 这些划线区分增加可读性
/* ---- Status Reply ---- */

// type StatusReply struct {
// 	status string
// }

// func MakeStatusReply(status string) *StatusReply {
// 	return &StatusReply{status: status}
// }

// func (s *StatusReply) ToBytes() []byte {
// 	return []byte("+" + s.status + CLCR)
// }

// /* ---- Error Reply ---- */

// type ErrorReply struct {
// 	status string
// }

// func MakeErrorReply(status string) *ErrorReply {
// 	return &ErrorReply{status: status}
// }

// func (e *ErrorReply) ToBytes() []byte {
// 	return []byte("-" + e.status + CLCR)
// }

// /* ---- Int Reply ---- */

// type IntReply struct {
// 	num int64
// }

// func MakeIntReply(num int64) *IntReply {
// 	return &IntReply{num: num}
// }

// func (i *IntReply) ToBytes() []byte {
// 	return []byte(":" + strconv.FormatInt(i.num, 10) + CLCR)
// }

// /* ---- Bulk Str Reply ---- */
// type BulkReply struct {
// 	bulk string
// }

// func MakeBulkReply(content string) *BulkReply {
// 	return &BulkReply{bulk: content}
// }

// func (b *BulkReply) ToBytes() []byte {
// 	len := strconv.FormatInt(int64(len(b.bulk)), 10)
// 	return []byte("$" + len + CLCR + b.bulk + CLCR)
// }

// type NullBulkReply struct {}


