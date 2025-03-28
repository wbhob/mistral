package mistral

func Ptr[T any](v T) *T {
	return &v
}

var auto ToolChoice = ToolChoice{
	ToolChoiceEnum: Ptr(TOOLCHOICEENUM_AUTO),
}

type ImageURL = ImageUrl
