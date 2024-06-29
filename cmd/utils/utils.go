package utils

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

type ProjectFile struct {
	Name     string
	Path     string
	TmplName string
	Data     any
}

func CreateProjectFile(projectPath *string, pkgName string, p ProjectFile) error {
	file, err := os.Create(filepath.Join(*projectPath, p.Path, p.Name))
	if err != nil {
		return err
	}
	defer file.Close()

	if p.TmplName != "" {
		//inject content into the file
		if err := InjectContentToFile(file, pkgName, p.TmplName, p.Data); err != nil {
			return err
		}
	}
	return nil
}

func InjectContentToFile(writer io.Writer, pkgName string, tmplName string, data any) error {
	tmplFile := filepath.Join("cmd/", pkgName, "files", tmplName)
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		return err
	}
	if err := tmpl.Execute(writer, data); err != nil {
		return err
	}
	return nil
}

func CreateAllProjectFiles(projectFiles []ProjectFile, projectPath *string, pkgName string) error {
	for _, projectFile := range projectFiles {
		if err := CreateProjectFile(projectPath, pkgName, projectFile); err != nil {
			return err
		}
	}
	return nil
}

func ToTitleCase(input string) string {
	var words []string
	var currentWord []rune

	for i, r := range input {
		if unicode.IsUpper(r) && i > 0 && (unicode.IsLower(rune(input[i-1])) || unicode.IsDigit(rune(input[i-1]))) {
			words = append(words, string(currentWord))
			currentWord = nil
		} else if r == '_' || r == '-' || r == ' ' {
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = nil
			}
			continue
		}
		currentWord = append(currentWord, unicode.ToLower(r))
	}
	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, " ")
}
