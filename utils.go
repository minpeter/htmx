package main

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"
)

func RenderTemplates(mainTemplateName, pagesDirPath string) ([]byte, error) {
	// 메인 템플릿 디렉토리
	mainTemplateDir := "templates/layouts/"

	// 템플릿 생성
	tmpl, err := template.New(mainTemplateName).ParseGlob(filepath.Join(mainTemplateDir, "*.tmpl"))
	if err != nil {
		return nil, err
	}

	fmt.Println("1")

	// 서브 템플릿 등록
	subTemplatePath := filepath.Join("templates/pages/", pagesDirPath+".tmpl")
	_, err = tmpl.ParseFiles(subTemplatePath)
	if err != nil {
		return nil, err
	}

	fmt.Println("2")

	// 렌더링 결과를 저장할 버퍼 생성
	var result bytes.Buffer

	// 템플릿 실행 및 결과를 버퍼에 쓰기
	err = tmpl.ExecuteTemplate(&result, mainTemplateName+".tmpl", nil)
	if err != nil {
		return nil, err
	}

	// 렌더링된 결과를 []byte로 반환
	return result.Bytes(), nil
}
