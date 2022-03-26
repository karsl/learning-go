package stub

type User struct {
}

type Pet struct {
	Name string
}

type Person struct {
}

type Entities interface {
	GetUser(id string) (User, error)
	GetPets(userID string) ([]Pet, error)
	GetChildren(userID string) ([]Person, error)
	GetFriends(userID string) ([]Person, error)
	SaveUser(user User) error
}

type Logic struct {
	Entities Entities
}

func (l Logic) GetPetNames(userId string) ([]string, error) {
	pets, err := l.Entities.GetPets(userId)
	if err != nil {
		return nil, err
	}
	out := make([]string, len(pets))
	for i, p := range pets {
		out[i] = p.Name
		//out = append(out, p.Name)
	}
	return out, nil
}

type GetPetNamesStub struct {
	Entities
}

//func (ps GetPetNamesStub) GetPets(userID string) ([]Pet, error) {
//	switch userID {
//	case "1":
//		return []Pet{{Name: "Bubbles"}}, nil
//	case "2":
//		return []Pet{{Name: "Stampy"}, {Name: "Snowball II"}}, nil
//	default:
//		return nil, fmt.Errorf("invalid id: %s", userID)
//	}
//}

func (ps GetPetNamesStub) GetUser(id string) (User, error) {
	return User{}, nil
}

func (ps GetPetNamesStub) GetChildren(userId string) ([]Person, error) {
	return nil, nil
}

func (ps GetPetNamesStub) GetFriends(userId string) ([]Person, error) {
	return nil, nil
}

func (ps GetPetNamesStub) SaveUser(user User) error {
	return nil
}
