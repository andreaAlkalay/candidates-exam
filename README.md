# SAP CxP Multi Cloud - Candidates Exam

## Prerequisites

1. Install the [Go Programming Language](https://golang.org/dl/)

1. Install [Git](https://git-scm.com/downloads)

1. Install the [Cloud Foundry Command Line Interface (CF-CLI)](https://github.com/cloudfoundry/cli#downloads)

1. [Optional] Install [IntelliJ Idea Community Edition](https://www.jetbrains.com/idea/#chooseYourEdition) or another IDE of your choice

1. Create ~/go/src

1. Set the environment variable GOPATH to ~/go

1. Get required Go dependencies (can take a while)

 ```
 $ go get github.com/cloudfoundry-community/go-cfclient
 $ go get github.com/go-martini/martini
 ```

## Exam

### Preparations
 
1. Get the following from your instructor:

  * URL for the Cloud Foundry API
  * Username and password for Cloud Foundry

1. Fork https://github.com/sapmulticloud/candidates-exam to your own GitHub user

1. Clone the forked repository to your computer

1. Copy the **cloudapps** folder from the GIT repository to **~/go/src**.

1. Go to ~/go/src/cloudapps and run main.go

  `$ go run main.go`
  
### Task 1: List Cloud Foundry apps

* Extend ClientService to retreieve the list of apps from Cloud Foundry

* Implement the **/apps** endpoint to return a JSON containing the list of apps

  The /apps endpoint should return a JSON similar to this example:
  ```
  [
    {
      "app": {
        "name": "firstApp",
        "space_guid": "xxx",
        "memory": 256,
        "instances": 1
      }
    },
    {
      "app": {
        "name": "secondApp",
        "space_guid": "xxx",
        "memory": 256,
        "instances": 1
      }
    }
  ]
  ```
  
  Reference: https://github.com/cloudfoundry-community/go-cfclient

### Task 2: Allow shuffling and sorting of the apps list

* Add the URL parameter **/sort**

  The URL **/apps/sort?shuffle** should returned a shuffled list  
  The URL **/apps/sort?byName** should return a list sorted by name
  
* Bonus: add unit tests for the shuffling and sorting methods

### Task 3: Show a welcome message

* The URL **/** should return a warm welcome message with the current date and time

### Task 4: Deploy the application to Cloud Foundry

* Use the CF CLI to deploy the application to your development space. **Use a new application name**, e.g. cloudapps-guyr.

  Note: You will need to install and use [Godep](https://github.com/tools/godep) to update your Go dependencies before pushing the application
  
  Hint: You only need to install Godep and run `godep save` in your application folder. Note that the godep binary will be installed in **~/go/bin**.

  Reference: [CF CLI](https://docs.cloudfoundry.org/cf-cli/)
  
### Task 5: Submit your test

* Push your code back to your forked repository

* Create a pull request to the original exam repository

  Note: In the pull request, add the URL of your deployed application

## Good luck!
