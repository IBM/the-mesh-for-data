// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

/*
// RecurrenceType defines how often the data loads should be repeated
// Relevant to implicit and explicit copies, streamed read flows
type RecurrenceType string

// RecurrenceType valid values
const (
	None    RecurrenceType = "none" // load once, and no repeat
	Daily   RecurrenceType = "daily"
	Weekly  RecurrenceType = "weekly"
	Monthly RecurrenceType = "monthly"
	Stream  RecurrenceType = "stream"
)
*/
// CommonDataContext contains the parameters common to both cataloged and uncataloged data
/*
type CommonDataContext struct {
	// IFdetails indicates the protocol and format expected by the data by the Data Scientist's application
	// +required
	IFdetails InterfaceDetails `json:"ifDetails"`

	// Flow indicates if the intent is to read or write data from the user's application
	// If purpose = ingest && Flow = COPY, data should be ingested from the designated data store into the automatically provisioned store in the managed environment
	// (future) If purpose = archive & Flow = COPY, data should be copied from the designated data store and sent to the automatically provisioned archive store
	// If nothing is indicated, then the assumption is it is a read flow.
	// +required
	Flow ModuleFlow `json:"flow"`

	// Recurrence indicates how often the copy should take place or if streamed input is needed
	// +optional
	Recurrence RecurrenceType `json:"recurrence"`
}
*/
/*
// DataResidency indicates the geography in which the data was created
// TODO: Values should come from a taxonomy
type DataResidency string

// DataSensitivity indicates the level of sensitivity of the data in the data set
// TODO: Values should be taken from a taxonomy
type DataSensitivity string

// ExternalDataContext provides the information for reading or writing data that is not cataloged in the data catalog.
// In this case the metadata has to be declared, rather than obtained from the data catalog.
type ExternalDataContext struct {
	// External data source not cataloged in the data catalog
	// +required
	ExternalStore DataStore `json:"externalStore"`

	// Residency indicates the geographical origin of the data, especially important if there is an personal or other sensitive data.
	// Assumes all data in the data set is from the same geography.
	// TODO: In the future perhaps we need to handle situations where there are multiple residencies
	// +required
	Residency DataResidency `json:"residency"`

	// Sensitivity indicates if there is any sensitive data in the data set.
	// Since there may be different types of data with different levels of sensitivity,
	// this is a list.  (ex: confidential, personal, special personal information)
	// No values means there is no sensitive data.
	// +optional
	Sensitivity []DataSensitivity `json:"sensitivity"`

	// CommonContext has additional parameters used for both cataloged and uncataloged data
	// +required
	//CommonContext CommonDataContext `json:"commonContext"`

	// IFdetails indicates the protocol and format expected by the data by the Data Scientist's application
	// +required
	IFdetails InterfaceDetails `json:"ifDetails"`

	// Flow indicates if the intent is to read or write data from the user's application
	// If purpose = ingest && Flow = COPY, data should be ingested from the designated data store into the automatically provisioned store in the managed environment
	// (future) If purpose = archive & Flow = COPY, data should be copied from the designated data store and sent to the automatically provisioned archive store
	// If nothing is indicated, then the assumption is it is a read flow.
	// +required
	Flow ModuleFlow `json:"flow"`

	// Recurrence indicates how often the copy should take place or if streamed input is needed
	// +optional
	Recurrence RecurrenceType `json:"recurrence"`
}

// DataContext indicates data set chosen by the Data Scientist to be read from or written to by his application,
// and includes information about the data format and technologies used by the application
// to access the data.
type DataContext struct {
	// DataSetID is a unique identifier of the dataset chosen from the data catalog for processing by the data user application.
	// +required
	// +kubebuilder:validation:MinLength=1
	DataSetID string `json:"dataSetID"`

	// IFdetails indicates the protocol and format expected by the data by the Data Scientist's application
	// +required
	IFdetails InterfaceDetails `json:"ifDetails"`

	// Flow indicates if the intent is to read or write data from the user's application
	// If purpose = ingest && Flow = COPY, data should be ingested from the designated data store into the automatically provisioned store in the managed environment
	// (future) If purpose = archive & Flow = COPY, data should be copied from the designated data store and sent to the automatically provisioned archive store
	// If nothing is indicated, then the assumption is it is a read flow.
	// +required
	Flow ModuleFlow `json:"flow"`

	// Recurrence indicates how often the copy should take place or if streamed input is needed
	// +optional
	Recurrence RecurrenceType `json:"recurrence"`

	// CommonContext has additional parameters used for both cataloged and uncataloged data
	// +required
	//	CommonContext CommonDataContext `json:"commonContext"`
}

// AppUserRole indicates the role required to use the application
type AppUserRole string

// ApplicationDetails provides information about the Data Scientist's application, which is deployed separately.
// The information provided is used to determine if the data should be altered in any way prior to its use,
// based on policies and rules defined in an external data policy manager.
type ApplicationDetails struct {
	// Purpose indicates the reason for the processing and the use of the data by the Data Scientist's application.
	// For ingest scenario this should be "ingest"
	// (future) For archive scenario this should be "archive"
	// +required
	Purpose string `json:"purpose,omitempty"`

	// ProcessingGeography indicates the state or country or union in which the data processing will take place.
	// This should be the same as the location of the cluster in which the manager is deployed.
	// TO BE OBTAINED FROM THE ENVIRONMENT
	// +optional
	ProcessingGeography string `json:"processingGeography,omitempty"`

	// Role indicates the position held or role filled by the Data Scientist as it relates to the processing of the
	// data he has chosen.
	// TO BE MOVED TO CREDENTIALS
	// +required
	Role AppUserRole `json:"role"`
}

// M4DApplicationSpec defines the desired state of M4DApplication.
type M4DApplicationSpec struct {

	// Selector enables to connect the resource to the application
	// Application labels should match the labels in the selector.
	// +optional
	Selector metav1.LabelSelector `json:"selector"`

	// AppInfo contains information describing the reasons and geography of the processing
	// that will be done by the Data Scientist's application.
	// +required
	AppInfo ApplicationDetails `json:"appInfo"`

	// Data contains the identifiers of the cataloged data to be read or written by the Data Scientist's application,
	// and the protocol used to access it and the format expected.
	// Note: either Data or ExternalData must have at least one entry
	// +optional
	Data []DataContext `json:"data"`

	// ExternalData contains the information about data outside the managed environment that is not cataloged in the data catalog.
	// +optional
	ExternalData []ExternalDataContext `json:"externalData"`
}

// ErrorMessages that are reported to the user
const (
	ReadAccessDenied            string = "Governance policies forbid access to the data"
	CopyNotAllowed              string = "Copy of the data is required but can not be done according to the governance policies."
	ModuleNotFound              string = "No module has been registered"
	InsufficientStorage         string = "No bucket was provisioned for implicit copy"
	InvalidClusterConfiguration string = "Cluster configuration does not support the requirements."
)

// Condition indices are static. Conditions always present in the status.
const (
	FailureConditionIndex int64 = 0
	ErrorConditionIndex   int64 = 1
)

// ConditionType represents a condition type
type ConditionType string

const (
	// ErrorCondition means that an error was encountered during blueprint construction
	ErrorCondition ConditionType = "Error"

	// FailureCondition means that a blueprint could not be constructed
	FailureCondition ConditionType = "Failure"
)

// Condition describes the state of a M4DApplication at a certain point.
type Condition struct {
	// Type of the condition
	Type ConditionType `json:"type"`
	// Status of the condition: true or false
	Status corev1.ConditionStatus `json:"status"`
	// Message contains the details of the current condition
	// +optional
	Message string `json:"message,omitempty"`
}

// ResourceReference contains resource identifier(name, namespace, kind)
type ResourceReference struct {
	// Name of the resource
	Name string `json:"name"`
	// Namespace of the resource
	Namespace string `json:"namespace"`
	// Kind of the resource (Blueprint, Plotter)
	Kind string `json:"kind"`
}

// M4DApplicationStatus defines the observed state of M4DApplication.
type M4DApplicationStatus struct {

	// Ready is true if a blueprint has been successfully orchestrated
	Ready bool `json:"ready,omitempty"`

	// Conditions represent the possible error and failure conditions
	// +optional
	Conditions []Condition `json:"conditions,omitempty"`

	// DataAccessInstructions indicate how the data user or his application may access the data (i.e. virtual endpoint).
	// Instructions are available upon successful orchestration.
	// For data cataloged by the system, this will include the catalog ID
	// +optional
	DataAccessInstructions string `json:"dataAccessInstructions,omitempty"`

	// ObservedGeneration is taken from the M4DApplication metadata.  This is used to determine during reconcile
	// whether reconcile was called because the desired state changed, or whether the Blueprint status changed.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Generated resource identifier
	// +optional
	Generated *ResourceReference `json:"generated,omitempty"`
}

// M4DApplication provides information about the application being used by a Data Scientist,
// the nature of the processing, and the data sets that the Data Scientist has chosen for processing by the application.
// The M4DApplication controller (aka pilot) obtains instructions regarding any governance related changes that must
// be performed on the data, identifies the modules capable of performing such changes, and finally
// generates the Blueprint which defines the secure runtime environment and all the components
// in it.  This runtime environment provides the Data Scientist's application with access to the data requested
// in a secure manner and without having to provide any credentials for the data sets.  The credentials are obtained automatically
// by the manager from an external credential management system, which may or may not be part of a data catalog.
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type M4DApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   M4DApplicationSpec   `json:"spec,omitempty"`
	Status M4DApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// M4DApplicationList contains a list of M4DApplication
type M4DApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []M4DApplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&M4DApplication{}, &M4DApplicationList{})
}
*/