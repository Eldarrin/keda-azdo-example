# keda-azdo-example

This is a sample project to demonstrate how to use KEDA with Azure DevOps Pipelines parents.

The deployment is for the reference container

The scaledjob is the scaler (this could be a scaled object)

The pipeline shows, create the deployment, register to ado, kill deployment, create scaler

Note: AZP_AGENT_NAME is only set in the reference deployment to register in ADO.


In production you would probably use API calls to check the deployment readiness etc, but this is just a demo.