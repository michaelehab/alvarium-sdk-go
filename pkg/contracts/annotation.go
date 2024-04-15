/*******************************************************************************
 * Copyright 2021 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
package contracts

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/oklog/ulid/v2"
)

// TagEnvKey is an environment key used to associate annotations with specific metadata,
// aiding in the linkage of scores across different layers of the stack. For instance, in the "app" layer,
// it is utilized to retrieve the commit SHA of the workload where the application is running,
// which is instrumental in tracing the impact on the current layer's score from the lower layers.
const TagEnvKey = "TAG"

// Annotation represents an individual criterion of evaluation in regard to a piece of data
type Annotation struct {
	Id          ulid.ULID      `json:"id,omitempty"`        // Id should probably be a ULID -- uniquely identifies the annotation itself
	Key         string         `json:"key,omitempty"`       // Key is the hash value of the data being annotated
	Hash        HashType       `json:"hash,omitempty"`      // Hash identifies which algorithm was used to construct the hash
	Host        string         `json:"host,omitempty"`      // Host is the hostname of the node making the annotation
	Tag         string         `json:"tag,omitempty"`       // Tag is the link between the current layer and the below layer
	Layer       LayerType      `json:"layer,omitempty"`     // Layer is the layer where the annotation was produced
	Kind        AnnotationType `json:"kind,omitempty"`      // Kind indicates what kind of annotation this is
	Signature   string         `json:"signature,omitempty"` // Signature contains the signature of the party making the annotation
	IsSatisfied bool           `json:"isSatisfied"`         // IsSatisfied indicates whether the criteria defining the annotation were fulfilled
	Timestamp   time.Time      `json:"timestamp,omitempty"` // Timestamp indicates when the annotation was created
}

// AnnotationList is an envelope for zero to many annotations
type AnnotationList struct {
	Items []Annotation `json:"items,omitempty"` // Items contains 0-many annotations
}

// Define a type for the getTagValue function
type TagValueGetter func(LayerType) string

// defaultGetTagValue retrieves the value associated with the tag field for a given layer.
func defaultGetTagValue(layer LayerType) string {
	switch layer {
	case Application:
		return os.Getenv(TagEnvKey)
	}
	return ""
}

// Global variable for the current tag value getter function
var CurrentTagValueGetter TagValueGetter = defaultGetTagValue

func SetCurrentTagValueGetter(getTag TagValueGetter) {
	CurrentTagValueGetter = getTag
}

// Wrapper function that uses TagValueGetter or falls back to default logic
func GetTagValue(layer LayerType) string {
	// Attempt to use the CurrentTagValueGetter first
	tagValue := CurrentTagValueGetter(layer)
	if tagValue != "" {
		return tagValue
	}

	// Fallback to default logic if CurrentTagValueGetter returned an empty string
	return defaultGetTagValue(layer)
}

/*
The way I see it is providing a way to override tag value logic for specific layers as needed.
This gives the opportunity to use custom logic or even closures that maintain states.
This is especially beneficial for tag values that are expensive to compute or retrieve and need to be fetched only once.

func CustomTagValueGetter() TagValueGetter {
	var cache = make(map[LayerType]string)

	return func(layer LayerType) string {
		switch layer {
		case SPECIFIC_LAYER:
			if val, ok := cache[layer]; ok {
				// Return the cached value if available
				return val
			} else {
				// Fetch and cache the value if not already in cache
				val := // Logic to get the tag value
				cache[layer] = val
				return val
			}
		}
	}
}

When the global variable for the current tag value getter function is changed to the custom one, it'll override the logic for certain layers.
But for other layers where tag value is simple to get like "os" or "app" layers, it'll use the default provided logic.

SetCurrentTagValueGetter(CustomTagValueGetter)
*/

// NewAnnotation is the constructor for an Annotation instance.
func NewAnnotation(key string, hash HashType, host string, layer LayerType, kind AnnotationType, satisfied bool) Annotation {
	return Annotation{
		Id:          NewULID(),
		Key:         key,
		Hash:        hash,
		Host:        host,
		Tag:         GetTagValue(layer),
		Layer:       layer,
		Kind:        kind,
		IsSatisfied: satisfied,
		Timestamp:   time.Now(),
	}
}

func (a *Annotation) UnmarshalJSON(data []byte) (err error) {
	type Alias struct {
		Id          ulid.ULID
		Key         string
		Hash        HashType
		Host        string
		Tag         string
		Layer       LayerType
		Kind        AnnotationType
		Signature   string
		IsSatisfied bool
		Timestamp   time.Time
	}
	x := Alias{}
	// Error with unmarshaling
	if err = json.Unmarshal(data, &x); err != nil {
		return err
	}

	if !x.Hash.Validate() {
		return fmt.Errorf("invalid HashType value provided %s", x.Hash)
	}

	if !x.Kind.Validate() {
		return fmt.Errorf("invalid AnnotationType value provided %s", x.Kind)
	}

	a.Id = x.Id
	a.Key = x.Key
	a.Hash = x.Hash
	a.Host = x.Host
	a.Tag = x.Tag
	a.Layer = x.Layer
	a.Kind = x.Kind
	a.Signature = x.Signature
	a.IsSatisfied = x.IsSatisfied
	a.Timestamp = x.Timestamp
	return nil
}
