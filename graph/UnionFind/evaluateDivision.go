package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	//  [["a","b"],["b","c"]], values = [2.0,3.0],
	p := map[string]string{}
	r := map[string]float64{}
	o := []float64{}

	for i, v := range values {
		// i=0, v=2, e=[a,b]  |  i=1, v=3,  e=[b,c] p = {a:b, b:b} r = {a:2, b:1}
		e := equations[i]
		for _, x := range e {
			// p = {} r= {} ||   p = {a:a, b:b} r = {a:1, b:1}  |  p = {a:b, b:b} r = {a:2, b:1} x =  b || x=c
			if _, ok := p[x]; !ok {
				p[x], r[x] = x, 1.0
			}
		}
		// p = {a:a, b:b} r = {a:1, b:1} |  |  p = {a:b, b:b, c:c} r = {a:2, b:1, c:1}
		union(p, r, e[0], e[1], v)
	}

	//  p = {a:b, b:c, c:c} r = {a:2, b:3, c:1}

	//  [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
	for _, q := range queries {
		// q = [a, c ]
		p0, r0 := find(p, r, q[0])
		p1, r1 := find(p, r, q[1])
		//  p = {a:c, b:c, c:c} r = {a:6, b:3, c:1}
		// c,6
		// c,1
		if len(p1) > 0 && p1 == p0 {
			o = append(o, r0/r1)
		} else {
			o = append(o, -1)
		}
		// 0 = {6/1, 3/6}

	}
	return o
}

func union(p map[string]string, r map[string]float64, i, j string, val float64) {
	// p = {a:a, b:b} r = {a:1, b:1} i=a j=b val =2 | |  p = {a:b, b:b, c:c} r = {a:2, b:1, c:1} i=b j=c val =3
	pi, ri := find(p, r, i)
	pj, rj := find(p, r, j)
	// a, 1, b, 1  | b, 1, c, 1
	if pi != pj {
		p[pi] = pj
		r[pi] = val * rj / ri
	}
	// p = {a:b, b:b} r = {a:2, b:1} |  p = {a:b, b:c, c:c} r = {a:2, b:3, c:1}
}

func find(p map[string]string, r map[string]float64, i string) (string, float64) {
	// p = {a:a, b:b} r = {a:1, b:1} i=a || i = b |  p = {a:b, b:b, c:c} r = {a:2, b:1, c:1} i = b| i=c
	//  p = {a:b, b:c, c:c} r = {a:2, b:3, c:1} i= a || i = b || i = c
	if p[i] != i {
		x, y := find(p, r, p[i])
		p[i] = x
		r[i] *= y
		// r = 3  p[a] = c
		// r[a] = 2*3 = 5
		// p = {a:c, b:c, c:c} r = {a:6, b:3, c:1}
	}
	// a, 1 || b, 1 | b, 1 | c, 1
	return p[i], r[i]
}

/****************************************/

func main() {
	eq := [][]string{{"a", "b"}, {"e", "f"}, {"b", "e"}}
	val := []float64{3.4, 1.4, 2.3}
	que := [][]string{{"b", "a"}, {"a", "f"}, {"f", "f"}, {"e", "e"}, {"c", "c"}, {"a", "c"}, {"f", "e"}}
	//que := [][]string{{"a", "f"}}
	o := calcEquation(eq, val, que)
	out := []float64{0.29412, 10.94800, 1.00000, 1.00000, -1.00000, -1.00000, 0.71429}
	fmt.Println(o)
	fmt.Println(out)

}
