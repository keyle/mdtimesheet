package main

import "testing"

func TestLineParsingFull(t *testing.T) {
	if parseLine("2020-01-06-wed :1200-1220-10m") != 10 {
		t.Fail()
	}
}

func TestLineParsingPartial(t *testing.T) {
	if parseLine("2020-01-06-wed :1200-1220-30m") != 0 {
		t.Fail()
	}
}

func TestLineParsingDay(t *testing.T) {
	if parseLine("2020-01-06-wed :0600-2220-30m") != 950 {
		t.Fail()
	}
}

func TestLineBadDay(t *testing.T) {
	if parseLine("2020-01-66-wed :0600-2220-30m") != 0 {
		t.Fail()
	}
}

func TestLineBadMonth(t *testing.T) {
	if parseLine("2020-50-02-wed :0600-2220-30m") != 0 {
		t.Fail()
	}
}

func TestLineBadYYYY(t *testing.T) {
	if parseLine("201-50-02-wed :0600-2220-30m") != 0 {
		t.Fail()
	}
}

func TestLineBadDaaay(t *testing.T) {
	if parseLine("2010-12-02-wedasdasdasd :0600-1220-30m") != 0 {
		t.Fail()
	}
}

func TestLineBadStart(t *testing.T) {
	if parseLine("2010-12-02-mon :04600-1220-30m") != 0 {
		t.Fail()
	}
}

func TestLineBadEnd(t *testing.T) {
	if parseLine("2010-12-02-mon :0400-12x20-30m") != 0 {
		t.Fail()
	}
}

func TestLineParsingNext(t *testing.T) {
	if parseLine("@next") != 0 {
		t.Fail()
	}
}

func TestLineParsingConsider(t *testing.T) {
	if parseLine("@consider") != 0 {
		t.Fail()
	}
}

func TestLineParsingBad(t *testing.T) {
	if parseLine("poiasdjaoi dasid apioasj") != 0 {
		t.Fail()
	}
}
