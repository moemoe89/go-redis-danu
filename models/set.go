package models

type SetRequest struct {
	Key		interface{}	`json:"key"`
	Value	interface{}	`json:"value"`
}

func Set(key,value string)(error){
	set := client.Set(key, value, 0)
	err := set.Err()
	if err != nil {
		return err
	}

	return nil
}