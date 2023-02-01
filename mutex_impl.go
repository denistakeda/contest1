package contest

type MyMutex chan struct{}

func New() Mutex {
	m := make(chan struct{}, 1)
	m <- struct{}{}
	return MyMutex(m)
}

func (m MyMutex) Lock() {
	<-m
}

func (m MyMutex) Unlock() {
	m <- struct{}{}
}

func (m MyMutex) LockChannel() <-chan struct{} {
	return m
}
