package persistence

import "fmt"

var daoProvidersMap_Language = make(map[int]func() LanguageDAO)

func RegisterLanguageDAO(daoType int, provider func() LanguageDAO) {
	daoProvidersMap_Language[daoType] = provider
}

func GetLanguageDAO() LanguageDAO {
	return GetSpecificLanguageDAO(defaultDAOType)
}

func GetSpecificLanguageDAO(daoType int) LanguageDAO {
	provider := daoProvidersMap_Language[daoType]
	if provider != nil {
		return provider() // call provider function
	} else {
		panic(fmt.Sprintf("Language DAO not found ( DAO type = %d )", daoType))
	}
}
