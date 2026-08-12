package tt

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	ttModel "stack-bm/internal/model/mkt/tt"
	ttRepo "stack-bm/internal/repository/mkt/tt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type TitleTemplateService struct {
	repo     *ttRepo.TitleTemplateRepository
	wordRepo *ttRepo.WordListRepository
}

func NewTitleTemplateService() *TitleTemplateService {
	return &TitleTemplateService{repo: ttRepo.NewTitleTemplateRepository(), wordRepo: ttRepo.NewWordListRepository()}
}

var wordPlaceholderRe = regexp.MustCompile(`\{([^\{\}]+)\}`)

func (s *TitleTemplateService) validateMaterials(materials []ttModel.TtTitleMaterial) error {
	for i, m := range materials {
		if strings.TrimSpace(m.Title) == "" {
			return errors.New("标题不能为空")
		}
		matches := wordPlaceholderRe.FindAllStringSubmatch(m.Title, -1)
		if len(matches) > 2 {
			return errors.New("第" + strconv.Itoa(i+1) + "行单个标题最多插入两个词包")
		}
		if len(matches) == 0 {
			continue
		}
		names := make([]string, 0, len(matches))
		for _, mm := range matches {
			names = append(names, mm[1])
		}
		list, err := s.wordRepo.ListByNames(names)
		if err != nil {
			return err
		}
		wordMap := map[string]int{}
		for _, w := range list {
			wordMap[w.Name] = w.CreativeWordID
		}
		wordList := make([]int, 0, len(names))
		for _, n := range names {
			id, ok := wordMap[n]
			if !ok {
				return errors.New("第" + strconv.Itoa(i+1) + "行标题词包配置词包有误，请检查后重新提交")
			}
			wordList = append(wordList, id)
		}
		materials[i].WordList = wordList
	}
	return nil
}

func (s *TitleTemplateService) Create(doc *ttModel.TitleTemplate, adminID int) error {
	if doc.TemplateName == "" {
		return errors.New("模板名称不能为空")
	}
	if err := s.validateMaterials(doc.TitleMaterials); err != nil {
		return err
	}
	doc.TitleNum = len(doc.TitleMaterials)
	doc.SysUserID = adminID
	doc.Display = ttModel.DisplayShow
	doc.CreatedAt = now()
	doc.UpdatedAt = now()
	return s.repo.Create(doc)
}

func (s *TitleTemplateService) FindPage(page, size int, keyword string) ([]ttModel.TitleTemplate, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword)
}

func (s *TitleTemplateService) FindByID(id string) (*ttModel.TitleTemplate, error) {
	return s.repo.FindByID(id)
}

func (s *TitleTemplateService) Update(id string, doc *ttModel.TitleTemplate) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("标题包模板不存在")
	}
	if doc.TemplateName == "" {
		return errors.New("模板名称不能为空")
	}
	if err := s.validateMaterials(doc.TitleMaterials); err != nil {
		return err
	}
	doc.ID = existing.ID
	doc.SysUserID = existing.SysUserID
	doc.Display = existing.Display
	doc.CreatedAt = existing.CreatedAt
	doc.UpdatedAt = now()
	doc.TitleNum = len(doc.TitleMaterials)
	return s.repo.Update(doc)
}

func (s *TitleTemplateService) Delete(id string) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("标题包模板不存在")
	}
	return s.repo.SoftDelete(id)
}

func (s *TitleTemplateService) Copy(id, newName string, adminID int) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("标题包模板不存在")
	}
	if newName == "" {
		return errors.New("模板名称不能为空")
	}
	existing.ID = bson.ObjectID{}
	existing.TemplateName = newName
	existing.SysUserID = adminID
	existing.Display = ttModel.DisplayShow
	existing.CreatedAt = now()
	existing.UpdatedAt = now()
	return s.repo.Create(existing)
}
