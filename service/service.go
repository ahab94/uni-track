package service

import UniTrack "github.com/ahab94/uni-track"

type Service struct {
	rt *UniTrack.Runtime
}

// NewUniTrackService exports service struct
func NewUniTrackService(rt *UniTrack.Runtime) *Service {
	return &Service{rt: rt}
}
