package simpleflask

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func CreateProject(pathToCreate string) error {
	projectPath, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := createProjectFolders(projectPath, pathToCreate); err != nil {
		return err
	}
	if err := createProjectFiles(projectPath, pathToCreate); err != nil {
		return err
	}
	return nil
}

func createProjectFolders(projectPath string, pathToCreate string) error {
	path := filepath.Join(projectPath, pathToCreate)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0751)
		if err != nil {
			log.Printf("Error creating directory %v\n", err)
			return err
		}

		//create app folder
		appPath := filepath.Join(path, "app")

		if err := os.Mkdir(appPath, 0754); err != nil {
			return err
		}

		staticPath := filepath.Join(appPath, "static")
		if err := os.Mkdir(staticPath, 0754); err != nil {
			return err
		}
		for _, folder := range []string{"css", "js", "imgs"} {
			if err := os.Mkdir(filepath.Join(staticPath, folder), 0754); err != nil {
				return err
			}
		}

		templatesPath := filepath.Join(appPath, "templates")
		if err := os.Mkdir(templatesPath, 0754); err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("Project with the same name already exists")
}
func createProjectFiles(projectPath string, pathToCreate string) error {
	path := filepath.Join(projectPath, pathToCreate)

	//create main.py
	mainFile, err := os.Create(filepath.Join(path, "main.py"))
	if err != nil {
		return err
	}
	defer mainFile.Close()

	//inject data into main.py
	if err := injectDataToFile(mainFile, "main.tmpl", nil); err != nil {
		return err
	}

	//create __init__.py
	initFile, err := os.Create(filepath.Join(path, "app", "__init__.py"))
	if err != nil {
		return err
	}

	//inject data into __init__.py
	if err := injectDataToFile(initFile, "init.tmpl", nil); err != nil {
		return err
	}

	//create config.py
	configFile, err := os.Create(filepath.Join(path, "app", "config.py"))
	if err != nil {
		return err
	}

	//inject data into config.py
	if err := injectDataToFile(configFile, "config.tmpl", nil); err != nil {
		return err
	}

	//create views.py
	viewsFile, err := os.Create(filepath.Join(path, "app", "views.py"))
	if err != nil {
		return err
	}

	//inject data into views.py
	if err := injectDataToFile(viewsFile, "views.tmpl", nil); err != nil {
		return err
	}

	//create models.py
	modelsFile, err := os.Create(filepath.Join(path, "app", "models.py"))
	if err != nil {
		return err
	}

	//inject data into models.py
	if err := injectDataToFile(modelsFile, "models.tmpl", nil); err != nil {
		return err
	}

	//create README.md
	readmeFile, err := os.Create(filepath.Join(path, "README.md"))
	if err != nil {
		return err
	}

	//inject data into README.md
	if err := injectDataToFile(readmeFile, "readme.tmpl", nil); err != nil {
		return err
	}

	//create requirements.txt
	requirementsFile, err := os.Create(filepath.Join(path, "requirements.txt"))
	if err != nil {
		return err
	}

	//inject data into requirements.txt
	if err := injectDataToFile(requirementsFile, "requirements.tmpl", nil); err != nil {
		return err
	}

	//create base.html
	baseFile, err := os.Create(filepath.Join(path, "app", "templates/base.html"))
	if err != nil {
		return err
	}

	//inject data into base.html
	if err := injectDataToFile(baseFile, "base_html.tmpl", nil); err != nil {
		return err
	}

	//create index.html
	indexFile, err := os.Create(filepath.Join(path, "app", "templates/index.html"))
	if err != nil {
		return err
	}

	//inject data into index.html
	if err := injectDataToFile(indexFile, "index_html.tmpl", nil); err != nil {
		return err
	}

	//create forms.py
	if _, err := os.Create(filepath.Join(path, "app", "forms.py")); err != nil {
		return err
	}

	//create .env
	if _, err := os.Create(filepath.Join(path, ".env")); err != nil {
		return err
	}

	//create styles.css
	if _, err := os.Create(filepath.Join(path, "app/static/css/styles.css")); err != nil {
		return err
	}

	//create index.js
	if _, err := os.Create(filepath.Join(path, "app/static/js/index.js")); err != nil {
		return err
	}

	return nil
}
func injectDataToFile(writer io.Writer, tmplName string, data any) error {
	tmplFile := "cmd/simpleflask/files/" + tmplName
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		return err
	}
	if err := tmpl.Execute(writer, data); err != nil {
		return err
	}
	return nil
}

// func createDir(path string,perm int){
//   if err := os.Mkdir(path, perm);err != nil{
//       return err
//     }
//     return nil
// }
