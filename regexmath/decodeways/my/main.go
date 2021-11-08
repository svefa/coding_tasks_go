package main

func numDecodings(s string) int {
	// 226 10
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	a, b := 1, 1
	r2 := (s[0]-'0')*10 + (s[1] - '0')

	/*  if  r2 == 10 || r2 == 20 {
	        a, b = 1,1 //a, a}
	    } else if  r2 <= 26 {
	        a, b = 1, 2  //  b, a + b}
	    } else if  r2 > 27 {
	        a, b = 1, 1 // a, a}
	    }
	*/
	for i := 1; i < len(s); i++ {
		r2 = (s[i-1]-'0')*10 + (s[i] - '0')

		// ...ab
		// ....06
		// ....ab
		if r2 == 0 || r2 == 30 || r2 == 40 || r2 == 50 || r2 == 60 || r2 == 70 || r2 == 80 || r2 == 90 {
			return 0
		} else if r2 < 10 {
			a, b = b, b
		} else if r2 == 10 || r2 == 20 {
			a, b = a, a
		} else if r2 <= 26 {
			a, b = b, a+b
		} else if r2 > 27 {
			a, b = b, b
		}

	}
	return b

}
