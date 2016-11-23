package models

type SAddRequest struct {
	Key		interface{}	`json:"key"`
	Value	interface{}	`json:"value"`
}

func SAdd(key,value string)(error){
	sAdd := client.SAdd(key,value)
	err := sAdd.Err()
	if err != nil {
		return err
	}

	return nil
}