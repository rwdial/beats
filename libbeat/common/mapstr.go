package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	EventMetadataKey = "_event_metadata"
	FieldsKey        = "fields"
	TagsKey          = "tags"
)

var ErrorFieldsIsNotMapStr = errors.New("the value stored in fields is not a MapStr")
var ErrorTagsIsNotStringArray = errors.New("the value stored in tags is not a []string")

// Commonly used map of things, used in JSON creation and the like.
type MapStr map[string]interface{}

// EventMetadata contains fields and tags that can be added to an event via
// configuration.
type EventMetadata struct {
	Fields          MapStr
	FieldsUnderRoot bool `config:"fields_under_root"`
	Tags            []string
}

// Eventer defines a type its ability to fill a MapStr.
type Eventer interface {
	// Add fields to MapStr.
	Event(event MapStr) error
}

// MapStrUnion creates a new MapStr containing the union of the
// key-value pairs of the two maps. If the same key is present in
// both, the key-value pairs from dict2 overwrite the ones from dict1.
func MapStrUnion(dict1 MapStr, dict2 MapStr) MapStr {
	dict := MapStr{}

	for k, v := range dict1 {
		dict[k] = v
	}

	for k, v := range dict2 {
		dict[k] = v
	}
	return dict
}

// Update copies all the key-value pairs from the
// d map overwriting any existing keys.
func (m MapStr) Update(d MapStr) {
	for k, v := range d {
		m[k] = v
	}
}

// Checks if a timestamp field exists and if it doesn't it adds
// one by using the injected now() function as a time source.
func (m MapStr) EnsureTimestampField(now func() time.Time) error {
	ts, exists := m["@timestamp"]
	if !exists {
		m["@timestamp"] = Time(now())
		return nil
	}

	_, is_common_time := ts.(Time)
	if is_common_time {
		// already perfect
		return nil
	}

	tstime, is_time := ts.(time.Time)
	if is_time {
		m["@timestamp"] = Time(tstime)
		return nil
	}

	tsstr, is_string := ts.(string)
	if is_string {
		var err error
		m["@timestamp"], err = ParseTime(tsstr)
		return err
	}
	return fmt.Errorf("Don't know how to convert %v to a Time value", ts)
}

// EnsureCountField sets the 'count' field to 1 if count does not already exist.
func (m MapStr) EnsureCountField() error {
	_, exists := m["count"]
	if !exists {
		m["count"] = 1
	}
	return nil
}

// String returns the MapStr as a JSON string.
func (m MapStr) String() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		return fmt.Sprintf("Not valid json: %v", err)
	}
	return string(bytes)
}

// MergeFields merges the top-level keys and values in each source hash (it does
// not perform a deep merge). If the same key exists in both, the value in
// fields takes precedence. If underRoot is true then the contents of the fields
// MapStr is merged with the value of the 'fields' key in ms.
//
// An error is returned if underRoot is true and the value of ms.fields is not a
// MapStr.
func MergeFields(ms, fields MapStr, underRoot bool) error {
	if ms == nil || fields == nil {
		return nil
	}

	fieldsMS := ms
	if !underRoot {
		f, ok := ms[FieldsKey]
		if !ok {
			fieldsMS = make(MapStr, len(fields))
			ms[FieldsKey] = fieldsMS
		} else {
			// Use existing 'fields' value.
			fieldsMS, ok = f.(MapStr)
			if !ok {
				return ErrorFieldsIsNotMapStr
			}
		}
	}

	// Add fields and override.
	for k, v := range fields {
		fieldsMS[k] = v
	}

	return nil
}

// AddTag appends a tag to the tags field of ms. If the tags field does not
// exist then it will be created. If the tags field exists and is not a []string
// then an error will be returned. It does not deduplicate the list of tags.
func AddTags(ms MapStr, tags []string) error {
	if ms == nil || len(tags) == 0 {
		return nil
	}

	tagsIfc, ok := ms[TagsKey]
	if !ok {
		ms[TagsKey] = tags
		return nil
	}

	existingTags, ok := tagsIfc.([]string)
	if !ok {
		return ErrorTagsIsNotStringArray
	}

	ms[TagsKey] = append(existingTags, tags...)
	return nil
}
