package library

import (
	"errors"
	"fmt"
)

type MusicEntry struct {
	Id string
	Name string
	Artist string
	Source string
	Type string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Add(music *MusicEntry) {
	if nil == music {
		return
	}
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) {
	if index < 0 || index >= len(m.musics) {
		fmt.Println("Index out of range")
		return
	}
	if index == 0 {
		m.musics = make([]MusicEntry, 0)
	} else if  index < len(m.musics) - 1{
		m.musics = append(m.musics[:index - 1], m.musics[index + 1:]...)
	} else  {
		m.musics = m.musics[:index - 1]
	}
}

func (m *MusicManager) List() {
	for i, music := range m.musics{
		fmt.Println(i ," : ", music.Name, music.Source, music.Type)
	}
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error){
	if index < 0 || index >= len(m.musics) {
		fmt.Println("Index out of range")
		err = errors.New("Index out of range")
		return nil, err
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager)Find(name string) *MusicEntry {
	for _, music := range m.musics {
		if music.Name == name {
			return &music
		}
	}
	return nil
}



