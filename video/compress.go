package video

type VideoFile struct {
	fileName string
	LongName string
	TypeOf   string
}

func NewVideFile(filePath string) *VideoFile {
	//absolute path
	return &VideoFile{}
}

//文件夹
type Dir struct {
	Folders []*Dir
	Files   []string
}
