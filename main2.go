package main2
	likeRepo := _driverFactory.NewLikeRepository(db)
	likeUsecase := _likeUseCase.NewLikeUsecase(likeRepo)
	likeCtrl := _likeController.NewLikeController(likeUsecase)

	commentRepo := _driverFactory.NewCommentRepository(db)
	commentUsecase := _commentUseCase.NewCommentUsecase(commentRepo)
	commentCtrl := _commentController.NewCommentController(commentUsecase)

	threadRepo := _driverFactory.NewThreadRepository(db)
	threadUsecase := _threadUseCase.NewThreadUsecase(threadRepo)
	threadCtrl := _threadController.NewThreadController(threadUsecase)

	savethreadRepo := _driverFactory.NewSavethreadRepository(db)
	savethreadUsecase := _savethreadUseCase.NewSavethreadUsecase(savethreadRepo)
	savethreadCtrl := _savethreadController.NewSavethreadController(savethreadUsecase)

	followRepo := _driverFactory.NewFollowRepository(db)
	followUsecase := _followUseCase.NewFollowUsecase(followRepo)
	followCtrl := _followController.NewFollowController(followUsecase)