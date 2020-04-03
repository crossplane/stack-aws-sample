# AWS Sample Stack

You can use this stack to spin up a private network as well as resource classes that will let you provision resources in that network.

# Installation

Requirements:
* Crossplane should be installed.
* [AWS Provider](https://github.com/crossplaneio/provider-aws) should be installed and its version should be at least 0.5.0

If you have crossplane-cli installed, you can use the following command to install:

```bash
# Do not forget to change <version> with the correct version.
kubectl crossplane stack install --cluster -n crossplane-system 'crossplane/stack-aws-sample:<version>' aws-sample
```

If you don't have crossplane-cli installed, you need to create the following YAML to install:

```yaml
apiVersion: stacks.crossplane.io/v1alpha1
kind: ClusterStackInstall
metadata:
  name: "aws-sample"
  namespace: crossplane-system
spec:
  package: "crossplane/stack-aws-sample:<version>"
```

# Usage Instructions

You can create the following YAML to trigger creation of:
* [`VPC`][vpc]
* [`Subnet`][subnet]
* [`Security Group`][securitygroup]
* [`Route Table`][routetable] 
* [`Internet Gateway`][internetgateway] 
* [`IAM Role`][iamrole] 
* [`IAM Role Policy Attachment`][iamrolepolicyattachment] 
* [`DB Subnet Group`][dbsubnetgroup]
* [`Provider`][provider] that points to credentials secret reference you supply

and the following resource classes with minimal hardware requirements that will let you create instances that are connected to that network.

* [`EKSClusterClass`][ekscluster-class]
* [`RDSInstanceClass`][rdsinstance-class]
* [`ReplicationGroupClass`][replicationgroup-class]


```yaml
apiVersion: aws.resourcepacks.crossplane.io/v1alpha1
kind: AWSSample
metadata:
  name: test
spec:
  region: us-west-2
  projectID: crossplane-playground
  credentialsSecretRef:
    name: aws-account-creds
    namespace: crossplane-system
    key: credentials
```

In Crossplane, the resource classes that are annotated with `resourceclass.crossplane.io/is-default-class: "true"` are used as default if the claim doesn't specify a resource class selector. The resource classes you create via the `AWSSample` instance above will deploy all of its resource classes as default. If you'd like those defaulting annotations to be removed, you need to add the following to `AWSSample` instance above:

```yaml
templatestacks.crossplane.io/remove-defaulting-annotations: true
```

## Build

Run `make`

## Test Locally

### Minikube

Run `make` and then run the following command to copy the image into your minikube node's image registry:

```bash
# Do not forget to specify <version>
docker save "crossplane/stack-aws-sample:<version>" | (eval "$(minikube docker-env --shell bash)" && docker load)
```

After running this, you can use the [installation](#installation) command and the image loaded into minikube node will be picked up. 

[vpc]: kustomize/aws/network/vpc.yaml
[subnet]: kustomize/aws/network/subnet.yaml
[securitygroup]: kustomize/aws/network/securitygroup.yaml
[routetable]: kustomize/aws/network/routetable.yaml
[internetgateway]: kustomize/aws/network/internetgateway.yaml
[iamrole]: kustomize/aws/identity/iamrole.yaml
[iamrolepolicyattachment]: kustomize/aws/identity/iamrolepolicyattachment.yaml
[dbsubnetgroup]: kustomize/aws/database/dbsubnetgroup.yaml
[provider]: kustomize/aws/provider.yaml
[ekscluster-class]: kustomize/aws/compute/eksclusterclass.yaml
[rdsinstance-class]: kustomize/aws/database/rdsinstanceclass.yaml
[replicationgroup-class]: kustomize/aws/cache/replicationgroupclass.yaml