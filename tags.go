package teamwork

import (
	"encoding/json"
	"fmt"
)

type ModTag struct {
	Content string `json:"content"`
}

func (conn *Connection) AddTag(taskId int, tag string) error {
	method := "PUT"
	url := fmt.Sprintf("%stasks/%d/tags.json", conn.Account.Url, taskId)
	fmt.Printf("adding tag %s to URL: %s\n", tag, url)
	m := &ModTag{Content: tag}
	mp := make(map[string]*ModTag)
	mp["tags"] = m
	data_s, _ := json.Marshal(mp)
	//fmt.Printf("JSON:\n%s\n", data_s)
	data := []byte(data_s)
	_, _, err := postrequest(conn.ApiToken, method, url, data)
	if err != nil {
		return fmt.Errorf("Cannot set tag %s on %d", tag, taskId)
	}
	return nil
}

type omg struct {
	Tags               map[string]string `json:"tags"`
	RemoveProvidedTags string            `json:"removeProvidedTags"`
}

func (conn *Connection) RemoveTag(taskId int, tag string) error {
	method := "PUT"
	url := fmt.Sprintf("%stasks/%d/tags.json", conn.Account.Url, taskId)
	fmt.Printf("removing tag %s to URL: %s\n", tag, url)
	o := &omg{Tags: make(map[string]string)}
	o.Tags["content"] = tag
	o.RemoveProvidedTags = "true"
	data_s, _ := json.Marshal(o)
	//fmt.Printf("JSON:\n%s\n", data_s)
	data := []byte(data_s)
	_, _, err := postrequest(conn.ApiToken, method, url, data)
	if err != nil {
		return fmt.Errorf("Cannot set tag %s on %d", tag, taskId)
	}
	return nil
}
