// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    listReturns, err := UnmarshalListReturns(bytes)
//    bytes, err = listReturns.Marshal()

package list

import "encoding/json"

func UnmarshalListReturns(data []byte) (ListReturns, error) {
	var r ListReturns
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListReturns) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ListReturns struct {
	Expand          *string                `json:"expand,omitempty"`    
	Issues          []Issue                `json:"issues"`              
	MaxResults      *int64                 `json:"maxResults,omitempty"`
	Names           map[string]interface{} `json:"names,omitempty"`     
	Schema          map[string]interface{} `json:"schema,omitempty"`    
	StartAt         *int64                 `json:"startAt,omitempty"`   
	Total           *int64                 `json:"total,omitempty"`     
	WarningMessages []string               `json:"warningMessages"`     
}

type Issue struct {
	Changelog                *Changelog             `json:"changelog,omitempty"`               
	Editmeta                 *EditMeta              `json:"editmeta,omitempty"`                
	Expand                   *string                `json:"expand,omitempty"`                  
	Fields                   map[string]interface{} `json:"fields,omitempty"`                  
	FieldsToInclude          map[string]interface{} `json:"fieldsToInclude,omitempty"`         
	ID                       *string                `json:"id,omitempty"`                      
	Key                      *string                `json:"key,omitempty"`                     
	Names                    map[string]interface{} `json:"names,omitempty"`                   
	Operations               *Opsbar                `json:"operations,omitempty"`              
	Properties               *Properties            `json:"properties,omitempty"`              
	RenderedFields           map[string]interface{} `json:"renderedFields,omitempty"`          
	Schema                   map[string]interface{} `json:"schema,omitempty"`                  
	Self                     *string                `json:"self,omitempty"`                    
	Transitions              []Transition           `json:"transitions"`                       
	VersionedRepresentations map[string]interface{} `json:"versionedRepresentations,omitempty"`
}

type Changelog struct {
	Histories  []ChangeHistory `json:"histories"`           
	MaxResults *int64          `json:"maxResults,omitempty"`
	StartAt    *int64          `json:"startAt,omitempty"`   
	Total      *int64          `json:"total,omitempty"`     
}

type ChangeHistory struct {
	Author          *User            `json:"author,omitempty"`         
	Created         *string          `json:"created,omitempty"`        
	HistoryMetadata *HistoryMetadata `json:"historyMetadata,omitempty"`
	ID              *string          `json:"id,omitempty"`             
	Items           []ChangeItem     `json:"items"`                    
}

type User struct {
	AccountID    *string                `json:"accountId,omitempty"`   
	Active       *bool                  `json:"active,omitempty"`      
	AvatarUrls   map[string]interface{} `json:"avatarUrls,omitempty"`  
	DisplayName  *string                `json:"displayName,omitempty"` 
	EmailAddress *string                `json:"emailAddress,omitempty"`
	Key          *string                `json:"key,omitempty"`         
	Name         *string                `json:"name,omitempty"`        
	Self         *string                `json:"self,omitempty"`        
	TimeZone     *string                `json:"timeZone,omitempty"`    
}

type HistoryMetadata struct {
	ActivityDescription    *string                `json:"activityDescription,omitempty"`   
	ActivityDescriptionKey *string                `json:"activityDescriptionKey,omitempty"`
	Actor                  *ActorClass            `json:"actor,omitempty"`                 
	Cause                  *CauseClass            `json:"cause,omitempty"`                 
	Description            *string                `json:"description,omitempty"`           
	DescriptionKey         *string                `json:"descriptionKey,omitempty"`        
	EmailDescription       *string                `json:"emailDescription,omitempty"`      
	EmailDescriptionKey    *string                `json:"emailDescriptionKey,omitempty"`   
	ExtraData              map[string]interface{} `json:"extraData,omitempty"`             
	Generator              *GeneratorClass        `json:"generator,omitempty"`             
	Type                   *string                `json:"type,omitempty"`                  
}

type ActorClass struct {
	AvatarURL      *string `json:"avatarUrl,omitempty"`     
	DisplayName    *string `json:"displayName,omitempty"`   
	DisplayNameKey *string `json:"displayNameKey,omitempty"`
	ID             *string `json:"id,omitempty"`            
	Type           *string `json:"type,omitempty"`          
	URL            *string `json:"url,omitempty"`           
}

type CauseClass struct {
	AvatarURL      *string `json:"avatarUrl,omitempty"`     
	DisplayName    *string `json:"displayName,omitempty"`   
	DisplayNameKey *string `json:"displayNameKey,omitempty"`
	ID             *string `json:"id,omitempty"`            
	Type           *string `json:"type,omitempty"`          
	URL            *string `json:"url,omitempty"`           
}

type GeneratorClass struct {
	AvatarURL      *string `json:"avatarUrl,omitempty"`     
	DisplayName    *string `json:"displayName,omitempty"`   
	DisplayNameKey *string `json:"displayNameKey,omitempty"`
	ID             *string `json:"id,omitempty"`            
	Type           *string `json:"type,omitempty"`          
	URL            *string `json:"url,omitempty"`           
}

type ChangeItem struct {
	Field      *string `json:"field,omitempty"`     
	FieldID    *string `json:"fieldId,omitempty"`   
	Fieldtype  *string `json:"fieldtype,omitempty"` 
	From       *string `json:"from,omitempty"`      
	FromString *string `json:"fromString,omitempty"`
	To         *string `json:"to,omitempty"`        
	ToString   *string `json:"toString,omitempty"`  
}

type EditMeta struct {
	Fields map[string]interface{} `json:"fields,omitempty"`
}

type Opsbar struct {
	LinkGroups []LinkGroupElement `json:"linkGroups"`
}

type LinkGroupElement struct {
	Groups     []GroupElement `json:"groups"`              
	Header     *SimpleLink    `json:"header,omitempty"`    
	ID         *string        `json:"id,omitempty"`        
	Links      []SimpleLink   `json:"links"`               
	StyleClass *string        `json:"styleClass,omitempty"`
	Weight     *int64         `json:"weight,omitempty"`    
}

type GroupElement struct {
	Groups     []GroupElement `json:"groups"`              
	Header     *SimpleLink    `json:"header,omitempty"`    
	ID         *string        `json:"id,omitempty"`        
	Links      []SimpleLink   `json:"links"`               
	StyleClass *string        `json:"styleClass,omitempty"`
	Weight     *int64         `json:"weight,omitempty"`    
}

type SimpleLink struct {
	Href       *string `json:"href,omitempty"`      
	IconClass  *string `json:"iconClass,omitempty"` 
	ID         *string `json:"id,omitempty"`        
	Label      *string `json:"label,omitempty"`     
	StyleClass *string `json:"styleClass,omitempty"`
	Title      *string `json:"title,omitempty"`     
	Weight     *int64  `json:"weight,omitempty"`    
}

type Properties struct {
	Properties map[string]interface{} `json:"properties,omitempty"`
}

type Transition struct {
	Expand    *string                `json:"expand,omitempty"`   
	Fields    map[string]interface{} `json:"fields,omitempty"`   
	HasScreen *bool                  `json:"hasScreen,omitempty"`
	ID        *string                `json:"id,omitempty"`       
	Name      *string                `json:"name,omitempty"`     
	To        *Status                `json:"to,omitempty"`       
}

type Status struct {
	Description    *string         `json:"description,omitempty"`   
	IconURL        *string         `json:"iconUrl,omitempty"`       
	ID             *string         `json:"id,omitempty"`            
	Name           *string         `json:"name,omitempty"`          
	Self           *string         `json:"self,omitempty"`          
	StatusCategory *StatusCategory `json:"statusCategory,omitempty"`
	StatusColor    *string         `json:"statusColor,omitempty"`   
}

type StatusCategory struct {
	ColorName *string `json:"colorName,omitempty"`
	ID        *int64  `json:"id,omitempty"`       
	Key       *string `json:"key,omitempty"`      
	Name      *string `json:"name,omitempty"`     
	Self      *string `json:"self,omitempty"`     
}
