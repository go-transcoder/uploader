package repositories

type UnityOfWork interface {
	StartTransaction()
	GetVideosRepo() VideosRepository
	Rollback()
	Commit()
}
