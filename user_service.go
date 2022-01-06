package typeform

type UserService struct {
	resource
}

func NewUserService(client Client) *UserService {
	return &UserService{
		resource: resource{
			client: client,
		},
	}
}

func (s UserService) RetrieveSelf() (User, error) {
	var created User
	err := s.resource.retrieve(resourceUserSelf, "", &created)
	if err != nil {
		return User{}, err
	}

	return created, err
}
