package _0211018

import "strings"

//type lazybuf struct {
//	s   string
//	buf []byte
//	w   int
//}
//
//func (b *lazybuf) append(c byte) {
//	if b.buf == nil {
//		if b.w < len(b.s) && b.s[b.w] == c {
//			b.w++
//			return
//		}
//		b.buf = make([]byte, len(b.s))
//		copy(b.buf, b.s[:b.w])
//	}
//	b.buf[b.w] = c
//	b.w++
//}
//
//func (b *lazybuf) index(i int) byte {
//	if b.buf != nil {
//		return b.buf[i]
//	}
//	return b.s[i]
//}
//
//func (b *lazybuf) string() string {
//	if b.buf == nil {
//		return b.s[:b.w]
//	}
//	return string(b.buf[:b.w])
//}
//
//func clean(path string) string {
//	if path == "" {
//		return "."
//	}
//
//	rooted := path[0] == '/'
//	n := len(path)
//
//	out := lazybuf{s: path}
//	r, dotdot := 0, 0
//	if rooted {
//		out.append('/')
//		r, dotdot = 1, 1
//	}
//
//	for r < n {
//		switch {
//		case path[r] == '/':
//			// empty path element
//			r++
//		case path[r] == '.' && (r+1 == n || path[r+1] == '/'):
//			// . element
//			r++
//		case path[r] == '.' && path[r+1] == '.' && (r+2 == n || path[r+2] == '/'):
//			// .. element: remove to last /
//			r += 2
//			switch {
//			case out.w > dotdot:
//				// can backtrack
//				out.w--
//				for out.w > dotdot && out.index(out.w) != '/' {
//					out.w--
//				}
//			case !rooted:
//				// cannot backtrack, but not rooted, so append .. element.
//				if out.w > 0 {
//					out.append('/')
//				}
//				out.append('.')
//				out.append('.')
//				dotdot = out.w
//			}
//		default:
//			// real path element.
//			// add slash if needed
//			if rooted && out.w != 1 || !rooted && out.w != 0 {
//				out.append('/')
//			}
//			// copy element
//			for ; r < n && path[r] != '/'; r++ {
//				out.append(path[r])
//			}
//		}
//	}
//
//	// Turn empty string into "."
//	if out.w == 0 {
//		return "."
//	}
//
//	return out.string()
//}
//func simplifyPath(path string) string {
//	return clean(path)
//}

//func simplifyPath(path string) string {
//	paths := strings.Split(path, "/") //以/作为分隔符分割字符串
//	stack := []string{}               //模拟栈，栈里面放每一段字符串
//
//	for _, v := range paths {
//		switch v {
//		case "..":
//			if len(stack) >= 1 {
//				stack = stack[:len(stack)-1]
//			}
//		case "", ".":
//			continue
//		default:
//			stack = append(stack, v)
//		}
//	}
//	return "/" + strings.Join(stack, "/")
//}

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	stack := []string{}
	for _, v := range paths {
		switch v {
		case "..":
			if len(stack) >= 1 {
				stack = stack[:len(stack)-1]
			}
		case "", ".":
			continue
		default:
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}
