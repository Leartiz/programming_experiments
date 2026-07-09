package canonical

type Bitset struct {
	a   []byte
	b   []byte
	cnt int
}

func Constructor(size int) Bitset {
	result := Bitset{
		a: make([]byte, size),
		b: make([]byte, size),
	}

	for i := range result.a {
		result.a[i] = '0'
	}
	for i := range result.b {
		result.b[i] = '1'
	}

	return result
}

func (this *Bitset) Fix(idx int) {
	if this.a[idx] == '0' {
		this.a[idx] = '1'
		this.b[idx] = '0'
		this.cnt++
	}
}

func (this *Bitset) Unfix(idx int) {
	if this.a[idx] == '1' {
		this.a[idx] = '0'
		this.b[idx] = '1'
		this.cnt--
	}
}

func (this *Bitset) Flip() {
	this.a, this.b = this.b, this.a
	this.cnt = len(this.a) - this.cnt
}

func (this *Bitset) All() bool {
	return this.cnt == len(this.a)
}

func (this *Bitset) One() bool {
	return this.cnt > 0
}

func (this *Bitset) Count() int {
	return this.cnt
}

func (this *Bitset) ToString() string {
	return string(this.a)
}
