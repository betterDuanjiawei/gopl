package main

import (
	"fmt"
	"go4.org/sort"
	"os"
	"text/tabwriter"
	"time"
)

func main()  {
	sort.Sort(byArtist(track))
	printTrack(track)
	sort.Sort(sort.Reverse(byArtist(track))) // 这个咋用啊?
	printTrack(track)

	sort.Sort(ByYear(track))
	printTrack(track)
	
	sort.Sort(customSort{track, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}

		return false
	}})
	printTrack(track)

	demo1()
}

func demo1()  {
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values))
	sort.Ints(values)
	fmt.Println(sort.IntsAreSorted(values))

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
	//false
	//true
	//[4 3 1 1]
	//false

}

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int)  {
	x[i], x[j] = x[j], x[i]
}

type ByYear []*Track
func (x ByYear) Len() int {
	return len(x)
}

func (x ByYear) Less(i, j int) bool {
	return x[i].Year < x[j].Year
}

func (x ByYear) Swap(i, j int)  {
	x[i], x[j] = x[j], x[i]
}

var track = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solving", "Smash", 2011, length("4m24s")},

}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTrack(tracks []*Track)  {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}

	tw.Flush() // 计算各列宽度并输出表格

}

type customSort	struct {
	t []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int  {
	return len(x.t)
}
func (x customSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}
func (x customSort) Swap(i, j int)  {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}
