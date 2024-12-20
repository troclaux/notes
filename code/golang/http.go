package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id
	jsonData, err := json.Marshal(data)

	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)

	if err != nil {
		return User{}, err
	}
	return user, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)

	if err != nil {
		return User{}, err
	}
	return user, nil

}

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return errors.New("request to delete user unsuccessful")
	}
	return nil
}
