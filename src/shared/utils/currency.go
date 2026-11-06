package veUtils

type Currency uint32

func (c Currency) Int() int {
	return int(c)
}
