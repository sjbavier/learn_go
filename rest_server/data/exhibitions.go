package data

type Exhibition struct {
	Title       string
	Description string
	Image       string
}

func GetAll() []Exhibition {
	return list
}

func Add(exhibition Exhibition) {
	list = append(list, exhibition)
}

var list = []Exhibition{
	{
		Title:       "title for exhibition",
		Description: "description of exhibition",
		Image:       "image-that-exists.jpg",
	},
	{
		Title:       "title 2",
		Description: "description 2",
		Image:       "image-2.jpg",
	},
}
