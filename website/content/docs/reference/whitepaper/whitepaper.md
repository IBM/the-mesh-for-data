title: Delivering Enterprise Compliance for Data Usage - The Mesh for Data

# Delivering Enterprise Compliance for Data Usage - The Mesh for Data

The Mesh for Data is a cloud native platform that connects, secures, controls and observes the use of data --- it orchestrates the services and facilities that involves the use of data, while addressing security and governance. The Mesh for Data ensures that the usage of data will be complaint with the prescribed data usage policies. The Mesh for Data automatically adds governance, security, and data access functionality to the user workload, enabling it to focus on the business goals delegating concerns the non business goals concerns to the mesh.

## The Challenge

Data is critical to derive value and insight for organizations and is a key element for businesses and governments to efficiently operate. However, easily taking advantage of data is challenging due to the need to address security, governance, discovery and lineage tracking, especially when deriving insight in the era of hybrid cloud.
In particular, with today's technology, after access has been granted to a data set, the organization loses control of what is done with the data and it is only via manual processes that it is possible to ensure the data is used only for allowed purposes in allowed locations.

In today’s world, there are three main personas involved in processing data and generating insight in a compliant and governed manner:
* A Data User such as a Data Scientist or analyst who is deriving business value from the data
* An Operator who is responsible for managing the data repositories and the compute infrastructure
* A Data Steward who is responsible for defining and enforcing the policies to protect data assets and ensure compliance with regulations

Each of these personas suffers from a range of difficulties. For instance, the data user often needs to endure a long and manual process to get permission to use the data needed for the task, the Operator needs to manage lots of manual copies as data is moved between clouds in a hybrid setting, and the Data Steward has to ensure compliance with policies in a dynamic and global regulatory environment. As a whole, the organization suffers from the lost opportunity cost due to the time it takes to be able to use data, the risk of reputational impact if data is lost and the risk of fines and reputation damage for non-compliance with regulations.

## The Vision

The Mesh for Data enables the above persona to achieve business goals by focusing on declaring and specifying what should be done and have the Mesh for Data enforce the required policies and regulation to ensure compliance.
Each of the personas interact with native tools to specify how the workload should behave.
* The Data User develops and runs the models, analytics or code using native tools such as notebooks and deploy the workload in containers specifying the workload business goal
* The Operator describe the infrastructure (e.g., cloud resources, data locations, networking) using Kuberentes mechanisms
* The Data Steward specifies policies and controls compliance through data catalogs and policy managers

![Vision](mesh-for-data-vision.png)

The Mesh for Data is a delivered as a combination of control plane orchestration and run-time modules. The Mesh for Data interacts with data catalogs and policy management tools, and deploys run-time modules that handle data access, security and governance in a holistic manner. The Mesh for Data control plane wraps the application workload (i.e., user core business logic) with run-time modules that are responsible for intermediating between the user code and external resource (e.g., data).

## The Concept

The Mesh for Data wraps the user workload by leveraging building blocks from Kubernetes and Istio. It enables running containerized application as-is, but intermediates all ingress and egress flows to and from the application. Internal communication between the application containers are not affected by the Mesh for Data.
The Mesh for Data leverages a declarative specification. provided as M4DApplication CRD, that describe the purpose of the workload and the required external resources (i.e., data, communication) that are needed to run.
Using the above information, the Mesh for Data isolates the application workload and limits communication between the application and the external world only through modules that are deployed by the Mesh for Data.

The Mesh for Data address the following main pillars:
* Connect: provide applications with access to data agnostic to the location of either the computation or the data, including the ability to schedule compute near the data or data mobility as needed by the computation. Moreover, The Mesh for Data transparently supports various application APIs and data source APIs.
* Secure: control access to data and deliver data to the applications taking into account the purpose of the computation, the location and the relevant policies that apply to the use of data. Moreover, control how the output of the computation is handled, communicated and stored.
* Control: orchestrate computation and data mobility based on requirements and policies.
* Observe: collection information on the use of data and provide metrics, logs, audit and data provenance for performance, governance and compliance.

## Architecture

The Mesh for Data is composed of a control plane augmented by connectors that interacts with external services (e.g., data catalog, policy manager, credentials) that act as an orchestration layer that deploys run-time components, called modules, which enforces compliance with the policies and requirements.

### The Control Plane

The Mesh for Data control plane is a set of Kubernetes operators which are collectively named the Manager, supported by registries and connectors.
The M4DApplication CRD is the main driver for the Manager, it describes and registers the application workload with the Mesh for Data. Applying the M4DApplication CRD provides the context for the workload, including the purpose for which it’s running, the data assets that it needs, and a selector (label) to identify the workload. Additional context such as geo-location is extracted from the running platform and resource registries that describe the resources available (e.g., compute and storage resources).

The Manager is responsible for obtaining the details of the data assets (e.g., location, APIs) and relevant policies, which reside in external sources by using connectors. The connectors are configured in the Manager's deployment config-map. All the control plane components runs in the same namespace including connectors. External components can run in their own namespace or location.
The Manager builds a deployment blueprint that isolates the workload from external resources and intermediates external interactions according to the policies defined in the policy manager, based on the context that is provided by M4DApplication and the running environment. The components of the blueprint are assembled based on information from the Module Registry, which describe the available modules resources, their features and function. The Mesh for Data does not manage end-users identity and credentials directly, and delegates this responsibility to the application to implement mechanisms such as end user authentication if required, e.g. using Istio authorization with JWT.

![Control Plane](mesh-for-data-control-plane.png)

The blueprint represents the data flow in the workload, and describes the modules that should be deployed, using Helm charts, and how they should be configured in order to fulfill security, governance, data management, and application requirements.
The modules represent the runtime part of the Mesh for Data. The Manager creates a blueprint for each `M4Dapplication`. The modules that are described in the blueprint are deployed in a separate namespace that is under the control of the Mesh for Data. This namespace is also separate from the application workload namespace.

The Blueprint Controller is responsible for taking the blueprint CRD, which is deployed in the runtime namespace, and deploying the modules using Helm charts thar are provided by each module.

### The Runtime

The Mesh for Data runtime is composed of the user workload and the modules that the blueprint controller deployed. The user workload runs in the namespace that the user has provided, while the components in the blueprint runs in their own namespace which is controlled by the Mesh for Data and is outside the control scope of the user.
When running in the Mesh for Data runtime, the user workload can only access data through the Mesh for Data modules; it no longer directly access data and services, but rather all communication is intermediated by the Mesh for Data modules.

The Modules are responsible to make sure that the usage of data is compliant with the policies defined in the policy manager. The Mesh for Data leverages mechanisms in Kubernetes and Istio to control and limit the application communication with external source. Note that communication between the pods associated with the same workload is not intermediated.

In the figure below we observe that the application container receives data only through ingest data module that are added by the control plane. The ingest module is configured to apply governance according to the policies in the Policy Manager. 
Moreover, the Mesh for Data can automatically add credentials to the data access request outside the control scope of the user and thus bypassing the need to share credentials with the user and remove the possibility of bypassing governance when accessing the data.
By have sole control over the distribution of credentials and enforcing access through the access modules, the Mesh for Data can apply governance policies, and ensure that consideration such as location of the data, the data residency and the location of the workload are take into account.

In the same way, the Mesh for Data can control the input of the application, and for example, ensure that the application can only write its results to explicitly allowed locations. Modules can also inspect and if needed modify the results in the same way that the inputs can be transformed. For instance, search for credit card numbers in the output, and ensure that the application does not export credit card numbers.

![The Runtime](mesh-for-data-runtime.png)

Since the runtime is built on top of Kubernetes, it can be deployed anywhere a client has available OpenShift clusters. And in particular, our vision is to make it  possible to deploy a single mesh in a hybrid environment, such that some of the mesh's containers run on one cluster and other containers run on a different cluster. This will support both hybrid and multi cloud deployments.

## Architecture - Key Components

### Control plane components

* **Mesh for Data Control Plane:** The Mesh for Data control plane is a component composed of a collection of Kubernetes controllers, connectors and registries that is used to define the way an application running on the Mesh for Data should be deployed, based upon inputs about the application to execute (provided in the `M4Dapplication`) and information retrieved from an external policy manager and data catalog. The control plane has three built-in controllers: application description (`M4Dapplication`) controller, blueprint controller, and implicit Copy Controller.

* **`M4DApplication` Controller:** Part of the Mesh for Data manager, responsible for compiling a blueprint that describe compliant data flows for the application workload. It processes the Mesh for Data application YAML (`M4DApplication`) and retrieves a set of modules to implement all governance and management actions based upon defined policies retrieved by the Policy Compiler component. Governance actions are re-generated upon any policy change. The controller retrieves information about the resources and modules that it can use. Specifically, it searches for modules that can satisfy the workload requirements (e.g., governance enforcement actions) taking into account performance considerations, the use of specific client SDKs and more. The controller generates a blueprint describing the entire data path, including the components injected into the data ingress and egress paths of the application.

* **Blueprint Controller:** Part of the Mesh for Data manager, a Kubernetes  orchestrator that processes the blueprint and deploys all runtime components accordingly.  

* **Implicit Copy Controller:** Part of the Mesh for Data manager, responsible for creating an implicit copy of a data set from one data store to another, potentially applying transformations dictated by enforcement actions. The implicit copy is a transient copy that is cleaned up when the application completes execution.  This module will be inserted into the blueprint by the `M4DApplication` controller depending upon the location of the data set, the location where the application will run and governance requirements.

* **Module Registry:** Describes the list of modules available to be deployed as the Mesh for Data runtime. The registry lists for each module its type, features, and supported APIs. Currently the module registry is built as a set of `M4DModule` CRDs

* **Compute and Storage Resource Registries:** The registries that are used by the Mesh for Data to understand the available infrastructure to the workload. These registries are currently conceptual and are not yet properly implemented.

* **Connector:** A control component configured through the Manager configmap, typically part of the control plane, that interacts with external resources in order to control the runtime behavior. A connector is a gRPC service that is deployed with the Mesh for Data, typically as part of the Mesh for Data namespace. Three connector types are supported: policy manager for keeping policies, data catalog that describes the data assets and credential manager that keeps data assets credentials.

* **Policy Manager Connector:** A connector service that interacts with an external policy manager and is responsible to communicate the enforcement actions that the Mesh for Data needs to apply on the application workloads. Enforcing data governance policies requires a Policy Manager or a Policy Decision Point (PDP) that dictates what enforcement actions need to take place. The Mesh for Data supports a wide and extensible set of enforcement actions to perform on data read, write or copy. These include transformation of data, verification of the data, and various restrictions on the external activity of an application that can access the data. A PDP returns a list of enforcement actions given a set of policies and specific context about the application and the data it uses. 
The Mesh for Data open source distribution includes an example PDP that is powered by Open Policy Agent (OPA). A Policy Compiler component handles combining lists of enforcement actions from multiple policy manager connectors, if configured.

* **Data Catalog Connector:** The Mesh for Data assumes the use of a data catalog which is responsible for keeping track of data in various store and locations. That is, keep technical details for data asset that are referenced from `M4DApplication` CRD, which include a link to the asset in the catalog. The catalog provides metadata about the asset such as tags and connection information that describes how to connect to the data source for accessing the data. The Mesh for Data uses the metadata provided by the catalog both to enable seamless connectivity to the data and as input for policy decisions. 
This enables the data user to transparently access the data with no regard for where the data resides.
The Mesh for Data does not include a built-it data catalog, but rather connects to an external catalog.

* **Credentials Manager Connector:** The Mesh for Data assumes that data access credentials of cataloged assets are securely stored in an enterprise credentials manager and are referenced from the data catalog. The Mesh for Data links to existing credential managers using connectors for reading data access credentials. The Mesh for Data contains an example connector for accessing credentials in HashiCorp Vault.

### Control plane CRDs

* **Mesh for Data Application:** This is a YAML (`M4Dapplication`) description of the application workload that will execute on Kubernetes and be wrapped by the Mesh for Data. It provides the context about the application such as the purpose for which it’s running, the data assets that it needs, and a label selector to identify the workload.  Example workloads are Watson Studio notebook, a Data Stage Job, etc., or post v1 it could be a bespoke application.

* **Blueprint:** The manager builds a blueprint for each `M4Dapplication` applied and processed by the control plane. The blueprint describes the information flows and modules for the workload, and includes services, jobs, and configurations that are injected into the data path by the control plane to fulfill security, governance, data management, and application requirements.  


### Runtime components

* **Mesh for Data Runtime:** The run-time is composed of the user workload and modules instances (derived from `M4DModule` CRDs) that intermediate between the application workload and the external world (e.g., data and services). The run-time is deployed according to the blueprint, which is generated by the `M4DApplication` Controller.

* **Mesh for Data Module:**  Modules are the run-time components of the Mesh for Data and allow controlling and intermediating the flow of information for the application workload. Modules can be open or proprietary and are defined in a registry. Modules are the way to package function such as side-cars, proxies, jobs, and configurations and make them available to the control plane to choose from, and to be deployed during run-time. 

* **JWT Token Injector Module:** The Mesh for Data includes a module that knows how to add a JWT token into a REST request that an application sends. The token injector retrieves the proper token from the credential connector and automatically adds it to the proper REST requests.

* **Isolation and Gateway Module:** A Kubernetes controller responsible for controlling the external interactions of the application workload. All communication from application workload pods to non-workload pods is controlled by this module. All traffic to external endpoints is routed through a gateway that can control and intermediate those flows. 

* **Apache Arrow Flight Data Access Module:** A Mesh for Data module that exposes an Arrow/Flight interface and is able to apply transformations on data accessed through the module.  This module is used to support governance on data that is directly accessed without an implicit copy.

# Glossary

* CRD: Custom resource definition
