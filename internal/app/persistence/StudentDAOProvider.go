package persistence

import "fmt"

var daoProvidersMap_Student = make(map[int]func() StudentDAO)

func RegisterStudentDAO(daoType int, provider func() StudentDAO) {
	daoProvidersMap_Student[daoType] = provider
}

func GetStudentDAO() StudentDAO {
	return GetSpecificStudentDAO(defaultDAOType)
}

func GetSpecificStudentDAO(daoType int) StudentDAO {
	provider := daoProvidersMap_Student[daoType]
	if provider != nil {
		return provider() // call provider function
	} else {
		panic(fmt.Sprintf("Student DAO not found ( DAO type = %d )", daoType))
	}
}
