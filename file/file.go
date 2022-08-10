/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>, December 2018
 */

package file

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func RelativeFilePath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println("error getting relative path")
	}
	return dir, err
}

func GetWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}

func FileExist(f string) bool {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func RenameFile(oldFile string, newFile string) error {
	if FileExist(oldFile) {
		err := os.Rename(oldFile, newFile)
		if err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return errors.New("unable to find file")
	}
}

func DeleteFile(f string) error {
	if FileExist(f) {
		err := os.Remove(f)
		if err != nil {
			log.Println("error deleting file")
			return err
		}
		return nil
	} else {
		return nil
	}
}

func ReadBytes(fp string) ([]byte, error) {
	var dat []byte
	dat, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Println("Unable to read bytes")
		return dat, err
	}
	return dat, nil
}

func ReadFile(fp string) (string, error) {
	dat, err := ReadBytes(fp)
	if err != nil {
		log.Println("Unable to read file")
		return "", err
	}

	return string(dat), nil
}

func AppendFile(content string, fileName string) error {
	//f, err := os.Create(fileName)
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println("Error creating file")
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(content))
	if err != nil {
		log.Println("Error writing file")
		return err
	}

	return nil
}

func WriteFile(content string, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		log.Println("Error creating file")
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(content))
	if err != nil {
		log.Println("Error writing file")
		return err
	}

	return nil
}

func FileSize(fp string) (int64, error) {
	fi, err := os.Stat(fp)
	if err != nil {
		return 0, err
	}

	return fi.Size(), nil
}

func ValidFileExtension(fp string, ext string) bool {
	return filepath.Ext(fp) != ext
}

func TimeFile() string {
	fileTimeFormat := "2006-01-02-15-04-05"
	t := time.Now()
	tf := strings.Replace(t.Format(fileTimeFormat), "+", "-", -1)
	return tf
}

func MakeDirectory(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	return err
}
