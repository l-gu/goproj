package persistence

import "fmt"

var daoProvidersMap_FooBar = make(map[int]func() FooBarDAO)

func RegisterFooBarDAO(daoType int, provider func() FooBarDAO) {
	daoProvidersMap_FooBar[daoType] = provider
}

func GetFooBarDAO() FooBarDAO {
	return GetSpecificFooBarDAO(defaultDAOType)
}

func GetSpecificFooBarDAO(daoType int) FooBarDAO {
	provider := daoProvidersMap_FooBar[daoType]
	if provider != nil {
		return provider() // call provider function
	} else {
		panic(fmt.Sprintf("FooBar DAO not found ( DAO type = %d )", daoType))
	}
}
