package solutions

import "github.com/kyokomi/emoji"

func GetMessage() string {
	s := emoji.Sprint("Hello :world_map:!")
	return s
}
