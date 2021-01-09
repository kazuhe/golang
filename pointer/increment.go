package main

func increment(p *int) {
	// パラメータpをデリファレンスして+1増分
	*p++
}

func thatWayIncrement(p int) {
	// ポインタを介さずパラメータpを+1増分
	p++
}
