# Sprint 4 Final product Completed

# Final Project demo 
[![demo](https://i9.ytimg.com/vi/j_HJzqQ3_Tc/mq1.jpg?sqp=CLSYg5MG&rs=AOn4CLByKro9-cvzoM3Lq-xOUxaSbQ4s3w)](https://www.youtube.com/embed/j_HJzqQ3_Tc)</br>
link: https://youtu.be/j_HJzqQ3_Tc



# Sprint 4 Final product Completed

# Sprint 4 Final step of the web app, FrontEnd added "Post" page, and combined all functionalities with backend 
Contributor: Yuhao Shi\
Framework: react\
Files and folders added or modified under "Frontend" folder.\
How to run: 1. Frontend: run "npm install" for react packages and "npm start" to start the application; 2. run "go build main" and "go run main.go" for the two backend under folder "backend" and "backend-Chat" seperatly. 3. Before running, install module "react-twitter-embed" and "react-chat-engine".

Funtionalities for this sprint: Added "post" page where users can post their tweets on the web for others to view. Combined frontend and backend with functionalities including login/register, post and view tweets, and real-time chatting between multiple users. Users can now register and login with the same username and password, post any texts on the web page for all other users to see, and have real-time public chatting with any user who joined the session while seeing their login and logout status. Also added additional feature: a simple embed Twitter Widget with tweet feed directly from UF's Twitter account.







# Sprint 3 Frontend Real Time Public Chat with mutiple clients
Contributor: Yuhao Shi\
Framework: react\
Temporary backend: Go\
Files and folders for this functionality: "Frontend" & "backend-Chat" under "main-branch"\
How to run: run "npm install" for react packages and "npm start" to start the application; run "go build main" for the temporary backend. (Also Showcased in the demo video).\
Funtionalities for this Sprint: Real-time Public Chat between multiple users was implemented, users will be able to send and read text massages in real time under "Messages" page. Everyone who joined the local host address "8080" will be able to see the chat history. Notifications will also be seen when new users joined or left the channel. The Frontend was modified so that it will send message data and recieve data from WebSockets. To implement these, I wrote a seperate backend under "backend-Chat" folder because our previous backend was not really capable of combining it with the frontend, so another backend needs to be written from scratch for the real time chatting to work. We will modify the backend even more in the future for this functionality. (To see exactly how it works, a demo video is provided as "Sprint3-Frontend Chat.mp4").


# Sprint 3 Frontend Support Chat and Support Page
Contributor: Fugang Deng\
Framework: react\
Package: axios, cypress , cypress testing library, ant-design, react-chat-engine\
Function: At the bottom-right of the website, there is a chatbot image surronding with bule color, once you move the mouse over, it will show "Support Chat Here!!!" and the bule side will disappear. Then, click that, there will be the UI show up to let you enter the email. Once, you enter the email, the chat window will show up and you can chat.\
Testing: The cypress testing require npm install cypress and npm install --save-dev cypress @testing-library/cypress, and put import '@testing-library/cypress/add-commands' in the cypress/support/command.js. I changed to 'test' in package.json to 'cypress open'. Once you do npm start, then open a new terminal type npm test. The cypress interface will show up and double the homepage.js, the test will begin. All the test codes are in the homepage.js in cypress/intergartion. You will see an error on the second test because we need backend to send something to do the fetch, other than that, all the test are good. For, the new functionality I did not write the code for the purpose of testing, so cypress cannot find the exact location of to hover and click, so more test details will show in the Sprint3SupportChat vedio. \
HowToRun: First, you should only put the files in the FrontEnd in one project. Then install these 3 packages as following:\
          npm i axios\
          npm install react-chat-engine --save\
          npm i @ant-design/icons\
          Finally, put npm run in the terminal.\
This specific functionality is placed in the branch "main-branch/FrontEnd"
# Sprint 3 Backend
Contributor: Dafei Du\
Framework: Gin\
Folders:\
backend/sprint3
Features:\
In this sprint, users can create posts, list posts, update posts and delete posts. Add more unit tests for this functionality. 
1. create posts: users need input their name and some content to create a post
2. list posts: users can view all posts, or search exactly one post by searching post id
3. update posts: users can edits their posts
4. delete posts: users can delete their posts

The previous functionalities are also optimized. 
How To Run: go run main.go (to see more details, please check sprint3-backend-demo.mov in backend/sprint3 folder.)

Contributor: Yuhan Jin\
Framework: Gin\
Features:

1.Modify the function of sprint2, follow the GIN framework in sprint1\

2.add the function of post in sprint3,Realize the user's dynamic update function (CURD)\

3.add Friends list viewing , adding new friend functions (only part of the code has been completed)
# Sprint 2 Backend
Contributor: Yuhan Jin\
Framework: Gin\
Features:The backend implements the function of real-time communication, creating a chat room that includes the following functions：

1.Broadcast function of users going online :When a user goes online, it will broadcast to all online users that they are online.

2.Online users' query：We can obtain the currently online users list through the query command. 

3.Force users to go offline when timeout： When the user does not do any operation within the specified time, the program will force the users to go offline. 

4.private and public chat：The private chat function is a point-to-point message transmission; only the two parties can see the content, while the public chat is visible to everyone.

# Sprint 2 Backend
Contributor: Dafei Du\
Framework: Gin\
Features: public chat, private chat, update username, log out, API documentation\
Functions: 

1. public chat: users can chat to everyone. Any online users can receive the message.

2. private chat: users can choose any online user to chat. They can send messages to each other. Others cannot see the messages.

3. update username: Users do not have their name at the beginning. They should update their username. After done that, they have unique name. This means other users cannot use the same username.

4. logout: If users do not want to chat, they can log out the account the quit the chat room. The status of the1. user will changed to inactive. 

5. In addition, I built a client to facilitate operation and combined with sprint1 and sprint2 together. 

6. API documentatio is implemented. It contains in backend folder called api.yaml. It can be viewed in swagger.io.


# Sprint 2 Frontend UI design & combine frontend and backend
Contributor: Yuhao Shi\
What's working: Based on the first draft, a chat page was added to display real time message textting. After navigating the page to "message" in taskbar, there are three components on the page including a friends list on the left, real time chatting box and a message input box in the middle, and a list of current online friends on the right. 
I also combined frontend and backend successfully by posting and retriving data between server and database through our local host API. Users now are able to register their account and login with their registered email and password. When clicking on the login button, users will be redirected to a page where a message will be displayed based on whether they are authenticated or not. 

Everything we worked on for Frontend were placed under "main-branch" and folder "FrontEnd".

IMPORTANT: as of right now, it seems like we are not able to access API and retrieve data from MySQL database on WINDOWS machines. But everything works perfectly fine on MacOS. It might seems to be an unsolved issue from the database. We might have to change the databse to something else in future sprints. 

# Sprint 2 Frontend Taskbar and Cypress Unit Testing
Contributor: Fugang Deng\
Framework: react\
Package: react-icon, react-router-dom, cypress , cypress testing library, bootstrap, react-bootstrap\
Function: At the top-left of the website, there is a three white bars, once you click that, the taskbar will show, show websites just simple titles, still need future implementation. I also help to combine the UI of Home, Login, Register and HomePost of my task bar.\
Testing: The cypress testing require npm install cypress and npm install --save-dev cypress @testing-library/cypress, and put import '@testing-library/cypress/add-commands' in the cypress/support/command.js. I changed to 'test' in package.json to 'cypress open'. Once you do npm start, then open a new terminal type npm test. The cypress interface will show up and double the homepage.js, the test will begin. All the test codes are in the homepage.js in cypress/intergartion. You will see an error on the second test because we need backend to send something to do the fetch, other than that, all the test are good. \
Components: Taskbar\
This specific functionality is placed in the branch "main-branch/FrontEnd"

# Frontend UI design
Contributor: Yuhao Shi\
What's working: This first draft implementation includes UI design and layout for Home Page, Login Page, Register Page, a global "taskbar" on top to redirect between pages, and a draft page after login. Users can sign up for a new account and then use the credentials to login. Each page is linked through React Route, however, the database that stores user credentials and server API that transfer information between registered users and login still need more work to collaborate with the backend.

The structure uses React Router to navigate between pages.

The implementation is placed in the branch "frontend #1".

Demo for the UI is uploaded in the branch seperately

# Sprint 1 Backend Tasktracker
Programming Language: Golang\
Framework: Gin\
Database: mySQL\
Package: jwt-go, cors, gin-gonic/gin, go-sql-driver/mysql, golang.org/x/crypto/bcrypt\
Function: regitser, login, logout\
Test:Postman

# Backend 
What's working: User authentication (JWT), Login, and Logout functions are implemented, and MySQL is used as the database based on the GIN framework. The next step will focus on the user's information sections: following, followers, Avatar, Comment，Post，Repost.

The implementation is placed in the branch "Sprint1-backend".

# Backend Code Contributions：
Dafei Du: datasource,  domain, serveices, main.go, video\
Yuhan Jin：app, controller, utils 


# Getting Started with Create React App

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in your browser.

The page will reload when you make changes.\
You may also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.\
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `npm run eject`

**Note: this is a one-way operation. Once you `eject`, you can't go back!**

If you aren't satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you're on your own.

You don't have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn't feel obligated to use this feature. However we understand that this tool wouldn't be useful if you couldn't customize it when you are ready for it.

## Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

To learn React, check out the [React documentation](https://reactjs.org/).

### Code Splitting

This section has moved here: [https://facebook.github.io/create-react-app/docs/code-splitting](https://facebook.github.io/create-react-app/docs/code-splitting)

### Analyzing the Bundle Size

This section has moved here: [https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size](https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size)

### Making a Progressive Web App

This section has moved here: [https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app](https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app)

### Advanced Configuration

This section has moved here: [https://facebook.github.io/create-react-app/docs/advanced-configuration](https://facebook.github.io/create-react-app/docs/advanced-configuration)

### Deployment

This section has moved here: [https://facebook.github.io/create-react-app/docs/deployment](https://facebook.github.io/create-react-app/docs/deployment)

### `npm run build` fails to minify

This section has moved here: [https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify](https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify)
