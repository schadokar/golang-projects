# go-to-do-app

This application is originally developed by [Shubham Chadokar](https://github.com/schadokar)

The Complete tutorial is available at [Medium](https://levelup.gitconnected.com/build-a-todo-app-in-golang-mongodb-and-react-e1357b4690a6)

#### Note:

This project is developed inside the `GOPATH` directory. The imports are non local packages means whatever you are importing you have to provide its complete path after the `$GOPATH/src`

##### For ex.

if the `GOPATH` is `C://users/go`
Your package path: `C://users/go/src/github.com/schadokar/practice-stuff/to-do-app/golang/implementation-1/server/src/pkg/models`
To import `models` package in your project, you have to import it as

```
"github.com/schadokar/practice-stuff/to-do-app/golang/implementation-1/server/src/pkg/models"
```

the complete path after the `src`

#### Note:

The directory is changed with respect to the original application.
The `middleware models router` packages are moved to `src/pkg` folder

##### Note: Checkout its offline version at http://getshitdone.surge.sh . To check the source code, switch to offlineToDoApp branch.

This is a to-do list application.

**Server: Golang  
Client: React, semantic-ui-react  
Database: Local MongoDB**

# Application Requirement

### golang server requirement

1. golang https://golang.org/dl/
2. gorilla/mux library for router `go get -u github.com/gorilla/mux`
3. mongo-driver library to connect with mongoDB `go get go.mongodb.org/mongo-driver`

### react client

From the Application directory

`create-react-app client`

# Start the application

1. Make sure your mongoDB is started
2. From server directory, open a terminal and run
   `go run main.go`
3. From client directory,  
   a. install all the dependencies using `npm install`  
   b. start client `npm start`

# Walk through the application

Open application at http://localhost:3000

### Index page

![](https://github.com/schadokar/go-to-do-app/blob/master/images/index.PNG)

### Create task

Enter a task and Enter

![](https://github.com/schadokar/go-to-do-app/blob/master/images/createTask.PNG)

### Task Complete

On completion of a task, click "done" Icon of the respective task card.

![](https://github.com/schadokar/go-to-do-app/blob/master/images/taskComplete.PNG)

You'll notice on completion of task, card's bottom line color changed from yellow to green.

### Undo a task

To undone a task, click on "undo" Icon,

![](https://github.com/schadokar/go-to-do-app/blob/master/images/createTask.PNG)

You'll notice on completion of task, card's bottom line color changed from green to yellow.

### Delete a task

To delete a task, click on "delete" Icon.

![](https://github.com/schadokar/go-to-do-app/blob/master/images/deletetask.PNG)

# References

https://godoc.org/go.mongodb.org/mongo-driver/mongo  
https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial  
https://vkt.sh/go-mongodb-driver-cookbook/

# License

MIT License

Copyright (c) 2019 Shubham Chadokar
