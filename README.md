
This time we combinde the frontend and backend together, and furthur development in more pages and their functionalities. Everything we worked on this time were ALL placed under "main-branch" default branch with clear history of committments. 

Everything for backend was placed under "BackEnd" folder, and everything for frontend was under "FrontEnd" folder.

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
Features: public chat, private chat, update username, log out\
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
