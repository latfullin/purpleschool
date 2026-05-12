package helpers

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type VerifyRecord struct {
	Hash      string    `json:"hash"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func SaveHash(hash string, email string) error {
	records, err := readAllHashes()
	if err != nil {
		return err
	}

	record := VerifyRecord{
		Hash:      hash,
		Email:     email,
		CreatedAt: time.Now(),
	}

	records = append(records, record)
	return writeAllHashes(records)
}

func ReadHash(hash string) (VerifyRecord, error) {
	records, err := readAllHashes()
	if err != nil {
		return VerifyRecord{}, err
	}

	for _, record := range records {
		if record.Hash == hash {
			return record, nil
		}
	}

	return VerifyRecord{}, os.ErrNotExist
}

func DelereHash(hash string) error {
	records, err := readAllHashes()
	if err != nil {
		return err
	}

	filtered := make([]VerifyRecord, 0, len(records))
	found := false
	for _, record := range records {
		if record.Hash == hash {
			found = true
			continue
		}
		filtered = append(filtered, record)
	}

	if !found {
		return os.ErrNotExist
	}

	return writeAllHashes(filtered)
}

func readAllHashes() ([]VerifyRecord, error) {
	content, err := os.ReadFile("verify.json")
	if errors.Is(err, os.ErrNotExist) {
		return []VerifyRecord{}, nil
	}
	if err != nil {
		return nil, err
	}
	if len(content) == 0 {
		return []VerifyRecord{}, nil
	}

	var records []VerifyRecord
	if err := json.Unmarshal(content, &records); err == nil {
		return records, nil
	}

	var single VerifyRecord
	if err := json.Unmarshal(content, &single); err == nil {
		if single.Hash == "" {
			return []VerifyRecord{}, nil
		}
		return []VerifyRecord{single}, nil
	}

	return nil, err
}

func writeAllHashes(records []VerifyRecord) error {
	content, err := json.Marshal(records)
	if err != nil {
		return err
	}

	return os.WriteFile("verify.json", content, 0o600)
}
