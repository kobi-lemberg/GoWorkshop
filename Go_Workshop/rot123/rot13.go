
package rot13



func Rot13(s string ) string {
	buf := make([]byte, len(s))
	for i:=0; i<len(s); i++{
		c := s[i]
		switch  {
		case ('A' <= c && c <= 'M')||('a' <= c && c <= 'm'):
			buf[i] = c+13

		case ('N' <= c && c <= 'Z')||('n' <= c && c <= 'z'):
			buf[i] = c-13
		default:
			buf[i] = c


		}
	}
	return string(buf)
}



