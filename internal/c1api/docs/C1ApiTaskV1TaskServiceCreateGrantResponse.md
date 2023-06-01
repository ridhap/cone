# C1ApiTaskV1TaskServiceCreateGrantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expanded** | Pointer to [**[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner**](C1ApiAppV1AppResourceServiceGetResponseExpandedInner.md) | The expanded field. | [optional] 
**TaskView** | Pointer to [**C1ApiTaskV1TaskView**](C1ApiTaskV1TaskView.md) |  | [optional] 

## Methods

### NewC1ApiTaskV1TaskServiceCreateGrantResponse

`func NewC1ApiTaskV1TaskServiceCreateGrantResponse() *C1ApiTaskV1TaskServiceCreateGrantResponse`

NewC1ApiTaskV1TaskServiceCreateGrantResponse instantiates a new C1ApiTaskV1TaskServiceCreateGrantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiTaskV1TaskServiceCreateGrantResponseWithDefaults

`func NewC1ApiTaskV1TaskServiceCreateGrantResponseWithDefaults() *C1ApiTaskV1TaskServiceCreateGrantResponse`

NewC1ApiTaskV1TaskServiceCreateGrantResponseWithDefaults instantiates a new C1ApiTaskV1TaskServiceCreateGrantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpanded

`func (o *C1ApiTaskV1TaskServiceCreateGrantResponse) GetExpanded() []C1ApiAppV1AppResourceServiceGetResponseExpandedInner`

GetExpanded returns the Expanded field if non-nil, zero value otherwise.

### GetExpandedOk

`func (o *C1ApiTaskV1TaskServiceCreateGrantResponse) GetExpandedOk() (*[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner, bool)`

GetExpandedOk returns a tuple with the Expanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanded

`func (o *C1ApiTaskV1TaskServiceCreateGrantResponse) SetExpanded(v []C1ApiAppV1AppResourceServiceGetResponseExpandedInner)`

SetExpanded sets Expanded field to given value.

### HasExpanded

`func (o *C1ApiTaskV1TaskServiceCreateGrantResponse) HasExpanded() bool`

HasExpanded returns a boolean if a field has been set.

### SetExpandedNil

`func (o *C1ApiTaskV1TaskServiceCreateGrantResponse) SetExpandedNil(b bool)`

 SetExpandedNil sets the value for Expanded to be an explicit nil

### UnsetExpanded
`func (o *C1ApiTaskV1TaskServiceCreateGrantResponse) UnsetExpanded()`

UnsetExpanded ensures that no value is present for Expanded, not even an explicit nil
### GetTaskView

`func (o *C1ApiTaskV1TaskServiceCreateGrantResponse) GetTaskView() C1ApiTaskV1TaskView`

GetTaskView returns the TaskView field if non-nil, zero value otherwise.

### GetTaskViewOk

`func (o *C1ApiTaskV1TaskServiceCreateGrantResponse) GetTaskViewOk() (*C1ApiTaskV1TaskView, bool)`

GetTaskViewOk returns a tuple with the TaskView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskView

`func (o *C1ApiTaskV1TaskServiceCreateGrantResponse) SetTaskView(v C1ApiTaskV1TaskView)`

SetTaskView sets TaskView field to given value.

### HasTaskView

`func (o *C1ApiTaskV1TaskServiceCreateGrantResponse) HasTaskView() bool`

HasTaskView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

