package main

func deleteAlbum(index int, albums []album) []album {
	return append(albums[:index], albums[index+1:]...)
}
