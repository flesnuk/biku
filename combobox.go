package main

type TimeAgo struct {
	Id   int
	Name string
}

func HowMuchTimeAgo() []*TimeAgo {
	return []*TimeAgo{
		{-1, "Last 24h"},
		{-7, "Last week"},
		{-30, "Last month"},
		{-90, "Last 3 months"},
		{-180, "Last 6 months"},
	}
}
