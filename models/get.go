package models

func Get(key string)(string,error){

	get := client.Get(key)
	err := get.Err()
	if err != nil {
		return "",err
	}

	result,err := get.Result()
	if err != nil {
		return "",err
	}

	return result,nil
}

func GetSadd(key string)([]string,error){

	get := client.SMembers(key)
	err := get.Err()
	if err != nil {
		return nil,err
	}

	result,err := get.Result()
	if err != nil {
		return nil,err
	}

	return result,nil
}

func GetHSet(key,field string)(string,error){

	getHSet := client.HGet(key,field)
	err := getHSet.Err()
	if err != nil {
		return "",err
	}

	result,err := getHSet.Result()
	if err != nil {
		return "",err
	}

	return result,nil
}