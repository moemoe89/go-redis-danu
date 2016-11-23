package models

func Delete(key string)(error){

	delete := client.Del(key)
	err := delete.Err()
	if err != nil {
		return err
	}

	return nil
}