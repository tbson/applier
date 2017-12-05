package constant


const PageSize int = 2
const HeaderLinkRaw string = "<%s?start=%d&direction=%s>; rel=\"%s\""

type POption struct {
	Start int
	Direction string
}
