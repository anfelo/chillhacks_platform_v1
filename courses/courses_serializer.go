package courses

import "encoding/json"

type CourseSerializer struct{}
type CoursesSerializer struct{}

func (i *CourseSerializer) Decode(input []byte) (*Course, error) {
	item := &Course{}
	if err := json.Unmarshal(input, item); err != nil {
		return nil, err
	}
	return item, nil
}

func (i *CourseSerializer) Encode(input *Course) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	return rawMsg, nil
}
