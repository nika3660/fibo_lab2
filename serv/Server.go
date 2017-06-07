package serv

//import "fmt"

type Server struct {
	Cache Cache
}

func InitServer(cacheSize int) *Server {
	return &Server{
		Cache{
			Size: cacheSize,
			History: make(map[int]int),
			CacheContainer: make(map[int]int),
		},
	}
}

func (server Server) GetCache() int {
	return server.Cache.Size
}

func (server Server) GetFibonacciNumber(number int) int {
	server.Cache.UpdateHistory(number)
	check := server.Cache.CheckValueInCache(number)
	if check {
		return server.Cache.GetFibonValue(number)
	} else {
		value := Fibon(number)
		server.Cache.AddNewValue(number, value)
		return value
	}
}