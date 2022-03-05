This time we combine the frontend and backend together

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
