# Implementation 1

This project is originally developed by [Elliot Forbes](https://github.com/elliotforbes)

The tutorial for this project is published on [Tutorial Edge](https://tutorialedge.net/projects/chat-system-in-go-and-react/)

### Note

There are couple of things which you have to take care of

#### In Part 2

The api/index.js it should be in src directory

#### In Part 3

Add chatHistory state in App.js

```
class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: []
    };
  }
```
