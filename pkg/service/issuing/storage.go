package issuing

import (
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"github.com/tavroi/ssi-service/pkg/storage"
)

type Storage struct {
	db storage.ServiceStorage
}

const namespace = "issuance_template"

func NewIssuingStorage(s storage.ServiceStorage) (*Storage, error) {
	if s == nil {
		return nil, errors.New("s cannot be nil")
	}
	return &Storage{
		db: s,
	}, nil
}

type StoredIssuanceTemplate struct {
	IssuanceTemplate IssuanceTemplate
}

func (s Storage) StoreIssuanceTemplate(template StoredIssuanceTemplate) error {
	if template.IssuanceTemplate.ID == "" {
		return errors.New("cannot store issuance template without an ID")
	}
	data, err := json.Marshal(template)
	if err != nil {
		return errors.Wrap(err, "marshalling template")
	}
	return s.db.Write(namespace, template.IssuanceTemplate.ID, data)
}

func (s Storage) GetIssuanceTemplate(id string) (*StoredIssuanceTemplate, error) {
	if id == "" {
		return nil, errors.New("cannot fetch issuance template without an ID")
	}
	data, err := s.db.Read(namespace, id)
	if err != nil {
		return nil, errors.Wrap(err, "reading from db")
	}
	if len(data) == 0 {
		return nil, errors.Errorf("issuance template not found with id: %s", id)
	}
	var st StoredIssuanceTemplate
	if err = json.Unmarshal(data, &st); err != nil {
		return nil, errors.Wrap(err, "unmarshalling template")
	}
	return &st, nil
}

func (s Storage) DeleteIssuanceTemplate(id string) error {
	if id == "" {
		return nil
	}
	if err := s.db.Delete(namespace, id); err != nil {
		return errors.Wrap(err, "deleting from db")
	}
	return nil
}

func (s Storage) ListIssuanceTemplates() ([]IssuanceTemplate, error) {
	m, err := s.db.ReadAll(namespace)
	if err != nil {
		return nil, errors.Wrap(err, "reading all")
	}
	ts := make([]IssuanceTemplate, len(m))
	i := 0
	for k, v := range m {
		if err = json.Unmarshal(v, &ts[i]); err != nil {
			return nil, errors.Wrapf(err, "unmarshalling template with key <%s>", k)
		}
		i++
	}
	return ts, nil
}
