package pipelines

type FSFilesStore struct {
}

type S3FilesStore struct {
}

type GCSFilesStore struct {
}

type FilesPipeline struct {
	*MediaPipeline
}


func (mp *FilesPipeline) MediaToDownload()  {

}

func (mp *FilesPipeline) GetMediaRequests()  {

}

func (mp *FilesPipeline) MediaDownloaded()  {

}

func (mp *FilesPipeline) MediaFailed()  {

}

func (mp *FilesPipeline) ItemCompleted()  {

}

func (mp *FilesPipeline) FileDownloaded()  {

}

func (mp *FilesPipeline) FilePath()  {

}

