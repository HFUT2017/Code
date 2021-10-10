package _0211010

import "math"

type StockPrice struct {
	Hash      map[int]int
	Max       int
	Min       int
	LastTime  int
	LastPrice int
}

func Constructor() StockPrice {
	return StockPrice{
		Hash:      map[int]int{},
		Min:       math.MaxInt32,
		Max:       math.MinInt32,
		LastTime:  math.MinInt32,
		LastPrice: 0,
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if value, ok := this.Hash[timestamp]; ok {
		if value == this.Max {
			this.Max = math.MinInt32
		}
		if value == this.Min {
			this.Min = math.MaxInt32
		}
	}
	this.Hash[timestamp] = price
	if this.Max == math.MinInt32 || this.Min == math.MaxInt32 {
		for _, value := range this.Hash {
			if value > this.Max {
				this.Max = value
			}
			if value < this.Min {
				this.Min = value
			}
		}
	}
	if price > this.Max {
		this.Max = price
	}
	if price < this.Min {
		this.Min = price
	}
	if timestamp >= this.LastTime {
		this.LastTime = timestamp
		this.LastPrice = price
	}
}

func (this *StockPrice) Current() int {
	return this.LastPrice
}

func (this *StockPrice) Maximum() int {
	return this.Max
}

func (this *StockPrice) Minimum() int {
	return this.Min
}
